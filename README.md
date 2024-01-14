# kvdb
In-memory key-value store. Can be used as a database or cache.

Currently supported data types:
- `Strings`

Proper documentation coming soon.

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
