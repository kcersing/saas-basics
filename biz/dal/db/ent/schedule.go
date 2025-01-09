// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/biz/dal/db/ent/schedule"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Schedule is the model entity for the Schedule schema.
type Schedule struct {
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
	// 类型
	Type string `json:"type,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 课程
	ProductID int64 `json:"product_id,omitempty"`
	// 时长
	Length int64 `json:"length,omitempty"`
	// 场地ID
	PlaceID int64 `json:"place_id,omitempty"`
	// 上课人数
	Num int64 `json:"num,omitempty"`
	// 剩余可约人数
	NumSurplus int64 `json:"num_surplus,omitempty"`
	// 日期
	Date string `json:"date,omitempty"`
	// 开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 开始时间
	EndTime time.Time `json:"end_time,omitempty"`
	// 课程价格
	Price float64 `json:"price,omitempty"`
	// 场馆名称
	VenueName string `json:"venue_name,omitempty"`
	// 场地名称
	PlaceName string `json:"place_name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ScheduleQuery when eager-loading is set.
	Edges        ScheduleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ScheduleEdges holds the relations/edges for other nodes in the graph.
type ScheduleEdges struct {
	// Members holds the value of the members edge.
	Members []*ScheduleMember `json:"members,omitempty"`
	// Coachs holds the value of the coachs edge.
	Coachs []*ScheduleCoach `json:"coachs,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e ScheduleEdges) MembersOrErr() ([]*ScheduleMember, error) {
	if e.loadedTypes[0] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// CoachsOrErr returns the Coachs value or an error if the edge
// was not loaded in eager-loading.
func (e ScheduleEdges) CoachsOrErr() ([]*ScheduleCoach, error) {
	if e.loadedTypes[1] {
		return e.Coachs, nil
	}
	return nil, &NotLoadedError{edge: "coachs"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Schedule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case schedule.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case schedule.FieldID, schedule.FieldDelete, schedule.FieldCreatedID, schedule.FieldStatus, schedule.FieldVenueID, schedule.FieldProductID, schedule.FieldLength, schedule.FieldPlaceID, schedule.FieldNum, schedule.FieldNumSurplus:
			values[i] = new(sql.NullInt64)
		case schedule.FieldType, schedule.FieldName, schedule.FieldDate, schedule.FieldVenueName, schedule.FieldPlaceName, schedule.FieldRemark:
			values[i] = new(sql.NullString)
		case schedule.FieldCreatedAt, schedule.FieldUpdatedAt, schedule.FieldStartTime, schedule.FieldEndTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Schedule fields.
func (s *Schedule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case schedule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int64(value.Int64)
		case schedule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case schedule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				s.UpdatedAt = value.Time
			}
		case schedule.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				s.Delete = value.Int64
			}
		case schedule.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				s.CreatedID = value.Int64
			}
		case schedule.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.Int64
			}
		case schedule.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				s.Type = value.String
			}
		case schedule.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case schedule.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				s.VenueID = value.Int64
			}
		case schedule.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				s.ProductID = value.Int64
			}
		case schedule.FieldLength:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field length", values[i])
			} else if value.Valid {
				s.Length = value.Int64
			}
		case schedule.FieldPlaceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field place_id", values[i])
			} else if value.Valid {
				s.PlaceID = value.Int64
			}
		case schedule.FieldNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num", values[i])
			} else if value.Valid {
				s.Num = value.Int64
			}
		case schedule.FieldNumSurplus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num_surplus", values[i])
			} else if value.Valid {
				s.NumSurplus = value.Int64
			}
		case schedule.FieldDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field date", values[i])
			} else if value.Valid {
				s.Date = value.String
			}
		case schedule.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				s.StartTime = value.Time
			}
		case schedule.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				s.EndTime = value.Time
			}
		case schedule.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				s.Price = value.Float64
			}
		case schedule.FieldVenueName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field venue_name", values[i])
			} else if value.Valid {
				s.VenueName = value.String
			}
		case schedule.FieldPlaceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field place_name", values[i])
			} else if value.Valid {
				s.PlaceName = value.String
			}
		case schedule.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				s.Remark = value.String
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Schedule.
// This includes values selected through modifiers, order, etc.
func (s *Schedule) Value(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// QueryMembers queries the "members" edge of the Schedule entity.
func (s *Schedule) QueryMembers() *ScheduleMemberQuery {
	return NewScheduleClient(s.config).QueryMembers(s)
}

// QueryCoachs queries the "coachs" edge of the Schedule entity.
func (s *Schedule) QueryCoachs() *ScheduleCoachQuery {
	return NewScheduleClient(s.config).QueryCoachs(s)
}

// Update returns a builder for updating this Schedule.
// Note that you need to call Schedule.Unwrap() before calling this method if this Schedule
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Schedule) Update() *ScheduleUpdateOne {
	return NewScheduleClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Schedule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Schedule) Unwrap() *Schedule {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Schedule is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Schedule) String() string {
	var builder strings.Builder
	builder.WriteString("Schedule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("created_at=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(s.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", s.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(s.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", s.VenueID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", s.ProductID))
	builder.WriteString(", ")
	builder.WriteString("length=")
	builder.WriteString(fmt.Sprintf("%v", s.Length))
	builder.WriteString(", ")
	builder.WriteString("place_id=")
	builder.WriteString(fmt.Sprintf("%v", s.PlaceID))
	builder.WriteString(", ")
	builder.WriteString("num=")
	builder.WriteString(fmt.Sprintf("%v", s.Num))
	builder.WriteString(", ")
	builder.WriteString("num_surplus=")
	builder.WriteString(fmt.Sprintf("%v", s.NumSurplus))
	builder.WriteString(", ")
	builder.WriteString("date=")
	builder.WriteString(s.Date)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(s.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(s.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", s.Price))
	builder.WriteString(", ")
	builder.WriteString("venue_name=")
	builder.WriteString(s.VenueName)
	builder.WriteString(", ")
	builder.WriteString("place_name=")
	builder.WriteString(s.PlaceName)
	builder.WriteString(", ")
	builder.WriteString("remark=")
	builder.WriteString(s.Remark)
	builder.WriteByte(')')
	return builder.String()
}

// Schedules is a parsable slice of Schedule.
type Schedules []*Schedule
