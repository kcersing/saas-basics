// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/orderamount"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderAmountQuery is the builder for querying OrderAmount entities.
type OrderAmountQuery struct {
	config
	ctx        *QueryContext
	order      []orderamount.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderAmount
	withOrder  *OrderQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderAmountQuery builder.
func (oaq *OrderAmountQuery) Where(ps ...predicate.OrderAmount) *OrderAmountQuery {
	oaq.predicates = append(oaq.predicates, ps...)
	return oaq
}

// Limit the number of records to be returned by this query.
func (oaq *OrderAmountQuery) Limit(limit int) *OrderAmountQuery {
	oaq.ctx.Limit = &limit
	return oaq
}

// Offset to start from.
func (oaq *OrderAmountQuery) Offset(offset int) *OrderAmountQuery {
	oaq.ctx.Offset = &offset
	return oaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oaq *OrderAmountQuery) Unique(unique bool) *OrderAmountQuery {
	oaq.ctx.Unique = &unique
	return oaq
}

// Order specifies how the records should be ordered.
func (oaq *OrderAmountQuery) Order(o ...orderamount.OrderOption) *OrderAmountQuery {
	oaq.order = append(oaq.order, o...)
	return oaq
}

// QueryOrder chains the current query on the "order" edge.
func (oaq *OrderAmountQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: oaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderamount.Table, orderamount.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderamount.OrderTable, orderamount.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(oaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderAmount entity from the query.
// Returns a *NotFoundError when no OrderAmount was found.
func (oaq *OrderAmountQuery) First(ctx context.Context) (*OrderAmount, error) {
	nodes, err := oaq.Limit(1).All(setContextOp(ctx, oaq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderamount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oaq *OrderAmountQuery) FirstX(ctx context.Context) *OrderAmount {
	node, err := oaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderAmount ID from the query.
// Returns a *NotFoundError when no OrderAmount ID was found.
func (oaq *OrderAmountQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oaq.Limit(1).IDs(setContextOp(ctx, oaq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderamount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oaq *OrderAmountQuery) FirstIDX(ctx context.Context) int64 {
	id, err := oaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderAmount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderAmount entity is found.
// Returns a *NotFoundError when no OrderAmount entities are found.
func (oaq *OrderAmountQuery) Only(ctx context.Context) (*OrderAmount, error) {
	nodes, err := oaq.Limit(2).All(setContextOp(ctx, oaq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderamount.Label}
	default:
		return nil, &NotSingularError{orderamount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oaq *OrderAmountQuery) OnlyX(ctx context.Context) *OrderAmount {
	node, err := oaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderAmount ID in the query.
// Returns a *NotSingularError when more than one OrderAmount ID is found.
// Returns a *NotFoundError when no entities are found.
func (oaq *OrderAmountQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oaq.Limit(2).IDs(setContextOp(ctx, oaq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderamount.Label}
	default:
		err = &NotSingularError{orderamount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oaq *OrderAmountQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := oaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderAmounts.
func (oaq *OrderAmountQuery) All(ctx context.Context) ([]*OrderAmount, error) {
	ctx = setContextOp(ctx, oaq.ctx, ent.OpQueryAll)
	if err := oaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderAmount, *OrderAmountQuery]()
	return withInterceptors[[]*OrderAmount](ctx, oaq, qr, oaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oaq *OrderAmountQuery) AllX(ctx context.Context) []*OrderAmount {
	nodes, err := oaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderAmount IDs.
func (oaq *OrderAmountQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if oaq.ctx.Unique == nil && oaq.path != nil {
		oaq.Unique(true)
	}
	ctx = setContextOp(ctx, oaq.ctx, ent.OpQueryIDs)
	if err = oaq.Select(orderamount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oaq *OrderAmountQuery) IDsX(ctx context.Context) []int64 {
	ids, err := oaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oaq *OrderAmountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oaq.ctx, ent.OpQueryCount)
	if err := oaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oaq, querierCount[*OrderAmountQuery](), oaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oaq *OrderAmountQuery) CountX(ctx context.Context) int {
	count, err := oaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oaq *OrderAmountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oaq.ctx, ent.OpQueryExist)
	switch _, err := oaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oaq *OrderAmountQuery) ExistX(ctx context.Context) bool {
	exist, err := oaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderAmountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oaq *OrderAmountQuery) Clone() *OrderAmountQuery {
	if oaq == nil {
		return nil
	}
	return &OrderAmountQuery{
		config:     oaq.config,
		ctx:        oaq.ctx.Clone(),
		order:      append([]orderamount.OrderOption{}, oaq.order...),
		inters:     append([]Interceptor{}, oaq.inters...),
		predicates: append([]predicate.OrderAmount{}, oaq.predicates...),
		withOrder:  oaq.withOrder.Clone(),
		// clone intermediate query.
		sql:  oaq.sql.Clone(),
		path: oaq.path,
	}
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (oaq *OrderAmountQuery) WithOrder(opts ...func(*OrderQuery)) *OrderAmountQuery {
	query := (&OrderClient{config: oaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oaq.withOrder = query
	return oaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderAmount.Query().
//		GroupBy(orderamount.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oaq *OrderAmountQuery) GroupBy(field string, fields ...string) *OrderAmountGroupBy {
	oaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderAmountGroupBy{build: oaq}
	grbuild.flds = &oaq.ctx.Fields
	grbuild.label = orderamount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.OrderAmount.Query().
//		Select(orderamount.FieldCreatedAt).
//		Scan(ctx, &v)
func (oaq *OrderAmountQuery) Select(fields ...string) *OrderAmountSelect {
	oaq.ctx.Fields = append(oaq.ctx.Fields, fields...)
	sbuild := &OrderAmountSelect{OrderAmountQuery: oaq}
	sbuild.label = orderamount.Label
	sbuild.flds, sbuild.scan = &oaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderAmountSelect configured with the given aggregations.
func (oaq *OrderAmountQuery) Aggregate(fns ...AggregateFunc) *OrderAmountSelect {
	return oaq.Select().Aggregate(fns...)
}

func (oaq *OrderAmountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oaq); err != nil {
				return err
			}
		}
	}
	for _, f := range oaq.ctx.Fields {
		if !orderamount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oaq.path != nil {
		prev, err := oaq.path(ctx)
		if err != nil {
			return err
		}
		oaq.sql = prev
	}
	return nil
}

func (oaq *OrderAmountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderAmount, error) {
	var (
		nodes       = []*OrderAmount{}
		_spec       = oaq.querySpec()
		loadedTypes = [1]bool{
			oaq.withOrder != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderAmount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderAmount{config: oaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oaq.withOrder; query != nil {
		if err := oaq.loadOrder(ctx, query, nodes, nil,
			func(n *OrderAmount, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oaq *OrderAmountQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*OrderAmount, init func(*OrderAmount), assign func(*OrderAmount, *Order)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*OrderAmount)
	for i := range nodes {
		fk := nodes[i].OrderID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(order.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (oaq *OrderAmountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oaq.querySpec()
	_spec.Node.Columns = oaq.ctx.Fields
	if len(oaq.ctx.Fields) > 0 {
		_spec.Unique = oaq.ctx.Unique != nil && *oaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oaq.driver, _spec)
}

func (oaq *OrderAmountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderamount.Table, orderamount.Columns, sqlgraph.NewFieldSpec(orderamount.FieldID, field.TypeInt64))
	_spec.From = oaq.sql
	if unique := oaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oaq.path != nil {
		_spec.Unique = true
	}
	if fields := oaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderamount.FieldID)
		for i := range fields {
			if fields[i] != orderamount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if oaq.withOrder != nil {
			_spec.Node.AddColumnOnce(orderamount.FieldOrderID)
		}
	}
	if ps := oaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oaq *OrderAmountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oaq.driver.Dialect())
	t1 := builder.Table(orderamount.Table)
	columns := oaq.ctx.Fields
	if len(columns) == 0 {
		columns = orderamount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oaq.sql != nil {
		selector = oaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oaq.ctx.Unique != nil && *oaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oaq.predicates {
		p(selector)
	}
	for _, p := range oaq.order {
		p(selector)
	}
	if offset := oaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderAmountGroupBy is the group-by builder for OrderAmount entities.
type OrderAmountGroupBy struct {
	selector
	build *OrderAmountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oagb *OrderAmountGroupBy) Aggregate(fns ...AggregateFunc) *OrderAmountGroupBy {
	oagb.fns = append(oagb.fns, fns...)
	return oagb
}

// Scan applies the selector query and scans the result into the given value.
func (oagb *OrderAmountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oagb.build.ctx, ent.OpQueryGroupBy)
	if err := oagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderAmountQuery, *OrderAmountGroupBy](ctx, oagb.build, oagb, oagb.build.inters, v)
}

func (oagb *OrderAmountGroupBy) sqlScan(ctx context.Context, root *OrderAmountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oagb.fns))
	for _, fn := range oagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oagb.flds)+len(oagb.fns))
		for _, f := range *oagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderAmountSelect is the builder for selecting fields of OrderAmount entities.
type OrderAmountSelect struct {
	*OrderAmountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oas *OrderAmountSelect) Aggregate(fns ...AggregateFunc) *OrderAmountSelect {
	oas.fns = append(oas.fns, fns...)
	return oas
}

// Scan applies the selector query and scans the result into the given value.
func (oas *OrderAmountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oas.ctx, ent.OpQuerySelect)
	if err := oas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderAmountQuery, *OrderAmountSelect](ctx, oas.OrderAmountQuery, oas, oas.inters, v)
}

func (oas *OrderAmountSelect) sqlScan(ctx context.Context, root *OrderAmountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oas.fns))
	for _, fn := range oas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
