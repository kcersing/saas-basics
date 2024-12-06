// Code generated by ent, DO NOT EDIT.

package messages

import (
	"time"

	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the messages type in the database.
	Label = "messages"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeleteAt holds the string denoting the delete_at field in the database.
	FieldDeleteAt = "delete_at"
	// FieldCreatedID holds the string denoting the created_id field in the database.
	FieldCreatedID = "created_id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldToUserID holds the string denoting the to_user_id field in the database.
	FieldToUserID = "to_user_id"
	// FieldFromUserID holds the string denoting the from_user_id field in the database.
	FieldFromUserID = "from_user_id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// Table holds the table name of the messages in the database.
	Table = "messages"
)

// Columns holds all SQL columns for messages fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeleteAt,
	FieldCreatedID,
	FieldType,
	FieldToUserID,
	FieldFromUserID,
	FieldContent,
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
	// DefaultCreatedID holds the default value on creation for the "created_id" field.
	DefaultCreatedID int64
)

// OrderOption defines the ordering options for the Messages queries.
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

// ByDeleteAt orders the results by the delete_at field.
func ByDeleteAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeleteAt, opts...).ToFunc()
}

// ByCreatedID orders the results by the created_id field.
func ByCreatedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedID, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByToUserID orders the results by the to_user_id field.
func ByToUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldToUserID, opts...).ToFunc()
}

// ByFromUserID orders the results by the from_user_id field.
func ByFromUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFromUserID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}
