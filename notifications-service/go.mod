module notifications-service

go 1.25.0

require project-dpwims/database v0.0.0

require project-dpwims/shared v0.0.0

replace project-dpwims/database => ../database

replace project-dpwims/shared => ../shared

require (
	github.com/eclipse/paho.mqtt.golang v1.5.1
	github.com/go-chi/chi/v5 v5.3.0
	github.com/go-sql-driver/mysql v1.10.0
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.3.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
)
