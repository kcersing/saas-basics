// Code generated by ent, DO NOT EDIT.

package messages

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldUpdatedAt, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldType, v))
}

// ToUserID applies equality check predicate on the "to_user_id" field. It's identical to ToUserIDEQ.
func ToUserID(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldToUserID, v))
}

// FromUserID applies equality check predicate on the "from_user_id" field. It's identical to FromUserIDEQ.
func FromUserID(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldFromUserID, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContainsFold(FieldType, v))
}

// ToUserIDEQ applies the EQ predicate on the "to_user_id" field.
func ToUserIDEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldToUserID, v))
}

// ToUserIDNEQ applies the NEQ predicate on the "to_user_id" field.
func ToUserIDNEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldToUserID, v))
}

// ToUserIDIn applies the In predicate on the "to_user_id" field.
func ToUserIDIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldToUserID, vs...))
}

// ToUserIDNotIn applies the NotIn predicate on the "to_user_id" field.
func ToUserIDNotIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldToUserID, vs...))
}

// ToUserIDGT applies the GT predicate on the "to_user_id" field.
func ToUserIDGT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldToUserID, v))
}

// ToUserIDGTE applies the GTE predicate on the "to_user_id" field.
func ToUserIDGTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldToUserID, v))
}

// ToUserIDLT applies the LT predicate on the "to_user_id" field.
func ToUserIDLT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldToUserID, v))
}

// ToUserIDLTE applies the LTE predicate on the "to_user_id" field.
func ToUserIDLTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldToUserID, v))
}

// ToUserIDContains applies the Contains predicate on the "to_user_id" field.
func ToUserIDContains(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContains(FieldToUserID, v))
}

// ToUserIDHasPrefix applies the HasPrefix predicate on the "to_user_id" field.
func ToUserIDHasPrefix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasPrefix(FieldToUserID, v))
}

// ToUserIDHasSuffix applies the HasSuffix predicate on the "to_user_id" field.
func ToUserIDHasSuffix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasSuffix(FieldToUserID, v))
}

// ToUserIDEqualFold applies the EqualFold predicate on the "to_user_id" field.
func ToUserIDEqualFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEqualFold(FieldToUserID, v))
}

// ToUserIDContainsFold applies the ContainsFold predicate on the "to_user_id" field.
func ToUserIDContainsFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContainsFold(FieldToUserID, v))
}

// FromUserIDEQ applies the EQ predicate on the "from_user_id" field.
func FromUserIDEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldFromUserID, v))
}

// FromUserIDNEQ applies the NEQ predicate on the "from_user_id" field.
func FromUserIDNEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldFromUserID, v))
}

// FromUserIDIn applies the In predicate on the "from_user_id" field.
func FromUserIDIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldFromUserID, vs...))
}

// FromUserIDNotIn applies the NotIn predicate on the "from_user_id" field.
func FromUserIDNotIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldFromUserID, vs...))
}

// FromUserIDGT applies the GT predicate on the "from_user_id" field.
func FromUserIDGT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldFromUserID, v))
}

// FromUserIDGTE applies the GTE predicate on the "from_user_id" field.
func FromUserIDGTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldFromUserID, v))
}

// FromUserIDLT applies the LT predicate on the "from_user_id" field.
func FromUserIDLT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldFromUserID, v))
}

// FromUserIDLTE applies the LTE predicate on the "from_user_id" field.
func FromUserIDLTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldFromUserID, v))
}

// FromUserIDContains applies the Contains predicate on the "from_user_id" field.
func FromUserIDContains(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContains(FieldFromUserID, v))
}

// FromUserIDHasPrefix applies the HasPrefix predicate on the "from_user_id" field.
func FromUserIDHasPrefix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasPrefix(FieldFromUserID, v))
}

// FromUserIDHasSuffix applies the HasSuffix predicate on the "from_user_id" field.
func FromUserIDHasSuffix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasSuffix(FieldFromUserID, v))
}

// FromUserIDEqualFold applies the EqualFold predicate on the "from_user_id" field.
func FromUserIDEqualFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEqualFold(FieldFromUserID, v))
}

// FromUserIDContainsFold applies the ContainsFold predicate on the "from_user_id" field.
func FromUserIDContainsFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContainsFold(FieldFromUserID, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Messages {
	return predicate.Messages(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Messages {
	return predicate.Messages(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Messages {
	return predicate.Messages(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Messages {
	return predicate.Messages(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Messages {
	return predicate.Messages(sql.FieldContainsFold(FieldContent, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Messages) predicate.Messages {
	return predicate.Messages(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Messages) predicate.Messages {
	return predicate.Messages(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Messages) predicate.Messages {
	return predicate.Messages(sql.NotPredicates(p))
}
