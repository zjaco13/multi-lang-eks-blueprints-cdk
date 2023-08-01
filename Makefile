# Dependencies
HOMEBREW_LIBS := go python rust cargo protobuf

VENV := venv
PIP := $(VENV)/bin/pip
PYTHON := $(VENV)/bin/python3
GO_BINARY_NAME := example-main
GO_SDK:= sdks/go
RUST_SDK:= sdks/rust
PYTHON_SDK:= sdks/python

.PHONY: codegen go-example python-example rust-example clean go-build

codegen:
	bash scripts/codegen.sh

go-build:
	cd $(GO_SDK) && go build -o $(GO_BINARY_NAME) example/main.go
go-example: go-build
	./$(GO_SDK)/$(GO_BINARY_NAME)

$(VENV)/bin/activate: sdks/python/requirements.txt
	python3 -m venv venv
	$(PIP) install -r $(PYTHON_SDK)/requirements.txt

python-example: $(VENV)/bin/activate
	$(PYTHON) $(PYTHON_SDK)/example.py

rust-example:
	cargo run --manifest-path=$(RUST_SDK)/Cargo.toml --example client-example

clean:
	rm -rf $(VENV) 
	rm -f $(GO_SDK)/$(GO_BINARY_NAME)
	rm -rf $(PYTHON_SDK)/src/__pycache__
	rm -rf $(PYTHON_SDK)/src/codegen/__pycache__
	rm -rf $(RUST_SDK)/target
