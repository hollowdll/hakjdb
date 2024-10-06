# Server configuration

HakjDB can be configured with a configuration file, environment variables and command line flags.

The configurations are processed in the following order:
1. Command line flags
2. Environment variables
3. Configuration file

This means that if for example port is configured with all the sources, it will process the value set by the command line flag when the configurations are loaded.

Below is a list of all configuration options

# Configuration options

Authentication enabled
- Enable authentication. If enabled, clients need to authenticate.
- Default: false
- Config file: auth_enabled
- Env variable: HAKJ_AUTH_ENABLED
- Command line flag: --auth-enabled

Debug enabled
- Enable debug mode. Some features are only enabled in debug mode. Can be true or false.
- Default: false
- Config file: debug_enabled
- Env variable: HAKJ_DEBUG_ENABLED
- Command line flag: --debug-enabled

Default database
- The name of the default database that is created at server startup.
- Default: "default"
- Config file: default_db
- Env variable: HAKJ_DEFAULT_DB
- Command line flag: --default-db

JWT token secret key
- Secret key used to sign JWT tokens. Should be long and secure.
- Default: ""
- Config file: auth_token_secret_key
- Env variable: HAKJ_AUTH_TOKEN_SECRET_KEY
- Command line flag: --auth-token-secret-key

JWT token TTL
- JWT token time to live in seconds. Once a JWT token is created, it expires after the number of seconds specified by this.
- Default: 900
- Config file: auth_token_ttl
- Env variable: HAKJ_AUTH_TOKEN_TTL
- Command line flag: --auth-token-ttl

Log level
- The log level. Can be debug, info, warning, error, or fatal.
- Default: "info"
- Config file: log_level
- Env variable: HAKJ_LOG_LEVEL
- Command line flag: --log-level

Log file enabled
- Enable the log file. If enabled, logs will be written to the log file. Can be true or false.
- Default: false
- Config file: logfile_enabled
- Env variable: HAKJ_LOGFILE_ENABLED
- Command line flag: --logfile-enabled

Max client connections
- The maximum number of active client connections allowed.
- Default: 1000
- Config file: max_client_connections
- Env variable: HAKJ_MAX_CLIENT_CONNECTIONS
- Command line flag: --max-client-connections

Port
- Server TCP/IP port. Ranges from 1 to 65535.
- Default: 12345
- Config file: port
- Env variable: HAKJ_PORT
- Command line flag: --port

TLS CA cert path
- Path to the TLS CA certificate file.
- Default: ""
- Config file: tls_ca_cert_path
- Env variable: HAKJ_TLS_CA_CERT_PATH
- Command line flag: --tls-ca-cert-path

TLS cert path
- Path to the TLS certificate file.
- Default: ""
- Config file: tls_cert_path
- Env variable: HAKJ_TLS_CERT_PATH
- Command line flag: --tls-cert-path

TLS client cert auth enabled
- Enable TLS client certificate authentication (mTLS). If enabled, clients need to provide a client certificate signed by the server's root CA. This has no effect if TLS is not enabled. Can be true or false.
- Default: false
- Config file: tls_client_cert_auth_enabled
- Env variable: HAKJ_TLS_CLIENT_CERT_AUTH_ENABLED
- Command line flag: --tls-client-cert-auth-enabled

TLS enabled
- Enable TLS. If enabled, connections will be encrypted. Can be true or false.
- Default: false
- Config file: tls_enabled
- Env variable: HAKJ_TLS_ENABLED
- Command line flag: --tls-enabled

TLS key path
- Path to the TLS private key file.
- Default: ""
- Config file: tls_private_key_path
- Env variable: HAKJ_TLS_PRIVATE_KEY_PATH
- Command line flag: --tls-private-key-path

Verbose logs enabled
- Enable TLS. Enable verbose logs. Verbose logs show more information and details. Typically used with debug log level for debugging purposes.
- Default: false
- Config file: verbose_logs_enabled
- Env variable: HAKJ_VERBOSE_LOGS_ENABLED
- Command line flag: --verbose-logs-enabled
