#!/bin/sh -e
# TODO run directly from Go?
_=${HSTRANSFER_SSH_HOST?"Please set the HSTRANSFER_SSH_HOST variable."}
_=${HSTRANSFER_SSH_SESSION_PATH?"Please set the HSTRANSFER_SSH_SESSION_PATH variable."}

ssh "$HSTRANSFER_SSH_HOST" mkdir -p "$HSTRANSFER_SSH_SESSION_PATH"
#ls -al /data_ready_to_upload/
rsync -rz --delete-after /data_ready_to_upload/ "$HSTRANSFER_SSH_HOST:$HSTRANSFER_SSH_SESSION_PATH/"
# -v
