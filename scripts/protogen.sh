#!/usr/bin/env sh

ROOT=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)

cd "$ROOT" || exit 1

OUT="$ROOT/internal/pkg/proto"
if [ -d "$OUT" ]; then
    rm -rf "$OUT"
fi
mkdir -p ${OUT}

GOPATH=$(go env GOPATH)
if [ "$GOBIN" = "" ]; then
  if [ "$GOPATH" = "" ]; then
    echo "Required env var GOPATH is not set; aborting with error; see the following documentation which can be invoked via the 'go help gopath' command."
    go help gopath
    exit 1
  fi

  echo "Optional env var GOBIN is not set; using default derived from GOPATH as: \"$GOPATH/bin\""
  GOBIN="$GOPATH/bin"
fi

PROTOC=`command -v protoc`
if [ "$PROTOC" = "" ]; then
  echo "Required "protoc" to be installed. Please visit https://github.com/protocolbuffers/protobuf/releases."
  exit 1
fi

echo "Compiling protobuf definitions.."

FILES=$(find ./api -type f -name "*.proto" -printf '%P ')

"${PROTOC}" \
  --proto_path=./api \
  --proto_path=/usr/local/include \
  --go_out="${OUT}" --go_opt=paths=source_relative \
  --go-grpc_out="${OUT}" --go-grpc_opt=paths=source_relative \
  ${FILES}


