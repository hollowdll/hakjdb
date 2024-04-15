# kvdb-cli hashmap delete

```sh
kvdb-cli hashmap delete [key] [field ...] [OPTIONS]
```

Removes the specified fields from the HashMap stored at a key.

This command can remove multiple fields.

## Arguments

- `key` - The name of the key.
- `field` - Field to remove.

## Options

- `-d`, `--database` - The database to use. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- The number of fields that were removed.
- Error message if not successful.

## Examples

```sh
# use the default database that is configured in the config file
kvdb-cli hashmap delete key1 field1
1

# specify the database to use
kvdb-cli hashmap delete key1 field2 -d default
1

# fields that do not exist are ignored
kvdb-cli hashmap delete key1 field1
0

# remove multiple fields
kvdb-cli hashmap delete key1 field3 field4 field5
3
```