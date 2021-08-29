#!/usr/bin/env bash

#export KEY_FILE="${HOME}/certificati/hetzner/orderbooks/ubuntu/private_ubuntu.key"
export KEY_FILE="${HOME}/certificati/hetzner/orderbooks/new/ubuntu/private_ubuntu.key"
#export KEY_FILE="'/home/nicco/Script utility/sshuser/tmp/ubuntu/private_ubuntu.key'"
export REMOTE_PATH='/home/nicco/orderbooks'
export LOCAL_BASEPATH="${GOPATH}/src/bitbucket.org/niccolo_castelli/orderbooks"
#export IP="95.216.144.221" #OLD
export IP="192.168.10.3"

rsync -chavzP -e "ssh -i $KEY_FILE"   "${LOCAL_BASEPATH}/orderbooks" nicco@${IP}:"${REMOTE_PATH}/orderbooks"
rsync -chavzP -e  "ssh -i $KEY_FILE"   "${LOCAL_BASEPATH}/orderbooks.service" nicco@${IP}:"${REMOTE_PATH}/orderbooks.service"

#rsync -chavzP -e  "ssh -i $KEY_FILE"    "${LOCAL_BASEPATH}/_server_home/" ubuntu@${IP}:"${REMOTE_PATH}/"

