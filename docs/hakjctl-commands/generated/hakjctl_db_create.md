## hakjctl db create

Create a new database

### Synopsis

Create a new database with the specified name. An optional description can be set with --description option.
Returns the name of the created database.


```
hakjctl db create NAME [flags]
```

### Examples

```
# Create database 'mydb' without description
hakjctl db create mydb

# Create database 'mydb2' with description
hakjctl db create mydb2 -d "Database description."
```

### Options

```
  -d, --description string   Description of the database
  -h, --help                 help for create
```

### SEE ALSO

* [hakjctl db](hakjctl_db.md)	 - Manage databases

