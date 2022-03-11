# `/api`

This directory contains [Protocol Buffers](https://developers.google.com/protocol-buffers) definition files to specify the communication protocol between client and server over [gRPC](https://grpc.io/).

## System Requirements

In order to generate Go code from these definition files, you will need the following system dependencies:

* `protoc` - The Protocol Buffers Compiler
* Go protocol compiler plugins `protoc-gen-go` and `protoc-gen-go-grpc`

For installation instructions, see the [gRPC Go Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/#prerequisites).