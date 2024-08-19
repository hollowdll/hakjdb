# API v0

The kvdb server API is implemented with gRPC and defined with Protocol Buffers. HTTP/2 is needed in the connections. Requests are made with RPCs (Remote Procedure Calls). Connections require a gRPC client.

RPCs are grouped to gRPC services that provide related RPCs. The protobuf gRPC service definitions are in the `api/` directory. It contains the currently maintained API version. All services and RPCs are documented there in files that have `.proto` file extension.

[Link to the API directory](../api/)

# Authentication

If authentication is enabled, the client needs to authorize itself with a JWT token. The token is sent in gRPC metadata and needs to be included in all requests, except the authentication endpoint, which returns a token if the provided password is correct.

The gRPC metadata key for the token is `auth-token`. The actual token to send is set as the value of the key.

# Common gRPC Metadata

This section lists common gRPC metadata that all RPCs use:

Request metadata:
- `auth-token`: The JWT token. Used to authorize the client if authentication is enabled. Not needed in the authentication RPC.

Response metadata:
- `api-version`: The gRPC API version. It is of format `0.0.0`.

# gRPC services

## AuthService

- Service name: `AuthService`
- Package: `api.v0.authpb`
- Proto file: `auth.proto`

[Link to the protobuf definitions](../api/v0/authpb/auth.proto)

AuthService provides RPCs related to authentication.

## EchoService

- Service name: `EchoService`
- Package: `api.v0.echopb`
- Proto file: `echo.proto`

[Link to the protobuf definitions](../api/v0/echopb/echo.proto)

EchoService provides connection utility RPCs.
The RPCs in this service can be used to test connections to the server
with minimal network overhead.

## ServerService

- Service name: `ServerService`
- Package: `api.v0.serverpb`
- Proto file: `server.proto`

[Link to the protobuf definitions](../api/v0/serverpb/server.proto)

ServerService provides RPCs related to the server.

## DBService

- Service name: `DBService`
- Package: `api.v0.dbpb`
- Proto file: `db.proto`

[Link to the protobuf definitions](../api/v0/dbpb/db.proto)

DBService provides RPCs related to databases.

## GeneralKVService

- Service name: `GeneralKVService`
- Package: `api.v0.kvpb`
- Proto file: `general_kv.proto`

[Link to the protobuf definitions](../api/v0/kvpb/general_kv.proto)

GeneralKVService provides RPCs related to general key management.

Common gRPC metadata keys for this service's RPCs:
- `database`: The database to use. If omitted, the default database is used.

## StringKVService

- Service name: `StringKVService`
- Package: `api.v0.kvpb`
- Proto file: `string_kv.proto`

[Link to the protobuf definitions](../api/v0/kvpb/string_kv.proto)

StringKVService provides RPCs related to String keys.

Common gRPC metadata keys for this service's RPCs:
- `database`: The database to use. If omitted, the default database is used.

## HashMapKVService

- Service name: `HashMapKVService`
- Package: `api.v0.kvpb`
- Proto file: `hashmap_kv.proto`

[Link to the protobuf definitions](../api/v0/kvpb/hashmap_kv.proto)

HashMapKVService provides RPCs related to HashMap keys.

Common gRPC metadata keys for this service's RPCs:
- `database`: The database to use. If omitted, the default database is used.

