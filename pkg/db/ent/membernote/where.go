// Code generated by ent, DO NOT EDIT.

package membernote

import (
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldUpdatedAt, v))
}

// Delete applies equality check predicate on the "delete" field. It's identical to DeleteEQ.
func Delete(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldDelete, v))
}

// CreatedID applies equality check predicate on the "created_id" field. It's identical to CreatedIDEQ.
func CreatedID(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldCreatedID, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldStatus, v))
}

// MemberID applies equality check predicate on the "member_id" field. It's identical to MemberIDEQ.
func MemberID(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldMemberID, v))
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldNote, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedAtIsNil applies the IsNil predicate on the "created_at" field.
func CreatedAtIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldCreatedAt))
}

// CreatedAtNotNil applies the NotNil predicate on the "created_at" field.
func CreatedAtNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldCreatedAt))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldUpdatedAt))
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldUpdatedAt))
}

// DeleteEQ applies the EQ predicate on the "delete" field.
func DeleteEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldDelete, v))
}

// DeleteNEQ applies the NEQ predicate on the "delete" field.
func DeleteNEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldDelete, v))
}

// DeleteIn applies the In predicate on the "delete" field.
func DeleteIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldDelete, vs...))
}

// DeleteNotIn applies the NotIn predicate on the "delete" field.
func DeleteNotIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldDelete, vs...))
}

// DeleteGT applies the GT predicate on the "delete" field.
func DeleteGT(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldDelete, v))
}

// DeleteGTE applies the GTE predicate on the "delete" field.
func DeleteGTE(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldDelete, v))
}

// DeleteLT applies the LT predicate on the "delete" field.
func DeleteLT(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldDelete, v))
}

// DeleteLTE applies the LTE predicate on the "delete" field.
func DeleteLTE(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldDelete, v))
}

// DeleteIsNil applies the IsNil predicate on the "delete" field.
func DeleteIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldDelete))
}

// DeleteNotNil applies the NotNil predicate on the "delete" field.
func DeleteNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldDelete))
}

// CreatedIDEQ applies the EQ predicate on the "created_id" field.
func CreatedIDEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldCreatedID, v))
}

// CreatedIDNEQ applies the NEQ predicate on the "created_id" field.
func CreatedIDNEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldCreatedID, v))
}

// CreatedIDIn applies the In predicate on the "created_id" field.
func CreatedIDIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldCreatedID, vs...))
}

// CreatedIDNotIn applies the NotIn predicate on the "created_id" field.
func CreatedIDNotIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldCreatedID, vs...))
}

// CreatedIDGT applies the GT predicate on the "created_id" field.
func CreatedIDGT(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldCreatedID, v))
}

// CreatedIDGTE applies the GTE predicate on the "created_id" field.
func CreatedIDGTE(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldCreatedID, v))
}

// CreatedIDLT applies the LT predicate on the "created_id" field.
func CreatedIDLT(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldCreatedID, v))
}

// CreatedIDLTE applies the LTE predicate on the "created_id" field.
func CreatedIDLTE(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldCreatedID, v))
}

// CreatedIDIsNil applies the IsNil predicate on the "created_id" field.
func CreatedIDIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldCreatedID))
}

// CreatedIDNotNil applies the NotNil predicate on the "created_id" field.
func CreatedIDNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldCreatedID))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldStatus))
}

// MemberIDEQ applies the EQ predicate on the "member_id" field.
func MemberIDEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldMemberID, v))
}

// MemberIDNEQ applies the NEQ predicate on the "member_id" field.
func MemberIDNEQ(v int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldMemberID, v))
}

// MemberIDIn applies the In predicate on the "member_id" field.
func MemberIDIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldMemberID, vs...))
}

// MemberIDNotIn applies the NotIn predicate on the "member_id" field.
func MemberIDNotIn(vs ...int64) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldMemberID, vs...))
}

// MemberIDIsNil applies the IsNil predicate on the "member_id" field.
func MemberIDIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldMemberID))
}

// MemberIDNotNil applies the NotNil predicate on the "member_id" field.
func MemberIDNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldMemberID))
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEQ(FieldNote, v))
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNEQ(FieldNote, v))
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIn(FieldNote, vs...))
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotIn(FieldNote, vs...))
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGT(FieldNote, v))
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldGTE(FieldNote, v))
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLT(FieldNote, v))
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldLTE(FieldNote, v))
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldContains(FieldNote, v))
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldHasPrefix(FieldNote, v))
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldHasSuffix(FieldNote, v))
}

// NoteIsNil applies the IsNil predicate on the "note" field.
func NoteIsNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldIsNull(FieldNote))
}

// NoteNotNil applies the NotNil predicate on the "note" field.
func NoteNotNil() predicate.MemberNote {
	return predicate.MemberNote(sql.FieldNotNull(FieldNote))
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldEqualFold(FieldNote, v))
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.MemberNote {
	return predicate.MemberNote(sql.FieldContainsFold(FieldNote, v))
}

// HasNotes applies the HasEdge predicate on the "notes" edge.
func HasNotes() predicate.MemberNote {
	return predicate.MemberNote(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NotesTable, NotesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNotesWith applies the HasEdge predicate on the "notes" edge with a given conditions (other predicates).
func HasNotesWith(preds ...predicate.Member) predicate.MemberNote {
	return predicate.MemberNote(func(s *sql.Selector) {
		step := newNotesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MemberNote) predicate.MemberNote {
	return predicate.MemberNote(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MemberNote) predicate.MemberNote {
	return predicate.MemberNote(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MemberNote) predicate.MemberNote {
	return predicate.MemberNote(sql.NotPredicates(p))
}
