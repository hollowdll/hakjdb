## kvdbctl set

Set the value of a String key

### Synopsis

Set the value of a String key.
Creates the key if it doesn't exist.
Overwrites the key if it is holding a value of another data type.


```
kvdbctl set KEY VALUE [flags]
```

### Examples

```
# Use the default database
kvdbctl set key1 "Hello world!"

# Specify the database to use
kvdbctl set key2 "value123" --database default
```

### Options

```
  -h, --help   help for set
```

### SEE ALSO

* [kvdbctl](kvdbctl.md)	 - CLI tool for kvdb key-value data store

