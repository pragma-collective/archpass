package repository

import (
	"github.com/google/uuid"
	"github.com/pragma-collective/archpass/ent"
	"github.com/pragma-collective/archpass/internal/domain/dto"
)

type IOrderRepository interface {
	Create(input dto.CreateOrderInput) (ent.Order, error)
	GetById(orderId int) (ent.Order, error)
	UpdateCheckoutId(orderId int, checkoutId uuid.UUID) error
	UpdateOrderPayment(orderId int, input dto.UpdateOrderInput) error
	GetPaidOrders(limit int) (ent.Orders, error)
	UpdateOrderMint(orderId int, tokenId int, mintTx string, transferTx string) error
}
