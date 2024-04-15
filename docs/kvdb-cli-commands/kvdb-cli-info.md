# kvdb-cli info

```sh
kvdb-cli info [OPTIONS]
```

Shows information about the kvdb server.

## Options

- `-h`, `--help` - Show help page.

## Returns

- Information about the server.
- Error message if not successful.

## Examples

```sh
kvdb-cli info
kvdb_version: 0.1.0
go_version: go1.21.6
db_count: 1
total_data_size: 0B
os: Linux 5.10.102.1-microsoft-standard-WSL2 x86_64
arch: amd64
process_id: 1
uptime_seconds: 54
tcp_port: 12345
```

Meaning of the fields:

- `kvdb_version`: Version of kvdb
- `go_version`: Version of go used to compile the server
- `db_count`: Number of databases
- `total_data_size`: Total amount of stored data in bytes
- `os`: Server operating system
- `arch`: Architecture which can be 32 or 64 bits
- `process_id`: PID of the server process
- `uptime_seconds`: Server process uptime in seconds
- `tcp_port`: Server TCP/IP port