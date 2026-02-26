module notifications-service

go 1.25.0

require project-dpwims/database v0.0.0

replace project-dpwims/database => ../database

require (
	github.com/eclipse/paho.mqtt.golang v1.5.1
	github.com/go-chi/chi/v5 v5.2.5
	github.com/go-sql-driver/mysql v1.9.3
)

require (
	filippo.io/edwards25519 v1.2.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sync v0.19.0 // indirect
)
