package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pragma-collective/archpass/internal/adapter/repository"
	"github.com/pragma-collective/archpass/internal/domain/dto"
	"net/http"
	"strconv"
)

func CoinbaseWebhook(c echo.Context) error {
	var input dto.WebhookPayload
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	// todo(jhudiel) - Remove afterwards.
	prettyJSON, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v", err)
	} else {
		fmt.Printf("Received webhook payload:\n%s", string(prettyJSON))
	}

	// todo(jhudiel) - Handle charge:failed
	if input.Event.Type != "charge:pending" {
		fmt.Print("Ignoring webhook from coinbase.", input.Event.Type)
		return c.NoContent(http.StatusOK)
	}

	ctx := context.Background()
	orderRepo := repository.NewOrderRepository(&ctx)

	orderId := input.Event.Data.Metadata.OrderId
	orderIdInt, err := strconv.Atoi(orderId)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: "Something went wrong"})
	}
	successEvents := input.Event.Data.Web3Data.SuccessEvents
	var paymentWalletAddress, paymentTransactionHash string

	if len(successEvents) > 0 {
		paymentWalletAddress = successEvents[0].Sender
		paymentTransactionHash = successEvents[0].TxHash
	}
	err = orderRepo.UpdateOrderPayment(orderIdInt, dto.UpdateOrderInput{
		PaymentReference:       input.Event.Data.Code,
		PaymentWalletAddress:   paymentWalletAddress,
		PaymentTransactionHash: paymentTransactionHash,
	})
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusBadRequest, dto.ErrorResponse{Message: err.Error()})
	}

	return c.NoContent(http.StatusOK)
}
