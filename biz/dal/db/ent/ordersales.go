// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/ordersales"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderSales is the model entity for the OrderSales schema.
type OrderSales struct {
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
	// 订单id
	OrderID int64 `json:"order_id,omitempty"`
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 销售id
	SalesID int64 `json:"sales_id,omitempty"`
	// Ratio holds the value of the "ratio" field.
	Ratio int64 `json:"ratio,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderSalesQuery when eager-loading is set.
	Edges        OrderSalesEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderSalesEdges holds the relations/edges for other nodes in the graph.
type OrderSalesEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderSalesEdges) OrderOrErr() (*Order, error) {
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
func (*OrderSales) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ordersales.FieldID, ordersales.FieldDelete, ordersales.FieldCreatedID, ordersales.FieldStatus, ordersales.FieldOrderID, ordersales.FieldMemberID, ordersales.FieldSalesID, ordersales.FieldRatio:
			values[i] = new(sql.NullInt64)
		case ordersales.FieldCreatedAt, ordersales.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderSales fields.
func (os *OrderSales) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ordersales.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			os.ID = int64(value.Int64)
		case ordersales.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				os.CreatedAt = value.Time
			}
		case ordersales.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				os.UpdatedAt = value.Time
			}
		case ordersales.FieldDelete:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete", values[i])
			} else if value.Valid {
				os.Delete = value.Int64
			}
		case ordersales.FieldCreatedID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_id", values[i])
			} else if value.Valid {
				os.CreatedID = value.Int64
			}
		case ordersales.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				os.Status = value.Int64
			}
		case ordersales.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				os.OrderID = value.Int64
			}
		case ordersales.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				os.MemberID = value.Int64
			}
		case ordersales.FieldSalesID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sales_id", values[i])
			} else if value.Valid {
				os.SalesID = value.Int64
			}
		case ordersales.FieldRatio:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ratio", values[i])
			} else if value.Valid {
				os.Ratio = value.Int64
			}
		default:
			os.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderSales.
// This includes values selected through modifiers, order, etc.
func (os *OrderSales) Value(name string) (ent.Value, error) {
	return os.selectValues.Get(name)
}

// QueryOrder queries the "order" edge of the OrderSales entity.
func (os *OrderSales) QueryOrder() *OrderQuery {
	return NewOrderSalesClient(os.config).QueryOrder(os)
}

// Update returns a builder for updating this OrderSales.
// Note that you need to call OrderSales.Unwrap() before calling this method if this OrderSales
// was returned from a transaction, and the transaction was committed or rolled back.
func (os *OrderSales) Update() *OrderSalesUpdateOne {
	return NewOrderSalesClient(os.config).UpdateOne(os)
}

// Unwrap unwraps the OrderSales entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (os *OrderSales) Unwrap() *OrderSales {
	_tx, ok := os.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderSales is not a transactional entity")
	}
	os.config.driver = _tx.drv
	return os
}

// String implements the fmt.Stringer.
func (os *OrderSales) String() string {
	var builder strings.Builder
	builder.WriteString("OrderSales(")
	builder.WriteString(fmt.Sprintf("id=%v, ", os.ID))
	builder.WriteString("created_at=")
	builder.WriteString(os.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(os.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("delete=")
	builder.WriteString(fmt.Sprintf("%v", os.Delete))
	builder.WriteString(", ")
	builder.WriteString("created_id=")
	builder.WriteString(fmt.Sprintf("%v", os.CreatedID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", os.Status))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", os.OrderID))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", os.MemberID))
	builder.WriteString(", ")
	builder.WriteString("sales_id=")
	builder.WriteString(fmt.Sprintf("%v", os.SalesID))
	builder.WriteString(", ")
	builder.WriteString("ratio=")
	builder.WriteString(fmt.Sprintf("%v", os.Ratio))
	builder.WriteByte(')')
	return builder.String()
}

// OrderSalesSlice is a parsable slice of OrderSales.
type OrderSalesSlice []*OrderSales
