module users-service

go 1.25

require project-dpwims/database v0.0.0

replace project-dpwims/database => ../database

require (
	github.com/DATA-DOG/go-sqlmock v1.5.2
	github.com/go-chi/chi/v5 v5.2.5
	github.com/go-sql-driver/mysql v1.9.3
	github.com/stretchr/testify v1.11.1
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
