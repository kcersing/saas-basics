// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberproductcourses"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MemberProductCourses is the model entity for the MemberProductCourses schema.
type MemberProductCourses struct {
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
	// 课名
	Name string `json:"name,omitempty"`
	// 节数
	Number int64 `json:"number,omitempty"`
	// 产品名称
	MemberProductID int64 `json:"member_product_id,omitempty"`
	// 课名称
	CoursesID int64 `json:"courses_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberProductCoursesQuery when eager-loading is set.
	Edges        MemberProductCoursesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberProductCoursesEdges holds the relations/edges for other nodes in the graph.
type MemberProductCoursesEdges struct {
	// NodeC holds the value of the nodeC edge.
	NodeC *MemberProduct `json:"nodeC,omitempty"`
	// NodeL holds the value of the nodeL edge.
	NodeL *MemberProduct `json:"nodeL,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// NodeCOrErr returns the NodeC value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberProductCoursesEdges) NodeCOrErr() (*MemberProduct, error) {
	if e.loadedTypes[0] {
		if e.NodeC == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: memberproduct.Label}
		}
		return e.NodeC, nil
	}
	return nil, &NotLoadedError{edge: "nodeC"}
}

// NodeLOrErr returns the NodeL value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberProductCoursesEdges) NodeLOrErr() (*MemberProduct, error) {
	if e.loadedTypes[1] {
		if e.NodeL == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: memberproduct.Label}
		}
		return e.NodeL, nil
	}
	return nil, &NotLoadedError{edge: "nodeL"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberProductCourses) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case memberproductcourses.FieldID, memberproductcourses.FieldDelete, memberproductcourses.FieldCreatedID, memberproductcourses.FieldStatus, memberproductcourses.FieldNumber, memberproductcourses.FieldMemberProductID, memberproductcourses.FieldCoursesID:
			values[i] = new(sql.NullInt64)
		case memberproductcourses.FieldType, memberproductcourses.FieldName:
			values[i] = new(sql.NullString)
		case memberproductcourses.FieldCreatedAt, memberproductcourses.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberProductCourses fields.
func (mpc *MemberProductCourses) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case memberproductcourses.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mpc.ID = int64(value.Int64)
		case memberproductcourses.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mpc.CreatedAt = value.Time
			}
		case memberproductcourses.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mpc.UpdatedAt = value.Time
			}
		case memberproductcourses.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				mpc.Delete = value.Int64
			}
		case memberproductcourses.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				mpc.CreatedID = value.Int64
			}
		case memberproductcourses.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mpc.Status = value.Int64
			}
		case memberproductcourses.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				mpc.Type = value.String
			}
		case memberproductcourses.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mpc.Name = value.String
			}
		case memberproductcourses.FieldNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field number", values[i])
			} else if value.Valid {
				mpc.Number = value.Int64
			}
		case memberproductcourses.FieldMemberProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_product_id", values[i])
			} else if value.Valid {
				mpc.MemberProductID = value.Int64
			}
		case memberproductcourses.FieldCoursesID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field courses_id", values[i])
			} else if value.Valid {
				mpc.CoursesID = value.Int64
			}
		default:
			mpc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberProductCourses.
// This includes values selected through modifiers, order, etc.
func (mpc *MemberProductCourses) Value(name string) (ent.Value, error) {
	return mpc.selectValues.Get(name)
}

// QueryNodeC queries the "nodeC" edge of the MemberProductCourses entity.
func (mpc *MemberProductCourses) QueryNodeC() *MemberProductQuery {
	return NewMemberProductCoursesClient(mpc.config).QueryNodeC(mpc)
}

// QueryNodeL queries the "nodeL" edge of the MemberProductCourses entity.
func (mpc *MemberProductCourses) QueryNodeL() *MemberProductQuery {
	return NewMemberProductCoursesClient(mpc.config).QueryNodeL(mpc)
}

// Update returns a builder for updating this MemberProductCourses.
// Note that you need to call MemberProductCourses.Unwrap() before calling this method if this MemberProductCourses
// was returned from a transaction, and the transaction was committed or rolled back.
func (mpc *MemberProductCourses) Update() *MemberProductCoursesUpdateOne {
	return NewMemberProductCoursesClient(mpc.config).UpdateOne(mpc)
}

// Unwrap unwraps the MemberProductCourses entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mpc *MemberProductCourses) Unwrap() *MemberProductCourses {
	_tx, ok := mpc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MemberProductCourses is not a transactional entity")
	}
	mpc.config.driver = _tx.drv
	return mpc
}

// String implements the fmt.Stringer.
func (mpc *MemberProductCourses) String() string {
	var builder strings.Builder
	builder.WriteString("MemberProductCourses(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mpc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mpc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mpc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", mpc.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", mpc.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mpc.Status))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(mpc.Type)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(mpc.Name)
	builder.WriteString(", ")
	builder.WriteString("number=")
	builder.WriteString(fmt.Sprintf("%v", mpc.Number))
	builder.WriteString(", ")
	builder.WriteString("member_product_id=")
	builder.WriteString(fmt.Sprintf("%v", mpc.MemberProductID))
	builder.WriteString(", ")
	builder.WriteString("courses_id=")
	builder.WriteString(fmt.Sprintf("%v", mpc.CoursesID))
	builder.WriteByte(')')
	return builder.String()
}

// MemberProductCoursesSlice is a parsable slice of MemberProductCourses.
type MemberProductCoursesSlice []*MemberProductCourses
