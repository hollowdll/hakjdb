# Overview
kvdb is an in-memory key-value store written in go language, using a simple client-server model. It can be used as a database, session storage, or cache. It is designed to be simple and easy to use. It may not be suitable for complex needs.

Data is stored at keys of different types. You can store data ranging from basic string values to large objects.

Instances are easily configurable with environment variables and a simple JSON file.

Components:
- `kvdbserver` - The server process
- `kvdb-cli` - CLI tool to manage the server
  
**Note!** This project is in early development.

# Data types

Currently supported data types:
- `String`
- `HashMap`

# Documentation

- [Using the CLI](./docs/kvdb-cli.md)
- [Using the server](./docs/kvdbserver.md)
- [Databases](./docs/databases.md)
- [Data types](./docs/datatypes.md)
- [API](./docs/api.md)

# Releases

Release notes are available [here](./docs/changelog/).

Binaries are available for download. Multiple platforms supported. You can download them [here](https://github.com/hollowdll/kvdb/releases).

However, it is recommended to build the binaries from source. Instructions on how to build from source below.

Releases are managed with [GoReleaser](https://goreleaser.com/).

# Build binaries

To build the binaries, you first need to install Go. Minimum version required is go1.22.

Instructions for installing Go can be found [here](https://go.dev/doc/install).

You may also need tools to work with gRPC and Protocol Buffers in Go. This is needed if you want to compile `.proto` files and generate Go code.

- [Protocol Buffer compiler](https://github.com/protocolbuffers/protobuf#protobuf-compiler-installation)
- [Quickstart](https://grpc.io/docs/languages/go/quickstart/)

If you just want to compile the binaries, then installing only Go is enough.

After you have successfully installed go, clone this repository.

Cloning with git:
```bash
git clone https://github.com/hollowdll/kvdb.git
```

Note: You can also download the source code for a specific release [here](https://github.com/hollowdll/kvdb/releases).

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

# Docker

Images are available in Docker Hub with multiple tags. Links below.

- [Repository](https://hub.docker.com/r/hakj/kvdb)
- [Tags](https://hub.docker.com/r/hakj/kvdb/tags)

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

## Integration test environment variables

Integration test environment is configurable. Below is a list of environment variables that can be used when running integration tests.

- `KVDB_PORT`: Test server TCP/IP port. Default port is 12345.
- `KVDB_HOST`: Test server address. The test client will try to connect to this. Default address is localhost.

# License

This project is licensed under MIT license. It is free and open source software.
