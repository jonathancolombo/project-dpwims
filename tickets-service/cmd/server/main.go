package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"

	"github.com/go-chi/chi/v5"
)

import (
	"tickets-service/internal/handlers"
	"tickets-service/internal/repositories"
	"tickets-service/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

const urlTickets = "/tickets"
const urlTicketsID = "/tickets/{id}"

// main, runs with this command in the terminal: docker compose --env-file ./env/develop.env up --build
func main() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, errorConnection := database.NewMySQLConnection(dsn)
	if errorConnection != nil {
		log.Fatal(errorConnection)
	}
	repository := repositories.NewMySQLRepositoryTickets(db)
	service := services.NewTicketService(repository)
	handler := handlers.NewTicketHandler(service)
	router := chi.NewRouter()

	router.Post(urlTickets, handler.CreateTicket)
	router.Get(urlTickets, handler.GetAllTickets)
	router.Get(urlTicketsID, handler.GetTicket)
	router.Delete(urlTicketsID, handler.DeleteTicket)
	router.Patch(urlTicketsID, handler.UpdateTicket)

	log.Println("Ticket Service running on port 8083 with url http://localhost:8083")
	errorHttp := http.ListenAndServe(":8083", router)
	if errorHttp != nil {
		return
	}
}
