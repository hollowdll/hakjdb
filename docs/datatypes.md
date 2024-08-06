# Data types

Each key holds a value that can have one of the following data types listed below.

## String

String stores sequences of bytes. They are stored as raw bits and are not encoded in any scheme e.g. UTF-8.

String can be used to store e.g. text or binary data.

## HashMap

HashMap stores a collection of field-value pairs. Each field is unique, which means that there cannot be multiple fields with the same name. Fields are stored as sequences of bytes encoded in UTF-8. Values use the data type String.

HashMap can be used to store e.g. a group of properties belonging to some entity. Can be useful if a lot of data belongs to the same key and needs to be stored in a structured and organized way.

The maximum number of fields a HashMap can store is 4,294,967,295 (2^32 - 1).
