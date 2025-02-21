// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/contest"
	"saas/biz/dal/db/ent/contestparticipant"
	"saas/biz/dal/db/ent/internal"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContestQuery is the builder for querying Contest entities.
type ContestQuery struct {
	config
	ctx                     *QueryContext
	order                   []contest.OrderOption
	inters                  []Interceptor
	predicates              []predicate.Contest
	withContestParticipants *ContestParticipantQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContestQuery builder.
func (cq *ContestQuery) Where(ps ...predicate.Contest) *ContestQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *ContestQuery) Limit(limit int) *ContestQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *ContestQuery) Offset(offset int) *ContestQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ContestQuery) Unique(unique bool) *ContestQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *ContestQuery) Order(o ...contest.OrderOption) *ContestQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryContestParticipants chains the current query on the "contest_participants" edge.
func (cq *ContestQuery) QueryContestParticipants() *ContestParticipantQuery {
	query := (&ContestParticipantClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contest.Table, contest.FieldID, selector),
			sqlgraph.To(contestparticipant.Table, contestparticipant.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, contest.ContestParticipantsTable, contest.ContestParticipantsColumn),
		)
		schemaConfig := cq.schemaConfig
		step.To.Schema = schemaConfig.ContestParticipant
		step.Edge.Schema = schemaConfig.ContestParticipant
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Contest entity from the query.
// Returns a *NotFoundError when no Contest was found.
func (cq *ContestQuery) First(ctx context.Context) (*Contest, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ContestQuery) FirstX(ctx context.Context) *Contest {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Contest ID from the query.
// Returns a *NotFoundError when no Contest ID was found.
func (cq *ContestQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ContestQuery) FirstIDX(ctx context.Context) int64 {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Contest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Contest entity is found.
// Returns a *NotFoundError when no Contest entities are found.
func (cq *ContestQuery) Only(ctx context.Context) (*Contest, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contest.Label}
	default:
		return nil, &NotSingularError{contest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ContestQuery) OnlyX(ctx context.Context) *Contest {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Contest ID in the query.
// Returns a *NotSingularError when more than one Contest ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ContestQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contest.Label}
	default:
		err = &NotSingularError{contest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ContestQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Contests.
func (cq *ContestQuery) All(ctx context.Context) ([]*Contest, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Contest, *ContestQuery]()
	return withInterceptors[[]*Contest](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *ContestQuery) AllX(ctx context.Context) []*Contest {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Contest IDs.
func (cq *ContestQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(contest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ContestQuery) IDsX(ctx context.Context) []int64 {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ContestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*ContestQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ContestQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ContestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
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
func (cq *ContestQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ContestQuery) Clone() *ContestQuery {
	if cq == nil {
		return nil
	}
	return &ContestQuery{
		config:                  cq.config,
		ctx:                     cq.ctx.Clone(),
		order:                   append([]contest.OrderOption{}, cq.order...),
		inters:                  append([]Interceptor{}, cq.inters...),
		predicates:              append([]predicate.Contest{}, cq.predicates...),
		withContestParticipants: cq.withContestParticipants.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithContestParticipants tells the query-builder to eager-load the nodes that are connected to
// the "contest_participants" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ContestQuery) WithContestParticipants(opts ...func(*ContestParticipantQuery)) *ContestQuery {
	query := (&ContestParticipantClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withContestParticipants = query
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
//	client.Contest.Query().
//		GroupBy(contest.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *ContestQuery) GroupBy(field string, fields ...string) *ContestGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContestGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = contest.Label
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
//	client.Contest.Query().
//		Select(contest.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *ContestQuery) Select(fields ...string) *ContestSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &ContestSelect{ContestQuery: cq}
	sbuild.label = contest.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContestSelect configured with the given aggregations.
func (cq *ContestQuery) Aggregate(fns ...AggregateFunc) *ContestSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *ContestQuery) prepareQuery(ctx context.Context) error {
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
		if !contest.ValidColumn(f) {
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

func (cq *ContestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Contest, error) {
	var (
		nodes       = []*Contest{}
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withContestParticipants != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Contest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Contest{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = cq.schemaConfig.Contest
	ctx = internal.NewSchemaConfigContext(ctx, cq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withContestParticipants; query != nil {
		if err := cq.loadContestParticipants(ctx, query, nodes,
			func(n *Contest) { n.Edges.ContestParticipants = []*ContestParticipant{} },
			func(n *Contest, e *ContestParticipant) {
				n.Edges.ContestParticipants = append(n.Edges.ContestParticipants, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ContestQuery) loadContestParticipants(ctx context.Context, query *ContestParticipantQuery, nodes []*Contest, init func(*Contest), assign func(*Contest, *ContestParticipant)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Contest)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(contestparticipant.FieldContestID)
	}
	query.Where(predicate.ContestParticipant(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(contest.ContestParticipantsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ContestID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "contest_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ContestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Schema = cq.schemaConfig.Contest
	ctx = internal.NewSchemaConfigContext(ctx, cq.schemaConfig)
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ContestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contest.FieldID)
		for i := range fields {
			if fields[i] != contest.FieldID {
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

func (cq *ContestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(contest.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = contest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(cq.schemaConfig.Contest)
	ctx = internal.NewSchemaConfigContext(ctx, cq.schemaConfig)
	selector.WithContext(ctx)
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

// ContestGroupBy is the group-by builder for Contest entities.
type ContestGroupBy struct {
	selector
	build *ContestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ContestGroupBy) Aggregate(fns ...AggregateFunc) *ContestGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *ContestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContestQuery, *ContestGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *ContestGroupBy) sqlScan(ctx context.Context, root *ContestQuery, v any) error {
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

// ContestSelect is the builder for selecting fields of Contest entities.
type ContestSelect struct {
	*ContestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *ContestSelect) Aggregate(fns ...AggregateFunc) *ContestSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ContestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContestQuery, *ContestSelect](ctx, cs.ContestQuery, cs, cs.inters, v)
}

func (cs *ContestSelect) sqlScan(ctx context.Context, root *ContestQuery, v any) error {
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
