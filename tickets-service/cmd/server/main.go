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

const urlPayments = "/payments"
const urlPaymentsID = "/payments/{id}"

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
	repositoryTickets := repositories.NewMySQLRepositoryTickets(db)
	ticketService := services.NewTicketService(repositoryTickets)
	ticketHandler := handlers.NewTicketHandler(ticketService)

	repositoryPayments := repositories.NewMySQLRepositoryPayments(db)
	paymentService := services.NewPaymentService(repositoryPayments)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	router := chi.NewRouter()

	router.Post(urlTickets, ticketHandler.CreateTicket)
	router.Get(urlTickets, ticketHandler.GetAllTickets)
	router.Get(urlTicketsID, ticketHandler.GetTicket)
	router.Delete(urlTicketsID, ticketHandler.DeleteTicket)
	router.Patch(urlTicketsID, ticketHandler.UpdateTicket)

	router.Post(urlPayments, paymentHandler.CreatePayment)
	router.Get(urlPayments, paymentHandler.GetAllPayments)
	router.Get(urlPaymentsID, paymentHandler.GetPayment)
	router.Delete(urlPaymentsID, paymentHandler.DeletePayment)
	router.Patch(urlPaymentsID, paymentHandler.UpdatePayment)

	log.Println("Ticket Service running on port 8083 with url http://localhost:8083")
	errorHttp := http.ListenAndServe(":8083", router)
	if errorHttp != nil {
		return
	}
}
