//go:build go1.8
// +build go1.8

package odbc

import (
	"context"
	"database/sql/driver"
)

// OpenConnector implements driver.DriverContext.
func (d *Driver) OpenConnector(dsn string) (driver.Connector, error) {

	drv := *d

	return &mcConnector{
		dsn: dsn, drv: drv,
	}, nil
}

type mcConnector struct {
	dsn string
	drv driver.Driver
}

// Connect implements driver.Connector.
// It returns a new mcConn.
// It is called by the database/sql package.
// It is not called directly by users.
// It is not safe for concurrent use.
// TODO: split the logic of mcConn.Open() and mcConn.OpenConnector()
func (ctor mcConnector) Connect(context.Context) (drv driver.Conn, err error) {

	//zap.S().Debugf("Connector Context connect %v", ctor.dsn)
	//cheating a little bit
	//TODO split the logic
	rc, err := ctor.Driver().Open(ctor.dsn)

	return rc, err

}

// Driver returns the underlying Driver of the Connector,
// mainly to maintain compatibility with the Driver method
// on sql.DB.
func (ctor mcConnector) Driver() driver.Driver {
	//zap.S().Debugf("Connector Driver %v", ctor.drv)
	return ctor.drv
}
