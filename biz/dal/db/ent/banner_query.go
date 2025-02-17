// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/banner"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BannerQuery is the builder for querying Banner entities.
type BannerQuery struct {
	config
	ctx        *QueryContext
	order      []banner.OrderOption
	inters     []Interceptor
	predicates []predicate.Banner
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BannerQuery builder.
func (bq *BannerQuery) Where(ps ...predicate.Banner) *BannerQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BannerQuery) Limit(limit int) *BannerQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BannerQuery) Offset(offset int) *BannerQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BannerQuery) Unique(unique bool) *BannerQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BannerQuery) Order(o ...banner.OrderOption) *BannerQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// First returns the first Banner entity from the query.
// Returns a *NotFoundError when no Banner was found.
func (bq *BannerQuery) First(ctx context.Context) (*Banner, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{banner.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BannerQuery) FirstX(ctx context.Context) *Banner {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Banner ID from the query.
// Returns a *NotFoundError when no Banner ID was found.
func (bq *BannerQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{banner.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BannerQuery) FirstIDX(ctx context.Context) int64 {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Banner entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Banner entity is found.
// Returns a *NotFoundError when no Banner entities are found.
func (bq *BannerQuery) Only(ctx context.Context) (*Banner, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{banner.Label}
	default:
		return nil, &NotSingularError{banner.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BannerQuery) OnlyX(ctx context.Context) *Banner {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Banner ID in the query.
// Returns a *NotSingularError when more than one Banner ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BannerQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{banner.Label}
	default:
		err = &NotSingularError{banner.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BannerQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Banners.
func (bq *BannerQuery) All(ctx context.Context) ([]*Banner, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryAll)
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Banner, *BannerQuery]()
	return withInterceptors[[]*Banner](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BannerQuery) AllX(ctx context.Context) []*Banner {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Banner IDs.
func (bq *BannerQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryIDs)
	if err = bq.Select(banner.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BannerQuery) IDsX(ctx context.Context) []int64 {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BannerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryCount)
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BannerQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BannerQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BannerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryExist)
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BannerQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BannerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BannerQuery) Clone() *BannerQuery {
	if bq == nil {
		return nil
	}
	return &BannerQuery{
		config:     bq.config,
		ctx:        bq.ctx.Clone(),
		order:      append([]banner.OrderOption{}, bq.order...),
		inters:     append([]Interceptor{}, bq.inters...),
		predicates: append([]predicate.Banner{}, bq.predicates...),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
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
//	client.Banner.Query().
//		GroupBy(banner.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BannerQuery) GroupBy(field string, fields ...string) *BannerGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BannerGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = banner.Label
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
//	client.Banner.Query().
//		Select(banner.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BannerQuery) Select(fields ...string) *BannerSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BannerSelect{BannerQuery: bq}
	sbuild.label = banner.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BannerSelect configured with the given aggregations.
func (bq *BannerQuery) Aggregate(fns ...AggregateFunc) *BannerSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BannerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !banner.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BannerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Banner, error) {
	var (
		nodes = []*Banner{}
		_spec = bq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Banner).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Banner{config: bq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bq *BannerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BannerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(banner.Table, banner.Columns, sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, banner.FieldID)
		for i := range fields {
			if fields[i] != banner.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BannerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(banner.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = banner.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BannerGroupBy is the group-by builder for Banner entities.
type BannerGroupBy struct {
	selector
	build *BannerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BannerGroupBy) Aggregate(fns ...AggregateFunc) *BannerGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BannerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, ent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BannerQuery, *BannerGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BannerGroupBy) sqlScan(ctx context.Context, root *BannerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BannerSelect is the builder for selecting fields of Banner entities.
type BannerSelect struct {
	*BannerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BannerSelect) Aggregate(fns ...AggregateFunc) *BannerSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BannerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, ent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BannerQuery, *BannerSelect](ctx, bs.BannerQuery, bs, bs.inters, v)
}

func (bs *BannerSelect) sqlScan(ctx context.Context, root *BannerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
