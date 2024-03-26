# Databases

Databases are like namespaces for grouping keys. They store the data in the server's memory. There is no limit to how many databases you can have.

# Naming

Databases have names that identify them. You cannot have two databases with the same name. The maximum length of a database name is 64 bytes. It is recommended to keep the names short to minimize memory usage.

The following list shows the allowed characters in a database name:

- Letters a-z and A-Z
- Numbers 0-9
- Symbols - and _

# Keyspace

In a database, each key is unique. That means you cannot have two keys with the same name in the same database. The maximum length of a key is 1024 bytes. It is recommended to keep keys short to minimize memory usage. Long keys may also impact performance and response time. The maximum number of keys a database can have is 4,294,967,295 (2^32 - 1).
