# kvdb-cli hashmap get

```sh
kvdb-cli hashmap get [key] [field] [OPTIONS]
```

Gets the value of a field in the HashMap stored at a key.

## Arguments

- `key` - The name of the key.
- `field` - The field whose value to get.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The value of the field.
- `(None)` if the key or field doesn't exist.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli hashmap get key1 field1
"value1"

# specify the database to use
kvdb-cli hashmap get key1 field1 -d default
"value1"

# key doesn't exist
kvdb-cli hashmap get key1234 field1
(None)

# field doesn't exist
kvdb-cli hashmap get key1 field123
(None)
```