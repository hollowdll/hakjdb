# kvdb-cli delete

```sh
kvdb-cli delete [key] [OPTIONS]
```

Deletes a key and the value it is holding.

## Arguments

- `key` - The name of the key.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- `true` if the key exists and was deleted.
- `false` if the key doesn't exist.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli delete key1
true

# specify the database to use
kvdb-cli delete key2 -d default
true

# the key doesn't exist anymore
kvdb-cli delete key1
false
```