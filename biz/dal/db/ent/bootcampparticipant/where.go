// Code generated by ent, DO NOT EDIT.

package bootcampparticipant

import (
	"saas/biz/dal/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldStatus, v))
}

// BootcampID applies equality check predicate on the "bootcamp_id" field. It's identical to BootcampIDEQ.
func BootcampID(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldBootcampID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldName, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldMobile, v))
}

// Fields applies equality check predicate on the "fields" field. It's identical to FieldsEQ.
func Fields(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldFields, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldOrderID, v))
}

// OrderSn applies equality check predicate on the "order_sn" field. It's identical to OrderSnEQ.
func OrderSn(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldOrderSn, v))
}

// Fee applies equality check predicate on the "fee" field. It's identical to FeeEQ.
func Fee(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldFee, v))
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldMemberID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldStatus))
}

// BootcampIDEQ applies the EQ predicate on the "bootcamp_id" field.
func BootcampIDEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldBootcampID, v))
}

// BootcampIDNEQ applies the NEQ predicate on the "bootcamp_id" field.
func BootcampIDNEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldBootcampID, v))
}

// BootcampIDIn applies the In predicate on the "bootcamp_id" field.
func BootcampIDIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldBootcampID, vs...))
}

// BootcampIDNotIn applies the NotIn predicate on the "bootcamp_id" field.
func BootcampIDNotIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldBootcampID, vs...))
}

// BootcampIDIsNil applies the IsNil predicate on the "bootcamp_id" field.
func BootcampIDIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldBootcampID))
}

// BootcampIDNotNil applies the NotNil predicate on the "bootcamp_id" field.
func BootcampIDNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldBootcampID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContainsFold(FieldName, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContainsFold(FieldMobile, v))
}

// FieldsEQ applies the EQ predicate on the "fields" field.
func FieldsEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldFields, v))
}

// FieldsNEQ applies the NEQ predicate on the "fields" field.
func FieldsNEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldFields, v))
}

// FieldsIn applies the In predicate on the "fields" field.
func FieldsIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldFields, vs...))
}

// FieldsNotIn applies the NotIn predicate on the "fields" field.
func FieldsNotIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldFields, vs...))
}

// FieldsGT applies the GT predicate on the "fields" field.
func FieldsGT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldFields, v))
}

// FieldsGTE applies the GTE predicate on the "fields" field.
func FieldsGTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldFields, v))
}

// FieldsLT applies the LT predicate on the "fields" field.
func FieldsLT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldFields, v))
}

// FieldsLTE applies the LTE predicate on the "fields" field.
func FieldsLTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldFields, v))
}

// FieldsContains applies the Contains predicate on the "fields" field.
func FieldsContains(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContains(FieldFields, v))
}

// FieldsHasPrefix applies the HasPrefix predicate on the "fields" field.
func FieldsHasPrefix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasPrefix(FieldFields, v))
}

// FieldsHasSuffix applies the HasSuffix predicate on the "fields" field.
func FieldsHasSuffix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasSuffix(FieldFields, v))
}

// FieldsIsNil applies the IsNil predicate on the "fields" field.
func FieldsIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldFields))
}

// FieldsNotNil applies the NotNil predicate on the "fields" field.
func FieldsNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldFields))
}

// FieldsEqualFold applies the EqualFold predicate on the "fields" field.
func FieldsEqualFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEqualFold(FieldFields, v))
}

// FieldsContainsFold applies the ContainsFold predicate on the "fields" field.
func FieldsContainsFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContainsFold(FieldFields, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldOrderID, v))
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldOrderID, v))
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldOrderID, v))
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldOrderID, v))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldOrderID))
}

// OrderSnEQ applies the EQ predicate on the "order_sn" field.
func OrderSnEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldOrderSn, v))
}

// OrderSnNEQ applies the NEQ predicate on the "order_sn" field.
func OrderSnNEQ(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldOrderSn, v))
}

// OrderSnIn applies the In predicate on the "order_sn" field.
func OrderSnIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldOrderSn, vs...))
}

// OrderSnNotIn applies the NotIn predicate on the "order_sn" field.
func OrderSnNotIn(vs ...string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldOrderSn, vs...))
}

