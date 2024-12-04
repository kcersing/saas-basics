// Code generated by ent, DO NOT EDIT.

package entrylogs

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the entrylogs type in the database.
	Label = "entry_logs"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldMemberProductID holds the string denoting the member_product_id field in the database.
	FieldMemberProductID = "member_product_id"
	// FieldMemberPropertyID holds the string denoting the member_property_id field in the database.
	FieldMemberPropertyID = "member_property_id"
	// FieldEntryTime holds the string denoting the entry_time field in the database.
	FieldEntryTime = "entry_time"
	// FieldLeavingTime holds the string denoting the leaving_time field in the database.
	FieldLeavingTime = "leaving_time"
	// EdgeVenues holds the string denoting the venues edge name in mutations.
	EdgeVenues = "venues"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the entrylogs in the database.
	Table = "entry_logs"
	// VenuesTable is the table that holds the venues relation/edge.
	VenuesTable = "entry_logs"
	// VenuesInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenuesInverseTable = "venue"
	// VenuesColumn is the table column denoting the venues relation/edge.
	VenuesColumn = "venue_id"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "entry_logs"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "member"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "member_id"
	// UsersTable is the table that holds the users relation/edge.
	UsersTable = "entry_logs"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "sys_users"
	// UsersColumn is the table column denoting the users relation/edge.
	UsersColumn = "user_id"
)

// Columns holds all SQL columns for entrylogs fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldMemberID,
	FieldUserID,
	FieldVenueID,
	FieldMemberProductID,
	FieldMemberPropertyID,
	FieldEntryTime,
	FieldLeavingTime,
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
	// DefaultMemberID holds the default value on creation for the "member_id" field.
	DefaultMemberID int64
	// DefaultUserID holds the default value on creation for the "user_id" field.
	DefaultUserID int64
)

// OrderOption defines the ordering options for the EntryLogs queries.
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

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByMemberProductID orders the results by the member_product_id field.
func ByMemberProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductID, opts...).ToFunc()
}

// ByMemberPropertyID orders the results by the member_property_id field.
func ByMemberPropertyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberPropertyID, opts...).ToFunc()
}

// ByEntryTime orders the results by the entry_time field.
func ByEntryTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntryTime, opts...).ToFunc()
}

// ByLeavingTime orders the results by the leaving_time field.
func ByLeavingTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLeavingTime, opts...).ToFunc()
}

// ByVenuesField orders the results by venues field.
func ByVenuesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVenuesStep(), sql.OrderByField(field, opts...))
	}
}

// ByMembersField orders the results by members field.
func ByMembersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembersStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersField orders the results by users field.
func ByUsersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), sql.OrderByField(field, opts...))
	}
}
func newVenuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VenuesTable, VenuesColumn),
	)
}
func newMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MembersTable, MembersColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
	)
}
