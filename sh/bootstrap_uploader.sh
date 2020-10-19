#!/bin/sh -xe
eval $(ssh-agent -s)
ssh-add /id_rsa

_=${HSTRANSFER_SSH_HOST?"Please set the HSTRANSFER_SSH_HOST variable."}
_=${HSTRANSFER_SSH_PATH?"Please set the HSTRANSFER_SSH_PATH variable."}
_=${HSTRANSFER_HTTP_PATH?"Please set the HSTRANSFER_HTTP_PATH variable."}
python3 -c "assert '$HSTRANSFER_SSH_PATH'.endswith('/'), 'HSTRANSFER_SSH_PATH=$HSTRANSFER_SSH_PATH should end with a /'"
python3 -c "assert '$HSTRANSFER_HTTP_PATH'.endswith('/'), 'HSTRANSFER_HTTP_PATH=$HSTRANSFER_HTTP_PATH should end with a /'"


export HSTRANSFER_SESSION_ID="$(uuidgen)"
export HSTRANSFER_SSH_SESSION_PATH="$HSTRANSFER_SSH_PATH$HSTRANSFER_SESSION_ID"
export HSTRANSFER_HTTP_SESSION_PATH="$HSTRANSFER_HTTP_PATH$HSTRANSFER_SESSION_ID"

ssh "$HSTRANSFER_SSH_HOST" echo "echo 'Sanity check: Hello world from remote server ($HSTRANSFER_SSH_HOST)'"

cd /go/src/hstransfer/
HSTRANSFER_MODE="uploader" ./hstransfer $@
