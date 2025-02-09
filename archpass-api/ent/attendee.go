// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/pragma-collective/archpass/ent/attendee"
	"github.com/pragma-collective/archpass/ent/event"
	"github.com/pragma-collective/archpass/ent/ticket"
	"github.com/pragma-collective/archpass/ent/user"
)

// Attendee is the model entity for the Attendee schema.
type Attendee struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// EventID holds the value of the "event_id" field.
	EventID int `json:"event_id,omitempty"`
	// TicketID holds the value of the "ticket_id" field.
	TicketID int `json:"ticket_id,omitempty"`
	// TokenID holds the value of the "token_id" field.
	TokenID int `json:"token_id,omitempty"`
	// TransactionHash holds the value of the "transaction_hash" field.
	TransactionHash string `json:"transaction_hash,omitempty"`
	// BlockNumber holds the value of the "block_number" field.
	BlockNumber int64 `json:"block_number,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AttendeeQuery when eager-loading is set.
	Edges        AttendeeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AttendeeEdges holds the relations/edges for other nodes in the graph.
type AttendeeEdges struct {
	// Event holds the value of the event edge.
	Event *Event `json:"event,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Ticket holds the value of the ticket edge.
	Ticket *Ticket `json:"ticket,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// EventOrErr returns the Event value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttendeeEdges) EventOrErr() (*Event, error) {
	if e.Event != nil {
		return e.Event, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: event.Label}
	}
	return nil, &NotLoadedError{edge: "event"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttendeeEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// TicketOrErr returns the Ticket value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AttendeeEdges) TicketOrErr() (*Ticket, error) {
	if e.Ticket != nil {
		return e.Ticket, nil
	} else if e.loadedTypes[2] {
		return nil, &NotFoundError{label: ticket.Label}
	}
	return nil, &NotLoadedError{edge: "ticket"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Attendee) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case attendee.FieldID, attendee.FieldUserID, attendee.FieldEventID, attendee.FieldTicketID, attendee.FieldTokenID, attendee.FieldBlockNumber:
			values[i] = new(sql.NullInt64)
		case attendee.FieldTransactionHash:
			values[i] = new(sql.NullString)
		case attendee.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Attendee fields.
func (a *Attendee) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case attendee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case attendee.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				a.UserID = int(value.Int64)
			}
		case attendee.FieldEventID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field event_id", values[i])
			} else if value.Valid {
				a.EventID = int(value.Int64)
			}
		case attendee.FieldTicketID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ticket_id", values[i])
			} else if value.Valid {
				a.TicketID = int(value.Int64)
			}
		case attendee.FieldTokenID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field token_id", values[i])
			} else if value.Valid {
				a.TokenID = int(value.Int64)
			}
		case attendee.FieldTransactionHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_hash", values[i])
			} else if value.Valid {
				a.TransactionHash = value.String
			}
		case attendee.FieldBlockNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block_number", values[i])
			} else if value.Valid {
				a.BlockNumber = value.Int64
			}
		case attendee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				a.CreatedAt = value.Time
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Attendee.
// This includes values selected through modifiers, order, etc.
func (a *Attendee) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryEvent queries the "event" edge of the Attendee entity.
func (a *Attendee) QueryEvent() *EventQuery {
	return NewAttendeeClient(a.config).QueryEvent(a)
}

// QueryUser queries the "user" edge of the Attendee entity.
func (a *Attendee) QueryUser() *UserQuery {
	return NewAttendeeClient(a.config).QueryUser(a)
}

// QueryTicket queries the "ticket" edge of the Attendee entity.
func (a *Attendee) QueryTicket() *TicketQuery {
	return NewAttendeeClient(a.config).QueryTicket(a)
}

// Update returns a builder for updating this Attendee.
// Note that you need to call Attendee.Unwrap() before calling this method if this Attendee
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Attendee) Update() *AttendeeUpdateOne {
	return NewAttendeeClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Attendee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Attendee) Unwrap() *Attendee {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Attendee is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Attendee) String() string {
	var builder strings.Builder
	builder.WriteString("Attendee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", a.UserID))
	builder.WriteString(", ")
	builder.WriteString("event_id=")
	builder.WriteString(fmt.Sprintf("%v", a.EventID))
	builder.WriteString(", ")
	builder.WriteString("ticket_id=")
	builder.WriteString(fmt.Sprintf("%v", a.TicketID))
	builder.WriteString(", ")
	builder.WriteString("token_id=")
	builder.WriteString(fmt.Sprintf("%v", a.TokenID))
	builder.WriteString(", ")
	builder.WriteString("transaction_hash=")
	builder.WriteString(a.TransactionHash)
	builder.WriteString(", ")
	builder.WriteString("block_number=")
	builder.WriteString(fmt.Sprintf("%v", a.BlockNumber))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(a.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Attendees is a parsable slice of Attendee.
type Attendees []*Attendee
