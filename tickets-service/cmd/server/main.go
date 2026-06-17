package main

import (
	"log"
	"net/http"
	"project-dpwims/database"
	util "project-dpwims/shared/utilities"

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
	dsn := util.ConstructDSN()
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

	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)

		chiRouter.Post(urlTickets, ticketHandler.CreateTicket)
		chiRouter.Post(urlPayments, paymentHandler.CreatePayment)
		chiRouter.Get("/tickets/user/{id}", ticketHandler.GetTicketsByUserID)

		chiRouter.Get(urlTicketsID, ticketHandler.GetTicket)
		chiRouter.Get(urlPaymentsID, paymentHandler.GetPayment)
		chiRouter.Delete(urlTicketsID, ticketHandler.DeleteTicket)
	})

	router.Group(func(chiRouter chi.Router) {
		chiRouter.Use(sharedAuth.ValidateJWT)
		chiRouter.Use(sharedAuth.RequireRole("admin"))
		chiRouter.Get(urlTickets, ticketHandler.GetAllTickets)
		chiRouter.Patch(urlTicketsID, ticketHandler.UpdateTicket)

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
