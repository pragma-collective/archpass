package handler

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/adapter/repository"
	"github.com/pragma-collective/archpass/internal/domain/dto"
	"net/http"
	"strconv"
)

func PublicGetOrder(c echo.Context) error {
	ctx := context.Background()
	orderIdParam := c.QueryParam("orderId")
	if len(orderIdParam) == 0 {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "'orderId' parameter is required."})
	}

	orderId, err := strconv.Atoi(orderIdParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: "Invalid 'orderId' parameter, must be a valid integer."})
	}

	orderRepo := repository.NewOrderRepository(&ctx)
	order, err := orderRepo.GetById(orderId)
	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, dto.ErrorResponse{Message: "Order not found."})
		}
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	paymentStatus := "pending"
	mintingStatus := "pending"

	if len(order.PaymentTransactionHash) > 0 {
		paymentStatus = "completed"
	}

	if order.TokenID != 0 {
		mintingStatus = "completed"
	}

	priceInDollars := fmt.Sprintf("%.2f", float64(order.PriceInCents)/100)
	orderDto := dto.PublicOrderDetails{
		Id:               order.ID,
		PaymentStatus:    paymentStatus,
		MintingStatus:    mintingStatus,
		Quantity:         1,
		TicketName:       order.Edges.Ticket.Name,
		TicketImageUrl:   order.Edges.Ticket.ImageURL,
		EventName:        order.Edges.Event.Name,
		Price:            priceInDollars,
		PaymentReference: order.PaymentReference,
		PaymentTxHash:    order.PaymentTransactionHash,
		TransferTxHash:   order.TransferTransactionHash,
		MinTxHash:        order.MintTransactionHash,
		TokenId:          order.TokenID,
	}

	return c.JSON(http.StatusOK, orderDto)

}
