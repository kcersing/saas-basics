// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/userscheduling"
	"saas/idl_gen/model/base"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// UserScheduling is the model entity for the UserScheduling schema.
type UserScheduling struct {
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
	// 日期
	Date base.UserSchedulingDate `json:"date,omitempty"`
	// 員工id
	UserID int64 `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSchedulingQuery when eager-loading is set.
	Edges        UserSchedulingEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserSchedulingEdges holds the relations/edges for other nodes in the graph.
type UserSchedulingEdges struct {
	// Users holds the value of the users edge.
	Users *User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSchedulingEdges) UsersOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Users == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserScheduling) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case userscheduling.FieldDate:
			values[i] = new([]byte)
		case userscheduling.FieldID, userscheduling.FieldDelete, userscheduling.FieldCreatedID, userscheduling.FieldStatus, userscheduling.FieldUserID:
			values[i] = new(sql.NullInt64)
		case userscheduling.FieldCreatedAt, userscheduling.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserScheduling fields.
func (us *UserScheduling) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userscheduling.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			us.ID = int64(value.Int64)
		case userscheduling.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				us.CreatedAt = value.Time
			}
		case userscheduling.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				us.UpdatedAt = value.Time
			}
		case userscheduling.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				us.Delete = value.Int64
			}
		case userscheduling.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				us.CreatedID = value.Int64
			}
		case userscheduling.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				us.Status = value.Int64
			}
		case userscheduling.FieldDate:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &us.Date); err != nil {
					return fmt.Errorf("unmarshal field date: %w", err)
				}
			}
		case userscheduling.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				us.UserID = value.Int64
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserScheduling.
// This includes values selected through modifiers, order, etc.
func (us *UserScheduling) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the UserScheduling entity.
func (us *UserScheduling) QueryUsers() *UserQuery {
	return NewUserSchedulingClient(us.config).QueryUsers(us)
}

// Update returns a builder for updating this UserScheduling.
// Note that you need to call UserScheduling.Unwrap() before calling this method if this UserScheduling
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserScheduling) Update() *UserSchedulingUpdateOne {
	return NewUserSchedulingClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserScheduling entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserScheduling) Unwrap() *UserScheduling {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserScheduling is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserScheduling) String() string {
	var builder strings.Builder
	builder.WriteString("UserScheduling(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("created_at=")
	builder.WriteString(us.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(us.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", us.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", us.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", us.Status))
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(fmt.Sprintf("%v", us.Date))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", us.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// UserSchedulings is a parsable slice of UserScheduling.
type UserSchedulings []*UserScheduling
