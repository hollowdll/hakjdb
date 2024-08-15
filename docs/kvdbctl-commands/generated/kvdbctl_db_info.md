## kvdbctl db info

Show information about a database

### Synopsis

Show information about a database. If the database name is not specified, shows information about the default database.

Meaning of the returned fields:
- name: Name of the database
- description: Description of the database
- created_at: UTC timestamp specifying when the database was created
- updated_at: UTC timestamp specifying when the database was last updated
- key_count: Number of keys stored in the database
- data_size: Size of the stored data in bytes


```
kvdbctl db info [flags]
```

### Examples

```
# Use the default database
kvdbctl db info

# Specify the database to use
kvdbctl db info --name "mydb"
```

### Options

```
  -h, --help          help for info
  -n, --name string   The name of the database
```

### SEE ALSO

* [kvdbctl db](kvdbctl_db.md)	 - Manage databases

