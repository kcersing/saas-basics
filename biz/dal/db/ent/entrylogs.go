// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/venue"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// EntryLogs is the model entity for the EntryLogs schema.
type EntryLogs struct {
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
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 用户id
	UserID int64 `json:"user_id,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 用户产品id
	MemberProductID int64 `json:"member_product_id,omitempty"`
	// 属性id
	MemberPropertyID int64 `json:"member_property_id,omitempty"`
	// 进场时间
	EntryTime time.Time `json:"entry_time,omitempty"`
	// 离场时间
	LeavingTime time.Time `json:"leaving_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntryLogsQuery when eager-loading is set.
	Edges        EntryLogsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EntryLogsEdges holds the relations/edges for other nodes in the graph.
type EntryLogsEdges struct {
	// Venues holds the value of the venues edge.
	Venues *Venue `json:"venues,omitempty"`
	// Members holds the value of the members edge.
	Members *Member `json:"members,omitempty"`
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// VenuesOrErr returns the Venues value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntryLogsEdges) VenuesOrErr() (*Venue, error) {
	if e.loadedTypes[0] {
		if e.Venues == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: venue.Label}
		}
		return e.Venues, nil
	}
	return nil, &NotLoadedError{edge: "venues"}
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntryLogsEdges) MembersOrErr() (*Member, error) {
	if e.loadedTypes[1] {
		if e.Members == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntryLogsEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EntryLogs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entrylogs.FieldID, entrylogs.FieldDelete, entrylogs.FieldCreatedID, entrylogs.FieldMemberID, entrylogs.FieldUserID, entrylogs.FieldVenueID, entrylogs.FieldMemberProductID, entrylogs.FieldMemberPropertyID:
			values[i] = new(sql.NullInt64)
		case entrylogs.FieldCreatedAt, entrylogs.FieldUpdatedAt, entrylogs.FieldEntryTime, entrylogs.FieldLeavingTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EntryLogs fields.
func (el *EntryLogs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entrylogs.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			el.ID = int64(value.Int64)
		case entrylogs.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				el.CreatedAt = value.Time
			}
		case entrylogs.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				el.UpdatedAt = value.Time
			}
		case entrylogs.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				el.Delete = value.Int64
			}
		case entrylogs.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				el.CreatedID = value.Int64
			}
		case entrylogs.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				el.MemberID = value.Int64
			}
		case entrylogs.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				el.UserID = value.Int64
			}
		case entrylogs.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				el.VenueID = value.Int64
			}
		case entrylogs.FieldMemberProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_product_id", values[i])
			} else if value.Valid {
				el.MemberProductID = value.Int64
			}
		case entrylogs.FieldMemberPropertyID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_property_id", values[i])
			} else if value.Valid {
				el.MemberPropertyID = value.Int64
			}
		case entrylogs.FieldEntryTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field entry_time", values[i])
			} else if value.Valid {
				el.EntryTime = value.Time
			}
		case entrylogs.FieldLeavingTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field leaving_time", values[i])
			} else if value.Valid {
				el.LeavingTime = value.Time
			}
		default:
			el.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EntryLogs.
// This includes values selected through modifiers, order, etc.
func (el *EntryLogs) Value(name string) (ent.Value, error) {
	return el.selectValues.Get(name)
}

// QueryVenues queries the "venues" edge of the EntryLogs entity.
func (el *EntryLogs) QueryVenues() *VenueQuery {
	return NewEntryLogsClient(el.config).QueryVenues(el)
}

// QueryMembers queries the "members" edge of the EntryLogs entity.
func (el *EntryLogs) QueryMembers() *MemberQuery {
	return NewEntryLogsClient(el.config).QueryMembers(el)
}

// QueryUsers queries the "users" edge of the EntryLogs entity.
func (el *EntryLogs) QueryUsers() *UserQuery {
	return NewEntryLogsClient(el.config).QueryUsers(el)
}

// Update returns a builder for updating this EntryLogs.
// Note that you need to call EntryLogs.Unwrap() before calling this method if this EntryLogs
// was returned from a transaction, and the transaction was committed or rolled back.
func (el *EntryLogs) Update() *EntryLogsUpdateOne {
	return NewEntryLogsClient(el.config).UpdateOne(el)
}

// Unwrap unwraps the EntryLogs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (el *EntryLogs) Unwrap() *EntryLogs {
	_tx, ok := el.config.driver.(*txDriver)
	if !ok {
		panic("ent: EntryLogs is not a transactional entity")
	}
	el.config.driver = _tx.drv
	return el
}

// String implements the fmt.Stringer.
func (el *EntryLogs) String() string {
	var builder strings.Builder
	builder.WriteString("EntryLogs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", el.ID))
	builder.WriteString("created_at=")
	builder.WriteString(el.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(el.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", el.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", el.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", el.MemberID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", el.UserID))
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", el.VenueID))
	builder.WriteString(", ")
	builder.WriteString("member_product_id=")
	builder.WriteString(fmt.Sprintf("%v", el.MemberProductID))
	builder.WriteString(", ")
	builder.WriteString("member_property_id=")
	builder.WriteString(fmt.Sprintf("%v", el.MemberPropertyID))
	builder.WriteString(", ")
	builder.WriteString("entry_time=")
	builder.WriteString(el.EntryTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("leaving_time=")
	builder.WriteString(el.LeavingTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// EntryLogsSlice is a parsable slice of EntryLogs.
type EntryLogsSlice []*EntryLogs
