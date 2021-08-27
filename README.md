
odbc driver based on  alexbrainman work with some changes for missing functions like


- names of the columns
- lenght of the columns
- data types name with an interim format "SQLDATATYPE!C_DATATYPE" i.e. "VARCHAR|SQL_C_CHAR"
- added ColumnTypeScanType : the driver user  is able to infer scan types for the columns more easily

be aware I have "cut some corners"  [ we need a better mapping for types /sizes ... ]  ( I need some basic functionalities for my use)  
in next revisions I'll add back such corners :-)


useful for example to access old sqlserver2000 or ibm as/400 iseries

odbc driver written in go. Implements database driver interface as used by standard database/sql package. It calls into odbc dll on Windows, and uses cgo (unixODBC) everywhere else.

To get started using odbc, have a look at the [wiki](../../wiki) pages.
