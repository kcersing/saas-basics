// Code generated by ent, DO NOT EDIT.

package memberproduct

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memberproduct type in the database.
	Label = "member_product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldSn holds the string denoting the sn field in the database.
	FieldSn = "sn"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldValidityAt holds the string denoting the validity_at field in the database.
	FieldValidityAt = "validity_at"
	// FieldCancelAt holds the string denoting the cancel_at field in the database.
	FieldCancelAt = "cancel_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeMemberProductPropertys holds the string denoting the member_product_propertys edge name in mutations.
	EdgeMemberProductPropertys = "member_product_propertys"
	// Table holds the table name of the memberproduct in the database.
	Table = "member_product"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "member_product"
	// OwnerInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	OwnerInverseTable = "member"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "member_id"
	// MemberProductPropertysTable is the table that holds the member_product_propertys relation/edge.
	MemberProductPropertysTable = "member_product_property"
	// MemberProductPropertysInverseTable is the table name for the MemberProductProperty entity.
	// It exists in this package in order to avoid circular dependency with the "memberproductproperty" package.
	MemberProductPropertysInverseTable = "member_product_property"
	// MemberProductPropertysColumn is the table column denoting the member_product_propertys relation/edge.
	MemberProductPropertysColumn = "member_product_id"
)

// Columns holds all SQL columns for memberproduct fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldSn,
	FieldType,
	FieldMemberID,
	FieldProductID,
	FieldName,
	FieldPrice,
	FieldValidityAt,
	FieldCancelAt,
	FieldStatus,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
)

// OrderOption defines the ordering options for the MemberProduct queries.
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

// BySn orders the results by the sn field.
func BySn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSn, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByValidityAt orders the results by the validity_at field.
func ByValidityAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidityAt, opts...).ToFunc()
}

// ByCancelAt orders the results by the cancel_at field.
func ByCancelAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCancelAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}

// ByMemberProductPropertysCount orders the results by member_product_propertys count.
func ByMemberProductPropertysCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMemberProductPropertysStep(), opts...)
	}
}

// ByMemberProductPropertys orders the results by member_product_propertys terms.
func ByMemberProductPropertys(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberProductPropertysStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newMemberProductPropertysStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberProductPropertysInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MemberProductPropertysTable, MemberProductPropertysColumn),
	)
}
