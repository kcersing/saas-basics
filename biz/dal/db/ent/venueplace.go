// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/idl_gen/model/base"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// VenuePlace is the model entity for the VenuePlace schema.
type VenuePlace struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// last delete  1:已删除
	Delete int64 `json:"delete,omitempty"`
	// created
	CreatedID int64 `json:"created_id,omitempty"`
	// 状态[1:正常,2:禁用]
	Status int64 `json:"status,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 分类
	Classify int64 `json:"classify,omitempty"`
	// pic | 照片
	Pic string `json:"pic,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 可容纳人数
	Number int64 `json:"number,omitempty"`
	// 是否展示:1展示;2不展示
	IsShow int64 `json:"is_show,omitempty"`
	// 是否展示;1开放;2关闭
	IsAccessible int64 `json:"is_accessible,omitempty"`
	// 是否预约:1可预约;2不可
	IsBooking int64 `json:"is_booking,omitempty"`
	// 详情
	Information string `json:"information,omitempty"`
	// 座位
	Seat []*base.Seat `json:"seat,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VenuePlaceQuery when eager-loading is set.
	Edges        VenuePlaceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// VenuePlaceEdges holds the relations/edges for other nodes in the graph.
type VenuePlaceEdges struct {
	// Venue holds the value of the venue edge.
	Venue *Venue `json:"venue,omitempty"`
	// Products holds the value of the products edge.
	Products []*Product `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// VenueOrErr returns the Venue value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VenuePlaceEdges) VenueOrErr() (*Venue, error) {
	if e.loadedTypes[0] {
		if e.Venue == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: venue.Label}
		}
		return e.Venue, nil
	}
	return nil, &NotLoadedError{edge: "venue"}
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e VenuePlaceEdges) ProductsOrErr() ([]*Product, error) {
	if e.loadedTypes[1] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VenuePlace) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case venueplace.FieldSeat:
			values[i] = new([]byte)
		case venueplace.FieldID, venueplace.FieldDelete, venueplace.FieldCreatedID, venueplace.FieldStatus, venueplace.FieldClassify, venueplace.FieldVenueID, venueplace.FieldNumber, venueplace.FieldIsShow, venueplace.FieldIsAccessible, venueplace.FieldIsBooking:
			values[i] = new(sql.NullInt64)
		case venueplace.FieldName, venueplace.FieldPic, venueplace.FieldInformation:
			values[i] = new(sql.NullString)
		case venueplace.FieldCreatedAt, venueplace.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VenuePlace fields.
func (vp *VenuePlace) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case venueplace.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vp.ID = int64(value.Int64)
		case venueplace.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				vp.CreatedAt = value.Time
			}
		case venueplace.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				vp.UpdatedAt = value.Time
			}
		case venueplace.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				vp.Delete = value.Int64
			}
		case venueplace.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				vp.CreatedID = value.Int64
			}
		case venueplace.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				vp.Status = value.Int64
			}
		case venueplace.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				vp.Name = value.String
			}
		case venueplace.FieldClassify:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field classify", values[i])
			} else if value.Valid {
				vp.Classify = value.Int64
			}
		case venueplace.FieldPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic", values[i])
			} else if value.Valid {
				vp.Pic = value.String
			}
		case venueplace.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				vp.VenueID = value.Int64
			}
		case venueplace.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				vp.Number = value.Int64
			}
		case venueplace.FieldIsShow:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_show", values[i])
			} else if value.Valid {
				vp.IsShow = value.Int64
			}
		case venueplace.FieldIsAccessible:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_accessible", values[i])
			} else if value.Valid {
				vp.IsAccessible = value.Int64
			}
		case venueplace.FieldIsBooking:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_booking", values[i])
			} else if value.Valid {
				vp.IsBooking = value.Int64
			}
		case venueplace.FieldInformation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field information", values[i])
			} else if value.Valid {
				vp.Information = value.String
			}
		case venueplace.FieldSeat:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field seat", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &vp.Seat); err != nil {
					return fmt.Errorf("unmarshal field seat: %w", err)
				}
			}
		default:
			vp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VenuePlace.
// This includes values selected through modifiers, order, etc.
func (vp *VenuePlace) Value(name string) (ent.Value, error) {
	return vp.selectValues.Get(name)
}

// QueryVenue queries the "venue" edge of the VenuePlace entity.
func (vp *VenuePlace) QueryVenue() *VenueQuery {
	return NewVenuePlaceClient(vp.config).QueryVenue(vp)
}

// QueryProducts queries the "products" edge of the VenuePlace entity.
func (vp *VenuePlace) QueryProducts() *ProductQuery {
	return NewVenuePlaceClient(vp.config).QueryProducts(vp)
}

// Update returns a builder for updating this VenuePlace.
// Note that you need to call VenuePlace.Unwrap() before calling this method if this VenuePlace
// was returned from a transaction, and the transaction was committed or rolled back.
func (vp *VenuePlace) Update() *VenuePlaceUpdateOne {
	return NewVenuePlaceClient(vp.config).UpdateOne(vp)
}

// Unwrap unwraps the VenuePlace entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vp *VenuePlace) Unwrap() *VenuePlace {
	_tx, ok := vp.config.driver.(*txDriver)
	if !ok {
		panic("ent: VenuePlace is not a transactional entity")
	}
	vp.config.driver = _tx.drv
	return vp
}

// String implements the fmt.Stringer.
func (vp *VenuePlace) String() string {
	var builder strings.Builder
	builder.WriteString("VenuePlace(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(vp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(vp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", vp.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", vp.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", vp.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(vp.Name)
	builder.WriteString(", ")
	builder.WriteString("classify=")
	builder.WriteString(fmt.Sprintf("%v", vp.Classify))
	builder.WriteString(", ")
	builder.WriteString("pic=")
	builder.WriteString(vp.Pic)
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", vp.VenueID))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", vp.Number))
	builder.WriteString(", ")
	builder.WriteString("is_show=")
	builder.WriteString(fmt.Sprintf("%v", vp.IsShow))
	builder.WriteString(", ")
	builder.WriteString("is_accessible=")
	builder.WriteString(fmt.Sprintf("%v", vp.IsAccessible))
	builder.WriteString(", ")
	builder.WriteString("is_booking=")
	builder.WriteString(fmt.Sprintf("%v", vp.IsBooking))
	builder.WriteString(", ")
	builder.WriteString("information=")
	builder.WriteString(vp.Information)
	builder.WriteString(", ")
	builder.WriteString("seat=")
	builder.WriteString(fmt.Sprintf("%v", vp.Seat))
	builder.WriteByte(')')
	return builder.String()
}

// VenuePlaces is a parsable slice of VenuePlace.
type VenuePlaces []*VenuePlace
