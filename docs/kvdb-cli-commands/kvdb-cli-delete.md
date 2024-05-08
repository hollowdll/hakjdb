# kvdb-cli delete

```sh
kvdb-cli delete [key ...] [OPTIONS]
```

Deletes the specified keys and the values they are holding. Ignores keys that do not exist.

This command can delete multiple keys.

## Arguments

- `key` - Key to delete.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The number of keys that were deleted.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli delete key1
1

# specify the database to use
kvdb-cli delete key2 -d default
1

# the key doesn't exist anymore
kvdb-cli delete key1
0

# delete multiple keys
kvdb-cli delete key3 key4 key5
3
```