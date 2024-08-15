## kvdbctl keytype

Get the data type of a key

### Synopsis

Get the data type of a key. Returns (None) if the key doesn't exist.

```
kvdbctl keytype KEY [flags]
```

### Examples

```
# Use the default database
kvdbctl keytype mykey

# Specify the database to use
kvdbctl keytype mykey --database default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for keytype
```

### SEE ALSO

* [kvdbctl](kvdbctl.md)	 - CLI tool for kvdb key-value data store

