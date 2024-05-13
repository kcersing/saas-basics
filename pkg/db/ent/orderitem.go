// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/orderitem"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// OrderItem is the model entity for the OrderItem schema.
type OrderItem struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 订单id
	OrderID int64 `json:"order_id,omitempty"`
	// 产品id
	ProductID int64 `json:"product_id,omitempty"`
	// 购买数量
	Quantity int64 `json:"quantity,omitempty"`
	// 关联会员产品id
	RelatedUserProductID int64 `json:"related_user_product_id,omitempty"`
	// 合同ID
	ContractID int64 `json:"contract_id,omitempty"`
	// 指定时间
	AssignAt time.Time `json:"assign_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderItemQuery when eager-loading is set.
	Edges        OrderItemEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderItemEdges holds the relations/edges for other nodes in the graph.
type OrderItemEdges struct {
	// Aufk holds the value of the aufk edge.
	Aufk *Order `json:"aufk,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AufkOrErr returns the Aufk value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderItemEdges) AufkOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Aufk == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Aufk, nil
	}
	return nil, &NotLoadedError{edge: "aufk"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderItem) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID, orderitem.FieldOrderID, orderitem.FieldProductID, orderitem.FieldQuantity, orderitem.FieldRelatedUserProductID, orderitem.FieldContractID:
			values[i] = new(sql.NullInt64)
		case orderitem.FieldCreatedAt, orderitem.FieldUpdatedAt, orderitem.FieldAssignAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderItem fields.
func (oi *OrderItem) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderitem.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oi.ID = int64(value.Int64)
		case orderitem.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oi.CreatedAt = value.Time
			}
		case orderitem.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oi.UpdatedAt = value.Time
			}
		case orderitem.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oi.OrderID = value.Int64
			}
		case orderitem.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				oi.ProductID = value.Int64
			}
		case orderitem.FieldQuantity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field quantity", values[i])
			} else if value.Valid {
				oi.Quantity = value.Int64
			}
		case orderitem.FieldRelatedUserProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field related_user_product_id", values[i])
			} else if value.Valid {
				oi.RelatedUserProductID = value.Int64
			}
		case orderitem.FieldContractID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field contract_id", values[i])
			} else if value.Valid {
				oi.ContractID = value.Int64
			}
		case orderitem.FieldAssignAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field assign_at", values[i])
			} else if value.Valid {
				oi.AssignAt = value.Time
			}
		default:
			oi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderItem.
// This includes values selected through modifiers, order, etc.
func (oi *OrderItem) Value(name string) (ent.Value, error) {
	return oi.selectValues.Get(name)
}

// QueryAufk queries the "aufk" edge of the OrderItem entity.
func (oi *OrderItem) QueryAufk() *OrderQuery {
	return NewOrderItemClient(oi.config).QueryAufk(oi)
}

// Update returns a builder for updating this OrderItem.
// Note that you need to call OrderItem.Unwrap() before calling this method if this OrderItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oi *OrderItem) Update() *OrderItemUpdateOne {
	return NewOrderItemClient(oi.config).UpdateOne(oi)
}

// Unwrap unwraps the OrderItem entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oi *OrderItem) Unwrap() *OrderItem {
	_tx, ok := oi.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderItem is not a transactional entity")
	}
	oi.config.driver = _tx.drv
	return oi
}

// String implements the fmt.Stringer.
func (oi *OrderItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderItem(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oi.ID))
	builder.WriteString("created_at=")
	builder.WriteString(oi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.OrderID))
	builder.WriteString(", ")
	builder.WriteString("product_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ProductID))
	builder.WriteString(", ")
	builder.WriteString("quantity=")
	builder.WriteString(fmt.Sprintf("%v", oi.Quantity))
	builder.WriteString(", ")
	builder.WriteString("related_user_product_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.RelatedUserProductID))
	builder.WriteString(", ")
	builder.WriteString("contract_id=")
	builder.WriteString(fmt.Sprintf("%v", oi.ContractID))
	builder.WriteString(", ")
	builder.WriteString("assign_at=")
	builder.WriteString(oi.AssignAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// OrderItems is a parsable slice of OrderItem.
type OrderItems []*OrderItem
