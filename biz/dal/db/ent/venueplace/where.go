// Code generated by ent, DO NOT EDIT.

package venueplace

import (
	"saas/biz/dal/db/ent/internal"
	"saas/biz/dal/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldName, v))
}

// Classify applies equality check predicate on the "classify" field. It's identical to ClassifyEQ.
func Classify(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldClassify, v))
}

// Pic applies equality check predicate on the "pic" field. It's identical to PicEQ.
func Pic(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldPic, v))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldVenueID, v))
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldNumber, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldType, v))
}

// IsShow applies equality check predicate on the "is_show" field. It's identical to IsShowEQ.
func IsShow(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldIsShow, v))
}

// IsAccessible applies equality check predicate on the "is_accessible" field. It's identical to IsAccessibleEQ.
func IsAccessible(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldIsAccessible, v))
}

// IsBooking applies equality check predicate on the "is_booking" field. It's identical to IsBookingEQ.
func IsBooking(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldIsBooking, v))
}

// Information applies equality check predicate on the "information" field. It's identical to InformationEQ.
func Information(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldInformation, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldStatus))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldContainsFold(FieldName, v))
}

// ClassifyEQ applies the EQ predicate on the "classify" field.
func ClassifyEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldClassify, v))
}

// ClassifyNEQ applies the NEQ predicate on the "classify" field.
func ClassifyNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldClassify, v))
}

// ClassifyIn applies the In predicate on the "classify" field.
func ClassifyIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldClassify, vs...))
}

// ClassifyNotIn applies the NotIn predicate on the "classify" field.
func ClassifyNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldClassify, vs...))
}

// ClassifyGT applies the GT predicate on the "classify" field.
func ClassifyGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldClassify, v))
}

// ClassifyGTE applies the GTE predicate on the "classify" field.
func ClassifyGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldClassify, v))
}

// ClassifyLT applies the LT predicate on the "classify" field.
func ClassifyLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldClassify, v))
}

// ClassifyLTE applies the LTE predicate on the "classify" field.
func ClassifyLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldClassify, v))
}

// ClassifyIsNil applies the IsNil predicate on the "classify" field.
func ClassifyIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldClassify))
}

// ClassifyNotNil applies the NotNil predicate on the "classify" field.
func ClassifyNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldClassify))
}

// PicEQ applies the EQ predicate on the "pic" field.
func PicEQ(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldPic, v))
}

// PicNEQ applies the NEQ predicate on the "pic" field.
func PicNEQ(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldPic, v))
}

// PicIn applies the In predicate on the "pic" field.
func PicIn(vs ...string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldPic, vs...))
}

// PicNotIn applies the NotIn predicate on the "pic" field.
func PicNotIn(vs ...string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldPic, vs...))
}

// PicGT applies the GT predicate on the "pic" field.
func PicGT(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldPic, v))
}

// PicGTE applies the GTE predicate on the "pic" field.
func PicGTE(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldPic, v))
}

// PicLT applies the LT predicate on the "pic" field.
func PicLT(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldPic, v))
}

// PicLTE applies the LTE predicate on the "pic" field.
func PicLTE(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldPic, v))
}

// PicContains applies the Contains predicate on the "pic" field.
func PicContains(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldContains(FieldPic, v))
}

// PicHasPrefix applies the HasPrefix predicate on the "pic" field.
func PicHasPrefix(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldHasPrefix(FieldPic, v))
}

// PicHasSuffix applies the HasSuffix predicate on the "pic" field.
func PicHasSuffix(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldHasSuffix(FieldPic, v))
}

// PicIsNil applies the IsNil predicate on the "pic" field.
func PicIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldPic))
}

// PicNotNil applies the NotNil predicate on the "pic" field.
func PicNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldPic))
}

// PicEqualFold applies the EqualFold predicate on the "pic" field.
func PicEqualFold(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEqualFold(FieldPic, v))
}

// PicContainsFold applies the ContainsFold predicate on the "pic" field.
func PicContainsFold(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldContainsFold(FieldPic, v))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDIsNil applies the IsNil predicate on the "venue_id" field.
func VenueIDIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldVenueID))
}

