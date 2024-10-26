// Code generated by ent, DO NOT EDIT.

package orderamount

import (
	"order/biz/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldOrderID, v))
}

// Total applies equality check predicate on the "total" field. It's identical to TotalEQ.
func Total(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldTotal, v))
}

// Actual applies equality check predicate on the "actual" field. It's identical to ActualEQ.
func Actual(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldActual, v))
}

// Residue applies equality check predicate on the "residue" field. It's identical to ResidueEQ.
func Residue(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldResidue, v))
}

// Remission applies equality check predicate on the "remission" field. It's identical to RemissionEQ.
func Remission(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldRemission, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotNull(FieldOrderID))
}

// TotalEQ applies the EQ predicate on the "total" field.
func TotalEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldTotal, v))
}

// TotalNEQ applies the NEQ predicate on the "total" field.
func TotalNEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldTotal, v))
}

// TotalIn applies the In predicate on the "total" field.
func TotalIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldTotal, vs...))
}

// TotalNotIn applies the NotIn predicate on the "total" field.
func TotalNotIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldTotal, vs...))
}

// TotalGT applies the GT predicate on the "total" field.
func TotalGT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldTotal, v))
}

// TotalGTE applies the GTE predicate on the "total" field.
func TotalGTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldTotal, v))
}

// TotalLT applies the LT predicate on the "total" field.
func TotalLT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldTotal, v))
}

// TotalLTE applies the LTE predicate on the "total" field.
func TotalLTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldTotal, v))
}

// TotalIsNil applies the IsNil predicate on the "total" field.
func TotalIsNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIsNull(FieldTotal))
}

// TotalNotNil applies the NotNil predicate on the "total" field.
func TotalNotNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotNull(FieldTotal))
}

// ActualEQ applies the EQ predicate on the "actual" field.
func ActualEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldActual, v))
}

// ActualNEQ applies the NEQ predicate on the "actual" field.
func ActualNEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldActual, v))
}

// ActualIn applies the In predicate on the "actual" field.
func ActualIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldActual, vs...))
}

// ActualNotIn applies the NotIn predicate on the "actual" field.
func ActualNotIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldActual, vs...))
}

// ActualGT applies the GT predicate on the "actual" field.
func ActualGT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldActual, v))
}

// ActualGTE applies the GTE predicate on the "actual" field.
func ActualGTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldActual, v))
}

// ActualLT applies the LT predicate on the "actual" field.
func ActualLT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldActual, v))
}

// ActualLTE applies the LTE predicate on the "actual" field.
func ActualLTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldActual, v))
}

// ActualIsNil applies the IsNil predicate on the "actual" field.
func ActualIsNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIsNull(FieldActual))
}

// ActualNotNil applies the NotNil predicate on the "actual" field.
func ActualNotNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotNull(FieldActual))
}

// ResidueEQ applies the EQ predicate on the "residue" field.
func ResidueEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldResidue, v))
}

// ResidueNEQ applies the NEQ predicate on the "residue" field.
func ResidueNEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldResidue, v))
}

// ResidueIn applies the In predicate on the "residue" field.
func ResidueIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldResidue, vs...))
}

// ResidueNotIn applies the NotIn predicate on the "residue" field.
func ResidueNotIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldResidue, vs...))
}

// ResidueGT applies the GT predicate on the "residue" field.
func ResidueGT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldResidue, v))
}

// ResidueGTE applies the GTE predicate on the "residue" field.
func ResidueGTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldResidue, v))
}

// ResidueLT applies the LT predicate on the "residue" field.
func ResidueLT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldResidue, v))
}

// ResidueLTE applies the LTE predicate on the "residue" field.
func ResidueLTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldResidue, v))
}

// ResidueIsNil applies the IsNil predicate on the "residue" field.
func ResidueIsNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIsNull(FieldResidue))
}

// ResidueNotNil applies the NotNil predicate on the "residue" field.
func ResidueNotNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotNull(FieldResidue))
}

// RemissionEQ applies the EQ predicate on the "remission" field.
func RemissionEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldEQ(FieldRemission, v))
}

// RemissionNEQ applies the NEQ predicate on the "remission" field.
func RemissionNEQ(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNEQ(FieldRemission, v))
}

// RemissionIn applies the In predicate on the "remission" field.
func RemissionIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIn(FieldRemission, vs...))
}

// RemissionNotIn applies the NotIn predicate on the "remission" field.
func RemissionNotIn(vs ...float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotIn(FieldRemission, vs...))
}

// RemissionGT applies the GT predicate on the "remission" field.
func RemissionGT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGT(FieldRemission, v))
}

// RemissionGTE applies the GTE predicate on the "remission" field.
func RemissionGTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldGTE(FieldRemission, v))
}

// RemissionLT applies the LT predicate on the "remission" field.
func RemissionLT(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLT(FieldRemission, v))
}

// RemissionLTE applies the LTE predicate on the "remission" field.
func RemissionLTE(v float64) predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldLTE(FieldRemission, v))
}

// RemissionIsNil applies the IsNil predicate on the "remission" field.
func RemissionIsNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldIsNull(FieldRemission))
}

// RemissionNotNil applies the NotNil predicate on the "remission" field.
func RemissionNotNil() predicate.OrderAmount {
	return predicate.OrderAmount(sql.FieldNotNull(FieldRemission))
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.OrderAmount {
	return predicate.OrderAmount(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.OrderAmount {
	return predicate.OrderAmount(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.OrderAmount) predicate.OrderAmount {
	return predicate.OrderAmount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.OrderAmount) predicate.OrderAmount {
	return predicate.OrderAmount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.OrderAmount) predicate.OrderAmount {
	return predicate.OrderAmount(sql.NotPredicates(p))
}
