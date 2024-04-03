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

Configurations are stored in a configuration file and can be changed there. This file is created with default configurations if it doesn't exist. The name of the configuration file is `.kvdbserver.yaml` and it is created to the [data directory](#data-directory). Configurations are stored in YAML format.

Here is a list of all configurations with their default values:

```yaml
debug_enabled: false
default_db: default
logfile_enabled: false
port: 12345
```

Meaning of fields:

- `debug_enabled`: Determines if debug mode is enabled. If enabled, debug messages are logged. Can be true or false.
- `default_db`: The name of the default database that is created at server startup.
- `logfile_enabled`: Determines if log file is enabled. If enabled, logs will be written to the log file. Can be true or false.
- `port`: Server's TCP/IP port. Ranges from 1 to 65535.

# Environment variables

It is also possible to change configurations with environment variables. Environment variables override configurations in the configuration file but the values are not saved. They are only read by the server process.

Here is a list of all environment variables:

- `KVDB_PORT`: Server TCP/IP port.
- `KVDB_PASSWORD`: Server password. If not set, password protection is disabled.
- `KVDB_DEBUG_ENABLED`: Determines if debug mode is enabled. If enabled, debug messages are logged. Can be true or false.
- `KVDB_DEFAULT_DB`: The name of the default database that is created at server startup.
- `KVDB_LOGFILE_ENABLED`: Determines if log file is enabled. If enabled, logs will be written to the log file. Can be true or false.

# Data directory

The server has a data directory `data/` that is created to the executable's parent directory if it doesn't exist. Server-specific files are stored in this directory.

Here is a list of files in this directory:
- Configuration file: `.kvdbserver.yaml`
- Log file: `kvdb.log`

# Logs

## Log types

There are five different types of logs:

- `Debug`: Debug messages
- `Info`: Informative messages
- `Error`: Error messages
- `Warning`: Warning messages
- `Fatal`: Fatal error messages.

Debug messages are disabled by default. You can enable them by turning debug mode on. Debug mode can be enabled by modifying the configuration file or with environment variable `KVDB_DEBUG_ENABLED` set to true.

## Log file

By default logs will be written only to the standard error stream (stderr). To write logs to a file, you need to enable the log file. Log file is intended only for debugging purposes as it decreases the server's performance by doing additional writes. 

The log file can be enabled in the configuration file or with environment variable `KVDB_LOGFILE_ENABLED`.

The name of the log file is `kvdb.log`. If log file is enabled, the file is created to the [data directory](#data-directory) at server startup.

# Security

The server can be password protected to prevent unauthorized use. When password protection is enabled, all clients must authenticate using password. By default, password protection is disabled.

Password protection can be enabled by setting password with environment variable `KVDB_PASSWORD`. The password is hashed using bcrypt before storing it in memory. The maximum password size is 72 bytes.

# Default database

When the server starts, it creates an empty default database 'default'. The name of the default database can be changed in the configuration file or with environment variable `KVDB_DEFAULT_DB`.