// VenueIDNotNil applies the NotNil predicate on the "venue_id" field.
func VenueIDNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldVenueID))
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldNumber, v))
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldNumber, v))
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldNumber, vs...))
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldNumber, vs...))
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldNumber, v))
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldNumber, v))
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldNumber, v))
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldNumber, v))
}

// NumberIsNil applies the IsNil predicate on the "number" field.
func NumberIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldNumber))
}

// NumberNotNil applies the NotNil predicate on the "number" field.
func NumberNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldNumber))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldType))
}

// IsShowEQ applies the EQ predicate on the "is_show" field.
func IsShowEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldIsShow, v))
}

// IsShowNEQ applies the NEQ predicate on the "is_show" field.
func IsShowNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldIsShow, v))
}

// IsShowIn applies the In predicate on the "is_show" field.
func IsShowIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldIsShow, vs...))
}

// IsShowNotIn applies the NotIn predicate on the "is_show" field.
func IsShowNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldIsShow, vs...))
}

// IsShowGT applies the GT predicate on the "is_show" field.
func IsShowGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldIsShow, v))
}

// IsShowGTE applies the GTE predicate on the "is_show" field.
func IsShowGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldIsShow, v))
}

// IsShowLT applies the LT predicate on the "is_show" field.
func IsShowLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldIsShow, v))
}

// IsShowLTE applies the LTE predicate on the "is_show" field.
func IsShowLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldIsShow, v))
}

// IsShowIsNil applies the IsNil predicate on the "is_show" field.
func IsShowIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldIsShow))
}

// IsShowNotNil applies the NotNil predicate on the "is_show" field.
func IsShowNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldIsShow))
}

// IsAccessibleEQ applies the EQ predicate on the "is_accessible" field.
func IsAccessibleEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldIsAccessible, v))
}

// IsAccessibleNEQ applies the NEQ predicate on the "is_accessible" field.
func IsAccessibleNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldIsAccessible, v))
}

// IsAccessibleIn applies the In predicate on the "is_accessible" field.
func IsAccessibleIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldIsAccessible, vs...))
}

// IsAccessibleNotIn applies the NotIn predicate on the "is_accessible" field.
func IsAccessibleNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldIsAccessible, vs...))
}

// IsAccessibleGT applies the GT predicate on the "is_accessible" field.
func IsAccessibleGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldIsAccessible, v))
}

// IsAccessibleGTE applies the GTE predicate on the "is_accessible" field.
func IsAccessibleGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldIsAccessible, v))
}

// IsAccessibleLT applies the LT predicate on the "is_accessible" field.
func IsAccessibleLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldIsAccessible, v))
}

// IsAccessibleLTE applies the LTE predicate on the "is_accessible" field.
func IsAccessibleLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldIsAccessible, v))
}

// IsAccessibleIsNil applies the IsNil predicate on the "is_accessible" field.
func IsAccessibleIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldIsAccessible))
}

// IsAccessibleNotNil applies the NotNil predicate on the "is_accessible" field.
func IsAccessibleNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldIsAccessible))
}

// IsBookingEQ applies the EQ predicate on the "is_booking" field.
func IsBookingEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldIsBooking, v))
}

// IsBookingNEQ applies the NEQ predicate on the "is_booking" field.
func IsBookingNEQ(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldIsBooking, v))
}

// IsBookingIn applies the In predicate on the "is_booking" field.
func IsBookingIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldIsBooking, vs...))
}

// IsBookingNotIn applies the NotIn predicate on the "is_booking" field.
func IsBookingNotIn(vs ...int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldIsBooking, vs...))
}

// IsBookingGT applies the GT predicate on the "is_booking" field.
func IsBookingGT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldIsBooking, v))
}

// IsBookingGTE applies the GTE predicate on the "is_booking" field.
func IsBookingGTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldIsBooking, v))
}

// IsBookingLT applies the LT predicate on the "is_booking" field.
func IsBookingLT(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldIsBooking, v))
}

