# kvdb-cli set

```sh
kvdb-cli set [key] [value] [OPTIONS]
```

Sets a key to hold a String value. Creates the key if it doesn't exist. Overwrites the key if it is holding a value of another data type.

## Arguments

- `key` - The name of the key.
- `value` - The String value to store.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page

## Returns

- `OK` if successful.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli set key1 "Hello world!"
OK

# specify the database to use
kvdb-cli set key2 "some value" -d default
OK
```