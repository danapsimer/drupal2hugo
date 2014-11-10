#!/bin/sh -e
# other useful value is 'html'
MODE=${1:-func}

cover()
{
    tmp=$(tempfile)
    go test -coverprofile=$tmp $@ && \
      go tool cover -$MODE=$tmp && \
      unlink $tmp
}


for dir in $(find src/drupal2hugo/ -type d); do
  pkg=$(echo $dir | cut -c5-)
  cover $pkg
done

