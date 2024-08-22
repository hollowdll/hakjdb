## hakjctl delete

Delete keys

### Synopsis

Delete the specified keys and the values they are holding.
Ignores keys that do not exist.
This command can delete multiple keys.
All the keys of a database can be deleted with --all option.
Returns the number of keys that were deleted or OK if all the keys were deleted.


```
hakjctl delete [KEY ...] [flags]
```

### Examples

```
# Use the default database
hakjctl delete key1

# Specify the database to use
hakjctl delete key2 -d default

# Delete multiple keys
hakjctl delete key3 key4 key5

# Delete all the keys
hakjctl delete --all
```

### Options

```
      --all               Delete all the keys of the database that is being used
  -d, --database string   The database to use. If not present, the default database is used
  -h, --help              help for delete
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

