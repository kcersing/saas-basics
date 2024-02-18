// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/rs/xid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldIntroduce holds the string denoting the introduce field in the database.
	FieldIntroduce = "introduce"
	// Table holds the table name of the user in the database.
	Table = "user"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAccountID,
	FieldUsername,
	FieldPassword,
	FieldMobile,
	FieldGender,
	FieldAge,
	FieldIntroduce,
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
	// DefaultAccountID holds the default value on creation for the "account_id" field.
	DefaultAccountID func() xid.ID
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender string
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByGender orders the results by the gender field.
func ByGender(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGender, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByIntroduce orders the results by the introduce field.
func ByIntroduce(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIntroduce, opts...).ToFunc()
}
