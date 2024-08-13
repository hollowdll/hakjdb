# Developer Documentation

This documentation explains things that can be useful for the developer.

# Generate gRPC API code from the proto files

When changes are made to the gRPC API proto files, new protobuf code needs to be generated.

Run the `compile_protos.sh` script found in the project root
```sh
./compile_protos.sh
```

You may need to give it execute permission
```sh
sudo chmod u+x compile_protos.sh
```

# Releasing new versions

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
