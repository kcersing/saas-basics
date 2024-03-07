// Code generated by ent, DO NOT EDIT.

package productproperty

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldUpdatedAt, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldType, v))
}

// SpuName applies equality check predicate on the "spu_name" field. It's identical to SpuNameEQ.
func SpuName(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldSpuName, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldDuration, v))
}

// Length applies equality check predicate on the "length" field. It's identical to LengthEQ.
func Length(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldLength, v))
}

// Count applies equality check predicate on the "count" field. It's identical to CountEQ.
func Count(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldCount, v))
}

// SpuPrice applies equality check predicate on the "spu_price" field. It's identical to SpuPriceEQ.
func SpuPrice(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldSpuPrice, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldHasSuffix(FieldType, v))
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIsNull(FieldType))
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotNull(FieldType))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldContainsFold(FieldType, v))
}

// SpuNameEQ applies the EQ predicate on the "spu_name" field.
func SpuNameEQ(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldSpuName, v))
}

// SpuNameNEQ applies the NEQ predicate on the "spu_name" field.
func SpuNameNEQ(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldSpuName, v))
}

// SpuNameIn applies the In predicate on the "spu_name" field.
func SpuNameIn(vs ...string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldSpuName, vs...))
}

// SpuNameNotIn applies the NotIn predicate on the "spu_name" field.
func SpuNameNotIn(vs ...string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldSpuName, vs...))
}

// SpuNameGT applies the GT predicate on the "spu_name" field.
func SpuNameGT(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldSpuName, v))
}

// SpuNameGTE applies the GTE predicate on the "spu_name" field.
func SpuNameGTE(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldSpuName, v))
}

// SpuNameLT applies the LT predicate on the "spu_name" field.
func SpuNameLT(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldSpuName, v))
}

// SpuNameLTE applies the LTE predicate on the "spu_name" field.
func SpuNameLTE(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldSpuName, v))
}

// SpuNameContains applies the Contains predicate on the "spu_name" field.
func SpuNameContains(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldContains(FieldSpuName, v))
}

// SpuNameHasPrefix applies the HasPrefix predicate on the "spu_name" field.
func SpuNameHasPrefix(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldHasPrefix(FieldSpuName, v))
}

// SpuNameHasSuffix applies the HasSuffix predicate on the "spu_name" field.
func SpuNameHasSuffix(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldHasSuffix(FieldSpuName, v))
}

// SpuNameIsNil applies the IsNil predicate on the "spu_name" field.
func SpuNameIsNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIsNull(FieldSpuName))
}

// SpuNameNotNil applies the NotNil predicate on the "spu_name" field.
func SpuNameNotNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotNull(FieldSpuName))
}

// SpuNameEqualFold applies the EqualFold predicate on the "spu_name" field.
func SpuNameEqualFold(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEqualFold(FieldSpuName, v))
}

// SpuNameContainsFold applies the ContainsFold predicate on the "spu_name" field.
func SpuNameContainsFold(v string) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldContainsFold(FieldSpuName, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldDuration, v))
}

// DurationIsNil applies the IsNil predicate on the "duration" field.
func DurationIsNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIsNull(FieldDuration))
}

// DurationNotNil applies the NotNil predicate on the "duration" field.
func DurationNotNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotNull(FieldDuration))
}

// LengthEQ applies the EQ predicate on the "length" field.
func LengthEQ(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldLength, v))
}

// LengthNEQ applies the NEQ predicate on the "length" field.
func LengthNEQ(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldLength, v))
}

// LengthIn applies the In predicate on the "length" field.
func LengthIn(vs ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldLength, vs...))
}

// LengthNotIn applies the NotIn predicate on the "length" field.
func LengthNotIn(vs ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldLength, vs...))
}

// LengthGT applies the GT predicate on the "length" field.
func LengthGT(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldLength, v))
}

// LengthGTE applies the GTE predicate on the "length" field.
func LengthGTE(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldLength, v))
}

