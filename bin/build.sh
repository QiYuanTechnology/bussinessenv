#!/bin/bash

dir=`pwd`

go mod tidy

build() {
	for d in $(ls ./$1); do
		echo "building $1/$d"
		pushd $dir/$1/$d >/dev/null
		CGO_ENABLED=0 GOOS=linux go build -o cmd -a -installsuffix cgo -ldflags '-w'
		popd >/dev/null
	done
}

build api
build srv
