// Code generated by ent, DO NOT EDIT.

package schedulemember

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the schedulemember type in the database.
	Label = "schedule_member"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldScheduleID holds the string denoting the schedule_id field in the database.
	FieldScheduleID = "schedule_id"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldMemberProductID holds the string denoting the member_product_id field in the database.
	FieldMemberProductID = "member_product_id"
	// FieldMemberProductPropertyID holds the string denoting the member_product_property_id field in the database.
	FieldMemberProductPropertyID = "member_product_property_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldSignStartTime holds the string denoting the sign_start_time field in the database.
	FieldSignStartTime = "sign_start_time"
	// FieldSignEndTime holds the string denoting the sign_end_time field in the database.
	FieldSignEndTime = "sign_end_time"
	// FieldMemberName holds the string denoting the member_name field in the database.
	FieldMemberName = "member_name"
	// FieldMemberProductName holds the string denoting the member_product_name field in the database.
	FieldMemberProductName = "member_product_name"
	// FieldMemberProductPropertyName holds the string denoting the member_product_property_name field in the database.
	FieldMemberProductPropertyName = "member_product_property_name"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// EdgeSchedule holds the string denoting the schedule edge name in mutations.
	EdgeSchedule = "schedule"
	// Table holds the table name of the schedulemember in the database.
	Table = "schedule_member"
	// ScheduleTable is the table that holds the schedule relation/edge.
	ScheduleTable = "schedule_member"
	// ScheduleInverseTable is the table name for the Schedule entity.
	// It exists in this package in order to avoid circular dependency with the "schedule" package.
	ScheduleInverseTable = "schedule"
	// ScheduleColumn is the table column denoting the schedule relation/edge.
	ScheduleColumn = "schedule_id"
)

// Columns holds all SQL columns for schedulemember fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldVenueID,
	FieldScheduleID,
	FieldMemberID,
	FieldMemberProductID,
	FieldMemberProductPropertyID,
	FieldType,
	FieldStartTime,
	FieldEndTime,
	FieldSignStartTime,
	FieldSignEndTime,
	FieldMemberName,
	FieldMemberProductName,
	FieldMemberProductPropertyName,
	FieldRemark,
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
	// DefaultStartTime holds the default value on creation for the "start_time" field.
	DefaultStartTime func() time.Time
	// DefaultEndTime holds the default value on creation for the "end_time" field.
	DefaultEndTime func() time.Time
	// DefaultSignStartTime holds the default value on creation for the "sign_start_time" field.
	DefaultSignStartTime func() time.Time
	// DefaultSignEndTime holds the default value on creation for the "sign_end_time" field.
	DefaultSignEndTime func() time.Time
)

// OrderOption defines the ordering options for the ScheduleMember queries.
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

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByScheduleID orders the results by the schedule_id field.
func ByScheduleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScheduleID, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByMemberProductID orders the results by the member_product_id field.
func ByMemberProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductID, opts...).ToFunc()
}

// ByMemberProductPropertyID orders the results by the member_product_property_id field.
func ByMemberProductPropertyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductPropertyID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// BySignStartTime orders the results by the sign_start_time field.
func BySignStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignStartTime, opts...).ToFunc()
}

// BySignEndTime orders the results by the sign_end_time field.
func BySignEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignEndTime, opts...).ToFunc()
}

// ByMemberName orders the results by the member_name field.
func ByMemberName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberName, opts...).ToFunc()
}

// ByMemberProductName orders the results by the member_product_name field.
func ByMemberProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductName, opts...).ToFunc()
}

// ByMemberProductPropertyName orders the results by the member_product_property_name field.
func ByMemberProductPropertyName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberProductPropertyName, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByScheduleField orders the results by schedule field.
func ByScheduleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newScheduleStep(), sql.OrderByField(field, opts...))
	}
}
func newScheduleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ScheduleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ScheduleTable, ScheduleColumn),
	)
}
