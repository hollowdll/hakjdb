# kvdb-cli db delete

```sh
kvdb-cli db delete [OPTIONS]
```

Deletes a database.

## Options

- `-n`, `--name` - The name of the database. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The name of the deleted database.
- Error message if not successful.

## Examples

```sh
# specify the database to delete
kvdb-cli db delete --name db1
db1

# delete the default database that is configured in the config file
kvdb-cli db delete
db2
```