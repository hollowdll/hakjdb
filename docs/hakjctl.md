# hakjctl

hakjctl is a CLI (command line interface) tool to control and interact with HakjDB servers.

# Configuration

Configurations are stored in a configuration file and can be changed there. hakjctl creates this file with default configurations to the executable's directory if it doesn't exist and tries to find the file there. The name of the configuration file is `hakjctl-config.yaml`. Configurations are stored in YAML format.

Below is a list of all configurations with their default values:

```yaml
command_timeout: 10
default_db: default
host: localhost
port: 12345
tls_cert_path: ""
tls_enabled: false
```

Meaning of fields:

- `command_timeout`: Command timeout in seconds. If a command doesn't get response before the timeout ends, it is cancelled.
- `default_db`: Default database to use. Commands use this database by default.
- `host`: Server's address to connect to. Can be hostname or IP address.
- `port`: Server's TCP/IP port. Ranges from 1 to 65535.
- `tls_cert_path`: The path to the TLS certificate file. The certificate has to be the server's certificate.
- `tls_enabled`: Use TLS when connecting to a kvdb server. Needed if TLS is enabled on the server. Can be true or false.

# Environment variables

Below is a list of all environment variables:

- `HAKJCTL_PASSWORD`: Provides password to the authenticate command if it reads the password from environment variable.

# TLS

If TLS is enabled on the server, you must enable TLS connection. In addition, you must configure the path to the server's certificate file. Currently no native mTLS support.

This can be done by modifying the configuration file:
```yaml
tls_cert_path: path/to/your/certificate
tls_enabled: true
```

Directory `tls/test-cert/` contains a certificate for testing purposes. Use it if the server is configured to use it. Otherwise use your own certificate.

# Commands

Detailed command documentation can be found [here](./hakjctl-commands/generated/).

# How to use

This section gives a general guide on how to use the hakjctl tool. Refer to the command documentation for detailed command usage.

## Help page

Show help page for a command:
```sh
hakjctl help
```
or
```sh
hakjctl --help
```
Every command has its own help page and `--help` command line flag.

## Version

Print the hakjctl version:
```sh
hakjctl --version
```

Print the client, server, and API versions:
```sh
hakjctl version
```
hakjctl needs to be properly configured to access the server to see the server and API versions.

## Connecting to a server

By default, hakjctl tries to connect to a HakjDB server at localhost in port 12345. Port 12345 is the server's default port. The default database to use is `default`.

To see the current connection settings, use command:
```sh
hakjctl connect show
```

It will show something like this:
```sh
host: localhost
port: 12345
database: default
```

To change the connection settings, use the following command with desired flags:
```sh
hakjctl connect set
```

For example, this changes the address and port:
```sh
hakjctl connect set -a some.other.host.or.IP -p 9000
```
Now hakjctl tries to connect to a different address and port.

This only changes the default database:
```sh
hakjctl connect set -d some-db
```
Now commands will use database some-db as the default database. Note that this does not create the database.

## Authentication

If authentication is enabled on the server, users must authenticate to the server using the server's password. Use `hakjctl authenticate` command for this.

The following command prompts the user to enter password:
```sh
hakjctl authenticate
```
The input is not displayed in order to improve security.

You can also pass the password as an argument, but this may be less secure:
```sh
hakjctl authenticate -p your-password
```

It is also possible to read the password from environment variable `HAKJCTL_PASSWORD`:
```sh
hakjctl authenticate --password-from-env
```

After successful authentication, the returned JWT token is stored in a token cache file in the user's cache directory. This directory varies between platforms. The token is read from the file in subsequent commands that need to send requests to the server. If the token expires, users need to authenticate again.

## Test connection

Sometimes it can be useful to test the connection to the server with minimal network overhead to see if it's still possible to connect.

The following command sends a message to the server and receives the same message back. The received message is then printed.
```sh
hakjctl echo "Hello?"
```
or send an empty message
```sh
hakjctl echo
```

## Server information

To show information about the server, use command:
```sh
hakjctl info
```

## Logs

Server logs can be fetched if the server's log file is enabled. If it is not enabled, this command will return an error. This command is intended for debugging purposes.

To get all server logs, use command:
```sh
hakjctl logs
```

## Create a database

In order to store data on the server, you need a database. A database is a namespace for keys. Each key stored in a database is unique to that database. You can't have two keys with the same name in one database. The server creates a default database, but you can also create your databases.

The following command creates a new database called 'my-db' with empty description:
```sh
hakjctl db create my-db
```

## Change the metadata of a database

After creating a database, you can change some metadata of it. Database metadata that can be changed is its name and description.

The following command changes the description of database 'my-db':
```sh
hakjctl db change my-db -d "New description."
```

## Delete a database

Deleting a database erases the database and all its data. When a database is deleted, it can no longer be accessed. The database and its data cannot be restored so use this with caution.

The following command deletes the database 'my-db' if it exists:
```sh
hakjctl db delete my-db
```

> [!WARNING]
> This command can have dangerous consequences. Use it with caution!

## Listing databases

To list all the databases on the server, use command:
```sh
hakjctl db ls
```
This lists the databases by their name.

