// Code generated by ent, DO NOT EDIT.

package orderpay

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orderpay type in the database.
	Label = "order_pay"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldRemission holds the string denoting the remission field in the database.
	FieldRemission = "remission"
	// FieldPay holds the string denoting the pay field in the database.
	FieldPay = "pay"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldPayWay holds the string denoting the pay_way field in the database.
	FieldPayWay = "pay_way"
	// FieldCreateID holds the string denoting the create_id field in the database.
	FieldCreateID = "create_id"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// Table holds the table name of the orderpay in the database.
	Table = "order_pay"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "order_pay"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "order"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_id"
)

// Columns holds all SQL columns for orderpay fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldOrderID,
	FieldRemission,
	FieldPay,
	FieldNote,
	FieldPayWay,
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
)

// OrderOption defines the ordering options for the OrderPay queries.
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

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByRemission orders the results by the remission field.
func ByRemission(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemission, opts...).ToFunc()
}

// ByPay orders the results by the pay field.
func ByPay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPay, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// ByPayWay orders the results by the pay_way field.
func ByPayWay(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPayWay, opts...).ToFunc()
}

// ByCreateID orders the results by the create_id field.
func ByCreateID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreateID, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
	)
}
