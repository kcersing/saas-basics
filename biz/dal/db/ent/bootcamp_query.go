// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/bootcamp"
	"saas/biz/dal/db/ent/bootcampparticipant"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BootcampQuery is the builder for querying Bootcamp entities.
type BootcampQuery struct {
	config
	ctx                      *QueryContext
	order                    []bootcamp.OrderOption
	inters                   []Interceptor
	predicates               []predicate.Bootcamp
	withBootcampParticipants *BootcampParticipantQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BootcampQuery builder.
func (bq *BootcampQuery) Where(ps ...predicate.Bootcamp) *BootcampQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BootcampQuery) Limit(limit int) *BootcampQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BootcampQuery) Offset(offset int) *BootcampQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BootcampQuery) Unique(unique bool) *BootcampQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BootcampQuery) Order(o ...bootcamp.OrderOption) *BootcampQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryBootcampParticipants chains the current query on the "bootcamp_participants" edge.
func (bq *BootcampQuery) QueryBootcampParticipants() *BootcampParticipantQuery {
	query := (&BootcampParticipantClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bootcamp.Table, bootcamp.FieldID, selector),
			sqlgraph.To(bootcampparticipant.Table, bootcampparticipant.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bootcamp.BootcampParticipantsTable, bootcamp.BootcampParticipantsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bootcamp entity from the query.
// Returns a *NotFoundError when no Bootcamp was found.
func (bq *BootcampQuery) First(ctx context.Context) (*Bootcamp, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bootcamp.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BootcampQuery) FirstX(ctx context.Context) *Bootcamp {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bootcamp ID from the query.
// Returns a *NotFoundError when no Bootcamp ID was found.
func (bq *BootcampQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bootcamp.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BootcampQuery) FirstIDX(ctx context.Context) int64 {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bootcamp entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bootcamp entity is found.
// Returns a *NotFoundError when no Bootcamp entities are found.
func (bq *BootcampQuery) Only(ctx context.Context) (*Bootcamp, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bootcamp.Label}
	default:
		return nil, &NotSingularError{bootcamp.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BootcampQuery) OnlyX(ctx context.Context) *Bootcamp {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bootcamp ID in the query.
// Returns a *NotSingularError when more than one Bootcamp ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BootcampQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bootcamp.Label}
	default:
		err = &NotSingularError{bootcamp.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BootcampQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Bootcamps.
func (bq *BootcampQuery) All(ctx context.Context) ([]*Bootcamp, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryAll)
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bootcamp, *BootcampQuery]()
	return withInterceptors[[]*Bootcamp](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BootcampQuery) AllX(ctx context.Context) []*Bootcamp {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bootcamp IDs.
func (bq *BootcampQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryIDs)
	if err = bq.Select(bootcamp.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BootcampQuery) IDsX(ctx context.Context) []int64 {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BootcampQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, ent.OpQueryCount)
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BootcampQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BootcampQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BootcampQuery) Exist(ctx context.Context) (bool, error) {
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
func (bq *BootcampQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BootcampQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BootcampQuery) Clone() *BootcampQuery {
	if bq == nil {
		return nil
	}
	return &BootcampQuery{
		config:                   bq.config,
		ctx:                      bq.ctx.Clone(),
		order:                    append([]bootcamp.OrderOption{}, bq.order...),
		inters:                   append([]Interceptor{}, bq.inters...),
		predicates:               append([]predicate.Bootcamp{}, bq.predicates...),
		withBootcampParticipants: bq.withBootcampParticipants.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithBootcampParticipants tells the query-builder to eager-load the nodes that are connected to
// the "bootcamp_participants" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BootcampQuery) WithBootcampParticipants(opts ...func(*BootcampParticipantQuery)) *BootcampQuery {
	query := (&BootcampParticipantClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withBootcampParticipants = query
	return bq
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
//	client.Bootcamp.Query().
//		GroupBy(bootcamp.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BootcampQuery) GroupBy(field string, fields ...string) *BootcampGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BootcampGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bootcamp.Label
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
//	client.Bootcamp.Query().
//		Select(bootcamp.FieldCreatedAt).
//		Scan(ctx, &v)
func (bq *BootcampQuery) Select(fields ...string) *BootcampSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BootcampSelect{BootcampQuery: bq}
	sbuild.label = bootcamp.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BootcampSelect configured with the given aggregations.
func (bq *BootcampQuery) Aggregate(fns ...AggregateFunc) *BootcampSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BootcampQuery) prepareQuery(ctx context.Context) error {
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
		if !bootcamp.ValidColumn(f) {
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

func (bq *BootcampQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bootcamp, error) {
	var (
		nodes       = []*Bootcamp{}
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withBootcampParticipants != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bootcamp).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bootcamp{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
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
	if query := bq.withBootcampParticipants; query != nil {
		if err := bq.loadBootcampParticipants(ctx, query, nodes,
			func(n *Bootcamp) { n.Edges.BootcampParticipants = []*BootcampParticipant{} },
			func(n *Bootcamp, e *BootcampParticipant) {
				n.Edges.BootcampParticipants = append(n.Edges.BootcampParticipants, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BootcampQuery) loadBootcampParticipants(ctx context.Context, query *BootcampParticipantQuery, nodes []*Bootcamp, init func(*Bootcamp), assign func(*Bootcamp, *BootcampParticipant)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Bootcamp)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(bootcampparticipant.FieldBootcampID)
	}
	query.Where(predicate.BootcampParticipant(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(bootcamp.BootcampParticipantsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BootcampID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "bootcamp_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BootcampQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BootcampQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bootcamp.Table, bootcamp.Columns, sqlgraph.NewFieldSpec(bootcamp.FieldID, field.TypeInt64))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bootcamp.FieldID)
		for i := range fields {
			if fields[i] != bootcamp.FieldID {
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

func (bq *BootcampQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bootcamp.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bootcamp.Columns
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

// BootcampGroupBy is the group-by builder for Bootcamp entities.
type BootcampGroupBy struct {
	selector
	build *BootcampQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BootcampGroupBy) Aggregate(fns ...AggregateFunc) *BootcampGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BootcampGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, ent.OpQueryGroupBy)
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BootcampQuery, *BootcampGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BootcampGroupBy) sqlScan(ctx context.Context, root *BootcampQuery, v any) error {
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

// BootcampSelect is the builder for selecting fields of Bootcamp entities.
type BootcampSelect struct {
	*BootcampQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BootcampSelect) Aggregate(fns ...AggregateFunc) *BootcampSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BootcampSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, ent.OpQuerySelect)
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BootcampQuery, *BootcampSelect](ctx, bs.BootcampQuery, bs, bs.inters, v)
}

func (bs *BootcampSelect) sqlScan(ctx context.Context, root *BootcampQuery, v any) error {
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
