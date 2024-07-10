// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/schedule"
	"saas/pkg/db/ent/schedulecoach"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleCoachQuery is the builder for querying ScheduleCoach entities.
type ScheduleCoachQuery struct {
	config
	ctx          *QueryContext
	order        []schedulecoach.OrderOption
	inters       []Interceptor
	predicates   []predicate.ScheduleCoach
	withSchedule *ScheduleQuery
	modifiers    []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScheduleCoachQuery builder.
func (scq *ScheduleCoachQuery) Where(ps ...predicate.ScheduleCoach) *ScheduleCoachQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ScheduleCoachQuery) Limit(limit int) *ScheduleCoachQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ScheduleCoachQuery) Offset(offset int) *ScheduleCoachQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ScheduleCoachQuery) Unique(unique bool) *ScheduleCoachQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ScheduleCoachQuery) Order(o ...schedulecoach.OrderOption) *ScheduleCoachQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QuerySchedule chains the current query on the "schedule" edge.
func (scq *ScheduleCoachQuery) QuerySchedule() *ScheduleQuery {
	query := (&ScheduleClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(schedulecoach.Table, schedulecoach.FieldID, selector),
			sqlgraph.To(schedule.Table, schedule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, schedulecoach.ScheduleTable, schedulecoach.ScheduleColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ScheduleCoach entity from the query.
// Returns a *NotFoundError when no ScheduleCoach was found.
func (scq *ScheduleCoachQuery) First(ctx context.Context) (*ScheduleCoach, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{schedulecoach.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ScheduleCoachQuery) FirstX(ctx context.Context) *ScheduleCoach {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScheduleCoach ID from the query.
// Returns a *NotFoundError when no ScheduleCoach ID was found.
func (scq *ScheduleCoachQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{schedulecoach.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ScheduleCoachQuery) FirstIDX(ctx context.Context) int64 {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScheduleCoach entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ScheduleCoach entity is found.
// Returns a *NotFoundError when no ScheduleCoach entities are found.
func (scq *ScheduleCoachQuery) Only(ctx context.Context) (*ScheduleCoach, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{schedulecoach.Label}
	default:
		return nil, &NotSingularError{schedulecoach.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ScheduleCoachQuery) OnlyX(ctx context.Context) *ScheduleCoach {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScheduleCoach ID in the query.
// Returns a *NotSingularError when more than one ScheduleCoach ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ScheduleCoachQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{schedulecoach.Label}
	default:
		err = &NotSingularError{schedulecoach.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ScheduleCoachQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScheduleCoaches.
func (scq *ScheduleCoachQuery) All(ctx context.Context) ([]*ScheduleCoach, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ScheduleCoach, *ScheduleCoachQuery]()
	return withInterceptors[[]*ScheduleCoach](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ScheduleCoachQuery) AllX(ctx context.Context) []*ScheduleCoach {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScheduleCoach IDs.
func (scq *ScheduleCoachQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(schedulecoach.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ScheduleCoachQuery) IDsX(ctx context.Context) []int64 {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ScheduleCoachQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ScheduleCoachQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ScheduleCoachQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ScheduleCoachQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, "Exist")
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *ScheduleCoachQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScheduleCoachQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ScheduleCoachQuery) Clone() *ScheduleCoachQuery {
	if scq == nil {
		return nil
	}
	return &ScheduleCoachQuery{
		config:       scq.config,
		ctx:          scq.ctx.Clone(),
		order:        append([]schedulecoach.OrderOption{}, scq.order...),
		inters:       append([]Interceptor{}, scq.inters...),
		predicates:   append([]predicate.ScheduleCoach{}, scq.predicates...),
		withSchedule: scq.withSchedule.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithSchedule tells the query-builder to eager-load the nodes that are connected to
// the "schedule" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ScheduleCoachQuery) WithSchedule(opts ...func(*ScheduleQuery)) *ScheduleCoachQuery {
	query := (&ScheduleClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withSchedule = query
	return scq
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
//	client.ScheduleCoach.Query().
//		GroupBy(schedulecoach.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ScheduleCoachQuery) GroupBy(field string, fields ...string) *ScheduleCoachGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScheduleCoachGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = schedulecoach.Label
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
//	client.ScheduleCoach.Query().
//		Select(schedulecoach.FieldCreatedAt).
//		Scan(ctx, &v)
func (scq *ScheduleCoachQuery) Select(fields ...string) *ScheduleCoachSelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ScheduleCoachSelect{ScheduleCoachQuery: scq}
	sbuild.label = schedulecoach.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScheduleCoachSelect configured with the given aggregations.
func (scq *ScheduleCoachQuery) Aggregate(fns ...AggregateFunc) *ScheduleCoachSelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ScheduleCoachQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !schedulecoach.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *ScheduleCoachQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ScheduleCoach, error) {
	var (
		nodes       = []*ScheduleCoach{}
		_spec       = scq.querySpec()
		loadedTypes = [1]bool{
			scq.withSchedule != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ScheduleCoach).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ScheduleCoach{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withSchedule; query != nil {
		if err := scq.loadSchedule(ctx, query, nodes, nil,
			func(n *ScheduleCoach, e *Schedule) { n.Edges.Schedule = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ScheduleCoachQuery) loadSchedule(ctx context.Context, query *ScheduleQuery, nodes []*ScheduleCoach, init func(*ScheduleCoach), assign func(*ScheduleCoach, *Schedule)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ScheduleCoach)
	for i := range nodes {
		fk := nodes[i].ScheduleID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(schedule.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "schedule_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (scq *ScheduleCoachQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *ScheduleCoachQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(schedulecoach.Table, schedulecoach.Columns, sqlgraph.NewFieldSpec(schedulecoach.FieldID, field.TypeInt64))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedulecoach.FieldID)
		for i := range fields {
			if fields[i] != schedulecoach.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if scq.withSchedule != nil {
			_spec.Node.AddColumnOnce(schedulecoach.FieldScheduleID)
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *ScheduleCoachQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(schedulecoach.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = schedulecoach.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range scq.modifiers {
		m(selector)
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scq *ScheduleCoachQuery) Modify(modifiers ...func(s *sql.Selector)) *ScheduleCoachSelect {
	scq.modifiers = append(scq.modifiers, modifiers...)
	return scq.Select()
}

// ScheduleCoachGroupBy is the group-by builder for ScheduleCoach entities.
type ScheduleCoachGroupBy struct {
	selector
	build *ScheduleCoachQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ScheduleCoachGroupBy) Aggregate(fns ...AggregateFunc) *ScheduleCoachGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ScheduleCoachGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScheduleCoachQuery, *ScheduleCoachGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ScheduleCoachGroupBy) sqlScan(ctx context.Context, root *ScheduleCoachQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScheduleCoachSelect is the builder for selecting fields of ScheduleCoach entities.
type ScheduleCoachSelect struct {
	*ScheduleCoachQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ScheduleCoachSelect) Aggregate(fns ...AggregateFunc) *ScheduleCoachSelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ScheduleCoachSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScheduleCoachQuery, *ScheduleCoachSelect](ctx, scs.ScheduleCoachQuery, scs, scs.inters, v)
}

func (scs *ScheduleCoachSelect) sqlScan(ctx context.Context, root *ScheduleCoachQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scs *ScheduleCoachSelect) Modify(modifiers ...func(s *sql.Selector)) *ScheduleCoachSelect {
	scs.modifiers = append(scs.modifiers, modifiers...)
	return scs
}
