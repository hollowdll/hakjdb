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

By default, hakjctl tries to connect to a HakjDB server at localhost in port 12345. Port 12345 is the server's default port. The default database to use is `default`. The commands use this database if database to use is not explicitly specified.

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

## Set String

Store a String key-value:
```sh
hakjctl set key1 "Hello World!"
```
This sets key "key1" to hold the String "Hello World!".

## Get String

Get the value of a String key:
```sh
hakjctl get key1
```
This prints the value of key "key1".

## Delete keys

Deleting a key removes it and its value from the database.

Delete the specified keys if they exist:
```sh
hakjctl delete key1 key2 key3
```

Delete all the keys of a database:
```sh
hakjctl delete --all
```

> [!WARNING]
> This command can have dangerous consequences. Use it with caution!

## List keys

Listing keys returns a list of keys present in a database. This command is intended for debugging purposes. The database will be blocked until the operation has finished.

The following command lists all the keys of a database:
```sh
hakjctl getkeys
```

## Get key type

The data type of a key can be printed.

The following command prints the data type of key "key1":
```sh
hakjctl keytype key1
```

The data type can be "String" or "HashMap".

## Set HashMap

Store a HashMap key-value and set fields in it:
```sh
hakjctl hashmap set key1 field1 value1 field2 value2
```
This sets key "key1" to hold a HashMap value and sets 2 field-value pairs in it.

## Get HashMap field values

It is possible to return the values of the specified fields of a HashMap, or return all the field-value pairs that exist in the HashMap.

Return the values of the specified fields:
```sh
hakjctl hashmap get key1 field1 field2
```
This returns the values of fields "field1" and "field2" in the HashMap stored at key "key1".

Or return all the field-value pairs:
```sh
hakjctl hashmap getall key1
```

If the HashMap doesn't contain any fields, nothing is printed.

## Remove fields from a HashMap

Remove the specified fields from a HashMap:
```sh
hakjctl hashmap delete key1 field1 field2
```
This removes the fields "field1" and "field2" along with the values they are holding from the HashMap stored at key "key1".

