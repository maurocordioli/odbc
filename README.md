based on  alexbrainman work with some improvemnts like
- names of the columns
- lenght of the columns
-
odbc driver written in go. Implements database driver interface as used by standard database/sql package. It calls into odbc dll on Windows, and uses cgo (unixODBC) everywhere else.

To get started using odbc, have a look at the [wiki](../../wiki) pages.
