// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package odbc

import (
	"database/sql/driver"
	"fmt"
	"io"

	"github.com/maurocordioli/odbc/api"
)

type Rows struct {
	os *ODBCStmt
}

func (r *Rows) Columns() []string {
	names := make([]string, len(r.os.Cols))
	for i := 0; i < len(names); i++ {
		names[i] = r.os.Cols[i].Name()
	}

	return names
}

func (r *Rows) Next(dest []driver.Value) error {
	ret := api.SQLFetch(r.os.h)
	if ret == api.SQL_NO_DATA {
		return io.EOF
	}
	if IsError(ret) {
		return NewError("SQLFetch", r.os.h)
	}
	for i := range dest {
		v, err := r.os.Cols[i].Value(r.os.h, i)
		if err != nil {
			return err
		}
		dest[i] = v
	}
	return nil
}

func (r *Rows) Close() error {
	return r.os.closeByRows()
}

func (r *Rows) HasNextResultSet() bool {
	return true
}

func (r *Rows) NextResultSet() error {
	ret := api.SQLMoreResults(r.os.h)
	if ret == api.SQL_NO_DATA {
		return io.EOF
	}
	if IsError(ret) {
		return NewError("SQLMoreResults", r.os.h)
	}

	err := r.os.BindColumns()
	if err != nil {
		return err
	}
	return nil
}

// ColumnTypeDatabaseTypeName return the database system type name.
func (r *Rows) ColumnTypeDatabaseTypeName(index int) string {

	switch x := r.os.Cols[index].(type) {
	case *BindableColumn:

		return sqlTypeString(x.SQLType) + "|" + cTypeString(x.CType)
	case *NonBindableColumn:

		return sqlTypeString(x.SQLType) + "|" + cTypeString(x.CType)
	}
	return ""
}

func (r *Rows) ColumnTypeLength(index int) (length int64, ok bool) {

	switch x := r.os.Cols[index].(type) {
	case *BindableColumn:
		//sizex, _ := sqlTypeLenghtFrmSize(x.SQLType, int64(x.Size))
		//zap.S().Infof("BindableColumn ************************** ColumnTypeLength %v %v %v %v", index, sizex, x.Size, int64(x.Len))
		return sqlTypeLenghtFrmSize(x.SQLType, int64(x.Size))
	case *NonBindableColumn:
		//zap.S().Infof("BindableColumn ************************** ColumnTypeLength %v", index)
		return 0, false
	}
	return 0, false
}

func sqlTypeLenghtFrmSize(ct api.SQLSMALLINT, size int64) (int64, bool) {

	switch ct {

	case api.SQL_UNKNOWN_TYPE:
		return size, false
	case
		api.SQL_CHAR:
		return size - 1, true
	case api.SQL_NUMERIC:
		return size, false
	case api.SQL_DECIMAL:
		return size, false
	case api.SQL_INTEGER:
		return size, false
	case api.SQL_SMALLINT:
		return size, false
	case api.SQL_FLOAT:
		return size, false
	case api.SQL_REAL:
		return size, false
	case api.SQL_DOUBLE:
		return size, false
	case api.SQL_DATETIME:
		return size, false
	//TODO: check if we can choose the sub type based on other info
	//case api.SQL_DATE:
	//	return "DATE"
	case api.SQL_TIME:
		return size, false
	case api.SQL_VARCHAR:
		return size - 1, false
	case api.SQL_TYPE_DATE:
		return size, false
	case api.SQL_TYPE_TIME:
		return size, false
	case api.SQL_TYPE_TIMESTAMP:
		return size, false
	case api.SQL_TIMESTAMP:
		return size, false
	case api.SQL_LONGVARCHAR:
		return size - 1, false
	case api.SQL_BINARY:
		return size, false
	case api.SQL_VARBINARY:
		return size - 1, true
	case api.SQL_LONGVARBINARY:
		return size - 1, true
	case api.SQL_BIGINT:
		return size, false
	case api.SQL_TINYINT:
		return size, false
	case api.SQL_BIT:
		return size, false
	case api.SQL_WCHAR:
		return size/2 - 1, true
	case api.SQL_WVARCHAR:
		return size/2 - 1, true
	case api.SQL_WLONGVARCHAR:
		return size/2 - 1, true
	case api.SQL_GUID:
		return size/2 - 1, false
	case api.SQL_SIGNED_OFFSET:
		return size, false
	case api.SQL_UNSIGNED_OFFSET:
		return size, false
	case api.SQL_SS_XML:
		return size, true
	case api.SQL_SS_TIME2:
		return size, false
	}

	return size, false
}

