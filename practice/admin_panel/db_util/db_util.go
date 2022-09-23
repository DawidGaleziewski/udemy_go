package dbutil

import (
	"database/sql"
	"fmt"
	"time"
)

type config struct {
	DriverName string
	Username   string
	Password   string
	Host       string
	DBName     string
	Charset    string
	TimeLayout string
}

func (c config) FormatConnectionURL() string {
	// Example format
	// "test_mysql:Test123!@tcp(127.0.0.1:3306)/gosql?charset=utf8"
	return fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v", c.Username, c.Password, c.Host, c.DBName, c.Charset)
}

func (c config) FormatTime(t time.Time) string {
	return t.Format(c.TimeLayout)
}

func (c config) OpenConnection() (*sql.DB, error) {
	return sql.Open(c.DriverName, c.FormatConnectionURL())
}

var Config = config{
	Username:   "test_mysql",
	Password:   "Test123!",
	Host:       "127.0.0.1:3306",
	DBName:     "gosql",
	Charset:    "utf8",
	DriverName: "mysql",
	TimeLayout: "2006-01-02 03:04:05",
}
