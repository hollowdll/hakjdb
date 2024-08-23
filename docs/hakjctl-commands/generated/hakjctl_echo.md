## hakjctl echo

Test connection

### Synopsis

Test connection to the server. Sends a message to the server and returns the same message back.
Can be useful for verifying that the server is still alive and can process requests.


```
hakjctl echo [MESSAGE] [flags]
```

### Examples

```
# Send an empty message
hakjctl echo

# Send message "Hello"
hakjctl echo "Hello"
```

### Options

```
  -h, --help   help for echo
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

