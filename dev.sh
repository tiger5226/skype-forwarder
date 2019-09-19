#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
APP_DIR="$DIR"

(

  hash reflex 2>/dev/null || go get github.com/cespare/reflex
  hash reflex 2>/dev/null || { echo >&2 'Make sure $GOPATH/bin is in your $PATH'; exit 1;  }

  hash go-bindata 2>/dev/null || go get github.com/jteeuwen/go-bindata/...

  cd "$APP_DIR"

  reflex --decoration=none --start-service=true --regex='\.go$' --inverse-regex='migration/bindata\.go' -- sh -c "go generate && go run *.go serve"
)