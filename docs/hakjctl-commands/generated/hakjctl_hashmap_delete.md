## hakjctl hashmap delete

Remove fields from a HashMap key value

### Synopsis

Remove the specified fields from a HashMap key value.
Ignores fields that do not exist. This command can remove multiple fields.
Returns the number of fields that were removed.


```
hakjctl hashmap delete KEY FIELD [FIELD ...] [flags]
```

### Examples

```
# Use the default database
hakjctl hashmap delete key1 field1

# Specify the database to use
hakjctl hashmap delete key1 field2 -d default

# Remove multiple fields
hakjctl hashmap delete key1 field3 field4 field5
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for delete
```

### SEE ALSO

* [hakjctl hashmap](hakjctl_hashmap.md)	 - Manage HashMap keys

