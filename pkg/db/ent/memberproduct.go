// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/memberproduct"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// MemberProduct is the model entity for the MemberProduct schema.
type MemberProduct struct {
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
	// 编号
	Sn string `json:"sn,omitempty"`
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 产品ID
	ProductID int64 `json:"product_id,omitempty"`
	// 场馆ID
	VenueID int64 `json:"venue_id,omitempty"`
	// 订单ID
	OrderID int64 `json:"order_id,omitempty"`
	// 产品名称
	Name string `json:"name,omitempty"`
	// 产品价格
	Price float64 `json:"price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MemberProductQuery when eager-loading is set.
	Edges        MemberProductEdges `json:"edges"`
	selectValues sql.SelectValues
}

// MemberProductEdges holds the relations/edges for other nodes in the graph.
type MemberProductEdges struct {
	// Members holds the value of the members edge.
	Members *Member `json:"members,omitempty"`
	// MemberProductPropertys holds the value of the member_product_propertys edge.
	MemberProductPropertys []*MemberProductProperty `json:"member_product_propertys,omitempty"`
	// MemberProductEntry holds the value of the member_product_entry edge.
	MemberProductEntry []*EntryLogs `json:"member_product_entry,omitempty"`
	// MemberProductContents holds the value of the member_product_contents edge.
	MemberProductContents []*MemberContract `json:"member_product_contents,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MemberProductEdges) MembersOrErr() (*Member, error) {
	if e.loadedTypes[0] {
		if e.Members == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: member.Label}
		}
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// MemberProductPropertysOrErr returns the MemberProductPropertys value or an error if the edge
// was not loaded in eager-loading.
func (e MemberProductEdges) MemberProductPropertysOrErr() ([]*MemberProductProperty, error) {
	if e.loadedTypes[1] {
		return e.MemberProductPropertys, nil
	}
	return nil, &NotLoadedError{edge: "member_product_propertys"}
}

// MemberProductEntryOrErr returns the MemberProductEntry value or an error if the edge
// was not loaded in eager-loading.
func (e MemberProductEdges) MemberProductEntryOrErr() ([]*EntryLogs, error) {
	if e.loadedTypes[2] {
		return e.MemberProductEntry, nil
	}
	return nil, &NotLoadedError{edge: "member_product_entry"}
}

// MemberProductContentsOrErr returns the MemberProductContents value or an error if the edge
// was not loaded in eager-loading.
func (e MemberProductEdges) MemberProductContentsOrErr() ([]*MemberContract, error) {
	if e.loadedTypes[3] {
		return e.MemberProductContents, nil
	}
	return nil, &NotLoadedError{edge: "member_product_contents"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MemberProduct) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case memberproduct.FieldPrice:
			values[i] = new(sql.NullFloat64)
		case memberproduct.FieldID, memberproduct.FieldStatus, memberproduct.FieldMemberID, memberproduct.FieldProductID, memberproduct.FieldVenueID, memberproduct.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case memberproduct.FieldSn, memberproduct.FieldName:
			values[i] = new(sql.NullString)
		case memberproduct.FieldCreatedAt, memberproduct.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MemberProduct fields.
func (mp *MemberProduct) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case memberproduct.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mp.ID = int64(value.Int64)
		case memberproduct.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mp.CreatedAt = value.Time
			}
		case memberproduct.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mp.UpdatedAt = value.Time
			}
		case memberproduct.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				mp.Status = value.Int64
			}
		case memberproduct.FieldSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sn", values[i])
			} else if value.Valid {
				mp.Sn = value.String
			}
		case memberproduct.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				mp.MemberID = value.Int64
			}
		case memberproduct.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				mp.ProductID = value.Int64
			}
		case memberproduct.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				mp.VenueID = value.Int64
			}
		case memberproduct.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				mp.OrderID = value.Int64
			}
		case memberproduct.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				mp.Name = value.String
			}
		case memberproduct.FieldPrice:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				mp.Price = value.Float64
			}
		default:
			mp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the MemberProduct.
// This includes values selected through modifiers, order, etc.
func (mp *MemberProduct) Value(name string) (ent.Value, error) {
	return mp.selectValues.Get(name)
}

// QueryMembers queries the "members" edge of the MemberProduct entity.
func (mp *MemberProduct) QueryMembers() *MemberQuery {
	return NewMemberProductClient(mp.config).QueryMembers(mp)
}

// QueryMemberProductPropertys queries the "member_product_propertys" edge of the MemberProduct entity.
func (mp *MemberProduct) QueryMemberProductPropertys() *MemberProductPropertyQuery {
	return NewMemberProductClient(mp.config).QueryMemberProductPropertys(mp)
}

// QueryMemberProductEntry queries the "member_product_entry" edge of the MemberProduct entity.
func (mp *MemberProduct) QueryMemberProductEntry() *EntryLogsQuery {
	return NewMemberProductClient(mp.config).QueryMemberProductEntry(mp)
}

// QueryMemberProductContents queries the "member_product_contents" edge of the MemberProduct entity.
func (mp *MemberProduct) QueryMemberProductContents() *MemberContractQuery {
	return NewMemberProductClient(mp.config).QueryMemberProductContents(mp)
}

// Update returns a builder for updating this MemberProduct.
// Note that you need to call MemberProduct.Unwrap() before calling this method if this MemberProduct
// was returned from a transaction, and the transaction was committed or rolled back.
func (mp *MemberProduct) Update() *MemberProductUpdateOne {
	return NewMemberProductClient(mp.config).UpdateOne(mp)
}

// Unwrap unwraps the MemberProduct entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mp *MemberProduct) Unwrap() *MemberProduct {
	_tx, ok := mp.config.driver.(*txDriver)
	if !ok {
		panic("ent: MemberProduct is not a transactional entity")
	}
	mp.config.driver = _tx.drv
	return mp
}

// String implements the fmt.Stringer.
func (mp *MemberProduct) String() string {
	var builder strings.Builder
	builder.WriteString("MemberProduct(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(mp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(mp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", mp.Status))
	builder.WriteString(", ")
	builder.WriteString("sn=")
	builder.WriteString(mp.Sn)
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.MemberID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.ProductID))
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.VenueID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", mp.OrderID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(mp.Name)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", mp.Price))
	builder.WriteByte(')')
	return builder.String()
}

// MemberProducts is a parsable slice of MemberProduct.
type MemberProducts []*MemberProduct
