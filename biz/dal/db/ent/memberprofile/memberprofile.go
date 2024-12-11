// Code generated by ent, DO NOT EDIT.

package memberprofile

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memberprofile type in the database.
	Label = "member_profile"
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
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldMobileAscription holds the string denoting the mobile_ascription field in the database.
	FieldMobileAscription = "mobile_ascription"
	// FieldFatherName holds the string denoting the father_name field in the database.
	FieldFatherName = "father_name"
	// FieldMotherName holds the string denoting the mother_name field in the database.
	FieldMotherName = "mother_name"
	// FieldGrade holds the string denoting the grade field in the database.
	FieldGrade = "grade"
	// FieldIntention holds the string denoting the intention field in the database.
	FieldIntention = "intention"
	// FieldSource holds the string denoting the source field in the database.
	FieldSource = "source"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// Table holds the table name of the memberprofile in the database.
	Table = "member_profile"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "member_profile"
	// ProfileInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	ProfileInverseTable = "member"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "member_id"
)

// Columns holds all SQL columns for memberprofile fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldMemberID,
	FieldMobileAscription,
	FieldFatherName,
	FieldMotherName,
	FieldGrade,
	FieldIntention,
	FieldSource,
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
	// DefaultMobileAscription holds the default value on creation for the "mobile_ascription" field.
	DefaultMobileAscription int64
	// DefaultGrade holds the default value on creation for the "grade" field.
	DefaultGrade int64
	// DefaultIntention holds the default value on creation for the "intention" field.
	DefaultIntention int64
	// DefaultSource holds the default value on creation for the "source" field.
	DefaultSource int64
)

// OrderOption defines the ordering options for the MemberProfile queries.
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

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByMobileAscription orders the results by the mobile_ascription field.
func ByMobileAscription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobileAscription, opts...).ToFunc()
}

// ByFatherName orders the results by the father_name field.
func ByFatherName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFatherName, opts...).ToFunc()
}

// ByMotherName orders the results by the mother_name field.
func ByMotherName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMotherName, opts...).ToFunc()
}

// ByGrade orders the results by the grade field.
func ByGrade(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGrade, opts...).ToFunc()
}

// ByIntention orders the results by the intention field.
func ByIntention(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntention, opts...).ToFunc()
}

// BySource orders the results by the source field.
func BySource(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSource, opts...).ToFunc()
}

// ByProfileField orders the results by profile field.
func ByProfileField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProfileStep(), sql.OrderByField(field, opts...))
	}
}
func newProfileStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProfileInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ProfileTable, ProfileColumn),
	)
}
