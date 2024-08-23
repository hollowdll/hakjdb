## hakjctl connect set

Change connection settings

### Synopsis

Change the connection settings used to connect to a HakjDB server. Only sets those that are specified.

```
hakjctl connect set [flags]
```

### Examples

```
# Change the host and port
hakjctl connect set --host 127.0.0.1 --port 9000

# Change only the default database
hakjctl connect set --database default
```

### Options

```
  -d, --database string   Default database to use
  -h, --help              help for set
  -a, --host string       Host or IP address
  -p, --port uint16       Port number
```

### SEE ALSO

* [hakjctl connect](hakjctl_connect.md)	 - Manage connection settings