// IsBookingLTE applies the LTE predicate on the "is_booking" field.
func IsBookingLTE(v int64) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldIsBooking, v))
}

// IsBookingIsNil applies the IsNil predicate on the "is_booking" field.
func IsBookingIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldIsBooking))
}

// IsBookingNotNil applies the NotNil predicate on the "is_booking" field.
func IsBookingNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldIsBooking))
}

// InformationEQ applies the EQ predicate on the "information" field.
func InformationEQ(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEQ(FieldInformation, v))
}

// InformationNEQ applies the NEQ predicate on the "information" field.
func InformationNEQ(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNEQ(FieldInformation, v))
}

// InformationIn applies the In predicate on the "information" field.
func InformationIn(vs ...string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIn(FieldInformation, vs...))
}

// InformationNotIn applies the NotIn predicate on the "information" field.
func InformationNotIn(vs ...string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotIn(FieldInformation, vs...))
}

// InformationGT applies the GT predicate on the "information" field.
func InformationGT(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGT(FieldInformation, v))
}

// InformationGTE applies the GTE predicate on the "information" field.
func InformationGTE(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldGTE(FieldInformation, v))
}

// InformationLT applies the LT predicate on the "information" field.
func InformationLT(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLT(FieldInformation, v))
}

// InformationLTE applies the LTE predicate on the "information" field.
func InformationLTE(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldLTE(FieldInformation, v))
}

// InformationContains applies the Contains predicate on the "information" field.
func InformationContains(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldContains(FieldInformation, v))
}

// InformationHasPrefix applies the HasPrefix predicate on the "information" field.
func InformationHasPrefix(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldHasPrefix(FieldInformation, v))
}

// InformationHasSuffix applies the HasSuffix predicate on the "information" field.
func InformationHasSuffix(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldHasSuffix(FieldInformation, v))
}

// InformationIsNil applies the IsNil predicate on the "information" field.
func InformationIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldInformation))
}

// InformationNotNil applies the NotNil predicate on the "information" field.
func InformationNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldInformation))
}

// InformationEqualFold applies the EqualFold predicate on the "information" field.
func InformationEqualFold(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldEqualFold(FieldInformation, v))
}

// InformationContainsFold applies the ContainsFold predicate on the "information" field.
func InformationContainsFold(v string) predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldContainsFold(FieldInformation, v))
}

// SeatIsNil applies the IsNil predicate on the "seat" field.
func SeatIsNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldIsNull(FieldSeat))
}

// SeatNotNil applies the NotNil predicate on the "seat" field.
func SeatNotNil() predicate.VenuePlace {
	return predicate.VenuePlace(sql.FieldNotNull(FieldSeat))
}

// HasVenue applies the HasEdge predicate on the "venue" edge.
func HasVenue() predicate.VenuePlace {
	return predicate.VenuePlace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VenueTable, VenueColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.VenuePlace
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVenueWith applies the HasEdge predicate on the "venue" edge with a given conditions (other predicates).
func HasVenueWith(preds ...predicate.Venue) predicate.VenuePlace {
	return predicate.VenuePlace(func(s *sql.Selector) {
		step := newVenueStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Venue
		step.Edge.Schema = schemaConfig.VenuePlace
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProducts applies the HasEdge predicate on the "products" edge.
func HasProducts() predicate.VenuePlace {
	return predicate.VenuePlace(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ProductsTable, ProductsPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Product
		step.Edge.Schema = schemaConfig.VenuePlaceProducts
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductsWith applies the HasEdge predicate on the "products" edge with a given conditions (other predicates).
func HasProductsWith(preds ...predicate.Product) predicate.VenuePlace {
	return predicate.VenuePlace(func(s *sql.Selector) {
		step := newProductsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Product
		step.Edge.Schema = schemaConfig.VenuePlaceProducts
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VenuePlace) predicate.VenuePlace {
	return predicate.VenuePlace(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VenuePlace) predicate.VenuePlace {
	return predicate.VenuePlace(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VenuePlace) predicate.VenuePlace {
	return predicate.VenuePlace(sql.NotPredicates(p))
}
