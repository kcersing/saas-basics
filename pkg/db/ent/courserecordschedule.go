// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/courserecordschedule"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CourseRecordSchedule is the model entity for the CourseRecordSchedule schema.
type CourseRecordSchedule struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 类型
	Type string `json:"type,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 场地ID
	PlaceID int64 `json:"place_id,omitempty"`
	// 教练ID
	CoachID int64 `json:"coach_id,omitempty"`
	// 上课人数
	Num int64 `json:"num,omitempty"`
	// 开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 开始时间
	EndTime time.Time `json:"end_time,omitempty"`
	// 课程价格
	Price float64 `json:"price,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseRecordScheduleQuery when eager-loading is set.
	Edges        CourseRecordScheduleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CourseRecordScheduleEdges holds the relations/edges for other nodes in the graph.
type CourseRecordScheduleEdges struct {
	// Users holds the value of the users edge.
	Users []*CourseRecordUser `json:"users,omitempty"`
	// Coach holds the value of the coach edge.
	Coach []*CourseRecordCoach `json:"coach,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e CourseRecordScheduleEdges) UsersOrErr() ([]*CourseRecordUser, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// CoachOrErr returns the Coach value or an error if the edge
// was not loaded in eager-loading.
func (e CourseRecordScheduleEdges) CoachOrErr() ([]*CourseRecordCoach, error) {
	if e.loadedTypes[1] {
		return e.Coach, nil
	}
	return nil, &NotLoadedError{edge: "coach"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CourseRecordSchedule) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case courserecordschedule.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case courserecordschedule.FieldID, courserecordschedule.FieldVenueID, courserecordschedule.FieldPlaceID, courserecordschedule.FieldCoachID, courserecordschedule.FieldNum, courserecordschedule.FieldStatus:
			values[i] = new(sql.NullInt64)
		case courserecordschedule.FieldType:
			values[i] = new(sql.NullString)
		case courserecordschedule.FieldCreatedAt, courserecordschedule.FieldUpdatedAt, courserecordschedule.FieldStartTime, courserecordschedule.FieldEndTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CourseRecordSchedule fields.
func (crs *CourseRecordSchedule) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case courserecordschedule.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			crs.ID = int64(value.Int64)
		case courserecordschedule.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				crs.CreatedAt = value.Time
			}
		case courserecordschedule.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				crs.UpdatedAt = value.Time
			}
		case courserecordschedule.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				crs.Type = value.String
			}
		case courserecordschedule.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				crs.VenueID = value.Int64
			}
		case courserecordschedule.FieldPlaceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field place_id", values[i])
			} else if value.Valid {
				crs.PlaceID = value.Int64
			}
		case courserecordschedule.FieldCoachID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coach_id", values[i])
			} else if value.Valid {
				crs.CoachID = value.Int64
			}
		case courserecordschedule.FieldNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num", values[i])
			} else if value.Valid {
				crs.Num = value.Int64
			}
		case courserecordschedule.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				crs.StartTime = value.Time
			}
		case courserecordschedule.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				crs.EndTime = value.Time
			}
		case courserecordschedule.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				crs.Price = value.Float64
			}
		case courserecordschedule.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				crs.Status = value.Int64
			}
		default:
			crs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CourseRecordSchedule.
// This includes values selected through modifiers, order, etc.
func (crs *CourseRecordSchedule) Value(name string) (ent.Value, error) {
	return crs.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the CourseRecordSchedule entity.
func (crs *CourseRecordSchedule) QueryUsers() *CourseRecordUserQuery {
	return NewCourseRecordScheduleClient(crs.config).QueryUsers(crs)
}

// QueryCoach queries the "coach" edge of the CourseRecordSchedule entity.
func (crs *CourseRecordSchedule) QueryCoach() *CourseRecordCoachQuery {
	return NewCourseRecordScheduleClient(crs.config).QueryCoach(crs)
}

// Update returns a builder for updating this CourseRecordSchedule.
// Note that you need to call CourseRecordSchedule.Unwrap() before calling this method if this CourseRecordSchedule
// was returned from a transaction, and the transaction was committed or rolled back.
func (crs *CourseRecordSchedule) Update() *CourseRecordScheduleUpdateOne {
	return NewCourseRecordScheduleClient(crs.config).UpdateOne(crs)
}

// Unwrap unwraps the CourseRecordSchedule entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (crs *CourseRecordSchedule) Unwrap() *CourseRecordSchedule {
	_tx, ok := crs.config.driver.(*txDriver)
	if !ok {
		panic("ent: CourseRecordSchedule is not a transactional entity")
	}
	crs.config.driver = _tx.drv
	return crs
}

// String implements the fmt.Stringer.
func (crs *CourseRecordSchedule) String() string {
	var builder strings.Builder
	builder.WriteString("CourseRecordSchedule(")
	builder.WriteString(fmt.Sprintf("id=%v, ", crs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(crs.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(crs.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(crs.Type)
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", crs.VenueID))
	builder.WriteString(", ")
	builder.WriteString("place_id=")
	builder.WriteString(fmt.Sprintf("%v", crs.PlaceID))
	builder.WriteString(", ")
	builder.WriteString("coach_id=")
	builder.WriteString(fmt.Sprintf("%v", crs.CoachID))
	builder.WriteString(", ")
	builder.WriteString("num=")
	builder.WriteString(fmt.Sprintf("%v", crs.Num))
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(crs.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(crs.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", crs.Price))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", crs.Status))
	builder.WriteByte(')')
	return builder.String()
}

// CourseRecordSchedules is a parsable slice of CourseRecordSchedule.
type CourseRecordSchedules []*CourseRecordSchedule
