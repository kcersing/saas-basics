// Code generated by ent, DO NOT EDIT.

package venuesms

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the venuesms type in the database.
	Label = "venue_sms"
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
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldNoticeCount holds the string denoting the notice_count field in the database.
	FieldNoticeCount = "notice_count"
	// FieldUsedNotice holds the string denoting the used_notice field in the database.
	FieldUsedNotice = "used_notice"
	// EdgeVenue holds the string denoting the venue edge name in mutations.
	EdgeVenue = "venue"
	// Table holds the table name of the venuesms in the database.
	Table = "sys_venue_sms"
	// VenueTable is the table that holds the venue relation/edge.
	VenueTable = "sys_venue_sms"
	// VenueInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenueInverseTable = "venue"
	// VenueColumn is the table column denoting the venue relation/edge.
	VenueColumn = "venue_id"
)

// Columns holds all SQL columns for venuesms fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldVenueID,
	FieldNoticeCount,
	FieldUsedNotice,
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
	// DefaultVenueID holds the default value on creation for the "venue_id" field.
	DefaultVenueID int64
	// DefaultNoticeCount holds the default value on creation for the "notice_count" field.
	DefaultNoticeCount int64
	// DefaultUsedNotice holds the default value on creation for the "used_notice" field.
	DefaultUsedNotice int64
)

// OrderOption defines the ordering options for the VenueSms queries.
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

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByNoticeCount orders the results by the notice_count field.
func ByNoticeCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNoticeCount, opts...).ToFunc()
}

// ByUsedNotice orders the results by the used_notice field.
func ByUsedNotice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsedNotice, opts...).ToFunc()
}

// ByVenueField orders the results by venue field.
func ByVenueField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVenueStep(), sql.OrderByField(field, opts...))
	}
}
func newVenueStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenueInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
	)
}
