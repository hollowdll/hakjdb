## kvdbctl getkeys

List keys

### Synopsis

List all the keys of a database.

```
kvdbctl getkeys [flags]
```

### Examples

```
# Use the default database
kvdbctl getkeys

# Specify the database to use
kvdbctl getkeys --database default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for getkeys
```

### SEE ALSO

* [kvdbctl](kvdbctl.md)	 - CLI tool for kvdb key-value data store

