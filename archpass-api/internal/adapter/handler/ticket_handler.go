package handler

import (
	"context"
	"fmt"
	"github.com/pragma-collective/archpass/internal/application/service"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/pragma-collective/archpass/internal/adapter/repository"
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

func CreateTicket(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	var input dto.CreateTicketInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	ctx := context.Background()
	ticketRepo := repository.NewTicketRepository(&ctx)
	eventRepo := repository.NewEventRepository(&ctx)
	irysService, err := service.NewIrysService()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Retrieve event details
	event, err := eventRepo.GetById(input.EventId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Generate ticket image
	imageBuffer, err := generateTicket(event.Name, event.Location, event.Date, "[YOUR NAME]", input.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Upload image.
	key := fmt.Sprintf("%d/%s", input.EventId, "ticket-image.png")
	imageUrl, err := irysService.UploadFile(key, imageBuffer, "image/png")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}
	input.ImageUrl = imageUrl

	ticket, err := ticketRepo.Create(input, claims.Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	metadata := service.MetadataInput{
		Name:        ticket.Name,
		Description: ticket.Description,
		Image:       imageUrl,
		Attributes: []map[string]interface{}{
			{
				"trait_type": "eventName",
				"value":      event.Name,
			},
			{
				"trait_type": "eventLocation",
				"value":      event.Location,
			},
			{
				"trait_type": "eventDate",
				"value":      event.Date,
			},
		},
		Version: "1.0",
	}
	metadataUrl, err := irysService.UploadMetadata(metadata)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	// Update token uri for ticket
	_, err = ticketRepo.UpdateBaseTokenUri(ticket.ID, metadataUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// Update ticket with baseTokenUri
	_, err = ticketRepo.UpdateBaseTokenUri(ticket.ID, metadataUrl)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.DashboardTicket{Id: ticket.ID,
		Name: ticket.Name, Description: ticket.Description, TicketSlug: ticket.TicketSlug,
		CreatedAt: ticket.CreatedAt, ModifiedAt: ticket.UpdatedAt, MintPrice: ticket.MintPrice,
		Quantity: ticket.Quantity,
	})
}

func ListDashboardTickets(c echo.Context) error {
	ctx := context.Background()

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "JWT Token missing or invalid."})
	}

	claims, ok := token.Claims.(*dto.JWTCustomClaims)
	if !ok {
		return c.JSON(http.StatusForbidden, dto.ErrorResponse{Message: "Unable to cast claims."})
	}

	eventIdParam := c.QueryParam("eventId")
	if eventIdParam == "" {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'eventId' parameter is required."})
	}
	id, err := strconv.Atoi(eventIdParam)
	if err != nil || id <= 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'eventId' parameter is invalid"})
	}

	ticketRepo := repository.NewTicketRepository(&ctx)
	tickets, err := ticketRepo.ListByEventId(id, claims.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error()})
	}

	ticketDtos := make([]dto.SimpleDashboardTicket, 0)

	for _, ticket := range tickets {
		ticketDto := dto.SimpleDashboardTicket{
			Id:        ticket.ID,
			Name:      ticket.Name,
			Quantity:  ticket.Quantity,
			MintPrice: ticket.MintPrice,
			ImageUrl:  ticket.ImageURL,
		}
		ticketDtos = append(ticketDtos, ticketDto)
	}

	return c.JSON(http.StatusOK, ticketDtos)
}

func GetEventTicket(c echo.Context) error {
	ctx := context.Background()

	ticketSlug := c.QueryParam("ticketSlug")
	if len(ticketSlug) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'ticketSlug' parameter is required."})
	}

	ticketRepo := repository.NewTicketRepository(&ctx)
	ticket, err := ticketRepo.GetBySlugAndEvent(ticketSlug)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SimplePublicTicket{
		Name:        ticket.Name,
		Description: ticket.Description,
		TicketSlug:  ticket.TicketSlug,
		MintPrice:   ticket.MintPrice,
		EventSlug:   ticket.Edges.Event.EventSlug,
		ImageUrl:    ticket.ImageURL,
		Location:    ticket.Edges.Event.Location,
		Date:        ticket.Edges.Event.Date,
		EventName:   ticket.Edges.Event.Name,
	})
}

func GetAttendeeTickets(c echo.Context) error {
	ctx := context.Background()

	walletAddress := c.QueryParam("walletAddress")
	if len(walletAddress) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'walletAddress' parameter is required."})
	}

	// Todo: validate if wallet address

	ticketRepo := repository.NewTicketRepository(&ctx)
	attendees, err := ticketRepo.GetByAttendee(walletAddress)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	user := attendees[0].Edges.User
	var tickets []dto.AttendeeSimpleTicket

	for _, attendee := range attendees {
		ticket := attendee.Edges.Ticket
		event := ticket.Edges.Event

		if ticket == nil || event == nil {
			continue
		}

		tickets = append(tickets, dto.AttendeeSimpleTicket{
			TicketSlug:      ticket.TicketSlug,
			EventName:       event.Name,
			TokenId:         attendee.TokenID,
			ContractAddress: ticket.ContractAddress,
		})
	}

	attendeeTicketDto := dto.AttendeeTicket{
		DisplayName:   user.DisplayName,
		Bio:           user.Bio,
		WalletAddress: walletAddress,
		Tickets:       tickets,
	}

	return c.JSON(http.StatusOK, attendeeTicketDto)
}
