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

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldStatus, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldOrderID, v))
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldMemberID, v))
}

// SalesID applies equality check predicate on the "sales_id" field. It's identical to SalesIDEQ.
func SalesID(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldSalesID, v))
}

// Ratio applies equality check predicate on the "ratio" field. It's identical to RatioEQ.
func Ratio(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldRatio, v))
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

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldCreatedAt))
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

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldStatus))
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

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldMemberID, v))
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldMemberID, v))
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldMemberID, vs...))
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldMemberID, vs...))
}

// MemberIDGT applies the GT predicate on the "member_id" field.
func MemberIDGT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldMemberID, v))
}

// MemberIDGTE applies the GTE predicate on the "member_id" field.
func MemberIDGTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldMemberID, v))
}

// MemberIDLT applies the LT predicate on the "member_id" field.
func MemberIDLT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldMemberID, v))
}

// MemberIDLTE applies the LTE predicate on the "member_id" field.
func MemberIDLTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldMemberID, v))
}

// MemberIDIsNil applies the IsNil predicate on the "member_id" field.
func MemberIDIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldMemberID))
}

// MemberIDNotNil applies the NotNil predicate on the "member_id" field.
func MemberIDNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldMemberID))
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

// RatioEQ applies the EQ predicate on the "ratio" field.
func RatioEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldEQ(FieldRatio, v))
}

// RatioNEQ applies the NEQ predicate on the "ratio" field.
func RatioNEQ(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNEQ(FieldRatio, v))
}

// RatioIn applies the In predicate on the "ratio" field.
func RatioIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIn(FieldRatio, vs...))
}

// RatioNotIn applies the NotIn predicate on the "ratio" field.
func RatioNotIn(vs ...int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotIn(FieldRatio, vs...))
}

// RatioGT applies the GT predicate on the "ratio" field.
func RatioGT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGT(FieldRatio, v))
}

// RatioGTE applies the GTE predicate on the "ratio" field.
func RatioGTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldGTE(FieldRatio, v))
}

// RatioLT applies the LT predicate on the "ratio" field.
func RatioLT(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLT(FieldRatio, v))
}

// RatioLTE applies the LTE predicate on the "ratio" field.
func RatioLTE(v int64) predicate.OrderSales {
	return predicate.OrderSales(sql.FieldLTE(FieldRatio, v))
}

// RatioIsNil applies the IsNil predicate on the "ratio" field.
func RatioIsNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldIsNull(FieldRatio))
}

// RatioNotNil applies the NotNil predicate on the "ratio" field.
func RatioNotNil() predicate.OrderSales {
	return predicate.OrderSales(sql.FieldNotNull(FieldRatio))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderSales {
	return predicate.OrderSales(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderSales {
	return predicate.OrderSales(func(s *sql.Selector) {
		step := newOrderStep()
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
