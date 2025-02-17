// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/bootcamp"
	"saas/biz/dal/db/ent/bootcampparticipant"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BootcampParticipantQuery is the builder for querying BootcampParticipant entities.
type BootcampParticipantQuery struct {
	config
	ctx           *QueryContext
	order         []bootcampparticipant.OrderOption
	inters        []Interceptor
	predicates    []predicate.BootcampParticipant
	withBootcamps *BootcampQuery
	withMembers   *MemberQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BootcampParticipantQuery builder.
func (bpq *BootcampParticipantQuery) Where(ps ...predicate.BootcampParticipant) *BootcampParticipantQuery {
	bpq.predicates = append(bpq.predicates, ps...)
	return bpq
}

// Limit the number of records to be returned by this query.
func (bpq *BootcampParticipantQuery) Limit(limit int) *BootcampParticipantQuery {
	bpq.ctx.Limit = &limit
	return bpq
}

// Offset to start from.
func (bpq *BootcampParticipantQuery) Offset(offset int) *BootcampParticipantQuery {
	bpq.ctx.Offset = &offset
	return bpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bpq *BootcampParticipantQuery) Unique(unique bool) *BootcampParticipantQuery {
	bpq.ctx.Unique = &unique
	return bpq
}

// Order specifies how the records should be ordered.
func (bpq *BootcampParticipantQuery) Order(o ...bootcampparticipant.OrderOption) *BootcampParticipantQuery {
	bpq.order = append(bpq.order, o...)
	return bpq
}

// QueryBootcamps chains the current query on the "bootcamps" edge.
func (bpq *BootcampParticipantQuery) QueryBootcamps() *BootcampQuery {
	query := (&BootcampClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bootcampparticipant.Table, bootcampparticipant.FieldID, selector),
			sqlgraph.To(bootcamp.Table, bootcamp.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bootcampparticipant.BootcampsTable, bootcampparticipant.BootcampsColumn),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMembers chains the current query on the "members" edge.
func (bpq *BootcampParticipantQuery) QueryMembers() *MemberQuery {
	query := (&MemberClient{config: bpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bootcampparticipant.Table, bootcampparticipant.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, bootcampparticipant.MembersTable, bootcampparticipant.MembersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(bpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BootcampParticipant entity from the query.
// Returns a *NotFoundError when no BootcampParticipant was found.
func (bpq *BootcampParticipantQuery) First(ctx context.Context) (*BootcampParticipant, error) {
	nodes, err := bpq.Limit(1).All(setContextOp(ctx, bpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bootcampparticipant.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) FirstX(ctx context.Context) *BootcampParticipant {
	node, err := bpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BootcampParticipant ID from the query.
// Returns a *NotFoundError when no BootcampParticipant ID was found.
func (bpq *BootcampParticipantQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bpq.Limit(1).IDs(setContextOp(ctx, bpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bootcampparticipant.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) FirstIDX(ctx context.Context) int64 {
	id, err := bpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BootcampParticipant entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BootcampParticipant entity is found.
// Returns a *NotFoundError when no BootcampParticipant entities are found.
func (bpq *BootcampParticipantQuery) Only(ctx context.Context) (*BootcampParticipant, error) {
	nodes, err := bpq.Limit(2).All(setContextOp(ctx, bpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bootcampparticipant.Label}
	default:
		return nil, &NotSingularError{bootcampparticipant.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) OnlyX(ctx context.Context) *BootcampParticipant {
	node, err := bpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BootcampParticipant ID in the query.
// Returns a *NotSingularError when more than one BootcampParticipant ID is found.
// Returns a *NotFoundError when no entities are found.
func (bpq *BootcampParticipantQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = bpq.Limit(2).IDs(setContextOp(ctx, bpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bootcampparticipant.Label}
	default:
		err = &NotSingularError{bootcampparticipant.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := bpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BootcampParticipants.
func (bpq *BootcampParticipantQuery) All(ctx context.Context) ([]*BootcampParticipant, error) {
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryAll)
	if err := bpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BootcampParticipant, *BootcampParticipantQuery]()
	return withInterceptors[[]*BootcampParticipant](ctx, bpq, qr, bpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) AllX(ctx context.Context) []*BootcampParticipant {
	nodes, err := bpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BootcampParticipant IDs.
func (bpq *BootcampParticipantQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if bpq.ctx.Unique == nil && bpq.path != nil {
		bpq.Unique(true)
	}
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryIDs)
	if err = bpq.Select(bootcampparticipant.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) IDsX(ctx context.Context) []int64 {
	ids, err := bpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bpq *BootcampParticipantQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryCount)
	if err := bpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bpq, querierCount[*BootcampParticipantQuery](), bpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) CountX(ctx context.Context) int {
	count, err := bpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bpq *BootcampParticipantQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bpq.ctx, ent.OpQueryExist)
	switch _, err := bpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bpq *BootcampParticipantQuery) ExistX(ctx context.Context) bool {
	exist, err := bpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BootcampParticipantQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bpq *BootcampParticipantQuery) Clone() *BootcampParticipantQuery {
	if bpq == nil {
		return nil
	}
	return &BootcampParticipantQuery{
		config:        bpq.config,
		ctx:           bpq.ctx.Clone(),
		order:         append([]bootcampparticipant.OrderOption{}, bpq.order...),
		inters:        append([]Interceptor{}, bpq.inters...),
		predicates:    append([]predicate.BootcampParticipant{}, bpq.predicates...),
		withBootcamps: bpq.withBootcamps.Clone(),
		withMembers:   bpq.withMembers.Clone(),
		// clone intermediate query.
		sql:  bpq.sql.Clone(),
		path: bpq.path,
	}
}

// WithBootcamps tells the query-builder to eager-load the nodes that are connected to
// the "bootcamps" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BootcampParticipantQuery) WithBootcamps(opts ...func(*BootcampQuery)) *BootcampParticipantQuery {
	query := (&BootcampClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withBootcamps = query
	return bpq
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (bpq *BootcampParticipantQuery) WithMembers(opts ...func(*MemberQuery)) *BootcampParticipantQuery {
	query := (&MemberClient{config: bpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bpq.withMembers = query
	return bpq
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
//	client.BootcampParticipant.Query().
//		GroupBy(bootcampparticipant.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bpq *BootcampParticipantQuery) GroupBy(field string, fields ...string) *BootcampParticipantGroupBy {
	bpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BootcampParticipantGroupBy{build: bpq}
	grbuild.flds = &bpq.ctx.Fields
	grbuild.label = bootcampparticipant.Label
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
//	client.BootcampParticipant.Query().
//		Select(bootcampparticipant.FieldCreatedAt).
//		Scan(ctx, &v)
func (bpq *BootcampParticipantQuery) Select(fields ...string) *BootcampParticipantSelect {
	bpq.ctx.Fields = append(bpq.ctx.Fields, fields...)
	sbuild := &BootcampParticipantSelect{BootcampParticipantQuery: bpq}
	sbuild.label = bootcampparticipant.Label
	sbuild.flds, sbuild.scan = &bpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BootcampParticipantSelect configured with the given aggregations.
func (bpq *BootcampParticipantQuery) Aggregate(fns ...AggregateFunc) *BootcampParticipantSelect {
	return bpq.Select().Aggregate(fns...)
}

func (bpq *BootcampParticipantQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bpq); err != nil {
				return err
			}
		}
	}
	for _, f := range bpq.ctx.Fields {
		if !bootcampparticipant.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bpq.path != nil {
		prev, err := bpq.path(ctx)
		if err != nil {
			return err
		}
		bpq.sql = prev
	}
	return nil
}

func (bpq *BootcampParticipantQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BootcampParticipant, error) {
	var (
		nodes       = []*BootcampParticipant{}
		_spec       = bpq.querySpec()
		loadedTypes = [2]bool{
			bpq.withBootcamps != nil,
			bpq.withMembers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BootcampParticipant).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BootcampParticipant{config: bpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bpq.withBootcamps; query != nil {
		if err := bpq.loadBootcamps(ctx, query, nodes, nil,
			func(n *BootcampParticipant, e *Bootcamp) { n.Edges.Bootcamps = e }); err != nil {
			return nil, err
		}
	}
	if query := bpq.withMembers; query != nil {
		if err := bpq.loadMembers(ctx, query, nodes,
			func(n *BootcampParticipant) { n.Edges.Members = []*Member{} },
			func(n *BootcampParticipant, e *Member) { n.Edges.Members = append(n.Edges.Members, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bpq *BootcampParticipantQuery) loadBootcamps(ctx context.Context, query *BootcampQuery, nodes []*BootcampParticipant, init func(*BootcampParticipant), assign func(*BootcampParticipant, *Bootcamp)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*BootcampParticipant)
	for i := range nodes {
		fk := nodes[i].BootcampID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(bootcamp.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "bootcamp_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (bpq *BootcampParticipantQuery) loadMembers(ctx context.Context, query *MemberQuery, nodes []*BootcampParticipant, init func(*BootcampParticipant), assign func(*BootcampParticipant, *Member)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*BootcampParticipant)
	nids := make(map[int64]map[*BootcampParticipant]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(bootcampparticipant.MembersTable)
		s.Join(joinT).On(s.C(member.FieldID), joinT.C(bootcampparticipant.MembersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(bootcampparticipant.MembersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(bootcampparticipant.MembersPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullInt64).Int64
				inValue := values[1].(*sql.NullInt64).Int64
				if nids[inValue] == nil {
					nids[inValue] = map[*BootcampParticipant]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Member](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "members" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (bpq *BootcampParticipantQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bpq.querySpec()
	_spec.Node.Columns = bpq.ctx.Fields
	if len(bpq.ctx.Fields) > 0 {
		_spec.Unique = bpq.ctx.Unique != nil && *bpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bpq.driver, _spec)
}

func (bpq *BootcampParticipantQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bootcampparticipant.Table, bootcampparticipant.Columns, sqlgraph.NewFieldSpec(bootcampparticipant.FieldID, field.TypeInt64))
	_spec.From = bpq.sql
	if unique := bpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bpq.path != nil {
		_spec.Unique = true
	}
	if fields := bpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bootcampparticipant.FieldID)
		for i := range fields {
			if fields[i] != bootcampparticipant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if bpq.withBootcamps != nil {
			_spec.Node.AddColumnOnce(bootcampparticipant.FieldBootcampID)
		}
	}
	if ps := bpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bpq *BootcampParticipantQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bpq.driver.Dialect())
	t1 := builder.Table(bootcampparticipant.Table)
	columns := bpq.ctx.Fields
	if len(columns) == 0 {
		columns = bootcampparticipant.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bpq.sql != nil {
		selector = bpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bpq.ctx.Unique != nil && *bpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bpq.predicates {
		p(selector)
	}
	for _, p := range bpq.order {
		p(selector)
	}
	if offset := bpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BootcampParticipantGroupBy is the group-by builder for BootcampParticipant entities.
type BootcampParticipantGroupBy struct {
	selector
	build *BootcampParticipantQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bpgb *BootcampParticipantGroupBy) Aggregate(fns ...AggregateFunc) *BootcampParticipantGroupBy {
	bpgb.fns = append(bpgb.fns, fns...)
	return bpgb
}

// Scan applies the selector query and scans the result into the given value.
func (bpgb *BootcampParticipantGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bpgb.build.ctx, ent.OpQueryGroupBy)
	if err := bpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BootcampParticipantQuery, *BootcampParticipantGroupBy](ctx, bpgb.build, bpgb, bpgb.build.inters, v)
}

func (bpgb *BootcampParticipantGroupBy) sqlScan(ctx context.Context, root *BootcampParticipantQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bpgb.fns))
	for _, fn := range bpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bpgb.flds)+len(bpgb.fns))
		for _, f := range *bpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BootcampParticipantSelect is the builder for selecting fields of BootcampParticipant entities.
type BootcampParticipantSelect struct {
	*BootcampParticipantQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bps *BootcampParticipantSelect) Aggregate(fns ...AggregateFunc) *BootcampParticipantSelect {
	bps.fns = append(bps.fns, fns...)
	return bps
}

// Scan applies the selector query and scans the result into the given value.
func (bps *BootcampParticipantSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bps.ctx, ent.OpQuerySelect)
	if err := bps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BootcampParticipantQuery, *BootcampParticipantSelect](ctx, bps.BootcampParticipantQuery, bps, bps.inters, v)
}

func (bps *BootcampParticipantSelect) sqlScan(ctx context.Context, root *BootcampParticipantQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bps.fns))
	for _, fn := range bps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
