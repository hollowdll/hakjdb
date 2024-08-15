## kvdbctl get

Get the value of a String key

### Synopsis

Get the value of a String key. Returns (None) if the key doesn't exist.

```
kvdbctl get KEY [flags]
```

### Examples

```
# Use the default database
kvdbctl get key1

# Specify the database to use
kvdbctl get key1 --database default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for get
```

### SEE ALSO

* [kvdbctl](kvdbctl.md)	 - CLI tool for kvdb key-value data store

