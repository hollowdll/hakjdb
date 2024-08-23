## hakjctl hashmap set

Set HashMap fields and values

### Synopsis

Set the specified fields and their values of a HashMap key value.
If the specified fields exist, they will be overwritten with the new values.
Creates the key if it doesn't exist.
Overwrites the key if it is holding a value of another data type.
This command can set multiple fields.
Returns the number of fields that were added.


```
hakjctl hashmap set KEY FIELD VALUE [FIELD VALUE ...] [flags]
```

### Examples

```
# Use the default database
hakjctl hashmap set key1 field1 "value1"

# Specify the database to use
hakjctl hashmap set key1 field1 "value1" -d default

# Set multiple fields
hakjctl hashmap set key1 field1 "value1" field2 "value2" field3 "value3"

# Update the values of existing fields
hakjctl hashmap set key1 field1 "value111" field2 "value222" field3 "value333"
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for set
```

### SEE ALSO

* [hakjctl hashmap](hakjctl_hashmap.md)	 - Manage HashMap keys

