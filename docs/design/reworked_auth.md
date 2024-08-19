# Reworked authentication mechanism design

- Use JWT instead of Bcrypt password hashing in every request (improved performance)
- Passwords are still hashed using Bcrypt and stored in the server's memory
- Hashing happens only when authenticating a connection
- Server creates a default root user that the server password belongs to
- Authentication can be enabled/disabled with a separate config
- The password can still be set with environment variable, but it does not enable authentication by default
- Client creates a gRPC connection and sends the credentials (password) to the server
- Server checks the credentials (Bcrypt hashing and compare with the stored hash)
- If successful, the server generates JWT token and returns it to the client
- Client creates another connection that includes the token so it can make authenticated API requests
- The token is sent in gRPC metadata and can be used with gRPC's PerRPCCredentials
- The server verifies the token and processes the request if the token is valid
