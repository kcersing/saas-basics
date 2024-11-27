// Code generated by ent, DO NOT EDIT.

package venue

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldStatus, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldName, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldAddress, v))
}

// AddressDetail applies equality check predicate on the "address_detail" field. It's identical to AddressDetailEQ.
func AddressDetail(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldAddressDetail, v))
}

// Latitude applies equality check predicate on the "latitude" field. It's identical to LatitudeEQ.
func Latitude(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldLatitude, v))
}

// Longitude applies equality check predicate on the "longitude" field. It's identical to LongitudeEQ.
func Longitude(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldLongitude, v))
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldMobile, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldEmail, v))
}

// Information applies equality check predicate on the "information" field. It's identical to InformationEQ.
func Information(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldInformation, v))
}

// Pic applies equality check predicate on the "pic" field. It's identical to PicEQ.
func Pic(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldPic, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldStatus))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldName, v))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldAddress))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldAddress, v))
}

// AddressDetailEQ applies the EQ predicate on the "address_detail" field.
func AddressDetailEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldAddressDetail, v))
}

// AddressDetailNEQ applies the NEQ predicate on the "address_detail" field.
func AddressDetailNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldAddressDetail, v))
}

// AddressDetailIn applies the In predicate on the "address_detail" field.
func AddressDetailIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldAddressDetail, vs...))
}

// AddressDetailNotIn applies the NotIn predicate on the "address_detail" field.
func AddressDetailNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldAddressDetail, vs...))
}

// AddressDetailGT applies the GT predicate on the "address_detail" field.
func AddressDetailGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldAddressDetail, v))
}

// AddressDetailGTE applies the GTE predicate on the "address_detail" field.
func AddressDetailGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldAddressDetail, v))
}

// AddressDetailLT applies the LT predicate on the "address_detail" field.
func AddressDetailLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldAddressDetail, v))
}

// AddressDetailLTE applies the LTE predicate on the "address_detail" field.
func AddressDetailLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldAddressDetail, v))
}

// AddressDetailContains applies the Contains predicate on the "address_detail" field.
func AddressDetailContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldAddressDetail, v))
}

// AddressDetailHasPrefix applies the HasPrefix predicate on the "address_detail" field.
func AddressDetailHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldAddressDetail, v))
}

// AddressDetailHasSuffix applies the HasSuffix predicate on the "address_detail" field.
func AddressDetailHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldAddressDetail, v))
}

// AddressDetailIsNil applies the IsNil predicate on the "address_detail" field.
func AddressDetailIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldAddressDetail))
}

// AddressDetailNotNil applies the NotNil predicate on the "address_detail" field.
func AddressDetailNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldAddressDetail))
}

// AddressDetailEqualFold applies the EqualFold predicate on the "address_detail" field.
func AddressDetailEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldAddressDetail, v))
}

// AddressDetailContainsFold applies the ContainsFold predicate on the "address_detail" field.
func AddressDetailContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldAddressDetail, v))
}

// LatitudeEQ applies the EQ predicate on the "latitude" field.
func LatitudeEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldLatitude, v))
}

// LatitudeNEQ applies the NEQ predicate on the "latitude" field.
func LatitudeNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldLatitude, v))
}

// LatitudeIn applies the In predicate on the "latitude" field.
func LatitudeIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldLatitude, vs...))
}

// LatitudeNotIn applies the NotIn predicate on the "latitude" field.
func LatitudeNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldLatitude, vs...))
}

// LatitudeGT applies the GT predicate on the "latitude" field.
func LatitudeGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldLatitude, v))
}

// LatitudeGTE applies the GTE predicate on the "latitude" field.
func LatitudeGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldLatitude, v))
}

// LatitudeLT applies the LT predicate on the "latitude" field.
func LatitudeLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldLatitude, v))
}

// LatitudeLTE applies the LTE predicate on the "latitude" field.
func LatitudeLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldLatitude, v))
}

// LatitudeContains applies the Contains predicate on the "latitude" field.
func LatitudeContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldLatitude, v))
}

