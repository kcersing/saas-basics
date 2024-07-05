// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldOrderID, v))
}

// ProductID applies equality check predicate on the "product_id" field. It's identical to ProductIDEQ.
func ProductID(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldProductID, v))
}

// RelatedUserProductID applies equality check predicate on the "related_user_product_id" field. It's identical to RelatedUserProductIDEQ.
func RelatedUserProductID(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldRelatedUserProductID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotNull(FieldOrderID))
}

// ProductIDEQ applies the EQ predicate on the "product_id" field.
func ProductIDEQ(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldProductID, v))
}

// ProductIDNEQ applies the NEQ predicate on the "product_id" field.
func ProductIDNEQ(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldProductID, v))
}

// ProductIDIn applies the In predicate on the "product_id" field.
func ProductIDIn(vs ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldProductID, vs...))
}

// ProductIDNotIn applies the NotIn predicate on the "product_id" field.
func ProductIDNotIn(vs ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldProductID, vs...))
}

// ProductIDGT applies the GT predicate on the "product_id" field.
func ProductIDGT(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldProductID, v))
}

// ProductIDGTE applies the GTE predicate on the "product_id" field.
func ProductIDGTE(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldProductID, v))
}

// ProductIDLT applies the LT predicate on the "product_id" field.
func ProductIDLT(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldProductID, v))
}

// ProductIDLTE applies the LTE predicate on the "product_id" field.
func ProductIDLTE(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldProductID, v))
}

// ProductIDIsNil applies the IsNil predicate on the "product_id" field.
func ProductIDIsNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIsNull(FieldProductID))
}

// ProductIDNotNil applies the NotNil predicate on the "product_id" field.
func ProductIDNotNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotNull(FieldProductID))
}

// RelatedUserProductIDEQ applies the EQ predicate on the "related_user_product_id" field.
func RelatedUserProductIDEQ(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldEQ(FieldRelatedUserProductID, v))
}

// RelatedUserProductIDNEQ applies the NEQ predicate on the "related_user_product_id" field.
func RelatedUserProductIDNEQ(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNEQ(FieldRelatedUserProductID, v))
}

// RelatedUserProductIDIn applies the In predicate on the "related_user_product_id" field.
func RelatedUserProductIDIn(vs ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIn(FieldRelatedUserProductID, vs...))
}

// RelatedUserProductIDNotIn applies the NotIn predicate on the "related_user_product_id" field.
func RelatedUserProductIDNotIn(vs ...int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotIn(FieldRelatedUserProductID, vs...))
}

// RelatedUserProductIDGT applies the GT predicate on the "related_user_product_id" field.
func RelatedUserProductIDGT(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGT(FieldRelatedUserProductID, v))
}

// RelatedUserProductIDGTE applies the GTE predicate on the "related_user_product_id" field.
func RelatedUserProductIDGTE(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldGTE(FieldRelatedUserProductID, v))
}

// RelatedUserProductIDLT applies the LT predicate on the "related_user_product_id" field.
func RelatedUserProductIDLT(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLT(FieldRelatedUserProductID, v))
}

// RelatedUserProductIDLTE applies the LTE predicate on the "related_user_product_id" field.
func RelatedUserProductIDLTE(v int64) predicate.OrderItem {
	return predicate.OrderItem(sql.FieldLTE(FieldRelatedUserProductID, v))
}

// RelatedUserProductIDIsNil applies the IsNil predicate on the "related_user_product_id" field.
func RelatedUserProductIDIsNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIsNull(FieldRelatedUserProductID))
}

// RelatedUserProductIDNotNil applies the NotNil predicate on the "related_user_product_id" field.
func RelatedUserProductIDNotNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotNull(FieldRelatedUserProductID))
}

// DataIsNil applies the IsNil predicate on the "data" field.
func DataIsNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldIsNull(FieldData))
}

// DataNotNil applies the NotNil predicate on the "data" field.
func DataNotNil() predicate.OrderItem {
	return predicate.OrderItem(sql.FieldNotNull(FieldData))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderItem {
	return predicate.OrderItem(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderItem) predicate.OrderItem {
	return predicate.OrderItem(sql.NotPredicates(p))
}
