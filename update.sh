#!/bin/sh -e
cd $(dirname $0)
HERE=$PWD

. ./env.sh

rm -rf bin pkg

for d in src/code.google.com src/github.com src/golang.org; do
  for pkg in $d/*/*; do
    if [ -d $pkg/.git ]; then
      cd $pkg
      git reset --hard
      cd ../../../..
    elif [ -d $pkg/.hg ]; then
      cd $pkg
      hg revert -C .
      cd ../../../..
    else
      rm -rf $pkg
    fi
  done
done

go get -u -x github.com/bmizerany/pq \
                 github.com/ziutek/mymysql/godrv \
                 github.com/rickb777/gorp \
                 github.com/go-sql-driver/mysql \
                 github.com/robertkrimen/terst \
                 github.com/robertkrimen/smuggol \
                 github.com/davecgh/go-spew/spew \
             || echo "Non-zero exit code from 'go get'."

./capture-current-deps.sh

go list ./...
