// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"saas/biz/dal/db/ent/schedule"
	"saas/biz/dal/db/ent/schedulemember"
	"saas/idl_gen/model/base"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ScheduleMember is the model entity for the ScheduleMember schema.
type ScheduleMember struct {
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
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 场地ID
	PlaceID int64 `json:"place_id,omitempty"`
	// 课程
	ProductID int64 `json:"product_id,omitempty"`
	// 课程ID
	ScheduleID int64 `json:"schedule_id,omitempty"`
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 会员购买课ID
	MemberProductID int64 `json:"member_product_id,omitempty"`
	// 类型
	Type string `json:"type,omitempty"`
	// 开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 结束时间
	EndTime time.Time `json:"end_time,omitempty"`
	// 上课签到时间
	SignStartTime time.Time `json:"sign_start_time,omitempty"`
	// 下课签到时间
	SignEndTime time.Time `json:"sign_end_time,omitempty"`
	// 座位
	Seat base.Seat `json:"seat,omitempty"`
	// 課包 1支持2不支持
	IsCourse int64 `json:"is_course,omitempty"`
	// 会员名称
	MemberName string `json:"member_name,omitempty"`
	// 会员产品名称
	MemberProductName string `json:"member_product_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScheduleMemberQuery when eager-loading is set.
	Edges        ScheduleMemberEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ScheduleMemberEdges holds the relations/edges for other nodes in the graph.
type ScheduleMemberEdges struct {
	// Schedule holds the value of the schedule edge.
	Schedule *Schedule `json:"schedule,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ScheduleMemberEdges) ScheduleOrErr() (*Schedule, error) {
	if e.loadedTypes[0] {
		if e.Schedule == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: schedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ScheduleMember) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case schedulemember.FieldSeat:
			values[i] = new([]byte)
		case schedulemember.FieldID, schedulemember.FieldDelete, schedulemember.FieldCreatedID, schedulemember.FieldStatus, schedulemember.FieldVenueID, schedulemember.FieldPlaceID, schedulemember.FieldProductID, schedulemember.FieldScheduleID, schedulemember.FieldMemberID, schedulemember.FieldMemberProductID, schedulemember.FieldIsCourse:
			values[i] = new(sql.NullInt64)
		case schedulemember.FieldType, schedulemember.FieldMemberName, schedulemember.FieldMemberProductName, schedulemember.FieldRemark:
			values[i] = new(sql.NullString)
		case schedulemember.FieldCreatedAt, schedulemember.FieldUpdatedAt, schedulemember.FieldStartTime, schedulemember.FieldEndTime, schedulemember.FieldSignStartTime, schedulemember.FieldSignEndTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ScheduleMember fields.
func (sm *ScheduleMember) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case schedulemember.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sm.ID = int64(value.Int64)
		case schedulemember.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				sm.CreatedAt = value.Time
			}
		case schedulemember.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				sm.UpdatedAt = value.Time
			}
		case schedulemember.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				sm.Delete = value.Int64
			}
		case schedulemember.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				sm.CreatedID = value.Int64
			}
		case schedulemember.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				sm.Status = value.Int64
			}
		case schedulemember.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				sm.VenueID = value.Int64
			}
		case schedulemember.FieldPlaceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field place_id", values[i])
			} else if value.Valid {
				sm.PlaceID = value.Int64
			}
		case schedulemember.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				sm.ProductID = value.Int64
			}
		case schedulemember.FieldScheduleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field schedule_id", values[i])
			} else if value.Valid {
				sm.ScheduleID = value.Int64
			}
		case schedulemember.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				sm.MemberID = value.Int64
			}
		case schedulemember.FieldMemberProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_product_id", values[i])
			} else if value.Valid {
				sm.MemberProductID = value.Int64
			}
		case schedulemember.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				sm.Type = value.String
			}
		case schedulemember.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				sm.StartTime = value.Time
			}
		case schedulemember.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				sm.EndTime = value.Time
			}
		case schedulemember.FieldSignStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_start_time", values[i])
			} else if value.Valid {
				sm.SignStartTime = value.Time
			}
		case schedulemember.FieldSignEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_end_time", values[i])
			} else if value.Valid {
				sm.SignEndTime = value.Time
			}
		case schedulemember.FieldSeat:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field seat", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &sm.Seat); err != nil {
					return fmt.Errorf("unmarshal field seat: %w", err)
				}
			}
		case schedulemember.FieldIsCourse:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field is_course", values[i])
			} else if value.Valid {
				sm.IsCourse = value.Int64
			}
		case schedulemember.FieldMemberName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_name", values[i])
			} else if value.Valid {
				sm.MemberName = value.String
			}
		case schedulemember.FieldMemberProductName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field member_product_name", values[i])
			} else if value.Valid {
				sm.MemberProductName = value.String
			}
		case schedulemember.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				sm.Remark = value.String
			}
		default:
			sm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ScheduleMember.
// This includes values selected through modifiers, order, etc.
func (sm *ScheduleMember) Value(name string) (ent.Value, error) {
	return sm.selectValues.Get(name)
}

// QuerySchedule queries the "schedule" edge of the ScheduleMember entity.
func (sm *ScheduleMember) QuerySchedule() *ScheduleQuery {
	return NewScheduleMemberClient(sm.config).QuerySchedule(sm)
}

// Update returns a builder for updating this ScheduleMember.
// Note that you need to call ScheduleMember.Unwrap() before calling this method if this ScheduleMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (sm *ScheduleMember) Update() *ScheduleMemberUpdateOne {
	return NewScheduleMemberClient(sm.config).UpdateOne(sm)
}

// Unwrap unwraps the ScheduleMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sm *ScheduleMember) Unwrap() *ScheduleMember {
	_tx, ok := sm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ScheduleMember is not a transactional entity")
	}
	sm.config.driver = _tx.drv
	return sm
}

// String implements the fmt.Stringer.
func (sm *ScheduleMember) String() string {
	var builder strings.Builder
	builder.WriteString("ScheduleMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", sm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(sm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(sm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", sm.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", sm.Status))
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.VenueID))
	builder.WriteString(", ")
	builder.WriteString("place_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.PlaceID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.ProductID))
	builder.WriteString(", ")
	builder.WriteString("schedule_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.ScheduleID))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.MemberID))
	builder.WriteString(", ")
	builder.WriteString("member_product_id=")
	builder.WriteString(fmt.Sprintf("%v", sm.MemberProductID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(sm.Type)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(sm.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(sm.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sign_start_time=")
	builder.WriteString(sm.SignStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sign_end_time=")
	builder.WriteString(sm.SignEndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("seat=")
	builder.WriteString(fmt.Sprintf("%v", sm.Seat))
	builder.WriteString(", ")
	builder.WriteString("is_course=")
	builder.WriteString(fmt.Sprintf("%v", sm.IsCourse))
	builder.WriteString(", ")
	builder.WriteString("member_name=")
	builder.WriteString(sm.MemberName)
	builder.WriteString(", ")
	builder.WriteString("member_product_name=")
	builder.WriteString(sm.MemberProductName)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(sm.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// ScheduleMembers is a parsable slice of ScheduleMember.
type ScheduleMembers []*ScheduleMember
