#!/bin/bash

set -e

PROTO_DIR=./proto
PROTO_FILES=$(find $PROTO_DIR -type f -iname "*.proto" -exec basename {} \;)
GO_DIR=sdks/go/proto
TS_DIR=server/lib
PY_DIR=sdks/python/proto

echo $PROTO_FILES

echo "Generating .go files!"
protoc \
	-I=$PROTO_DIR \
	--go_out=$GO_DIR --go_opt=paths=source_relative \
	--go-grpc_out=$GO_DIR --go-grpc_opt=paths=source_relative \
	$PROTO_FILES

echo "Generating .ts files!"
protoc \
	-I=$PROTO_DIR \
	--plugin=./server/node_modules/.bin/protoc-gen-ts_proto \
	--ts_proto_out=$TS_DIR --ts_proto_opt=env=node,outputServices=grpc-js,outputIndex=true \
	$PROTO_FILES

echo "Generating .py files"
python3 -m grpc_tools.protoc \
	-I=$PROTO_DIR \
	--python_out=$PY_DIR --pyi_out=$PY_DIR \
	--grpc_python_out=$PY_DIR \
	$PROTO_FILES