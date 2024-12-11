// Code generated by ent, DO NOT EDIT.

package memberdetails

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the memberdetails type in the database.
	Label = "member_details"
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
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWecom holds the string denoting the wecom field in the database.
	FieldWecom = "wecom"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldMoneySum holds the string denoting the money_sum field in the database.
	FieldMoneySum = "money_sum"
	// FieldProductID holds the string denoting the product_id field in the database.
	FieldProductID = "product_id"
	// FieldProductName holds the string denoting the product_name field in the database.
	FieldProductName = "product_name"
	// FieldProductVenue holds the string denoting the product_venue field in the database.
	FieldProductVenue = "product_venue"
	// FieldProductVenueName holds the string denoting the product_venue_name field in the database.
	FieldProductVenueName = "product_venue_name"
	// FieldEntrySum holds the string denoting the entry_sum field in the database.
	FieldEntrySum = "entry_sum"
	// FieldEntryLastTime holds the string denoting the entry_last_time field in the database.
	FieldEntryLastTime = "entry_last_time"
	// FieldEntryDeadlineTime holds the string denoting the entry_deadline_time field in the database.
	FieldEntryDeadlineTime = "entry_deadline_time"
	// FieldClassLastTime holds the string denoting the class_last_time field in the database.
	FieldClassLastTime = "class_last_time"
	// FieldRelationUID holds the string denoting the relation_uid field in the database.
	FieldRelationUID = "relation_uid"
	// FieldRelationUname holds the string denoting the relation_uname field in the database.
	FieldRelationUname = "relation_uname"
	// FieldRelationMid holds the string denoting the relation_mid field in the database.
	FieldRelationMid = "relation_mid"
	// FieldRelationMame holds the string denoting the relation_mame field in the database.
	FieldRelationMame = "relation_mame"
	// EdgeInfo holds the string denoting the info edge name in mutations.
	EdgeInfo = "info"
	// Table holds the table name of the memberdetails in the database.
	Table = "member_details"
	// InfoTable is the table that holds the info relation/edge.
	InfoTable = "member_details"
	// InfoInverseTable is the table name for the Member entity.
	// It exists in this package in order to avoid circular dependency with the "member" package.
	InfoInverseTable = "member"
	// InfoColumn is the table column denoting the info relation/edge.
	InfoColumn = "member_id"
)

// Columns holds all SQL columns for memberdetails fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldMemberID,
	FieldEmail,
	FieldWecom,
	FieldGender,
	FieldBirthday,
	FieldMoneySum,
	FieldProductID,
	FieldProductName,
	FieldProductVenue,
	FieldProductVenueName,
	FieldEntrySum,
	FieldEntryLastTime,
	FieldEntryDeadlineTime,
	FieldClassLastTime,
	FieldRelationUID,
	FieldRelationUname,
	FieldRelationMid,
	FieldRelationMame,
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
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender int64
	// DefaultMoneySum holds the default value on creation for the "money_sum" field.
	DefaultMoneySum float64
	// DefaultProductID holds the default value on creation for the "product_id" field.
	DefaultProductID int64
	// DefaultProductVenue holds the default value on creation for the "product_venue" field.
	DefaultProductVenue int64
	// DefaultEntrySum holds the default value on creation for the "entry_sum" field.
	DefaultEntrySum int64
	// DefaultRelationUID holds the default value on creation for the "relation_uid" field.
	DefaultRelationUID int64
	// DefaultRelationMid holds the default value on creation for the "relation_mid" field.
	DefaultRelationMid int64
)

// OrderOption defines the ordering options for the MemberDetails queries.
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

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByWecom orders the results by the wecom field.
func ByWecom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWecom, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByMoneySum orders the results by the money_sum field.
func ByMoneySum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMoneySum, opts...).ToFunc()
}

// ByProductID orders the results by the product_id field.
func ByProductID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductID, opts...).ToFunc()
}

// ByProductName orders the results by the product_name field.
func ByProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductName, opts...).ToFunc()
}

// ByProductVenue orders the results by the product_venue field.
func ByProductVenue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductVenue, opts...).ToFunc()
}

// ByProductVenueName orders the results by the product_venue_name field.
func ByProductVenueName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductVenueName, opts...).ToFunc()
}

// ByEntrySum orders the results by the entry_sum field.
func ByEntrySum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntrySum, opts...).ToFunc()
}

// ByEntryLastTime orders the results by the entry_last_time field.
func ByEntryLastTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntryLastTime, opts...).ToFunc()
}

// ByEntryDeadlineTime orders the results by the entry_deadline_time field.
func ByEntryDeadlineTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEntryDeadlineTime, opts...).ToFunc()
}

// ByClassLastTime orders the results by the class_last_time field.
func ByClassLastTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClassLastTime, opts...).ToFunc()
}

// ByRelationUID orders the results by the relation_uid field.
func ByRelationUID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelationUID, opts...).ToFunc()
}

// ByRelationUname orders the results by the relation_uname field.
func ByRelationUname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelationUname, opts...).ToFunc()
}

// ByRelationMid orders the results by the relation_mid field.
func ByRelationMid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelationMid, opts...).ToFunc()
}

// ByRelationMame orders the results by the relation_mame field.
func ByRelationMame(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelationMame, opts...).ToFunc()
}

// ByInfoField orders the results by info field.
func ByInfoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInfoStep(), sql.OrderByField(field, opts...))
	}
}
func newInfoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InfoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InfoTable, InfoColumn),
	)
}
