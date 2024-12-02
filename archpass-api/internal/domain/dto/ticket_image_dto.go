package dto

type CreateTicketImageInput struct {
	EventName     string `json:"eventName"`
	EventLocation string `json:"eventLocation"`
	EventDate     string `json:"eventDate"`
	AttendeeName  string `json:"attendeeName"`
	TicketName    string `json:"ticketName"`
}