// LengthLT applies the LT predicate on the "length" field.
func LengthLT(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldLength, v))
}

// LengthLTE applies the LTE predicate on the "length" field.
func LengthLTE(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldLength, v))
}

// LengthIsNil applies the IsNil predicate on the "length" field.
func LengthIsNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIsNull(FieldLength))
}

// LengthNotNil applies the NotNil predicate on the "length" field.
func LengthNotNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotNull(FieldLength))
}

// CountEQ applies the EQ predicate on the "count" field.
func CountEQ(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldCount, v))
}

// CountNEQ applies the NEQ predicate on the "count" field.
func CountNEQ(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldCount, v))
}

// CountIn applies the In predicate on the "count" field.
func CountIn(vs ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldCount, vs...))
}

// CountNotIn applies the NotIn predicate on the "count" field.
func CountNotIn(vs ...int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldCount, vs...))
}

// CountGT applies the GT predicate on the "count" field.
func CountGT(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldCount, v))
}

// CountGTE applies the GTE predicate on the "count" field.
func CountGTE(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldCount, v))
}

// CountLT applies the LT predicate on the "count" field.
func CountLT(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldCount, v))
}

// CountLTE applies the LTE predicate on the "count" field.
func CountLTE(v int64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldCount, v))
}

// CountIsNil applies the IsNil predicate on the "count" field.
func CountIsNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIsNull(FieldCount))
}

// CountNotNil applies the NotNil predicate on the "count" field.
func CountNotNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotNull(FieldCount))
}

// SpuPriceEQ applies the EQ predicate on the "spu_price" field.
func SpuPriceEQ(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldEQ(FieldSpuPrice, v))
}

// SpuPriceNEQ applies the NEQ predicate on the "spu_price" field.
func SpuPriceNEQ(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNEQ(FieldSpuPrice, v))
}

// SpuPriceIn applies the In predicate on the "spu_price" field.
func SpuPriceIn(vs ...float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIn(FieldSpuPrice, vs...))
}

// SpuPriceNotIn applies the NotIn predicate on the "spu_price" field.
func SpuPriceNotIn(vs ...float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotIn(FieldSpuPrice, vs...))
}

// SpuPriceGT applies the GT predicate on the "spu_price" field.
func SpuPriceGT(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGT(FieldSpuPrice, v))
}

// SpuPriceGTE applies the GTE predicate on the "spu_price" field.
func SpuPriceGTE(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldGTE(FieldSpuPrice, v))
}

// SpuPriceLT applies the LT predicate on the "spu_price" field.
func SpuPriceLT(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLT(FieldSpuPrice, v))
}

// SpuPriceLTE applies the LTE predicate on the "spu_price" field.
func SpuPriceLTE(v float64) predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldLTE(FieldSpuPrice, v))
}

// SpuPriceIsNil applies the IsNil predicate on the "spu_price" field.
func SpuPriceIsNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldIsNull(FieldSpuPrice))
}

// SpuPriceNotNil applies the NotNil predicate on the "spu_price" field.
func SpuPriceNotNil() predicate.ProductProperty {
	return predicate.ProductProperty(sql.FieldNotNull(FieldSpuPrice))
}

// HasProduct applies the HasEdge predicate on the "product" edge.
func HasProduct() predicate.ProductProperty {
	return predicate.ProductProperty(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ProductTable, ProductPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProductWith applies the HasEdge predicate on the "product" edge with a given conditions (other predicates).
func HasProductWith(preds ...predicate.Product) predicate.ProductProperty {
	return predicate.ProductProperty(func(s *sql.Selector) {
		step := newProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductProperty) predicate.ProductProperty {
	return predicate.ProductProperty(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductProperty) predicate.ProductProperty {
	return predicate.ProductProperty(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductProperty) predicate.ProductProperty {
	return predicate.ProductProperty(sql.NotPredicates(p))
}
