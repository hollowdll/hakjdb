# API

The kvdb API uses gRPC. You need a gRPC client to connect to the server. All requests to the server are made with Remote Procedure Calls.

The protobuf gRPC service definitions are in the `proto/kvdbserverpb/` directory. This directory contains all the .proto files needed for a gRPC client.

[Link to the directory](../proto/kvdbserverpb/)

# Authentication

If the server is password protected, the client needs to authenticate with password. The password is sent in gRPC metadata and needs to be included in all requests.

The gRPC metadata key for password is `password`. The actual password to send is set as the value of the key.

# gRPC services

## ServerService

The server service `ServerService` is defined in the `server.proto` file. This service contains RPCs to work with operations related to the kvdb server.

[Link to the protobuf definitions](../proto/kvdbserverpb/server.proto)

Common gRPC metadata for this service's RPCs:
- `password`: The server password if the server is password protected.

RPCs:
- [GetServerInfo](./rpc/server/getserverinfo.md)
- [GetLogs](./rpc/server/getlogs.md)

## DatabaseService

The database service `DatabaseService` is defined in the `db.proto` file. This service contains RPCs to work with operations related to databases.

[Link to the protobuf definitions](../proto/kvdbserverpb/db.proto)

Common gRPC metadata for this service's RPCs:
- `password`: The server password if the server is password protected.

RPCs:
- [CreateDatabase](./rpc/database/createdatabase.md)
- [GetAllDatabases](./rpc/database/getalldatabases.md)
- [GetDatabaseInfo](./rpc/database/getdatabaseinfo.md)
- [DeleteDatabase](./rpc/database/deletedatabase.md)

## StorageService

The storage service `StorageService` is defined in the `storage.proto` file. This service contains RPCs to work with operations related to data storage.

[Link to the protobuf definitions](../proto/kvdbserverpb/storage.proto)

Common gRPC metadata for this service's RPCs:
- `password`: The server password if the server is password protected.
- `database`: The database to use.

RPCs:
- GetTypeOfKey
- SetString
- GetString
- SetHashMap
- GetHashMapFieldValue
- GetAllHashMapFieldsAndValues
- DeleteHashMapFields
- DeleteKey
- DeleteAllKeys
- GetKeys
