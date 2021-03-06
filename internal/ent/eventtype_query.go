// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/event"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/eventtype"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/predicate"
	"github.com/google/uuid"
)

// EventTypeQuery is the builder for querying EventType entities.
type EventTypeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EventType
	// eager-loading edges.
	withEvents *EventQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EventTypeQuery builder.
func (etq *EventTypeQuery) Where(ps ...predicate.EventType) *EventTypeQuery {
	etq.predicates = append(etq.predicates, ps...)
	return etq
}

// Limit adds a limit step to the query.
func (etq *EventTypeQuery) Limit(limit int) *EventTypeQuery {
	etq.limit = &limit
	return etq
}

// Offset adds an offset step to the query.
func (etq *EventTypeQuery) Offset(offset int) *EventTypeQuery {
	etq.offset = &offset
	return etq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (etq *EventTypeQuery) Unique(unique bool) *EventTypeQuery {
	etq.unique = &unique
	return etq
}

// Order adds an order step to the query.
func (etq *EventTypeQuery) Order(o ...OrderFunc) *EventTypeQuery {
	etq.order = append(etq.order, o...)
	return etq
}

// QueryEvents chains the current query on the "events" edge.
func (etq *EventTypeQuery) QueryEvents() *EventQuery {
	query := &EventQuery{config: etq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := etq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(eventtype.Table, eventtype.FieldID, selector),
			sqlgraph.To(event.Table, event.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, eventtype.EventsTable, eventtype.EventsColumn),
		)
		fromU = sqlgraph.SetNeighbors(etq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EventType entity from the query.
// Returns a *NotFoundError when no EventType was found.
func (etq *EventTypeQuery) First(ctx context.Context) (*EventType, error) {
	nodes, err := etq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{eventtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (etq *EventTypeQuery) FirstX(ctx context.Context) *EventType {
	node, err := etq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EventType ID from the query.
// Returns a *NotFoundError when no EventType ID was found.
func (etq *EventTypeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = etq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{eventtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (etq *EventTypeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := etq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EventType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EventType entity is found.
// Returns a *NotFoundError when no EventType entities are found.
func (etq *EventTypeQuery) Only(ctx context.Context) (*EventType, error) {
	nodes, err := etq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{eventtype.Label}
	default:
		return nil, &NotSingularError{eventtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (etq *EventTypeQuery) OnlyX(ctx context.Context) *EventType {
	node, err := etq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EventType ID in the query.
// Returns a *NotSingularError when more than one EventType ID is found.
// Returns a *NotFoundError when no entities are found.
func (etq *EventTypeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = etq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = &NotSingularError{eventtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (etq *EventTypeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := etq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EventTypes.
func (etq *EventTypeQuery) All(ctx context.Context) ([]*EventType, error) {
	if err := etq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return etq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (etq *EventTypeQuery) AllX(ctx context.Context) []*EventType {
	nodes, err := etq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EventType IDs.
func (etq *EventTypeQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := etq.Select(eventtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (etq *EventTypeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := etq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (etq *EventTypeQuery) Count(ctx context.Context) (int, error) {
	if err := etq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return etq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (etq *EventTypeQuery) CountX(ctx context.Context) int {
	count, err := etq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (etq *EventTypeQuery) Exist(ctx context.Context) (bool, error) {
	if err := etq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return etq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (etq *EventTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := etq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EventTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (etq *EventTypeQuery) Clone() *EventTypeQuery {
	if etq == nil {
		return nil
	}
	return &EventTypeQuery{
		config:     etq.config,
		limit:      etq.limit,
		offset:     etq.offset,
		order:      append([]OrderFunc{}, etq.order...),
		predicates: append([]predicate.EventType{}, etq.predicates...),
		withEvents: etq.withEvents.Clone(),
		// clone intermediate query.
		sql:    etq.sql.Clone(),
		path:   etq.path,
		unique: etq.unique,
	}
}

// WithEvents tells the query-builder to eager-load the nodes that are connected to
// the "events" edge. The optional arguments are used to configure the query builder of the edge.
func (etq *EventTypeQuery) WithEvents(opts ...func(*EventQuery)) *EventTypeQuery {
	query := &EventQuery{config: etq.config}
	for _, opt := range opts {
		opt(query)
	}
	etq.withEvents = query
	return etq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EventType.Query().
//		GroupBy(eventtype.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (etq *EventTypeQuery) GroupBy(field string, fields ...string) *EventTypeGroupBy {
	group := &EventTypeGroupBy{config: etq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := etq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return etq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.EventType.Query().
//		Select(eventtype.FieldName).
//		Scan(ctx, &v)
//
func (etq *EventTypeQuery) Select(fields ...string) *EventTypeSelect {
	etq.fields = append(etq.fields, fields...)
	return &EventTypeSelect{EventTypeQuery: etq}
}

func (etq *EventTypeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range etq.fields {
		if !eventtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if etq.path != nil {
		prev, err := etq.path(ctx)
		if err != nil {
			return err
		}
		etq.sql = prev
	}
	return nil
}

func (etq *EventTypeQuery) sqlAll(ctx context.Context) ([]*EventType, error) {
	var (
		nodes       = []*EventType{}
		_spec       = etq.querySpec()
		loadedTypes = [1]bool{
			etq.withEvents != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EventType{config: etq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, etq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := etq.withEvents; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[uuid.UUID]*EventType)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Events = []*Event{}
		}
		query.withFKs = true
		query.Where(predicate.Event(func(s *sql.Selector) {
			s.Where(sql.InValues(eventtype.EventsColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.event_type_events
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "event_type_events" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "event_type_events" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Events = append(node.Edges.Events, n)
		}
	}

	return nodes, nil
}

func (etq *EventTypeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := etq.querySpec()
	_spec.Node.Columns = etq.fields
	if len(etq.fields) > 0 {
		_spec.Unique = etq.unique != nil && *etq.unique
	}
	return sqlgraph.CountNodes(ctx, etq.driver, _spec)
}

func (etq *EventTypeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := etq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (etq *EventTypeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   eventtype.Table,
			Columns: eventtype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventtype.FieldID,
			},
		},
		From:   etq.sql,
		Unique: true,
	}
	if unique := etq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := etq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, eventtype.FieldID)
		for i := range fields {
			if fields[i] != eventtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := etq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := etq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := etq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := etq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (etq *EventTypeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(etq.driver.Dialect())
	t1 := builder.Table(eventtype.Table)
	columns := etq.fields
	if len(columns) == 0 {
		columns = eventtype.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if etq.sql != nil {
		selector = etq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if etq.unique != nil && *etq.unique {
		selector.Distinct()
	}
	for _, p := range etq.predicates {
		p(selector)
	}
	for _, p := range etq.order {
		p(selector)
	}
	if offset := etq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := etq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EventTypeGroupBy is the group-by builder for EventType entities.
type EventTypeGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (etgb *EventTypeGroupBy) Aggregate(fns ...AggregateFunc) *EventTypeGroupBy {
	etgb.fns = append(etgb.fns, fns...)
	return etgb
}

// Scan applies the group-by query and scans the result into the given value.
func (etgb *EventTypeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := etgb.path(ctx)
	if err != nil {
		return err
	}
	etgb.sql = query
	return etgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (etgb *EventTypeGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := etgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EventTypeGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (etgb *EventTypeGroupBy) StringsX(ctx context.Context) []string {
	v, err := etgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = etgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (etgb *EventTypeGroupBy) StringX(ctx context.Context) string {
	v, err := etgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EventTypeGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (etgb *EventTypeGroupBy) IntsX(ctx context.Context) []int {
	v, err := etgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = etgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (etgb *EventTypeGroupBy) IntX(ctx context.Context) int {
	v, err := etgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EventTypeGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (etgb *EventTypeGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := etgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = etgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (etgb *EventTypeGroupBy) Float64X(ctx context.Context) float64 {
	v, err := etgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(etgb.fields) > 1 {
		return nil, errors.New("ent: EventTypeGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := etgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (etgb *EventTypeGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := etgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (etgb *EventTypeGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = etgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (etgb *EventTypeGroupBy) BoolX(ctx context.Context) bool {
	v, err := etgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (etgb *EventTypeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range etgb.fields {
		if !eventtype.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := etgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := etgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (etgb *EventTypeGroupBy) sqlQuery() *sql.Selector {
	selector := etgb.sql.Select()
	aggregation := make([]string, 0, len(etgb.fns))
	for _, fn := range etgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(etgb.fields)+len(etgb.fns))
		for _, f := range etgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(etgb.fields...)...)
}

// EventTypeSelect is the builder for selecting fields of EventType entities.
type EventTypeSelect struct {
	*EventTypeQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ets *EventTypeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ets.prepareQuery(ctx); err != nil {
		return err
	}
	ets.sql = ets.EventTypeQuery.sqlQuery(ctx)
	return ets.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ets *EventTypeSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ets.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EventTypeSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ets *EventTypeSelect) StringsX(ctx context.Context) []string {
	v, err := ets.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ets.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ets *EventTypeSelect) StringX(ctx context.Context) string {
	v, err := ets.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EventTypeSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ets *EventTypeSelect) IntsX(ctx context.Context) []int {
	v, err := ets.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ets.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ets *EventTypeSelect) IntX(ctx context.Context) int {
	v, err := ets.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EventTypeSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ets *EventTypeSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ets.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ets.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ets *EventTypeSelect) Float64X(ctx context.Context) float64 {
	v, err := ets.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ets.fields) > 1 {
		return nil, errors.New("ent: EventTypeSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ets.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ets *EventTypeSelect) BoolsX(ctx context.Context) []bool {
	v, err := ets.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ets *EventTypeSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ets.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{eventtype.Label}
	default:
		err = fmt.Errorf("ent: EventTypeSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ets *EventTypeSelect) BoolX(ctx context.Context) bool {
	v, err := ets.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ets *EventTypeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ets.sql.Query()
	if err := ets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
