// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/logs"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LogsQuery is the builder for querying Logs entities.
type LogsQuery struct {
	config
	ctx        *QueryContext
	order      []logs.OrderOption
	inters     []Interceptor
	predicates []predicate.Logs
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogsQuery builder.
func (lq *LogsQuery) Where(ps ...predicate.Logs) *LogsQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit the number of records to be returned by this query.
func (lq *LogsQuery) Limit(limit int) *LogsQuery {
	lq.ctx.Limit = &limit
	return lq
}

// Offset to start from.
func (lq *LogsQuery) Offset(offset int) *LogsQuery {
	lq.ctx.Offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LogsQuery) Unique(unique bool) *LogsQuery {
	lq.ctx.Unique = &unique
	return lq
}

// Order specifies how the records should be ordered.
func (lq *LogsQuery) Order(o ...logs.OrderOption) *LogsQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// First returns the first Logs entity from the query.
// Returns a *NotFoundError when no Logs was found.
func (lq *LogsQuery) First(ctx context.Context) (*Logs, error) {
	nodes, err := lq.Limit(1).All(setContextOp(ctx, lq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LogsQuery) FirstX(ctx context.Context) *Logs {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Logs ID from the query.
// Returns a *NotFoundError when no Logs ID was found.
func (lq *LogsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lq.Limit(1).IDs(setContextOp(ctx, lq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LogsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Logs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Logs entity is found.
// Returns a *NotFoundError when no Logs entities are found.
func (lq *LogsQuery) Only(ctx context.Context) (*Logs, error) {
	nodes, err := lq.Limit(2).All(setContextOp(ctx, lq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logs.Label}
	default:
		return nil, &NotSingularError{logs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LogsQuery) OnlyX(ctx context.Context) *Logs {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Logs ID in the query.
// Returns a *NotSingularError when more than one Logs ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LogsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = lq.Limit(2).IDs(setContextOp(ctx, lq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logs.Label}
	default:
		err = &NotSingularError{logs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LogsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogsSlice.
func (lq *LogsQuery) All(ctx context.Context) ([]*Logs, error) {
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryAll)
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Logs, *LogsQuery]()
	return withInterceptors[[]*Logs](ctx, lq, qr, lq.inters)
}

// AllX is like All, but panics if an error occurs.
func (lq *LogsQuery) AllX(ctx context.Context) []*Logs {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Logs IDs.
func (lq *LogsQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if lq.ctx.Unique == nil && lq.path != nil {
		lq.Unique(true)
	}
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryIDs)
	if err = lq.Select(logs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LogsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LogsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryCount)
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, lq, querierCount[*LogsQuery](), lq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LogsQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LogsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, lq.ctx, ent.OpQueryExist)
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LogsQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LogsQuery) Clone() *LogsQuery {
	if lq == nil {
		return nil
	}
	return &LogsQuery{
		config:     lq.config,
		ctx:        lq.ctx.Clone(),
		order:      append([]logs.OrderOption{}, lq.order...),
		inters:     append([]Interceptor{}, lq.inters...),
		predicates: append([]predicate.Logs{}, lq.predicates...),
		// clone intermediate query.
		sql:  lq.sql.Clone(),
		path: lq.path,
	}
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
//	client.Logs.Query().
//		GroupBy(logs.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LogsQuery) GroupBy(field string, fields ...string) *LogsGroupBy {
	lq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &LogsGroupBy{build: lq}
	grbuild.flds = &lq.ctx.Fields
	grbuild.label = logs.Label
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
//	client.Logs.Query().
//		Select(logs.FieldCreatedAt).
//		Scan(ctx, &v)
func (lq *LogsQuery) Select(fields ...string) *LogsSelect {
	lq.ctx.Fields = append(lq.ctx.Fields, fields...)
	sbuild := &LogsSelect{LogsQuery: lq}
	sbuild.label = logs.Label
	sbuild.flds, sbuild.scan = &lq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a LogsSelect configured with the given aggregations.
func (lq *LogsQuery) Aggregate(fns ...AggregateFunc) *LogsSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LogsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range lq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, lq); err != nil {
				return err
			}
		}
	}
	for _, f := range lq.ctx.Fields {
		if !logs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LogsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Logs, error) {
	var (
		nodes = []*Logs{}
		_spec = lq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Logs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Logs{config: lq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lq *LogsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	_spec.Node.Columns = lq.ctx.Fields
	if len(lq.ctx.Fields) > 0 {
		_spec.Unique = lq.ctx.Unique != nil && *lq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LogsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(logs.Table, logs.Columns, sqlgraph.NewFieldSpec(logs.FieldID, field.TypeInt64))
	_spec.From = lq.sql
	if unique := lq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if lq.path != nil {
		_spec.Unique = true
	}
	if fields := lq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logs.FieldID)
		for i := range fields {
			if fields[i] != logs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LogsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(logs.Table)
	columns := lq.ctx.Fields
	if len(columns) == 0 {
		columns = logs.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.ctx.Unique != nil && *lq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LogsGroupBy is the group-by builder for Logs entities.
type LogsGroupBy struct {
	selector
	build *LogsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LogsGroupBy) Aggregate(fns ...AggregateFunc) *LogsGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the selector query and scans the result into the given value.
func (lgb *LogsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, lgb.build.ctx, ent.OpQueryGroupBy)
	if err := lgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogsQuery, *LogsGroupBy](ctx, lgb.build, lgb, lgb.build.inters, v)
}

func (lgb *LogsGroupBy) sqlScan(ctx context.Context, root *LogsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*lgb.flds)+len(lgb.fns))
		for _, f := range *lgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*lgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// LogsSelect is the builder for selecting fields of Logs entities.
type LogsSelect struct {
	*LogsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LogsSelect) Aggregate(fns ...AggregateFunc) *LogsSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LogsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ls.ctx, ent.OpQuerySelect)
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*LogsQuery, *LogsSelect](ctx, ls.LogsQuery, ls, ls.inters, v)
}

func (ls *LogsSelect) sqlScan(ctx context.Context, root *LogsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
