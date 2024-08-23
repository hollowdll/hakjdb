## hakjctl getkeys

List keys

### Synopsis

List all the keys of a database.

```
hakjctl getkeys [flags]
```

### Examples

```
# Use the default database
hakjctl getkeys

# Specify the database to use
hakjctl getkeys -d default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for getkeys
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

