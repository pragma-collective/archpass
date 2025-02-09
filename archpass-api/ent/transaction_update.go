// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pragma-collective/archpass/ent/predicate"
	"github.com/pragma-collective/archpass/ent/transaction"
)

// TransactionUpdate is the builder for updating Transaction entities.
type TransactionUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionMutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tu *TransactionUpdate) Where(ps ...predicate.Transaction) *TransactionUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetEventType sets the "event_type" field.
func (tu *TransactionUpdate) SetEventType(s string) *TransactionUpdate {
	tu.mutation.SetEventType(s)
	return tu
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableEventType(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetEventType(*s)
	}
	return tu
}

// SetWalletAddress sets the "wallet_address" field.
func (tu *TransactionUpdate) SetWalletAddress(s string) *TransactionUpdate {
	tu.mutation.SetWalletAddress(s)
	return tu
}

// SetNillableWalletAddress sets the "wallet_address" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableWalletAddress(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetWalletAddress(*s)
	}
	return tu
}

// SetTransactionHash sets the "transaction_hash" field.
func (tu *TransactionUpdate) SetTransactionHash(s string) *TransactionUpdate {
	tu.mutation.SetTransactionHash(s)
	return tu
}

// SetNillableTransactionHash sets the "transaction_hash" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableTransactionHash(s *string) *TransactionUpdate {
	if s != nil {
		tu.SetTransactionHash(*s)
	}
	return tu
}

// SetBlockNumber sets the "block_number" field.
func (tu *TransactionUpdate) SetBlockNumber(i int64) *TransactionUpdate {
	tu.mutation.ResetBlockNumber()
	tu.mutation.SetBlockNumber(i)
	return tu
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (tu *TransactionUpdate) SetNillableBlockNumber(i *int64) *TransactionUpdate {
	if i != nil {
		tu.SetBlockNumber(*i)
	}
	return tu
}

// AddBlockNumber adds i to the "block_number" field.
func (tu *TransactionUpdate) AddBlockNumber(i int64) *TransactionUpdate {
	tu.mutation.AddBlockNumber(i)
	return tu
}

// Mutation returns the TransactionMutation object of the builder.
func (tu *TransactionUpdate) Mutation() *TransactionMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TransactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.EventType(); ok {
		_spec.SetField(transaction.FieldEventType, field.TypeString, value)
	}
	if value, ok := tu.mutation.WalletAddress(); ok {
		_spec.SetField(transaction.FieldWalletAddress, field.TypeString, value)
	}
	if value, ok := tu.mutation.TransactionHash(); ok {
		_spec.SetField(transaction.FieldTransactionHash, field.TypeString, value)
	}
	if value, ok := tu.mutation.BlockNumber(); ok {
		_spec.SetField(transaction.FieldBlockNumber, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedBlockNumber(); ok {
		_spec.AddField(transaction.FieldBlockNumber, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TransactionUpdateOne is the builder for updating a single Transaction entity.
type TransactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TransactionMutation
}

// SetEventType sets the "event_type" field.
func (tuo *TransactionUpdateOne) SetEventType(s string) *TransactionUpdateOne {
	tuo.mutation.SetEventType(s)
	return tuo
}

// SetNillableEventType sets the "event_type" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableEventType(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetEventType(*s)
	}
	return tuo
}

// SetWalletAddress sets the "wallet_address" field.
func (tuo *TransactionUpdateOne) SetWalletAddress(s string) *TransactionUpdateOne {
	tuo.mutation.SetWalletAddress(s)
	return tuo
}

// SetNillableWalletAddress sets the "wallet_address" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableWalletAddress(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetWalletAddress(*s)
	}
	return tuo
}

// SetTransactionHash sets the "transaction_hash" field.
func (tuo *TransactionUpdateOne) SetTransactionHash(s string) *TransactionUpdateOne {
	tuo.mutation.SetTransactionHash(s)
	return tuo
}

// SetNillableTransactionHash sets the "transaction_hash" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableTransactionHash(s *string) *TransactionUpdateOne {
	if s != nil {
		tuo.SetTransactionHash(*s)
	}
	return tuo
}

// SetBlockNumber sets the "block_number" field.
func (tuo *TransactionUpdateOne) SetBlockNumber(i int64) *TransactionUpdateOne {
	tuo.mutation.ResetBlockNumber()
	tuo.mutation.SetBlockNumber(i)
	return tuo
}

// SetNillableBlockNumber sets the "block_number" field if the given value is not nil.
func (tuo *TransactionUpdateOne) SetNillableBlockNumber(i *int64) *TransactionUpdateOne {
	if i != nil {
		tuo.SetBlockNumber(*i)
	}
	return tuo
}

// AddBlockNumber adds i to the "block_number" field.
func (tuo *TransactionUpdateOne) AddBlockNumber(i int64) *TransactionUpdateOne {
	tuo.mutation.AddBlockNumber(i)
	return tuo
}

// Mutation returns the TransactionMutation object of the builder.
func (tuo *TransactionUpdateOne) Mutation() *TransactionMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TransactionUpdate builder.
func (tuo *TransactionUpdateOne) Where(ps ...predicate.Transaction) *TransactionUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TransactionUpdateOne) Select(field string, fields ...string) *TransactionUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Transaction entity.
func (tuo *TransactionUpdateOne) Save(ctx context.Context) (*Transaction, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionUpdateOne) SaveX(ctx context.Context) *Transaction {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TransactionUpdateOne) sqlSave(ctx context.Context) (_node *Transaction, err error) {
	_spec := sqlgraph.NewUpdateSpec(transaction.Table, transaction.Columns, sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Transaction.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, transaction.FieldID)
		for _, f := range fields {
			if !transaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != transaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.EventType(); ok {
		_spec.SetField(transaction.FieldEventType, field.TypeString, value)
	}
	if value, ok := tuo.mutation.WalletAddress(); ok {
		_spec.SetField(transaction.FieldWalletAddress, field.TypeString, value)
	}
	if value, ok := tuo.mutation.TransactionHash(); ok {
		_spec.SetField(transaction.FieldTransactionHash, field.TypeString, value)
	}
	if value, ok := tuo.mutation.BlockNumber(); ok {
		_spec.SetField(transaction.FieldBlockNumber, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedBlockNumber(); ok {
		_spec.AddField(transaction.FieldBlockNumber, field.TypeInt64, value)
	}
	_node = &Transaction{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
