// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"order/biz/ent/migrate"

	"order/biz/ent/order"
	"order/biz/ent/orderamount"
	"order/biz/ent/orderitem"
	"order/biz/ent/orderpay"
	"order/biz/ent/ordersales"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.




	Schema *migrate.Schema
	// Order is the client for interacting with the Order builders.
	Order *OrderClient
	// OrderAmount is the client for interacting with the OrderAmount builders.
	OrderAmount *OrderAmountClient
	// OrderItem is the client for interacting with the OrderItem builders.
	OrderItem *OrderItemClient
	// OrderPay is the client for interacting with the OrderPay builders.
	OrderPay *OrderPayClient
	// OrderSales is the client for interacting with the OrderSales builders.
	OrderSales *OrderSalesClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Order = NewOrderClient(c.config)
	c.OrderAmount = NewOrderAmountClient(c.config)
	c.OrderItem = NewOrderItemClient(c.config)
	c.OrderPay = NewOrderPayClient(c.config)
	c.OrderSales = NewOrderSalesClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Order:       NewOrderClient(cfg),
		OrderAmount: NewOrderAmountClient(cfg),
		OrderItem:   NewOrderItemClient(cfg),
		OrderPay:    NewOrderPayClient(cfg),
		OrderSales:  NewOrderSalesClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Order:       NewOrderClient(cfg),
		OrderAmount: NewOrderAmountClient(cfg),
		OrderItem:   NewOrderItemClient(cfg),
		OrderPay:    NewOrderPayClient(cfg),
		OrderSales:  NewOrderSalesClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Order.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Order.Use(hooks...)
	c.OrderAmount.Use(hooks...)
	c.OrderItem.Use(hooks...)
	c.OrderPay.Use(hooks...)
	c.OrderSales.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Order.Intercept(interceptors...)
	c.OrderAmount.Intercept(interceptors...)
	c.OrderItem.Intercept(interceptors...)
	c.OrderPay.Intercept(interceptors...)
	c.OrderSales.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *OrderMutation:
		return c.Order.mutate(ctx, m)
	case *OrderAmountMutation:
		return c.OrderAmount.mutate(ctx, m)
	case *OrderItemMutation:
		return c.OrderItem.mutate(ctx, m)
	case *OrderPayMutation:
		return c.OrderPay.mutate(ctx, m)
	case *OrderSalesMutation:
		return c.OrderSales.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// OrderClient is a client for the Order schema.
type OrderClient struct {
	config
}

