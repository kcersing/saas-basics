// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/pkg/db/ent/membercontract"
	"saas/pkg/db/ent/membercontractcontent"
	"saas/pkg/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberContractContentQuery is the builder for querying MemberContractContent entities.
type MemberContractContentQuery struct {
	config
	ctx          *QueryContext
	order        []membercontractcontent.OrderOption
	inters       []Interceptor
	predicates   []predicate.MemberContractContent
	withContract *MemberContractQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberContractContentQuery builder.
func (mccq *MemberContractContentQuery) Where(ps ...predicate.MemberContractContent) *MemberContractContentQuery {
	mccq.predicates = append(mccq.predicates, ps...)
	return mccq
}

// Limit the number of records to be returned by this query.
func (mccq *MemberContractContentQuery) Limit(limit int) *MemberContractContentQuery {
	mccq.ctx.Limit = &limit
	return mccq
}

// Offset to start from.
func (mccq *MemberContractContentQuery) Offset(offset int) *MemberContractContentQuery {
	mccq.ctx.Offset = &offset
	return mccq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mccq *MemberContractContentQuery) Unique(unique bool) *MemberContractContentQuery {
	mccq.ctx.Unique = &unique
	return mccq
}

// Order specifies how the records should be ordered.
func (mccq *MemberContractContentQuery) Order(o ...membercontractcontent.OrderOption) *MemberContractContentQuery {
	mccq.order = append(mccq.order, o...)
	return mccq
}

// QueryContract chains the current query on the "contract" edge.
func (mccq *MemberContractContentQuery) QueryContract() *MemberContractQuery {
	query := (&MemberContractClient{config: mccq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mccq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mccq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membercontractcontent.Table, membercontractcontent.FieldID, selector),
			sqlgraph.To(membercontract.Table, membercontract.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membercontractcontent.ContractTable, membercontractcontent.ContractColumn),
		)
		fromU = sqlgraph.SetNeighbors(mccq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MemberContractContent entity from the query.
// Returns a *NotFoundError when no MemberContractContent was found.
func (mccq *MemberContractContentQuery) First(ctx context.Context) (*MemberContractContent, error) {
	nodes, err := mccq.Limit(1).All(setContextOp(ctx, mccq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{membercontractcontent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mccq *MemberContractContentQuery) FirstX(ctx context.Context) *MemberContractContent {
	node, err := mccq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MemberContractContent ID from the query.
// Returns a *NotFoundError when no MemberContractContent ID was found.
func (mccq *MemberContractContentQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mccq.Limit(1).IDs(setContextOp(ctx, mccq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{membercontractcontent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mccq *MemberContractContentQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mccq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MemberContractContent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MemberContractContent entity is found.
// Returns a *NotFoundError when no MemberContractContent entities are found.
func (mccq *MemberContractContentQuery) Only(ctx context.Context) (*MemberContractContent, error) {
	nodes, err := mccq.Limit(2).All(setContextOp(ctx, mccq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{membercontractcontent.Label}
	default:
		return nil, &NotSingularError{membercontractcontent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mccq *MemberContractContentQuery) OnlyX(ctx context.Context) *MemberContractContent {
	node, err := mccq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MemberContractContent ID in the query.
// Returns a *NotSingularError when more than one MemberContractContent ID is found.
// Returns a *NotFoundError when no entities are found.
func (mccq *MemberContractContentQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mccq.Limit(2).IDs(setContextOp(ctx, mccq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{membercontractcontent.Label}
	default:
		err = &NotSingularError{membercontractcontent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mccq *MemberContractContentQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mccq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MemberContractContents.
func (mccq *MemberContractContentQuery) All(ctx context.Context) ([]*MemberContractContent, error) {
	ctx = setContextOp(ctx, mccq.ctx, "All")
	if err := mccq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MemberContractContent, *MemberContractContentQuery]()
	return withInterceptors[[]*MemberContractContent](ctx, mccq, qr, mccq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mccq *MemberContractContentQuery) AllX(ctx context.Context) []*MemberContractContent {
	nodes, err := mccq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MemberContractContent IDs.
func (mccq *MemberContractContentQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mccq.ctx.Unique == nil && mccq.path != nil {
		mccq.Unique(true)
	}
	ctx = setContextOp(ctx, mccq.ctx, "IDs")
	if err = mccq.Select(membercontractcontent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mccq *MemberContractContentQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mccq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mccq *MemberContractContentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mccq.ctx, "Count")
	if err := mccq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mccq, querierCount[*MemberContractContentQuery](), mccq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mccq *MemberContractContentQuery) CountX(ctx context.Context) int {
	count, err := mccq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mccq *MemberContractContentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mccq.ctx, "Exist")
	switch _, err := mccq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mccq *MemberContractContentQuery) ExistX(ctx context.Context) bool {
	exist, err := mccq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberContractContentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mccq *MemberContractContentQuery) Clone() *MemberContractContentQuery {
	if mccq == nil {
		return nil
	}
	return &MemberContractContentQuery{
		config:       mccq.config,
		ctx:          mccq.ctx.Clone(),
		order:        append([]membercontractcontent.OrderOption{}, mccq.order...),
		inters:       append([]Interceptor{}, mccq.inters...),
		predicates:   append([]predicate.MemberContractContent{}, mccq.predicates...),
		withContract: mccq.withContract.Clone(),
		// clone intermediate query.
		sql:  mccq.sql.Clone(),
		path: mccq.path,
	}
}

// WithContract tells the query-builder to eager-load the nodes that are connected to
// the "contract" edge. The optional arguments are used to configure the query builder of the edge.
func (mccq *MemberContractContentQuery) WithContract(opts ...func(*MemberContractQuery)) *MemberContractContentQuery {
	query := (&MemberContractClient{config: mccq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mccq.withContract = query
	return mccq
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
//	client.MemberContractContent.Query().
//		GroupBy(membercontractcontent.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mccq *MemberContractContentQuery) GroupBy(field string, fields ...string) *MemberContractContentGroupBy {
	mccq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberContractContentGroupBy{build: mccq}
	grbuild.flds = &mccq.ctx.Fields
	grbuild.label = membercontractcontent.Label
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
//	client.MemberContractContent.Query().
//		Select(membercontractcontent.FieldCreatedAt).
//		Scan(ctx, &v)
func (mccq *MemberContractContentQuery) Select(fields ...string) *MemberContractContentSelect {
	mccq.ctx.Fields = append(mccq.ctx.Fields, fields...)
	sbuild := &MemberContractContentSelect{MemberContractContentQuery: mccq}
	sbuild.label = membercontractcontent.Label
	sbuild.flds, sbuild.scan = &mccq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberContractContentSelect configured with the given aggregations.
func (mccq *MemberContractContentQuery) Aggregate(fns ...AggregateFunc) *MemberContractContentSelect {
	return mccq.Select().Aggregate(fns...)
}

func (mccq *MemberContractContentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mccq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mccq); err != nil {
				return err
			}
		}
	}
	for _, f := range mccq.ctx.Fields {
		if !membercontractcontent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mccq.path != nil {
		prev, err := mccq.path(ctx)
		if err != nil {
			return err
		}
		mccq.sql = prev
	}
	return nil
}

func (mccq *MemberContractContentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MemberContractContent, error) {
	var (
		nodes       = []*MemberContractContent{}
		_spec       = mccq.querySpec()
		loadedTypes = [1]bool{
			mccq.withContract != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MemberContractContent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MemberContractContent{config: mccq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mccq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mccq.withContract; query != nil {
		if err := mccq.loadContract(ctx, query, nodes, nil,
			func(n *MemberContractContent, e *MemberContract) { n.Edges.Contract = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mccq *MemberContractContentQuery) loadContract(ctx context.Context, query *MemberContractQuery, nodes []*MemberContractContent, init func(*MemberContractContent), assign func(*MemberContractContent, *MemberContract)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MemberContractContent)
	for i := range nodes {
		fk := nodes[i].MemberContractID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(membercontract.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_contract_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mccq *MemberContractContentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mccq.querySpec()
	_spec.Node.Columns = mccq.ctx.Fields
	if len(mccq.ctx.Fields) > 0 {
		_spec.Unique = mccq.ctx.Unique != nil && *mccq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mccq.driver, _spec)
}

func (mccq *MemberContractContentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(membercontractcontent.Table, membercontractcontent.Columns, sqlgraph.NewFieldSpec(membercontractcontent.FieldID, field.TypeInt64))
	_spec.From = mccq.sql
	if unique := mccq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mccq.path != nil {
		_spec.Unique = true
	}
	if fields := mccq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membercontractcontent.FieldID)
		for i := range fields {
			if fields[i] != membercontractcontent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mccq.withContract != nil {
			_spec.Node.AddColumnOnce(membercontractcontent.FieldMemberContractID)
		}
	}
	if ps := mccq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mccq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mccq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mccq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mccq *MemberContractContentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mccq.driver.Dialect())
	t1 := builder.Table(membercontractcontent.Table)
	columns := mccq.ctx.Fields
	if len(columns) == 0 {
		columns = membercontractcontent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mccq.sql != nil {
		selector = mccq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mccq.ctx.Unique != nil && *mccq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mccq.predicates {
		p(selector)
	}
	for _, p := range mccq.order {
		p(selector)
	}
	if offset := mccq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mccq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MemberContractContentGroupBy is the group-by builder for MemberContractContent entities.
type MemberContractContentGroupBy struct {
	selector
	build *MemberContractContentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mccgb *MemberContractContentGroupBy) Aggregate(fns ...AggregateFunc) *MemberContractContentGroupBy {
	mccgb.fns = append(mccgb.fns, fns...)
	return mccgb
}

// Scan applies the selector query and scans the result into the given value.
func (mccgb *MemberContractContentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mccgb.build.ctx, "GroupBy")
	if err := mccgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberContractContentQuery, *MemberContractContentGroupBy](ctx, mccgb.build, mccgb, mccgb.build.inters, v)
}

func (mccgb *MemberContractContentGroupBy) sqlScan(ctx context.Context, root *MemberContractContentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mccgb.fns))
	for _, fn := range mccgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mccgb.flds)+len(mccgb.fns))
		for _, f := range *mccgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mccgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mccgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberContractContentSelect is the builder for selecting fields of MemberContractContent entities.
type MemberContractContentSelect struct {
	*MemberContractContentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mccs *MemberContractContentSelect) Aggregate(fns ...AggregateFunc) *MemberContractContentSelect {
	mccs.fns = append(mccs.fns, fns...)
	return mccs
}

// Scan applies the selector query and scans the result into the given value.
func (mccs *MemberContractContentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mccs.ctx, "Select")
	if err := mccs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberContractContentQuery, *MemberContractContentSelect](ctx, mccs.MemberContractContentQuery, mccs, mccs.inters, v)
}

func (mccs *MemberContractContentSelect) sqlScan(ctx context.Context, root *MemberContractContentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mccs.fns))
	for _, fn := range mccs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mccs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mccs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
