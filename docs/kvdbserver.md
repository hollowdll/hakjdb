# kvdbserver

kvdbserver is the server process that listens to requests from kvdb clients. It is responsible for managing the server, databases and data.

The server API is implemented with gRPC and therefore requires HTTP/2. gRPC uses Protocol Buffers data format. Requests are made with RPCs (Remote Procedure Calls) and need a gRPC client.

# How to use

## Running the binary

After you have the server binary, you can run it with:

```bash
./kvdbserver
```
This runs it from the current working directory. Make sure to be in the same directory as the binary.

Now the server should be running with default configurations. The default TCP/IP port is 12345.

## Running with Docker

Another way to run the server is by using Docker. Instructions [here](../README.md#docker)

# Configuration

Configurations are saved to a configuration file and can be changed there. This file is created with default configurations if it doesn't exist. The name of the configuration file is `.kvdbserver.json` and it is created to the data directory. Configurations are saved in JSON format.

Here is a list of all configurations with their default values:

```json
{
  "debug_enabled": false,
  "default_db": "default",
  "port": 12345
}
```

Meaning of fields:

- `debug_enabled`: Specifies if debug mode is enabled. If enabled, debug messages are logged. Can be true or false.
- `default_db`: The name of the default database that is created at server startup.
- `port`: Server's TCP/IP port. Ranges from 1 to 65535.

# Environment variables

It is also possible to change configurations with environment variables. Environment variables override configurations in the configuration file but the values are not saved. They are only read by the server process.

Here is a list of all environment variables:

- `KVDB_PORT`: Server TCP/IP port.
- `KVDB_PASSWORD`: Server password. If not set, password protection is disabled.
- `KVDB_DEBUG_ENABLED`: Controls whether debug mode is enabled. If true, debug messages are logged.
- `KVDB_DEFAULT_DB`: The name of the default database that is created at server startup.

# Data directory

The server has a data directory `data/` that is created to the executable's parent directory if it doesn't exist. Server specific data is saved to this directory. Currently only configuration file gets saved here.

# Logs

There are five different types of logs:

- `Debug`: Debug messages
- `Info`: Informative messages
- `Error`: Error messages
- `Warning`: Warning messages
- `Fatal`: Fatal error messages.

Debug messages are disabled by default. You can enable them by turning debug mode on. Debug mode can be enabled by modifying the configuration file or with environment variable `KVDB_DEBUG_ENABLED` set to true.

# Security

The server can be password protected to prevent unauthorized use. When password protection is enabled, all clients must authenticate using password. By default, password protection is disabled.

Password protection can be enabled by setting password with environment variable `KVDB_PASSWORD`. The password is hashed using bcrypt and stored in the server's memory.

# Default database

When the server starts, it creates an empty default database 'default'. The name of the default database can be changed in the configuration file or with environment variable `KVDB_DEFAULT_DB`.
