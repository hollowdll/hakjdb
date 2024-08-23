## hakjctl set

Set the value of a String key

### Synopsis

Set the value of a String key.
Creates the key if it doesn't exist.
Overwrites the key if it is holding a value of another data type.


```
hakjctl set KEY VALUE [flags]
```

### Examples

```
# Use the default database
hakjctl set key1 "Hello world!"

# Specify the database to use
hakjctl set key2 "value123" -d default
```

### Options

```
  -h, --help   help for set
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

