module trains-service

go 1.25

require project-dpwims/database v0.0.0

replace project-dpwims/database => ../database

require (
	filippo.io/edwards25519 v1.1.1 // indirect
	github.com/go-chi/chi/v5 v5.2.5
	github.com/go-sql-driver/mysql v1.9.3
	github.com/google/uuid v1.6.0
)
