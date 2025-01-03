// Code generated by ent, DO NOT EDIT.

package venueplace

import (
	"saas/idl_gen/model/base"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the venueplace type in the database.
	Label = "venue_place"
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
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldClassify holds the string denoting the classify field in the database.
	FieldClassify = "classify"
	// FieldPic holds the string denoting the pic field in the database.
	FieldPic = "pic"
	// FieldVenueID holds the string denoting the venue_id field in the database.
	FieldVenueID = "venue_id"
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldIsShow holds the string denoting the is_show field in the database.
	FieldIsShow = "is_show"
	// FieldIsAccessible holds the string denoting the is_accessible field in the database.
	FieldIsAccessible = "is_accessible"
	// FieldIsBooking holds the string denoting the is_booking field in the database.
	FieldIsBooking = "is_booking"
	// FieldInformation holds the string denoting the information field in the database.
	FieldInformation = "information"
	// FieldSeat holds the string denoting the seat field in the database.
	FieldSeat = "seat"
	// EdgeVenue holds the string denoting the venue edge name in mutations.
	EdgeVenue = "venue"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// Table holds the table name of the venueplace in the database.
	Table = "venue_place"
	// VenueTable is the table that holds the venue relation/edge.
	VenueTable = "venue_place"
	// VenueInverseTable is the table name for the Venue entity.
	// It exists in this package in order to avoid circular dependency with the "venue" package.
	VenueInverseTable = "venue"
	// VenueColumn is the table column denoting the venue relation/edge.
	VenueColumn = "venue_id"
	// ProductsTable is the table that holds the products relation/edge. The primary key declared below.
	ProductsTable = "venue_place_products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "product"
)

// Columns holds all SQL columns for venueplace fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDelete,
	FieldCreatedID,
	FieldStatus,
	FieldName,
	FieldClassify,
	FieldPic,
	FieldVenueID,
	FieldNumber,
	FieldType,
	FieldIsShow,
	FieldIsAccessible,
	FieldIsBooking,
	FieldInformation,
	FieldSeat,
}

var (
	// ProductsPrimaryKey and ProductsColumn2 are the table columns denoting the
	// primary key for the products relation (M2M).
	ProductsPrimaryKey = []string{"venue_place_id", "product_id"}
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
	// DefaultIsShow holds the default value on creation for the "is_show" field.
	DefaultIsShow int64
	// DefaultIsAccessible holds the default value on creation for the "is_accessible" field.
	DefaultIsAccessible int64
	// DefaultIsBooking holds the default value on creation for the "is_booking" field.
	DefaultIsBooking int64
	// DefaultSeat holds the default value on creation for the "seat" field.
	DefaultSeat [][]*base.Seat
)

// OrderOption defines the ordering options for the VenuePlace queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByClassify orders the results by the classify field.
func ByClassify(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldClassify, opts...).ToFunc()
}

// ByPic orders the results by the pic field.
func ByPic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPic, opts...).ToFunc()
}

// ByVenueID orders the results by the venue_id field.
func ByVenueID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVenueID, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByIsShow orders the results by the is_show field.
func ByIsShow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsShow, opts...).ToFunc()
}

// ByIsAccessible orders the results by the is_accessible field.
func ByIsAccessible(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsAccessible, opts...).ToFunc()
}

// ByIsBooking orders the results by the is_booking field.
func ByIsBooking(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsBooking, opts...).ToFunc()
}

// ByInformation orders the results by the information field.
func ByInformation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInformation, opts...).ToFunc()
}

// ByVenueField orders the results by venue field.
func ByVenueField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVenueStep(), sql.OrderByField(field, opts...))
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
func newVenueStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenueInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
	)
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProductsTable, ProductsPrimaryKey...),
	)
}