func sqlTypeString(ct api.SQLSMALLINT) string {

	switch ct {

	case api.SQL_UNKNOWN_TYPE:
		return "SQL_UNKNOWN_TYPE"
	case
		api.SQL_CHAR:
		return "CHAR"
	case api.SQL_NUMERIC:
		return "NUMERIC"
	case api.SQL_DECIMAL:
		return "DECIMAL"
	case api.SQL_INTEGER:
		return "INTEGER"
	case api.SQL_SMALLINT:
		return "SMALLINT"
	case api.SQL_FLOAT:
		return "FLOAT"
	case api.SQL_REAL:
		return "REAL"
	case api.SQL_DOUBLE:
		return "DOUBLE"
	case api.SQL_DATETIME:
		return "DATATIME"
	//TODO: check if we can choose the sub type based on other info
	//case api.SQL_DATE:
	//	return "DATE"
	case api.SQL_TIME:
		return "TIME"
	case api.SQL_VARCHAR:
		return "VARCHAR"
	case api.SQL_TYPE_DATE:
		return "TYPE_DATE"
	case api.SQL_TYPE_TIME:
		return "TYPE_TIME"
	case api.SQL_TYPE_TIMESTAMP:
		return "TYPE_TIMESTAMP"
	case api.SQL_TIMESTAMP:
		return "TIMESTAMP"
	case api.SQL_LONGVARCHAR:
		return "LONGVARCHAR"
	case api.SQL_BINARY:
		return "BINARY"
	case api.SQL_VARBINARY:
		return "VARBINARY"
	case api.SQL_LONGVARBINARY:
		return "LONGVARBINARY"
	case api.SQL_BIGINT:
		return "BIGINT"
	case api.SQL_TINYINT:
		return "TINYINT"
	case api.SQL_BIT:
		return "BIT"
	case api.SQL_WCHAR:
		return "WCHAR"
	case api.SQL_WVARCHAR:
		return "WVARCHAR"
	case api.SQL_WLONGVARCHAR:
		return "WLONGVARCHAR"
	case api.SQL_GUID:
		return "GUID"
	case api.SQL_SIGNED_OFFSET:
		return "SIGNED_OFFSET"
	case api.SQL_UNSIGNED_OFFSET:
		return "UNSIGNE_OFFSET"
	case api.SQL_SS_XML:
		return "SS_XML"
	case api.SQL_SS_TIME2:
		return "SS_TIME2"
	}

	return fmt.Sprintf("UNKNOWN %v", ct)
}

func cTypeString(ct api.SQLSMALLINT) string {
	switch ct {
	case api.SQL_C_CHAR:
		return "SQL_C_CHAR"
	case api.SQL_C_LONG:
		return "SQL_C_LONG"
	case api.SQL_C_SHORT:
		return "SQL_C_SHORT"
	case api.SQL_C_FLOAT:
		return "SQL_C_FLOAT"
	case api.SQL_C_DOUBLE:
		return "SQL_C_DOUBLE"
	case api.SQL_C_NUMERIC:
		return "SQL_C_NUMERIC"
	case api.SQL_C_DATE:
		return "SQL_C_DATE"
	case api.SQL_C_TIME:
		return "SQL_C_TIME"
	case api.SQL_C_TYPE_TIMESTAMP:
		return "SQL_C_TYPE_TIMESTAMP"
	case api.SQL_C_TIMESTAMP:
		return "SQL_C_TIMESTAMP"
	case api.SQL_C_BINARY:
		return "SQL_C_BINARY"
	case api.SQL_C_BIT:
		return "SQL_C_BIT"
	case api.SQL_C_WCHAR:
		return "SQL_C_WCHAR"
	case api.SQL_C_DEFAULT:
		return "SQL_C_DEFAULT"
	case api.SQL_C_SBIGINT:
		return "SQL_C_SBIGINT"
	case api.SQL_C_UBIGINT:
		return "SQL_C_UBIGINT"
	case api.SQL_C_GUID:
		return "SQL_C_GUID"

	}
	return ""
}
