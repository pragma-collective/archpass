package handler

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/adapter/repository"
	"github.com/pragma-collective/archpass/internal/application/service"
	"github.com/pragma-collective/archpass/internal/domain/dto"
	"net/http"
	"os"
)

func CreateCheckout(c echo.Context) error {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Something went wrong."})
	}

	archPassWebUrl, ok := os.LookupEnv("ARCHPASS_WEB_URL")
	if !ok {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Something went wrong."})
	}

	var input dto.CreateCheckoutInput
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	ctx := context.Background()
	ticketRepo := repository.NewTicketRepository(&ctx)
	orderRepo := repository.NewOrderRepository(&ctx)
	commerceService, err := service.NewCoinbaseCommerceService()
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	ticket, err := ticketRepo.GetByContractAddress(input.ContractAddress)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Order not found."})
		}
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	order, err := orderRepo.Create(dto.CreateOrderInput{
		TicketId:      ticket.ID,
		EventId:       ticket.EventID,
		WalletAddress: input.WalletAddress,
		Price:         ticket.PriceInCents,
		CCCheckoutId:  uuid.New(), // temporary uuid
	})

	charge, err := commerceService.CreateCharge(service.ChargeInput{
		Name:        ticket.Name,
		Description: ticket.Description,
		PricingType: "fixed_price",
		LocalPrice: service.PriceInput{
			Amount:   fmt.Sprintf("%.2f", float64(ticket.PriceInCents)/100),
			Currency: "USD",
		},
		Metadata: map[string]string{
			"order_id":         fmt.Sprintf("%v", order.ID),
			"contract_address": input.ContractAddress,
			"ticket_id":        fmt.Sprintf("%v", ticket.ID),
		},
		RedirectURL: fmt.Sprintf("%s/orders/success?orderId=%v", archPassWebUrl, order.ID),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: fmt.Sprintf("Failed to create charge: %v", err)})
	}

	err = orderRepo.UpdateCheckoutId(order.ID, charge.Data.ID)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.CheckoutResponse{
		PaymentUrl: charge.Data.HostedURL,
		ChargeID:   charge.Data.ID,
	})
}
