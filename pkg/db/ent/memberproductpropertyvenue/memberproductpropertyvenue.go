// Code generated by ent, DO NOT EDIT.

package memberproductpropertyvenue

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memberproductpropertyvenue type in the database.
	Label = "member_product_property_venue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldMemberProductPropertyID holds the string denoting the member_product_property_id field in the database.
	FieldMemberProductPropertyID = "member_product_property_id"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the memberproductpropertyvenue in the database.
	Table = "member_product_property_venue"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "member_product_property_venue"
	// OwnerInverseTable is the table name for the MemberProductProperty entity.
	// It exists in this package in order to avoid circular dependency with the "memberproductproperty" package.
	OwnerInverseTable = "member_product_property"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "member_product_property_id"
)

// Columns holds all SQL columns for memberproductpropertyvenue fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldVenueID,
	FieldMemberProductPropertyID,
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

// OrderOption defines the ordering options for the MemberProductPropertyVenue queries.
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

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByMemberProductPropertyID orders the results by the member_product_property_id field.
func ByMemberProductPropertyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductPropertyID, opts...).ToFunc()
}

// ByOwnerField orders the results by owner field.
func ByOwnerField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOwnerStep(), sql.OrderByField(field, opts...))
	}
}
func newOwnerStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OwnerInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
	)
}