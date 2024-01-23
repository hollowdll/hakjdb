# kvdb-cli

kvdb-cli is a CLI (command line interface) to interact with a kvdb server.

# Configuration file

All configurations are saved to a configuration file and can be changed there. kvdb-cli creates this file with default configurations if it doesn't exist. The name of the configuration file is `.kvdb-cli.json`.

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

By default, kvdb-cli tries to connect to a kvdb server at address 127.0.0.1 in port 12345. 127.0.0.1 is the same as your machine's localhost. Port 12345 is the server's default port.

To see the current connection settings, use the following command:
```bash
$ kvdb-cli connect show
```

It will show something like this:
```bash
Host: localhost
Port: 12345
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
