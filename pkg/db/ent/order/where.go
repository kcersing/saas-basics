// Code generated by ent, DO NOT EDIT.

package order

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderSn applies equality check predicate on the "order_sn" field. It's identical to OrderSnEQ.
func OrderSn(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderSn, v))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldVenueID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUserID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldSource, v))
}

// Device applies equality check predicate on the "device" field. It's identical to DeviceEQ.
func Device(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDevice, v))
}

// CompletionAt applies equality check predicate on the "completion_at" field. It's identical to CompletionAtEQ.
func CompletionAt(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCompletionAt, v))
}

// CreateID applies equality check predicate on the "create_id" field. It's identical to CreateIDEQ.
func CreateID(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreateID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldUpdatedAt, v))
}

// OrderSnEQ applies the EQ predicate on the "order_sn" field.
func OrderSnEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldOrderSn, v))
}

// OrderSnNEQ applies the NEQ predicate on the "order_sn" field.
func OrderSnNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldOrderSn, v))
}

// OrderSnIn applies the In predicate on the "order_sn" field.
func OrderSnIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldOrderSn, vs...))
}

// OrderSnNotIn applies the NotIn predicate on the "order_sn" field.
func OrderSnNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldOrderSn, vs...))
}

// OrderSnGT applies the GT predicate on the "order_sn" field.
func OrderSnGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldOrderSn, v))
}

// OrderSnGTE applies the GTE predicate on the "order_sn" field.
func OrderSnGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldOrderSn, v))
}

// OrderSnLT applies the LT predicate on the "order_sn" field.
func OrderSnLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldOrderSn, v))
}

// OrderSnLTE applies the LTE predicate on the "order_sn" field.
func OrderSnLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldOrderSn, v))
}

// OrderSnContains applies the Contains predicate on the "order_sn" field.
func OrderSnContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldOrderSn, v))
}

// OrderSnHasPrefix applies the HasPrefix predicate on the "order_sn" field.
func OrderSnHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldOrderSn, v))
}

// OrderSnHasSuffix applies the HasSuffix predicate on the "order_sn" field.
func OrderSnHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldOrderSn, v))
}

// OrderSnIsNil applies the IsNil predicate on the "order_sn" field.
func OrderSnIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldOrderSn))
}

// OrderSnNotNil applies the NotNil predicate on the "order_sn" field.
func OrderSnNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldOrderSn))
}

// OrderSnEqualFold applies the EqualFold predicate on the "order_sn" field.
func OrderSnEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldOrderSn, v))
}

// OrderSnContainsFold applies the ContainsFold predicate on the "order_sn" field.
func OrderSnContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldOrderSn, v))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDGT applies the GT predicate on the "venue_id" field.
func VenueIDGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldVenueID, v))
}

// VenueIDGTE applies the GTE predicate on the "venue_id" field.
func VenueIDGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldVenueID, v))
}

// VenueIDLT applies the LT predicate on the "venue_id" field.
func VenueIDLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldVenueID, v))
}

// VenueIDLTE applies the LTE predicate on the "venue_id" field.
func VenueIDLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldVenueID, v))
}

// VenueIDIsNil applies the IsNil predicate on the "venue_id" field.
func VenueIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldVenueID))
}

// VenueIDNotNil applies the NotNil predicate on the "venue_id" field.
func VenueIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldVenueID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldUserID, v))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldUserID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldStatus))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldSource, v))
}

// SourceIsNil applies the IsNil predicate on the "source" field.
func SourceIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldSource))
}

// SourceNotNil applies the NotNil predicate on the "source" field.
func SourceNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldSource))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldSource, v))
}

// DeviceEQ applies the EQ predicate on the "device" field.
func DeviceEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldDevice, v))
}

// DeviceNEQ applies the NEQ predicate on the "device" field.
func DeviceNEQ(v string) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldDevice, v))
}

// DeviceIn applies the In predicate on the "device" field.
func DeviceIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldDevice, vs...))
}

// DeviceNotIn applies the NotIn predicate on the "device" field.
func DeviceNotIn(vs ...string) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldDevice, vs...))
}

// DeviceGT applies the GT predicate on the "device" field.
func DeviceGT(v string) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldDevice, v))
}

// DeviceGTE applies the GTE predicate on the "device" field.
func DeviceGTE(v string) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldDevice, v))
}

// DeviceLT applies the LT predicate on the "device" field.
func DeviceLT(v string) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldDevice, v))
}

// DeviceLTE applies the LTE predicate on the "device" field.
func DeviceLTE(v string) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldDevice, v))
}

// DeviceContains applies the Contains predicate on the "device" field.
func DeviceContains(v string) predicate.Order {
	return predicate.Order(sql.FieldContains(FieldDevice, v))
}

// DeviceHasPrefix applies the HasPrefix predicate on the "device" field.
func DeviceHasPrefix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasPrefix(FieldDevice, v))
}

