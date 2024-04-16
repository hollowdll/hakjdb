# kvdb-cli connect show

```sh
kvdb-cli connect show [OPTIONS]
```

Shows the currently configured connection settings used to connect to a kvdb server.

## Options

- `-h`, `--help` - Show help page.

## Returns

- The currently configured connection settings.
- Error message if not successful.

## Examples

```sh
kvdb-cli connect show
host: localhost
port: 12345
database: default
```