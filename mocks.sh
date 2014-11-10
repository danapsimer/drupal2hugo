#! /bin/sh -e
. ./env.sh

[ ! -d src/code.google.com ] && ./update.sh
[ ! -x bin/mockgen ] && go install code.google.com/p/gomock/...

DIR=src/mock/net/http
mkdir -p $DIR
bin/mockgen -package http net/http Handler > $DIR/mock_handler.go
bin/mockgen -package http net/http ResponseWriter > $DIR/mock_responsewriter.go
gofmt -w $DIR/mock_responsewriter.go