// NewOrderClient returns a client for the Order from the given config.
func NewOrderClient(c config) *OrderClient {
	return &OrderClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `order.Hooks(f(g(h())))`.
func (c *OrderClient) Use(hooks ...Hook) {
	c.hooks.Order = append(c.hooks.Order, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `order.Intercept(f(g(h())))`.
func (c *OrderClient) Intercept(interceptors ...Interceptor) {
	c.inters.Order = append(c.inters.Order, interceptors...)
}

// Create returns a builder for creating a Order entity.
func (c *OrderClient) Create() *OrderCreate {
	mutation := newOrderMutation(c.config, OpCreate)
	return &OrderCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Order entities.
func (c *OrderClient) CreateBulk(builders ...*OrderCreate) *OrderCreateBulk {
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderClient) MapCreateBulk(slice any, setFunc func(*OrderCreate, int)) *OrderCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderCreateBulk{err: fmt.Errorf("calling to OrderClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Order.
func (c *OrderClient) Update() *OrderUpdate {
	mutation := newOrderMutation(c.config, OpUpdate)
	return &OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderClient) UpdateOne(o *Order) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrder(o))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderClient) UpdateOneID(id int64) *OrderUpdateOne {
	mutation := newOrderMutation(c.config, OpUpdateOne, withOrderID(id))
	return &OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Order.
func (c *OrderClient) Delete() *OrderDelete {
	mutation := newOrderMutation(c.config, OpDelete)
	return &OrderDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderClient) DeleteOne(o *Order) *OrderDeleteOne {
	return c.DeleteOneID(o.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderClient) DeleteOneID(id int64) *OrderDeleteOne {
	builder := c.Delete().Where(order.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderDeleteOne{builder}
}

// Query returns a query builder for Order.
func (c *OrderClient) Query() *OrderQuery {
	return &OrderQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrder},
		inters: c.Interceptors(),
	}
}

// Get returns a Order entity by its id.
func (c *OrderClient) Get(ctx context.Context, id int64) (*Order, error) {
	return c.Query().Where(order.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderClient) GetX(ctx context.Context, id int64) *Order {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAmount queries the amount edge of a Order.
func (c *OrderClient) QueryAmount(o *Order) *OrderAmountQuery {
	query := (&OrderAmountClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(orderamount.Table, orderamount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.AmountTable, order.AmountColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryItem queries the item edge of a Order.
func (c *OrderClient) QueryItem(o *Order) *OrderItemQuery {
	query := (&OrderItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(orderitem.Table, orderitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.ItemTable, order.ItemColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPay queries the pay edge of a Order.
func (c *OrderClient) QueryPay(o *Order) *OrderPayQuery {
	query := (&OrderPayClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(orderpay.Table, orderpay.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.PayTable, order.PayColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySales queries the sales edge of a Order.
func (c *OrderClient) QuerySales(o *Order) *OrderSalesQuery {
	query := (&OrderSalesClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := o.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(order.Table, order.FieldID, id),
			sqlgraph.To(ordersales.Table, ordersales.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, order.SalesTable, order.SalesColumn),
		)
		fromV = sqlgraph.Neighbors(o.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderClient) Hooks() []Hook {
	return c.hooks.Order
}

// Interceptors returns the client interceptors.
func (c *OrderClient) Interceptors() []Interceptor {
	return c.inters.Order
}

func (c *OrderClient) mutate(ctx context.Context, m *OrderMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Order mutation op: %q", m.Op())
	}
}

// OrderAmountClient is a client for the OrderAmount schema.
type OrderAmountClient struct {
	config
}

// NewOrderAmountClient returns a client for the OrderAmount from the given config.
func NewOrderAmountClient(c config) *OrderAmountClient {
	return &OrderAmountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `orderamount.Hooks(f(g(h())))`.
func (c *OrderAmountClient) Use(hooks ...Hook) {
	c.hooks.OrderAmount = append(c.hooks.OrderAmount, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `orderamount.Intercept(f(g(h())))`.
func (c *OrderAmountClient) Intercept(interceptors ...Interceptor) {
	c.inters.OrderAmount = append(c.inters.OrderAmount, interceptors...)
}

// Create returns a builder for creating a OrderAmount entity.
func (c *OrderAmountClient) Create() *OrderAmountCreate {
	mutation := newOrderAmountMutation(c.config, OpCreate)
	return &OrderAmountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OrderAmount entities.
func (c *OrderAmountClient) CreateBulk(builders ...*OrderAmountCreate) *OrderAmountCreateBulk {
	return &OrderAmountCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderAmountClient) MapCreateBulk(slice any, setFunc func(*OrderAmountCreate, int)) *OrderAmountCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderAmountCreateBulk{err: fmt.Errorf("calling to OrderAmountClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderAmountCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderAmountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OrderAmount.
func (c *OrderAmountClient) Update() *OrderAmountUpdate {
	mutation := newOrderAmountMutation(c.config, OpUpdate)
	return &OrderAmountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderAmountClient) UpdateOne(oa *OrderAmount) *OrderAmountUpdateOne {
	mutation := newOrderAmountMutation(c.config, OpUpdateOne, withOrderAmount(oa))
	return &OrderAmountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderAmountClient) UpdateOneID(id int64) *OrderAmountUpdateOne {
	mutation := newOrderAmountMutation(c.config, OpUpdateOne, withOrderAmountID(id))
	return &OrderAmountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OrderAmount.
func (c *OrderAmountClient) Delete() *OrderAmountDelete {
	mutation := newOrderAmountMutation(c.config, OpDelete)
	return &OrderAmountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderAmountClient) DeleteOne(oa *OrderAmount) *OrderAmountDeleteOne {
	return c.DeleteOneID(oa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderAmountClient) DeleteOneID(id int64) *OrderAmountDeleteOne {
	builder := c.Delete().Where(orderamount.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderAmountDeleteOne{builder}
}

// Query returns a query builder for OrderAmount.
func (c *OrderAmountClient) Query() *OrderAmountQuery {
	return &OrderAmountQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrderAmount},
		inters: c.Interceptors(),
	}
}

// Get returns a OrderAmount entity by its id.
func (c *OrderAmountClient) Get(ctx context.Context, id int64) (*OrderAmount, error) {
	return c.Query().Where(orderamount.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderAmountClient) GetX(ctx context.Context, id int64) *OrderAmount {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrder queries the order edge of a OrderAmount.
func (c *OrderAmountClient) QueryOrder(oa *OrderAmount) *OrderQuery {
	query := (&OrderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := oa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(orderamount.Table, orderamount.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderamount.OrderTable, orderamount.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(oa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderAmountClient) Hooks() []Hook {
	return c.hooks.OrderAmount
}

// Interceptors returns the client interceptors.
func (c *OrderAmountClient) Interceptors() []Interceptor {
	return c.inters.OrderAmount
}

func (c *OrderAmountClient) mutate(ctx context.Context, m *OrderAmountMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderAmountCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderAmountUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderAmountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderAmountDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OrderAmount mutation op: %q", m.Op())
	}
}

// OrderItemClient is a client for the OrderItem schema.
type OrderItemClient struct {
	config
}

// NewOrderItemClient returns a client for the OrderItem from the given config.
func NewOrderItemClient(c config) *OrderItemClient {
	return &OrderItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `orderitem.Hooks(f(g(h())))`.
func (c *OrderItemClient) Use(hooks ...Hook) {
	c.hooks.OrderItem = append(c.hooks.OrderItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `orderitem.Intercept(f(g(h())))`.
func (c *OrderItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.OrderItem = append(c.inters.OrderItem, interceptors...)
}

// Create returns a builder for creating a OrderItem entity.
func (c *OrderItemClient) Create() *OrderItemCreate {
	mutation := newOrderItemMutation(c.config, OpCreate)
	return &OrderItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OrderItem entities.
func (c *OrderItemClient) CreateBulk(builders ...*OrderItemCreate) *OrderItemCreateBulk {
	return &OrderItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderItemClient) MapCreateBulk(slice any, setFunc func(*OrderItemCreate, int)) *OrderItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderItemCreateBulk{err: fmt.Errorf("calling to OrderItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OrderItem.
func (c *OrderItemClient) Update() *OrderItemUpdate {
	mutation := newOrderItemMutation(c.config, OpUpdate)
	return &OrderItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderItemClient) UpdateOne(oi *OrderItem) *OrderItemUpdateOne {
	mutation := newOrderItemMutation(c.config, OpUpdateOne, withOrderItem(oi))
	return &OrderItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderItemClient) UpdateOneID(id int64) *OrderItemUpdateOne {
	mutation := newOrderItemMutation(c.config, OpUpdateOne, withOrderItemID(id))
	return &OrderItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OrderItem.
func (c *OrderItemClient) Delete() *OrderItemDelete {
	mutation := newOrderItemMutation(c.config, OpDelete)
	return &OrderItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderItemClient) DeleteOne(oi *OrderItem) *OrderItemDeleteOne {
	return c.DeleteOneID(oi.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderItemClient) DeleteOneID(id int64) *OrderItemDeleteOne {
	builder := c.Delete().Where(orderitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderItemDeleteOne{builder}
}

// Query returns a query builder for OrderItem.
func (c *OrderItemClient) Query() *OrderItemQuery {
	return &OrderItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrderItem},
		inters: c.Interceptors(),
	}
}

// Get returns a OrderItem entity by its id.
func (c *OrderItemClient) Get(ctx context.Context, id int64) (*OrderItem, error) {
	return c.Query().Where(orderitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderItemClient) GetX(ctx context.Context, id int64) *OrderItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrder queries the order edge of a OrderItem.
func (c *OrderItemClient) QueryOrder(oi *OrderItem) *OrderQuery {
	query := (&OrderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := oi.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(orderitem.Table, orderitem.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderitem.OrderTable, orderitem.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(oi.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderItemClient) Hooks() []Hook {
	return c.hooks.OrderItem
}

// Interceptors returns the client interceptors.
func (c *OrderItemClient) Interceptors() []Interceptor {
	return c.inters.OrderItem
}

func (c *OrderItemClient) mutate(ctx context.Context, m *OrderItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OrderItem mutation op: %q", m.Op())
	}
}

// OrderPayClient is a client for the OrderPay schema.
type OrderPayClient struct {
	config
}

// NewOrderPayClient returns a client for the OrderPay from the given config.
func NewOrderPayClient(c config) *OrderPayClient {
	return &OrderPayClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `orderpay.Hooks(f(g(h())))`.
func (c *OrderPayClient) Use(hooks ...Hook) {
	c.hooks.OrderPay = append(c.hooks.OrderPay, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `orderpay.Intercept(f(g(h())))`.
func (c *OrderPayClient) Intercept(interceptors ...Interceptor) {
	c.inters.OrderPay = append(c.inters.OrderPay, interceptors...)
}

// Create returns a builder for creating a OrderPay entity.
func (c *OrderPayClient) Create() *OrderPayCreate {
	mutation := newOrderPayMutation(c.config, OpCreate)
	return &OrderPayCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OrderPay entities.
func (c *OrderPayClient) CreateBulk(builders ...*OrderPayCreate) *OrderPayCreateBulk {
	return &OrderPayCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderPayClient) MapCreateBulk(slice any, setFunc func(*OrderPayCreate, int)) *OrderPayCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderPayCreateBulk{err: fmt.Errorf("calling to OrderPayClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderPayCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderPayCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OrderPay.
func (c *OrderPayClient) Update() *OrderPayUpdate {
	mutation := newOrderPayMutation(c.config, OpUpdate)
	return &OrderPayUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderPayClient) UpdateOne(op *OrderPay) *OrderPayUpdateOne {
	mutation := newOrderPayMutation(c.config, OpUpdateOne, withOrderPay(op))
	return &OrderPayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderPayClient) UpdateOneID(id int64) *OrderPayUpdateOne {
	mutation := newOrderPayMutation(c.config, OpUpdateOne, withOrderPayID(id))
	return &OrderPayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OrderPay.
func (c *OrderPayClient) Delete() *OrderPayDelete {
	mutation := newOrderPayMutation(c.config, OpDelete)
	return &OrderPayDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderPayClient) DeleteOne(op *OrderPay) *OrderPayDeleteOne {
	return c.DeleteOneID(op.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderPayClient) DeleteOneID(id int64) *OrderPayDeleteOne {
	builder := c.Delete().Where(orderpay.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderPayDeleteOne{builder}
}

// Query returns a query builder for OrderPay.
func (c *OrderPayClient) Query() *OrderPayQuery {
	return &OrderPayQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrderPay},
		inters: c.Interceptors(),
	}
}

// Get returns a OrderPay entity by its id.
func (c *OrderPayClient) Get(ctx context.Context, id int64) (*OrderPay, error) {
	return c.Query().Where(orderpay.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderPayClient) GetX(ctx context.Context, id int64) *OrderPay {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrder queries the order edge of a OrderPay.
func (c *OrderPayClient) QueryOrder(op *OrderPay) *OrderQuery {
	query := (&OrderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := op.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(orderpay.Table, orderpay.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderpay.OrderTable, orderpay.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(op.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderPayClient) Hooks() []Hook {
	return c.hooks.OrderPay
}

// Interceptors returns the client interceptors.
func (c *OrderPayClient) Interceptors() []Interceptor {
	return c.inters.OrderPay
}

func (c *OrderPayClient) mutate(ctx context.Context, m *OrderPayMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderPayCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderPayUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderPayUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderPayDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OrderPay mutation op: %q", m.Op())
	}
}

// OrderSalesClient is a client for the OrderSales schema.
type OrderSalesClient struct {
	config
}

// NewOrderSalesClient returns a client for the OrderSales from the given config.
func NewOrderSalesClient(c config) *OrderSalesClient {
	return &OrderSalesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ordersales.Hooks(f(g(h())))`.
func (c *OrderSalesClient) Use(hooks ...Hook) {
	c.hooks.OrderSales = append(c.hooks.OrderSales, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ordersales.Intercept(f(g(h())))`.
func (c *OrderSalesClient) Intercept(interceptors ...Interceptor) {
	c.inters.OrderSales = append(c.inters.OrderSales, interceptors...)
}

// Create returns a builder for creating a OrderSales entity.
func (c *OrderSalesClient) Create() *OrderSalesCreate {
	mutation := newOrderSalesMutation(c.config, OpCreate)
	return &OrderSalesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of OrderSales entities.
func (c *OrderSalesClient) CreateBulk(builders ...*OrderSalesCreate) *OrderSalesCreateBulk {
	return &OrderSalesCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *OrderSalesClient) MapCreateBulk(slice any, setFunc func(*OrderSalesCreate, int)) *OrderSalesCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &OrderSalesCreateBulk{err: fmt.Errorf("calling to OrderSalesClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*OrderSalesCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &OrderSalesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for OrderSales.
func (c *OrderSalesClient) Update() *OrderSalesUpdate {
	mutation := newOrderSalesMutation(c.config, OpUpdate)
	return &OrderSalesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *OrderSalesClient) UpdateOne(os *OrderSales) *OrderSalesUpdateOne {
	mutation := newOrderSalesMutation(c.config, OpUpdateOne, withOrderSales(os))
	return &OrderSalesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *OrderSalesClient) UpdateOneID(id int64) *OrderSalesUpdateOne {
	mutation := newOrderSalesMutation(c.config, OpUpdateOne, withOrderSalesID(id))
	return &OrderSalesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for OrderSales.
func (c *OrderSalesClient) Delete() *OrderSalesDelete {
	mutation := newOrderSalesMutation(c.config, OpDelete)
	return &OrderSalesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *OrderSalesClient) DeleteOne(os *OrderSales) *OrderSalesDeleteOne {
	return c.DeleteOneID(os.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *OrderSalesClient) DeleteOneID(id int64) *OrderSalesDeleteOne {
	builder := c.Delete().Where(ordersales.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &OrderSalesDeleteOne{builder}
}

// Query returns a query builder for OrderSales.
func (c *OrderSalesClient) Query() *OrderSalesQuery {
	return &OrderSalesQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeOrderSales},
		inters: c.Interceptors(),
	}
}

// Get returns a OrderSales entity by its id.
func (c *OrderSalesClient) Get(ctx context.Context, id int64) (*OrderSales, error) {
	return c.Query().Where(ordersales.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *OrderSalesClient) GetX(ctx context.Context, id int64) *OrderSales {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryOrder queries the order edge of a OrderSales.
func (c *OrderSalesClient) QueryOrder(os *OrderSales) *OrderQuery {
	query := (&OrderClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := os.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ordersales.Table, ordersales.FieldID, id),
			sqlgraph.To(order.Table, order.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ordersales.OrderTable, ordersales.OrderColumn),
		)
		fromV = sqlgraph.Neighbors(os.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *OrderSalesClient) Hooks() []Hook {
	return c.hooks.OrderSales
}

// Interceptors returns the client interceptors.
func (c *OrderSalesClient) Interceptors() []Interceptor {
	return c.inters.OrderSales
}

func (c *OrderSalesClient) mutate(ctx context.Context, m *OrderSalesMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&OrderSalesCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&OrderSalesUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&OrderSalesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&OrderSalesDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown OrderSales mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Order, OrderAmount, OrderItem, OrderPay, OrderSales []ent.Hook
	}
	inters struct {
		Order, OrderAmount, OrderItem, OrderPay, OrderSales []ent.Interceptor
	}
)
