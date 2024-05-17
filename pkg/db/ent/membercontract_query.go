// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/membercontract"
	"saas/pkg/db/ent/membercontractcontent"
	"saas/pkg/db/ent/memberproduct"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberContractQuery is the builder for querying MemberContract entities.
type MemberContractQuery struct {
	config
	ctx               *QueryContext
	order             []membercontract.OrderOption
	inters            []Interceptor
	predicates        []predicate.MemberContract
	withContent       *MemberContractContentQuery
	withMember        *MemberQuery
	withOrder         *OrderQuery
	withMemberProduct *MemberProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberContractQuery builder.
func (mcq *MemberContractQuery) Where(ps ...predicate.MemberContract) *MemberContractQuery {
	mcq.predicates = append(mcq.predicates, ps...)
	return mcq
}

// Limit the number of records to be returned by this query.
func (mcq *MemberContractQuery) Limit(limit int) *MemberContractQuery {
	mcq.ctx.Limit = &limit
	return mcq
}

// Offset to start from.
func (mcq *MemberContractQuery) Offset(offset int) *MemberContractQuery {
	mcq.ctx.Offset = &offset
	return mcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mcq *MemberContractQuery) Unique(unique bool) *MemberContractQuery {
	mcq.ctx.Unique = &unique
	return mcq
}

// Order specifies how the records should be ordered.
func (mcq *MemberContractQuery) Order(o ...membercontract.OrderOption) *MemberContractQuery {
	mcq.order = append(mcq.order, o...)
	return mcq
}

// QueryContent chains the current query on the "content" edge.
func (mcq *MemberContractQuery) QueryContent() *MemberContractContentQuery {
	query := (&MemberContractContentClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membercontract.Table, membercontract.FieldID, selector),
			sqlgraph.To(membercontractcontent.Table, membercontractcontent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, membercontract.ContentTable, membercontract.ContentColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMember chains the current query on the "member" edge.
func (mcq *MemberContractQuery) QueryMember() *MemberQuery {
	query := (&MemberClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membercontract.Table, membercontract.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membercontract.MemberTable, membercontract.MemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrder chains the current query on the "order" edge.
func (mcq *MemberContractQuery) QueryOrder() *OrderQuery {
	query := (&OrderClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membercontract.Table, membercontract.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membercontract.OrderTable, membercontract.OrderColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberProduct chains the current query on the "member_product" edge.
func (mcq *MemberContractQuery) QueryMemberProduct() *MemberProductQuery {
	query := (&MemberProductClient{config: mcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(membercontract.Table, membercontract.FieldID, selector),
			sqlgraph.To(memberproduct.Table, memberproduct.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, membercontract.MemberProductTable, membercontract.MemberProductColumn),
		)
		fromU = sqlgraph.SetNeighbors(mcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MemberContract entity from the query.
// Returns a *NotFoundError when no MemberContract was found.
func (mcq *MemberContractQuery) First(ctx context.Context) (*MemberContract, error) {
	nodes, err := mcq.Limit(1).All(setContextOp(ctx, mcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{membercontract.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mcq *MemberContractQuery) FirstX(ctx context.Context) *MemberContract {
	node, err := mcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MemberContract ID from the query.
// Returns a *NotFoundError when no MemberContract ID was found.
func (mcq *MemberContractQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mcq.Limit(1).IDs(setContextOp(ctx, mcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{membercontract.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mcq *MemberContractQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MemberContract entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MemberContract entity is found.
// Returns a *NotFoundError when no MemberContract entities are found.
func (mcq *MemberContractQuery) Only(ctx context.Context) (*MemberContract, error) {
	nodes, err := mcq.Limit(2).All(setContextOp(ctx, mcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{membercontract.Label}
	default:
		return nil, &NotSingularError{membercontract.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mcq *MemberContractQuery) OnlyX(ctx context.Context) *MemberContract {
	node, err := mcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MemberContract ID in the query.
// Returns a *NotSingularError when more than one MemberContract ID is found.
// Returns a *NotFoundError when no entities are found.
func (mcq *MemberContractQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mcq.Limit(2).IDs(setContextOp(ctx, mcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{membercontract.Label}
	default:
		err = &NotSingularError{membercontract.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mcq *MemberContractQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MemberContracts.
func (mcq *MemberContractQuery) All(ctx context.Context) ([]*MemberContract, error) {
	ctx = setContextOp(ctx, mcq.ctx, "All")
	if err := mcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MemberContract, *MemberContractQuery]()
	return withInterceptors[[]*MemberContract](ctx, mcq, qr, mcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mcq *MemberContractQuery) AllX(ctx context.Context) []*MemberContract {
	nodes, err := mcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MemberContract IDs.
func (mcq *MemberContractQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mcq.ctx.Unique == nil && mcq.path != nil {
		mcq.Unique(true)
	}
	ctx = setContextOp(ctx, mcq.ctx, "IDs")
	if err = mcq.Select(membercontract.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mcq *MemberContractQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mcq *MemberContractQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mcq.ctx, "Count")
	if err := mcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mcq, querierCount[*MemberContractQuery](), mcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mcq *MemberContractQuery) CountX(ctx context.Context) int {
	count, err := mcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mcq *MemberContractQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mcq.ctx, "Exist")
	switch _, err := mcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mcq *MemberContractQuery) ExistX(ctx context.Context) bool {
	exist, err := mcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberContractQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mcq *MemberContractQuery) Clone() *MemberContractQuery {
	if mcq == nil {
		return nil
	}
	return &MemberContractQuery{
		config:            mcq.config,
		ctx:               mcq.ctx.Clone(),
		order:             append([]membercontract.OrderOption{}, mcq.order...),
		inters:            append([]Interceptor{}, mcq.inters...),
		predicates:        append([]predicate.MemberContract{}, mcq.predicates...),
		withContent:       mcq.withContent.Clone(),
		withMember:        mcq.withMember.Clone(),
		withOrder:         mcq.withOrder.Clone(),
		withMemberProduct: mcq.withMemberProduct.Clone(),
		// clone intermediate query.
		sql:  mcq.sql.Clone(),
		path: mcq.path,
	}
}

// WithContent tells the query-builder to eager-load the nodes that are connected to
// the "content" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MemberContractQuery) WithContent(opts ...func(*MemberContractContentQuery)) *MemberContractQuery {
	query := (&MemberContractContentClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withContent = query
	return mcq
}

// WithMember tells the query-builder to eager-load the nodes that are connected to
// the "member" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MemberContractQuery) WithMember(opts ...func(*MemberQuery)) *MemberContractQuery {
	query := (&MemberClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withMember = query
	return mcq
}

// WithOrder tells the query-builder to eager-load the nodes that are connected to
// the "order" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MemberContractQuery) WithOrder(opts ...func(*OrderQuery)) *MemberContractQuery {
	query := (&OrderClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withOrder = query
	return mcq
}

// WithMemberProduct tells the query-builder to eager-load the nodes that are connected to
// the "member_product" edge. The optional arguments are used to configure the query builder of the edge.
func (mcq *MemberContractQuery) WithMemberProduct(opts ...func(*MemberProductQuery)) *MemberContractQuery {
	query := (&MemberProductClient{config: mcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mcq.withMemberProduct = query
	return mcq
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
//	client.MemberContract.Query().
//		GroupBy(membercontract.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mcq *MemberContractQuery) GroupBy(field string, fields ...string) *MemberContractGroupBy {
	mcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberContractGroupBy{build: mcq}
	grbuild.flds = &mcq.ctx.Fields
	grbuild.label = membercontract.Label
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
//	client.MemberContract.Query().
//		Select(membercontract.FieldCreatedAt).
//		Scan(ctx, &v)
func (mcq *MemberContractQuery) Select(fields ...string) *MemberContractSelect {
	mcq.ctx.Fields = append(mcq.ctx.Fields, fields...)
	sbuild := &MemberContractSelect{MemberContractQuery: mcq}
	sbuild.label = membercontract.Label
	sbuild.flds, sbuild.scan = &mcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberContractSelect configured with the given aggregations.
func (mcq *MemberContractQuery) Aggregate(fns ...AggregateFunc) *MemberContractSelect {
	return mcq.Select().Aggregate(fns...)
}

func (mcq *MemberContractQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mcq); err != nil {
				return err
			}
		}
	}
	for _, f := range mcq.ctx.Fields {
		if !membercontract.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mcq.path != nil {
		prev, err := mcq.path(ctx)
		if err != nil {
			return err
		}
		mcq.sql = prev
	}
	return nil
}

func (mcq *MemberContractQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MemberContract, error) {
	var (
		nodes       = []*MemberContract{}
		_spec       = mcq.querySpec()
		loadedTypes = [4]bool{
			mcq.withContent != nil,
			mcq.withMember != nil,
			mcq.withOrder != nil,
			mcq.withMemberProduct != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MemberContract).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MemberContract{config: mcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mcq.withContent; query != nil {
		if err := mcq.loadContent(ctx, query, nodes,
			func(n *MemberContract) { n.Edges.Content = []*MemberContractContent{} },
			func(n *MemberContract, e *MemberContractContent) { n.Edges.Content = append(n.Edges.Content, e) }); err != nil {
			return nil, err
		}
	}
	if query := mcq.withMember; query != nil {
		if err := mcq.loadMember(ctx, query, nodes, nil,
			func(n *MemberContract, e *Member) { n.Edges.Member = e }); err != nil {
			return nil, err
		}
	}
	if query := mcq.withOrder; query != nil {
		if err := mcq.loadOrder(ctx, query, nodes, nil,
			func(n *MemberContract, e *Order) { n.Edges.Order = e }); err != nil {
			return nil, err
		}
	}
	if query := mcq.withMemberProduct; query != nil {
		if err := mcq.loadMemberProduct(ctx, query, nodes, nil,
			func(n *MemberContract, e *MemberProduct) { n.Edges.MemberProduct = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mcq *MemberContractQuery) loadContent(ctx context.Context, query *MemberContractContentQuery, nodes []*MemberContract, init func(*MemberContract), assign func(*MemberContract, *MemberContractContent)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MemberContract)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(membercontractcontent.FieldMemberContractID)
	}
	query.Where(predicate.MemberContractContent(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(membercontract.ContentColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberContractID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_contract_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mcq *MemberContractQuery) loadMember(ctx context.Context, query *MemberQuery, nodes []*MemberContract, init func(*MemberContract), assign func(*MemberContract, *Member)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MemberContract)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (mcq *MemberContractQuery) loadOrder(ctx context.Context, query *OrderQuery, nodes []*MemberContract, init func(*MemberContract), assign func(*MemberContract, *Order)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MemberContract)
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
func (mcq *MemberContractQuery) loadMemberProduct(ctx context.Context, query *MemberProductQuery, nodes []*MemberContract, init func(*MemberContract), assign func(*MemberContract, *MemberProduct)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MemberContract)
	for i := range nodes {
		fk := nodes[i].MemberProductID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(memberproduct.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_product_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (mcq *MemberContractQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mcq.querySpec()
	_spec.Node.Columns = mcq.ctx.Fields
	if len(mcq.ctx.Fields) > 0 {
		_spec.Unique = mcq.ctx.Unique != nil && *mcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mcq.driver, _spec)
}

func (mcq *MemberContractQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(membercontract.Table, membercontract.Columns, sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64))
	_spec.From = mcq.sql
	if unique := mcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mcq.path != nil {
		_spec.Unique = true
	}
	if fields := mcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, membercontract.FieldID)
		for i := range fields {
			if fields[i] != membercontract.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mcq.withMember != nil {
			_spec.Node.AddColumnOnce(membercontract.FieldMemberID)
		}
		if mcq.withOrder != nil {
			_spec.Node.AddColumnOnce(membercontract.FieldOrderID)
		}
		if mcq.withMemberProduct != nil {
			_spec.Node.AddColumnOnce(membercontract.FieldMemberProductID)
		}
	}
	if ps := mcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mcq *MemberContractQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mcq.driver.Dialect())
	t1 := builder.Table(membercontract.Table)
	columns := mcq.ctx.Fields
	if len(columns) == 0 {
		columns = membercontract.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mcq.sql != nil {
		selector = mcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mcq.ctx.Unique != nil && *mcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mcq.predicates {
		p(selector)
	}
	for _, p := range mcq.order {
		p(selector)
	}
	if offset := mcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MemberContractGroupBy is the group-by builder for MemberContract entities.
type MemberContractGroupBy struct {
	selector
	build *MemberContractQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mcgb *MemberContractGroupBy) Aggregate(fns ...AggregateFunc) *MemberContractGroupBy {
	mcgb.fns = append(mcgb.fns, fns...)
	return mcgb
}

// Scan applies the selector query and scans the result into the given value.
func (mcgb *MemberContractGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcgb.build.ctx, "GroupBy")
	if err := mcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberContractQuery, *MemberContractGroupBy](ctx, mcgb.build, mcgb, mcgb.build.inters, v)
}

func (mcgb *MemberContractGroupBy) sqlScan(ctx context.Context, root *MemberContractQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mcgb.fns))
	for _, fn := range mcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mcgb.flds)+len(mcgb.fns))
		for _, f := range *mcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberContractSelect is the builder for selecting fields of MemberContract entities.
type MemberContractSelect struct {
	*MemberContractQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mcs *MemberContractSelect) Aggregate(fns ...AggregateFunc) *MemberContractSelect {
	mcs.fns = append(mcs.fns, fns...)
	return mcs
}

// Scan applies the selector query and scans the result into the given value.
func (mcs *MemberContractSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mcs.ctx, "Select")
	if err := mcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberContractQuery, *MemberContractSelect](ctx, mcs.MemberContractQuery, mcs, mcs.inters, v)
}

func (mcs *MemberContractSelect) sqlScan(ctx context.Context, root *MemberContractQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mcs.fns))
	for _, fn := range mcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
