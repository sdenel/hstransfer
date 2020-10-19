#!/bin/sh -xe

_=${HSTRANSFER_HTTP_SESSION_PATH?"Please set the HSTRANSFER_HTTP_SESSION_PATH variable."}

echo "    Building downloader [HSTRANSFER_HTTP_SESSION_PATH=$HSTRANSFER_HTTP_SESSION_PATH]"
# Flags:
# "-s -w" reduces the program size. See https://stackoverflow.com/a/37468877
# Building the binary with session variables ensures it always have a different hash ;-)
LD_FLAGS="-s -w -X main.httpSessionPath=$HSTRANSFER_HTTP_SESSION_PATH"
echo "    LD_FLAGS=$LD_FLAGS"
# https://stackoverflow.com/a/36308464
CGO_ENABLED=0 go build -ldflags "$LD_FLAGS" -o downloader-linux
GOOS=windows GOARCH=amd64 go build -ldflags "$LD_FLAGS" -o downloader-windows.exe
echo "      Done building downloader. Moving binaries to /data_ready_to_upload/..."

mkdir -p /data_ready_to_upload/
mv downloader-linux /data_ready_to_upload/downloader-linux
mv downloader-windows.exe /data_ready_to_upload/downloader-windows.exe
# ls + md5sum, see https://stackoverflow.com/a/6874612
#find /data_ready_to_upload/ -type f -exec sh -c 'printf "%s %s \n" "$(ls -l $1)" "$(md5sum $1)"' '' '{}' '{}' \;
echo "      Done moving binaries."