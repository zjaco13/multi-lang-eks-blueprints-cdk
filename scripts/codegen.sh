#!/bin/bash

set -e

PROTO_DIR=./proto
PROTO_FILES=$(find $PROTO_DIR -type f -iname "*.proto" -exec basename {} \;)
PROTO_FILE=./proto/cluster.proto
GO_DIR=./sdks/go/codegen
PY_DIR=./sdks/python/src/codegen
RS_DIR=./sdks/rust/src/proto

echo $PROTO_FILES

echo "Generating .go files!"
protoc \
	-I=$PROTO_DIR \
	--go_out=$GO_DIR --go_opt=paths=source_relative \
	--go-grpc_out=$GO_DIR --go-grpc_opt=paths=source_relative \
	$PROTO_FILES

echo "Generating .py files"
python3 -m grpc_tools.protoc \
	-I=$PROTO_DIR \
	--python_out=$PY_DIR --pyi_out=$PY_DIR \
	--grpc_python_out=$PY_DIR \
	$PROTO_FILES
sed -i "" "s/import team/from . import team/g" $PY_DIR/cluster_pb2.py
sed -i "" "s/import addons/from . import addons/g" $PY_DIR/cluster_pb2.py
sed -i "" "s/import cluster/from . import cluster/g" $PY_DIR/cluster_pb2.py
sed -i "" "s/import resource/from . import resource/g" $PY_DIR/cluster_pb2.py

sed -i "" "s/import team/from . import team/g" $PY_DIR/cluster_pb2_grpc.py
sed -i "" "s/import addons/from . import addons/g" $PY_DIR/cluster_pb2_grpc.py
sed -i "" "s/import cluster/from . import cluster/g" $PY_DIR/cluster_pb2_grpc.py
sed -i "" "s/import resource/from . import resource/g" $PY_DIR/cluster_pb2_grpc.py

echo "Generating .rs files"
cargo build --manifest-path=./sdks/rust/Cargo.toml
