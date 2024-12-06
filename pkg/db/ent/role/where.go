// Code generated by ent, DO NOT EDIT.

package role

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldUpdatedAt, v))
}

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDeleteAt, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldName, v))
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldValue, v))
}

// DefaultRouter applies equality check predicate on the "default_router" field. It's identical to DefaultRouterEQ.
func DefaultRouter(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDefaultRouter, v))
}

// Remark applies equality check predicate on the "remark" field. It's identical to RemarkEQ.
func Remark(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRemark, v))
}

// OrderNo applies equality check predicate on the "order_no" field. It's identical to OrderNoEQ.
func OrderNo(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldOrderNo, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldUpdatedAt, v))
}

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldDeleteAt, v))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldCreatedID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Role {
	return predicate.Role(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Role {
	return predicate.Role(sql.FieldNotNull(FieldStatus))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldName, v))
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldValue, v))
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldValue, v))
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldValue, vs...))
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldValue, vs...))
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldValue, v))
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldValue, v))
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldValue, v))
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldValue, v))
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldValue, v))
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldValue, v))
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldValue, v))
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldValue, v))
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldValue, v))
}

// DefaultRouterEQ applies the EQ predicate on the "default_router" field.
func DefaultRouterEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldDefaultRouter, v))
}

// DefaultRouterNEQ applies the NEQ predicate on the "default_router" field.
func DefaultRouterNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldDefaultRouter, v))
}

// DefaultRouterIn applies the In predicate on the "default_router" field.
func DefaultRouterIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldDefaultRouter, vs...))
}

// DefaultRouterNotIn applies the NotIn predicate on the "default_router" field.
func DefaultRouterNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldDefaultRouter, vs...))
}

// DefaultRouterGT applies the GT predicate on the "default_router" field.
func DefaultRouterGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldDefaultRouter, v))
}

// DefaultRouterGTE applies the GTE predicate on the "default_router" field.
func DefaultRouterGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldDefaultRouter, v))
}

// DefaultRouterLT applies the LT predicate on the "default_router" field.
func DefaultRouterLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldDefaultRouter, v))
}

// DefaultRouterLTE applies the LTE predicate on the "default_router" field.
func DefaultRouterLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldDefaultRouter, v))
}

// DefaultRouterContains applies the Contains predicate on the "default_router" field.
func DefaultRouterContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldDefaultRouter, v))
}

// DefaultRouterHasPrefix applies the HasPrefix predicate on the "default_router" field.
func DefaultRouterHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldDefaultRouter, v))
}

// DefaultRouterHasSuffix applies the HasSuffix predicate on the "default_router" field.
func DefaultRouterHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldDefaultRouter, v))
}

// DefaultRouterEqualFold applies the EqualFold predicate on the "default_router" field.
func DefaultRouterEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldDefaultRouter, v))
}

// DefaultRouterContainsFold applies the ContainsFold predicate on the "default_router" field.
func DefaultRouterContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldDefaultRouter, v))
}

// RemarkEQ applies the EQ predicate on the "remark" field.
func RemarkEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldRemark, v))
}

// RemarkNEQ applies the NEQ predicate on the "remark" field.
func RemarkNEQ(v string) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldRemark, v))
}

// RemarkIn applies the In predicate on the "remark" field.
func RemarkIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldRemark, vs...))
}

// RemarkNotIn applies the NotIn predicate on the "remark" field.
func RemarkNotIn(vs ...string) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldRemark, vs...))
}

// RemarkGT applies the GT predicate on the "remark" field.
func RemarkGT(v string) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldRemark, v))
}

// RemarkGTE applies the GTE predicate on the "remark" field.
func RemarkGTE(v string) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldRemark, v))
}

// RemarkLT applies the LT predicate on the "remark" field.
func RemarkLT(v string) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldRemark, v))
}

// RemarkLTE applies the LTE predicate on the "remark" field.
func RemarkLTE(v string) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldRemark, v))
}

// RemarkContains applies the Contains predicate on the "remark" field.
func RemarkContains(v string) predicate.Role {
	return predicate.Role(sql.FieldContains(FieldRemark, v))
}

// RemarkHasPrefix applies the HasPrefix predicate on the "remark" field.
func RemarkHasPrefix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasPrefix(FieldRemark, v))
}

// RemarkHasSuffix applies the HasSuffix predicate on the "remark" field.
func RemarkHasSuffix(v string) predicate.Role {
	return predicate.Role(sql.FieldHasSuffix(FieldRemark, v))
}

// RemarkEqualFold applies the EqualFold predicate on the "remark" field.
func RemarkEqualFold(v string) predicate.Role {
	return predicate.Role(sql.FieldEqualFold(FieldRemark, v))
}

// RemarkContainsFold applies the ContainsFold predicate on the "remark" field.
func RemarkContainsFold(v string) predicate.Role {
	return predicate.Role(sql.FieldContainsFold(FieldRemark, v))
}

// OrderNoEQ applies the EQ predicate on the "order_no" field.
func OrderNoEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldEQ(FieldOrderNo, v))
}

// OrderNoNEQ applies the NEQ predicate on the "order_no" field.
func OrderNoNEQ(v int64) predicate.Role {
	return predicate.Role(sql.FieldNEQ(FieldOrderNo, v))
}

// OrderNoIn applies the In predicate on the "order_no" field.
func OrderNoIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldIn(FieldOrderNo, vs...))
}

// OrderNoNotIn applies the NotIn predicate on the "order_no" field.
func OrderNoNotIn(vs ...int64) predicate.Role {
	return predicate.Role(sql.FieldNotIn(FieldOrderNo, vs...))
}

// OrderNoGT applies the GT predicate on the "order_no" field.
func OrderNoGT(v int64) predicate.Role {
	return predicate.Role(sql.FieldGT(FieldOrderNo, v))
}

// OrderNoGTE applies the GTE predicate on the "order_no" field.
func OrderNoGTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldGTE(FieldOrderNo, v))
}

// OrderNoLT applies the LT predicate on the "order_no" field.
func OrderNoLT(v int64) predicate.Role {
	return predicate.Role(sql.FieldLT(FieldOrderNo, v))
}

// OrderNoLTE applies the LTE predicate on the "order_no" field.
func OrderNoLTE(v int64) predicate.Role {
	return predicate.Role(sql.FieldLTE(FieldOrderNo, v))
}

// HasMenus applies the HasEdge predicate on the "menus" edge.
func HasMenus() predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, MenusTable, MenusPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMenusWith applies the HasEdge predicate on the "menus" edge with a given conditions (other predicates).
func HasMenusWith(preds ...predicate.Menu) predicate.Role {
	return predicate.Role(func(s *sql.Selector) {
		step := newMenusStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Role) predicate.Role {
	return predicate.Role(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Role) predicate.Role {
	return predicate.Role(sql.NotPredicates(p))
}
