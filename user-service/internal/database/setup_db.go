package database

import (
	"database/sql"
	"log"
	"os"
	"strings"
)

const initDbPath = "internal/database/scripts/init.sql"
const testDbPath = "internal/database/scripts/test.sql"

func RunInitScripts(database *sql.DB) {
	runSQLFile(database, initDbPath)
	runSQLFile(database, testDbPath)
}

// runSQLFile read a file .sql into a path and execute it
func runSQLFile(db *sql.DB, path string) {
	content, errorPath := os.ReadFile(path)
	if errorPath != nil {
		log.Fatalf("cannot read %s: %v", path, errorPath)
	}
	var separator = ";"
	queries := strings.Split(string(content), separator)

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, errorExecution := db.Exec(query)
		if errorExecution != nil {
			log.Fatalf("error executing query %query: %v", query, errorExecution)
		}
	}
}
