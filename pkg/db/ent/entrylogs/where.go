// Code generated by ent, DO NOT EDIT.

package entrylogs

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldUpdatedAt, v))
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldMemberID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldUserID, v))
}

// VenueID applies equality check predicate on the "venue_id" field. It's identical to VenueIDEQ.
func VenueID(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldVenueID, v))
}

// MemberProductID applies equality check predicate on the "member_product_id" field. It's identical to MemberProductIDEQ.
func MemberProductID(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldMemberProductID, v))
}

// MemberPropertyID applies equality check predicate on the "member_property_id" field. It's identical to MemberPropertyIDEQ.
func MemberPropertyID(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldMemberPropertyID, v))
}

// EntryTime applies equality check predicate on the "entry_time" field. It's identical to EntryTimeEQ.
func EntryTime(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldEntryTime, v))
}

// LeavingTime applies equality check predicate on the "leaving_time" field. It's identical to LeavingTimeEQ.
func LeavingTime(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldLeavingTime, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLTE(FieldUpdatedAt, v))
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldMemberID, v))
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldMemberID, v))
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldMemberID, vs...))
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldMemberID, vs...))
}

// MemberIDIsNil applies the IsNil predicate on the "member_id" field.
func MemberIDIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldMemberID))
}

// MemberIDNotNil applies the NotNil predicate on the "member_id" field.
func MemberIDNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldMemberID))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldUserID))
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldUserID))
}

// VenueIDEQ applies the EQ predicate on the "venue_id" field.
func VenueIDEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldVenueID, v))
}

// VenueIDNEQ applies the NEQ predicate on the "venue_id" field.
func VenueIDNEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldVenueID, v))
}

// VenueIDIn applies the In predicate on the "venue_id" field.
func VenueIDIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldVenueID, vs...))
}

// VenueIDNotIn applies the NotIn predicate on the "venue_id" field.
func VenueIDNotIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldVenueID, vs...))
}

// VenueIDIsNil applies the IsNil predicate on the "venue_id" field.
func VenueIDIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldVenueID))
}

// VenueIDNotNil applies the NotNil predicate on the "venue_id" field.
func VenueIDNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldVenueID))
}

// MemberProductIDEQ applies the EQ predicate on the "member_product_id" field.
func MemberProductIDEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldMemberProductID, v))
}

// MemberProductIDNEQ applies the NEQ predicate on the "member_product_id" field.
func MemberProductIDNEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldMemberProductID, v))
}

// MemberProductIDIn applies the In predicate on the "member_product_id" field.
func MemberProductIDIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldMemberProductID, vs...))
}

// MemberProductIDNotIn applies the NotIn predicate on the "member_product_id" field.
func MemberProductIDNotIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldMemberProductID, vs...))
}

// MemberProductIDIsNil applies the IsNil predicate on the "member_product_id" field.
func MemberProductIDIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldMemberProductID))
}

// MemberProductIDNotNil applies the NotNil predicate on the "member_product_id" field.
func MemberProductIDNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldMemberProductID))
}

// MemberPropertyIDEQ applies the EQ predicate on the "member_property_id" field.
func MemberPropertyIDEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldMemberPropertyID, v))
}

// MemberPropertyIDNEQ applies the NEQ predicate on the "member_property_id" field.
func MemberPropertyIDNEQ(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldMemberPropertyID, v))
}

// MemberPropertyIDIn applies the In predicate on the "member_property_id" field.
func MemberPropertyIDIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldMemberPropertyID, vs...))
}

// MemberPropertyIDNotIn applies the NotIn predicate on the "member_property_id" field.
func MemberPropertyIDNotIn(vs ...int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldMemberPropertyID, vs...))
}

// MemberPropertyIDGT applies the GT predicate on the "member_property_id" field.
func MemberPropertyIDGT(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGT(FieldMemberPropertyID, v))
}

// MemberPropertyIDGTE applies the GTE predicate on the "member_property_id" field.
func MemberPropertyIDGTE(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGTE(FieldMemberPropertyID, v))
}

// MemberPropertyIDLT applies the LT predicate on the "member_property_id" field.
func MemberPropertyIDLT(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLT(FieldMemberPropertyID, v))
}

// MemberPropertyIDLTE applies the LTE predicate on the "member_property_id" field.
func MemberPropertyIDLTE(v int64) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLTE(FieldMemberPropertyID, v))
}

// MemberPropertyIDIsNil applies the IsNil predicate on the "member_property_id" field.
func MemberPropertyIDIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldMemberPropertyID))
}

