// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/biz/dal/db/ent/memberproduct"
	"saas/biz/dal/db/ent/memberproductcourses"
	"saas/biz/dal/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProductCoursesUpdate is the builder for updating MemberProductCourses entities.
type MemberProductCoursesUpdate struct {
	config
	hooks    []Hook
	mutation *MemberProductCoursesMutation
}

// Where appends a list predicates to the MemberProductCoursesUpdate builder.
func (mpcu *MemberProductCoursesUpdate) Where(ps ...predicate.MemberProductCourses) *MemberProductCoursesUpdate {
	mpcu.mutation.Where(ps...)
	return mpcu
}

// SetUpdatedAt sets the "updated_at" field.
func (mpcu *MemberProductCoursesUpdate) SetUpdatedAt(t time.Time) *MemberProductCoursesUpdate {
	mpcu.mutation.SetUpdatedAt(t)
	return mpcu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mpcu *MemberProductCoursesUpdate) ClearUpdatedAt() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearUpdatedAt()
	return mpcu
}

// SetDelete sets the "delete" field.
func (mpcu *MemberProductCoursesUpdate) SetDelete(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.ResetDelete()
	mpcu.mutation.SetDelete(i)
	return mpcu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableDelete(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetDelete(*i)
	}
	return mpcu
}

// AddDelete adds i to the "delete" field.
func (mpcu *MemberProductCoursesUpdate) AddDelete(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.AddDelete(i)
	return mpcu
}

// ClearDelete clears the value of the "delete" field.
func (mpcu *MemberProductCoursesUpdate) ClearDelete() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearDelete()
	return mpcu
}

// SetCreatedID sets the "created_id" field.
func (mpcu *MemberProductCoursesUpdate) SetCreatedID(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.ResetCreatedID()
	mpcu.mutation.SetCreatedID(i)
	return mpcu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableCreatedID(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetCreatedID(*i)
	}
	return mpcu
}

// AddCreatedID adds i to the "created_id" field.
func (mpcu *MemberProductCoursesUpdate) AddCreatedID(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.AddCreatedID(i)
	return mpcu
}

// ClearCreatedID clears the value of the "created_id" field.
func (mpcu *MemberProductCoursesUpdate) ClearCreatedID() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearCreatedID()
	return mpcu
}

// SetStatus sets the "status" field.
func (mpcu *MemberProductCoursesUpdate) SetStatus(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.ResetStatus()
	mpcu.mutation.SetStatus(i)
	return mpcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableStatus(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetStatus(*i)
	}
	return mpcu
}

// AddStatus adds i to the "status" field.
func (mpcu *MemberProductCoursesUpdate) AddStatus(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.AddStatus(i)
	return mpcu
}

// ClearStatus clears the value of the "status" field.
func (mpcu *MemberProductCoursesUpdate) ClearStatus() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearStatus()
	return mpcu
}

// SetType sets the "type" field.
func (mpcu *MemberProductCoursesUpdate) SetType(s string) *MemberProductCoursesUpdate {
	mpcu.mutation.SetType(s)
	return mpcu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableType(s *string) *MemberProductCoursesUpdate {
	if s != nil {
		mpcu.SetType(*s)
	}
	return mpcu
}

// ClearType clears the value of the "type" field.
func (mpcu *MemberProductCoursesUpdate) ClearType() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearType()
	return mpcu
}

// SetName sets the "name" field.
func (mpcu *MemberProductCoursesUpdate) SetName(s string) *MemberProductCoursesUpdate {
	mpcu.mutation.SetName(s)
	return mpcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableName(s *string) *MemberProductCoursesUpdate {
	if s != nil {
		mpcu.SetName(*s)
	}
	return mpcu
}

// ClearName clears the value of the "name" field.
func (mpcu *MemberProductCoursesUpdate) ClearName() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearName()
	return mpcu
}

// SetNumber sets the "number" field.
func (mpcu *MemberProductCoursesUpdate) SetNumber(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.ResetNumber()
	mpcu.mutation.SetNumber(i)
	return mpcu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableNumber(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetNumber(*i)
	}
	return mpcu
}

// AddNumber adds i to the "number" field.
func (mpcu *MemberProductCoursesUpdate) AddNumber(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.AddNumber(i)
	return mpcu
}

// ClearNumber clears the value of the "number" field.
func (mpcu *MemberProductCoursesUpdate) ClearNumber() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearNumber()
	return mpcu
}

