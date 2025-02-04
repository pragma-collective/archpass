package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("ticket_id"),
		field.Int("event_id"),
		field.String("wallet_address"),
		field.Int64("price_in_cents"),
		field.UUID("cc_checkout_id", uuid.UUID{}),
		field.Time("created_at").
			Default(time.Now),
		field.Time("modified_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("payment_reference").
			Optional(),
		field.String("payment_wallet_address").
			Optional(),
		field.String("payment_transaction_hash").
			Optional(),
		field.Int("token_id").
			Optional(),
		field.String("mint_transaction_hash").
			Optional(),
		field.String("transfer_transaction_hash").
			Optional(),
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("orders").
			Field("event_id").
			Unique().
			Required(),
		edge.From("ticket", Ticket.Type).
			Ref("orders").
			Field("ticket_id").
			Unique().
			Required(),
	}
}
