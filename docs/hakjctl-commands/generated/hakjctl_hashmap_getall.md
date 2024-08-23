## hakjctl hashmap getall

Get all the fields and values of a HashMap key value

### Synopsis

Get all the fields and values of a HashMap key value. Returns (None) if the key doesn't exist.

```
hakjctl hashmap getall KEY [flags]
```

### Examples

```
# Use the default database
hakjctl hashmap getall key1

# Specify the database to use
hakjctl hashmap getall key1 -d default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for getall
```

### SEE ALSO

* [hakjctl hashmap](hakjctl_hashmap.md)	 - Manage HashMap keys

