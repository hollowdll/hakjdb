# kvdb-cli deletekeys

```sh
kvdb-cli deletekeys [OPTIONS]
```

Deletes all the keys of a database.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- `OK` if successful.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli deletekeys
OK

# specify the database to use
kvdb-cli deletekeys -d default
OK
```