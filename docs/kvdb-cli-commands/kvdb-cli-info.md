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
kvdb_version: 0.10.0
go_version: go1.22.0
db_count: 1
total_data_size: 0 B
os: Linux 5.10.102.1-microsoft-standard-WSL2 x86_64
arch: amd64
process_id: 1
uptime_seconds: 54
tcp_port: 12345
default_db: default
memory_alloc: 0.8 MB
memory_total_alloc: 0.8 MB
memory_sys: 7.0 MB
tls_enabled: no
password_enabled: no
logfile_enabled: no
debug_enabled: yes
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
- `default_db`: The default database that the server uses
- `memory_alloc`: Allocated memory in megabytes
- `memory_total_alloc`: Total allocated memory in megabytes
- `memory_sys`: Total memory obtained from the OS in megabytes
- `tls_enabled`: If TLS is enabled. Yes or no.
- `password_enabled`: If password protection is enabled. Yes or no.
- `logfile_enabled`: If the log file is enabled. Yes or no.
- `debug_enabled`: If debug mode is enabled. Yes or no.
