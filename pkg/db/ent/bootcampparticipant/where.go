// Code generated by ent, DO NOT EDIT.

package bootcampparticipant

import (
	"saas/pkg/db/ent/predicate"
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

// DeleteAt applies equality check predicate on the "delete_at" field. It's identical to DeleteAtEQ.
func DeleteAt(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldDeleteAt, v))
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

// DeleteAtEQ applies the EQ predicate on the "delete_at" field.
func DeleteAtEQ(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldEQ(FieldDeleteAt, v))
}

// DeleteAtNEQ applies the NEQ predicate on the "delete_at" field.
func DeleteAtNEQ(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNEQ(FieldDeleteAt, v))
}

// DeleteAtIn applies the In predicate on the "delete_at" field.
func DeleteAtIn(vs ...time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldIn(FieldDeleteAt, vs...))
}

// DeleteAtNotIn applies the NotIn predicate on the "delete_at" field.
func DeleteAtNotIn(vs ...time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldNotIn(FieldDeleteAt, vs...))
}

// DeleteAtGT applies the GT predicate on the "delete_at" field.
func DeleteAtGT(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGT(FieldDeleteAt, v))
}

// DeleteAtGTE applies the GTE predicate on the "delete_at" field.
func DeleteAtGTE(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldGTE(FieldDeleteAt, v))
}

// DeleteAtLT applies the LT predicate on the "delete_at" field.
func DeleteAtLT(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLT(FieldDeleteAt, v))
}

// DeleteAtLTE applies the LTE predicate on the "delete_at" field.
func DeleteAtLTE(v time.Time) predicate.BootcampParticipant {
	return predicate.BootcampParticipant(sql.FieldLTE(FieldDeleteAt, v))
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

// HasBootcamp applies the HasEdge predicate on the "bootcamp" edge.
func HasBootcamp() predicate.BootcampParticipant {
	return predicate.BootcampParticipant(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, BootcampTable, BootcampColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBootcampWith applies the HasEdge predicate on the "bootcamp" edge with a given conditions (other predicates).
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
