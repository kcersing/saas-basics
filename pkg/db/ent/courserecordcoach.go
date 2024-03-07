// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/courserecordcoach"
	"saas/pkg/db/ent/courserecordschedule"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// CourseRecordCoach is the model entity for the CourseRecordCoach schema.
type CourseRecordCoach struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 教练ID
	CoachID int64 `json:"coach_id,omitempty"`
	// 课程ID
	CourseRecordScheduleID int64 `json:"course_record_schedule_id,omitempty"`
	// 类型
	Type string `json:"type,omitempty"`
	// 开始时间
	StartTime time.Time `json:"start_time,omitempty"`
	// 开始时间
	EndTime time.Time `json:"end_time,omitempty"`
	// 上课签到时间
	SignStartTime time.Time `json:"sign_start_time,omitempty"`
	// 下课签到时间
	SignNdTime time.Time `json:"sign_nd_time,omitempty"`
	// 状态
	Status int64 `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CourseRecordCoachQuery when eager-loading is set.
	Edges        CourseRecordCoachEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CourseRecordCoachEdges holds the relations/edges for other nodes in the graph.
type CourseRecordCoachEdges struct {
	// Schedule holds the value of the schedule edge.
	Schedule *CourseRecordSchedule `json:"schedule,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ScheduleOrErr returns the Schedule value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CourseRecordCoachEdges) ScheduleOrErr() (*CourseRecordSchedule, error) {
	if e.loadedTypes[0] {
		if e.Schedule == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: courserecordschedule.Label}
		}
		return e.Schedule, nil
	}
	return nil, &NotLoadedError{edge: "schedule"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CourseRecordCoach) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case courserecordcoach.FieldID, courserecordcoach.FieldVenueID, courserecordcoach.FieldCoachID, courserecordcoach.FieldCourseRecordScheduleID, courserecordcoach.FieldStatus:
			values[i] = new(sql.NullInt64)
		case courserecordcoach.FieldType:
			values[i] = new(sql.NullString)
		case courserecordcoach.FieldCreatedAt, courserecordcoach.FieldUpdatedAt, courserecordcoach.FieldStartTime, courserecordcoach.FieldEndTime, courserecordcoach.FieldSignStartTime, courserecordcoach.FieldSignNdTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CourseRecordCoach fields.
func (crc *CourseRecordCoach) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case courserecordcoach.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			crc.ID = int64(value.Int64)
		case courserecordcoach.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				crc.CreatedAt = value.Time
			}
		case courserecordcoach.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				crc.UpdatedAt = value.Time
			}
		case courserecordcoach.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				crc.VenueID = value.Int64
			}
		case courserecordcoach.FieldCoachID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field coach_id", values[i])
			} else if value.Valid {
				crc.CoachID = value.Int64
			}
		case courserecordcoach.FieldCourseRecordScheduleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field course_record_schedule_id", values[i])
			} else if value.Valid {
				crc.CourseRecordScheduleID = value.Int64
			}
		case courserecordcoach.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				crc.Type = value.String
			}
		case courserecordcoach.FieldStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field start_time", values[i])
			} else if value.Valid {
				crc.StartTime = value.Time
			}
		case courserecordcoach.FieldEndTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field end_time", values[i])
			} else if value.Valid {
				crc.EndTime = value.Time
			}
		case courserecordcoach.FieldSignStartTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_start_time", values[i])
			} else if value.Valid {
				crc.SignStartTime = value.Time
			}
		case courserecordcoach.FieldSignNdTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field sign_nd_time", values[i])
			} else if value.Valid {
				crc.SignNdTime = value.Time
			}
		case courserecordcoach.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				crc.Status = value.Int64
			}
		default:
			crc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CourseRecordCoach.
// This includes values selected through modifiers, order, etc.
func (crc *CourseRecordCoach) Value(name string) (ent.Value, error) {
	return crc.selectValues.Get(name)
}

// QuerySchedule queries the "schedule" edge of the CourseRecordCoach entity.
func (crc *CourseRecordCoach) QuerySchedule() *CourseRecordScheduleQuery {
	return NewCourseRecordCoachClient(crc.config).QuerySchedule(crc)
}

// Update returns a builder for updating this CourseRecordCoach.
// Note that you need to call CourseRecordCoach.Unwrap() before calling this method if this CourseRecordCoach
// was returned from a transaction, and the transaction was committed or rolled back.
func (crc *CourseRecordCoach) Update() *CourseRecordCoachUpdateOne {
	return NewCourseRecordCoachClient(crc.config).UpdateOne(crc)
}

// Unwrap unwraps the CourseRecordCoach entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (crc *CourseRecordCoach) Unwrap() *CourseRecordCoach {
	_tx, ok := crc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CourseRecordCoach is not a transactional entity")
	}
	crc.config.driver = _tx.drv
	return crc
}

// String implements the fmt.Stringer.
func (crc *CourseRecordCoach) String() string {
	var builder strings.Builder
	builder.WriteString("CourseRecordCoach(")
	builder.WriteString(fmt.Sprintf("id=%v, ", crc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(crc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(crc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", crc.VenueID))
	builder.WriteString(", ")
	builder.WriteString("coach_id=")
	builder.WriteString(fmt.Sprintf("%v", crc.CoachID))
	builder.WriteString(", ")
	builder.WriteString("course_record_schedule_id=")
	builder.WriteString(fmt.Sprintf("%v", crc.CourseRecordScheduleID))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(crc.Type)
	builder.WriteString(", ")
	builder.WriteString("start_time=")
	builder.WriteString(crc.StartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("end_time=")
	builder.WriteString(crc.EndTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sign_start_time=")
	builder.WriteString(crc.SignStartTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("sign_nd_time=")
	builder.WriteString(crc.SignNdTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", crc.Status))
	builder.WriteByte(')')
	return builder.String()
}

// CourseRecordCoaches is a parsable slice of CourseRecordCoach.
type CourseRecordCoaches []*CourseRecordCoach
