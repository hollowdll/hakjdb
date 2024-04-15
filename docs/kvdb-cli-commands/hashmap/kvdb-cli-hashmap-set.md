# kvdb-cli hashmap set

```sh
kvdb-cli hashmap set [key] [field value ...] [OPTIONS]
```

Sets the specified fields and their values in the HashMap stored at a key. If the specified fields exist, they will be overwritten with the new values. Creates the key if it doesn't exist. Overwrites the key if it is holding a value of another data type.

This command can set multiple fields.

## Arguments

- `key` - The name of the key.
- `field` - Field to set.
- `value` - Value to be stored in the field.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The number of fields that were added.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli hashmap set key1 field1 "value1"
1

# specify the database to use
kvdb-cli hashmap set key1 field1 "value1" -d default
0

# set multiple fields
kvdb-cli hashmap set key1 field1 "value1" field2 "value2" field3 "value3"
2

# update existing fields without adding new ones
kvdb-cli hashmap set key1 field1 "value111" field2 "value222" field3 "value333"
0
```