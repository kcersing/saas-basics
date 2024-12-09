// Code generated by ent, DO NOT EDIT.

package member

import (
	"saas/biz/dal/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldStatus, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldPassword, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldName, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldNickname, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldMobile, v))
}

// Avatar applies equality check predicate on the "avatar" field. It's identical to AvatarEQ.
func Avatar(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldAvatar, v))
}

// Condition applies equality check predicate on the "condition" field. It's identical to ConditionEQ.
func Condition(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCondition, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldStatus))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldPassword))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldPassword, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldName, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldNickname))
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldNickname))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldNickname, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldMobile, v))
}

// AvatarEQ applies the EQ predicate on the "avatar" field.
func AvatarEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldAvatar, v))
}

// AvatarNEQ applies the NEQ predicate on the "avatar" field.
func AvatarNEQ(v string) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldAvatar, v))
}

// AvatarIn applies the In predicate on the "avatar" field.
func AvatarIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldAvatar, vs...))
}

// AvatarNotIn applies the NotIn predicate on the "avatar" field.
func AvatarNotIn(vs ...string) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldAvatar, vs...))
}

// AvatarGT applies the GT predicate on the "avatar" field.
func AvatarGT(v string) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldAvatar, v))
}

// AvatarGTE applies the GTE predicate on the "avatar" field.
func AvatarGTE(v string) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldAvatar, v))
}

// AvatarLT applies the LT predicate on the "avatar" field.
func AvatarLT(v string) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldAvatar, v))
}

// AvatarLTE applies the LTE predicate on the "avatar" field.
func AvatarLTE(v string) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldAvatar, v))
}

// AvatarContains applies the Contains predicate on the "avatar" field.
func AvatarContains(v string) predicate.Member {
	return predicate.Member(sql.FieldContains(FieldAvatar, v))
}

// AvatarHasPrefix applies the HasPrefix predicate on the "avatar" field.
func AvatarHasPrefix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasPrefix(FieldAvatar, v))
}

// AvatarHasSuffix applies the HasSuffix predicate on the "avatar" field.
func AvatarHasSuffix(v string) predicate.Member {
	return predicate.Member(sql.FieldHasSuffix(FieldAvatar, v))
}

// AvatarIsNil applies the IsNil predicate on the "avatar" field.
func AvatarIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldAvatar))
}

// AvatarNotNil applies the NotNil predicate on the "avatar" field.
func AvatarNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldAvatar))
}

// AvatarEqualFold applies the EqualFold predicate on the "avatar" field.
func AvatarEqualFold(v string) predicate.Member {
	return predicate.Member(sql.FieldEqualFold(FieldAvatar, v))
}

// AvatarContainsFold applies the ContainsFold predicate on the "avatar" field.
func AvatarContainsFold(v string) predicate.Member {
	return predicate.Member(sql.FieldContainsFold(FieldAvatar, v))
}

// ConditionEQ applies the EQ predicate on the "condition" field.
func ConditionEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldEQ(FieldCondition, v))
}

// ConditionNEQ applies the NEQ predicate on the "condition" field.
func ConditionNEQ(v int64) predicate.Member {
	return predicate.Member(sql.FieldNEQ(FieldCondition, v))
}

// ConditionIn applies the In predicate on the "condition" field.
func ConditionIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldIn(FieldCondition, vs...))
}

// ConditionNotIn applies the NotIn predicate on the "condition" field.
func ConditionNotIn(vs ...int64) predicate.Member {
	return predicate.Member(sql.FieldNotIn(FieldCondition, vs...))
}

// ConditionGT applies the GT predicate on the "condition" field.
func ConditionGT(v int64) predicate.Member {
	return predicate.Member(sql.FieldGT(FieldCondition, v))
}

// ConditionGTE applies the GTE predicate on the "condition" field.
func ConditionGTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldGTE(FieldCondition, v))
}

// ConditionLT applies the LT predicate on the "condition" field.
func ConditionLT(v int64) predicate.Member {
	return predicate.Member(sql.FieldLT(FieldCondition, v))
}

// ConditionLTE applies the LTE predicate on the "condition" field.
func ConditionLTE(v int64) predicate.Member {
	return predicate.Member(sql.FieldLTE(FieldCondition, v))
}

// ConditionIsNil applies the IsNil predicate on the "condition" field.
func ConditionIsNil() predicate.Member {
	return predicate.Member(sql.FieldIsNull(FieldCondition))
}

// ConditionNotNil applies the NotNil predicate on the "condition" field.
func ConditionNotNil() predicate.Member {
	return predicate.Member(sql.FieldNotNull(FieldCondition))
}

// HasMemberDetails applies the HasEdge predicate on the "member_details" edge.
func HasMemberDetails() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MemberDetailsTable, MemberDetailsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberDetailsWith applies the HasEdge predicate on the "member_details" edge with a given conditions (other predicates).
func HasMemberDetailsWith(preds ...predicate.MemberDetails) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newMemberDetailsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberNotes applies the HasEdge predicate on the "member_notes" edge.
func HasMemberNotes() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MemberNotesTable, MemberNotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberNotesWith applies the HasEdge predicate on the "member_notes" edge with a given conditions (other predicates).
func HasMemberNotesWith(preds ...predicate.MemberNote) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newMemberNotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberOrders applies the HasEdge predicate on the "member_orders" edge.
func HasMemberOrders() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MemberOrdersTable, MemberOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberOrdersWith applies the HasEdge predicate on the "member_orders" edge with a given conditions (other predicates).
func HasMemberOrdersWith(preds ...predicate.Order) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newMemberOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberEntry applies the HasEdge predicate on the "member_entry" edge.
func HasMemberEntry() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MemberEntryTable, MemberEntryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberEntryWith applies the HasEdge predicate on the "member_entry" edge with a given conditions (other predicates).
func HasMemberEntryWith(preds ...predicate.EntryLogs) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newMemberEntryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberContents applies the HasEdge predicate on the "member_contents" edge.
func HasMemberContents() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MemberContentsTable, MemberContentsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberContentsWith applies the HasEdge predicate on the "member_contents" edge with a given conditions (other predicates).
func HasMemberContentsWith(preds ...predicate.MemberContract) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newMemberContentsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParticipants applies the HasEdge predicate on the "participants" edge.
func HasParticipants() predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, ParticipantsTable, ParticipantsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParticipantsWith applies the HasEdge predicate on the "participants" edge with a given conditions (other predicates).
func HasParticipantsWith(preds ...predicate.ContestParticipant) predicate.Member {
	return predicate.Member(func(s *sql.Selector) {
		step := newParticipantsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Member) predicate.Member {
	return predicate.Member(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Member) predicate.Member {
	return predicate.Member(sql.NotPredicates(p))
}
