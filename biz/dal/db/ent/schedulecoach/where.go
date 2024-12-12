// Code generated by ent, DO NOT EDIT.

package schedulecoach

import (
	"saas/biz/dal/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldStatus, v))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldVenueID, v))
}

// CoachID applies equality check predicate on the "coach_id" field. It's identical to CoachIDEQ.
func CoachID(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCoachID, v))
}

// ScheduleID applies equality check predicate on the "schedule_id" field. It's identical to ScheduleIDEQ.
func ScheduleID(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldScheduleID, v))
}

// ScheduleName applies equality check predicate on the "schedule_name" field. It's identical to ScheduleNameEQ.
func ScheduleName(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldScheduleName, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldType, v))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "end_time" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldEndTime, v))
}

// SignStartTime applies equality check predicate on the "sign_start_time" field. It's identical to SignStartTimeEQ.
func SignStartTime(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldSignStartTime, v))
}

// SignEndTime applies equality check predicate on the "sign_end_time" field. It's identical to SignEndTimeEQ.
func SignEndTime(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldSignEndTime, v))
}

// CoachName applies equality check predicate on the "coach_name" field. It's identical to CoachNameEQ.
func CoachName(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCoachName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldStatus))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDGT applies the GT predicate on the "venue_id" field.
func VenueIDGT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldVenueID, v))
}

// VenueIDGTE applies the GTE predicate on the "venue_id" field.
func VenueIDGTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldVenueID, v))
}

// VenueIDLT applies the LT predicate on the "venue_id" field.
func VenueIDLT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldVenueID, v))
}

// VenueIDLTE applies the LTE predicate on the "venue_id" field.
func VenueIDLTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldVenueID, v))
}

// VenueIDIsNil applies the IsNil predicate on the "venue_id" field.
func VenueIDIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldVenueID))
}

// VenueIDNotNil applies the NotNil predicate on the "venue_id" field.
func VenueIDNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldVenueID))
}

// CoachIDEQ applies the EQ predicate on the "coach_id" field.
func CoachIDEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCoachID, v))
}

// CoachIDNEQ applies the NEQ predicate on the "coach_id" field.
func CoachIDNEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldCoachID, v))
}

// CoachIDIn applies the In predicate on the "coach_id" field.
func CoachIDIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldCoachID, vs...))
}

// CoachIDNotIn applies the NotIn predicate on the "coach_id" field.
func CoachIDNotIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldCoachID, vs...))
}

// CoachIDGT applies the GT predicate on the "coach_id" field.
func CoachIDGT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldCoachID, v))
}

// CoachIDGTE applies the GTE predicate on the "coach_id" field.
func CoachIDGTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldCoachID, v))
}

// CoachIDLT applies the LT predicate on the "coach_id" field.
func CoachIDLT(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldCoachID, v))
}

// CoachIDLTE applies the LTE predicate on the "coach_id" field.
func CoachIDLTE(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldCoachID, v))
}

// CoachIDIsNil applies the IsNil predicate on the "coach_id" field.
func CoachIDIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldCoachID))
}

// CoachIDNotNil applies the NotNil predicate on the "coach_id" field.
func CoachIDNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldCoachID))
}

// ScheduleIDEQ applies the EQ predicate on the "schedule_id" field.
func ScheduleIDEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldScheduleID, v))
}

// ScheduleIDNEQ applies the NEQ predicate on the "schedule_id" field.
func ScheduleIDNEQ(v int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldScheduleID, v))
}

// ScheduleIDIn applies the In predicate on the "schedule_id" field.
func ScheduleIDIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldScheduleID, vs...))
}

// ScheduleIDNotIn applies the NotIn predicate on the "schedule_id" field.
func ScheduleIDNotIn(vs ...int64) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldScheduleID, vs...))
}

// ScheduleIDIsNil applies the IsNil predicate on the "schedule_id" field.
func ScheduleIDIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldScheduleID))
}

// ScheduleIDNotNil applies the NotNil predicate on the "schedule_id" field.
func ScheduleIDNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldScheduleID))
}

// ScheduleNameEQ applies the EQ predicate on the "schedule_name" field.
func ScheduleNameEQ(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldScheduleName, v))
}

// ScheduleNameNEQ applies the NEQ predicate on the "schedule_name" field.
func ScheduleNameNEQ(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldScheduleName, v))
}

