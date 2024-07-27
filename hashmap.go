package kvdb

// HashMapFieldValue represents a field value in a HashMap.
type HashMapFieldValue struct {
	// Value is the value the field is holding.
	Value string
	// Ok is true if the field exists. Otherwise false.
	Ok bool
}
