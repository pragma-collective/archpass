// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/pragma-collective/archpass/ent/attendee"
	"github.com/pragma-collective/archpass/ent/event"
	"github.com/pragma-collective/archpass/ent/predicate"
	"github.com/pragma-collective/archpass/ent/ticket"
	"github.com/pragma-collective/archpass/ent/user"
)

// AttendeeQuery is the builder for querying Attendee entities.
type AttendeeQuery struct {
	config
	ctx        *QueryContext
	order      []attendee.OrderOption
	inters     []Interceptor
	predicates []predicate.Attendee
	withEvent  *EventQuery
	withUser   *UserQuery
	withTicket *TicketQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AttendeeQuery builder.
func (aq *AttendeeQuery) Where(ps ...predicate.Attendee) *AttendeeQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *AttendeeQuery) Limit(limit int) *AttendeeQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *AttendeeQuery) Offset(offset int) *AttendeeQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *AttendeeQuery) Unique(unique bool) *AttendeeQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *AttendeeQuery) Order(o ...attendee.OrderOption) *AttendeeQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryEvent chains the current query on the "event" edge.
func (aq *AttendeeQuery) QueryEvent() *EventQuery {
	query := (&EventClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendee.Table, attendee.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendee.EventTable, attendee.EventColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (aq *AttendeeQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendee.Table, attendee.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, attendee.UserTable, attendee.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTicket chains the current query on the "ticket" edge.
func (aq *AttendeeQuery) QueryTicket() *TicketQuery {
	query := (&TicketClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(attendee.Table, attendee.FieldID, selector),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, attendee.TicketTable, attendee.TicketColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Attendee entity from the query.
// Returns a *NotFoundError when no Attendee was found.
func (aq *AttendeeQuery) First(ctx context.Context) (*Attendee, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{attendee.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *AttendeeQuery) FirstX(ctx context.Context) *Attendee {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Attendee ID from the query.
// Returns a *NotFoundError when no Attendee ID was found.
func (aq *AttendeeQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{attendee.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *AttendeeQuery) FirstIDX(ctx context.Context) int {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Attendee entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Attendee entity is found.
// Returns a *NotFoundError when no Attendee entities are found.
func (aq *AttendeeQuery) Only(ctx context.Context) (*Attendee, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{attendee.Label}
	default:
		return nil, &NotSingularError{attendee.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *AttendeeQuery) OnlyX(ctx context.Context) *Attendee {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Attendee ID in the query.
// Returns a *NotSingularError when more than one Attendee ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *AttendeeQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{attendee.Label}
	default:
		err = &NotSingularError{attendee.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *AttendeeQuery) OnlyIDX(ctx context.Context) int {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Attendees.
func (aq *AttendeeQuery) All(ctx context.Context) ([]*Attendee, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryAll)
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Attendee, *AttendeeQuery]()
	return withInterceptors[[]*Attendee](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *AttendeeQuery) AllX(ctx context.Context) []*Attendee {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Attendee IDs.
func (aq *AttendeeQuery) IDs(ctx context.Context) (ids []int, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryIDs)
	if err = aq.Select(attendee.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *AttendeeQuery) IDsX(ctx context.Context) []int {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *AttendeeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryCount)
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*AttendeeQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *AttendeeQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *AttendeeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, ent.OpQueryExist)
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *AttendeeQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AttendeeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *AttendeeQuery) Clone() *AttendeeQuery {
	if aq == nil {
		return nil
	}
	return &AttendeeQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]attendee.OrderOption{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Attendee{}, aq.predicates...),
		withEvent:  aq.withEvent.Clone(),
		withUser:   aq.withUser.Clone(),
		withTicket: aq.withTicket.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithEvent tells the query-builder to eager-load the nodes that are connected to
// the "event" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttendeeQuery) WithEvent(opts ...func(*EventQuery)) *AttendeeQuery {
	query := (&EventClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withEvent = query
	return aq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttendeeQuery) WithUser(opts ...func(*UserQuery)) *AttendeeQuery {
	query := (&UserClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withUser = query
	return aq
}

// WithTicket tells the query-builder to eager-load the nodes that are connected to
// the "ticket" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *AttendeeQuery) WithTicket(opts ...func(*TicketQuery)) *AttendeeQuery {
	query := (&TicketClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withTicket = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Attendee.Query().
//		GroupBy(attendee.FieldUserID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *AttendeeQuery) GroupBy(field string, fields ...string) *AttendeeGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AttendeeGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = attendee.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int `json:"user_id,omitempty"`
//	}
//
//	client.Attendee.Query().
//		Select(attendee.FieldUserID).
//		Scan(ctx, &v)
func (aq *AttendeeQuery) Select(fields ...string) *AttendeeSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &AttendeeSelect{AttendeeQuery: aq}
	sbuild.label = attendee.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AttendeeSelect configured with the given aggregations.
func (aq *AttendeeQuery) Aggregate(fns ...AggregateFunc) *AttendeeSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *AttendeeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !attendee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *AttendeeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Attendee, error) {
	var (
		nodes       = []*Attendee{}
		_spec       = aq.querySpec()
		loadedTypes = [3]bool{
			aq.withEvent != nil,
			aq.withUser != nil,
			aq.withTicket != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Attendee).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Attendee{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withEvent; query != nil {
		if err := aq.loadEvent(ctx, query, nodes, nil,
			func(n *Attendee, e *Event) { n.Edges.Event = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withUser; query != nil {
		if err := aq.loadUser(ctx, query, nodes, nil,
			func(n *Attendee, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withTicket; query != nil {
		if err := aq.loadTicket(ctx, query, nodes, nil,
			func(n *Attendee, e *Ticket) { n.Edges.Ticket = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *AttendeeQuery) loadEvent(ctx context.Context, query *EventQuery, nodes []*Attendee, init func(*Attendee), assign func(*Attendee, *Event)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Attendee)
	for i := range nodes {
		fk := nodes[i].EventID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(event.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "event_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttendeeQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Attendee, init func(*Attendee), assign func(*Attendee, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Attendee)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *AttendeeQuery) loadTicket(ctx context.Context, query *TicketQuery, nodes []*Attendee, init func(*Attendee), assign func(*Attendee, *Ticket)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Attendee)
	for i := range nodes {
		fk := nodes[i].TicketID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(ticket.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "ticket_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *AttendeeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *AttendeeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(attendee.Table, attendee.Columns, sqlgraph.NewFieldSpec(attendee.FieldID, field.TypeInt))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attendee.FieldID)
		for i := range fields {
			if fields[i] != attendee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if aq.withEvent != nil {
			_spec.Node.AddColumnOnce(attendee.FieldEventID)
		}
		if aq.withUser != nil {
			_spec.Node.AddColumnOnce(attendee.FieldUserID)
		}
		if aq.withTicket != nil {
			_spec.Node.AddColumnOnce(attendee.FieldTicketID)
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *AttendeeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(attendee.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = attendee.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AttendeeGroupBy is the group-by builder for Attendee entities.
type AttendeeGroupBy struct {
	selector
	build *AttendeeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *AttendeeGroupBy) Aggregate(fns ...AggregateFunc) *AttendeeGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *AttendeeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, ent.OpQueryGroupBy)
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttendeeQuery, *AttendeeGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *AttendeeGroupBy) sqlScan(ctx context.Context, root *AttendeeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AttendeeSelect is the builder for selecting fields of Attendee entities.
type AttendeeSelect struct {
	*AttendeeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *AttendeeSelect) Aggregate(fns ...AggregateFunc) *AttendeeSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *AttendeeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, ent.OpQuerySelect)
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AttendeeQuery, *AttendeeSelect](ctx, as.AttendeeQuery, as, as.inters, v)
}

func (as *AttendeeSelect) sqlScan(ctx context.Context, root *AttendeeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
