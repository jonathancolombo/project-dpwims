module auth-service

go 1.25.0

require users-service v0.0.0

require project-dpwims/database v0.0.0

require project-dpwims/shared v0.0.0

replace project-dpwims/shared => ../shared

replace users-service => ../users-service

replace project-dpwims/database => ../database

require (
	github.com/go-chi/chi/v5 v5.3.0
	github.com/go-sql-driver/mysql v1.10.0
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
