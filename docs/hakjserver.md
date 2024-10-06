# hakjserver

hakjserver is the HakjDB server process that listens for requests from clients. It is responsible for managing the server, databases and keys.

# Directory

It is recommended to place the binary in a directory called `hakjserver`. The server generates some files relative to the binary's parent directory, so they can be easily found there.

# How to use

## Running the binary

After you have the server binary, you can run it with:

```bash
./hakjserver
```
This runs it from the current working directory. Make sure to be in the same directory as the binary.

Now the server should be running with default configurations. The default TCP/IP port is 12345.

You can pass `--help` flag to see all the different command line flags
```sh
./hakjserver --help
```

To change the port for example, you can use
```sh
./hakjserver --port 8080
```
Now the server listens for requests in port 8080.

## Running with Docker

Another way to run the server is by using Docker. Instructions [here](../README.md#docker)

# Configuration

Configurations are stored in a configuration file and can be changed there. This file is created with default configurations if it doesn't exist. The name of the configuration file is `hakjserver-config.yaml` and it is created to the [data directory](#data-directory). The server tries to find the file in this directory. Configurations are stored in YAML format.

It is also possible to change configurations with environment variables and command line flags. Environment variables override configurations in the configuration file. Command line flags override both sources.

See all configuration options [here](./configuration.md).

Below is a template configuration file with default values:

```yaml
auth_enabled: false
auth_token_secret_key: ""
auth_token_ttl: 900
debug_enabled: false
default_db: default
log_level: info
logfile_enabled: false
max_client_connections: 1000
password: ""
port: 12345
tls_ca_cert_path: ""
tls_cert_path: ""
tls_client_cert_auth_enabled: false
tls_enabled: false
tls_private_key_path: ""
verbose_logs_enabled: false
```

# Runtime configurable features

The following features can be changed at runtime without restarting the server:

- Log file enabled
- Verbose logs enabled
- Authentication enabled
- Log level
- Password
- JWT token TTL
- JWT token secret key
- Max active client connections

# Port

Port is the TCP/IP network socket port where the gRPC server of the server process listens for requests.

# Debug mode

Below is a list of features that are only enabled in debug mode:
- gRPC server reflection.

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

Debug log level is typically used with verbose logs. Verbose logs can be enabled to show more information about requests.

## Log file

By default logs will be written only to the standard error stream (stderr). To write logs to a file, you need to enable the log file. Log file is intended only for debugging purposes as it decreases the server's performance by doing additional writes. 

The name of the log file is `hakjserver.log`. If log file is enabled, the file is created to the [data directory](#data-directory) at server startup. The server tries to find the file in this directory.

# Security

## Authentication

The server can enable authentication to prevent unauthorized use. By default, authentication is disabled. When authentication is enabled, all clients must authenticate using password. Authentication process returns a JWT token that the client can use to make authorized requests. Passwords and JWT tokens are not logged in order to improve security.

By default, the server uses empty password and JWT secret key. These should be set as the default values are insecure. By default, JWT tokens expire 15 minutes after they have been created. This can be changed.

The password is hashed using bcrypt before storing it in memory. Authentication process compares the provided password to this password by hashing it with bcrypt. The maximum password size is 72 bytes.

## TLS

Connections can be encrypted with TLS/SSL. The server has native support for this.

When TLS is enabled, all non-TLS connections will be denied. Make sure that the client is connecting with TLS and proper configurations.

Directory `tls/test-cert/` in the repository root contains self-signed X.509 key pairs signed by a CA for testing and development purposes. It also contains self-signed CA. They can be used to test TLS locally. Alternatively, use your own certificates.

Server certificate and private key need to be signed by a CA. The CA certificate is needed when using client certificate authentication. Client certificates are verified against this CA.

Example with config file:
```yaml
tls_enabled: true
tls_cert_path: /path/to/your/certificate
tls_private_key_path: /path/to/your/privatekey
```

It is also possible to enable mutual TLS (mTLS) for client certificate authentication. When this is enabled, clients need to provide a client certificate signed by the server's root CA when establishing the TLS connection.

Example with config file:
```yaml
tls_client_cert_auth_enabled: true
tls_ca_cert_path: /path/to/your/ca-certificate
```

Or with `HAKJ_TLS_CLIENT_CERT_AUTH_ENABLED` and `HAKJ_TLS_CA_CERT_PATH` environment variables.

# Default database

When the server starts, it creates an empty default database 'default'. The name of this default database can be changed.

# Connections

The maximum number of active client connections can be limited. Client connections are gRPC clients that are connected to the server. By default, the server allows 1000 active connections. This can be changed.
