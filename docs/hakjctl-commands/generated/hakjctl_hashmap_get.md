## hakjctl hashmap get

Get field values of a HashMap key value

### Synopsis

Get the values of the specified fields of a HashMap key value.
This command can return multiple values. Returns (None) if the key or field doesn't exist.


```
hakjctl hashmap get KEY FIELD [FIELD ...] [flags]
```

### Examples

```
# Use the default database
hakjctl hashmap get key1 field1

# Specify the database to use
hakjctl hashmap get key1 field1 -d default

# Return multiple values
hakjctl hashmap get key1 field1 field2 field3
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for get
```

### SEE ALSO

* [hakjctl hashmap](hakjctl_hashmap.md)	 - Manage HashMap keys

