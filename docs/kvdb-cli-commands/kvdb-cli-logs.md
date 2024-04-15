# kvdb-cli logs

```sh
kvdb-cli logs [OPTIONS]
```

Gets logs from the server if the server's log file is enabled. Currently gets all the logs.

## Options

- `-h`, `--help` - Show help page.

## Returns

- The server logs if the log file is enabled.
- Error message if not successful.

## Examples

```sh
kvdb-cli logs
2024-02-22T11:16:43.292+02:00 [Info] Some log
2024-02-22T11:16:43.293+02:00 [Info] Another log
2024-02-22T11:16:43.294+02:00 [Error] error message
```