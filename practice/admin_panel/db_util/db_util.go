package dbutil

import (
	"database/sql"
	"fmt"
	"strings"
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

var Config = config{
	Username:   "test_mysql",
	Password:   "Test123!",
	Host:       "127.0.0.1:3306",
	DBName:     "gosql",
	Charset:    "utf8",
	DriverName: "mysql",
	TimeLayout: "2006-01-02 03:04:05",
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

type Query map[string]string

func (query Query) Where() string {
	sqlQueryString := ""
	var whereConditionsSlice []string
	for rowName, condition := range query {
		whereConditionsSlice = append(whereConditionsSlice, fmt.Sprintf("%v=\"%v\"", rowName, condition))

	}

	whereConditionsQuery := ` WHERE `
	if len(whereConditionsSlice) > 0 {
		whereConditionsQuery += strings.Join(whereConditionsSlice, "AND ")
	}

	if len(whereConditionsQuery) > 0 {
		sqlQueryString += whereConditionsQuery
	}

	return sqlQueryString
}

func (query Query) Select(tableName string) string {
	sqlQueryString := fmt.Sprintf("SELECT * from %v.%v", Config.DBName, tableName) + query.Where()
	return sqlQueryString
}

func (query Query) Delete(tableName string) string {
	sqlQueryString := fmt.Sprintf("DELETE FROM %v.%v", Config.DBName, tableName) + query.Where()
	return sqlQueryString
}
