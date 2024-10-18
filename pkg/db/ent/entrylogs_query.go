// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"saas/pkg/db/ent/entrylogs"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/memberproduct"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/user"
	"saas/pkg/db/ent/venue"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntryLogsQuery is the builder for querying EntryLogs entities.
type EntryLogsQuery struct {
	config
	ctx                *QueryContext
	order              []entrylogs.OrderOption
	inters             []Interceptor
	predicates         []predicate.EntryLogs
	withVenues         *VenueQuery
	withMembers        *MemberQuery
	withUsers          *UserQuery
	withMemberProducts *MemberProductQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EntryLogsQuery builder.
func (elq *EntryLogsQuery) Where(ps ...predicate.EntryLogs) *EntryLogsQuery {
	elq.predicates = append(elq.predicates, ps...)
	return elq
}

// Limit the number of records to be returned by this query.
func (elq *EntryLogsQuery) Limit(limit int) *EntryLogsQuery {
	elq.ctx.Limit = &limit
	return elq
}

// Offset to start from.
func (elq *EntryLogsQuery) Offset(offset int) *EntryLogsQuery {
	elq.ctx.Offset = &offset
	return elq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (elq *EntryLogsQuery) Unique(unique bool) *EntryLogsQuery {
	elq.ctx.Unique = &unique
	return elq
}

// Order specifies how the records should be ordered.
func (elq *EntryLogsQuery) Order(o ...entrylogs.OrderOption) *EntryLogsQuery {
	elq.order = append(elq.order, o...)
	return elq
}

// QueryVenues chains the current query on the "venues" edge.
func (elq *EntryLogsQuery) QueryVenues() *VenueQuery {
	query := (&VenueClient{config: elq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := elq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entrylogs.Table, entrylogs.FieldID, selector),
			sqlgraph.To(venue.Table, venue.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entrylogs.VenuesTable, entrylogs.VenuesColumn),
		)
		fromU = sqlgraph.SetNeighbors(elq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMembers chains the current query on the "members" edge.
func (elq *EntryLogsQuery) QueryMembers() *MemberQuery {
	query := (&MemberClient{config: elq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := elq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entrylogs.Table, entrylogs.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entrylogs.MembersTable, entrylogs.MembersColumn),
		)
		fromU = sqlgraph.SetNeighbors(elq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (elq *EntryLogsQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: elq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := elq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entrylogs.Table, entrylogs.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entrylogs.UsersTable, entrylogs.UsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(elq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMemberProducts chains the current query on the "member_products" edge.
func (elq *EntryLogsQuery) QueryMemberProducts() *MemberProductQuery {
	query := (&MemberProductClient{config: elq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := elq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(entrylogs.Table, entrylogs.FieldID, selector),
			sqlgraph.To(memberproduct.Table, memberproduct.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, entrylogs.MemberProductsTable, entrylogs.MemberProductsColumn),
		)
		fromU = sqlgraph.SetNeighbors(elq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EntryLogs entity from the query.
// Returns a *NotFoundError when no EntryLogs was found.
func (elq *EntryLogsQuery) First(ctx context.Context) (*EntryLogs, error) {
	nodes, err := elq.Limit(1).All(setContextOp(ctx, elq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{entrylogs.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (elq *EntryLogsQuery) FirstX(ctx context.Context) *EntryLogs {
	node, err := elq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EntryLogs ID from the query.
// Returns a *NotFoundError when no EntryLogs ID was found.
func (elq *EntryLogsQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = elq.Limit(1).IDs(setContextOp(ctx, elq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{entrylogs.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (elq *EntryLogsQuery) FirstIDX(ctx context.Context) int64 {
	id, err := elq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EntryLogs entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EntryLogs entity is found.
// Returns a *NotFoundError when no EntryLogs entities are found.
func (elq *EntryLogsQuery) Only(ctx context.Context) (*EntryLogs, error) {
	nodes, err := elq.Limit(2).All(setContextOp(ctx, elq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{entrylogs.Label}
	default:
		return nil, &NotSingularError{entrylogs.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (elq *EntryLogsQuery) OnlyX(ctx context.Context) *EntryLogs {
	node, err := elq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EntryLogs ID in the query.
// Returns a *NotSingularError when more than one EntryLogs ID is found.
// Returns a *NotFoundError when no entities are found.
func (elq *EntryLogsQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = elq.Limit(2).IDs(setContextOp(ctx, elq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{entrylogs.Label}
	default:
		err = &NotSingularError{entrylogs.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (elq *EntryLogsQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := elq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EntryLogsSlice.
func (elq *EntryLogsQuery) All(ctx context.Context) ([]*EntryLogs, error) {
	ctx = setContextOp(ctx, elq.ctx, "All")
	if err := elq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EntryLogs, *EntryLogsQuery]()
	return withInterceptors[[]*EntryLogs](ctx, elq, qr, elq.inters)
}

// AllX is like All, but panics if an error occurs.
func (elq *EntryLogsQuery) AllX(ctx context.Context) []*EntryLogs {
	nodes, err := elq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EntryLogs IDs.
func (elq *EntryLogsQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if elq.ctx.Unique == nil && elq.path != nil {
		elq.Unique(true)
	}
	ctx = setContextOp(ctx, elq.ctx, "IDs")
	if err = elq.Select(entrylogs.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (elq *EntryLogsQuery) IDsX(ctx context.Context) []int64 {
	ids, err := elq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (elq *EntryLogsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, elq.ctx, "Count")
	if err := elq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, elq, querierCount[*EntryLogsQuery](), elq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (elq *EntryLogsQuery) CountX(ctx context.Context) int {
	count, err := elq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (elq *EntryLogsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, elq.ctx, "Exist")
	switch _, err := elq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (elq *EntryLogsQuery) ExistX(ctx context.Context) bool {
	exist, err := elq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EntryLogsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (elq *EntryLogsQuery) Clone() *EntryLogsQuery {
	if elq == nil {
		return nil
	}
	return &EntryLogsQuery{
		config:             elq.config,
		ctx:                elq.ctx.Clone(),
		order:              append([]entrylogs.OrderOption{}, elq.order...),
		inters:             append([]Interceptor{}, elq.inters...),
		predicates:         append([]predicate.EntryLogs{}, elq.predicates...),
		withVenues:         elq.withVenues.Clone(),
		withMembers:        elq.withMembers.Clone(),
		withUsers:          elq.withUsers.Clone(),
		withMemberProducts: elq.withMemberProducts.Clone(),
		// clone intermediate query.
		sql:  elq.sql.Clone(),
		path: elq.path,
	}
}

// WithVenues tells the query-builder to eager-load the nodes that are connected to
// the "venues" edge. The optional arguments are used to configure the query builder of the edge.
func (elq *EntryLogsQuery) WithVenues(opts ...func(*VenueQuery)) *EntryLogsQuery {
	query := (&VenueClient{config: elq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	elq.withVenues = query
	return elq
}

// WithMembers tells the query-builder to eager-load the nodes that are connected to
// the "members" edge. The optional arguments are used to configure the query builder of the edge.
func (elq *EntryLogsQuery) WithMembers(opts ...func(*MemberQuery)) *EntryLogsQuery {
	query := (&MemberClient{config: elq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	elq.withMembers = query
	return elq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (elq *EntryLogsQuery) WithUsers(opts ...func(*UserQuery)) *EntryLogsQuery {
	query := (&UserClient{config: elq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	elq.withUsers = query
	return elq
}

// WithMemberProducts tells the query-builder to eager-load the nodes that are connected to
// the "member_products" edge. The optional arguments are used to configure the query builder of the edge.
func (elq *EntryLogsQuery) WithMemberProducts(opts ...func(*MemberProductQuery)) *EntryLogsQuery {
	query := (&MemberProductClient{config: elq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	elq.withMemberProducts = query
	return elq
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
//	client.EntryLogs.Query().
//		GroupBy(entrylogs.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (elq *EntryLogsQuery) GroupBy(field string, fields ...string) *EntryLogsGroupBy {
	elq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EntryLogsGroupBy{build: elq}
	grbuild.flds = &elq.ctx.Fields
	grbuild.label = entrylogs.Label
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
//	client.EntryLogs.Query().
//		Select(entrylogs.FieldCreatedAt).
//		Scan(ctx, &v)
func (elq *EntryLogsQuery) Select(fields ...string) *EntryLogsSelect {
	elq.ctx.Fields = append(elq.ctx.Fields, fields...)
	sbuild := &EntryLogsSelect{EntryLogsQuery: elq}
	sbuild.label = entrylogs.Label
	sbuild.flds, sbuild.scan = &elq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EntryLogsSelect configured with the given aggregations.
func (elq *EntryLogsQuery) Aggregate(fns ...AggregateFunc) *EntryLogsSelect {
	return elq.Select().Aggregate(fns...)
}

func (elq *EntryLogsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range elq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, elq); err != nil {
				return err
			}
		}
	}
	for _, f := range elq.ctx.Fields {
		if !entrylogs.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if elq.path != nil {
		prev, err := elq.path(ctx)
		if err != nil {
			return err
		}
		elq.sql = prev
	}
	return nil
}

func (elq *EntryLogsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EntryLogs, error) {
	var (
		nodes       = []*EntryLogs{}
		_spec       = elq.querySpec()
		loadedTypes = [4]bool{
			elq.withVenues != nil,
			elq.withMembers != nil,
			elq.withUsers != nil,
			elq.withMemberProducts != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EntryLogs).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EntryLogs{config: elq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, elq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := elq.withVenues; query != nil {
		if err := elq.loadVenues(ctx, query, nodes, nil,
			func(n *EntryLogs, e *Venue) { n.Edges.Venues = e }); err != nil {
			return nil, err
		}
	}
	if query := elq.withMembers; query != nil {
		if err := elq.loadMembers(ctx, query, nodes, nil,
			func(n *EntryLogs, e *Member) { n.Edges.Members = e }); err != nil {
			return nil, err
		}
	}
	if query := elq.withUsers; query != nil {
		if err := elq.loadUsers(ctx, query, nodes, nil,
			func(n *EntryLogs, e *User) { n.Edges.Users = e }); err != nil {
			return nil, err
		}
	}
	if query := elq.withMemberProducts; query != nil {
		if err := elq.loadMemberProducts(ctx, query, nodes, nil,
			func(n *EntryLogs, e *MemberProduct) { n.Edges.MemberProducts = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (elq *EntryLogsQuery) loadVenues(ctx context.Context, query *VenueQuery, nodes []*EntryLogs, init func(*EntryLogs), assign func(*EntryLogs, *Venue)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*EntryLogs)
	for i := range nodes {
		fk := nodes[i].VenueID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(venue.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "venue_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (elq *EntryLogsQuery) loadMembers(ctx context.Context, query *MemberQuery, nodes []*EntryLogs, init func(*EntryLogs), assign func(*EntryLogs, *Member)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*EntryLogs)
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
func (elq *EntryLogsQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*EntryLogs, init func(*EntryLogs), assign func(*EntryLogs, *User)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*EntryLogs)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (elq *EntryLogsQuery) loadMemberProducts(ctx context.Context, query *MemberProductQuery, nodes []*EntryLogs, init func(*EntryLogs), assign func(*EntryLogs, *MemberProduct)) error {
	ids := make([]int64, 0, len(nodes))
	nodeids := make(map[int64][]*EntryLogs)
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

func (elq *EntryLogsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := elq.querySpec()
	_spec.Node.Columns = elq.ctx.Fields
	if len(elq.ctx.Fields) > 0 {
		_spec.Unique = elq.ctx.Unique != nil && *elq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, elq.driver, _spec)
}

func (elq *EntryLogsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(entrylogs.Table, entrylogs.Columns, sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64))
	_spec.From = elq.sql
	if unique := elq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if elq.path != nil {
		_spec.Unique = true
	}
	if fields := elq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entrylogs.FieldID)
		for i := range fields {
			if fields[i] != entrylogs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if elq.withVenues != nil {
			_spec.Node.AddColumnOnce(entrylogs.FieldVenueID)
		}
		if elq.withMembers != nil {
			_spec.Node.AddColumnOnce(entrylogs.FieldMemberID)
		}
		if elq.withUsers != nil {
			_spec.Node.AddColumnOnce(entrylogs.FieldUserID)
		}
		if elq.withMemberProducts != nil {
			_spec.Node.AddColumnOnce(entrylogs.FieldMemberProductID)
		}
	}
	if ps := elq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := elq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := elq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := elq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (elq *EntryLogsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(elq.driver.Dialect())
	t1 := builder.Table(entrylogs.Table)
	columns := elq.ctx.Fields
	if len(columns) == 0 {
		columns = entrylogs.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if elq.sql != nil {
		selector = elq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if elq.ctx.Unique != nil && *elq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range elq.predicates {
		p(selector)
	}
	for _, p := range elq.order {
		p(selector)
	}
	if offset := elq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := elq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EntryLogsGroupBy is the group-by builder for EntryLogs entities.
type EntryLogsGroupBy struct {
	selector
	build *EntryLogsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (elgb *EntryLogsGroupBy) Aggregate(fns ...AggregateFunc) *EntryLogsGroupBy {
	elgb.fns = append(elgb.fns, fns...)
	return elgb
}

// Scan applies the selector query and scans the result into the given value.
func (elgb *EntryLogsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, elgb.build.ctx, "GroupBy")
	if err := elgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntryLogsQuery, *EntryLogsGroupBy](ctx, elgb.build, elgb, elgb.build.inters, v)
}

func (elgb *EntryLogsGroupBy) sqlScan(ctx context.Context, root *EntryLogsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(elgb.fns))
	for _, fn := range elgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*elgb.flds)+len(elgb.fns))
		for _, f := range *elgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*elgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := elgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EntryLogsSelect is the builder for selecting fields of EntryLogs entities.
type EntryLogsSelect struct {
	*EntryLogsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (els *EntryLogsSelect) Aggregate(fns ...AggregateFunc) *EntryLogsSelect {
	els.fns = append(els.fns, fns...)
	return els
}

// Scan applies the selector query and scans the result into the given value.
func (els *EntryLogsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, els.ctx, "Select")
	if err := els.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EntryLogsQuery, *EntryLogsSelect](ctx, els.EntryLogsQuery, els, els.inters, v)
}

func (els *EntryLogsSelect) sqlScan(ctx context.Context, root *EntryLogsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(els.fns))
	for _, fn := range els.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*els.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := els.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
