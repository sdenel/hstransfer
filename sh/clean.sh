#!/bin/sh -xe
# TODO run directly from Go?
_=${HSTRANSFER_SSH_HOST?"Please set the HSTRANSFER_SSH_HOST variable."}
_=${HSTRANSFER_SSH_SESSION_PATH?"Please set the HSTRANSFER_SSH_SESSION_PATH variable."}

echo "Deleting remote dir: $HSTRANSFER_SSH_SESSION_PATH"
ssh "$HSTRANSFER_SSH_HOST" rm -rf "$HSTRANSFER_SSH_SESSION_PATH"
