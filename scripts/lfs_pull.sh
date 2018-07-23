#!/bin/sh

if [ $EUID != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

curl -s https://packagecloud.io/install/repositories/github/git-lfs/script.deb.sh | sudo bash

cd $GOPATH/src/github.com/bernardoaraujor/corinda
git lfs pull