// ScheduleNameIn applies the In predicate on the "schedule_name" field.
func ScheduleNameIn(vs ...string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldScheduleName, vs...))
}

// ScheduleNameNotIn applies the NotIn predicate on the "schedule_name" field.
func ScheduleNameNotIn(vs ...string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldScheduleName, vs...))
}

// ScheduleNameGT applies the GT predicate on the "schedule_name" field.
func ScheduleNameGT(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldScheduleName, v))
}

// ScheduleNameGTE applies the GTE predicate on the "schedule_name" field.
func ScheduleNameGTE(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldScheduleName, v))
}

// ScheduleNameLT applies the LT predicate on the "schedule_name" field.
func ScheduleNameLT(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldScheduleName, v))
}

// ScheduleNameLTE applies the LTE predicate on the "schedule_name" field.
func ScheduleNameLTE(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldScheduleName, v))
}

// ScheduleNameContains applies the Contains predicate on the "schedule_name" field.
func ScheduleNameContains(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldContains(FieldScheduleName, v))
}

// ScheduleNameHasPrefix applies the HasPrefix predicate on the "schedule_name" field.
func ScheduleNameHasPrefix(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldHasPrefix(FieldScheduleName, v))
}

// ScheduleNameHasSuffix applies the HasSuffix predicate on the "schedule_name" field.
func ScheduleNameHasSuffix(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldHasSuffix(FieldScheduleName, v))
}

// ScheduleNameIsNil applies the IsNil predicate on the "schedule_name" field.
func ScheduleNameIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldScheduleName))
}

// ScheduleNameNotNil applies the NotNil predicate on the "schedule_name" field.
func ScheduleNameNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldScheduleName))
}

// ScheduleNameEqualFold applies the EqualFold predicate on the "schedule_name" field.
func ScheduleNameEqualFold(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEqualFold(FieldScheduleName, v))
}

// ScheduleNameContainsFold applies the ContainsFold predicate on the "schedule_name" field.
func ScheduleNameContainsFold(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldContainsFold(FieldScheduleName, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldContainsFold(FieldType, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldStartTime, v))
}

// StartTimeIsNil applies the IsNil predicate on the "start_time" field.
func StartTimeIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldStartTime))
}

// StartTimeNotNil applies the NotNil predicate on the "start_time" field.
func StartTimeNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldStartTime))
}

// EndTimeEQ applies the EQ predicate on the "end_time" field.
func EndTimeEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "end_time" field.
func EndTimeNEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "end_time" field.
func EndTimeIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "end_time" field.
func EndTimeNotIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "end_time" field.
func EndTimeGT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "end_time" field.
func EndTimeGTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "end_time" field.
func EndTimeLT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "end_time" field.
func EndTimeLTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldEndTime, v))
}

// EndTimeIsNil applies the IsNil predicate on the "end_time" field.
func EndTimeIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldEndTime))
}

// EndTimeNotNil applies the NotNil predicate on the "end_time" field.
func EndTimeNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldEndTime))
}

// SignStartTimeEQ applies the EQ predicate on the "sign_start_time" field.
func SignStartTimeEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldSignStartTime, v))
}

// SignStartTimeNEQ applies the NEQ predicate on the "sign_start_time" field.
func SignStartTimeNEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldSignStartTime, v))
}

// SignStartTimeIn applies the In predicate on the "sign_start_time" field.
func SignStartTimeIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldSignStartTime, vs...))
}

// SignStartTimeNotIn applies the NotIn predicate on the "sign_start_time" field.
func SignStartTimeNotIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldSignStartTime, vs...))
}

// SignStartTimeGT applies the GT predicate on the "sign_start_time" field.
func SignStartTimeGT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldSignStartTime, v))
}

// SignStartTimeGTE applies the GTE predicate on the "sign_start_time" field.
func SignStartTimeGTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldSignStartTime, v))
}

// SignStartTimeLT applies the LT predicate on the "sign_start_time" field.
func SignStartTimeLT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldSignStartTime, v))
}

// SignStartTimeLTE applies the LTE predicate on the "sign_start_time" field.
func SignStartTimeLTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldSignStartTime, v))
}

// SignStartTimeIsNil applies the IsNil predicate on the "sign_start_time" field.
func SignStartTimeIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldSignStartTime))
}

// SignStartTimeNotNil applies the NotNil predicate on the "sign_start_time" field.
func SignStartTimeNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldSignStartTime))
}