// OrderSnGT applies the GT predicate on the "order_sn" field.
func OrderSnGT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldOrderSn, v))
}

// OrderSnGTE applies the GTE predicate on the "order_sn" field.
func OrderSnGTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldOrderSn, v))
}

// OrderSnLT applies the LT predicate on the "order_sn" field.
func OrderSnLT(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldOrderSn, v))
}

// OrderSnLTE applies the LTE predicate on the "order_sn" field.
func OrderSnLTE(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldOrderSn, v))
}

// OrderSnContains applies the Contains predicate on the "order_sn" field.
func OrderSnContains(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContains(FieldOrderSn, v))
}

// OrderSnHasPrefix applies the HasPrefix predicate on the "order_sn" field.
func OrderSnHasPrefix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasPrefix(FieldOrderSn, v))
}

// OrderSnHasSuffix applies the HasSuffix predicate on the "order_sn" field.
func OrderSnHasSuffix(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldHasSuffix(FieldOrderSn, v))
}

// OrderSnIsNil applies the IsNil predicate on the "order_sn" field.
func OrderSnIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldOrderSn))
}

// OrderSnNotNil applies the NotNil predicate on the "order_sn" field.
func OrderSnNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldOrderSn))
}

// OrderSnEqualFold applies the EqualFold predicate on the "order_sn" field.
func OrderSnEqualFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEqualFold(FieldOrderSn, v))
}

// OrderSnContainsFold applies the ContainsFold predicate on the "order_sn" field.
func OrderSnContainsFold(v string) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldContainsFold(FieldOrderSn, v))
}

// FeeEQ applies the EQ predicate on the "fee" field.
func FeeEQ(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldFee, v))
}

// FeeNEQ applies the NEQ predicate on the "fee" field.
func FeeNEQ(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldFee, v))
}

// FeeIn applies the In predicate on the "fee" field.
func FeeIn(vs ...float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldFee, vs...))
}

// FeeNotIn applies the NotIn predicate on the "fee" field.
func FeeNotIn(vs ...float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldFee, vs...))
}

// FeeGT applies the GT predicate on the "fee" field.
func FeeGT(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldFee, v))
}

// FeeGTE applies the GTE predicate on the "fee" field.
func FeeGTE(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldFee, v))
}

// FeeLT applies the LT predicate on the "fee" field.
func FeeLT(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldFee, v))
}

// FeeLTE applies the LTE predicate on the "fee" field.
func FeeLTE(v float64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldFee, v))
}

// FeeIsNil applies the IsNil predicate on the "fee" field.
func FeeIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldFee))
}

// FeeNotNil applies the NotNil predicate on the "fee" field.
func FeeNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldFee))
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldMemberID, v))
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldMemberID, v))
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldMemberID, vs...))
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldMemberID, vs...))
}

// MemberIDGT applies the GT predicate on the "member_id" field.
func MemberIDGT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldMemberID, v))
}

// MemberIDGTE applies the GTE predicate on the "member_id" field.
func MemberIDGTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldMemberID, v))
}

// MemberIDLT applies the LT predicate on the "member_id" field.
func MemberIDLT(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldMemberID, v))
}

// MemberIDLTE applies the LTE predicate on the "member_id" field.
func MemberIDLTE(v int64) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldMemberID, v))
}

// MemberIDIsNil applies the IsNil predicate on the "member_id" field.
func MemberIDIsNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIsNull(FieldMemberID))
}

// MemberIDNotNil applies the NotNil predicate on the "member_id" field.
func MemberIDNotNil() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotNull(FieldMemberID))
}

// HasBootcamp applies the HasEdge predicate on the "Bootcamp" edge.
func HasBootcamp() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BootcampTable, BootcampColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBootcampWith applies the HasEdge predicate on the "Bootcamp" edge with a given conditions (other predicates).
func HasBootcampWith(preds ...predicate.Bootcamp) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(func(s *sql.Selector) {
		step := newBootcampStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, MembersTable, MembersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.Member) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.BootcampParticipant) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.BootcampParticipant) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.BootcampParticipant) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.NotPredicates(p))
}
