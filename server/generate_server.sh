#!/usr/bin/env bash

WD=$PWD
OUT_PATH="$GOPATH/src/github.com/niccoloCastelli/orderbooks/server/proto/"
PROTO_PATH="$GOPATH/src/github.com/niccoloCastelli/orderbooks/server/proto"
mkdir -p $OUT_PATH/js



#

cd $PROTO_PATH
protoc orderbooks.proto --gogofast_out=plugins=grpc,Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types:$OUT_PATH \
--proto_path=$HOME/protoc/include \
--proto_path=${GOPATH}/src \
--proto_path=.
go fmt $OUT_PATH/websocket.go

python3 -m grpc_tools.protoc -I. --python_out=./python --grpc_python_out=./python --mypy_out=./python orderbooks.proto


cd $OUT_PATH
ls *.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
cd $WD
