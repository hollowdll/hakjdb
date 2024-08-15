## kvdbctl connect set

Change connection settings

### Synopsis

Change the connection settings used to connect to a server. Only sets those that are specified.

```
kvdbctl connect set [flags]
```

### Examples

```
# Change the host and port
kvdbctl connect set --host 127.0.0.1 --port 9000

# Change only the default database
kvdbctl connect set --database default
```

### Options

```
  -d, --database string   Default database to use
  -h, --help              help for set
  -a, --host string       Host or IP address
  -p, --port uint16       Port number
```

### SEE ALSO

* [kvdbctl connect](kvdbctl_connect.md)	 - Manage connection settings

