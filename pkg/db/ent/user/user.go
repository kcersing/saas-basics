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
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldSideMode holds the string denoting the side_mode field in the database.
	FieldSideMode = "side_mode"
	// FieldBaseColor holds the string denoting the base_color field in the database.
	FieldBaseColor = "base_color"
	// FieldActiveColor holds the string denoting the active_color field in the database.
	FieldActiveColor = "active_color"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldWecom holds the string denoting the wecom field in the database.
	FieldWecom = "wecom"
	// FieldJob holds the string denoting the job field in the database.
	FieldJob = "job"
	// FieldOrganization holds the string denoting the organization field in the database.
	FieldOrganization = "organization"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// Table holds the table name of the user in the database.
	Table = "sys_users"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "sys_tokens"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "sys_tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "user_token"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldUsername,
	FieldPassword,
	FieldNickname,
	FieldSideMode,
	FieldBaseColor,
	FieldActiveColor,
	FieldRoleID,
	FieldMobile,
	FieldEmail,
	FieldWecom,
	FieldJob,
	FieldOrganization,
	FieldAvatar,
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
	// DefaultSideMode holds the default value on creation for the "side_mode" field.
	DefaultSideMode string
	// DefaultBaseColor holds the default value on creation for the "base_color" field.
	DefaultBaseColor string
	// DefaultActiveColor holds the default value on creation for the "active_color" field.
	DefaultActiveColor string
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

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByNickname orders the results by the nickname field.
func ByNickname(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNickname, opts...).ToFunc()
}

// BySideMode orders the results by the side_mode field.
func BySideMode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSideMode, opts...).ToFunc()
}

// ByBaseColor orders the results by the base_color field.
func ByBaseColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBaseColor, opts...).ToFunc()
}

// ByActiveColor orders the results by the active_color field.
func ByActiveColor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldActiveColor, opts...).ToFunc()
}

// ByRoleID orders the results by the role_id field.
func ByRoleID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRoleID, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByWecom orders the results by the wecom field.
func ByWecom(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldWecom, opts...).ToFunc()
}

// ByJob orders the results by the job field.
func ByJob(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldJob, opts...).ToFunc()
}

// ByOrganization orders the results by the organization field.
func ByOrganization(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOrganization, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByTokenField orders the results by token field.
func ByTokenField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTokenStep(), sql.OrderByField(field, opts...))
	}
}
func newTokenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TokenInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, TokenTable, TokenColumn),
	)
}
