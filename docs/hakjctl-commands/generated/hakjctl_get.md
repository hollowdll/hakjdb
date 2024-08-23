## hakjctl get

Get the value of a String key

### Synopsis

Get the value of a String key. Returns (None) if the key doesn't exist.

```
hakjctl get KEY [flags]
```

### Examples

```
# Use the default database
hakjctl get key1

# Specify the database to use
hakjctl get key1 -d default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for get
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

