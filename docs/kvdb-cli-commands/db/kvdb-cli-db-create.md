# kvdb-cli db create

```sh
kvdb-cli db create [OPTIONS]
```

Creates a new database.

## Options

- `-n`, `--name` - The name of the database. Required.
- `-h`, `--help` - Show help page.

## Returns

- The name of the created database.
- Error message if not successful.

## Examples

```sh
kvdb-cli db create --name db1
db1

kvdb-cli db create -n db2
db2
```