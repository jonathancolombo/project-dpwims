package models

// Ticket represents a ticket entity with its attributes.
type Ticket struct {
	UUID       string       `json:"uuid"`
	UserId     int64        `json:"user_id"`
	TrainUUID  string       `json:"train_id"`
	ScheduleID int64        `json:"schedule_id"`
	SeatNumber int          `json:"seat_number"`
	Price      float32      `json:"price"`
	Status     TicketStatus `json:"status"`
}

// TicketStatus represents the status of a ticket.
type TicketStatus string

const (
	TicketStatusBooked   TicketStatus = "booked"
	TicketStatusCanceled TicketStatus = "canceled"
	TicketStatusUsed     TicketStatus = "used"
)
