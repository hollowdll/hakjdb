# SetString

SetString sets a key to hold a String value. Creates the key if it doesn't exist. Overwrites the key if it is holding a value of another data type.

## Request

SetStringRequest

## Response

SetStringResponse

## Errors

- `NOT_FOUND` - The database to use was not found.
- `INVALID_ARGUMENT` - The key is invalid.
- `FAILED_PRECONDITION` - The database has reached or exceeded the maximum key limit.