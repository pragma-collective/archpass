// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pragma-collective/archpass/ent/attendee"
	"github.com/pragma-collective/archpass/ent/event"
	"github.com/pragma-collective/archpass/ent/ticket"
	"github.com/pragma-collective/archpass/ent/user"
)

// AttendeeCreate is the builder for creating a Attendee entity.
type AttendeeCreate struct {
	config
	mutation *AttendeeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUserID sets the "user_id" field.
func (ac *AttendeeCreate) SetUserID(i int) *AttendeeCreate {
	ac.mutation.SetUserID(i)
	return ac
}

// SetEventID sets the "event_id" field.
func (ac *AttendeeCreate) SetEventID(i int) *AttendeeCreate {
	ac.mutation.SetEventID(i)
	return ac
}

// SetTicketID sets the "ticket_id" field.
func (ac *AttendeeCreate) SetTicketID(i int) *AttendeeCreate {
	ac.mutation.SetTicketID(i)
	return ac
}

// SetTokenID sets the "token_id" field.
func (ac *AttendeeCreate) SetTokenID(i int) *AttendeeCreate {
	ac.mutation.SetTokenID(i)
	return ac
}

// SetTransactionHash sets the "transaction_hash" field.
func (ac *AttendeeCreate) SetTransactionHash(s string) *AttendeeCreate {
	ac.mutation.SetTransactionHash(s)
	return ac
}

// SetBlockNumber sets the "block_number" field.
func (ac *AttendeeCreate) SetBlockNumber(i int64) *AttendeeCreate {
	ac.mutation.SetBlockNumber(i)
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AttendeeCreate) SetCreatedAt(t time.Time) *AttendeeCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AttendeeCreate) SetNillableCreatedAt(t *time.Time) *AttendeeCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetEvent sets the "event" edge to the Event entity.
func (ac *AttendeeCreate) SetEvent(e *Event) *AttendeeCreate {
	return ac.SetEventID(e.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ac *AttendeeCreate) SetUser(u *User) *AttendeeCreate {
	return ac.SetUserID(u.ID)
}

// SetTicket sets the "ticket" edge to the Ticket entity.
func (ac *AttendeeCreate) SetTicket(t *Ticket) *AttendeeCreate {
	return ac.SetTicketID(t.ID)
}

// Mutation returns the AttendeeMutation object of the builder.
func (ac *AttendeeCreate) Mutation() *AttendeeMutation {
	return ac.mutation
}

// Save creates the Attendee in the database.
func (ac *AttendeeCreate) Save(ctx context.Context) (*Attendee, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AttendeeCreate) SaveX(ctx context.Context) *Attendee {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AttendeeCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AttendeeCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AttendeeCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := attendee.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AttendeeCreate) check() error {
	if _, ok := ac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Attendee.user_id"`)}
	}
	if _, ok := ac.mutation.EventID(); !ok {
		return &ValidationError{Name: "event_id", err: errors.New(`ent: missing required field "Attendee.event_id"`)}
	}
	if _, ok := ac.mutation.TicketID(); !ok {
		return &ValidationError{Name: "ticket_id", err: errors.New(`ent: missing required field "Attendee.ticket_id"`)}
	}
	if _, ok := ac.mutation.TokenID(); !ok {
		return &ValidationError{Name: "token_id", err: errors.New(`ent: missing required field "Attendee.token_id"`)}
	}
	if _, ok := ac.mutation.TransactionHash(); !ok {
		return &ValidationError{Name: "transaction_hash", err: errors.New(`ent: missing required field "Attendee.transaction_hash"`)}
	}
	if _, ok := ac.mutation.BlockNumber(); !ok {
		return &ValidationError{Name: "block_number", err: errors.New(`ent: missing required field "Attendee.block_number"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Attendee.created_at"`)}
	}
	if len(ac.mutation.EventIDs()) == 0 {
		return &ValidationError{Name: "event", err: errors.New(`ent: missing required edge "Attendee.event"`)}
	}
	if len(ac.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Attendee.user"`)}
	}
	if len(ac.mutation.TicketIDs()) == 0 {
		return &ValidationError{Name: "ticket", err: errors.New(`ent: missing required edge "Attendee.ticket"`)}
	}
	return nil
}

