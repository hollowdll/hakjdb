# kvdbserver

kvdbserver is the server process that listens for requests from kvdb clients. It is responsible for managing the server, databases and keys.

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

Configurations are stored in a configuration file and can be changed there. This file is created with default configurations if it doesn't exist. The name of the configuration file is `kvdbserver-config.yaml` and it is created to the [data directory](#data-directory). Configurations are stored in YAML format.

Below is a list of all configurations with their default values:

```yaml
debug_enabled: false
default_db: default
logfile_enabled: false
max_client_connections: 1000
port: 12345
tls_cert_path: ""
tls_enabled: false
tls_private_key_path: ""
log_level: info
verbose_logs_enabled: false
```

Meaning of fields:

- `debug_enabled`: Enable debug mode. Some features are only enabled in debug mode. Can be true or false.
- `default_db`: The name of the default database that is created at server startup.
- `logfile_enabled`: Enable the log file. If enabled, logs will be written to the log file. Can be true or false.
- `max_client_connections`: The maximum number of active client connections allowed.
- `port`: Server's TCP/IP port. Ranges from 1 to 65535.
- `tls_cert_path`: The path to the TLS certificate file.
- `tls_enabled`: Enable TLS. If enabled, connections will be encrypted. Can be true or false.
- `tls_private_key_path`: The path to the TLS private key.
- `log_level`: The log level. Can be debug, info, warning, error, or fatal.
- `verbose_logs_enabled`: Enable verbose logs. Verbose logs show more information and details. Typically used with debug log level for debugging purposes.

# Environment variables

It is also possible to change configurations with environment variables. Environment variables override configurations in the configuration file.

Below is a list of all environment variables:

- `KVDB_PORT`: Server TCP/IP port. Ranges from 1 to 65535.
- `KVDB_PASSWORD`: Server password. If not set, password protection is disabled.
- `KVDB_DEBUG_ENABLED`: Enable debug mode. Some features are only enabled in debug mode. Can be true or false.
- `KVDB_DEFAULT_DB`: The name of the default database that is created at server startup.
- `KVDB_LOGFILE_ENABLED`: Enable the log file. If enabled, logs will be written to the log file. Can be true or false.
- `KVDB_TLS_ENABLED`: Enable TLS. If enabled, connections will be encrypted. Can be true or false.
- `KVDB_TLS_CERT_PATH`: The path to the TLS certificate file.
- `KVDB_TLS_PRIVATE_KEY_PATH`: The path to the TLS private key.
- `KVDB_MAX_CLIENT_CONNECTIONS`: The maximum number of active client connections allowed.
- `KVDB_LOG_LEVEL`: The log level. Can be debug, info, warning, error, or fatal.
- `KVDB_VERBOSE_LOGS_ENABLED`: Enable verbose logs. Verbose logs show more information and details. Typically used with debug log level for debugging purposes.

# Debug mode

Below is a list of features that are only enabled in debug mode:

- gRPC server reflection.

Debug mode can be enabled in the configuration file or with environment variable `KVDB_DEBUG_ENABLED`.

# Data directory

The server has a data directory `data/` that is created to the executable's parent directory if it doesn't exist. Server-specific files are stored in this directory.

Here is a list of files in this directory:
- Configuration file: `kvdbserver-config.yaml`
- Log file: `kvdbserver.log`

# Logs

## Log levels

There are five different types of logs:

- `Debug`: Debug messages (Level 0)
- `Info`: Informative messages (Level 1)
- `Warning`: Warning messages (Level 2)
- `Error`: Error messages (Level 3)
- `Fatal`: Fatal error messages (Level 4)

You can control what is logged with log level. Debug is the lowest level and Fatal is the highest level. E.g. Debug shows all types of logs and Error shows only error and fatal logs.

Log level can be configured in the configuration file or with environment variable `KVDB_LOG_LEVEL`.

## Log file

By default logs will be written only to the standard error stream (stderr). To write logs to a file, you need to enable the log file. Log file is intended only for debugging purposes as it decreases the server's performance by doing additional writes. 

The log file can be enabled in the configuration file or with environment variable `KVDB_LOGFILE_ENABLED`.

The name of the log file is `kvdbserver.log`. If log file is enabled, the file is created to the [data directory](#data-directory) at server startup.

# Security

## Password protection

The server can be password protected to prevent unauthorized use. When password protection is enabled, all clients must authenticate using password. By default, password protection is disabled.

Password protection can be enabled by setting password with environment variable `KVDB_PASSWORD`. The password is hashed using bcrypt before storing it in memory. The maximum password size is 72 bytes.

## TLS

Connections can be encrypted with TLS/SSL. The server has native support for this.

Directory `tls/test-cert/` contains a X.509 certificate and private key for testing purposes. It can be used to test TLS locally. Alternatively, use your own certificate.

TLS can be enabled by modifying the configuration file or with environment variables.

With config file:
```yaml
tls_cert_path: path/to/your/certificate
tls_enabled: true
tls_private_key_path: path/to/your/privatekey
```

Or with `KVDB_TLS_ENABLED`, `KVDB_TLS_CERT_PATH`, and `KVDB_TLS_PRIVATE_KEY_PATH` environment variables. Look at [Environment variables](#environment-variables) for possible values.

When TLS is enabled, all non-TLS connections will be denied. Make sure that the client is connecting with TLS and using the certificate that you configured.

# Default database

When the server starts, it creates an empty default database 'default'. The name of the default database can be changed in the configuration file or with environment variable `KVDB_DEFAULT_DB`.

# Connections

The maximum number of active client connections can be limited. Client connections are gRPC clients that are connected to the server. By default, the server allows 1000 active connections. This can be changed in the configuration file or with environment variable `KVDB_MAX_CLIENT_CONNECTIONS`.
