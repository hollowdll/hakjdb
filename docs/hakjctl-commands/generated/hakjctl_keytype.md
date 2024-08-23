## hakjctl keytype

Get the data type of a key

### Synopsis

Get the data type of a key. Returns (None) if the key doesn't exist.

```
hakjctl keytype KEY [flags]
```

### Examples

```
# Use the default database
hakjctl keytype mykey

# Specify the database to use
hakjctl keytype mykey -d default
```

### Options

```
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for keytype
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

