// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/membercontract"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberproductcourses"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProductQuery is the builder for querying MemberProduct entities.
type MemberProductQuery struct {
	config
	ctx                       *QueryContext
	order                     []memberproduct.OrderOption
	inters                    []Interceptor
	predicates                []predicate.MemberProduct
	withMembers               *MemberQuery
	withMemberProductEntry    *EntryLogsQuery
	withMemberProductContents *MemberContractQuery
	withMemberCourses         *MemberProductCoursesQuery
	withMemberLessons         *MemberProductCoursesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MemberProductQuery builder.
func (mpq *MemberProductQuery) Where(ps ...predicate.MemberProduct) *MemberProductQuery {
	mpq.predicates = append(mpq.predicates, ps...)
	return mpq
}

// Limit the number of records to be returned by this query.
func (mpq *MemberProductQuery) Limit(limit int) *MemberProductQuery {
	mpq.ctx.Limit = &limit
	return mpq
}

// Offset to start from.
func (mpq *MemberProductQuery) Offset(offset int) *MemberProductQuery {
	mpq.ctx.Offset = &offset
	return mpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mpq *MemberProductQuery) Unique(unique bool) *MemberProductQuery {
	mpq.ctx.Unique = &unique
	return mpq
}

// Order specifies how the records should be ordered.
func (mpq *MemberProductQuery) Order(o ...memberproduct.OrderOption) *MemberProductQuery {
	mpq.order = append(mpq.order, o...)
	return mpq
}

// QueryMembers chains the current query on the "members" edge.
func (mpq *MemberProductQuery) QueryMembers() *MemberQuery {
	query := (&MemberClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(memberproduct.Table, memberproduct.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, memberproduct.MembersTable, memberproduct.MembersColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberProductEntry chains the current query on the "member_product_entry" edge.
func (mpq *MemberProductQuery) QueryMemberProductEntry() *EntryLogsQuery {
	query := (&EntryLogsClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(memberproduct.Table, memberproduct.FieldID, selector),
			sqlgraph.To(entrylogs.Table, entrylogs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, memberproduct.MemberProductEntryTable, memberproduct.MemberProductEntryColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberProductContents chains the current query on the "member_product_contents" edge.
func (mpq *MemberProductQuery) QueryMemberProductContents() *MemberContractQuery {
	query := (&MemberContractClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(memberproduct.Table, memberproduct.FieldID, selector),
			sqlgraph.To(membercontract.Table, membercontract.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, memberproduct.MemberProductContentsTable, memberproduct.MemberProductContentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberCourses chains the current query on the "memberCourses" edge.
func (mpq *MemberProductQuery) QueryMemberCourses() *MemberProductCoursesQuery {
	query := (&MemberProductCoursesClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(memberproduct.Table, memberproduct.FieldID, selector),
			sqlgraph.To(memberproductcourses.Table, memberproductcourses.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, memberproduct.MemberCoursesTable, memberproduct.MemberCoursesColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberLessons chains the current query on the "memberLessons" edge.
func (mpq *MemberProductQuery) QueryMemberLessons() *MemberProductCoursesQuery {
	query := (&MemberProductCoursesClient{config: mpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(memberproduct.Table, memberproduct.FieldID, selector),
			sqlgraph.To(memberproductcourses.Table, memberproductcourses.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, memberproduct.MemberLessonsTable, memberproduct.MemberLessonsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first MemberProduct entity from the query.
// Returns a *NotFoundError when no MemberProduct was found.
func (mpq *MemberProductQuery) First(ctx context.Context) (*MemberProduct, error) {
	nodes, err := mpq.Limit(1).All(setContextOp(ctx, mpq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{memberproduct.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mpq *MemberProductQuery) FirstX(ctx context.Context) *MemberProduct {
	node, err := mpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MemberProduct ID from the query.
// Returns a *NotFoundError when no MemberProduct ID was found.
func (mpq *MemberProductQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(1).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{memberproduct.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mpq *MemberProductQuery) FirstIDX(ctx context.Context) int64 {
	id, err := mpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MemberProduct entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MemberProduct entity is found.
// Returns a *NotFoundError when no MemberProduct entities are found.
func (mpq *MemberProductQuery) Only(ctx context.Context) (*MemberProduct, error) {
	nodes, err := mpq.Limit(2).All(setContextOp(ctx, mpq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{memberproduct.Label}
	default:
		return nil, &NotSingularError{memberproduct.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mpq *MemberProductQuery) OnlyX(ctx context.Context) *MemberProduct {
	node, err := mpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MemberProduct ID in the query.
// Returns a *NotSingularError when more than one MemberProduct ID is found.
// Returns a *NotFoundError when no entities are found.
func (mpq *MemberProductQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = mpq.Limit(2).IDs(setContextOp(ctx, mpq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{memberproduct.Label}
	default:
		err = &NotSingularError{memberproduct.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mpq *MemberProductQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := mpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MemberProducts.
func (mpq *MemberProductQuery) All(ctx context.Context) ([]*MemberProduct, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryAll)
	if err := mpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MemberProduct, *MemberProductQuery]()
	return withInterceptors[[]*MemberProduct](ctx, mpq, qr, mpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mpq *MemberProductQuery) AllX(ctx context.Context) []*MemberProduct {
	nodes, err := mpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MemberProduct IDs.
func (mpq *MemberProductQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if mpq.ctx.Unique == nil && mpq.path != nil {
		mpq.Unique(true)
	}
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryIDs)
	if err = mpq.Select(memberproduct.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mpq *MemberProductQuery) IDsX(ctx context.Context) []int64 {
	ids, err := mpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mpq *MemberProductQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryCount)
	if err := mpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mpq, querierCount[*MemberProductQuery](), mpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mpq *MemberProductQuery) CountX(ctx context.Context) int {
	count, err := mpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mpq *MemberProductQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mpq.ctx, ent.OpQueryExist)
	switch _, err := mpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mpq *MemberProductQuery) ExistX(ctx context.Context) bool {
	exist, err := mpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MemberProductQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mpq *MemberProductQuery) Clone() *MemberProductQuery {
	if mpq == nil {
		return nil
	}
	return &MemberProductQuery{
		config:                    mpq.config,
		ctx:                       mpq.ctx.Clone(),
		order:                     append([]memberproduct.OrderOption{}, mpq.order...),
		inters:                    append([]Interceptor{}, mpq.inters...),
		predicates:                append([]predicate.MemberProduct{}, mpq.predicates...),
		withMembers:               mpq.withMembers.Clone(),
		withMemberProductEntry:    mpq.withMemberProductEntry.Clone(),
		withMemberProductContents: mpq.withMemberProductContents.Clone(),
		withMemberCourses:         mpq.withMemberCourses.Clone(),
		withMemberLessons:         mpq.withMemberLessons.Clone(),
		// clone intermediate query.
		sql:  mpq.sql.Clone(),
		path: mpq.path,
	}
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MemberProductQuery) WithMembers(opts ...func(*MemberQuery)) *MemberProductQuery {
	query := (&MemberClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMembers = query
	return mpq
}

// WithMemberProductEntry tells the query-builder to eager-load the nodes that are connected to
// the "member_product_entry" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MemberProductQuery) WithMemberProductEntry(opts ...func(*EntryLogsQuery)) *MemberProductQuery {
	query := (&EntryLogsClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMemberProductEntry = query
	return mpq
}

// WithMemberProductContents tells the query-builder to eager-load the nodes that are connected to
// the "member_product_contents" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MemberProductQuery) WithMemberProductContents(opts ...func(*MemberContractQuery)) *MemberProductQuery {
	query := (&MemberContractClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMemberProductContents = query
	return mpq
}

// WithMemberCourses tells the query-builder to eager-load the nodes that are connected to
// the "memberCourses" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MemberProductQuery) WithMemberCourses(opts ...func(*MemberProductCoursesQuery)) *MemberProductQuery {
	query := (&MemberProductCoursesClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMemberCourses = query
	return mpq
}

// WithMemberLessons tells the query-builder to eager-load the nodes that are connected to
// the "memberLessons" edge. The optional arguments are used to configure the query builder of the edge.
func (mpq *MemberProductQuery) WithMemberLessons(opts ...func(*MemberProductCoursesQuery)) *MemberProductQuery {
	query := (&MemberProductCoursesClient{config: mpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mpq.withMemberLessons = query
	return mpq
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
//	client.MemberProduct.Query().
//		GroupBy(memberproduct.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mpq *MemberProductQuery) GroupBy(field string, fields ...string) *MemberProductGroupBy {
	mpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MemberProductGroupBy{build: mpq}
	grbuild.flds = &mpq.ctx.Fields
	grbuild.label = memberproduct.Label
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
//	client.MemberProduct.Query().
//		Select(memberproduct.FieldCreatedAt).
//		Scan(ctx, &v)
func (mpq *MemberProductQuery) Select(fields ...string) *MemberProductSelect {
	mpq.ctx.Fields = append(mpq.ctx.Fields, fields...)
	sbuild := &MemberProductSelect{MemberProductQuery: mpq}
	sbuild.label = memberproduct.Label
	sbuild.flds, sbuild.scan = &mpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MemberProductSelect configured with the given aggregations.
func (mpq *MemberProductQuery) Aggregate(fns ...AggregateFunc) *MemberProductSelect {
	return mpq.Select().Aggregate(fns...)
}

func (mpq *MemberProductQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mpq); err != nil {
				return err
			}
		}
	}
	for _, f := range mpq.ctx.Fields {
		if !memberproduct.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mpq.path != nil {
		prev, err := mpq.path(ctx)
		if err != nil {
			return err
		}
		mpq.sql = prev
	}
	return nil
}

func (mpq *MemberProductQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MemberProduct, error) {
	var (
		nodes       = []*MemberProduct{}
		_spec       = mpq.querySpec()
		loadedTypes = [5]bool{
			mpq.withMembers != nil,
			mpq.withMemberProductEntry != nil,
			mpq.withMemberProductContents != nil,
			mpq.withMemberCourses != nil,
			mpq.withMemberLessons != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MemberProduct).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MemberProduct{config: mpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mpq.withMembers; query != nil {
		if err := mpq.loadMembers(ctx, query, nodes, nil,
			func(n *MemberProduct, e *Member) { n.Edges.Members = e }); err != nil {
			return nil, err
		}
	}
	if query := mpq.withMemberProductEntry; query != nil {
		if err := mpq.loadMemberProductEntry(ctx, query, nodes,
			func(n *MemberProduct) { n.Edges.MemberProductEntry = []*EntryLogs{} },
			func(n *MemberProduct, e *EntryLogs) {
				n.Edges.MemberProductEntry = append(n.Edges.MemberProductEntry, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := mpq.withMemberProductContents; query != nil {
		if err := mpq.loadMemberProductContents(ctx, query, nodes,
			func(n *MemberProduct) { n.Edges.MemberProductContents = []*MemberContract{} },
			func(n *MemberProduct, e *MemberContract) {
				n.Edges.MemberProductContents = append(n.Edges.MemberProductContents, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := mpq.withMemberCourses; query != nil {
		if err := mpq.loadMemberCourses(ctx, query, nodes,
			func(n *MemberProduct) { n.Edges.MemberCourses = []*MemberProductCourses{} },
			func(n *MemberProduct, e *MemberProductCourses) {
				n.Edges.MemberCourses = append(n.Edges.MemberCourses, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := mpq.withMemberLessons; query != nil {
		if err := mpq.loadMemberLessons(ctx, query, nodes,
			func(n *MemberProduct) { n.Edges.MemberLessons = []*MemberProductCourses{} },
			func(n *MemberProduct, e *MemberProductCourses) {
				n.Edges.MemberLessons = append(n.Edges.MemberLessons, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mpq *MemberProductQuery) loadMembers(ctx context.Context, query *MemberQuery, nodes []*MemberProduct, init func(*MemberProduct), assign func(*MemberProduct, *Member)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*MemberProduct)
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
func (mpq *MemberProductQuery) loadMemberProductEntry(ctx context.Context, query *EntryLogsQuery, nodes []*MemberProduct, init func(*MemberProduct), assign func(*MemberProduct, *EntryLogs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MemberProduct)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(entrylogs.FieldMemberProductID)
	}
	query.Where(predicate.EntryLogs(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(memberproduct.MemberProductEntryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mpq *MemberProductQuery) loadMemberProductContents(ctx context.Context, query *MemberContractQuery, nodes []*MemberProduct, init func(*MemberProduct), assign func(*MemberProduct, *MemberContract)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MemberProduct)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(membercontract.FieldMemberProductID)
	}
	query.Where(predicate.MemberContract(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(memberproduct.MemberProductContentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mpq *MemberProductQuery) loadMemberCourses(ctx context.Context, query *MemberProductCoursesQuery, nodes []*MemberProduct, init func(*MemberProduct), assign func(*MemberProduct, *MemberProductCourses)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MemberProduct)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(memberproductcourses.FieldMemberProductID)
	}
	query.Where(predicate.MemberProductCourses(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(memberproduct.MemberCoursesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (mpq *MemberProductQuery) loadMemberLessons(ctx context.Context, query *MemberProductCoursesQuery, nodes []*MemberProduct, init func(*MemberProduct), assign func(*MemberProduct, *MemberProductCourses)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*MemberProduct)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(memberproductcourses.FieldMemberProductID)
	}
	query.Where(predicate.MemberProductCourses(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(memberproduct.MemberLessonsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MemberProductID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "member_product_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mpq *MemberProductQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mpq.querySpec()
	_spec.Node.Columns = mpq.ctx.Fields
	if len(mpq.ctx.Fields) > 0 {
		_spec.Unique = mpq.ctx.Unique != nil && *mpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mpq.driver, _spec)
}

func (mpq *MemberProductQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(memberproduct.Table, memberproduct.Columns, sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64))
	_spec.From = mpq.sql
	if unique := mpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mpq.path != nil {
		_spec.Unique = true
	}
	if fields := mpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memberproduct.FieldID)
		for i := range fields {
			if fields[i] != memberproduct.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if mpq.withMembers != nil {
			_spec.Node.AddColumnOnce(memberproduct.FieldMemberID)
		}
	}
	if ps := mpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mpq *MemberProductQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mpq.driver.Dialect())
	t1 := builder.Table(memberproduct.Table)
	columns := mpq.ctx.Fields
	if len(columns) == 0 {
		columns = memberproduct.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mpq.sql != nil {
		selector = mpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mpq.ctx.Unique != nil && *mpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mpq.predicates {
		p(selector)
	}
	for _, p := range mpq.order {
		p(selector)
	}
	if offset := mpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MemberProductGroupBy is the group-by builder for MemberProduct entities.
type MemberProductGroupBy struct {
	selector
	build *MemberProductQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mpgb *MemberProductGroupBy) Aggregate(fns ...AggregateFunc) *MemberProductGroupBy {
	mpgb.fns = append(mpgb.fns, fns...)
	return mpgb
}

// Scan applies the selector query and scans the result into the given value.
func (mpgb *MemberProductGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mpgb.build.ctx, ent.OpQueryGroupBy)
	if err := mpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberProductQuery, *MemberProductGroupBy](ctx, mpgb.build, mpgb, mpgb.build.inters, v)
}

func (mpgb *MemberProductGroupBy) sqlScan(ctx context.Context, root *MemberProductQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mpgb.fns))
	for _, fn := range mpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mpgb.flds)+len(mpgb.fns))
		for _, f := range *mpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MemberProductSelect is the builder for selecting fields of MemberProduct entities.
type MemberProductSelect struct {
	*MemberProductQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mps *MemberProductSelect) Aggregate(fns ...AggregateFunc) *MemberProductSelect {
	mps.fns = append(mps.fns, fns...)
	return mps
}

// Scan applies the selector query and scans the result into the given value.
func (mps *MemberProductSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mps.ctx, ent.OpQuerySelect)
	if err := mps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MemberProductQuery, *MemberProductSelect](ctx, mps.MemberProductQuery, mps, mps.inters, v)
}

func (mps *MemberProductSelect) sqlScan(ctx context.Context, root *MemberProductQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mps.fns))
	for _, fn := range mps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
