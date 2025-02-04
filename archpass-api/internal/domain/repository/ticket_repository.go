package repository

import (
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

type ITicketRepository interface {
	Create(input dto.CreateTicketInput, userId int) (ent.Ticket, error)
	ListByEventId(eventId int, userId int) (ent.Tickets, error)
	GetBySlugAndEvent(ticketSlug string) (ent.Ticket, error)
	GetByAttendee(walletAddress string) (ent.Attendees, error)
	GetByContractAddress(walletAddress string) (ent.Ticket, error)
	GetByTicketHash(hash string) (ent.Ticket, error)
	UpdateBaseTokenUri(ticketId int, baseTokenUri string) (ent.Ticket, error)
}
