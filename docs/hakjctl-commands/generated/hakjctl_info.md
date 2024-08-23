## hakjctl info

Show information about the server

### Synopsis

Show information about the server.
Meaning of the returned fields:

General
- server_version: Version of the server.
- api_version: Version of the API.
- go_version: Version of Go used to compile the server.
- os: Server operating system.
- arch: Architecture which can be 32 or 64 bits.
- process_id: PID of the server process.
- uptime_seconds: Server process uptime in seconds.
- tcp_port: Server TCP/IP port.
- tls_enabled: If TLS is enabled. Yes or no.
- auth_enabled: If authentication is enabled. Yes or no.
- logfile_enabled: If the log file is enabled. Yes or no.
- debug_enabled: If debug mode is enabled. Yes or no.

Databases
- db_count: Number of databases.
- default_db: The default database that the server uses.

Data storage
- total_data_size: Total amount of stored data in bytes.
- total_keys: Total number of keys stored on the server.

Client connections
- client_connections: Number of active client connections.
- max_client_connections: Maximum number of active client connections allowed.

Memory consumption
- memory_alloc: Allocated memory in megabytes.
- memory_total_alloc: Total allocated memory in megabytes.
- memory_sys: Total memory obtained from the OS in megabytes.


```
hakjctl info [flags]
```

### Examples

```
# Show all information
hakjctl info
```

### Options

```
  -h, --help   help for info
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

