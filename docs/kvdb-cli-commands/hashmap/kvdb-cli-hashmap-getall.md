# kvdb-cli hashmap getall

```sh
kvdb-cli hashmap getall [key] [OPTIONS]
```

Gets all the fields and values of the HashMap stored at a key.

## Arguments

- `key` - The name of the key.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- All the fields and values of the HashMap.
- `(None)` if the key doesn't exist.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli hashmap getall key1
1) "field1": "value1"
2) "field2": "value2"
3) "field3": "value3"

# specify the database to use
kvdb-cli hashmap getall key1 -d default
1) "field1": "value1"
2) "field2": "value2"
3) "field3": "value3"

# key doesn't exist
kvdb-cli hashmap getall key1234
(None)
```