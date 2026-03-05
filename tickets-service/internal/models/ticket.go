package models

// Ticket represents a ticket entity with its attributes.
type Ticket struct {
	UUID       string       `json:"uuid"`
	UserId     int64        `json:"user_id,omitempty"`
	TrainUUID  string       `json:"train_id,omitempty"`
	ScheduleID int64        `json:"schedule_id,omitempty"`
	SeatNumber string       `json:"seat_number,omitempty"`
	Price      float32      `json:"price:omitempty"`
	Status     TicketStatus `json:"status"`
}

// TicketStatus represents the status of a ticket.
type TicketStatus string

const (
	TicketStatusBooked   TicketStatus = "booked"
	TicketStatusCanceled TicketStatus = "canceled"
	TicketStatusUsed     TicketStatus = "used"
)

// UpdateTicket represents the fields that can be updated for a ticket entity.
type UpdateTicket struct {
	UserId     int64        `json:"user_id,omitempty"`
	TrainUUID  string       `json:"train_id,omitempty"`
	ScheduleID int64        `json:"schedule_id,omitempty"`
	SeatNumber string       `json:"seat_number,omitempty"`
	Price      float32      `json:"price,omitempty"`
	Status     TicketStatus `json:"status,omitempty"`
}