// SignEndTimeEQ applies the EQ predicate on the "sign_end_time" field.
func SignEndTimeEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldSignEndTime, v))
}

// SignEndTimeNEQ applies the NEQ predicate on the "sign_end_time" field.
func SignEndTimeNEQ(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldSignEndTime, v))
}

// SignEndTimeIn applies the In predicate on the "sign_end_time" field.
func SignEndTimeIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldSignEndTime, vs...))
}

// SignEndTimeNotIn applies the NotIn predicate on the "sign_end_time" field.
func SignEndTimeNotIn(vs ...time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldSignEndTime, vs...))
}

// SignEndTimeGT applies the GT predicate on the "sign_end_time" field.
func SignEndTimeGT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldSignEndTime, v))
}

// SignEndTimeGTE applies the GTE predicate on the "sign_end_time" field.
func SignEndTimeGTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldSignEndTime, v))
}

// SignEndTimeLT applies the LT predicate on the "sign_end_time" field.
func SignEndTimeLT(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldSignEndTime, v))
}

// SignEndTimeLTE applies the LTE predicate on the "sign_end_time" field.
func SignEndTimeLTE(v time.Time) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldSignEndTime, v))
}

// SignEndTimeIsNil applies the IsNil predicate on the "sign_end_time" field.
func SignEndTimeIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldSignEndTime))
}

// SignEndTimeNotNil applies the NotNil predicate on the "sign_end_time" field.
func SignEndTimeNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldSignEndTime))
}

// CoachNameEQ applies the EQ predicate on the "coach_name" field.
func CoachNameEQ(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEQ(FieldCoachName, v))
}

// CoachNameNEQ applies the NEQ predicate on the "coach_name" field.
func CoachNameNEQ(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNEQ(FieldCoachName, v))
}

// CoachNameIn applies the In predicate on the "coach_name" field.
func CoachNameIn(vs ...string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIn(FieldCoachName, vs...))
}

// CoachNameNotIn applies the NotIn predicate on the "coach_name" field.
func CoachNameNotIn(vs ...string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotIn(FieldCoachName, vs...))
}

// CoachNameGT applies the GT predicate on the "coach_name" field.
func CoachNameGT(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGT(FieldCoachName, v))
}

// CoachNameGTE applies the GTE predicate on the "coach_name" field.
func CoachNameGTE(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldGTE(FieldCoachName, v))
}

// CoachNameLT applies the LT predicate on the "coach_name" field.
func CoachNameLT(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLT(FieldCoachName, v))
}

// CoachNameLTE applies the LTE predicate on the "coach_name" field.
func CoachNameLTE(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldLTE(FieldCoachName, v))
}

// CoachNameContains applies the Contains predicate on the "coach_name" field.
func CoachNameContains(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldContains(FieldCoachName, v))
}

// CoachNameHasPrefix applies the HasPrefix predicate on the "coach_name" field.
func CoachNameHasPrefix(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldHasPrefix(FieldCoachName, v))
}

// CoachNameHasSuffix applies the HasSuffix predicate on the "coach_name" field.
func CoachNameHasSuffix(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldHasSuffix(FieldCoachName, v))
}

// CoachNameIsNil applies the IsNil predicate on the "coach_name" field.
func CoachNameIsNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldIsNull(FieldCoachName))
}

// CoachNameNotNil applies the NotNil predicate on the "coach_name" field.
func CoachNameNotNil() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldNotNull(FieldCoachName))
}

// CoachNameEqualFold applies the EqualFold predicate on the "coach_name" field.
func CoachNameEqualFold(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldEqualFold(FieldCoachName, v))
}

// CoachNameContainsFold applies the ContainsFold predicate on the "coach_name" field.
func CoachNameContainsFold(v string) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.FieldContainsFold(FieldCoachName, v))
}

// HasSchedule applies the HasEdge predicate on the "schedule" edge.
func HasSchedule() predicate.ScheduleCoach {
	return predicate.ScheduleCoach(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ScheduleTable, ScheduleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasScheduleWith applies the HasEdge predicate on the "schedule" edge with a given conditions (other predicates).
func HasScheduleWith(preds ...predicate.Schedule) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(func(s *sql.Selector) {
		step := newScheduleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ScheduleCoach) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ScheduleCoach) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ScheduleCoach) predicate.ScheduleCoach {
	return predicate.ScheduleCoach(sql.NotPredicates(p))
}
