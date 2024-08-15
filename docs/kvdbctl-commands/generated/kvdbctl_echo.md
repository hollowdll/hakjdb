## kvdbctl echo

Test connection

### Synopsis

Test connection to the server. Sends a message to the server and returns the same message back.
Can be useful for verifying that the server is still alive and can process requests.


```
kvdbctl echo [MESSAGE] [flags]
```

### Examples

```
# Send an empty message
kvdbctl echo

# Send message "Hello"
kvdbctl echo "Hello"
```

### Options

```
  -h, --help   help for echo
```

### SEE ALSO

* [kvdbctl](kvdbctl.md)	 - CLI tool for kvdb key-value data store

