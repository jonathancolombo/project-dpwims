module auth-service

go 1.25.0

require users-service v0.0.0
require project-dpwims/database v0.0.0

require (
	github.com/go-chi/chi/v5 v5.2.5
	github.com/go-sql-driver/mysql v1.9.3
	project-dpwims/database v0.0.0
)

require filippo.io/edwards25519 v1.2.0 // indirect

replace users-service => ../users-service

replace project-dpwims/database => ../database
