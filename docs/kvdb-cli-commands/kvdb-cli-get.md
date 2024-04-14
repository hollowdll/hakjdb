# kvdb-cli get

```sh
kvdb-cli get [key] [OPTIONS]
```

Returns the value a String key is holding.

## Arguments

- `key` - The name of the key.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page

## Returns

- The value the key is holding if the key exists.
- `(None)` if the key doesn't exist.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli get key1
"Hello world!"

# specify the database to use
kvdb-cli get key1 -d default
"Hello world!"

# key doesn't exist
kvdb-cli get key1234
(None)
```