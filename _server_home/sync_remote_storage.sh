#!/usr/bin/env bash

export HOME="/home/nicco"
#export KEY_FILE="${HOME}/certificati/hetzner/orderbooks/ubuntu/private_ubuntu.key"
export KEY_FILE="${HOME}/certificati/hetzner/orderbooks/new/ubuntu/private_ubuntu.key"
export REMOTE_PATH='/mnt/data/orderbooks'

#export LOCAL_BASEPATH="${GOPATH}/src/bitbucket.org/niccolo_castelli/orderbooks/data/"
export LOCAL_BASEPATH="/media/nicco/DATA_SSD/order_books/"

#export IP="95.216.144.221" #OLD
export IP="192.168.10.3"

mkdir -p ${LOCAL_BASEPATH}

rsync -chavzP -e "ssh -i $KEY_FILE" --exclude="*.csv" nicco@${IP}:"/mnt/data/orderbooks/" "${LOCAL_BASEPATH}"
