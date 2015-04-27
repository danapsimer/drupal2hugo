#!/bin/sh -ex

mkdir -p src/github.com/bmizerany/pq
cd src/github.com/bmizerany/pq
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach da2b95e392c172df0be322eadc19a90b0f771548
else 
  git clone git@github.com:bmizerany/pq.git .
fi
cd ../../../..

mkdir -p src/github.com/davecgh/go-spew
cd src/github.com/davecgh/go-spew
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach fc32781af5e85e548d3f1abaf0fa3dbe8a72495c
else
  git clone git@github.com:davecgh/go-spew.git .
fi
cd ../../../..

mkdir -p src/github.com/go-sql-driver/mysql
cd src/github.com/go-sql-driver/mysql
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach 7d52f0fc9e81d12e0d0a2f8b6749fb96058d73d4
else
  git clone git@github.com:go-sql-driver/mysql.git .
fi
cd ../../../..

mkdir -p src/github.com/rickb777/gorp
cd src/github.com/rickb777/gorp
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach 5d19ebd22fdcf3b9fb1c8cbfcf1d4e78e102f1e0
else
  git clone git@github.com:rickb777/gorp.git .
fi
cd ../../../..

mkdir -p src/github.com/robertkrimen/smuggol
cd src/github.com/robertkrimen/smuggol
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach 3d74d482c057d0a5f7d0206bc4d45f956c09ddef
else
  git clone git@github.com:robertkrimen/smuggol.git .
fi
cd ../../../..

mkdir -p src/github.com/robertkrimen/terst
cd src/github.com/robertkrimen/terst
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach 4b1c60b7cc23825033c7cecf3e985a41f6e87b53
else
  git clone git@github.com:robertkrimen/terst.git .
fi
cd ../../../..

mkdir -p src/github.com/ziutek/mymysql
cd src/github.com/ziutek/mymysql
if [[ -d ./.git ]]; then
  git reset -q --hard
  git fetch
  git checkout -q --detach e08c2f35356576b3c3690c252fe5dca728ae9292
else
  git clone git@github.com:ziutek/mymysql.git .
fi
cd ../../../..
