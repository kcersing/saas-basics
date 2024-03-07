// Code generated by ent, DO NOT EDIT.

package ordersales

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldOrderID, v))
}

// SalesID applies equality check predicate on the "sales_id" field. It's identical to SalesIDEQ.
func SalesID(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldSalesID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldOrderID))
}

// SalesIDEQ applies the EQ predicate on the "sales_id" field.
func SalesIDEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldSalesID, v))
}

// SalesIDNEQ applies the NEQ predicate on the "sales_id" field.
func SalesIDNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldSalesID, v))
}

// SalesIDIn applies the In predicate on the "sales_id" field.
func SalesIDIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldSalesID, vs...))
}

// SalesIDNotIn applies the NotIn predicate on the "sales_id" field.
func SalesIDNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldSalesID, vs...))
}

// SalesIDGT applies the GT predicate on the "sales_id" field.
func SalesIDGT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldSalesID, v))
}

// SalesIDGTE applies the GTE predicate on the "sales_id" field.
func SalesIDGTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldSalesID, v))
}

// SalesIDLT applies the LT predicate on the "sales_id" field.
func SalesIDLT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldSalesID, v))
}

// SalesIDLTE applies the LTE predicate on the "sales_id" field.
func SalesIDLTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldSalesID, v))
}

// SalesIDIsNil applies the IsNil predicate on the "sales_id" field.
func SalesIDIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldSalesID))
}

// SalesIDNotNil applies the NotNil predicate on the "sales_id" field.
func SalesIDNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldSalesID))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.OrderSales {
	return predicate.OrderSales(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.Order) predicate.OrderSales {
	return predicate.OrderSales(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderSales) predicate.OrderSales {
	return predicate.OrderSales(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderSales) predicate.OrderSales {
	return predicate.OrderSales(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderSales) predicate.OrderSales {
	return predicate.OrderSales(sql.NotPredicates(p))
}
