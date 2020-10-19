#!/bin/sh -xe
# Ran during docker build.
echo "Executing integration test..."

#
# Launch a dummy static web server
#
# Nginx or apache would require more configuraiton steps
mkdir -p /tmp/static_web_server_content
cd /tmp/static_web_server_content
python3 -m http.server --bind 127.0.0.1 &# Will launch a static web server on port 8000

#
# Install ssh server
#
ssh-keygen -A

ssh-keygen -f /root/.ssh/id_rsa -q -N ''
echo "root:root" | chpasswd
cat /root/.ssh/id_rsa.pub >/root/.ssh/authorized_keys

/usr/sbin/sshd

ssh -oStrictHostKeyChecking=no root@localhost -- exit # Ensures it works fine + add host key to /root/.ssh/known_hosts

#
# Launch hstransfer
#
# Create a dummy directory we would like to sync:
mkdir /directory_to_upload/
echo "hello world $(uuidgen)" >/directory_to_upload/somefile.txt
# Launch uploader with once
cp /root/.ssh/id_rsa /id_rsa
export HSTRANSFER_SSH_HOST="root@localhost"
export HSTRANSFER_SSH_PATH="/tmp/static_web_server_content/something/"
export HSTRANSFER_HTTP_PATH="http://localhost:8000/something/"
set -o pipefail
/go/src/hstransfer/sh/bootstrap_uploader.sh --run-once | tee uploader.log
# Get downloader
ENCRYPTION_KEY="$(cat uploader.log | grep "encryptionKey" | tail -n 1 | cut -d ' ' -f 2)"
DOWNLOADER_URL="$(cat uploader.log | grep "http://" | tail -n 1)"
echo "ENCRYPTION_KEY=$ENCRYPTION_KEY"
echo "DOWNLOADER_URL=$DOWNLOADER_URL"
wget -nv "$DOWNLOADER_URL"
chmod +x downloader-linux
# Run it
HSTRANSFER_ENCRYPTION_KEY=$ENCRYPTION_KEY ./downloader-linux --run-once
# TODO we should be able to say where to download content.
WANTED_MD5_SUM="$(md5sum /directory_to_upload/somefile.txt | cut -d ' ' -f 1)"
RESULT_MD5_SUM="$(md5sum hstransfer/somefile.txt | cut -d ' ' -f 1)"
python3 -c "assert '$WANTED_MD5_SUM' == '$RESULT_MD5_SUM'"