// SetNumberSurplus sets the "number_surplus" field.
func (mpcu *MemberProductCoursesUpdate) SetNumberSurplus(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.ResetNumberSurplus()
	mpcu.mutation.SetNumberSurplus(i)
	return mpcu
}

// SetNillableNumberSurplus sets the "number_surplus" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableNumberSurplus(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetNumberSurplus(*i)
	}
	return mpcu
}

// AddNumberSurplus adds i to the "number_surplus" field.
func (mpcu *MemberProductCoursesUpdate) AddNumberSurplus(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.AddNumberSurplus(i)
	return mpcu
}

// ClearNumberSurplus clears the value of the "number_surplus" field.
func (mpcu *MemberProductCoursesUpdate) ClearNumberSurplus() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearNumberSurplus()
	return mpcu
}

// SetMemberProductID sets the "member_product_id" field.
func (mpcu *MemberProductCoursesUpdate) SetMemberProductID(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.SetMemberProductID(i)
	return mpcu
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableMemberProductID(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetMemberProductID(*i)
	}
	return mpcu
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (mpcu *MemberProductCoursesUpdate) ClearMemberProductID() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearMemberProductID()
	return mpcu
}

// SetCoursesID sets the "courses_id" field.
func (mpcu *MemberProductCoursesUpdate) SetCoursesID(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.ResetCoursesID()
	mpcu.mutation.SetCoursesID(i)
	return mpcu
}

// SetNillableCoursesID sets the "courses_id" field if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableCoursesID(i *int64) *MemberProductCoursesUpdate {
	if i != nil {
		mpcu.SetCoursesID(*i)
	}
	return mpcu
}

// AddCoursesID adds i to the "courses_id" field.
func (mpcu *MemberProductCoursesUpdate) AddCoursesID(i int64) *MemberProductCoursesUpdate {
	mpcu.mutation.AddCoursesID(i)
	return mpcu
}

// ClearCoursesID clears the value of the "courses_id" field.
func (mpcu *MemberProductCoursesUpdate) ClearCoursesID() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearCoursesID()
	return mpcu
}

// SetNodeCID sets the "nodeC" edge to the MemberProduct entity by ID.
func (mpcu *MemberProductCoursesUpdate) SetNodeCID(id int64) *MemberProductCoursesUpdate {
	mpcu.mutation.SetNodeCID(id)
	return mpcu
}

// SetNillableNodeCID sets the "nodeC" edge to the MemberProduct entity by ID if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableNodeCID(id *int64) *MemberProductCoursesUpdate {
	if id != nil {
		mpcu = mpcu.SetNodeCID(*id)
	}
	return mpcu
}

// SetNodeC sets the "nodeC" edge to the MemberProduct entity.
func (mpcu *MemberProductCoursesUpdate) SetNodeC(m *MemberProduct) *MemberProductCoursesUpdate {
	return mpcu.SetNodeCID(m.ID)
}

// SetNodeLID sets the "nodeL" edge to the MemberProduct entity by ID.
func (mpcu *MemberProductCoursesUpdate) SetNodeLID(id int64) *MemberProductCoursesUpdate {
	mpcu.mutation.SetNodeLID(id)
	return mpcu
}

// SetNillableNodeLID sets the "nodeL" edge to the MemberProduct entity by ID if the given value is not nil.
func (mpcu *MemberProductCoursesUpdate) SetNillableNodeLID(id *int64) *MemberProductCoursesUpdate {
	if id != nil {
		mpcu = mpcu.SetNodeLID(*id)
	}
	return mpcu
}

// SetNodeL sets the "nodeL" edge to the MemberProduct entity.
func (mpcu *MemberProductCoursesUpdate) SetNodeL(m *MemberProduct) *MemberProductCoursesUpdate {
	return mpcu.SetNodeLID(m.ID)
}

// Mutation returns the MemberProductCoursesMutation object of the builder.
func (mpcu *MemberProductCoursesUpdate) Mutation() *MemberProductCoursesMutation {
	return mpcu.mutation
}

// ClearNodeC clears the "nodeC" edge to the MemberProduct entity.
func (mpcu *MemberProductCoursesUpdate) ClearNodeC() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearNodeC()
	return mpcu
}