func (ac *AttendeeCreate) sqlSave(ctx context.Context) (*Attendee, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AttendeeCreate) createSpec() (*Attendee, *sqlgraph.CreateSpec) {
	var (
		_node = &Attendee{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(attendee.Table, sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.TokenID(); ok {
		_spec.SetField(attendee.FieldTokenID, field.TypeInt, value)
		_node.TokenID = value
	}
	if value, ok := ac.mutation.TransactionHash(); ok {
		_spec.SetField(attendee.FieldTransactionHash, field.TypeString, value)
		_node.TransactionHash = value
	}
	if value, ok := ac.mutation.BlockNumber(); ok {
		_spec.SetField(attendee.FieldBlockNumber, field.TypeInt64, value)
		_node.BlockNumber = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(attendee.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := ac.mutation.EventIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.EventTable,
			Columns: []string{attendee.EventColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EventID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attendee.UserTable,
			Columns: []string{attendee.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ac.mutation.TicketIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   attendee.TicketTable,
			Columns: []string{attendee.TicketColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TicketID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Attendee.Create().
//		SetUserID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttendeeUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (ac *AttendeeCreate) OnConflict(opts ...sql.ConflictOption) *AttendeeUpsertOne {
	ac.conflict = opts
	return &AttendeeUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Attendee.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AttendeeCreate) OnConflictColumns(columns ...string) *AttendeeUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AttendeeUpsertOne{
		create: ac,
	}
}

type (
	// AttendeeUpsertOne is the builder for "upsert"-ing
	//  one Attendee node.
	AttendeeUpsertOne struct {
		create *AttendeeCreate
	}

	// AttendeeUpsert is the "OnConflict" setter.
	AttendeeUpsert struct {
		*sql.UpdateSet
	}
)

// SetUserID sets the "user_id" field.
func (u *AttendeeUpsert) SetUserID(v int) *AttendeeUpsert {
	u.Set(attendee.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateUserID() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldUserID)
	return u
}

// SetEventID sets the "event_id" field.
func (u *AttendeeUpsert) SetEventID(v int) *AttendeeUpsert {
	u.Set(attendee.FieldEventID, v)
	return u
}

// UpdateEventID sets the "event_id" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateEventID() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldEventID)
	return u
}

// SetTicketID sets the "ticket_id" field.
func (u *AttendeeUpsert) SetTicketID(v int) *AttendeeUpsert {
	u.Set(attendee.FieldTicketID, v)
	return u
}

// UpdateTicketID sets the "ticket_id" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateTicketID() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldTicketID)
	return u
}

// SetTokenID sets the "token_id" field.
func (u *AttendeeUpsert) SetTokenID(v int) *AttendeeUpsert {
	u.Set(attendee.FieldTokenID, v)
	return u
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateTokenID() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldTokenID)
	return u
}

// AddTokenID adds v to the "token_id" field.
func (u *AttendeeUpsert) AddTokenID(v int) *AttendeeUpsert {
	u.Add(attendee.FieldTokenID, v)
	return u
}

// SetTransactionHash sets the "transaction_hash" field.
func (u *AttendeeUpsert) SetTransactionHash(v string) *AttendeeUpsert {
	u.Set(attendee.FieldTransactionHash, v)
	return u
}

// UpdateTransactionHash sets the "transaction_hash" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateTransactionHash() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldTransactionHash)
	return u
}

// SetBlockNumber sets the "block_number" field.
func (u *AttendeeUpsert) SetBlockNumber(v int64) *AttendeeUpsert {
	u.Set(attendee.FieldBlockNumber, v)
	return u
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateBlockNumber() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldBlockNumber)
	return u
}

// AddBlockNumber adds v to the "block_number" field.
func (u *AttendeeUpsert) AddBlockNumber(v int64) *AttendeeUpsert {
	u.Add(attendee.FieldBlockNumber, v)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AttendeeUpsert) SetCreatedAt(v time.Time) *AttendeeUpsert {
	u.Set(attendee.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AttendeeUpsert) UpdateCreatedAt() *AttendeeUpsert {
	u.SetExcluded(attendee.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Attendee.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AttendeeUpsertOne) UpdateNewValues() *AttendeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Attendee.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AttendeeUpsertOne) Ignore() *AttendeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttendeeUpsertOne) DoNothing() *AttendeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttendeeCreate.OnConflict
// documentation for more info.
func (u *AttendeeUpsertOne) Update(set func(*AttendeeUpsert)) *AttendeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttendeeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *AttendeeUpsertOne) SetUserID(v int) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateUserID() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateUserID()
	})
}

// SetEventID sets the "event_id" field.
func (u *AttendeeUpsertOne) SetEventID(v int) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetEventID(v)
	})
}

// UpdateEventID sets the "event_id" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateEventID() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateEventID()
	})
}

// SetTicketID sets the "ticket_id" field.
func (u *AttendeeUpsertOne) SetTicketID(v int) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetTicketID(v)
	})
}

// UpdateTicketID sets the "ticket_id" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateTicketID() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateTicketID()
	})
}

// SetTokenID sets the "token_id" field.
func (u *AttendeeUpsertOne) SetTokenID(v int) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetTokenID(v)
	})
}

