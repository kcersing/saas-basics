// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/internal"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/schedulemember"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleMemberQuery is the builder for querying ScheduleMember entities.
type ScheduleMemberQuery struct {
	config
	ctx          *QueryContext
	order        []schedulemember.OrderOption
	inters       []Interceptor
	predicates   []predicate.ScheduleMember
	withSchedule *ScheduleQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ScheduleMemberQuery builder.
func (smq *ScheduleMemberQuery) Where(ps ...predicate.ScheduleMember) *ScheduleMemberQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit the number of records to be returned by this query.
func (smq *ScheduleMemberQuery) Limit(limit int) *ScheduleMemberQuery {
	smq.ctx.Limit = &limit
	return smq
}

// Offset to start from.
func (smq *ScheduleMemberQuery) Offset(offset int) *ScheduleMemberQuery {
	smq.ctx.Offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *ScheduleMemberQuery) Unique(unique bool) *ScheduleMemberQuery {
	smq.ctx.Unique = &unique
	return smq
}

// Order specifies how the records should be ordered.
func (smq *ScheduleMemberQuery) Order(o ...schedulemember.OrderOption) *ScheduleMemberQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// QuerySchedule chains the current query on the "schedule" edge.
func (smq *ScheduleMemberQuery) QuerySchedule() *ScheduleQuery {
	query := (&ScheduleClient{config: smq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(schedulemember.Table, schedulemember.FieldID, selector),
			sqlgraph.To(schedule.Table, schedule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, schedulemember.ScheduleTable, schedulemember.ScheduleColumn),
		)
		schemaConfig := smq.schemaConfig
		step.To.Schema = schemaConfig.Schedule
		step.Edge.Schema = schemaConfig.ScheduleMember
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ScheduleMember entity from the query.
// Returns a *NotFoundError when no ScheduleMember was found.
func (smq *ScheduleMemberQuery) First(ctx context.Context) (*ScheduleMember, error) {
	nodes, err := smq.Limit(1).All(setContextOp(ctx, smq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{schedulemember.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *ScheduleMemberQuery) FirstX(ctx context.Context) *ScheduleMember {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ScheduleMember ID from the query.
// Returns a *NotFoundError when no ScheduleMember ID was found.
func (smq *ScheduleMemberQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = smq.Limit(1).IDs(setContextOp(ctx, smq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{schedulemember.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *ScheduleMemberQuery) FirstIDX(ctx context.Context) int64 {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ScheduleMember entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ScheduleMember entity is found.
// Returns a *NotFoundError when no ScheduleMember entities are found.
func (smq *ScheduleMemberQuery) Only(ctx context.Context) (*ScheduleMember, error) {
	nodes, err := smq.Limit(2).All(setContextOp(ctx, smq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{schedulemember.Label}
	default:
		return nil, &NotSingularError{schedulemember.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *ScheduleMemberQuery) OnlyX(ctx context.Context) *ScheduleMember {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ScheduleMember ID in the query.
// Returns a *NotSingularError when more than one ScheduleMember ID is found.
// Returns a *NotFoundError when no entities are found.
func (smq *ScheduleMemberQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = smq.Limit(2).IDs(setContextOp(ctx, smq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{schedulemember.Label}
	default:
		err = &NotSingularError{schedulemember.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *ScheduleMemberQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScheduleMembers.
func (smq *ScheduleMemberQuery) All(ctx context.Context) ([]*ScheduleMember, error) {
	ctx = setContextOp(ctx, smq.ctx, "All")
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ScheduleMember, *ScheduleMemberQuery]()
	return withInterceptors[[]*ScheduleMember](ctx, smq, qr, smq.inters)
}

// AllX is like All, but panics if an error occurs.
func (smq *ScheduleMemberQuery) AllX(ctx context.Context) []*ScheduleMember {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ScheduleMember IDs.
func (smq *ScheduleMemberQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if smq.ctx.Unique == nil && smq.path != nil {
		smq.Unique(true)
	}
	ctx = setContextOp(ctx, smq.ctx, "IDs")
	if err = smq.Select(schedulemember.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *ScheduleMemberQuery) IDsX(ctx context.Context) []int64 {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *ScheduleMemberQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, smq.ctx, "Count")
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, smq, querierCount[*ScheduleMemberQuery](), smq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (smq *ScheduleMemberQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *ScheduleMemberQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, smq.ctx, "Exist")
	switch _, err := smq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *ScheduleMemberQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ScheduleMemberQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *ScheduleMemberQuery) Clone() *ScheduleMemberQuery {
	if smq == nil {
		return nil
	}
	return &ScheduleMemberQuery{
		config:       smq.config,
		ctx:          smq.ctx.Clone(),
		order:        append([]schedulemember.OrderOption{}, smq.order...),
		inters:       append([]Interceptor{}, smq.inters...),
		predicates:   append([]predicate.ScheduleMember{}, smq.predicates...),
		withSchedule: smq.withSchedule.Clone(),
		// clone intermediate query.
		sql:  smq.sql.Clone(),
		path: smq.path,
	}
}

// WithSchedule tells the query-builder to eager-load the nodes that are connected to
// the "schedule" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *ScheduleMemberQuery) WithSchedule(opts ...func(*ScheduleQuery)) *ScheduleMemberQuery {
	query := (&ScheduleClient{config: smq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	smq.withSchedule = query
	return smq
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
//	client.ScheduleMember.Query().
//		GroupBy(schedulemember.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (smq *ScheduleMemberQuery) GroupBy(field string, fields ...string) *ScheduleMemberGroupBy {
	smq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ScheduleMemberGroupBy{build: smq}
	grbuild.flds = &smq.ctx.Fields
	grbuild.label = schedulemember.Label
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
//	client.ScheduleMember.Query().
//		Select(schedulemember.FieldCreatedAt).
//		Scan(ctx, &v)
func (smq *ScheduleMemberQuery) Select(fields ...string) *ScheduleMemberSelect {
	smq.ctx.Fields = append(smq.ctx.Fields, fields...)
	sbuild := &ScheduleMemberSelect{ScheduleMemberQuery: smq}
	sbuild.label = schedulemember.Label
	sbuild.flds, sbuild.scan = &smq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ScheduleMemberSelect configured with the given aggregations.
func (smq *ScheduleMemberQuery) Aggregate(fns ...AggregateFunc) *ScheduleMemberSelect {
	return smq.Select().Aggregate(fns...)
}

func (smq *ScheduleMemberQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range smq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, smq); err != nil {
				return err
			}
		}
	}
	for _, f := range smq.ctx.Fields {
		if !schedulemember.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *ScheduleMemberQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ScheduleMember, error) {
	var (
		nodes       = []*ScheduleMember{}
		_spec       = smq.querySpec()
		loadedTypes = [1]bool{
			smq.withSchedule != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ScheduleMember).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ScheduleMember{config: smq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = smq.schemaConfig.ScheduleMember
	ctx = internal.NewSchemaConfigContext(ctx, smq.schemaConfig)
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := smq.withSchedule; query != nil {
		if err := smq.loadSchedule(ctx, query, nodes, nil,
			func(n *ScheduleMember, e *Schedule) { n.Edges.Schedule = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (smq *ScheduleMemberQuery) loadSchedule(ctx context.Context, query *ScheduleQuery, nodes []*ScheduleMember, init func(*ScheduleMember), assign func(*ScheduleMember, *Schedule)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*ScheduleMember)
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

func (smq *ScheduleMemberQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	_spec.Node.Schema = smq.schemaConfig.ScheduleMember
	ctx = internal.NewSchemaConfigContext(ctx, smq.schemaConfig)
	_spec.Node.Columns = smq.ctx.Fields
	if len(smq.ctx.Fields) > 0 {
		_spec.Unique = smq.ctx.Unique != nil && *smq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *ScheduleMemberQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(schedulemember.Table, schedulemember.Columns, sqlgraph.NewFieldSpec(schedulemember.FieldID, field.TypeInt64))
	_spec.From = smq.sql
	if unique := smq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if smq.path != nil {
		_spec.Unique = true
	}
	if fields := smq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schedulemember.FieldID)
		for i := range fields {
			if fields[i] != schedulemember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if smq.withSchedule != nil {
			_spec.Node.AddColumnOnce(schedulemember.FieldScheduleID)
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *ScheduleMemberQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(schedulemember.Table)
	columns := smq.ctx.Fields
	if len(columns) == 0 {
		columns = schedulemember.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smq.ctx.Unique != nil && *smq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(smq.schemaConfig.ScheduleMember)
	ctx = internal.NewSchemaConfigContext(ctx, smq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScheduleMemberGroupBy is the group-by builder for ScheduleMember entities.
type ScheduleMemberGroupBy struct {
	selector
	build *ScheduleMemberQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *ScheduleMemberGroupBy) Aggregate(fns ...AggregateFunc) *ScheduleMemberGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the selector query and scans the result into the given value.
func (smgb *ScheduleMemberGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smgb.build.ctx, "GroupBy")
	if err := smgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScheduleMemberQuery, *ScheduleMemberGroupBy](ctx, smgb.build, smgb, smgb.build.inters, v)
}

func (smgb *ScheduleMemberGroupBy) sqlScan(ctx context.Context, root *ScheduleMemberQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*smgb.flds)+len(smgb.fns))
		for _, f := range *smgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*smgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ScheduleMemberSelect is the builder for selecting fields of ScheduleMember entities.
type ScheduleMemberSelect struct {
	*ScheduleMemberQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sms *ScheduleMemberSelect) Aggregate(fns ...AggregateFunc) *ScheduleMemberSelect {
	sms.fns = append(sms.fns, fns...)
	return sms
}

// Scan applies the selector query and scans the result into the given value.
func (sms *ScheduleMemberSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sms.ctx, "Select")
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ScheduleMemberQuery, *ScheduleMemberSelect](ctx, sms.ScheduleMemberQuery, sms, sms.inters, v)
}

func (sms *ScheduleMemberSelect) sqlScan(ctx context.Context, root *ScheduleMemberQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sms.fns))
	for _, fn := range sms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
