// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDelete holds the string denoting the delete field in the database.
	FieldDelete = "delete"
	// FieldCreatedID holds the string denoting the created_id field in the database.
	FieldCreatedID = "created_id"
	// FieldOrderSn holds the string denoting the order_sn field in the database.
	FieldOrderSn = "order_sn"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldNature holds the string denoting the nature field in the database.
	FieldNature = "nature"
	// FieldProductType holds the string denoting the product_type field in the database.
	FieldProductType = "product_type"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// FieldDevice holds the string denoting the device field in the database.
	FieldDevice = "device"
	// FieldCompletionAt holds the string denoting the completion_at field in the database.
	FieldCompletionAt = "completion_at"
	// FieldCreateID holds the string denoting the create_id field in the database.
	FieldCreateID = "create_id"
	// EdgeAmount holds the string denoting the amount edge name in mutations.
	EdgeAmount = "amount"
	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// EdgePay holds the string denoting the pay edge name in mutations.
	EdgePay = "pay"
	// EdgeOrderContents holds the string denoting the order_contents edge name in mutations.
	EdgeOrderContents = "order_contents"
	// EdgeSales holds the string denoting the sales edge name in mutations.
	EdgeSales = "sales"
	// EdgeOrderVenues holds the string denoting the order_venues edge name in mutations.
	EdgeOrderVenues = "order_venues"
	// EdgeOrderMembers holds the string denoting the order_members edge name in mutations.
	EdgeOrderMembers = "order_members"
	// EdgeOrderCreates holds the string denoting the order_creates edge name in mutations.
	EdgeOrderCreates = "order_creates"
	// Table holds the table name of the order in the database.
	Table = "order"
	// AmountTable is the table that holds the amount relation/edge.
	AmountTable = "order_amount"
	// AmountInverseTable is the table name for the OrderAmount entity.
	// It exists in this package in order to avoid circular dependency with the "orderamount" package.
	AmountInverseTable = "order_amount"
	// AmountColumn is the table column denoting the amount relation/edge.
	AmountColumn = "order_id"
	// ItemTable is the table that holds the item relation/edge.
	ItemTable = "order_item"
	// ItemInverseTable is the table name for the OrderItem entity.
	// It exists in this package in order to avoid circular dependency with the "orderitem" package.
	ItemInverseTable = "order_item"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "order_id"
	// PayTable is the table that holds the pay relation/edge.
	PayTable = "order_pay"
	// PayInverseTable is the table name for the OrderPay entity.
	// It exists in this package in order to avoid circular dependency with the "orderpay" package.
	PayInverseTable = "order_pay"
	// PayColumn is the table column denoting the pay relation/edge.
	PayColumn = "order_id"
	// OrderContentsTable is the table that holds the order_contents relation/edge.
	OrderContentsTable = "member_contract"
	// OrderContentsInverseTable is the table name for the MemberContract entity.
	// It exists in this package in order to avoid circular dependency with the "membercontract" package.
	OrderContentsInverseTable = "member_contract"
	// OrderContentsColumn is the table column denoting the order_contents relation/edge.
	OrderContentsColumn = "order_id"
	// SalesTable is the table that holds the sales relation/edge.
	SalesTable = "order_sales"
	// SalesInverseTable is the table name for the OrderSales entity.
	// It exists in this package in order to avoid circular dependency with the "ordersales" package.
	SalesInverseTable = "order_sales"
	// SalesColumn is the table column denoting the sales relation/edge.
	SalesColumn = "order_id"
	// OrderVenuesTable is the table that holds the order_venues relation/edge.
	OrderVenuesTable = "order"
	// OrderVenuesInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	OrderVenuesInverseTable = "venue"
	// OrderVenuesColumn is the table column denoting the order_venues relation/edge.
	OrderVenuesColumn = "venue_id"
	// OrderMembersTable is the table that holds the order_members relation/edge.
	OrderMembersTable = "order"
	// OrderMembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	OrderMembersInverseTable = "member"
	// OrderMembersColumn is the table column denoting the order_members relation/edge.
	OrderMembersColumn = "member_id"
	// OrderCreatesTable is the table that holds the order_creates relation/edge.
	OrderCreatesTable = "order"
	// OrderCreatesInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OrderCreatesInverseTable = "sys_users"
	// OrderCreatesColumn is the table column denoting the order_creates relation/edge.
	OrderCreatesColumn = "create_id"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldOrderSn,
	FieldVenueID,
	FieldMemberID,
	FieldNature,
	FieldProductType,
	FieldStatus,
	FieldSource,
	FieldDevice,
	FieldCompletionAt,
	FieldCreateID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultDelete holds the default value on creation for the "delete" field.
	DefaultDelete int64
	// DefaultCreatedID holds the default value on creation for the "created_id" field.
	DefaultCreatedID int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultSource holds the default value on creation for the "source" field.
	DefaultSource string
	// DefaultDevice holds the default value on creation for the "device" field.
	DefaultDevice string
)

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDelete orders the results by the delete field.
func ByDelete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDelete, opts...).ToFunc()
}

// ByCreatedID orders the results by the created_id field.
func ByCreatedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedID, opts...).ToFunc()
}

// ByOrderSn orders the results by the order_sn field.
func ByOrderSn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderSn, opts...).ToFunc()
}

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByNature orders the results by the nature field.
func ByNature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNature, opts...).ToFunc()
}

// ByProductType orders the results by the product_type field.
func ByProductType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductType, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByDevice orders the results by the device field.
func ByDevice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDevice, opts...).ToFunc()
}

// ByCompletionAt orders the results by the completion_at field.
func ByCompletionAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompletionAt, opts...).ToFunc()
}

// ByCreateID orders the results by the create_id field.
func ByCreateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateID, opts...).ToFunc()
}

// ByAmountCount orders the results by amount count.
func ByAmountCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAmountStep(), opts...)
	}
}

// ByAmount orders the results by amount terms.
func ByAmount(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAmountStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByItemCount orders the results by item count.
func ByItemCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newItemStep(), opts...)
	}
}

// ByItem orders the results by item terms.
func ByItem(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByPayCount orders the results by pay count.
func ByPayCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPayStep(), opts...)
	}
}

// ByPay orders the results by pay terms.
func ByPay(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPayStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrderContentsCount orders the results by order_contents count.
func ByOrderContentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOrderContentsStep(), opts...)
	}
}

// ByOrderContents orders the results by order_contents terms.
func ByOrderContents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderContentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySalesCount orders the results by sales count.
func BySalesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSalesStep(), opts...)
	}
}

// BySales orders the results by sales terms.
func BySales(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSalesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOrderVenuesField orders the results by order_venues field.
func ByOrderVenuesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderVenuesStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrderMembersField orders the results by order_members field.
func ByOrderMembersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderMembersStep(), sql.OrderByField(field, opts...))
	}
}

// ByOrderCreatesField orders the results by order_creates field.
func ByOrderCreatesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderCreatesStep(), sql.OrderByField(field, opts...))
	}
}
func newAmountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AmountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AmountTable, AmountColumn),
	)
}
func newItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ItemTable, ItemColumn),
	)
}
func newPayStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PayInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PayTable, PayColumn),
	)
}
func newOrderContentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderContentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, OrderContentsTable, OrderContentsColumn),
	)
}
func newSalesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SalesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, SalesTable, SalesColumn),
	)
}
func newOrderVenuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderVenuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderVenuesTable, OrderVenuesColumn),
	)
}
func newOrderMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderMembersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderMembersTable, OrderMembersColumn),
	)
}
func newOrderCreatesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderCreatesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderCreatesTable, OrderCreatesColumn),
	)
}
