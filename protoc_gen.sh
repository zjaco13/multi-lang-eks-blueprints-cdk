#!/bin/bash

set -e

PROTO_FILES=$(find ./proto -name "*.proto")
PROTO_DIR=proto
GO_DIR=go/codegen
TS_DIR=ts/lib
PY_DIR=python

echo $PROTO_FILES

echo "Generating .go files!"
protoc \
	--proto_path=$PROTO_DIR \
	--go_out=$GO_DIR --go_opt=paths=source_relative \
	--go-grpc_out=$GO_DIR --go-grpc_opt=paths=source_relative \
	$PROTO_FILES

echo "Generating .ts files!"
protoc \
	--proto_path=$PROTO_DIR \
	--ts_out=$TS_DIR --ts_opt=target=node,outputServices=grpc-js \
	$PROTO_FILES

echo "Generating .py files"
python3 -m grpc_tools.protoc \
	--proto_path=$PROTO_DIR \
	--python_out=$PY_DIR --pyi_out=$PY_DIR \
	--grpc_python_out=$PY_DIR \
	$PROTO_FILES