// LatitudeHasPrefix applies the HasPrefix predicate on the "latitude" field.
func LatitudeHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldLatitude, v))
}

// LatitudeHasSuffix applies the HasSuffix predicate on the "latitude" field.
func LatitudeHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldLatitude, v))
}

// LatitudeIsNil applies the IsNil predicate on the "latitude" field.
func LatitudeIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldLatitude))
}

// LatitudeNotNil applies the NotNil predicate on the "latitude" field.
func LatitudeNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldLatitude))
}

// LatitudeEqualFold applies the EqualFold predicate on the "latitude" field.
func LatitudeEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldLatitude, v))
}

// LatitudeContainsFold applies the ContainsFold predicate on the "latitude" field.
func LatitudeContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldLatitude, v))
}

// LongitudeEQ applies the EQ predicate on the "longitude" field.
func LongitudeEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldLongitude, v))
}

// LongitudeNEQ applies the NEQ predicate on the "longitude" field.
func LongitudeNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldLongitude, v))
}

// LongitudeIn applies the In predicate on the "longitude" field.
func LongitudeIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldLongitude, vs...))
}

// LongitudeNotIn applies the NotIn predicate on the "longitude" field.
func LongitudeNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldLongitude, vs...))
}

// LongitudeGT applies the GT predicate on the "longitude" field.
func LongitudeGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldLongitude, v))
}

// LongitudeGTE applies the GTE predicate on the "longitude" field.
func LongitudeGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldLongitude, v))
}

// LongitudeLT applies the LT predicate on the "longitude" field.
func LongitudeLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldLongitude, v))
}

// LongitudeLTE applies the LTE predicate on the "longitude" field.
func LongitudeLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldLongitude, v))
}

// LongitudeContains applies the Contains predicate on the "longitude" field.
func LongitudeContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldLongitude, v))
}

// LongitudeHasPrefix applies the HasPrefix predicate on the "longitude" field.
func LongitudeHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldLongitude, v))
}

// LongitudeHasSuffix applies the HasSuffix predicate on the "longitude" field.
func LongitudeHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldLongitude, v))
}

// LongitudeIsNil applies the IsNil predicate on the "longitude" field.
func LongitudeIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldLongitude))
}

// LongitudeNotNil applies the NotNil predicate on the "longitude" field.
func LongitudeNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldLongitude))
}

// LongitudeEqualFold applies the EqualFold predicate on the "longitude" field.
func LongitudeEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldLongitude, v))
}

// LongitudeContainsFold applies the ContainsFold predicate on the "longitude" field.
func LongitudeContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldLongitude, v))
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldMobile, v))
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldMobile, v))
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldMobile, vs...))
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldMobile, vs...))
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldMobile, v))
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldMobile, v))
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldMobile, v))
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldMobile, v))
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldMobile, v))
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldMobile, v))
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldMobile, v))
}

// MobileIsNil applies the IsNil predicate on the "mobile" field.
func MobileIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldMobile))
}

// MobileNotNil applies the NotNil predicate on the "mobile" field.
func MobileNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldMobile))
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldMobile, v))
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldMobile, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailIsNil applies the IsNil predicate on the "email" field.
func EmailIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldEmail))
}

// EmailNotNil applies the NotNil predicate on the "email" field.
func EmailNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldEmail))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldEmail, v))
}

// InformationEQ applies the EQ predicate on the "information" field.
func InformationEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldInformation, v))
}

// InformationNEQ applies the NEQ predicate on the "information" field.
func InformationNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldInformation, v))
}

// InformationIn applies the In predicate on the "information" field.
func InformationIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldInformation, vs...))
}

// InformationNotIn applies the NotIn predicate on the "information" field.
func InformationNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldInformation, vs...))
}

// InformationGT applies the GT predicate on the "information" field.
func InformationGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldInformation, v))
}

// InformationGTE applies the GTE predicate on the "information" field.
func InformationGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldInformation, v))
}

// InformationLT applies the LT predicate on the "information" field.
func InformationLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldInformation, v))
}

