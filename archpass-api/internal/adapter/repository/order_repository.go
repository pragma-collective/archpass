package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/pragma-collective/archpass/ent"
	order2 "github.com/pragma-collective/archpass/ent/order"
	"github.com/pragma-collective/archpass/internal/adapter/database"
	"github.com/pragma-collective/archpass/internal/domain/dto"
	"github.com/pragma-collective/archpass/internal/domain/repository"
)

type OrderRepository struct {
	client ent.Client
	ctx    *context.Context
}

func NewOrderRepository(ctx *context.Context) repository.IOrderRepository {
	return &OrderRepository{
		client: *database.EntClient(),
		ctx:    ctx,
	}
}

func (o *OrderRepository) Create(input dto.CreateOrderInput) (ent.Order, error) {
	order, err := o.client.Order.
		Create().
		SetTicketID(input.TicketId).
		SetEventID(input.EventId).
		SetWalletAddress(input.WalletAddress).
		SetPriceInCents(input.Price).
		SetCcCheckoutID(input.CCCheckoutId).
		Save(*o.ctx)
	if err != nil {
		return ent.Order{}, err
	}

	return *order, nil
}

func (o *OrderRepository) UpdateCheckoutId(orderId int, checkoutId uuid.UUID) error {
	_, err := o.client.Order.
		UpdateOneID(orderId).
		SetCcCheckoutID(checkoutId).
		Save(*o.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) UpdateOrderPayment(orderId int, input dto.UpdateOrderInput) error {
	_, err := o.client.Order.
		UpdateOneID(orderId).
		SetPaymentWalletAddress(input.PaymentWalletAddress).
		SetPaymentReference(input.PaymentReference).
		SetPaymentTransactionHash(input.PaymentTransactionHash).
		Save(*o.ctx)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) GetById(orderId int) (ent.Order, error) {
	order, err := o.client.Order.
		Query().
		Where(order2.IDEQ(orderId)).
		WithEvent().
		WithTicket().
		Only(*o.ctx)
	if err != nil {
		return ent.Order{}, err
	}

	return *order, nil
}

func (o *OrderRepository) GetPaidOrders(limit int) (ent.Orders, error) {
	orders, err := o.client.Order.
		Query().
		Where(
			order2.And(
				order2.PaymentReferenceNotNil(),
				order2.PaymentTransactionHashNotNil(),
				order2.TokenIDIsNil(),
			),
		).
		WithEvent().
		WithTicket().
		Limit(limit).
		All(*o.ctx)
	if err != nil {
		return ent.Orders{}, err
	}

	return orders, nil
}

func (o *OrderRepository) UpdateOrderMint(orderId int, tokenId int, mintTx string, transferTx string) error {
	_, err := o.client.Order.
		UpdateOneID(orderId).
		SetTokenID(tokenId).
		SetMintTransactionHash(mintTx).
		SetTransferTransactionHash(transferTx).
		Save(*o.ctx)
	if err != nil {
		return err
	}

	return nil
}
