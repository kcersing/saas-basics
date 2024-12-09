// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/biz/dal/db/ent/bootcamp"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Bootcamp is the model entity for the Bootcamp schema.
type Bootcamp struct {
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
	// 训练营名称
	Name string `json:"name,omitempty"`
	// 报名人数
	SignNumber int64 `json:"sign_number,omitempty"`
	// 报名开始时间
	SignStartAt time.Time `json:"sign_start_at,omitempty"`
	// 报名结束时间
	SignEndAt time.Time `json:"sign_end_at,omitempty"`
	// 训练营人数
	Number int64 `json:"number,omitempty"`
	// 训练营开始时间
	StartAt time.Time `json:"start_at,omitempty"`
	// 训练营结束时间
	EndAt time.Time `json:"end_at,omitempty"`
	// 训练营图片
	Pic string `json:"pic,omitempty"`
	// 主办方
	Sponsor string `json:"sponsor,omitempty"`
	// 费用
	Fee float64 `json:"fee,omitempty"`
	// 是否有费用 1 无 2 有
	IsFee int64 `json:"is_fee,omitempty"`
	// 是否支持取消报名 0支持 1不支持
	IsCancel int64 `json:"is_cancel,omitempty"`
	// 取消时间
	CancelTime int64 `json:"cancel_time,omitempty"`
	// 详情
	Detail string `json:"detail,omitempty"`
	// 报名信息
	SignFields string `json:"sign_fields,omitempty"`
	// 状态[1:未报名;2:报名中;3:未开始;4:进行中;5:结束]
	Condition int64 `json:"condition,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BootcampQuery when eager-loading is set.
	Edges        BootcampEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BootcampEdges holds the relations/edges for other nodes in the graph.
type BootcampEdges struct {
	// BootcampParticipants holds the value of the bootcamp_participants edge.
	BootcampParticipants []*BootcampParticipant `json:"bootcamp_participants,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BootcampParticipantsOrErr returns the BootcampParticipants value or an error if the edge
// was not loaded in eager-loading.
func (e BootcampEdges) BootcampParticipantsOrErr() ([]*BootcampParticipant, error) {
	if e.loadedTypes[0] {
		return e.BootcampParticipants, nil
	}
	return nil, &NotLoadedError{edge: "bootcamp_participants"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bootcamp) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bootcamp.FieldFee:
			values[i] = new(sql.NullFloat64)
		case bootcamp.FieldID, bootcamp.FieldDelete, bootcamp.FieldCreatedID, bootcamp.FieldStatus, bootcamp.FieldSignNumber, bootcamp.FieldNumber, bootcamp.FieldIsFee, bootcamp.FieldIsCancel, bootcamp.FieldCancelTime, bootcamp.FieldCondition:
			values[i] = new(sql.NullInt64)
		case bootcamp.FieldName, bootcamp.FieldPic, bootcamp.FieldSponsor, bootcamp.FieldDetail, bootcamp.FieldSignFields:
			values[i] = new(sql.NullString)
		case bootcamp.FieldCreatedAt, bootcamp.FieldUpdatedAt, bootcamp.FieldSignStartAt, bootcamp.FieldSignEndAt, bootcamp.FieldStartAt, bootcamp.FieldEndAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bootcamp fields.
func (b *Bootcamp) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bootcamp.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int64(value.Int64)
		case bootcamp.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case bootcamp.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case bootcamp.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				b.Delete = value.Int64
			}
		case bootcamp.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				b.CreatedID = value.Int64
			}
		case bootcamp.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				b.Status = value.Int64
			}
		case bootcamp.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = value.String
			}
		case bootcamp.FieldSignNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sign_number", values[i])
			} else if value.Valid {
				b.SignNumber = value.Int64
			}
		case bootcamp.FieldSignStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_start_at", values[i])
			} else if value.Valid {
				b.SignStartAt = value.Time
			}
		case bootcamp.FieldSignEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_end_at", values[i])
			} else if value.Valid {
				b.SignEndAt = value.Time
			}
		case bootcamp.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				b.Number = value.Int64
			}
		case bootcamp.FieldStartAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				b.StartAt = value.Time
			}
		case bootcamp.FieldEndAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_at", values[i])
			} else if value.Valid {
				b.EndAt = value.Time
			}
		case bootcamp.FieldPic:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pic", values[i])
			} else if value.Valid {
				b.Pic = value.String
			}
		case bootcamp.FieldSponsor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sponsor", values[i])
			} else if value.Valid {
				b.Sponsor = value.String
			}
		case bootcamp.FieldFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				b.Fee = value.Float64
			}
		case bootcamp.FieldIsFee:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_fee", values[i])
			} else if value.Valid {
				b.IsFee = value.Int64
			}
		case bootcamp.FieldIsCancel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_cancel", values[i])
			} else if value.Valid {
				b.IsCancel = value.Int64
			}
		case bootcamp.FieldCancelTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field cancel_time", values[i])
			} else if value.Valid {
				b.CancelTime = value.Int64
			}
		case bootcamp.FieldDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field detail", values[i])
			} else if value.Valid {
				b.Detail = value.String
			}
		case bootcamp.FieldSignFields:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign_fields", values[i])
			} else if value.Valid {
				b.SignFields = value.String
			}
		case bootcamp.FieldCondition:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field condition", values[i])
			} else if value.Valid {
				b.Condition = value.Int64
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Bootcamp.
// This includes values selected through modifiers, order, etc.
func (b *Bootcamp) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryBootcampParticipants queries the "bootcamp_participants" edge of the Bootcamp entity.
func (b *Bootcamp) QueryBootcampParticipants() *BootcampParticipantQuery {
	return NewBootcampClient(b.config).QueryBootcampParticipants(b)
}

// Update returns a builder for updating this Bootcamp.
// Note that you need to call Bootcamp.Unwrap() before calling this method if this Bootcamp
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bootcamp) Update() *BootcampUpdateOne {
	return NewBootcampClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Bootcamp entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Bootcamp) Unwrap() *Bootcamp {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bootcamp is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bootcamp) String() string {
	var builder strings.Builder
	builder.WriteString("Bootcamp(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", b.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", b.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", b.Status))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(b.Name)
	builder.WriteString(", ")
	builder.WriteString("sign_number=")
	builder.WriteString(fmt.Sprintf("%v", b.SignNumber))
	builder.WriteString(", ")
	builder.WriteString("sign_start_at=")
	builder.WriteString(b.SignStartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sign_end_at=")
	builder.WriteString(b.SignEndAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", b.Number))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(b.StartAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_at=")
	builder.WriteString(b.EndAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("pic=")
	builder.WriteString(b.Pic)
	builder.WriteString(", ")
	builder.WriteString("sponsor=")
	builder.WriteString(b.Sponsor)
	builder.WriteString(", ")
	builder.WriteString("fee=")
	builder.WriteString(fmt.Sprintf("%v", b.Fee))
	builder.WriteString(", ")
	builder.WriteString("is_fee=")
	builder.WriteString(fmt.Sprintf("%v", b.IsFee))
	builder.WriteString(", ")
	builder.WriteString("is_cancel=")
	builder.WriteString(fmt.Sprintf("%v", b.IsCancel))
	builder.WriteString(", ")
	builder.WriteString("cancel_time=")
	builder.WriteString(fmt.Sprintf("%v", b.CancelTime))
	builder.WriteString(", ")
	builder.WriteString("detail=")
	builder.WriteString(b.Detail)
	builder.WriteString(", ")
	builder.WriteString("sign_fields=")
	builder.WriteString(b.SignFields)
	builder.WriteString(", ")
	builder.WriteString("condition=")
	builder.WriteString(fmt.Sprintf("%v", b.Condition))
	builder.WriteByte(')')
	return builder.String()
}

// Bootcamps is a parsable slice of Bootcamp.
type Bootcamps []*Bootcamp
