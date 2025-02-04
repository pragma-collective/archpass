package dto

import "github.com/google/uuid"

type CreateOrderInput struct {
	TicketId      int       `json:"ticketId"`
	EventId       int       `json:"eventId"`
	WalletAddress string    `json:"walletAddress"`
	Price         int64     `json:"price"`
	CCCheckoutId  uuid.UUID `json:"ccCheckoutId"`
}

type UpdateOrderInput struct {
	PaymentWalletAddress   string `json:"walletPaymentAddress"`
	PaymentReference       string `json:"paymentReference"`
	PaymentTransactionHash string `json:"paymentTransactionHash"`
}

type PublicOrderDetails struct {
	Id               int    `json:"id"`
	PaymentStatus    string `json:"paymentStatus"`
	MintingStatus    string `json:"mintingStatus"`
	Quantity         int    `json:"quantity"`
	TicketName       string `json:"ticketName"`
	TicketImageUrl   string `json:"ticketImageUrl"`
	EventName        string `json:"eventName"`
	Price            string `json:"price"`
	PaymentReference string `json:"paymentReference"`
	PaymentTxHash    string `json:"paymentTransactionHash"`
	TransferTxHash   string `json:"transferTransactionHash"`
	MinTxHash        string `json:"mintTransactionHash"`
	TokenId          int    `json:"tokenId"`
}
