# kvdb
In-memory key-value store. Can be used as a database or cache.

# Data types

Currently supported data types:
- `String`

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

# Docker

Images will eventually be available in an image registry. For now you have to build them.

## Build the server image

Make sure to be in the project root
```bash
cd kvdb
```
Alpine Linux
```bash
docker build -f "./docker/Dockerfile.alpine" -t kvdb:alpine .
```
Example of running the image
```bash
docker run -p 12345:12345 --rm -it kvdb:alpine
```

You can now access the server with `kvdb-cli` through port `12345`.
