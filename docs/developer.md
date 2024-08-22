# Developer Documentation

This documentation explains things that can be useful for the developer.

## Generate gRPC API code from the proto files

When changes are made to the gRPC API proto files, new protobuf code needs to be generated.

Run the `compile_protos.sh` script found in the project root
```sh
./compile_protos.sh
```

You may need to give it execute permission
```sh
sudo chmod u+x compile_protos.sh
```

## Releasing new versions

Make sure to update the version in `version/version.go` source code file.

Add a new changelog documentation file to the `docs/changelog/` directory that explains what is included in the release. Changelog files should have format `v<version_number>.md` where `<version_number>` is the version number of the release. For example `v0.1.0.md`.

There is a file `.goreleaser.yaml` in the project root that declares the contents of a release. Releases are automated with a CI workflow and use goreleaser. When a new release is triggered, the CI workflow begins and makes a new release along with the changelog, built binaries, and Docker images.

Releases work with git tags. When a new tag is pushed to the repository, it triggers a new release.

goreleaser documentation [here](https://goreleaser.com/quick-start/).

Example release
```sh
git tag -a v0.1.0 -m "Release v0.1.0"
git push origin v0.1.0
```
Replace `v0.1.0` with the actual release version.

## Update the gRPC API version

- Update the `APIVersion` constant in `version/version.go` source code file.
- Update the API version comment in all .proto files in `api/` directory.
- Treat the API version as a whole in all proto files so it is clear which API version the proto file belongs to.

## Generate and update kvdbctl command documentations

Run the script `gen_kvdbctl_command_docs.sh` from the project root
```sh
./gen_kvdbctl_command_docs.sh
```

This generates the updated command documentation and places it in `docs/kvdbctl-commands/generated` directory.

## Generate self-signed TLS certificate and private key for testing

Currently no native mTLS support so only server certificate and private key.

Directory `tls/test-cert` has a cert.conf for self-signed certificate configuration. Certificates should be placed there.

Example of generating certificate file and private key using openssl:
```sh
sudo openssl req -x509 -nodes -days 3650 -newkey rsa:2048 -keyout hakjserver.key -out hakjserver.crt -config cert.conf
```
