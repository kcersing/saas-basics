// Code generated by ent, DO NOT EDIT.

package banner

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldName, v))
}

// Pic applies equality check predicate on the "pic" field. It's identical to PicEQ.
func Pic(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldPic, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldLink, v))
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldIsShow, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldStatus))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldName, v))
}

// PicEQ applies the EQ predicate on the "pic" field.
func PicEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldPic, v))
}

// PicNEQ applies the NEQ predicate on the "pic" field.
func PicNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldPic, v))
}

// PicIn applies the In predicate on the "pic" field.
func PicIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldPic, vs...))
}

// PicNotIn applies the NotIn predicate on the "pic" field.
func PicNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldPic, vs...))
}

// PicGT applies the GT predicate on the "pic" field.
func PicGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldPic, v))
}

// PicGTE applies the GTE predicate on the "pic" field.
func PicGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldPic, v))
}

// PicLT applies the LT predicate on the "pic" field.
func PicLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldPic, v))
}

// PicLTE applies the LTE predicate on the "pic" field.
func PicLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldPic, v))
}

// PicContains applies the Contains predicate on the "pic" field.
func PicContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldPic, v))
}

// PicHasPrefix applies the HasPrefix predicate on the "pic" field.
func PicHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldPic, v))
}

// PicHasSuffix applies the HasSuffix predicate on the "pic" field.
func PicHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldPic, v))
}

// PicEqualFold applies the EqualFold predicate on the "pic" field.
func PicEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldPic, v))
}

// PicContainsFold applies the ContainsFold predicate on the "pic" field.
func PicContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldPic, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.Banner {
	return predicate.Banner(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.Banner {
	return predicate.Banner(sql.FieldContainsFold(FieldLink, v))
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldEQ(FieldIsShow, v))
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldNEQ(FieldIsShow, v))
}

// IsShowIn applies the In predicate on the "is_show" field.
func IsShowIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldIn(FieldIsShow, vs...))
}

// IsShowNotIn applies the NotIn predicate on the "is_show" field.
func IsShowNotIn(vs ...int64) predicate.Banner {
	return predicate.Banner(sql.FieldNotIn(FieldIsShow, vs...))
}

// IsShowGT applies the GT predicate on the "is_show" field.
func IsShowGT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGT(FieldIsShow, v))
}

// IsShowGTE applies the GTE predicate on the "is_show" field.
func IsShowGTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldGTE(FieldIsShow, v))
}

// IsShowLT applies the LT predicate on the "is_show" field.
func IsShowLT(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLT(FieldIsShow, v))
}

// IsShowLTE applies the LTE predicate on the "is_show" field.
func IsShowLTE(v int64) predicate.Banner {
	return predicate.Banner(sql.FieldLTE(FieldIsShow, v))
}

// IsShowIsNil applies the IsNil predicate on the "is_show" field.
func IsShowIsNil() predicate.Banner {
	return predicate.Banner(sql.FieldIsNull(FieldIsShow))
}

// IsShowNotNil applies the NotNil predicate on the "is_show" field.
func IsShowNotNil() predicate.Banner {
	return predicate.Banner(sql.FieldNotNull(FieldIsShow))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Banner) predicate.Banner {
	return predicate.Banner(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Banner) predicate.Banner {
	return predicate.Banner(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Banner) predicate.Banner {
	return predicate.Banner(sql.NotPredicates(p))
}
