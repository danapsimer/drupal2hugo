#!/bin/sh -e
. ./env.sh

[ ! -d src/github.com ] && ./update.sh
#[ ! -d src/mock ] && ./mocks.sh

mkdir -p src/drupal2hugo/util
VFILE=src/drupal2hugo/util/version.go

echo "// Generated automatically" > $VFILE
echo "package util" >> $VFILE
echo "const HgTip     = \"$(hg heads . |grep changeset: |head -1 |cut -f3 -d:)\"" >> $VFILE
echo "const HgPath    = \"$(hg paths default)\"" >> $VFILE
echo "const HgBranch  = \"$(hg branch)\"" >> $VFILE
echo "const BuildDate = \"$(date '+%FT%T')\"" >> $VFILE
echo "const BuildYear = \"$(date '+%Y')\"" >> $VFILE
echo "const BuildUser = \"$USER\"" >> $VFILE

echo go install ...
go install github.com/... code.google.com/... drupal2hugo/...

echo go test ...
go test drupal2hugo/...
