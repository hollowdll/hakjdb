# kvdb-cli

kvdb-cli is a CLI (command line interface) to interact with a kvdb server.

# Configuration

Configurations are saved to a configuration file and can be changed there. kvdb-cli creates this file with default configurations if it doesn't exist. The name of the configuration file is `.kvdb-cli.json`. Configurations are saved in JSON format.

Here is a list of all configurations with their default values:

```json
{
  "default_db": "default",
  "host": "localhost",
  "port": 12345
}
```

Meaning of fields:

- `default_db`: Default database to use. Commands use this database by default.
- `host`: Server's address to connect to. Can be hostname or IP address.
- `port`: Server's TCP/IP port. Ranges from 1 to 65535.

# Environment variables

Here is a list of all environment variables:

- `KVDBCLI_PASSWORD`: Provides password to access password protected server.

# Password

If the server is password protected, you can provide password with environment variable `KVDBCLI_PASSWORD`. kvdb-cli reads the value and sends it to the server in every request to perform authentication.

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

By default, kvdb-cli tries to connect to a kvdb server at address 127.0.0.1 in port 12345. 127.0.0.1 is the same as your machine's localhost. Port 12345 is the server's default port. The default database to use is 'default'.

To see the current connection settings, use the following command:
```bash
$ kvdb-cli connect show
```

It will show something like this:
```bash
Host: localhost
Port: 12345
Database: default
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

This only changes the default database to use:
```bash
$ kvdb-cli connect set -d some-db
```
Now commands will use database some-db as the default database. Note that this does not create the database.

## Server information

To show information about the server, use command:
```bash
$ kvdb-cli info
```

This will show something like this:
```bash
kvdb_version: 0.1.0
go_version: go1.21.6
db_count: 0
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

## Creating a database

In order to store data on the server, you need to create a database. A database is like a namespace for keys. Each key stored in a database is unique to that database. You can't have two keys with the same name in one database.

You can create a new database with the following command:
```bash
$ kvdb-cli db create -n name-of-your-db
```
Option -n specifies the name of the database you want to create. The maximum length of a database name is 32 bytes.

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

## Set string

To store a string value, you need to set a key to hold it. If the key already exists with some value, it is overwritten.

To set a string value, use command:
```bash
$ kvdb-cli set [key] [value] -d db-name
```
[key] is the name of the key and [value] is the string value to store. Option -d specifies the name of the database.

For example:
```bash
$ kvdb-cli set message "Hello World!" -d db0
OK
```
This would set key "message" to hold string "Hello World!" in database db0.

Keep in mind that there are some limitations to keys. The maximum length of a key is 1024 bytes. It is recommended to keep keys short to minimize memory usage. Long keys may also impact performance and response time.

## Get string

To get a string value, you need to retrieve it with the key that is holding the value.

To get a string value, use command:
```bash
$ kvdb-cli get [key] -d db-name
```
[key] is the name of the key holding the value to retrieve. Option -d specifies the name of the database.

For example:
```bash
$ kvdb-cli get message -d db0
"Hello World!"
```
This would get the string value that key "message" is holding in database db0.

If the key doesn't exist, a special value (None) is returned:
```bash
$ kvdb-cli get message123 -d db0
(None)
```

## Delete key

Deleting a key removes the key and the value it's holding. Does nothing if the key does not exist.

To delete a key, use command:
```bash
$ kvdb-cli delete [key] -d db-name
```
[key] is the name of the key to delete. Option -d specifies the name of the database.

For example:
```bash
$ kvdb-cli delete message -d db0
true
```
This would delete key "message" in database db0. If the key exists and was deleted, this returns true.

If the key does not exist, this returns false:
```bash
$ kvdb-cli delete message -d db0
false
```
