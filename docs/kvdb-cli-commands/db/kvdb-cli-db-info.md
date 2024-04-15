# kvdb-cli db info

```sh
kvdb-cli db info [OPTIONS]
```

Shows information about a database.

## Options

- `-n`, `--name` - The name of the database. If not specified, the default database is used.
- `-h`, `--help` - Show help page.

## Returns

- Information about the database.
- Error message if not successful.

## Examples

```sh
kvdb-cli db info
name: default
created_at: 2024-04-15T19:31:13Z00:00
updated_at: 2024-04-15T19:31:13Z00:00
key_count: 0
data_size: 0B
```

Meaning of the fields:

- `name`: Name of the database
- `created_at`: UTC timestamp specifying when the database was created
- `updated_at`: UTC timestamp specifying when the database was last updated
- `key_count`: Number of keys stored in the database
- `data_size`: Size of the stored data in bytes