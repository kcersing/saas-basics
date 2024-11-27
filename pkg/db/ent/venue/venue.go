// Code generated by ent, DO NOT EDIT.

package venue

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the venue type in the database.
	Label = "venue"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldAddressDetail holds the string denoting the address_detail field in the database.
	FieldAddressDetail = "address_detail"
	// FieldLatitude holds the string denoting the latitude field in the database.
	FieldLatitude = "latitude"
	// FieldLongitude holds the string denoting the longitude field in the database.
	FieldLongitude = "longitude"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldInformation holds the string denoting the information field in the database.
	FieldInformation = "information"
	// FieldPic holds the string denoting the pic field in the database.
	FieldPic = "pic"
	// EdgePlaces holds the string denoting the places edge name in mutations.
	EdgePlaces = "places"
	// EdgeVenueOrders holds the string denoting the venue_orders edge name in mutations.
	EdgeVenueOrders = "venue_orders"
	// Table holds the table name of the venue in the database.
	Table = "venue"
	// PlacesTable is the table that holds the places relation/edge.
	PlacesTable = "venue_place"
	// PlacesInverseTable is the table name for the VenuePlace entity.
	// It exists in this package in order to avoid circular dependency with the "venueplace" package.
	PlacesInverseTable = "venue_place"
	// PlacesColumn is the table column denoting the places relation/edge.
	PlacesColumn = "venue_id"
	// VenueOrdersTable is the table that holds the venue_orders relation/edge.
	VenueOrdersTable = "order"
	// VenueOrdersInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	VenueOrdersInverseTable = "order"
	// VenueOrdersColumn is the table column denoting the venue_orders relation/edge.
	VenueOrdersColumn = "venue_id"
)

// Columns holds all SQL columns for venue fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldName,
	FieldAddress,
	FieldAddressDetail,
	FieldLatitude,
	FieldLongitude,
	FieldMobile,
	FieldEmail,
	FieldInformation,
	FieldPic,
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
)

// OrderOption defines the ordering options for the Venue queries.
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

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByAddressDetail orders the results by the address_detail field.
func ByAddressDetail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddressDetail, opts...).ToFunc()
}

// ByLatitude orders the results by the latitude field.
func ByLatitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLatitude, opts...).ToFunc()
}

// ByLongitude orders the results by the longitude field.
func ByLongitude(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLongitude, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByInformation orders the results by the information field.
func ByInformation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldInformation, opts...).ToFunc()
}

// ByPic orders the results by the pic field.
func ByPic(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPic, opts...).ToFunc()
}

// ByPlacesCount orders the results by places count.
func ByPlacesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlacesStep(), opts...)
	}
}

// ByPlaces orders the results by places terms.
func ByPlaces(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlacesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVenueOrdersCount orders the results by venue_orders count.
func ByVenueOrdersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVenueOrdersStep(), opts...)
	}
}

// ByVenueOrders orders the results by venue_orders terms.
func ByVenueOrders(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVenueOrdersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPlacesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlacesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PlacesTable, PlacesColumn),
	)
}
func newVenueOrdersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VenueOrdersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VenueOrdersTable, VenueOrdersColumn),
	)
}
