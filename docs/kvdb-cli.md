# kvdb-cli

kvdb-cli is a CLI (command line interface) to interact with a kvdb server.

# Configuration file

All configurations are saved to a configuration file and can be changed there. kvdb-cli creates this file with default configurations if it doesn't exist. The name of the configuration file is `.kvdb-cli.json`.

# Usage

## How to use

Show help page for a command:
```bash
$ kvdb-cli help
```
or
```bash
$ kvdb-cli --help
```
Every command has its own help page.

## Connecting to a server

By default, kvdb-cli tries to connect to a kvdb server at address 127.0.0.1 in port 12345. 127.0.0.1 is the same as your machine's localhost. Port 12345 is the server's default port.

To see the current connection settings, use the following command:
```bash
$ kvdb-cli connect show
```

It will show something like this:
```bash
Host: localhost
Port: 12345
```

To change the connection settings, use the following command with desired flags:
```bash
$ kvdb-cli connect set
```

For example, this changes the address and port:
```bash
$ kvdb-cli connect set -a some.other.host.or.IP -p 9000
```
Now kvdb-cli tries to connect to a different address and port.

## Server information

To show information about the server, use command:
```bash
$ kvdb-cli info
```

This will show something like this:
```bash
server_version: 0.0.0
go_version: go1.21.0
db_count: 0
total_data_size: 0B
os: Microsoft Windows [Version 10.0.19045.3930]
arch: amd64
process_id: 2484
uptime_seconds: 2646
tcp_port: 12345
```
Meaning of the fields:

- `server_version`: Version of the server
- `go_version`: Version of go used to compile the server
- `db_count`: Number of databases
- `total_data_size`: Total amount of stored data in bytes
- `os`: Server operating system
- `arch`: Architecture which can be 32 or 64 bits
- `process_id`: PID of the server process
- `uptime_seconds`: Server process uptime in seconds
- `tcp_port`: Server TCP/IP port

## Creating a database

In order to store data on the server, you need to create a database. A database is like a namespace for keys. Each key stored in a database is unique to that database. You can't have two keys with the same name in one database.

You can create a new database with the following command:
```bash
$ kvdb-cli db create -n name-of-your-db
```
Option -n specifies the name of the database you want to create. The maximum number of characters a name can have is 32.

Database names are designed to be short. It is recommended to keep them short to minimize memory usage.

The following list shows the allowed characters a database name can contain:

- Letters a-z and A-Z
- Numbers 0-9
- Symbols - and _

## Listing databases

To list all the databases on the server, use command:
```bash
$ kvdb-cli db ls
db0
db1
db2
```
This lists the databases by their name.

## Database information

To show information about a database, use command:
```bash
$ kvdb-cli db info -n name-of-your-db
```
Option -n specifies the name of the database.

Output is something like this:
```bash
name: name-of-your-db
created_at: 2024-01-23T19:31:13Z00:00
updated_at: 2024-01-23T19:31:13Z00:00
key_count: 0
data_size: 0B
```
Meaning of the fields:

- `name`: Name of the database
- `created_at`: UTC timestamp specifying when the database was created
- `updated_at`: UTC timestamp specifying when the database was last updated
- `key_count`: Number of keys stored in the database
- `data_size`: Size of the stored data in bytes
