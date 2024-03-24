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

To see the current connection settings, use command:
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

This only changes the default database:
```bash
$ kvdb-cli connect set -d some-db
```
Now commands will use database some-db as the default database. Note that this does not create the database.

## Server information

To show information about the server, use command:
```bash
$ kvdb-cli info
```

Output is something like this:
```bash
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

## Logs

Server logs can be fetched if the server's log file is enabled. If it is not enabled, this command will return an error. This command is intended for debugging purposes.

To get all logs, use command:
```bash
$ kvdb-cli logs
```

Output is something like this:
```bash
2024-02-22T11:16:43.292+02:00 [Info] Some log
2024-02-22T11:16:43.293+02:00 [Info] Another log
```

## Creating a database

In order to store data on the server, you need to create a database. A database is like a namespace for keys. Each key stored in a database is unique to that database. You can't have two keys with the same name in one database.

To create a new database, use command:
```bash
$ kvdb-cli db create -n name-of-your-db
```
- Option -n specifies the name of the database you want to create. The maximum length of a database name is 32 bytes.

Database names are designed to be short. It is recommended to keep them short to minimize memory usage.

The following list shows the allowed characters a database name can contain:

- Letters a-z and A-Z
- Numbers 0-9
- Symbols - and _

## Deleting a database

Deleting a database erases the database and all its data. When a database is deleted, it can no longer be accessed. The database and its data cannot be restored so use this command with caution.

To delete a database, use command:
```bash
$ kvdb-cli db delete -n name-of-your-db
```
- Option -n specifies the name of the database you want to delete. If not specified, the default database is used.

If delete was successful, the command prints the name of the deleted database.

> [!WARNING]
> This command can have dangerous consequences. Use it with caution!

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
- Option -n specifies the name of the database. If not specified, the default database is used.

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

To store a string value, you need to set a key to hold it. If the key already exists with some value, it is overwritten. Creates the key if it doesn't exist.

To set a string value, use command:
```bash
$ kvdb-cli set [key] [value] -d db-name
```
- [key] is the name of the key and [value] is the string value to store.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli set message "Hello World!" -d db0
OK
```
This sets key "message" to hold string "Hello World!" in database db0.

## Get string

To get a string value, you need to retrieve it with the key that is holding the value.

To get a string value, use command:
```bash
$ kvdb-cli get [key] -d db-name
```
- [key] is the name of the key holding the value to retrieve.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli get message -d db0
"Hello World!"
```
This gets the string value that key "message" is holding in database db0.

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
- [key] is the name of the key to delete.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli delete message -d db0
true
```
This deletes key "message" in database db0. If the key exists and was deleted, this returns true.

If the key does not exist, this returns false:
```bash
$ kvdb-cli delete message -d db0
false
```

## Delete all keys

Deleting all the keys of a database removes the keys and the values they are holding. This can be used to remove all the data stored in a database. The database will be blocked until the operation has finished.

To delete all the keys of a database, use command:
```bash
$ kvdb-cli deletekeys -d db-name
```
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli deletekeys
OK
```
This deletes all the keys of the default database.

## Get keys

Getting keys returns a list of keys present in a database. This command is intended for debugging purposes. The database will be blocked until the operation has finished.

To get all the keys of a database, use command:
```bash
$ kvdb-cli getkeys -d db-name
```
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli getkeys
1) key1
2) key2
3) key3
```
This returns all the keys of the default database.

## Set HashMap

To store a HashMap and set fields in it, you need to set a key to hold the HashMap. If the key already exists with some fields, the fields are overwritten with the new values. Creates a new HashMap if the key doesn't exist.

To set HashMap field values, use command:
```bash
$ kvdb-cli hashmap set [key] [field value ...] -d db-name
```
- [key] is the name of the key.
- [field] is the name of a field.
- [value] is the value of a field.
- Option -d specifies the name of the database. If not specified, the default database is used.

The command can be used to set multiple fields.

For example:
```bash
$ kvdb-cli hashmap set key1 name "John" age "35"
OK
```
This sets key "key1" to hold a HashMap with fields "name" and "age" set to their respective values.

## Get HashMap field value

To get a HashMap field value, you need to retrieve it with the key that is holding the HashMap.

To get a HashMap field value, use command:
```bash
$ kvdb-cli hashmap get [key] [field] -d db-name
```
- [key] is the name of the key holding the HashMap.
- [field] is the field whose value should be returned.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli hashmap get key1 name
"John"
```
This gets the value of field "name" in the HashMap that "key1" is holding.

If the key or field doesn't exist, a special value (None) is returned:
```bash
$ kvdb-cli hashmap get key123 name
(None)
```
```bash
$ kvdb-cli hashmap get key1 field123
(None)
```

## Remove fields from a HashMap

To remove fields from a HashMap, use command:
```bash
$ kvdb-cli hashmap delete [key] [field ...] -d db-name
```
- [key] is the name of the key holding the HashMap.
- [field] is a field to be removed.
- Option -d specifies the name of the database. If not specified, the default database is used.

The command can be used to remove multiple fields.

For example:
```bash
$ kvdb-cli hashmap delete key1 field1
1
```
This removes the field "field1" from the HashMap that "key1" is holding. The returned integer is the number of fields that were removed.

If the key doesn't exist, a special value (None) is returned:
```bash
$ kvdb-cli hashmap delete key1234 field1
(None)
```

Fields that do not exist are ignored:
```bash
$ kvdb-cli hashmap delete key1 field12345
0
```

## Get key type

To get the data type of the value a key is holding, use command:
```bash
$ kvdb-cli keytype [key] -d db-name
```
- [key] is the name of the key.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```bash
$ kvdb-cli keytype string-key
"String"
```
This gets the data type of the value that key 'string-key' is holding.

If the key doesn't exist, a special value (None) is returned:
```bash
$ kvdb-cli keytype this-key-does-not-exist
(None)
```
