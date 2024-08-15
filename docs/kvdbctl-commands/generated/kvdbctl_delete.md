## kvdbctl delete

Delete keys

### Synopsis

Delete the specified keys and the values they are holding.
Ignores keys that do not exist.
This command can delete multiple keys.
All the keys of a database can be deleted with --all option.
Returns the number of keys that were deleted or OK if all the keys were deleted.


```
kvdbctl delete [KEY ...] [flags]
```

### Examples

```
# Use the default database
kvdbctl delete key1

# Specify the database to use
kvdbctl delete key2 --database default

# Delete multiple keys
kvdbctl delete key3 key4 key5

# Delete all the keys
kvdbctl delete --all
```

### Options

```
      --all               Delete all the keys of the database that is being used
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for delete
```

### SEE ALSO

* [kvdbctl](kvdbctl.md)	 - CLI tool for kvdb key-value data store

