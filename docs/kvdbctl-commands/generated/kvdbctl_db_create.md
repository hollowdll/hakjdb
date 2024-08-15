## kvdbctl db create

Create a new database

### Synopsis

Create a new database with the specified name. An optional description can be set with --description option.
Returns the name of the created database.


```
kvdbctl db create NAME [flags]
```

### Examples

```
# Create database 'mydb' without description
kvdbctl db create mydb

# Create database 'mydb2' with description
kvdbctl db create mydb2 --description "Database description."
```

### Options

```
  -d, --description string   Description of the database
  -h, --help                 help for create
```

### SEE ALSO

* [kvdbctl db](kvdbctl_db.md)	 - Manage databases

