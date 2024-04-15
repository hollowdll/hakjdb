# kvdb-cli getkeys

```sh
kvdb-cli getkeys [OPTIONS]
```

Get all the keys of a database.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- All the keys of a database.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli getkeys
1) key1
2) key2
3) key3

# specify the database to use
kvdb-cli getkeys -d default
1) key1
2) key2
3) key3
```