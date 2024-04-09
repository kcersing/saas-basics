// Code generated by ent, DO NOT EDIT.

package courserecordschedule

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldStatus, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldType, v))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldVenueID, v))
}

// PlaceID applies equality check predicate on the "place_id" field. It's identical to PlaceIDEQ.
func PlaceID(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldPlaceID, v))
}

// Num applies equality check predicate on the "num" field. It's identical to NumEQ.
func Num(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldNum, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldEndTime, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldPrice, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldStatus))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldContainsFold(FieldType, v))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDGT applies the GT predicate on the "venue_id" field.
func VenueIDGT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldVenueID, v))
}

// VenueIDGTE applies the GTE predicate on the "venue_id" field.
func VenueIDGTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldVenueID, v))
}

// VenueIDLT applies the LT predicate on the "venue_id" field.
func VenueIDLT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldVenueID, v))
}

// VenueIDLTE applies the LTE predicate on the "venue_id" field.
func VenueIDLTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldVenueID, v))
}

// VenueIDIsNil applies the IsNil predicate on the "venue_id" field.
func VenueIDIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldVenueID))
}

// VenueIDNotNil applies the NotNil predicate on the "venue_id" field.
func VenueIDNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldVenueID))
}

// PlaceIDEQ applies the EQ predicate on the "place_id" field.
func PlaceIDEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldPlaceID, v))
}

// PlaceIDNEQ applies the NEQ predicate on the "place_id" field.
func PlaceIDNEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldPlaceID, v))
}

// PlaceIDIn applies the In predicate on the "place_id" field.
func PlaceIDIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldPlaceID, vs...))
}

// PlaceIDNotIn applies the NotIn predicate on the "place_id" field.
func PlaceIDNotIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldPlaceID, vs...))
}

// PlaceIDGT applies the GT predicate on the "place_id" field.
func PlaceIDGT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldPlaceID, v))
}

// PlaceIDGTE applies the GTE predicate on the "place_id" field.
func PlaceIDGTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldPlaceID, v))
}

// PlaceIDLT applies the LT predicate on the "place_id" field.
func PlaceIDLT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldPlaceID, v))
}

// PlaceIDLTE applies the LTE predicate on the "place_id" field.
func PlaceIDLTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldPlaceID, v))
}

// PlaceIDIsNil applies the IsNil predicate on the "place_id" field.
func PlaceIDIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldPlaceID))
}

// PlaceIDNotNil applies the NotNil predicate on the "place_id" field.
func PlaceIDNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldPlaceID))
}

// NumEQ applies the EQ predicate on the "num" field.
func NumEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldNum, v))
}

// NumNEQ applies the NEQ predicate on the "num" field.
func NumNEQ(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldNum, v))
}

// NumIn applies the In predicate on the "num" field.
func NumIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldNum, vs...))
}

// NumNotIn applies the NotIn predicate on the "num" field.
func NumNotIn(vs ...int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldNum, vs...))
}

// NumGT applies the GT predicate on the "num" field.
func NumGT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldNum, v))
}

// NumGTE applies the GTE predicate on the "num" field.
func NumGTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldNum, v))
}

// NumLT applies the LT predicate on the "num" field.
func NumLT(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldNum, v))
}

// NumLTE applies the LTE predicate on the "num" field.
func NumLTE(v int64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldNum, v))
}

// NumIsNil applies the IsNil predicate on the "num" field.
func NumIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldNum))
}

// NumNotNil applies the NotNil predicate on the "num" field.
func NumNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldNum))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeIsNil applies the IsNil predicate on the "start_time" field.
func StartTimeIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldStartTime))
}

// StartTimeNotNil applies the NotNil predicate on the "start_time" field.
func StartTimeNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldStartTime))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldEndTime, v))
}

// EndTimeIsNil applies the IsNil predicate on the "end_time" field.
func EndTimeIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldEndTime))
}

// EndTimeNotNil applies the NotNil predicate on the "end_time" field.
func EndTimeNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldEndTime))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldLTE(FieldPrice, v))
}

// PriceIsNil applies the IsNil predicate on the "price" field.
func PriceIsNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldIsNull(FieldPrice))
}

// PriceNotNil applies the NotNil predicate on the "price" field.
func PriceNotNil() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.FieldNotNull(FieldPrice))
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MembersTable, MembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.CourseRecordMember) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCoachs applies the HasEdge predicate on the "coachs" edge.
func HasCoachs() predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CoachsTable, CoachsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoachsWith applies the HasEdge predicate on the "coachs" edge with a given conditions (other predicates).
func HasCoachsWith(preds ...predicate.CourseRecordCoach) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(func(s *sql.Selector) {
		step := newCoachsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CourseRecordSchedule) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CourseRecordSchedule) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CourseRecordSchedule) predicate.CourseRecordSchedule {
	return predicate.CourseRecordSchedule(sql.NotPredicates(p))
}
