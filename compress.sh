#!/usr/bin/env bash
PROJECT_PATH="/home/nicco/go/src/github.com/niccoloCastelli/orderbooks"
cd $PROJECT_PATH
./orderbooks --log_level 1 --config config_gzip.json compress 0