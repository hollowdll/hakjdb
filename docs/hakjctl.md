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

# Authentication

if authentication is enabled on the server, users must authenticate to the server using the server's password. Use `hakjctl authenticate` command for this.

# TLS

If TLS is enabled on the server, you must enable TLS connection. In addition, you must configure the path to the server's certificate file. Currently no native mTLS support.

This can be done by modifying the configuration file:
```yaml
tls_cert_path: path/to/your/certificate
tls_enabled: true
```

Directory `tls/test-cert/` contains a certificate for testing purposes. Use it if the server is configured to use it. Otherwise use your own certificate.