// InformationLTE applies the LTE predicate on the "information" field.
func InformationLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldInformation, v))
}

// InformationContains applies the Contains predicate on the "information" field.
func InformationContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldInformation, v))
}

// InformationHasPrefix applies the HasPrefix predicate on the "information" field.
func InformationHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldInformation, v))
}

// InformationHasSuffix applies the HasSuffix predicate on the "information" field.
func InformationHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldInformation, v))
}

// InformationIsNil applies the IsNil predicate on the "information" field.
func InformationIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldInformation))
}

// InformationNotNil applies the NotNil predicate on the "information" field.
func InformationNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldInformation))
}

// InformationEqualFold applies the EqualFold predicate on the "information" field.
func InformationEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldInformation, v))
}

// InformationContainsFold applies the ContainsFold predicate on the "information" field.
func InformationContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldInformation, v))
}

// PicEQ applies the EQ predicate on the "pic" field.
func PicEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEQ(FieldPic, v))
}

// PicNEQ applies the NEQ predicate on the "pic" field.
func PicNEQ(v string) predicate.Venue {
	return predicate.Venue(sql.FieldNEQ(FieldPic, v))
}

// PicIn applies the In predicate on the "pic" field.
func PicIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldIn(FieldPic, vs...))
}

// PicNotIn applies the NotIn predicate on the "pic" field.
func PicNotIn(vs ...string) predicate.Venue {
	return predicate.Venue(sql.FieldNotIn(FieldPic, vs...))
}

// PicGT applies the GT predicate on the "pic" field.
func PicGT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGT(FieldPic, v))
}

// PicGTE applies the GTE predicate on the "pic" field.
func PicGTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldGTE(FieldPic, v))
}

// PicLT applies the LT predicate on the "pic" field.
func PicLT(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLT(FieldPic, v))
}

// PicLTE applies the LTE predicate on the "pic" field.
func PicLTE(v string) predicate.Venue {
	return predicate.Venue(sql.FieldLTE(FieldPic, v))
}

// PicContains applies the Contains predicate on the "pic" field.
func PicContains(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContains(FieldPic, v))
}

// PicHasPrefix applies the HasPrefix predicate on the "pic" field.
func PicHasPrefix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasPrefix(FieldPic, v))
}

// PicHasSuffix applies the HasSuffix predicate on the "pic" field.
func PicHasSuffix(v string) predicate.Venue {
	return predicate.Venue(sql.FieldHasSuffix(FieldPic, v))
}

// PicIsNil applies the IsNil predicate on the "pic" field.
func PicIsNil() predicate.Venue {
	return predicate.Venue(sql.FieldIsNull(FieldPic))
}

// PicNotNil applies the NotNil predicate on the "pic" field.
func PicNotNil() predicate.Venue {
	return predicate.Venue(sql.FieldNotNull(FieldPic))
}

// PicEqualFold applies the EqualFold predicate on the "pic" field.
func PicEqualFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldEqualFold(FieldPic, v))
}

// PicContainsFold applies the ContainsFold predicate on the "pic" field.
func PicContainsFold(v string) predicate.Venue {
	return predicate.Venue(sql.FieldContainsFold(FieldPic, v))
}

// HasPlaces applies the HasEdge predicate on the "places" edge.
func HasPlaces() predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PlacesTable, PlacesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPlacesWith applies the HasEdge predicate on the "places" edge with a given conditions (other predicates).
func HasPlacesWith(preds ...predicate.VenuePlace) predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		step := newPlacesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasVenueOrders applies the HasEdge predicate on the "venue_orders" edge.
func HasVenueOrders() predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VenueOrdersTable, VenueOrdersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVenueOrdersWith applies the HasEdge predicate on the "venue_orders" edge with a given conditions (other predicates).
func HasVenueOrdersWith(preds ...predicate.Order) predicate.Venue {
	return predicate.Venue(func(s *sql.Selector) {
		step := newVenueOrdersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Venue) predicate.Venue {
	return predicate.Venue(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Venue) predicate.Venue {
	return predicate.Venue(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Venue) predicate.Venue {
	return predicate.Venue(sql.NotPredicates(p))
}
