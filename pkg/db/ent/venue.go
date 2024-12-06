// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/venue"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Venue is the model entity for the Venue schema.
type Venue struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// last delete time
	DeleteAt time.Time `json:"delete_at,omitempty"`
	// created
	CreatedID int64 `json:"created_id,omitempty"`
	// 状态[1:正常,2:禁用]
	Status int64 `json:"status,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 地址 省/市/区
	Address string `json:"address,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty"`
	// 维度
	Latitude string `json:"latitude,omitempty"`
	// 经度
	Longitude string `json:"longitude,omitempty"`
	// 联系电话
	Mobile string `json:"mobile,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 详情
	Information string `json:"information,omitempty"`
	// pic | 照片
	Pic string `json:"pic,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VenueQuery when eager-loading is set.
	Edges        VenueEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VenueEdges holds the relations/edges for other nodes in the graph.
type VenueEdges struct {
	// Places holds the value of the places edge.
	Places []*VenuePlace `json:"places,omitempty"`
	// VenueOrders holds the value of the venue_orders edge.
	VenueOrders []*Order `json:"venue_orders,omitempty"`
	// VenueEntry holds the value of the venue_entry edge.
	VenueEntry []*EntryLogs `json:"venue_entry,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PlacesOrErr returns the Places value or an error if the edge
// was not loaded in eager-loading.
func (e VenueEdges) PlacesOrErr() ([]*VenuePlace, error) {
	if e.loadedTypes[0] {
		return e.Places, nil
	}
	return nil, &NotLoadedError{edge: "places"}
}

// VenueOrdersOrErr returns the VenueOrders value or an error if the edge
// was not loaded in eager-loading.
func (e VenueEdges) VenueOrdersOrErr() ([]*Order, error) {
	if e.loadedTypes[1] {
		return e.VenueOrders, nil
	}
	return nil, &NotLoadedError{edge: "venue_orders"}
}

// VenueEntryOrErr returns the VenueEntry value or an error if the edge
// was not loaded in eager-loading.
func (e VenueEdges) VenueEntryOrErr() ([]*EntryLogs, error) {
	if e.loadedTypes[2] {
		return e.VenueEntry, nil
	}
	return nil, &NotLoadedError{edge: "venue_entry"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Venue) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case venue.FieldID, venue.FieldCreatedID, venue.FieldStatus:
			values[i] = new(sql.NullInt64)
		case venue.FieldName, venue.FieldAddress, venue.FieldAddressDetail, venue.FieldLatitude, venue.FieldLongitude, venue.FieldMobile, venue.FieldEmail, venue.FieldInformation, venue.FieldPic:
			values[i] = new(sql.NullString)
		case venue.FieldCreatedAt, venue.FieldUpdatedAt, venue.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Venue fields.
func (v *Venue) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case venue.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int64(value.Int64)
		case venue.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				v.CreatedAt = value.Time
			}
		case venue.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				v.UpdatedAt = value.Time
			}
		case venue.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				v.DeleteAt = value.Time
			}
		case venue.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				v.CreatedID = value.Int64
			}
		case venue.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				v.Status = value.Int64
			}
		case venue.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				v.Name = value.String
			}
		case venue.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				v.Address = value.String
			}
		case venue.FieldAddressDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address_detail", values[i])
			} else if value.Valid {
				v.AddressDetail = value.String
			}
		case venue.FieldLatitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field latitude", values[i])
			} else if value.Valid {
				v.Latitude = value.String
			}
		case venue.FieldLongitude:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field longitude", values[i])
			} else if value.Valid {
				v.Longitude = value.String
			}
		case venue.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				v.Mobile = value.String
			}
		case venue.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				v.Email = value.String
			}
		case venue.FieldInformation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field information", values[i])
			} else if value.Valid {
				v.Information = value.String
			}
		case venue.FieldPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic", values[i])
			} else if value.Valid {
				v.Pic = value.String
			}
		default:
			v.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Venue.
// This includes values selected through modifiers, order, etc.
func (v *Venue) Value(name string) (ent.Value, error) {
	return v.selectValues.Get(name)
}

// QueryPlaces queries the "places" edge of the Venue entity.
func (v *Venue) QueryPlaces() *VenuePlaceQuery {
	return NewVenueClient(v.config).QueryPlaces(v)
}

// QueryVenueOrders queries the "venue_orders" edge of the Venue entity.
func (v *Venue) QueryVenueOrders() *OrderQuery {
	return NewVenueClient(v.config).QueryVenueOrders(v)
}

// QueryVenueEntry queries the "venue_entry" edge of the Venue entity.
func (v *Venue) QueryVenueEntry() *EntryLogsQuery {
	return NewVenueClient(v.config).QueryVenueEntry(v)
}

// Update returns a builder for updating this Venue.
// Note that you need to call Venue.Unwrap() before calling this method if this Venue
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Venue) Update() *VenueUpdateOne {
	return NewVenueClient(v.config).UpdateOne(v)
}

// Unwrap unwraps the Venue entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Venue) Unwrap() *Venue {
	_tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Venue is not a transactional entity")
	}
	v.config.driver = _tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Venue) String() string {
	var builder strings.Builder
	builder.WriteString("Venue(")
	builder.WriteString(fmt.Sprintf("id=%v, ", v.ID))
	builder.WriteString("created_at=")
	builder.WriteString(v.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(v.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(v.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", v.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", v.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(v.Name)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(v.Address)
	builder.WriteString(", ")
	builder.WriteString("address_detail=")
	builder.WriteString(v.AddressDetail)
	builder.WriteString(", ")
	builder.WriteString("latitude=")
	builder.WriteString(v.Latitude)
	builder.WriteString(", ")
	builder.WriteString("longitude=")
	builder.WriteString(v.Longitude)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(v.Mobile)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(v.Email)
	builder.WriteString(", ")
	builder.WriteString("information=")
	builder.WriteString(v.Information)
	builder.WriteString(", ")
	builder.WriteString("pic=")
	builder.WriteString(v.Pic)
	builder.WriteByte(')')
	return builder.String()
}

// Venues is a parsable slice of Venue.
type Venues []*Venue
