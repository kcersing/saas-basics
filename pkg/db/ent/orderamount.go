// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/orderamount"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderAmount is the model entity for the OrderAmount schema.
type OrderAmount struct {
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
	// 订单id
	OrderID int64 `json:"order_id,omitempty"`
	// 总金额
	Total float64 `json:"total,omitempty"`
	// 实际已付款
	Actual float64 `json:"actual,omitempty"`
	// 未支付金额
	Residue float64 `json:"residue,omitempty"`
	// 减免
	Remission float64 `json:"remission,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderAmountQuery when eager-loading is set.
	Edges        OrderAmountEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderAmountEdges holds the relations/edges for other nodes in the graph.
type OrderAmountEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderAmountEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderAmount) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderamount.FieldTotal, orderamount.FieldActual, orderamount.FieldResidue, orderamount.FieldRemission:
			values[i] = new(sql.NullFloat64)
		case orderamount.FieldID, orderamount.FieldCreatedID, orderamount.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case orderamount.FieldCreatedAt, orderamount.FieldUpdatedAt, orderamount.FieldDeleteAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderAmount fields.
func (oa *OrderAmount) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderamount.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oa.ID = int64(value.Int64)
		case orderamount.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oa.CreatedAt = value.Time
			}
		case orderamount.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oa.UpdatedAt = value.Time
			}
		case orderamount.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				oa.DeleteAt = value.Time
			}
		case orderamount.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				oa.CreatedID = value.Int64
			}
		case orderamount.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oa.OrderID = value.Int64
			}
		case orderamount.FieldTotal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value.Valid {
				oa.Total = value.Float64
			}
		case orderamount.FieldActual:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field actual", values[i])
			} else if value.Valid {
				oa.Actual = value.Float64
			}
		case orderamount.FieldResidue:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field residue", values[i])
			} else if value.Valid {
				oa.Residue = value.Float64
			}
		case orderamount.FieldRemission:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field remission", values[i])
			} else if value.Valid {
				oa.Remission = value.Float64
			}
		default:
			oa.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderAmount.
// This includes values selected through modifiers, order, etc.
func (oa *OrderAmount) Value(name string) (ent.Value, error) {
	return oa.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the OrderAmount entity.
func (oa *OrderAmount) QueryOrder() *OrderQuery {
	return NewOrderAmountClient(oa.config).QueryOrder(oa)
}

// Update returns a builder for updating this OrderAmount.
// Note that you need to call OrderAmount.Unwrap() before calling this method if this OrderAmount
// was returned from a transaction, and the transaction was committed or rolled back.
func (oa *OrderAmount) Update() *OrderAmountUpdateOne {
	return NewOrderAmountClient(oa.config).UpdateOne(oa)
}

// Unwrap unwraps the OrderAmount entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oa *OrderAmount) Unwrap() *OrderAmount {
	_tx, ok := oa.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderAmount is not a transactional entity")
	}
	oa.config.driver = _tx.drv
	return oa
}

// String implements the fmt.Stringer.
func (oa *OrderAmount) String() string {
	var builder strings.Builder
	builder.WriteString("OrderAmount(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oa.ID))
	builder.WriteString("created_at=")
	builder.WriteString(oa.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oa.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete_at=")
	builder.WriteString(oa.DeleteAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", oa.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", oa.OrderID))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", oa.Total))
	builder.WriteString(", ")
	builder.WriteString("actual=")
	builder.WriteString(fmt.Sprintf("%v", oa.Actual))
	builder.WriteString(", ")
	builder.WriteString("residue=")
	builder.WriteString(fmt.Sprintf("%v", oa.Residue))
	builder.WriteString(", ")
	builder.WriteString("remission=")
	builder.WriteString(fmt.Sprintf("%v", oa.Remission))
	builder.WriteByte(')')
	return builder.String()
}

// OrderAmounts is a parsable slice of OrderAmount.
type OrderAmounts []*OrderAmount
