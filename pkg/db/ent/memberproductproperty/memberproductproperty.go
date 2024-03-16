// Code generated by ent, DO NOT EDIT.

package memberproductproperty

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memberproductproperty type in the database.
	Label = "member_product_property"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMemberProductID holds the string denoting the member_product_id field in the database.
	FieldMemberProductID = "member_product_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldCountSurplus holds the string denoting the count_surplus field in the database.
	FieldCountSurplus = "count_surplus"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeMemberProductPropertyVenues holds the string denoting the member_product_property_venues edge name in mutations.
	EdgeMemberProductPropertyVenues = "member_product_property_venues"
	// Table holds the table name of the memberproductproperty in the database.
	Table = "member_product_property"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "member_product_property"
	// OwnerInverseTable is the table name for the MemberProduct entity.
	// It exists in this package in order to avoid circular dependency with the "memberproduct" package.
	OwnerInverseTable = "member_product"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "member_product_id"
	// MemberProductPropertyVenuesTable is the table that holds the member_product_property_venues relation/edge.
	MemberProductPropertyVenuesTable = "member_product_property_venue"
	// MemberProductPropertyVenuesInverseTable is the table name for the MemberProductPropertyVenue entity.
	// It exists in this package in order to avoid circular dependency with the "memberproductpropertyvenue" package.
	MemberProductPropertyVenuesInverseTable = "member_product_property_venue"
	// MemberProductPropertyVenuesColumn is the table column denoting the member_product_property_venues relation/edge.
	MemberProductPropertyVenuesColumn = "member_product_property_id"
)

// Columns holds all SQL columns for memberproductproperty fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMemberProductID,
	FieldType,
	FieldName,
	FieldDuration,
	FieldLength,
	FieldCount,
	FieldCountSurplus,
	FieldPrice,
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
	// DefaultCount holds the default value on creation for the "count" field.
	DefaultCount int64
	// DefaultCountSurplus holds the default value on creation for the "count_surplus" field.
	DefaultCountSurplus int64
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
)

// OrderOption defines the ordering options for the MemberProductProperty queries.
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

// ByMemberProductID orders the results by the member_product_id field.
func ByMemberProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByLength orders the results by the length field.
func ByLength(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLength, opts...).ToFunc()
}

// ByCount orders the results by the count field.
func ByCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCount, opts...).ToFunc()
}

// ByCountSurplus orders the results by the count_surplus field.
func ByCountSurplus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCountSurplus, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
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

// ByMemberProductPropertyVenuesCount orders the results by member_product_property_venues count.
func ByMemberProductPropertyVenuesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMemberProductPropertyVenuesStep(), opts...)
	}
}

// ByMemberProductPropertyVenues orders the results by member_product_property_venues terms.
func ByMemberProductPropertyVenues(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberProductPropertyVenuesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}
func newMemberProductPropertyVenuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberProductPropertyVenuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MemberProductPropertyVenuesTable, MemberProductPropertyVenuesColumn),
	)
}
