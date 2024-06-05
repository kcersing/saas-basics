// Code generated by ent, DO NOT EDIT.

package schedule

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the schedule type in the database.
	Label = "schedule"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldPropertyID holds the string denoting the property_id field in the database.
	FieldPropertyID = "property_id"
	// FieldPlaceID holds the string denoting the place_id field in the database.
	FieldPlaceID = "place_id"
	// FieldNum holds the string denoting the num field in the database.
	FieldNum = "num"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldStartTime holds the string denoting the start_time field in the database.
	FieldStartTime = "start_time"
	// FieldEndTime holds the string denoting the end_time field in the database.
	FieldEndTime = "end_time"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldRemark holds the string denoting the remark field in the database.
	FieldRemark = "remark"
	// FieldVenueName holds the string denoting the venue_name field in the database.
	FieldVenueName = "venue_name"
	// FieldPlaceName holds the string denoting the place_name field in the database.
	FieldPlaceName = "place_name"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeCoachs holds the string denoting the coachs edge name in mutations.
	EdgeCoachs = "coachs"
	// Table holds the table name of the schedule in the database.
	Table = "schedule"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "schedule_member"
	// MembersInverseTable is the table name for the ScheduleMember entity.
	// It exists in this package in order to avoid circular dependency with the "schedulemember" package.
	MembersInverseTable = "schedule_member"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "schedule_id"
	// CoachsTable is the table that holds the coachs relation/edge.
	CoachsTable = "schedule_coach"
	// CoachsInverseTable is the table name for the ScheduleCoach entity.
	// It exists in this package in order to avoid circular dependency with the "schedulecoach" package.
	CoachsInverseTable = "schedule_coach"
	// CoachsColumn is the table column denoting the coachs relation/edge.
	CoachsColumn = "schedule_id"
)

// Columns holds all SQL columns for schedule fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldType,
	FieldName,
	FieldVenueID,
	FieldPropertyID,
	FieldPlaceID,
	FieldNum,
	FieldDate,
	FieldStartTime,
	FieldEndTime,
	FieldPrice,
	FieldRemark,
	FieldVenueName,
	FieldPlaceName,
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
	// DefaultPrice holds the default value on creation for the "price" field.
	DefaultPrice float64
)

// OrderOption defines the ordering options for the Schedule queries.
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

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByPropertyID orders the results by the property_id field.
func ByPropertyID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPropertyID, opts...).ToFunc()
}

// ByPlaceID orders the results by the place_id field.
func ByPlaceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlaceID, opts...).ToFunc()
}

// ByNum orders the results by the num field.
func ByNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNum, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByStartTime orders the results by the start_time field.
func ByStartTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartTime, opts...).ToFunc()
}

// ByEndTime orders the results by the end_time field.
func ByEndTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndTime, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByRemark orders the results by the remark field.
func ByRemark(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRemark, opts...).ToFunc()
}

// ByVenueName orders the results by the venue_name field.
func ByVenueName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueName, opts...).ToFunc()
}

// ByPlaceName orders the results by the place_name field.
func ByPlaceName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlaceName, opts...).ToFunc()
}

// ByMembersCount orders the results by members count.
func ByMembersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembersStep(), opts...)
	}
}

// ByMembers orders the results by members terms.
func ByMembers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCoachsCount orders the results by coachs count.
func ByCoachsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCoachsStep(), opts...)
	}
}

// ByCoachs orders the results by coachs terms.
func ByCoachs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCoachsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MembersTable, MembersColumn),
	)
}
func newCoachsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CoachsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CoachsTable, CoachsColumn),
	)
}
