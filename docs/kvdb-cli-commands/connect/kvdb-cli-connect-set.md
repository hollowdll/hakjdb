# kvdb-cli connect set

```sh
kvdb-cli connect set [OPTIONS]
```

Changes the connection settings used to connect to a kvdb server.

## Options

- `-d`, `--database` - The database that commands will use by default.
- `-p`, `--port` - The port number of the server.
- `-a` `--host` - The hostname or IP address of the server.
- `-h`, `--help` - Show help page.

## Returns

- Nothing if successful.
- Error message if not successful.

## Examples

```sh
# change the address and port
kvdb-cli connect set -a 127.0.0.1 -p 9000

# change only the database
kvdb-cli connect set -d default
```