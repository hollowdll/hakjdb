## hakjctl authenticate

Authenticate to the server

### Synopsis

Authenticate to the server using password. If no options provided, prompts the user to enter password.

```
hakjctl authenticate [flags]
```

### Examples

```
# Prompt password
hakjctl authenticate

# Pass password as argument
hakjctl authenticate -p your-password

# Read password from environment variable HAKJCTL_PASSWORD
hakjctl authenticate --password-from-env
```

### Options

```
  -h, --help                help for authenticate
  -p, --password string     The password to use
      --password-from-env   Read password from environment variable HAKJCTL_PASSWORD
```

### SEE ALSO

* [hakjctl](hakjctl.md)	 - CLI tool for HakjDB key-value data store

