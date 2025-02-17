// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/community"
	"saas/biz/dal/db/ent/communityparticipant"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommunityQuery is the builder for querying Community entities.
type CommunityQuery struct {
	config
	ctx                       *QueryContext
	order                     []community.OrderOption
	inters                    []Interceptor
	predicates                []predicate.Community
	withCommunityParticipants *CommunityParticipantQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommunityQuery builder.
func (cq *CommunityQuery) Where(ps ...predicate.Community) *CommunityQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CommunityQuery) Limit(limit int) *CommunityQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CommunityQuery) Offset(offset int) *CommunityQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CommunityQuery) Unique(unique bool) *CommunityQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CommunityQuery) Order(o ...community.OrderOption) *CommunityQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryCommunityParticipants chains the current query on the "community_participants" edge.
func (cq *CommunityQuery) QueryCommunityParticipants() *CommunityParticipantQuery {
	query := (&CommunityParticipantClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(community.Table, community.FieldID, selector),
			sqlgraph.To(communityparticipant.Table, communityparticipant.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, community.CommunityParticipantsTable, community.CommunityParticipantsColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Community entity from the query.
// Returns a *NotFoundError when no Community was found.
func (cq *CommunityQuery) First(ctx context.Context) (*Community, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{community.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommunityQuery) FirstX(ctx context.Context) *Community {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Community ID from the query.
// Returns a *NotFoundError when no Community ID was found.
func (cq *CommunityQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{community.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CommunityQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Community entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Community entity is found.
// Returns a *NotFoundError when no Community entities are found.
func (cq *CommunityQuery) Only(ctx context.Context) (*Community, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{community.Label}
	default:
		return nil, &NotSingularError{community.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommunityQuery) OnlyX(ctx context.Context) *Community {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Community ID in the query.
// Returns a *NotSingularError when more than one Community ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CommunityQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{community.Label}
	default:
		err = &NotSingularError{community.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CommunityQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Communities.
func (cq *CommunityQuery) All(ctx context.Context) ([]*Community, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryAll)
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Community, *CommunityQuery]()
	return withInterceptors[[]*Community](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommunityQuery) AllX(ctx context.Context) []*Community {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Community IDs.
func (cq *CommunityQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryIDs)
	if err = cq.Select(community.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommunityQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommunityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryCount)
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CommunityQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommunityQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommunityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, ent.OpQueryExist)
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommunityQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommunityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommunityQuery) Clone() *CommunityQuery {
	if cq == nil {
		return nil
	}
	return &CommunityQuery{
		config:                    cq.config,
		ctx:                       cq.ctx.Clone(),
		order:                     append([]community.OrderOption{}, cq.order...),
		inters:                    append([]Interceptor{}, cq.inters...),
		predicates:                append([]predicate.Community{}, cq.predicates...),
		withCommunityParticipants: cq.withCommunityParticipants.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithCommunityParticipants tells the query-builder to eager-load the nodes that are connected to
// the "community_participants" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CommunityQuery) WithCommunityParticipants(opts ...func(*CommunityParticipantQuery)) *CommunityQuery {
	query := (&CommunityParticipantClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withCommunityParticipants = query
	return cq
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
//	client.Community.Query().
//		GroupBy(community.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CommunityQuery) GroupBy(field string, fields ...string) *CommunityGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommunityGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = community.Label
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
//	client.Community.Query().
//		Select(community.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CommunityQuery) Select(fields ...string) *CommunitySelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CommunitySelect{CommunityQuery: cq}
	sbuild.label = community.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommunitySelect configured with the given aggregations.
func (cq *CommunityQuery) Aggregate(fns ...AggregateFunc) *CommunitySelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CommunityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !community.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CommunityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Community, error) {
	var (
		nodes       = []*Community{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withCommunityParticipants != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Community).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Community{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withCommunityParticipants; query != nil {
		if err := cq.loadCommunityParticipants(ctx, query, nodes,
			func(n *Community) { n.Edges.CommunityParticipants = []*CommunityParticipant{} },
			func(n *Community, e *CommunityParticipant) {
				n.Edges.CommunityParticipants = append(n.Edges.CommunityParticipants, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CommunityQuery) loadCommunityParticipants(ctx context.Context, query *CommunityParticipantQuery, nodes []*Community, init func(*Community), assign func(*Community, *CommunityParticipant)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Community)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(communityparticipant.FieldCommunityID)
	}
	query.Where(predicate.CommunityParticipant(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(community.CommunityParticipantsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CommunityID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "community_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CommunityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CommunityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(community.Table, community.Columns, sqlgraph.NewFieldSpec(community.FieldID, field.TypeInt64))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, community.FieldID)
		for i := range fields {
			if fields[i] != community.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CommunityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(community.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = community.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommunityGroupBy is the group-by builder for Community entities.
type CommunityGroupBy struct {
	selector
	build *CommunityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommunityGroupBy) Aggregate(fns ...AggregateFunc) *CommunityGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommunityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, ent.OpQueryGroupBy)
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommunityQuery, *CommunityGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CommunityGroupBy) sqlScan(ctx context.Context, root *CommunityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommunitySelect is the builder for selecting fields of Community entities.
type CommunitySelect struct {
	*CommunityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommunitySelect) Aggregate(fns ...AggregateFunc) *CommunitySelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommunitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, ent.OpQuerySelect)
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommunityQuery, *CommunitySelect](ctx, cs.CommunityQuery, cs, cs.inters, v)
}

func (cs *CommunitySelect) sqlScan(ctx context.Context, root *CommunityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
