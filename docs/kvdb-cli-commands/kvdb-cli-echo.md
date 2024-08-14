# kvdb-cli echo

Syntax
```sh
kvdb-cli echo [MESSAGE]
```

Test connection to the server. Sends a message to the server and returns the same message back. Can be useful for verifying that the server is still alive and can process requests.

## Arguments

- `MESSAGE` - The message to be sent.

## Options

- `-h`, `--help` - Show help page.

## Returns

- The same message that was sent to the server.

## Examples

```sh
# send an empty message
kvdb-cli echo
""

# send message "hello"
kvdb-cli echo "hello"
"hello"
```
