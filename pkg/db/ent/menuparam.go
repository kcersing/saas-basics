// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/menuparam"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MenuParam is the model entity for the MenuParam schema.
type MenuParam struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// pass parameters via params or query | 参数类型
	Type string `json:"type,omitempty"`
	// the key of parameters | 参数键
	Key string `json:"key,omitempty"`
	// the value of parameters | 参数值
	Value string `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuParamQuery when eager-loading is set.
	Edges        MenuParamEdges `json:"edges"`
	menu_params  *int64
	selectValues sql.SelectValues
}

// MenuParamEdges holds the relations/edges for other nodes in the graph.
type MenuParamEdges struct {
	// Menus holds the value of the menus edge.
	Menus *Menu `json:"menus,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MenusOrErr returns the Menus value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuParamEdges) MenusOrErr() (*Menu, error) {
	if e.loadedTypes[0] {
		if e.Menus == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: menu.Label}
		}
		return e.Menus, nil
	}
	return nil, &NotLoadedError{edge: "menus"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MenuParam) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case menuparam.FieldID:
			values[i] = new(sql.NullInt64)
		case menuparam.FieldType, menuparam.FieldKey, menuparam.FieldValue:
			values[i] = new(sql.NullString)
		case menuparam.FieldCreatedAt, menuparam.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case menuparam.ForeignKeys[0]: // menu_params
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MenuParam fields.
func (mp *MenuParam) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case menuparam.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mp.ID = int64(value.Int64)
		case menuparam.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mp.CreatedAt = value.Time
			}
		case menuparam.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mp.UpdatedAt = value.Time
			}
		case menuparam.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mp.Type = value.String
			}
		case menuparam.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				mp.Key = value.String
			}
		case menuparam.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				mp.Value = value.String
			}
		case menuparam.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field menu_params", value)
			} else if value.Valid {
				mp.menu_params = new(int64)
				*mp.menu_params = int64(value.Int64)
			}
		default:
			mp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the MenuParam.
// This includes values selected through modifiers, order, etc.
func (mp *MenuParam) GetValue(name string) (ent.Value, error) {
	return mp.selectValues.Get(name)
}

// QueryMenus queries the "menus" edge of the MenuParam entity.
func (mp *MenuParam) QueryMenus() *MenuQuery {
	return NewMenuParamClient(mp.config).QueryMenus(mp)
}

// Update returns a builder for updating this MenuParam.
// Note that you need to call MenuParam.Unwrap() before calling this method if this MenuParam
// was returned from a transaction, and the transaction was committed or rolled back.
func (mp *MenuParam) Update() *MenuParamUpdateOne {
	return NewMenuParamClient(mp.config).UpdateOne(mp)
}

// Unwrap unwraps the MenuParam entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mp *MenuParam) Unwrap() *MenuParam {
	_tx, ok := mp.config.driver.(*txDriver)
	if !ok {
		panic("ent: MenuParam is not a transactional entity")
	}
	mp.config.driver = _tx.drv
	return mp
}

// String implements the fmt.Stringer.
func (mp *MenuParam) String() string {
	var builder strings.Builder
	builder.WriteString("MenuParam(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(mp.Type)
	builder.WriteString(", ")
	builder.WriteString("key=")
	builder.WriteString(mp.Key)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(mp.Value)
	builder.WriteByte(')')
	return builder.String()
}

// MenuParams is a parsable slice of MenuParam.
type MenuParams []*MenuParam
