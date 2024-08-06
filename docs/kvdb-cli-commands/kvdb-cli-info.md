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
kvdb_version: 0.12.0
go_version: go1.22.4
os: Linux 5.15.153.1-microsoft-standard-WSL2 x86_64
arch: amd64
process_id: 19458
uptime_seconds: 339
tcp_port: 12345
tls_enabled: no
password_enabled: no
logfile_enabled: yes
debug_enabled: yes

** Databases **
db_count: 1
default_db: default

** Data storage **
total_data_size: 0 B
total_keys: 0

** Clients **
client_connections: 1
max_client_connections: 1000

** Memory **
memory_alloc: 1.0 MB
memory_total_alloc: 1.0 MB
memory_sys: 7.1 MB
```

Meaning of the fields:

General information about the server
- `kvdb_version`: Version of kvdb.
- `go_version`: Version of go used to compile the server.
- `os`: Server operating system.
- `arch`: Architecture which can be 32 or 64 bits.
- `process_id`: PID of the server process.
- `uptime_seconds`: Server process uptime in seconds.
- `tcp_port`: Server TCP/IP port.
- `tls_enabled`: If TLS is enabled. Yes or no.
- `password_enabled`: If password protection is enabled. Yes or no.
- `logfile_enabled`: If the log file is enabled. Yes or no.
- `debug_enabled`: If debug mode is enabled. Yes or no.

Information about databases
- `db_count`: Number of databases.
- `default_db`: The default database that the server uses.

Information about data storage
- `total_data_size`: Total amount of stored data in bytes.
- `total_keys`: Total number of keys stored on the server.

Information about client connections
- `client_connections`: Number of active client connections.
- `max_client_connections`: Maximum number of active client connections allowed.

Information about memory consumption
- `memory_alloc`: Allocated memory in megabytes.
- `memory_total_alloc`: Total allocated memory in megabytes.
- `memory_sys`: Total memory obtained from the OS in megabytes.
