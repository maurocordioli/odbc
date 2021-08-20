
odbc driver based on  alexbrainman work with some changes like
useful for example to access old sqlserver2000 or ibm as/400 iseries

- names of the columns
- lenght of the columns
- data types name with an interim format "SQLDATATYPE!C_DATATYPE"


odbc driver written in go. Implements database driver interface as used by standard database/sql package. It calls into odbc dll on Windows, and uses cgo (unixODBC) everywhere else.

To get started using odbc, have a look at the [wiki](../../wiki) pages.
