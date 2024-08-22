## hakjctl db change

Change a database

### Synopsis

Change the metadata of the specified database. Returns the name of the changed database.

```
hakjctl db change NAME [flags]
```

### Examples

```
# Change the name of database 'mydb'
hakjctl db change mydb -n "my-new-db"

# Change the description of database 'mydb'
hakjctl db change mydb -d "New database description."

# Change the name and description of database 'mydb'
hakjctl db change mydb -n "my-new-db" -d "New database description."
```

### Options

```
  -d, --description string   New description of the database
  -h, --help                 help for change
  -n, --name string          New name of the database
```

### SEE ALSO

* [hakjctl db](hakjctl_db.md)	 - Manage databases

