// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VenueQuery is the builder for querying Venue entities.
type VenueQuery struct {
	config
	ctx             *QueryContext
	order           []venue.OrderOption
	inters          []Interceptor
	predicates      []predicate.Venue
	withPlaces      *VenuePlaceQuery
	withVenueOrders *OrderQuery
	withVenueEntry  *EntryLogsQuery
	withUsers       *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VenueQuery builder.
func (vq *VenueQuery) Where(ps ...predicate.Venue) *VenueQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VenueQuery) Limit(limit int) *VenueQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VenueQuery) Offset(offset int) *VenueQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VenueQuery) Unique(unique bool) *VenueQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VenueQuery) Order(o ...venue.OrderOption) *VenueQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryPlaces chains the current query on the "places" edge.
func (vq *VenueQuery) QueryPlaces() *VenuePlaceQuery {
	query := (&VenuePlaceClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, selector),
			sqlgraph.To(venueplace.Table, venueplace.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, venue.PlacesTable, venue.PlacesColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVenueOrders chains the current query on the "venue_orders" edge.
func (vq *VenueQuery) QueryVenueOrders() *OrderQuery {
	query := (&OrderClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, selector),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, venue.VenueOrdersTable, venue.VenueOrdersColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVenueEntry chains the current query on the "venue_entry" edge.
func (vq *VenueQuery) QueryVenueEntry() *EntryLogsQuery {
	query := (&EntryLogsClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, selector),
			sqlgraph.To(entrylogs.Table, entrylogs.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, venue.VenueEntryTable, venue.VenueEntryColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUsers chains the current query on the "users" edge.
func (vq *VenueQuery) QueryUsers() *UserQuery {
	query := (&UserClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(venue.Table, venue.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, venue.UsersTable, venue.UsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Venue entity from the query.
// Returns a *NotFoundError when no Venue was found.
func (vq *VenueQuery) First(ctx context.Context) (*Venue, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{venue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VenueQuery) FirstX(ctx context.Context) *Venue {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Venue ID from the query.
// Returns a *NotFoundError when no Venue ID was found.
func (vq *VenueQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{venue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VenueQuery) FirstIDX(ctx context.Context) int64 {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Venue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Venue entity is found.
// Returns a *NotFoundError when no Venue entities are found.
func (vq *VenueQuery) Only(ctx context.Context) (*Venue, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{venue.Label}
	default:
		return nil, &NotSingularError{venue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VenueQuery) OnlyX(ctx context.Context) *Venue {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Venue ID in the query.
// Returns a *NotSingularError when more than one Venue ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VenueQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{venue.Label}
	default:
		err = &NotSingularError{venue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VenueQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Venues.
func (vq *VenueQuery) All(ctx context.Context) ([]*Venue, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Venue, *VenueQuery]()
	return withInterceptors[[]*Venue](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VenueQuery) AllX(ctx context.Context) []*Venue {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Venue IDs.
func (vq *VenueQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(venue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VenueQuery) IDsX(ctx context.Context) []int64 {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VenueQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VenueQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VenueQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VenueQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VenueQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VenueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VenueQuery) Clone() *VenueQuery {
	if vq == nil {
		return nil
	}
	return &VenueQuery{
		config:          vq.config,
		ctx:             vq.ctx.Clone(),
		order:           append([]venue.OrderOption{}, vq.order...),
		inters:          append([]Interceptor{}, vq.inters...),
		predicates:      append([]predicate.Venue{}, vq.predicates...),
		withPlaces:      vq.withPlaces.Clone(),
		withVenueOrders: vq.withVenueOrders.Clone(),
		withVenueEntry:  vq.withVenueEntry.Clone(),
		withUsers:       vq.withUsers.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithPlaces tells the query-builder to eager-load the nodes that are connected to
// the "places" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VenueQuery) WithPlaces(opts ...func(*VenuePlaceQuery)) *VenueQuery {
	query := (&VenuePlaceClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withPlaces = query
	return vq
}

// WithVenueOrders tells the query-builder to eager-load the nodes that are connected to
// the "venue_orders" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VenueQuery) WithVenueOrders(opts ...func(*OrderQuery)) *VenueQuery {
	query := (&OrderClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVenueOrders = query
	return vq
}

// WithVenueEntry tells the query-builder to eager-load the nodes that are connected to
// the "venue_entry" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VenueQuery) WithVenueEntry(opts ...func(*EntryLogsQuery)) *VenueQuery {
	query := (&EntryLogsClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withVenueEntry = query
	return vq
}

// WithUsers tells the query-builder to eager-load the nodes that are connected to
// the "users" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VenueQuery) WithUsers(opts ...func(*UserQuery)) *VenueQuery {
	query := (&UserClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withUsers = query
	return vq
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
//	client.Venue.Query().
//		GroupBy(venue.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VenueQuery) GroupBy(field string, fields ...string) *VenueGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VenueGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = venue.Label
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
//	client.Venue.Query().
//		Select(venue.FieldCreatedAt).
//		Scan(ctx, &v)
func (vq *VenueQuery) Select(fields ...string) *VenueSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VenueSelect{VenueQuery: vq}
	sbuild.label = venue.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VenueSelect configured with the given aggregations.
func (vq *VenueQuery) Aggregate(fns ...AggregateFunc) *VenueSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VenueQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !venue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VenueQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Venue, error) {
	var (
		nodes       = []*Venue{}
		_spec       = vq.querySpec()
		loadedTypes = [4]bool{
			vq.withPlaces != nil,
			vq.withVenueOrders != nil,
			vq.withVenueEntry != nil,
			vq.withUsers != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Venue).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Venue{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withPlaces; query != nil {
		if err := vq.loadPlaces(ctx, query, nodes,
			func(n *Venue) { n.Edges.Places = []*VenuePlace{} },
			func(n *Venue, e *VenuePlace) { n.Edges.Places = append(n.Edges.Places, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVenueOrders; query != nil {
		if err := vq.loadVenueOrders(ctx, query, nodes,
			func(n *Venue) { n.Edges.VenueOrders = []*Order{} },
			func(n *Venue, e *Order) { n.Edges.VenueOrders = append(n.Edges.VenueOrders, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withVenueEntry; query != nil {
		if err := vq.loadVenueEntry(ctx, query, nodes,
			func(n *Venue) { n.Edges.VenueEntry = []*EntryLogs{} },
			func(n *Venue, e *EntryLogs) { n.Edges.VenueEntry = append(n.Edges.VenueEntry, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withUsers; query != nil {
		if err := vq.loadUsers(ctx, query, nodes,
			func(n *Venue) { n.Edges.Users = []*User{} },
			func(n *Venue, e *User) { n.Edges.Users = append(n.Edges.Users, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VenueQuery) loadPlaces(ctx context.Context, query *VenuePlaceQuery, nodes []*Venue, init func(*Venue), assign func(*Venue, *VenuePlace)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Venue)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(venueplace.FieldVenueID)
	}
	query.Where(predicate.VenuePlace(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(venue.PlacesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VenueID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "venue_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VenueQuery) loadVenueOrders(ctx context.Context, query *OrderQuery, nodes []*Venue, init func(*Venue), assign func(*Venue, *Order)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Venue)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(order.FieldVenueID)
	}
	query.Where(predicate.Order(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(venue.VenueOrdersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VenueID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "venue_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VenueQuery) loadVenueEntry(ctx context.Context, query *EntryLogsQuery, nodes []*Venue, init func(*Venue), assign func(*Venue, *EntryLogs)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Venue)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(entrylogs.FieldVenueID)
	}
	query.Where(predicate.EntryLogs(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(venue.VenueEntryColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.VenueID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "venue_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VenueQuery) loadUsers(ctx context.Context, query *UserQuery, nodes []*Venue, init func(*Venue), assign func(*Venue, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int64]*Venue)
	nids := make(map[int64]map[*Venue]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(venue.UsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(venue.UsersPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(venue.UsersPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(venue.UsersPrimaryKey[1]))
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
					nids[inValue] = map[*Venue]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (vq *VenueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VenueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(venue.Table, venue.Columns, sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, venue.FieldID)
		for i := range fields {
			if fields[i] != venue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VenueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(venue.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = venue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VenueGroupBy is the group-by builder for Venue entities.
type VenueGroupBy struct {
	selector
	build *VenueQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VenueGroupBy) Aggregate(fns ...AggregateFunc) *VenueGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VenueGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VenueQuery, *VenueGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VenueGroupBy) sqlScan(ctx context.Context, root *VenueQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VenueSelect is the builder for selecting fields of Venue entities.
type VenueSelect struct {
	*VenueQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VenueSelect) Aggregate(fns ...AggregateFunc) *VenueSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VenueSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VenueQuery, *VenueSelect](ctx, vs.VenueQuery, vs, vs.inters, v)
}

func (vs *VenueSelect) sqlScan(ctx context.Context, root *VenueQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