// MemberPropertyIDNotNil applies the NotNil predicate on the "member_property_id" field.
func MemberPropertyIDNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldMemberPropertyID))
}

// EntryTimeEQ applies the EQ predicate on the "entry_time" field.
func EntryTimeEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldEntryTime, v))
}

// EntryTimeNEQ applies the NEQ predicate on the "entry_time" field.
func EntryTimeNEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldEntryTime, v))
}

// EntryTimeIn applies the In predicate on the "entry_time" field.
func EntryTimeIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldEntryTime, vs...))
}

// EntryTimeNotIn applies the NotIn predicate on the "entry_time" field.
func EntryTimeNotIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldEntryTime, vs...))
}

// EntryTimeGT applies the GT predicate on the "entry_time" field.
func EntryTimeGT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGT(FieldEntryTime, v))
}

// EntryTimeGTE applies the GTE predicate on the "entry_time" field.
func EntryTimeGTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGTE(FieldEntryTime, v))
}

// EntryTimeLT applies the LT predicate on the "entry_time" field.
func EntryTimeLT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLT(FieldEntryTime, v))
}

// EntryTimeLTE applies the LTE predicate on the "entry_time" field.
func EntryTimeLTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLTE(FieldEntryTime, v))
}

// EntryTimeIsNil applies the IsNil predicate on the "entry_time" field.
func EntryTimeIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldEntryTime))
}

// EntryTimeNotNil applies the NotNil predicate on the "entry_time" field.
func EntryTimeNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldEntryTime))
}

// LeavingTimeEQ applies the EQ predicate on the "leaving_time" field.
func LeavingTimeEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldEQ(FieldLeavingTime, v))
}

// LeavingTimeNEQ applies the NEQ predicate on the "leaving_time" field.
func LeavingTimeNEQ(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNEQ(FieldLeavingTime, v))
}

// LeavingTimeIn applies the In predicate on the "leaving_time" field.
func LeavingTimeIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIn(FieldLeavingTime, vs...))
}

// LeavingTimeNotIn applies the NotIn predicate on the "leaving_time" field.
func LeavingTimeNotIn(vs ...time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotIn(FieldLeavingTime, vs...))
}

// LeavingTimeGT applies the GT predicate on the "leaving_time" field.
func LeavingTimeGT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGT(FieldLeavingTime, v))
}

// LeavingTimeGTE applies the GTE predicate on the "leaving_time" field.
func LeavingTimeGTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldGTE(FieldLeavingTime, v))
}

// LeavingTimeLT applies the LT predicate on the "leaving_time" field.
func LeavingTimeLT(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLT(FieldLeavingTime, v))
}

// LeavingTimeLTE applies the LTE predicate on the "leaving_time" field.
func LeavingTimeLTE(v time.Time) predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldLTE(FieldLeavingTime, v))
}

// LeavingTimeIsNil applies the IsNil predicate on the "leaving_time" field.
func LeavingTimeIsNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldIsNull(FieldLeavingTime))
}

// LeavingTimeNotNil applies the NotNil predicate on the "leaving_time" field.
func LeavingTimeNotNil() predicate.EntryLogs {
	return predicate.EntryLogs(sql.FieldNotNull(FieldLeavingTime))
}

// HasVenues applies the HasEdge predicate on the "venues" edge.
func HasVenues() predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, VenuesTable, VenuesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVenuesWith applies the HasEdge predicate on the "venues" edge with a given conditions (other predicates).
func HasVenuesWith(preds ...predicate.Venue) predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := newVenuesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMembers applies the HasEdge predicate on the "members" edge.
func HasMembers() predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MembersTable, MembersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMembersWith applies the HasEdge predicate on the "members" edge with a given conditions (other predicates).
func HasMembersWith(preds ...predicate.Member) predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := newMembersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMemberProducts applies the HasEdge predicate on the "member_products" edge.
func HasMemberProducts() predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberProductsTable, MemberProductsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberProductsWith applies the HasEdge predicate on the "member_products" edge with a given conditions (other predicates).
func HasMemberProductsWith(preds ...predicate.MemberProduct) predicate.EntryLogs {
	return predicate.EntryLogs(func(s *sql.Selector) {
		step := newMemberProductsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.EntryLogs) predicate.EntryLogs {
	return predicate.EntryLogs(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.EntryLogs) predicate.EntryLogs {
	return predicate.EntryLogs(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.EntryLogs) predicate.EntryLogs {
	return predicate.EntryLogs(sql.NotPredicates(p))
}
