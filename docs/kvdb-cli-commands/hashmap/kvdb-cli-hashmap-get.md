# kvdb-cli hashmap get

```sh
kvdb-cli hashmap get [key] [field ...] [OPTIONS]
```

Gets the values of the specified fields in the HashMap stored at a key.

This command can return multiple values.

## Arguments

- `key` - The name of the key.
- `field` - The field whose value to get.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The values of the specified fields.
- `(None)` if the key or field doesn't exist.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli hashmap get key1 field1
1) "field1": "value1"

# specify the database to use
kvdb-cli hashmap get key1 field1 -d default
1) "field1": "value1"

# return multiple values
kvdb-cli hashmap get key1 field1 field2 field3
1) "field1": "value1"
2) "field2": "value2"
3) "field3": "value3"

# key doesn't exist
kvdb-cli hashmap get key1234 field1
(None)

# field doesn't exist
kvdb-cli hashmap get key1 field123
1) "field123": (None)
```
