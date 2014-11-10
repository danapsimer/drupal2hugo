#!/bin/sh -e
cd $(dirname $0)
HERE=$PWD

. ./env.sh

fetchFreshDepsFromInternet()
{
    set -x
    rm -rf bin pkg dependencies.tar.gz

    for d in src/github.com; do
      for pkg in $d/*/*; do
        if [ ! -d $pkg/.git -a ! -d $pkg/.hg ]; then
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
                 || echo "Non-zero exit code from 'go get'."

    for examples in src/github.com/ziutek/mymysql/examples \
                    $(find src/github.com/mattn/go-sqlite3/_example -type d); do
        cd $examples
        for src in *.go; do
            [ -f "$src" ] && mv -v "$src" "$src.txt"
        done
        cd $HERE
    done

    [ -d src/code.google.com ] && GOOGLE=src/code.google.com
    [ -d src/github.com      ] && GITHUB=src/github.com
    [ -d src/launchpad.net   ] && LAUNCHPAD=src/launchpad.net

    tar cvf dependencies.tar --exclude-vcs $GITHUB $LAUNCHPAD $GOOGLE
    gzip dependencies.tar

    go list ./...
}

unpackTarball()
{
    rm -rf bin pkg src/github.com src/code.google.com
    tar zxvf dependencies.tar.gz
}

case "$1" in
    unpack|local|tarball) unpackTarball ;;
    help|-*)              echo "Usage: $0 [unpack | download]" ;;
    *)                    fetchFreshDepsFromInternet ;;
esac

