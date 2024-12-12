// Code generated by ent, DO NOT EDIT.

package dictionarydetail

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the dictionarydetail type in the database.
	Label = "dictionary_detail"
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
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// FieldDictionaryID holds the string denoting the dictionary_id field in the database.
	FieldDictionaryID = "dictionary_id"
	// EdgeDictionary holds the string denoting the dictionary edge name in mutations.
	EdgeDictionary = "dictionary"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the dictionarydetail in the database.
	Table = "sys_dictionary_details"
	// DictionaryTable is the table that holds the dictionary relation/edge.
	DictionaryTable = "sys_dictionary_details"
	// DictionaryInverseTable is the table name for the Dictionary entity.
	// It exists in this package in order to avoid circular dependency with the "dictionary" package.
	DictionaryInverseTable = "sys_dictionaries"
	// DictionaryColumn is the table column denoting the dictionary relation/edge.
	DictionaryColumn = "dictionary_id"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_tag"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "sys_users"
	// ProductsTable is the table that holds the products relation/edge. The primary key declared below.
	ProductsTable = "user_tag"
	// ProductsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ProductsInverseTable = "sys_users"
)

// Columns holds all SQL columns for dictionarydetail fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldTitle,
	FieldKey,
	FieldValue,
	FieldDictionaryID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "sys_dictionary_details"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"product_tag",
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "dictionary_detail_id"}
	// ProductsPrimaryKey and ProductsColumn2 are the table columns denoting the
	// primary key for the products relation (M2M).
	ProductsPrimaryKey = []string{"user_id", "dictionary_detail_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
)

// OrderOption defines the ordering options for the DictionaryDetail queries.
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

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByKey orders the results by the key field.
func ByKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldKey, opts...).ToFunc()
}

// ByValue orders the results by the value field.
func ByValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldValue, opts...).ToFunc()
}

// ByDictionaryID orders the results by the dictionary_id field.
func ByDictionaryID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDictionaryID, opts...).ToFunc()
}

// ByDictionaryField orders the results by dictionary field.
func ByDictionaryField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDictionaryStep(), sql.OrderByField(field, opts...))
	}
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newDictionaryStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DictionaryInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DictionaryTable, DictionaryColumn),
	)
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
	)
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProductsTable, ProductsPrimaryKey...),
	)
}
