package dto

import "github.com/google/uuid"

type CreateCheckoutInput struct {
	ContractAddress string `json:"contractAddress"`
	WalletAddress   string `json:"walletAddress"`
}

type CheckoutResponse struct {
	PaymentUrl string    `json:"paymentUrl"`
	ChargeID   uuid.UUID `json:"chargeId"`
}
