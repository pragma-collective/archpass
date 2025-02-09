// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/pragma-collective/archpass/ent/event"
	"github.com/pragma-collective/archpass/ent/order"
	"github.com/pragma-collective/archpass/ent/ticket"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TicketID holds the value of the "ticket_id" field.
	TicketID int `json:"ticket_id,omitempty"`
	// EventID holds the value of the "event_id" field.
	EventID int `json:"event_id,omitempty"`
	// WalletAddress holds the value of the "wallet_address" field.
	WalletAddress string `json:"wallet_address,omitempty"`
	// PriceInCents holds the value of the "price_in_cents" field.
	PriceInCents int64 `json:"price_in_cents,omitempty"`
	// CcCheckoutID holds the value of the "cc_checkout_id" field.
	CcCheckoutID uuid.UUID `json:"cc_checkout_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// ModifiedAt holds the value of the "modified_at" field.
	ModifiedAt time.Time `json:"modified_at,omitempty"`
	// PaymentReference holds the value of the "payment_reference" field.
	PaymentReference string `json:"payment_reference,omitempty"`
	// PaymentWalletAddress holds the value of the "payment_wallet_address" field.
	PaymentWalletAddress string `json:"payment_wallet_address,omitempty"`
	// PaymentTransactionHash holds the value of the "payment_transaction_hash" field.
	PaymentTransactionHash string `json:"payment_transaction_hash,omitempty"`
	// TokenID holds the value of the "token_id" field.
	TokenID int `json:"token_id,omitempty"`
	// MintTransactionHash holds the value of the "mint_transaction_hash" field.
	MintTransactionHash string `json:"mint_transaction_hash,omitempty"`
	// TransferTransactionHash holds the value of the "transfer_transaction_hash" field.
	TransferTransactionHash string `json:"transfer_transaction_hash,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges        OrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// Ticket holds the value of the ticket edge.
	Ticket *Ticket `json:"ticket,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) EventOrErr() (*Event, error) {
	if e.Event != nil {
		return e.Event, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: event.Label}
	}
	return nil, &NotLoadedError{edge: "event"}
}

// TicketOrErr returns the Ticket value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) TicketOrErr() (*Ticket, error) {
	if e.Ticket != nil {
		return e.Ticket, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: ticket.Label}
	}
	return nil, &NotLoadedError{edge: "ticket"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldTicketID, order.FieldEventID, order.FieldPriceInCents, order.FieldTokenID:
			values[i] = new(sql.NullInt64)
		case order.FieldWalletAddress, order.FieldPaymentReference, order.FieldPaymentWalletAddress, order.FieldPaymentTransactionHash, order.FieldMintTransactionHash, order.FieldTransferTransactionHash:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldModifiedAt:
			values[i] = new(sql.NullTime)
		case order.FieldCcCheckoutID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldTicketID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_id", values[i])
			} else if value.Valid {
				o.TicketID = int(value.Int64)
			}
		case order.FieldEventID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value.Valid {
				o.EventID = int(value.Int64)
			}
		case order.FieldWalletAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wallet_address", values[i])
			} else if value.Valid {
				o.WalletAddress = value.String
			}
		case order.FieldPriceInCents:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price_in_cents", values[i])
			} else if value.Valid {
				o.PriceInCents = value.Int64
			}
		case order.FieldCcCheckoutID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field cc_checkout_id", values[i])
			} else if value != nil {
				o.CcCheckoutID = *value
			}
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field modified_at", values[i])
			} else if value.Valid {
				o.ModifiedAt = value.Time
			}
		case order.FieldPaymentReference:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_reference", values[i])
			} else if value.Valid {
				o.PaymentReference = value.String
			}
		case order.FieldPaymentWalletAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_wallet_address", values[i])
			} else if value.Valid {
				o.PaymentWalletAddress = value.String
			}
		case order.FieldPaymentTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field payment_transaction_hash", values[i])
			} else if value.Valid {
				o.PaymentTransactionHash = value.String
			}
		case order.FieldTokenID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field token_id", values[i])
			} else if value.Valid {
				o.TokenID = int(value.Int64)
			}
		case order.FieldMintTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mint_transaction_hash", values[i])
			} else if value.Valid {
				o.MintTransactionHash = value.String
			}
		case order.FieldTransferTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transfer_transaction_hash", values[i])
			} else if value.Valid {
				o.TransferTransactionHash = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryEvent queries the "event" edge of the Order entity.
func (o *Order) QueryEvent() *EventQuery {
	return NewOrderClient(o.config).QueryEvent(o)
}

// QueryTicket queries the "ticket" edge of the Order entity.
func (o *Order) QueryTicket() *TicketQuery {
	return NewOrderClient(o.config).QueryTicket(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("ticket_id=")
	builder.WriteString(fmt.Sprintf("%v", o.TicketID))
	builder.WriteString(", ")
	builder.WriteString("event_id=")
	builder.WriteString(fmt.Sprintf("%v", o.EventID))
	builder.WriteString(", ")
	builder.WriteString("wallet_address=")
	builder.WriteString(o.WalletAddress)
	builder.WriteString(", ")
	builder.WriteString("price_in_cents=")
	builder.WriteString(fmt.Sprintf("%v", o.PriceInCents))
	builder.WriteString(", ")
	builder.WriteString("cc_checkout_id=")
	builder.WriteString(fmt.Sprintf("%v", o.CcCheckoutID))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("modified_at=")
	builder.WriteString(o.ModifiedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("payment_reference=")
	builder.WriteString(o.PaymentReference)
	builder.WriteString(", ")
	builder.WriteString("payment_wallet_address=")
	builder.WriteString(o.PaymentWalletAddress)
	builder.WriteString(", ")
	builder.WriteString("payment_transaction_hash=")
	builder.WriteString(o.PaymentTransactionHash)
	builder.WriteString(", ")
	builder.WriteString("token_id=")
	builder.WriteString(fmt.Sprintf("%v", o.TokenID))
	builder.WriteString(", ")
	builder.WriteString("mint_transaction_hash=")
	builder.WriteString(o.MintTransactionHash)
	builder.WriteString(", ")
	builder.WriteString("transfer_transaction_hash=")
	builder.WriteString(o.TransferTransactionHash)
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
