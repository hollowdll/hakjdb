# SetHashMap

SetHashMap sets the specified fields and their values in the HashMap stored at a key. If the specified fields exist, they will be overwritten with the new values. Creates the key if it doesn't exist. Overwrites the key if it is holding a value of another data type.

## Request

SetHashMapRequest

## Response

SetHashMapResponse

## Errors

- `NOT_FOUND` - The database to use was not found.
- `INVALID_ARGUMENT` - The key is invalid.
- `FAILED_PRECONDITION` - The database has reached or exceeded the maximum key limit.