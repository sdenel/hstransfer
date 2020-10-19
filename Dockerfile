FROM golang:1.15-alpine as base
# openssh-client is for scp, util-linux for uuidgen
RUN apk add openssh-client util-linux git build-base findutils rsync zlib
# Is there a better way to get test dependancies?!
RUN go get github.com/stretchr/testify/assert
# integration test dependancies.
# TODO: we could use a multi stage pipeline to avoid those dependancies (+ git, assert, ...)
RUN apk add python3 openssh-server wget


COPY . /go/src/hstransfer/
WORKDIR /go/src/hstransfer/
# tests.
RUN time go test -v -cover ./lib
# Build
RUN time go build
# Integration test
RUN ./integration_test.sh

ENTRYPOINT ["/go/src/hstransfer/sh/bootstrap_uploader.sh"]