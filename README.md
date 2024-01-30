# Overview
kvdb is an in-memory key-value store written in go language. It can be used as a database or cache.

Components:
- `kvdbserver` - The server process
- `kvdb-cli` - CLI tool to manage the server
  
**Note!** This project is in early development.

# Data types

Currently supported data types:
- `String`

# Documentation

- [Using the CLI](./docs/kvdb-cli.md)
- [Using the server](./docs/kvdbserver.md)
- [Databases](./docs/databases.md)
- [Data types](./docs/datatypes.md)

# Releases

Release notes are available [here](./docs/changelog/).

Binaries are available for download. Multiple platforms supported. You can download them [here](https://github.com/hollowdll/kvdb/releases).

However, it is recommended to build the binaries from source. Instructions on how to build from source below.

# Build binaries

To build the binaries, you first need to install Go.

Instructions [here](https://go.dev/doc/install).

You may also need tools to work with gRPC and Protocol Buffers in Go. This is needed if you want to compile `.proto` files and generate Go code.

- [Protocol Buffer compiler](https://github.com/protocolbuffers/protobuf#protobuf-compiler-installation)
- [Quickstart](https://grpc.io/docs/languages/go/quickstart/)

After you have successfully installed go, clone this repository.

Cloning with git:
```bash
git clone https://github.com/hollowdll/kvdb.git
```

Change directory to the project root:
```bash
cd kvdb
```

Get the dependencies:
```bash
go mod tidy
```

Build the server:
```bash
go build -o ./bin/kvdbserver/ ./cmd/kvdbserver/
```

Build the CLI:
```bash
go build -o ./bin/kvdb-cli/ ./cmd/kvdb-cli/
```

These will build the binaries to `bin/` directory in the project root. You can change the output directory and binary names to whatever you like by modifying the path with `-o` flag.

For more advanced build, use `go help build` to see more build options.

# Environment variables

## Server

- `KVDB_PORT`: Server TCP/IP port.
- `KVDB_DEBUG_ENABLED`: Controls whether debug mode is enabled. If true, debug messages are logged.

## Tests

- `KVDB_PORT`: Test server TCP/IP port. Default port is 12345.
- `KVDB_HOST`: Test server address. Default address is localhost.

# Docker

Images are available in Docker Hub. See [repository](https://hub.docker.com/r/hakj/kvdb).

## Pull the server image

```bash
docker pull hakj/kvdb
```

## Build the server image

Make sure to be in the project root
```bash
cd kvdb
```
Latest tag
```bash
docker build -f "./Dockerfile.bookworm" -t kvdb:latest .
```
Debian based image
```bash
docker build -f "./Dockerfile.bookworm" -t kvdb:bookworm .
```
Alpine Linux based image
```bash
docker build -f "./Dockerfile.alpine" -t kvdb:alpine .
```

These commands build the image only for a single architecture. If you want to build multi-arch images for other platforms, read [this](https://docs.docker.com/build/building/multi-platform/).

## Start a container

Example of starting a container
```bash
docker run -p 12345:12345 --rm -it kvdb
```
This binds the host's port `12345` to the container's port `12345` so you can access the server outside the container.

# Running tests

Change directory to the project root:
```bash
cd kvdb
```

Run all tests:
```bash
go test ./...
```

Run only integration tests:
```bash
go test ./tests/integration
```