// DeviceHasSuffix applies the HasSuffix predicate on the "device" field.
func DeviceHasSuffix(v string) predicate.Order {
	return predicate.Order(sql.FieldHasSuffix(FieldDevice, v))
}

// DeviceIsNil applies the IsNil predicate on the "device" field.
func DeviceIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldDevice))
}

// DeviceNotNil applies the NotNil predicate on the "device" field.
func DeviceNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldDevice))
}

// DeviceEqualFold applies the EqualFold predicate on the "device" field.
func DeviceEqualFold(v string) predicate.Order {
	return predicate.Order(sql.FieldEqualFold(FieldDevice, v))
}

// DeviceContainsFold applies the ContainsFold predicate on the "device" field.
func DeviceContainsFold(v string) predicate.Order {
	return predicate.Order(sql.FieldContainsFold(FieldDevice, v))
}

// CompletionAtEQ applies the EQ predicate on the "completion_at" field.
func CompletionAtEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCompletionAt, v))
}

// CompletionAtNEQ applies the NEQ predicate on the "completion_at" field.
func CompletionAtNEQ(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCompletionAt, v))
}

// CompletionAtIn applies the In predicate on the "completion_at" field.
func CompletionAtIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCompletionAt, vs...))
}

// CompletionAtNotIn applies the NotIn predicate on the "completion_at" field.
func CompletionAtNotIn(vs ...time.Time) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCompletionAt, vs...))
}

// CompletionAtGT applies the GT predicate on the "completion_at" field.
func CompletionAtGT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCompletionAt, v))
}

// CompletionAtGTE applies the GTE predicate on the "completion_at" field.
func CompletionAtGTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCompletionAt, v))
}

// CompletionAtLT applies the LT predicate on the "completion_at" field.
func CompletionAtLT(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCompletionAt, v))
}

// CompletionAtLTE applies the LTE predicate on the "completion_at" field.
func CompletionAtLTE(v time.Time) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCompletionAt, v))
}

// CompletionAtIsNil applies the IsNil predicate on the "completion_at" field.
func CompletionAtIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldCompletionAt))
}

// CompletionAtNotNil applies the NotNil predicate on the "completion_at" field.
func CompletionAtNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldCompletionAt))
}

// CreateIDEQ applies the EQ predicate on the "create_id" field.
func CreateIDEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldEQ(FieldCreateID, v))
}

// CreateIDNEQ applies the NEQ predicate on the "create_id" field.
func CreateIDNEQ(v int64) predicate.Order {
	return predicate.Order(sql.FieldNEQ(FieldCreateID, v))
}

// CreateIDIn applies the In predicate on the "create_id" field.
func CreateIDIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldIn(FieldCreateID, vs...))
}

// CreateIDNotIn applies the NotIn predicate on the "create_id" field.
func CreateIDNotIn(vs ...int64) predicate.Order {
	return predicate.Order(sql.FieldNotIn(FieldCreateID, vs...))
}

// CreateIDGT applies the GT predicate on the "create_id" field.
func CreateIDGT(v int64) predicate.Order {
	return predicate.Order(sql.FieldGT(FieldCreateID, v))
}

// CreateIDGTE applies the GTE predicate on the "create_id" field.
func CreateIDGTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldGTE(FieldCreateID, v))
}

// CreateIDLT applies the LT predicate on the "create_id" field.
func CreateIDLT(v int64) predicate.Order {
	return predicate.Order(sql.FieldLT(FieldCreateID, v))
}

// CreateIDLTE applies the LTE predicate on the "create_id" field.
func CreateIDLTE(v int64) predicate.Order {
	return predicate.Order(sql.FieldLTE(FieldCreateID, v))
}

// CreateIDIsNil applies the IsNil predicate on the "create_id" field.
func CreateIDIsNil() predicate.Order {
	return predicate.Order(sql.FieldIsNull(FieldCreateID))
}

// CreateIDNotNil applies the NotNil predicate on the "create_id" field.
func CreateIDNotNil() predicate.Order {
	return predicate.Order(sql.FieldNotNull(FieldCreateID))
}

// HasAmount applies the HasEdge predicate on the "amount" edge.
func HasAmount() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AmountTable, AmountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAmountWith applies the HasEdge predicate on the "amount" edge with a given conditions (other predicates).
func HasAmountWith(preds ...predicate.OrderAmount) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newAmountStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasItem applies the HasEdge predicate on the "item" edge.
func HasItem() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemWith applies the HasEdge predicate on the "item" edge with a given conditions (other predicates).
func HasItemWith(preds ...predicate.OrderItem) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPay applies the HasEdge predicate on the "pay" edge.
func HasPay() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PayTable, PayColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPayWith applies the HasEdge predicate on the "pay" edge with a given conditions (other predicates).
func HasPayWith(preds ...predicate.OrderPay) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newPayStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSales applies the HasEdge predicate on the "sales" edge.
func HasSales() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SalesTable, SalesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSalesWith applies the HasEdge predicate on the "sales" edge with a given conditions (other predicates).
func HasSalesWith(preds ...predicate.OrderSales) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := newSalesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(sql.NotPredicates(p))
}