// ClearNodeL clears the "nodeL" edge to the MemberProduct entity.
func (mpcu *MemberProductCoursesUpdate) ClearNodeL() *MemberProductCoursesUpdate {
	mpcu.mutation.ClearNodeL()
	return mpcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mpcu *MemberProductCoursesUpdate) Save(ctx context.Context) (int, error) {
	mpcu.defaults()
	return withHooks(ctx, mpcu.sqlSave, mpcu.mutation, mpcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpcu *MemberProductCoursesUpdate) SaveX(ctx context.Context) int {
	affected, err := mpcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mpcu *MemberProductCoursesUpdate) Exec(ctx context.Context) error {
	_, err := mpcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpcu *MemberProductCoursesUpdate) ExecX(ctx context.Context) {
	if err := mpcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpcu *MemberProductCoursesUpdate) defaults() {
	if _, ok := mpcu.mutation.UpdatedAt(); !ok && !mpcu.mutation.UpdatedAtCleared() {
		v := memberproductcourses.UpdateDefaultUpdatedAt()
		mpcu.mutation.SetUpdatedAt(v)
	}
}

func (mpcu *MemberProductCoursesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(memberproductcourses.Table, memberproductcourses.Columns, sqlgraph.NewFieldSpec(memberproductcourses.FieldID, field.TypeInt64))
	if ps := mpcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mpcu.mutation.CreatedAtCleared() {
		_spec.ClearField(memberproductcourses.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mpcu.mutation.UpdatedAt(); ok {
		_spec.SetField(memberproductcourses.FieldUpdatedAt, field.TypeTime, value)
	}
	if mpcu.mutation.UpdatedAtCleared() {
		_spec.ClearField(memberproductcourses.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mpcu.mutation.Delete(); ok {
		_spec.SetField(memberproductcourses.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := mpcu.mutation.AddedDelete(); ok {
		_spec.AddField(memberproductcourses.FieldDelete, field.TypeInt64, value)
	}
	if mpcu.mutation.DeleteCleared() {
		_spec.ClearField(memberproductcourses.FieldDelete, field.TypeInt64)
	}
	if value, ok := mpcu.mutation.CreatedID(); ok {
		_spec.SetField(memberproductcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := mpcu.mutation.AddedCreatedID(); ok {
		_spec.AddField(memberproductcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if mpcu.mutation.CreatedIDCleared() {
		_spec.ClearField(memberproductcourses.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := mpcu.mutation.Status(); ok {
		_spec.SetField(memberproductcourses.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := mpcu.mutation.AddedStatus(); ok {
		_spec.AddField(memberproductcourses.FieldStatus, field.TypeInt64, value)
	}
	if mpcu.mutation.StatusCleared() {
		_spec.ClearField(memberproductcourses.FieldStatus, field.TypeInt64)
	}
	if value, ok := mpcu.mutation.GetType(); ok {
		_spec.SetField(memberproductcourses.FieldType, field.TypeString, value)
	}
	if mpcu.mutation.TypeCleared() {
		_spec.ClearField(memberproductcourses.FieldType, field.TypeString)
	}
	if value, ok := mpcu.mutation.Name(); ok {
		_spec.SetField(memberproductcourses.FieldName, field.TypeString, value)
	}
	if mpcu.mutation.NameCleared() {
		_spec.ClearField(memberproductcourses.FieldName, field.TypeString)
	}
	if value, ok := mpcu.mutation.Number(); ok {
		_spec.SetField(memberproductcourses.FieldNumber, field.TypeInt64, value)
	}
	if value, ok := mpcu.mutation.AddedNumber(); ok {
		_spec.AddField(memberproductcourses.FieldNumber, field.TypeInt64, value)
	}
	if mpcu.mutation.NumberCleared() {
		_spec.ClearField(memberproductcourses.FieldNumber, field.TypeInt64)
	}
	if value, ok := mpcu.mutation.NumberSurplus(); ok {
		_spec.SetField(memberproductcourses.FieldNumberSurplus, field.TypeInt64, value)
	}
	if value, ok := mpcu.mutation.AddedNumberSurplus(); ok {
		_spec.AddField(memberproductcourses.FieldNumberSurplus, field.TypeInt64, value)
	}
	if mpcu.mutation.NumberSurplusCleared() {
		_spec.ClearField(memberproductcourses.FieldNumberSurplus, field.TypeInt64)
	}
	if value, ok := mpcu.mutation.CoursesID(); ok {
		_spec.SetField(memberproductcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if value, ok := mpcu.mutation.AddedCoursesID(); ok {
		_spec.AddField(memberproductcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if mpcu.mutation.CoursesIDCleared() {
		_spec.ClearField(memberproductcourses.FieldCoursesID, field.TypeInt64)
	}
	if mpcu.mutation.NodeCCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeCTable,
			Columns: []string{memberproductcourses.NodeCColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpcu.mutation.NodeCIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeCTable,
			Columns: []string{memberproductcourses.NodeCColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mpcu.mutation.NodeLCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeLTable,
			Columns: []string{memberproductcourses.NodeLColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpcu.mutation.NodeLIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeLTable,
			Columns: []string{memberproductcourses.NodeLColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mpcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberproductcourses.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mpcu.mutation.done = true
	return n, nil
}

// MemberProductCoursesUpdateOne is the builder for updating a single MemberProductCourses entity.
type MemberProductCoursesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberProductCoursesMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetUpdatedAt(t time.Time) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.SetUpdatedAt(t)
	return mpcuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearUpdatedAt() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearUpdatedAt()
	return mpcuo
}

// SetDelete sets the "delete" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetDelete(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ResetDelete()
	mpcuo.mutation.SetDelete(i)
	return mpcuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableDelete(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetDelete(*i)
	}
	return mpcuo
}

// AddDelete adds i to the "delete" field.
func (mpcuo *MemberProductCoursesUpdateOne) AddDelete(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.AddDelete(i)
	return mpcuo
}

// ClearDelete clears the value of the "delete" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearDelete() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearDelete()
	return mpcuo
}

// SetCreatedID sets the "created_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetCreatedID(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ResetCreatedID()
	mpcuo.mutation.SetCreatedID(i)
	return mpcuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableCreatedID(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetCreatedID(*i)
	}
	return mpcuo
}

// AddCreatedID adds i to the "created_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) AddCreatedID(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.AddCreatedID(i)
	return mpcuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearCreatedID() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearCreatedID()
	return mpcuo
}

// SetStatus sets the "status" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetStatus(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ResetStatus()
	mpcuo.mutation.SetStatus(i)
	return mpcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableStatus(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetStatus(*i)
	}
	return mpcuo
}

// AddStatus adds i to the "status" field.
func (mpcuo *MemberProductCoursesUpdateOne) AddStatus(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.AddStatus(i)
	return mpcuo
}

// ClearStatus clears the value of the "status" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearStatus() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearStatus()
	return mpcuo
}

// SetType sets the "type" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetType(s string) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.SetType(s)
	return mpcuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableType(s *string) *MemberProductCoursesUpdateOne {
	if s != nil {
		mpcuo.SetType(*s)
	}
	return mpcuo
}

// ClearType clears the value of the "type" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearType() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearType()
	return mpcuo
}

// SetName sets the "name" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetName(s string) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.SetName(s)
	return mpcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableName(s *string) *MemberProductCoursesUpdateOne {
	if s != nil {
		mpcuo.SetName(*s)
	}
	return mpcuo
}

// ClearName clears the value of the "name" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearName() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearName()
	return mpcuo
}

// SetNumber sets the "number" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetNumber(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ResetNumber()
	mpcuo.mutation.SetNumber(i)
	return mpcuo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableNumber(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetNumber(*i)
	}
	return mpcuo
}

// AddNumber adds i to the "number" field.
func (mpcuo *MemberProductCoursesUpdateOne) AddNumber(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.AddNumber(i)
	return mpcuo
}

// ClearNumber clears the value of the "number" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearNumber() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearNumber()
	return mpcuo
}

// SetNumberSurplus sets the "number_surplus" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetNumberSurplus(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ResetNumberSurplus()
	mpcuo.mutation.SetNumberSurplus(i)
	return mpcuo
}

// SetNillableNumberSurplus sets the "number_surplus" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableNumberSurplus(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetNumberSurplus(*i)
	}
	return mpcuo
}

// AddNumberSurplus adds i to the "number_surplus" field.
func (mpcuo *MemberProductCoursesUpdateOne) AddNumberSurplus(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.AddNumberSurplus(i)
	return mpcuo
}

// ClearNumberSurplus clears the value of the "number_surplus" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearNumberSurplus() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearNumberSurplus()
	return mpcuo
}

// SetMemberProductID sets the "member_product_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetMemberProductID(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.SetMemberProductID(i)
	return mpcuo
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableMemberProductID(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetMemberProductID(*i)
	}
	return mpcuo
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearMemberProductID() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearMemberProductID()
	return mpcuo
}

// SetCoursesID sets the "courses_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) SetCoursesID(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ResetCoursesID()
	mpcuo.mutation.SetCoursesID(i)
	return mpcuo
}

// SetNillableCoursesID sets the "courses_id" field if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableCoursesID(i *int64) *MemberProductCoursesUpdateOne {
	if i != nil {
		mpcuo.SetCoursesID(*i)
	}
	return mpcuo
}

// AddCoursesID adds i to the "courses_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) AddCoursesID(i int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.AddCoursesID(i)
	return mpcuo
}

// ClearCoursesID clears the value of the "courses_id" field.
func (mpcuo *MemberProductCoursesUpdateOne) ClearCoursesID() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearCoursesID()
	return mpcuo
}

// SetNodeCID sets the "nodeC" edge to the MemberProduct entity by ID.
func (mpcuo *MemberProductCoursesUpdateOne) SetNodeCID(id int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.SetNodeCID(id)
	return mpcuo
}

// SetNillableNodeCID sets the "nodeC" edge to the MemberProduct entity by ID if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableNodeCID(id *int64) *MemberProductCoursesUpdateOne {
	if id != nil {
		mpcuo = mpcuo.SetNodeCID(*id)
	}
	return mpcuo
}

// SetNodeC sets the "nodeC" edge to the MemberProduct entity.
func (mpcuo *MemberProductCoursesUpdateOne) SetNodeC(m *MemberProduct) *MemberProductCoursesUpdateOne {
	return mpcuo.SetNodeCID(m.ID)
}

// SetNodeLID sets the "nodeL" edge to the MemberProduct entity by ID.
func (mpcuo *MemberProductCoursesUpdateOne) SetNodeLID(id int64) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.SetNodeLID(id)
	return mpcuo
}

// SetNillableNodeLID sets the "nodeL" edge to the MemberProduct entity by ID if the given value is not nil.
func (mpcuo *MemberProductCoursesUpdateOne) SetNillableNodeLID(id *int64) *MemberProductCoursesUpdateOne {
	if id != nil {
		mpcuo = mpcuo.SetNodeLID(*id)
	}
	return mpcuo
}

// SetNodeL sets the "nodeL" edge to the MemberProduct entity.
func (mpcuo *MemberProductCoursesUpdateOne) SetNodeL(m *MemberProduct) *MemberProductCoursesUpdateOne {
	return mpcuo.SetNodeLID(m.ID)
}

// Mutation returns the MemberProductCoursesMutation object of the builder.
func (mpcuo *MemberProductCoursesUpdateOne) Mutation() *MemberProductCoursesMutation {
	return mpcuo.mutation
}

// ClearNodeC clears the "nodeC" edge to the MemberProduct entity.
func (mpcuo *MemberProductCoursesUpdateOne) ClearNodeC() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearNodeC()
	return mpcuo
}

