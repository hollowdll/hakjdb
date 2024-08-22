# hakjserver

hakjserver is the HakjDB server process that listens for requests from clients. It is responsible for managing the server, databases and keys.

# How to use

## Running the binary

After you have the server binary, you can run it with:

```bash
./hakjserver
```
This runs it from the current working directory. Make sure to be in the same directory as the binary.

Now the server should be running with default configurations. The default TCP/IP port is 12345.

## Running with Docker

Another way to run the server is by using Docker. Instructions [here](../README.md#docker)

# Configuration

Configurations are stored in a configuration file and can be changed there. This file is created with default configurations if it doesn't exist. The name of the configuration file is `hakjserver-config.yaml` and it is created to the [data directory](#data-directory). The server tries to find the file in this directory. Configurations are stored in YAML format.

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
auth_enabled: false
auth_token_secret_key: ""
auth_token_ttl: 900
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
- `auth_enabled`: Enable authentication. If enabled, clients need to authenticate.
- `auth_token_secret_key`: Secret key used to sign JWT tokens. Should be long and secure.
- `auth_token_ttl`: JWT token time to live in seconds. Once a JWT token is created, it expires after the number of seconds specified by this.

# Environment variables

It is also possible to change configurations with environment variables. Environment variables override configurations in the configuration file.

Below is a list of all environment variables:

- `HAKJ_PORT`: Server TCP/IP port. Ranges from 1 to 65535.
- `HAKJ_PASSWORD`: Server password. Password is used in authentication.
- `HAKJ_DEBUG_ENABLED`: Enable debug mode. Some features are only enabled in debug mode. Can be true or false.
- `HAKJ_DEFAULT_DB`: The name of the default database that is created at server startup.
- `HAKJ_LOGFILE_ENABLED`: Enable the log file. If enabled, logs will be written to the log file. Can be true or false.
- `HAKJ_TLS_ENABLED`: Enable TLS. If enabled, connections will be encrypted. Can be true or false.
- `HAKJ_TLS_CERT_PATH`: The path to the TLS certificate file.
- `HAKJ_TLS_PRIVATE_KEY_PATH`: The path to the TLS private key.
- `HAKJ_MAX_CLIENT_CONNECTIONS`: The maximum number of active client connections allowed.
- `HAKJ_LOG_LEVEL`: The log level. Can be debug, info, warning, error, or fatal.
- `HAKJ_VERBOSE_LOGS_ENABLED`: Enable verbose logs. Verbose logs show more information and details. Typically used with debug log level for debugging purposes.
- `HAKJ_AUTH_ENABLED`: Enable authentication. If enabled, clients need to authenticate.
- `HAKJ_AUTH_TOKEN_SECRET_KEY`: Secret key used to sign JWT tokens. Should be long and secure.
- `HAKJ_AUTH_TOKEN_TTL`: JWT token time to live in seconds. Once a JWT token is created, it expires after the number of seconds specified by this.

# Port

Port is the TCP/IP network socket port where the gRPC server of the server process listens for requests. It can be configured in the configuration file or with environment variable.

- Config file: `port`
- Env variable: `HAKJ_PORT`

# Debug mode

Below is a list of features that are only enabled in debug mode:

- gRPC server reflection.

Debug mode can be enabled in the configuration file or with environment variable.

- Config file: `debug_enabled`
- Env variable: `HAKJ_DEBUG_ENABLED`

# Data directory

The server has a data directory `data/` that is created to the executable's parent directory if it doesn't exist. Server-specific files are stored in this directory.

Below is a list of files in this directory:
- Configuration file: `hakjserver-config.yaml`
- Log file: `hakjserver.log`

# Logs

## Log levels

There are five different types of logs:

- `Debug`: Debug messages (Level 0)
- `Info`: Informative messages (Level 1)
- `Warning`: Warning messages (Level 2)
- `Error`: Error messages (Level 3)
- `Fatal`: Fatal error messages (Level 4)

You can control what is logged with log level. Debug is the lowest level and Fatal is the highest level. E.g. Debug shows all types of logs and Error shows only error and fatal logs.

Log level can be configured in the configuration file or with environment variable.

- Config file: `log_level`
- Env variable: `HAKJ_LOG_LEVEL`

Debug log level is typically used with verbose logs. Verbose logs show more information about requests. It can be enabled in the configuration file or with environment variable.

- Config file: `verbose_logs_enabled`
- Env variable: `HAKJ_VERBOSE_LOGS_ENABLED`

## Log file

By default logs will be written only to the standard error stream (stderr). To write logs to a file, you need to enable the log file. Log file is intended only for debugging purposes as it decreases the server's performance by doing additional writes. 

The name of the log file is `hakjserver.log`. If log file is enabled, the file is created to the [data directory](#data-directory) at server startup. The server tries to find the file in this directory.

The log file can be enabled in the configuration file or with environment variable.

- Config file: `logfile_enabled`
- Env variable: `HAKJ_LOGFILE_ENABLED`

# Security

## Authentication

The server can enable authentication to prevent unauthorized use. By default, authentication is disabled. When authentication is enabled, all clients must authenticate using password. Authentication process returns a JWT token that the client can use to make authorized requests. Passwords and JWT tokens are not logged in order to improve security.

Authentication can be enabled in the configuration file or with environment variable. By default, the server uses empty password and JWT secret key. These should be set as the default values are insecure. By default, JWT tokens expire 15 minutes after they have been created. This can be changed.

Password can be set with environment variable `KVDB_PASSWORD`. The password is hashed using bcrypt before storing it in memory. Authentication process compares the provided password to this password by hashing it with bcrypt. The maximum password size is 72 bytes.

Enable authentication
- Config file: `auth_enabled`
- Env variable: `HAKJ_AUTH_ENABLED`

JWT token time to live
- Config file: `auth_token_ttl`
- Env variable: `HAKJ_AUTH_TOKEN_TTL`

JWT token secret key
- Config file: `auth_token_secret_key`
- Env variable: `HAKJ_AUTH_TOKEN_SECRET_KEY`

## TLS

Connections can be encrypted with TLS/SSL. The server has native support for this.

When TLS is enabled, all non-TLS connections will be denied. Make sure that the client is connecting with TLS and using the certificate that you configured.

Directory `tls/test-cert/` contains a X.509 certificate and private key for testing purposes. It can be used to test TLS locally. Alternatively, use your own certificate.

TLS can be enabled by modifying the configuration file or with environment variables.

With config file:
```yaml
tls_cert_path: path/to/your/certificate
tls_enabled: true
tls_private_key_path: path/to/your/privatekey
```

Or with `KVDB_TLS_ENABLED`, `KVDB_TLS_CERT_PATH`, and `KVDB_TLS_PRIVATE_KEY_PATH` environment variables.

# Default database

When the server starts, it creates an empty default database 'default'. The name of the default database can be changed in the configuration file or with environment variable.

- Config file: `default_db`
- Env variable: `HAKJ_DEFAULT_DB`

# Connections

The maximum number of active client connections can be limited. Client connections are gRPC clients that are connected to the server. By default, the server allows 1000 active connections. This can be changed in the configuration file or with environment variable.

- Config file: `max_client_connections`
- Env variable: `HAKJ_MAX_CLIENT_CONNECTIONS`
