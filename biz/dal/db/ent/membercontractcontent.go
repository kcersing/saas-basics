// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/biz/dal/db/ent/membercontract"
	"saas/biz/dal/db/ent/membercontractcontent"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MemberContractContent is the model entity for the MemberContractContent schema.
type MemberContractContent struct {
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
	// 合同ID
	MemberContractID int64 `json:"member_contract_id,omitempty"`
	// content | 内容
	Content string `json:"content,omitempty"`
	// sign_img | 会员签字b64 预处理
	SignImg string `json:"sign_img,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberContractContentQuery when eager-loading is set.
	Edges        MemberContractContentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberContractContentEdges holds the relations/edges for other nodes in the graph.
type MemberContractContentEdges struct {
	// Contract holds the value of the contract edge.
	Contract *MemberContract `json:"contract,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ContractOrErr returns the Contract value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberContractContentEdges) ContractOrErr() (*MemberContract, error) {
	if e.loadedTypes[0] {
		if e.Contract == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: membercontract.Label}
		}
		return e.Contract, nil
	}
	return nil, &NotLoadedError{edge: "contract"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberContractContent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case membercontractcontent.FieldID, membercontractcontent.FieldDelete, membercontractcontent.FieldCreatedID, membercontractcontent.FieldMemberContractID:
			values[i] = new(sql.NullInt64)
		case membercontractcontent.FieldContent, membercontractcontent.FieldSignImg:
			values[i] = new(sql.NullString)
		case membercontractcontent.FieldCreatedAt, membercontractcontent.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberContractContent fields.
func (mcc *MemberContractContent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case membercontractcontent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mcc.ID = int64(value.Int64)
		case membercontractcontent.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mcc.CreatedAt = value.Time
			}
		case membercontractcontent.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mcc.UpdatedAt = value.Time
			}
		case membercontractcontent.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				mcc.Delete = value.Int64
			}
		case membercontractcontent.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				mcc.CreatedID = value.Int64
			}
		case membercontractcontent.FieldMemberContractID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_contract_id", values[i])
			} else if value.Valid {
				mcc.MemberContractID = value.Int64
			}
		case membercontractcontent.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				mcc.Content = value.String
			}
		case membercontractcontent.FieldSignImg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sign_img", values[i])
			} else if value.Valid {
				mcc.SignImg = value.String
			}
		default:
			mcc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberContractContent.
// This includes values selected through modifiers, order, etc.
func (mcc *MemberContractContent) Value(name string) (ent.Value, error) {
	return mcc.selectValues.Get(name)
}

// QueryContract queries the "contract" edge of the MemberContractContent entity.
func (mcc *MemberContractContent) QueryContract() *MemberContractQuery {
	return NewMemberContractContentClient(mcc.config).QueryContract(mcc)
}

// Update returns a builder for updating this MemberContractContent.
// Note that you need to call MemberContractContent.Unwrap() before calling this method if this MemberContractContent
// was returned from a transaction, and the transaction was committed or rolled back.
func (mcc *MemberContractContent) Update() *MemberContractContentUpdateOne {
	return NewMemberContractContentClient(mcc.config).UpdateOne(mcc)
}

// Unwrap unwraps the MemberContractContent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mcc *MemberContractContent) Unwrap() *MemberContractContent {
	_tx, ok := mcc.config.driver.(*txDriver)
	if !ok {
		panic("ent: MemberContractContent is not a transactional entity")
	}
	mcc.config.driver = _tx.drv
	return mcc
}

// String implements the fmt.Stringer.
func (mcc *MemberContractContent) String() string {
	var builder strings.Builder
	builder.WriteString("MemberContractContent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mcc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mcc.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mcc.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", mcc.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", mcc.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("member_contract_id=")
	builder.WriteString(fmt.Sprintf("%v", mcc.MemberContractID))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(mcc.Content)
	builder.WriteString(", ")
	builder.WriteString("sign_img=")
	builder.WriteString(mcc.SignImg)
	builder.WriteByte(')')
	return builder.String()
}

// MemberContractContents is a parsable slice of MemberContractContent.
type MemberContractContents []*MemberContractContent
