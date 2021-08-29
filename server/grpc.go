package server

import (
	"github.com/niccoloCastelli/orderbooks/config"
	orderbooks "github.com/niccoloCastelli/orderbooks/server/proto"
	"github.com/niccoloCastelli/orderbooks/utils"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var (
	grpcServer *grpc.Server
)

func NewGrpcServer(logger *zerolog.Logger, conf config.ServerConfig) error {
	lis, err := net.Listen("tcp", conf.Host)
	if err != nil {
		return err
	}
	logger.Info().Msg("server initialization...")
	grpcServer = grpc.NewServer()
	srv, err := newOrderbooksGrpcServer(logger, conf)
	if err != nil {
		return err
	}

	orderbooks.RegisterOrderBooksServer(grpcServer, srv)
	reflection.Register(grpcServer)
	errChan := make(chan error)
	exitChan := utils.MakeExitChan()
	go func(errChan chan error, lis net.Listener, grpcServer *grpc.Server) {
		if err := grpcServer.Serve(lis); err != nil {
			logger.Err(err).Send()
			errChan <- err
		}
	}(errChan, lis, grpcServer)
	logger.Info().Msgf("GRPC server on: %s", conf.Host)
	for {
		select {
		case err := <-errChan:
			logger.Error().Err(err).Msg("Grpc error")
			return err
		case <-exitChan:
			logger.Warn().Msg("grpc closed")
			return nil
		}
	}

}
