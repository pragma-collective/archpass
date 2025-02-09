// Code generated by ent, DO NOT EDIT.

package event

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the event type in the database.
	Label = "event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldEventSlug holds the string denoting the event_slug field in the database.
	FieldEventSlug = "event_slug"
	// FieldStartDate holds the string denoting the start_date field in the database.
	FieldStartDate = "start_date"
	// FieldEndDate holds the string denoting the end_date field in the database.
	FieldEndDate = "end_date"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldEventHash holds the string denoting the event_hash field in the database.
	FieldEventHash = "event_hash"
	// FieldContractAddress holds the string denoting the contract_address field in the database.
	FieldContractAddress = "contract_address"
	// FieldTransactionHash holds the string denoting the transaction_hash field in the database.
	FieldTransactionHash = "transaction_hash"
	// FieldBlockNumber holds the string denoting the block_number field in the database.
	FieldBlockNumber = "block_number"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldModifiedAt holds the string denoting the modified_at field in the database.
	FieldModifiedAt = "modified_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeTickets holds the string denoting the tickets edge name in mutations.
	EdgeTickets = "tickets"
	// EdgeAttendees holds the string denoting the attendees edge name in mutations.
	EdgeAttendees = "attendees"
	// EdgeOrders holds the string denoting the orders edge name in mutations.
	EdgeOrders = "orders"
	// Table holds the table name of the event in the database.
	Table = "events"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "events"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// TicketsTable is the table that holds the tickets relation/edge.
	TicketsTable = "tickets"
	// TicketsInverseTable is the table name for the Ticket entity.
	// It exists in this package in order to avoid circular dependency with the "ticket" package.
	TicketsInverseTable = "tickets"
	// TicketsColumn is the table column denoting the tickets relation/edge.
	TicketsColumn = "event_id"
	// AttendeesTable is the table that holds the attendees relation/edge.
	AttendeesTable = "attendees"
	// AttendeesInverseTable is the table name for the Attendee entity.
	// It exists in this package in order to avoid circular dependency with the "attendee" package.
	AttendeesInverseTable = "attendees"
	// AttendeesColumn is the table column denoting the attendees relation/edge.
	AttendeesColumn = "event_id"
	// OrdersTable is the table that holds the orders relation/edge.
	OrdersTable = "orders"
	// OrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrdersInverseTable = "orders"
	// OrdersColumn is the table column denoting the orders relation/edge.
	OrdersColumn = "event_id"
)

// Columns holds all SQL columns for event fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldEventSlug,
	FieldStartDate,
	FieldEndDate,
	FieldDate,
	FieldLocation,
	FieldImageURL,
	FieldUserID,
	FieldEventHash,
	FieldContractAddress,
	FieldTransactionHash,
	FieldBlockNumber,
	FieldCreatedAt,
	FieldModifiedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// EventSlugValidator is a validator for the "event_slug" field. It is called by the builders before save.
	EventSlugValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultModifiedAt holds the default value on creation for the "modified_at" field.
	DefaultModifiedAt func() time.Time
	// UpdateDefaultModifiedAt holds the default value on update for the "modified_at" field.
	UpdateDefaultModifiedAt func() time.Time
)

// OrderOption defines the ordering options for the Event queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByEventSlug orders the results by the event_slug field.
func ByEventSlug(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventSlug, opts...).ToFunc()
}

// ByStartDate orders the results by the start_date field.
func ByStartDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartDate, opts...).ToFunc()
}

// ByEndDate orders the results by the end_date field.
func ByEndDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndDate, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByImageURL orders the results by the image_url field.
func ByImageURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldImageURL, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByEventHash orders the results by the event_hash field.
func ByEventHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEventHash, opts...).ToFunc()
}

// ByContractAddress orders the results by the contract_address field.
func ByContractAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContractAddress, opts...).ToFunc()
}

// ByTransactionHash orders the results by the transaction_hash field.
func ByTransactionHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTransactionHash, opts...).ToFunc()
}

// ByBlockNumber orders the results by the block_number field.
func ByBlockNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBlockNumber, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByModifiedAt orders the results by the modified_at field.
func ByModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifiedAt, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}

// ByTicketsCount orders the results by tickets count.
func ByTicketsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTicketsStep(), opts...)
	}
}

// ByTickets orders the results by tickets terms.
func ByTickets(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTicketsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAttendeesCount orders the results by attendees count.
func ByAttendeesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAttendeesStep(), opts...)
	}
}

// ByAttendees orders the results by attendees terms.
func ByAttendees(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAttendeesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrdersCount orders the results by orders count.
func ByOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrdersStep(), opts...)
	}
}

// ByOrders orders the results by orders terms.
func ByOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
	)
}
func newTicketsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TicketsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TicketsTable, TicketsColumn),
	)
}
func newAttendeesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AttendeesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AttendeesTable, AttendeesColumn),
	)
}
func newOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrdersTable, OrdersColumn),
	)
}
