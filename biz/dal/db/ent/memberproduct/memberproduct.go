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
	// FieldDelete holds the string denoting the delete field in the database.
	FieldDelete = "delete"
	// FieldCreatedID holds the string denoting the created_id field in the database.
	FieldCreatedID = "created_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldSn holds the string denoting the sn field in the database.
	FieldSn = "sn"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldSubType holds the string denoting the sub_type field in the database.
	FieldSubType = "sub_type"
	// FieldMemberID holds the string denoting the member_id field in the database.
	FieldMemberID = "member_id"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldOrderID holds the string denoting the order_id field in the database.
	FieldOrderID = "order_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldFee holds the string denoting the fee field in the database.
	FieldFee = "fee"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldLength holds the string denoting the length field in the database.
	FieldLength = "length"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldNumberSurplus holds the string denoting the number_surplus field in the database.
	FieldNumberSurplus = "number_surplus"
	// FieldIsCourse holds the string denoting the is_course field in the database.
	FieldIsCourse = "is_course"
	// FieldDeadline holds the string denoting the deadline field in the database.
	FieldDeadline = "deadline"
	// FieldValidityAt holds the string denoting the validity_at field in the database.
	FieldValidityAt = "validity_at"
	// FieldCancelAt holds the string denoting the cancel_at field in the database.
	FieldCancelAt = "cancel_at"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeMemberProductEntry holds the string denoting the member_product_entry edge name in mutations.
	EdgeMemberProductEntry = "member_product_entry"
	// EdgeMemberProductContents holds the string denoting the member_product_contents edge name in mutations.
	EdgeMemberProductContents = "member_product_contents"
	// EdgeMemberCourses holds the string denoting the membercourses edge name in mutations.
	EdgeMemberCourses = "memberCourses"
	// EdgeMemberLessons holds the string denoting the memberlessons edge name in mutations.
	EdgeMemberLessons = "memberLessons"
	// Table holds the table name of the memberproduct in the database.
	Table = "member_product"
	// MembersTable is the table that holds the members relation/edge.
	MembersTable = "member_product"
	// MembersInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	MembersInverseTable = "member"
	// MembersColumn is the table column denoting the members relation/edge.
	MembersColumn = "member_id"
	// MemberProductEntryTable is the table that holds the member_product_entry relation/edge.
	MemberProductEntryTable = "entry_logs"
	// MemberProductEntryInverseTable is the table name for the EntryLogs entity.
	// It exists in this package in order to avoid circular dependency with the "entrylogs" package.
	MemberProductEntryInverseTable = "entry_logs"
	// MemberProductEntryColumn is the table column denoting the member_product_entry relation/edge.
	MemberProductEntryColumn = "member_product_id"
	// MemberProductContentsTable is the table that holds the member_product_contents relation/edge.
	MemberProductContentsTable = "member_contract"
	// MemberProductContentsInverseTable is the table name for the MemberContract entity.
	// It exists in this package in order to avoid circular dependency with the "membercontract" package.
	MemberProductContentsInverseTable = "member_contract"
	// MemberProductContentsColumn is the table column denoting the member_product_contents relation/edge.
	MemberProductContentsColumn = "member_product_id"
	// MemberCoursesTable is the table that holds the memberCourses relation/edge.
	MemberCoursesTable = "member_product_courses"
	// MemberCoursesInverseTable is the table name for the MemberProductCourses entity.
	// It exists in this package in order to avoid circular dependency with the "memberproductcourses" package.
	MemberCoursesInverseTable = "member_product_courses"
	// MemberCoursesColumn is the table column denoting the memberCourses relation/edge.
	MemberCoursesColumn = "member_product_id"
	// MemberLessonsTable is the table that holds the memberLessons relation/edge.
	MemberLessonsTable = "member_product_courses"
	// MemberLessonsInverseTable is the table name for the MemberProductCourses entity.
	// It exists in this package in order to avoid circular dependency with the "memberproductcourses" package.
	MemberLessonsInverseTable = "member_product_courses"
	// MemberLessonsColumn is the table column denoting the memberLessons relation/edge.
	MemberLessonsColumn = "member_product_id"
)