## Database information

To show information about a database, use command:
```sh
hakjctl db info -n my-db
```
This shows information about database 'my-db'.

## Set string

To store a string value, you need to set a key to hold it. If the key already exists with some value, it is overwritten. Creates the key if it doesn't exist.

To set a string value, use command:
```sh
kvdb-cli set [key] [value] -d db-name
```
- [key] is the name of the key and [value] is the string value to store.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli set message "Hello World!" -d db0
OK
```
This sets key "message" to hold string "Hello World!" in database db0.

## Get string

To get a string value, you need to retrieve it with the key that is holding the value.

To get a string value, use command:
```sh
kvdb-cli get [key] -d db-name
```
- [key] is the name of the key holding the value to retrieve.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli get message -d db0
"Hello World!"
```
This gets the string value that key "message" is holding in database db0.

If the key doesn't exist, a special value (None) is returned:
```sh
kvdb-cli get message123 -d db0
(None)
```

## Delete keys

To delete keys, use command:
```sh
kvdb-cli delete [key ...] -d db-name
```
- [key] is the key to delete.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli delete message -d db0
1
```
This deletes key "message" in database db0. The number of keys that were deleted is returned.

The key doesn't exist anymore so 0 keys were deleted:
```sh
kvdb-cli delete message -d db0
0
```

## Delete all keys

Deleting all the keys of a database removes the keys and the values they are holding. This can be used to remove all the data stored in a database. The database will be blocked until the operation has finished.

To delete all the keys of a database, use command:
```sh
kvdb-cli deletekeys -d db-name
```
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli deletekeys
OK
```
This deletes all the keys of the default database.

## Get keys

Getting keys returns a list of keys present in a database. This command is intended for debugging purposes. The database will be blocked until the operation has finished.

To get all the keys of a database, use command:
```sh
kvdb-cli getkeys -d db-name
```
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli getkeys
1) key1
2) key2
3) key3
```
This returns all the keys of the default database.

## Set HashMap

To store a HashMap and set fields in it, you need to set a key to hold the HashMap. If the key already exists with some fields, the fields are overwritten with the new values. Creates a new HashMap if the key doesn't exist.

To set HashMap field values, use command:
```sh
kvdb-cli hashmap set [key] [field value ...] -d db-name
```
- [key] is the name of the key.
- [field] is the name of a field.
- [value] is the value of a field.
- Option -d specifies the name of the database. If not specified, the default database is used.

The command can be used to set multiple fields.

For example:
```sh
kvdb-cli hashmap set key1 name "John" age "35"
2
```
This sets key "key1" to hold a HashMap with fields "name" and "age" set to their respective values. The returned integer is the number of fields that were added.

## Get HashMap field values

To get HashMap field values, you need to retrieve them with the key that is holding the HashMap.

To get HashMap field values, use command:
```sh
kvdb-cli hashmap get [key] [field ...] -d db-name
```
- [key] is the name of the key holding the HashMap.
- [field] is the field whose value should be returned.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli hashmap get key1 name
1) "name": "John"
```
This gets the value of field "name" in the HashMap that "key1" is holding.

If the key or field doesn't exist, a special value (None) is returned.

Key doesn't exist:
```sh
kvdb-cli hashmap get key123 name
(None)
```

Field doesn't exist:
```sh
kvdb-cli hashmap get key1 field123
1) "field123": (None)
```

## Get all HashMap fields and values

To get all the fields and values of a HashMap, use command:
```sh
kvdb-cli hashmap getall [key] -d db-name
```
- [key] is the name of the key holding the HashMap.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli hashmap getall key1
1) "field1": "value1"
2) "field2": "value2"
3) "field3": "value3"
```
This gets all the fields and values of the HashMap that "key1" is holding.

If the key doesn't exist, a special value (None) is returned:
```sh
kvdb-cli hashmap getall key123
(None)
```

If the HashMap doesn't contain any fields, nothing is printed.

## Remove fields from a HashMap

To remove fields from a HashMap, use command:
```sh
kvdb-cli hashmap delete [key] [field ...] -d db-name
```
- [key] is the name of the key holding the HashMap.
- [field] is a field to be removed.
- Option -d specifies the name of the database. If not specified, the default database is used.

The command can be used to remove multiple fields.

For example:
```sh
kvdb-cli hashmap delete key1 field1
1
```
This removes the field "field1" from the HashMap that "key1" is holding. The returned integer is the number of fields that were removed.

If the key doesn't exist, a special value (None) is returned:
```sh
kvdb-cli hashmap delete key1234 field1
(None)
```

Fields that do not exist are ignored:
```sh
kvdb-cli hashmap delete key1 field12345
0
```

## Get key type

To get the data type of the value a key is holding, use command:
```sh
kvdb-cli keytype [key] -d db-name
```
- [key] is the name of the key.
- Option -d specifies the name of the database. If not specified, the default database is used.

For example:
```sh
kvdb-cli keytype string-key
"String"
```
This gets the data type of the value that key 'string-key' is holding.

If the key doesn't exist, a special value (None) is returned:
```sh
kvdb-cli keytype this-key-does-not-exist
(None)
```

