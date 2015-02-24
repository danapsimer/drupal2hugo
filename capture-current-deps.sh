#!/bin/sh -e
cd $(dirname $0)
HERE=$PWD

. ./env.sh

echo '#!/bin/sh -ex' > get-current-deps.sh

current()
{
    echo "$@" >> $HERE/get-current-deps.sh
}

cd src
for host in github.com code.google.com golang.org; do
    if [ -d $host ]; then
        for repo in $host/*/*; do
            current ""
            cd $repo

            if [ -d .git ]; then
                current "mkdir -p src/$repo"
                current "cd src/$repo"
                rev=$(git log -1 | grep commit | sed 's#commit ##')
                current "git reset -q --hard"
                current "git fetch"
                current "git checkout -q --detach $rev"
                current "cd ../../../.."
            fi

            if [ -d .hg ]; then
                current "mkdir -p src/$repo"
                current "cd src/$repo"
                rev=$(hg id -i)
                current "hg revert -q -C ."
                current "hg pull"
                current "hg update -r $rev"
                current "cd ../../../.."
            fi

            cd ../../..
        done
    fi
done
cd ..

chmod +x get-current-deps.sh
