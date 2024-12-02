// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/contest"
	"saas/pkg/db/ent/contestparticipant"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ContestParticipant is the model entity for the ContestParticipant schema.
type ContestParticipant struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 状态[0:禁用;1:正常]
	Status int64 `json:"status,omitempty"`
	// 赛事id
	ContestID int64 `json:"contest_id,omitempty"`
	// 名称
	Name string `json:"name,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 更多
	Fields string `json:"fields,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContestParticipantQuery when eager-loading is set.
	Edges        ContestParticipantEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ContestParticipantEdges holds the relations/edges for other nodes in the graph.
type ContestParticipantEdges struct {
	// Contest holds the value of the contest edge.
	Contest *Contest `json:"contest,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ContestOrErr returns the Contest value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContestParticipantEdges) ContestOrErr() (*Contest, error) {
	if e.loadedTypes[0] {
		if e.Contest == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: contest.Label}
		}
		return e.Contest, nil
	}
	return nil, &NotLoadedError{edge: "contest"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ContestParticipant) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case contestparticipant.FieldID, contestparticipant.FieldStatus, contestparticipant.FieldContestID:
			values[i] = new(sql.NullInt64)
		case contestparticipant.FieldName, contestparticipant.FieldMobile, contestparticipant.FieldFields:
			values[i] = new(sql.NullString)
		case contestparticipant.FieldCreatedAt, contestparticipant.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ContestParticipant fields.
func (cp *ContestParticipant) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case contestparticipant.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			cp.ID = int64(value.Int64)
		case contestparticipant.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cp.CreatedAt = value.Time
			}
		case contestparticipant.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cp.UpdatedAt = value.Time
			}
		case contestparticipant.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				cp.Status = value.Int64
			}
		case contestparticipant.FieldContestID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contest_id", values[i])
			} else if value.Valid {
				cp.ContestID = value.Int64
			}
		case contestparticipant.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cp.Name = value.String
			}
		case contestparticipant.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				cp.Mobile = value.String
			}
		case contestparticipant.FieldFields:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fields", values[i])
			} else if value.Valid {
				cp.Fields = value.String
			}
		default:
			cp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ContestParticipant.
// This includes values selected through modifiers, order, etc.
func (cp *ContestParticipant) Value(name string) (ent.Value, error) {
	return cp.selectValues.Get(name)
}

// QueryContest queries the "contest" edge of the ContestParticipant entity.
func (cp *ContestParticipant) QueryContest() *ContestQuery {
	return NewContestParticipantClient(cp.config).QueryContest(cp)
}

// Update returns a builder for updating this ContestParticipant.
// Note that you need to call ContestParticipant.Unwrap() before calling this method if this ContestParticipant
// was returned from a transaction, and the transaction was committed or rolled back.
func (cp *ContestParticipant) Update() *ContestParticipantUpdateOne {
	return NewContestParticipantClient(cp.config).UpdateOne(cp)
}

// Unwrap unwraps the ContestParticipant entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cp *ContestParticipant) Unwrap() *ContestParticipant {
	_tx, ok := cp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ContestParticipant is not a transactional entity")
	}
	cp.config.driver = _tx.drv
	return cp
}

// String implements the fmt.Stringer.
func (cp *ContestParticipant) String() string {
	var builder strings.Builder
	builder.WriteString("ContestParticipant(")
	builder.WriteString(fmt.Sprintf("id=%v, ", cp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(cp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(cp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", cp.Status))
	builder.WriteString(", ")
	builder.WriteString("contest_id=")
	builder.WriteString(fmt.Sprintf("%v", cp.ContestID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(cp.Name)
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(cp.Mobile)
	builder.WriteString(", ")
	builder.WriteString("fields=")
	builder.WriteString(cp.Fields)
	builder.WriteByte(')')
	return builder.String()
}

// ContestParticipants is a parsable slice of ContestParticipant.
type ContestParticipants []*ContestParticipant
