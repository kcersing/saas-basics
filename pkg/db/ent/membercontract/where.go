// Code generated by ent, DO NOT EDIT.

package membercontract

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldStatus, v))
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldMemberID, v))
}

// ContractID applies equality check predicate on the "contract_id" field. It's identical to ContractIDEQ.
func ContractID(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldContractID, v))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldOrderID, v))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldVenueID, v))
}

// MemberProductID applies equality check predicate on the "member_product_id" field. It's identical to MemberProductIDEQ.
func MemberProductID(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldMemberProductID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldName, v))
}

// Sign applies equality check predicate on the "sign" field. It's identical to SignEQ.
func Sign(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldSign, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldStatus))
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldMemberID, v))
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldMemberID, v))
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldMemberID, vs...))
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldMemberID, vs...))
}

// MemberIDIsNil applies the IsNil predicate on the "member_id" field.
func MemberIDIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldMemberID))
}

// MemberIDNotNil applies the NotNil predicate on the "member_id" field.
func MemberIDNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldMemberID))
}

// ContractIDEQ applies the EQ predicate on the "contract_id" field.
func ContractIDEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldContractID, v))
}

// ContractIDNEQ applies the NEQ predicate on the "contract_id" field.
func ContractIDNEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldContractID, v))
}

// ContractIDIn applies the In predicate on the "contract_id" field.
func ContractIDIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldContractID, vs...))
}

// ContractIDNotIn applies the NotIn predicate on the "contract_id" field.
func ContractIDNotIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldContractID, vs...))
}

// ContractIDGT applies the GT predicate on the "contract_id" field.
func ContractIDGT(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldContractID, v))
}

// ContractIDGTE applies the GTE predicate on the "contract_id" field.
func ContractIDGTE(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldContractID, v))
}

// ContractIDLT applies the LT predicate on the "contract_id" field.
func ContractIDLT(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldContractID, v))
}

// ContractIDLTE applies the LTE predicate on the "contract_id" field.
func ContractIDLTE(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldContractID, v))
}

// ContractIDIsNil applies the IsNil predicate on the "contract_id" field.
func ContractIDIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldContractID))
}

// ContractIDNotNil applies the NotNil predicate on the "contract_id" field.
func ContractIDNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldContractID))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDIsNil applies the IsNil predicate on the "order_id" field.
func OrderIDIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldOrderID))
}

// OrderIDNotNil applies the NotNil predicate on the "order_id" field.
func OrderIDNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldOrderID))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDGT applies the GT predicate on the "venue_id" field.
func VenueIDGT(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldVenueID, v))
}

// VenueIDGTE applies the GTE predicate on the "venue_id" field.
func VenueIDGTE(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldVenueID, v))
}

// VenueIDLT applies the LT predicate on the "venue_id" field.
func VenueIDLT(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldVenueID, v))
}

// VenueIDLTE applies the LTE predicate on the "venue_id" field.
func VenueIDLTE(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldVenueID, v))
}

// VenueIDIsNil applies the IsNil predicate on the "venue_id" field.
func VenueIDIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldVenueID))
}

// VenueIDNotNil applies the NotNil predicate on the "venue_id" field.
func VenueIDNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldVenueID))
}

// MemberProductIDEQ applies the EQ predicate on the "member_product_id" field.
func MemberProductIDEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldMemberProductID, v))
}

// MemberProductIDNEQ applies the NEQ predicate on the "member_product_id" field.
func MemberProductIDNEQ(v int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldMemberProductID, v))
}

// MemberProductIDIn applies the In predicate on the "member_product_id" field.
func MemberProductIDIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldMemberProductID, vs...))
}

// MemberProductIDNotIn applies the NotIn predicate on the "member_product_id" field.
func MemberProductIDNotIn(vs ...int64) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldMemberProductID, vs...))
}

// MemberProductIDIsNil applies the IsNil predicate on the "member_product_id" field.
func MemberProductIDIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldMemberProductID))
}

// MemberProductIDNotNil applies the NotNil predicate on the "member_product_id" field.
func MemberProductIDNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldMemberProductID))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldContainsFold(FieldName, v))
}

// SignEQ applies the EQ predicate on the "sign" field.
func SignEQ(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEQ(FieldSign, v))
}

// SignNEQ applies the NEQ predicate on the "sign" field.
func SignNEQ(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNEQ(FieldSign, v))
}

// SignIn applies the In predicate on the "sign" field.
func SignIn(vs ...string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIn(FieldSign, vs...))
}

// SignNotIn applies the NotIn predicate on the "sign" field.
func SignNotIn(vs ...string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotIn(FieldSign, vs...))
}

// SignGT applies the GT predicate on the "sign" field.
func SignGT(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGT(FieldSign, v))
}

// SignGTE applies the GTE predicate on the "sign" field.
func SignGTE(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldGTE(FieldSign, v))
}

// SignLT applies the LT predicate on the "sign" field.
func SignLT(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLT(FieldSign, v))
}

// SignLTE applies the LTE predicate on the "sign" field.
func SignLTE(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldLTE(FieldSign, v))
}

// SignContains applies the Contains predicate on the "sign" field.
func SignContains(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldContains(FieldSign, v))
}

// SignHasPrefix applies the HasPrefix predicate on the "sign" field.
func SignHasPrefix(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldHasPrefix(FieldSign, v))
}

// SignHasSuffix applies the HasSuffix predicate on the "sign" field.
func SignHasSuffix(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldHasSuffix(FieldSign, v))
}

// SignIsNil applies the IsNil predicate on the "sign" field.
func SignIsNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldIsNull(FieldSign))
}

// SignNotNil applies the NotNil predicate on the "sign" field.
func SignNotNil() predicate.MemberContract {
	return predicate.MemberContract(sql.FieldNotNull(FieldSign))
}

// SignEqualFold applies the EqualFold predicate on the "sign" field.
func SignEqualFold(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldEqualFold(FieldSign, v))
}

// SignContainsFold applies the ContainsFold predicate on the "sign" field.
func SignContainsFold(v string) predicate.MemberContract {
	return predicate.MemberContract(sql.FieldContainsFold(FieldSign, v))
}

// HasContent applies the HasEdge predicate on the "content" edge.
func HasContent() predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ContentTable, ContentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContentWith applies the HasEdge predicate on the "content" edge with a given conditions (other predicates).
func HasContentWith(preds ...predicate.MemberContractContent) predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := newContentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.Member) predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := newMemberStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOrder applies the HasEdge predicate on the "order" edge.
func HasOrder() predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOrderWith applies the HasEdge predicate on the "order" edge with a given conditions (other predicates).
func HasOrderWith(preds ...predicate.Order) predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := newOrderStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberProduct applies the HasEdge predicate on the "member_product" edge.
func HasMemberProduct() predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberProductTable, MemberProductColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberProductWith applies the HasEdge predicate on the "member_product" edge with a given conditions (other predicates).
func HasMemberProductWith(preds ...predicate.MemberProduct) predicate.MemberContract {
	return predicate.MemberContract(func(s *sql.Selector) {
		step := newMemberProductStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemberContract) predicate.MemberContract {
	return predicate.MemberContract(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemberContract) predicate.MemberContract {
	return predicate.MemberContract(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemberContract) predicate.MemberContract {
	return predicate.MemberContract(sql.NotPredicates(p))
}
