// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"order/biz/dal/mysql/ent/order"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 订单编号
	OrderSn string `json:"order_sn,omitempty"`
	// 场馆id
	VenueID int64 `json:"venue_id,omitempty"`
	// 会员id
	MemberID int64 `json:"member_id,omitempty"`
	// 会员产品id
	MemberProductID int64 `json:"member_product_id,omitempty"`
	// 状态 | [0:正常;1:禁用]
	Status int64 `json:"status,omitempty"`
	// 订单来源
	Source string `json:"source,omitempty"`
	// 设备来源
	Device string `json:"device,omitempty"`
	// 业务类型
	Nature int64 `json:"nature,omitempty"`
	// 订单完成时间
	CompletionAt time.Time `json:"completion_at,omitempty"`
	// 创建人id
	CreateID int64 `json:"create_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges        OrderEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Amount holds the value of the amount edge.
	Amount []*OrderAmount `json:"amount,omitempty"`
	// Item holds the value of the item edge.
	Item []*OrderItem `json:"item,omitempty"`
	// Pay holds the value of the pay edge.
	Pay []*OrderPay `json:"pay,omitempty"`
	// Sales holds the value of the sales edge.
	Sales []*OrderSales `json:"sales,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// AmountOrErr returns the Amount value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) AmountOrErr() ([]*OrderAmount, error) {
	if e.loadedTypes[0] {
		return e.Amount, nil
	}
	return nil, &NotLoadedError{edge: "amount"}
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) ItemOrErr() ([]*OrderItem, error) {
	if e.loadedTypes[1] {
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// PayOrErr returns the Pay value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) PayOrErr() ([]*OrderPay, error) {
	if e.loadedTypes[2] {
		return e.Pay, nil
	}
	return nil, &NotLoadedError{edge: "pay"}
}

// SalesOrErr returns the Sales value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) SalesOrErr() ([]*OrderSales, error) {
	if e.loadedTypes[3] {
		return e.Sales, nil
	}
	return nil, &NotLoadedError{edge: "sales"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldVenueID, order.FieldMemberID, order.FieldMemberProductID, order.FieldStatus, order.FieldNature, order.FieldCreateID:
			values[i] = new(sql.NullInt64)
		case order.FieldOrderSn, order.FieldSource, order.FieldDevice:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdatedAt, order.FieldCompletionAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int64(value.Int64)
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.FieldOrderSn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field order_sn", values[i])
			} else if value.Valid {
				o.OrderSn = value.String
			}
		case order.FieldVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field venue_id", values[i])
			} else if value.Valid {
				o.VenueID = value.Int64
			}
		case order.FieldMemberID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_id", values[i])
			} else if value.Valid {
				o.MemberID = value.Int64
			}
		case order.FieldMemberProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field member_product_id", values[i])
			} else if value.Valid {
				o.MemberProductID = value.Int64
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = value.Int64
			}
		case order.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				o.Source = value.String
			}
		case order.FieldDevice:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field device", values[i])
			} else if value.Valid {
				o.Device = value.String
			}
		case order.FieldNature:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nature", values[i])
			} else if value.Valid {
				o.Nature = value.Int64
			}
		case order.FieldCompletionAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field completion_at", values[i])
			} else if value.Valid {
				o.CompletionAt = value.Time
			}
		case order.FieldCreateID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_id", values[i])
			} else if value.Valid {
				o.CreateID = value.Int64
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryAmount queries the "amount" edge of the Order entity.
func (o *Order) QueryAmount() *OrderAmountQuery {
	return NewOrderClient(o.config).QueryAmount(o)
}

// QueryItem queries the "item" edge of the Order entity.
func (o *Order) QueryItem() *OrderItemQuery {
	return NewOrderClient(o.config).QueryItem(o)
}

// QueryPay queries the "pay" edge of the Order entity.
func (o *Order) QueryPay() *OrderPayQuery {
	return NewOrderClient(o.config).QueryPay(o)
}

// QuerySales queries the "sales" edge of the Order entity.
func (o *Order) QuerySales() *OrderSalesQuery {
	return NewOrderClient(o.config).QuerySales(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("order_sn=")
	builder.WriteString(o.OrderSn)
	builder.WriteString(", ")
	builder.WriteString("venue_id=")
	builder.WriteString(fmt.Sprintf("%v", o.VenueID))
	builder.WriteString(", ")
	builder.WriteString("member_id=")
	builder.WriteString(fmt.Sprintf("%v", o.MemberID))
	builder.WriteString(", ")
	builder.WriteString("member_product_id=")
	builder.WriteString(fmt.Sprintf("%v", o.MemberProductID))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", ")
	builder.WriteString("source=")
	builder.WriteString(o.Source)
	builder.WriteString(", ")
	builder.WriteString("device=")
	builder.WriteString(o.Device)
	builder.WriteString(", ")
	builder.WriteString("nature=")
	builder.WriteString(fmt.Sprintf("%v", o.Nature))
	builder.WriteString(", ")
	builder.WriteString("completion_at=")
	builder.WriteString(o.CompletionAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("create_id=")
	builder.WriteString(fmt.Sprintf("%v", o.CreateID))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
