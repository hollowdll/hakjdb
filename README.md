# Overview
kvdb is an in-memory key-value store. It can be used as a database or cache.

Components:
- `kvdbserver` - The server process
- `kvdb-cli` - CLI tool to manage the server
  
**Note!** This project is in early development. No releases yet, but we are close to v0.1.0. Lots of cool stuff coming after that.

# Data types

Currently supported data types:
- `String`

# Getting started

Coming soon.

# How to use

Coming soon.

# Build binaries

To build the binaries, you first need to install Go.

Instructions [here](https://go.dev/doc/install)

You may also need tools to work with gRPC and Protocol Buffers in Go. This is needed if you want to compile `.proto` files and generate Go code.

- [Protocol Buffer compiler](https://github.com/protocolbuffers/protobuf#protobuf-compiler-installation)
- [Quickstart](https://grpc.io/docs/languages/go/quickstart/)

After you have successfully installed go, clone this repository.

Change directory to the project root
```bash
cd kvdb
```

Build the server binary
```bash
go build -o ./bin/kvdbserver/ ./cmd/kvdbserver/
```

Build the CLI tool
```bash
go build -o ./bin/kvdb-cli/ ./cmd/kvdb-cli/
```

These will build the binaries to `bin/` directory in the project root. You can change the output directory and binary names to whatever you like by modifying the path with `-o` flag.

For more advanced build, use `go help build` to see more build options.

# Environment

## Server

Environment variables for the server

| Name            | Default value | Description |
|-----------------|---------------|-------------|
| KVDB_PORT       | 12345         | Server TCP/IP port. |
| KVDB_DEBUG_MODE | false         | Controls whether debug mode is enabled. If enabled, debug logs will show. |

## Tests

Environment variables for tests

| Name            | Default value | Description |
|-----------------|---------------|-------------|
| KVDB_PORT       | 12345         | Test server TCP/IP port. |
| KVDB_HOST       | localhost     | Test server address. |

# Docker

Images are available in Docker Hub. See [repository](https://hub.docker.com/r/hakj/kvdb)

## Pull the server image

```bash
docker pull hakj/kvdb
```

## Build the server image

Make sure to be in the project root
```bash
cd kvdb
```
Debian based image
```bash
docker build -f "./docker/Dockerfile.bookworm" -t kvdb:latest .
```
Alpine Linux based image
```bash
docker build -f "./docker/Dockerfile.alpine" -t kvdb:alpine .
```
Example of running the image
```bash
docker run -p 12345:12345 --rm -it kvdb:alpine
```
You can access the server through port `12345`.
