# API

The kvdb API uses gRPC. You need a gRPC client to connect to the server. All requests to the server are made with Remote Procedure Calls.

The protobuf gRPC service definitions are in the `proto/kvdbserver/` directory. This directory contains all the .proto files needed to build the gRPC client.

[Link to the directory](../proto/kvdbserver/)

# gRPC services

## ServerService

The server service `ServerService` is defined in the `server.proto` file. This service contains RPCs to work with operations related to the kvdb server.

RPCs:
- GetServerInfo
- GetLogs

[Link to the protobuf definitions](../proto/kvdbserver/server.proto)

## DatabaseService

The database service `DatabaseService` is defined in the `db.proto` file. This service contains RPCs to work with operations related to databases.

RPCs:
- CreateDatabase
- GetAllDatabases
- GetDatabaseInfo
- DeleteDatabase

[Link to the protobuf definitions](../proto/kvdbserver/db.proto)

## StorageService

The storage service `StorageService` is defined in the `storage.proto` file. This service contains RPCs to work with operations related to data storage.

RPCs:
- SetString
- GetString
- DeleteKey
- DeleteAllKeys
- GetKeys

[Link to the protobuf definitions](../proto/kvdbserver/storage.proto)