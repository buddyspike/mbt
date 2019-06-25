#!/bin/sh

set -e

DIR=$(pwd)
LIBGIT2_PATH=$DIR/vendor/libgit2
export PKG_CONFIG_PATH="$LIBGIT2_PATH/build"
export CGO_LDFLAGS="$(pkg-config --libs --static $LIBGIT2_PATH/build/libgit2.pc)"

go test -cpuprofile -bench -run=^$ github.com/mbtproject/mbt/lib
