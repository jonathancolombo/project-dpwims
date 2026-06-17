package utilities

import (
	"fmt"
	"os"
)

const keyHost = "DB_HOST"
const keyPort = "DB_PORT"
const keyUser = "DB_USER"
const keyPassword = "DB_PASSWORD"
const keyDatabaseName = "DB_NAME"
const KeyContentType = "Content-TrainType"
const ValueAppJson = "application/json"

// ConstructDSN construct data source name with some parameters reading env. variables
func ConstructDSN() string {
	host := os.Getenv(keyHost)
	port := os.Getenv(keyPort)
	user := os.Getenv(keyUser)
	password := os.Getenv(keyPassword)
	name := os.Getenv(keyDatabaseName)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)
	return dsn
}
