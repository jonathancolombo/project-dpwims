package database

import "database/sql"

func NewMySQLConnection() (*sql.DB, error) {
	var dataSourceName = "root:password@tcp(mysql:3306)/userservice?parseTime=true"
	return sql.Open("mysql", dataSourceName)
}