// Columns holds all SQL columns for memberproduct fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldSn,
	FieldType,
	FieldSubType,
	FieldMemberID,
	FieldProductID,
	FieldVenueID,
	FieldOrderID,
	FieldName,
	FieldPrice,
	FieldFee,
	FieldDuration,
	FieldLength,
	FieldNumber,
	FieldNumberSurplus,
	FieldIsCourse,
	FieldDeadline,
	FieldValidityAt,
	FieldCancelAt,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int64
	// DefaultSubType holds the default value on creation for the "sub_type" field.
	DefaultSubType string
	// DefaultNumber holds the default value on creation for the "number" field.
	DefaultNumber int64
	// DefaultNumberSurplus holds the default value on creation for the "number_surplus" field.
	DefaultNumberSurplus int64
	// DefaultIsCourse holds the default value on creation for the "is_course" field.
	DefaultIsCourse int64
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

// BySn orders the results by the sn field.
func BySn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSn, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// BySubType orders the results by the sub_type field.
func BySubType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubType, opts...).ToFunc()
}

// ByMemberID orders the results by the member_id field.
func ByMemberID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMemberID, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByOrderID orders the results by the order_id field.
func ByOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrderID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByFee orders the results by the fee field.
func ByFee(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFee, opts...).ToFunc()
}

// ByDuration orders the results by the duration field.
func ByDuration(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDuration, opts...).ToFunc()
}

// ByLength orders the results by the length field.
func ByLength(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLength, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByNumberSurplus orders the results by the number_surplus field.
func ByNumberSurplus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumberSurplus, opts...).ToFunc()
}

// ByIsCourse orders the results by the is_course field.
func ByIsCourse(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCourse, opts...).ToFunc()
}

// ByDeadline orders the results by the deadline field.
func ByDeadline(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeadline, opts...).ToFunc()
}

// ByValidityAt orders the results by the validity_at field.
func ByValidityAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValidityAt, opts...).ToFunc()
}

// ByCancelAt orders the results by the cancel_at field.
func ByCancelAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCancelAt, opts...).ToFunc()
}

// ByMembersField orders the results by members field.
func ByMembersField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembersStep(), sql.OrderByField(field, opts...))
	}
}

// ByMemberProductEntryCount orders the results by member_product_entry count.
func ByMemberProductEntryCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMemberProductEntryStep(), opts...)
	}
}

// ByMemberProductEntry orders the results by member_product_entry terms.
func ByMemberProductEntry(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberProductEntryStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMemberProductContentsCount orders the results by member_product_contents count.
func ByMemberProductContentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMemberProductContentsStep(), opts...)
	}
}

// ByMemberProductContents orders the results by member_product_contents terms.
func ByMemberProductContents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberProductContentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMemberCoursesCount orders the results by memberCourses count.
func ByMemberCoursesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMemberCoursesStep(), opts...)
	}
}

// ByMemberCourses orders the results by memberCourses terms.
func ByMemberCourses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberCoursesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMemberLessonsCount orders the results by memberLessons count.
func ByMemberLessonsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMemberLessonsStep(), opts...)
	}
}

// ByMemberLessons orders the results by memberLessons terms.
func ByMemberLessons(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMemberLessonsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newMembersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, MembersTable, MembersColumn),
	)
}
func newMemberProductEntryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberProductEntryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MemberProductEntryTable, MemberProductEntryColumn),
	)
}
func newMemberProductContentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberProductContentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MemberProductContentsTable, MemberProductContentsColumn),
	)
}
func newMemberCoursesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberCoursesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MemberCoursesTable, MemberCoursesColumn),
	)
}
func newMemberLessonsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MemberLessonsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MemberLessonsTable, MemberLessonsColumn),
	)
}
