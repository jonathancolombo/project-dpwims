package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"project-dpwims/database"

	sharedAuth "project-dpwims/shared/auth"

	"github.com/go-chi/chi/v5"

	"tickets-service/internal/handlers"
	"tickets-service/internal/repositories"
	"tickets-service/internal/services"

	_ "github.com/go-sql-driver/mysql"
)

const urlTickets = "/tickets"
const urlTicketsID = "/tickets/{uuid}"

const urlPayments = "/payments"
const urlPaymentsID = "/payments/{uuid}"

// main, runs with this command in the terminal: docker compose --env-file ./env/.env up --build
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
	paymentService := services.NewPaymentService(repositoryPayments, repositoryTickets)
	paymentHandler := handlers.NewPaymentHandler(paymentService)

	router := chi.NewRouter()

	// ROTTE PUBBLICHE
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)

		chiRouter.Post(urlTickets, ticketHandler.CreateTicket)
		chiRouter.Post(urlPayments, paymentHandler.CreatePayment)
		chiRouter.Get("/tickets/user/{id}", ticketHandler.GetTicketsByUserID)

		chiRouter.Get(urlTicketsID, ticketHandler.GetTicket)
		chiRouter.Get(urlPaymentsID, paymentHandler.GetPayment)
	})

	// ROTTE ADMIN
	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Use(sharedAuth.RequireRole("admin"))
		// Ticket admin
		chiRouter.Get(urlTickets, ticketHandler.GetAllTickets)
		chiRouter.Delete(urlTicketsID, ticketHandler.DeleteTicket)
		chiRouter.Patch(urlTicketsID, ticketHandler.UpdateTicket)

		// Payments admin
		chiRouter.Get(urlPayments, paymentHandler.GetAllPayments)
		chiRouter.Delete(urlPaymentsID, paymentHandler.DeletePayment)
		chiRouter.Patch(urlPaymentsID, paymentHandler.UpdatePayment)
	})

	log.Println("Ticket Service running on port 8083 with url http://localhost:8083")
	errorHttp := http.ListenAndServe(":8083", router)
	if errorHttp != nil {
		return
	}
}