// ClearNodeL clears the "nodeL" edge to the MemberProduct entity.
func (mpcuo *MemberProductCoursesUpdateOne) ClearNodeL() *MemberProductCoursesUpdateOne {
	mpcuo.mutation.ClearNodeL()
	return mpcuo
}

// Where appends a list predicates to the MemberProductCoursesUpdate builder.
func (mpcuo *MemberProductCoursesUpdateOne) Where(ps ...predicate.MemberProductCourses) *MemberProductCoursesUpdateOne {
	mpcuo.mutation.Where(ps...)
	return mpcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mpcuo *MemberProductCoursesUpdateOne) Select(field string, fields ...string) *MemberProductCoursesUpdateOne {
	mpcuo.fields = append([]string{field}, fields...)
	return mpcuo
}

// Save executes the query and returns the updated MemberProductCourses entity.
func (mpcuo *MemberProductCoursesUpdateOne) Save(ctx context.Context) (*MemberProductCourses, error) {
	mpcuo.defaults()
	return withHooks(ctx, mpcuo.sqlSave, mpcuo.mutation, mpcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mpcuo *MemberProductCoursesUpdateOne) SaveX(ctx context.Context) *MemberProductCourses {
	node, err := mpcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mpcuo *MemberProductCoursesUpdateOne) Exec(ctx context.Context) error {
	_, err := mpcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpcuo *MemberProductCoursesUpdateOne) ExecX(ctx context.Context) {
	if err := mpcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpcuo *MemberProductCoursesUpdateOne) defaults() {
	if _, ok := mpcuo.mutation.UpdatedAt(); !ok && !mpcuo.mutation.UpdatedAtCleared() {
		v := memberproductcourses.UpdateDefaultUpdatedAt()
		mpcuo.mutation.SetUpdatedAt(v)
	}
}

func (mpcuo *MemberProductCoursesUpdateOne) sqlSave(ctx context.Context) (_node *MemberProductCourses, err error) {
	_spec := sqlgraph.NewUpdateSpec(memberproductcourses.Table, memberproductcourses.Columns, sqlgraph.NewFieldSpec(memberproductcourses.FieldID, field.TypeInt64))
	id, ok := mpcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemberProductCourses.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mpcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memberproductcourses.FieldID)
		for _, f := range fields {
			if !memberproductcourses.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != memberproductcourses.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mpcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if mpcuo.mutation.CreatedAtCleared() {
		_spec.ClearField(memberproductcourses.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := mpcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(memberproductcourses.FieldUpdatedAt, field.TypeTime, value)
	}
	if mpcuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(memberproductcourses.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := mpcuo.mutation.Delete(); ok {
		_spec.SetField(memberproductcourses.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := mpcuo.mutation.AddedDelete(); ok {
		_spec.AddField(memberproductcourses.FieldDelete, field.TypeInt64, value)
	}
	if mpcuo.mutation.DeleteCleared() {
		_spec.ClearField(memberproductcourses.FieldDelete, field.TypeInt64)
	}
	if value, ok := mpcuo.mutation.CreatedID(); ok {
		_spec.SetField(memberproductcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := mpcuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(memberproductcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if mpcuo.mutation.CreatedIDCleared() {
		_spec.ClearField(memberproductcourses.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := mpcuo.mutation.Status(); ok {
		_spec.SetField(memberproductcourses.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := mpcuo.mutation.AddedStatus(); ok {
		_spec.AddField(memberproductcourses.FieldStatus, field.TypeInt64, value)
	}
	if mpcuo.mutation.StatusCleared() {
		_spec.ClearField(memberproductcourses.FieldStatus, field.TypeInt64)
	}
	if value, ok := mpcuo.mutation.GetType(); ok {
		_spec.SetField(memberproductcourses.FieldType, field.TypeString, value)
	}
	if mpcuo.mutation.TypeCleared() {
		_spec.ClearField(memberproductcourses.FieldType, field.TypeString)
	}
	if value, ok := mpcuo.mutation.Name(); ok {
		_spec.SetField(memberproductcourses.FieldName, field.TypeString, value)
	}
	if mpcuo.mutation.NameCleared() {
		_spec.ClearField(memberproductcourses.FieldName, field.TypeString)
	}
	if value, ok := mpcuo.mutation.Number(); ok {
		_spec.SetField(memberproductcourses.FieldNumber, field.TypeInt64, value)
	}
	if value, ok := mpcuo.mutation.AddedNumber(); ok {
		_spec.AddField(memberproductcourses.FieldNumber, field.TypeInt64, value)
	}
	if mpcuo.mutation.NumberCleared() {
		_spec.ClearField(memberproductcourses.FieldNumber, field.TypeInt64)
	}
	if value, ok := mpcuo.mutation.NumberSurplus(); ok {
		_spec.SetField(memberproductcourses.FieldNumberSurplus, field.TypeInt64, value)
	}
	if value, ok := mpcuo.mutation.AddedNumberSurplus(); ok {
		_spec.AddField(memberproductcourses.FieldNumberSurplus, field.TypeInt64, value)
	}
	if mpcuo.mutation.NumberSurplusCleared() {
		_spec.ClearField(memberproductcourses.FieldNumberSurplus, field.TypeInt64)
	}
	if value, ok := mpcuo.mutation.CoursesID(); ok {
		_spec.SetField(memberproductcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if value, ok := mpcuo.mutation.AddedCoursesID(); ok {
		_spec.AddField(memberproductcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if mpcuo.mutation.CoursesIDCleared() {
		_spec.ClearField(memberproductcourses.FieldCoursesID, field.TypeInt64)
	}
	if mpcuo.mutation.NodeCCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeCTable,
			Columns: []string{memberproductcourses.NodeCColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpcuo.mutation.NodeCIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeCTable,
			Columns: []string{memberproductcourses.NodeCColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mpcuo.mutation.NodeLCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeLTable,
			Columns: []string{memberproductcourses.NodeLColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mpcuo.mutation.NodeLIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductcourses.NodeLTable,
			Columns: []string{memberproductcourses.NodeLColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MemberProductCourses{config: mpcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mpcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberproductcourses.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mpcuo.mutation.done = true
	return _node, nil
}
