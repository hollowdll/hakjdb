# kvdb-cli keytype

```sh
kvdb-cli keytype [key] [OPTIONS]
```

Gets the data type of the value a key is holding.

## Arguments

- `key` - The name of the key.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The data type.
- `(None)` if the key doesn't exist.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli keytype string-key
"String"

kvdb-cli keytype hashmap-key
"HashMap"

# specify the database to use
kvdb-cli keytype string-key -d default
"String"

# key doesn't exist
kvdb-cli keytype this-key-does-not-exist
(None)
```