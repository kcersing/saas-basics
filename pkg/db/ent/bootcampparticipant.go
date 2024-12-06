// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/bootcamp"
	"saas/pkg/db/ent/bootcampparticipant"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BootcampParticipant is the model entity for the BootcampParticipant schema.
type BootcampParticipant struct {
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
	// 赛事id
	BootcampID int64 `json:"bootcamp_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 更多
	Fields string `json:"fields,omitempty"`
	// 订单ID
	OrderID int64 `json:"order_id,omitempty"`
	// 订单编号
	OrderSn string `json:"order_sn,omitempty"`
	// 费用
	Fee float64 `json:"fee,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BootcampParticipantQuery when eager-loading is set.
	Edges        BootcampParticipantEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BootcampParticipantEdges holds the relations/edges for other nodes in the graph.
type BootcampParticipantEdges struct {
	// Bootcamp holds the value of the bootcamp edge.
	Bootcamp *Bootcamp `json:"bootcamp,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BootcampOrErr returns the Bootcamp value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BootcampParticipantEdges) BootcampOrErr() (*Bootcamp, error) {
	if e.loadedTypes[0] {
		if e.Bootcamp == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: bootcamp.Label}
		}
		return e.Bootcamp, nil
	}
	return nil, &NotLoadedError{edge: "bootcamp"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BootcampParticipant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bootcampparticipant.FieldFee:
			values[i] = new(sql.NullFloat64)
		case bootcampparticipant.FieldID, bootcampparticipant.FieldCreatedID, bootcampparticipant.FieldStatus, bootcampparticipant.FieldBootcampID, bootcampparticipant.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case bootcampparticipant.FieldName, bootcampparticipant.FieldMobile, bootcampparticipant.FieldFields, bootcampparticipant.FieldOrderSn:
			values[i] = new(sql.NullString)
		case bootcampparticipant.FieldCreatedAt, bootcampparticipant.FieldUpdatedAt, bootcampparticipant.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BootcampParticipant fields.
func (bp *BootcampParticipant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bootcampparticipant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bp.ID = int64(value.Int64)
		case bootcampparticipant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bp.CreatedAt = value.Time
			}
		case bootcampparticipant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bp.UpdatedAt = value.Time
			}
		case bootcampparticipant.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				bp.DeleteAt = value.Time
			}
		case bootcampparticipant.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				bp.CreatedID = value.Int64
			}
		case bootcampparticipant.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				bp.Status = value.Int64
			}
		case bootcampparticipant.FieldBootcampID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field bootcamp_id", values[i])
			} else if value.Valid {
				bp.BootcampID = value.Int64
			}
		case bootcampparticipant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				bp.Name = value.String
			}
		case bootcampparticipant.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				bp.Mobile = value.String
			}
		case bootcampparticipant.FieldFields:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fields", values[i])
			} else if value.Valid {
				bp.Fields = value.String
			}
		case bootcampparticipant.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				bp.OrderID = value.Int64
			}
		case bootcampparticipant.FieldOrderSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_sn", values[i])
			} else if value.Valid {
				bp.OrderSn = value.String
			}
		case bootcampparticipant.FieldFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				bp.Fee = value.Float64
			}
		default:
			bp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BootcampParticipant.
// This includes values selected through modifiers, order, etc.
func (bp *BootcampParticipant) Value(name string) (ent.Value, error) {
	return bp.selectValues.Get(name)
}

// QueryBootcamp queries the "bootcamp" edge of the BootcampParticipant entity.
func (bp *BootcampParticipant) QueryBootcamp() *BootcampQuery {
	return NewBootcampParticipantClient(bp.config).QueryBootcamp(bp)
}

// Update returns a builder for updating this BootcampParticipant.
// Note that you need to call BootcampParticipant.Unwrap() before calling this method if this BootcampParticipant
// was returned from a transaction, and the transaction was committed or rolled back.
func (bp *BootcampParticipant) Update() *BootcampParticipantUpdateOne {
	return NewBootcampParticipantClient(bp.config).UpdateOne(bp)
}

// Unwrap unwraps the BootcampParticipant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bp *BootcampParticipant) Unwrap() *BootcampParticipant {
	_tx, ok := bp.config.driver.(*txDriver)
	if !ok {
		panic("ent: BootcampParticipant is not a transactional entity")
	}
	bp.config.driver = _tx.drv
	return bp
}

// String implements the fmt.Stringer.
func (bp *BootcampParticipant) String() string {
	var builder strings.Builder
	builder.WriteString("BootcampParticipant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(bp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(bp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(bp.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", bp.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", bp.Status))
	builder.WriteString(", ")
	builder.WriteString("bootcamp_id=")
	builder.WriteString(fmt.Sprintf("%v", bp.BootcampID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(bp.Name)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(bp.Mobile)
	builder.WriteString(", ")
	builder.WriteString("fields=")
	builder.WriteString(bp.Fields)
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", bp.OrderID))
	builder.WriteString(", ")
	builder.WriteString("order_sn=")
	builder.WriteString(bp.OrderSn)
	builder.WriteString(", ")
	builder.WriteString("fee=")
	builder.WriteString(fmt.Sprintf("%v", bp.Fee))
	builder.WriteByte(')')
	return builder.String()
}

// BootcampParticipants is a parsable slice of BootcampParticipant.
type BootcampParticipants []*BootcampParticipant
