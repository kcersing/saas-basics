// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
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
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldFunctions holds the string denoting the functions field in the database.
	FieldFunctions = "functions"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldJobTime holds the string denoting the job_time field in the database.
	FieldJobTime = "job_time"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldDefaultVenueID holds the string denoting the default_venue_id field in the database.
	FieldDefaultVenueID = "default_venue_id"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeTags holds the string denoting the tags edge name in mutations.
	EdgeTags = "tags"
	// EdgeCreatedOrders holds the string denoting the created_orders edge name in mutations.
	EdgeCreatedOrders = "created_orders"
	// EdgeUserEntry holds the string denoting the user_entry edge name in mutations.
	EdgeUserEntry = "user_entry"
	// EdgeVenues holds the string denoting the venues edge name in mutations.
	EdgeVenues = "venues"
	// EdgeUserScheduling holds the string denoting the user_scheduling edge name in mutations.
	EdgeUserScheduling = "user_scheduling"
	// Table holds the table name of the user in the database.
	Table = "sys_users"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "sys_tokens"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "sys_tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "user_token"
	// TagsTable is the table that holds the tags relation/edge. The primary key declared below.
	TagsTable = "user_tags"
	// TagsInverseTable is the table name for the DictionaryDetail entity.
	// It exists in this package in order to avoid circular dependency with the "dictionarydetail" package.
	TagsInverseTable = "sys_dictionary_details"
	// CreatedOrdersTable is the table that holds the created_orders relation/edge.
	CreatedOrdersTable = "order"
	// CreatedOrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	CreatedOrdersInverseTable = "order"
	// CreatedOrdersColumn is the table column denoting the created_orders relation/edge.
	CreatedOrdersColumn = "created_id"
	// UserEntryTable is the table that holds the user_entry relation/edge.
	UserEntryTable = "entry_logs"
	// UserEntryInverseTable is the table name for the EntryLogs entity.
	// It exists in this package in order to avoid circular dependency with the "entrylogs" package.
	UserEntryInverseTable = "entry_logs"
	// UserEntryColumn is the table column denoting the user_entry relation/edge.
	UserEntryColumn = "user_id"
	// VenuesTable is the table that holds the venues relation/edge. The primary key declared below.
	VenuesTable = "user_venues"
	// VenuesInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenuesInverseTable = "venue"
	// UserSchedulingTable is the table that holds the user_scheduling relation/edge.
	UserSchedulingTable = "sys_user_scheduling"
	// UserSchedulingInverseTable is the table name for the UserScheduling entity.
	// It exists in this package in order to avoid circular dependency with the "userscheduling" package.
	UserSchedulingInverseTable = "sys_user_scheduling"
	// UserSchedulingColumn is the table column denoting the user_scheduling relation/edge.
	UserSchedulingColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldMobile,
	FieldName,
	FieldGender,
	FieldUsername,
	FieldPassword,
	FieldFunctions,
	FieldType,
	FieldJobTime,
	FieldRoleID,
	FieldDefaultVenueID,
	FieldAvatar,
	FieldDetail,
}

var (
	// TagsPrimaryKey and TagsColumn2 are the table columns denoting the
	// primary key for the tags relation (M2M).
	TagsPrimaryKey = []string{"user_id", "dictionary_detail_id"}
	// VenuesPrimaryKey and VenuesColumn2 are the table columns denoting the
	// primary key for the venues relation (M2M).
	VenuesPrimaryKey = []string{"user_id", "venue_id"}
)

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
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender int64
	// DefaultType holds the default value on creation for the "type" field.
	DefaultType int64
	// DefaultJobTime holds the default value on creation for the "job_time" field.
	DefaultJobTime int64
	// DefaultRoleID holds the default value on creation for the "role_id" field.
	DefaultRoleID int64
)

// OrderOption defines the ordering options for the User queries.
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

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByFunctions orders the results by the functions field.
func ByFunctions(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFunctions, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByJobTime orders the results by the job_time field.
func ByJobTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJobTime, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByDefaultVenueID orders the results by the default_venue_id field.
func ByDefaultVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDefaultVenueID, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByDetail orders the results by the detail field.
func ByDetail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDetail, opts...).ToFunc()
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}

// ByTagsCount orders the results by tags count.
func ByTagsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTagsStep(), opts...)
	}
}

// ByTags orders the results by tags terms.
func ByTags(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCreatedOrdersCount orders the results by created_orders count.
func ByCreatedOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCreatedOrdersStep(), opts...)
	}
}

// ByCreatedOrders orders the results by created_orders terms.
func ByCreatedOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatedOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserEntryCount orders the results by user_entry count.
func ByUserEntryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserEntryStep(), opts...)
	}
}

// ByUserEntry orders the results by user_entry terms.
func ByUserEntry(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserEntryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVenuesCount orders the results by venues count.
func ByVenuesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVenuesStep(), opts...)
	}
}

// ByVenues orders the results by venues terms.
func ByVenues(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVenuesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUserSchedulingCount orders the results by user_scheduling count.
func ByUserSchedulingCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUserSchedulingStep(), opts...)
	}
}

// ByUserScheduling orders the results by user_scheduling terms.
func ByUserScheduling(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserSchedulingStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, TokenTable, TokenColumn),
	)
}
func newTagsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
	)
}
func newCreatedOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatedOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CreatedOrdersTable, CreatedOrdersColumn),
	)
}
func newUserEntryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserEntryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserEntryTable, UserEntryColumn),
	)
}
func newVenuesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenuesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, VenuesTable, VenuesPrimaryKey...),
	)
}
func newUserSchedulingStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserSchedulingInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UserSchedulingTable, UserSchedulingColumn),
	)
}