// AddTokenID adds v to the "token_id" field.
func (u *AttendeeUpsertOne) AddTokenID(v int) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.AddTokenID(v)
	})
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateTokenID() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateTokenID()
	})
}

// SetTransactionHash sets the "transaction_hash" field.
func (u *AttendeeUpsertOne) SetTransactionHash(v string) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetTransactionHash(v)
	})
}

// UpdateTransactionHash sets the "transaction_hash" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateTransactionHash() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateTransactionHash()
	})
}

// SetBlockNumber sets the "block_number" field.
func (u *AttendeeUpsertOne) SetBlockNumber(v int64) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetBlockNumber(v)
	})
}

// AddBlockNumber adds v to the "block_number" field.
func (u *AttendeeUpsertOne) AddBlockNumber(v int64) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.AddBlockNumber(v)
	})
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateBlockNumber() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateBlockNumber()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AttendeeUpsertOne) SetCreatedAt(v time.Time) *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AttendeeUpsertOne) UpdateCreatedAt() *AttendeeUpsertOne {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AttendeeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttendeeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttendeeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AttendeeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AttendeeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AttendeeCreateBulk is the builder for creating many Attendee entities in bulk.
type AttendeeCreateBulk struct {
	config
	err      error
	builders []*AttendeeCreate
	conflict []sql.ConflictOption
}

// Save creates the Attendee entities in the database.
func (acb *AttendeeCreateBulk) Save(ctx context.Context) ([]*Attendee, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Attendee, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AttendeeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AttendeeCreateBulk) SaveX(ctx context.Context) []*Attendee {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AttendeeCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AttendeeCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Attendee.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AttendeeUpsert) {
//			SetUserID(v+v).
//		}).
//		Exec(ctx)
func (acb *AttendeeCreateBulk) OnConflict(opts ...sql.ConflictOption) *AttendeeUpsertBulk {
	acb.conflict = opts
	return &AttendeeUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Attendee.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AttendeeCreateBulk) OnConflictColumns(columns ...string) *AttendeeUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AttendeeUpsertBulk{
		create: acb,
	}
}

// AttendeeUpsertBulk is the builder for "upsert"-ing
// a bulk of Attendee nodes.
type AttendeeUpsertBulk struct {
	create *AttendeeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Attendee.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AttendeeUpsertBulk) UpdateNewValues() *AttendeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Attendee.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AttendeeUpsertBulk) Ignore() *AttendeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AttendeeUpsertBulk) DoNothing() *AttendeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AttendeeCreateBulk.OnConflict
// documentation for more info.
func (u *AttendeeUpsertBulk) Update(set func(*AttendeeUpsert)) *AttendeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AttendeeUpsert{UpdateSet: update})
	}))
	return u
}

// SetUserID sets the "user_id" field.
func (u *AttendeeUpsertBulk) SetUserID(v int) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateUserID() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateUserID()
	})
}

// SetEventID sets the "event_id" field.
func (u *AttendeeUpsertBulk) SetEventID(v int) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetEventID(v)
	})
}

// UpdateEventID sets the "event_id" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateEventID() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateEventID()
	})
}

// SetTicketID sets the "ticket_id" field.
func (u *AttendeeUpsertBulk) SetTicketID(v int) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetTicketID(v)
	})
}

// UpdateTicketID sets the "ticket_id" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateTicketID() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateTicketID()
	})
}

// SetTokenID sets the "token_id" field.
func (u *AttendeeUpsertBulk) SetTokenID(v int) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetTokenID(v)
	})
}

// AddTokenID adds v to the "token_id" field.
func (u *AttendeeUpsertBulk) AddTokenID(v int) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.AddTokenID(v)
	})
}

// UpdateTokenID sets the "token_id" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateTokenID() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateTokenID()
	})
}

// SetTransactionHash sets the "transaction_hash" field.
func (u *AttendeeUpsertBulk) SetTransactionHash(v string) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetTransactionHash(v)
	})
}

// UpdateTransactionHash sets the "transaction_hash" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateTransactionHash() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateTransactionHash()
	})
}

// SetBlockNumber sets the "block_number" field.
func (u *AttendeeUpsertBulk) SetBlockNumber(v int64) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetBlockNumber(v)
	})
}

// AddBlockNumber adds v to the "block_number" field.
func (u *AttendeeUpsertBulk) AddBlockNumber(v int64) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.AddBlockNumber(v)
	})
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateBlockNumber() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateBlockNumber()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *AttendeeUpsertBulk) SetCreatedAt(v time.Time) *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AttendeeUpsertBulk) UpdateCreatedAt() *AttendeeUpsertBulk {
	return u.Update(func(s *AttendeeUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *AttendeeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AttendeeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AttendeeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AttendeeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
