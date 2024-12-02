// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/contest"
	"saas/pkg/db/ent/contestparticipant"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContestUpdate is the builder for updating Contest entities.
type ContestUpdate struct {
	config
	hooks    []Hook
	mutation *ContestMutation
}

// Where appends a list predicates to the ContestUpdate builder.
func (cu *ContestUpdate) Where(ps ...predicate.Contest) *ContestUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ContestUpdate) SetUpdatedAt(t time.Time) *ContestUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetStatus sets the "status" field.
func (cu *ContestUpdate) SetStatus(i int64) *ContestUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(i)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableStatus(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetStatus(*i)
	}
	return cu
}

// AddStatus adds i to the "status" field.
func (cu *ContestUpdate) AddStatus(i int64) *ContestUpdate {
	cu.mutation.AddStatus(i)
	return cu
}

// ClearStatus clears the value of the "status" field.
func (cu *ContestUpdate) ClearStatus() *ContestUpdate {
	cu.mutation.ClearStatus()
	return cu
}

// SetName sets the "name" field.
func (cu *ContestUpdate) SetName(s string) *ContestUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableName(s *string) *ContestUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// ClearName clears the value of the "name" field.
func (cu *ContestUpdate) ClearName() *ContestUpdate {
	cu.mutation.ClearName()
	return cu
}

// SetSignNumber sets the "sign_number" field.
func (cu *ContestUpdate) SetSignNumber(i int64) *ContestUpdate {
	cu.mutation.ResetSignNumber()
	cu.mutation.SetSignNumber(i)
	return cu
}

// SetNillableSignNumber sets the "sign_number" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableSignNumber(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetSignNumber(*i)
	}
	return cu
}

// AddSignNumber adds i to the "sign_number" field.
func (cu *ContestUpdate) AddSignNumber(i int64) *ContestUpdate {
	cu.mutation.AddSignNumber(i)
	return cu
}

// ClearSignNumber clears the value of the "sign_number" field.
func (cu *ContestUpdate) ClearSignNumber() *ContestUpdate {
	cu.mutation.ClearSignNumber()
	return cu
}

// SetSignStartAt sets the "sign_start_at" field.
func (cu *ContestUpdate) SetSignStartAt(t time.Time) *ContestUpdate {
	cu.mutation.SetSignStartAt(t)
	return cu
}

// SetNillableSignStartAt sets the "sign_start_at" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableSignStartAt(t *time.Time) *ContestUpdate {
	if t != nil {
		cu.SetSignStartAt(*t)
	}
	return cu
}

// ClearSignStartAt clears the value of the "sign_start_at" field.
func (cu *ContestUpdate) ClearSignStartAt() *ContestUpdate {
	cu.mutation.ClearSignStartAt()
	return cu
}

// SetSignEndAt sets the "sign_end_at" field.
func (cu *ContestUpdate) SetSignEndAt(t time.Time) *ContestUpdate {
	cu.mutation.SetSignEndAt(t)
	return cu
}

// SetNillableSignEndAt sets the "sign_end_at" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableSignEndAt(t *time.Time) *ContestUpdate {
	if t != nil {
		cu.SetSignEndAt(*t)
	}
	return cu
}

// ClearSignEndAt clears the value of the "sign_end_at" field.
func (cu *ContestUpdate) ClearSignEndAt() *ContestUpdate {
	cu.mutation.ClearSignEndAt()
	return cu
}

// SetNumber sets the "number" field.
func (cu *ContestUpdate) SetNumber(i int64) *ContestUpdate {
	cu.mutation.ResetNumber()
	cu.mutation.SetNumber(i)
	return cu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableNumber(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetNumber(*i)
	}
	return cu
}

// AddNumber adds i to the "number" field.
func (cu *ContestUpdate) AddNumber(i int64) *ContestUpdate {
	cu.mutation.AddNumber(i)
	return cu
}

// ClearNumber clears the value of the "number" field.
func (cu *ContestUpdate) ClearNumber() *ContestUpdate {
	cu.mutation.ClearNumber()
	return cu
}

// SetStartAt sets the "start_at" field.
func (cu *ContestUpdate) SetStartAt(t time.Time) *ContestUpdate {
	cu.mutation.SetStartAt(t)
	return cu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableStartAt(t *time.Time) *ContestUpdate {
	if t != nil {
		cu.SetStartAt(*t)
	}
	return cu
}

// ClearStartAt clears the value of the "start_at" field.
func (cu *ContestUpdate) ClearStartAt() *ContestUpdate {
	cu.mutation.ClearStartAt()
	return cu
}

// SetEndAt sets the "end_at" field.
func (cu *ContestUpdate) SetEndAt(t time.Time) *ContestUpdate {
	cu.mutation.SetEndAt(t)
	return cu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableEndAt(t *time.Time) *ContestUpdate {
	if t != nil {
		cu.SetEndAt(*t)
	}
	return cu
}

// ClearEndAt clears the value of the "end_at" field.
func (cu *ContestUpdate) ClearEndAt() *ContestUpdate {
	cu.mutation.ClearEndAt()
	return cu
}

// SetPic sets the "pic" field.
func (cu *ContestUpdate) SetPic(s string) *ContestUpdate {
	cu.mutation.SetPic(s)
	return cu
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (cu *ContestUpdate) SetNillablePic(s *string) *ContestUpdate {
	if s != nil {
		cu.SetPic(*s)
	}
	return cu
}

// ClearPic clears the value of the "pic" field.
func (cu *ContestUpdate) ClearPic() *ContestUpdate {
	cu.mutation.ClearPic()
	return cu
}

// SetSponsor sets the "sponsor" field.
func (cu *ContestUpdate) SetSponsor(s string) *ContestUpdate {
	cu.mutation.SetSponsor(s)
	return cu
}

// SetNillableSponsor sets the "sponsor" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableSponsor(s *string) *ContestUpdate {
	if s != nil {
		cu.SetSponsor(*s)
	}
	return cu
}

// ClearSponsor clears the value of the "sponsor" field.
func (cu *ContestUpdate) ClearSponsor() *ContestUpdate {
	cu.mutation.ClearSponsor()
	return cu
}

// SetFee sets the "fee" field.
func (cu *ContestUpdate) SetFee(f float64) *ContestUpdate {
	cu.mutation.ResetFee()
	cu.mutation.SetFee(f)
	return cu
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableFee(f *float64) *ContestUpdate {
	if f != nil {
		cu.SetFee(*f)
	}
	return cu
}

// AddFee adds f to the "fee" field.
func (cu *ContestUpdate) AddFee(f float64) *ContestUpdate {
	cu.mutation.AddFee(f)
	return cu
}

// ClearFee clears the value of the "fee" field.
func (cu *ContestUpdate) ClearFee() *ContestUpdate {
	cu.mutation.ClearFee()
	return cu
}

// SetIsCancel sets the "is_cancel" field.
func (cu *ContestUpdate) SetIsCancel(i int64) *ContestUpdate {
	cu.mutation.ResetIsCancel()
	cu.mutation.SetIsCancel(i)
	return cu
}

// SetNillableIsCancel sets the "is_cancel" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableIsCancel(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetIsCancel(*i)
	}
	return cu
}

// AddIsCancel adds i to the "is_cancel" field.
func (cu *ContestUpdate) AddIsCancel(i int64) *ContestUpdate {
	cu.mutation.AddIsCancel(i)
	return cu
}

// ClearIsCancel clears the value of the "is_cancel" field.
func (cu *ContestUpdate) ClearIsCancel() *ContestUpdate {
	cu.mutation.ClearIsCancel()
	return cu
}

// SetCancelTime sets the "cancel_time" field.
func (cu *ContestUpdate) SetCancelTime(i int64) *ContestUpdate {
	cu.mutation.ResetCancelTime()
	cu.mutation.SetCancelTime(i)
	return cu
}

// SetNillableCancelTime sets the "cancel_time" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableCancelTime(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetCancelTime(*i)
	}
	return cu
}

// AddCancelTime adds i to the "cancel_time" field.
func (cu *ContestUpdate) AddCancelTime(i int64) *ContestUpdate {
	cu.mutation.AddCancelTime(i)
	return cu
}

// ClearCancelTime clears the value of the "cancel_time" field.
func (cu *ContestUpdate) ClearCancelTime() *ContestUpdate {
	cu.mutation.ClearCancelTime()
	return cu
}

// SetDetail sets the "detail" field.
func (cu *ContestUpdate) SetDetail(s string) *ContestUpdate {
	cu.mutation.SetDetail(s)
	return cu
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableDetail(s *string) *ContestUpdate {
	if s != nil {
		cu.SetDetail(*s)
	}
	return cu
}

// ClearDetail clears the value of the "detail" field.
func (cu *ContestUpdate) ClearDetail() *ContestUpdate {
	cu.mutation.ClearDetail()
	return cu
}

// SetSignFields sets the "sign_fields" field.
func (cu *ContestUpdate) SetSignFields(s string) *ContestUpdate {
	cu.mutation.SetSignFields(s)
	return cu
}

// SetNillableSignFields sets the "sign_fields" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableSignFields(s *string) *ContestUpdate {
	if s != nil {
		cu.SetSignFields(*s)
	}
	return cu
}

// ClearSignFields clears the value of the "sign_fields" field.
func (cu *ContestUpdate) ClearSignFields() *ContestUpdate {
	cu.mutation.ClearSignFields()
	return cu
}

// SetCondition sets the "condition" field.
func (cu *ContestUpdate) SetCondition(i int64) *ContestUpdate {
	cu.mutation.ResetCondition()
	cu.mutation.SetCondition(i)
	return cu
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (cu *ContestUpdate) SetNillableCondition(i *int64) *ContestUpdate {
	if i != nil {
		cu.SetCondition(*i)
	}
	return cu
}

// AddCondition adds i to the "condition" field.
func (cu *ContestUpdate) AddCondition(i int64) *ContestUpdate {
	cu.mutation.AddCondition(i)
	return cu
}

// ClearCondition clears the value of the "condition" field.
func (cu *ContestUpdate) ClearCondition() *ContestUpdate {
	cu.mutation.ClearCondition()
	return cu
}

// AddContestParticipantIDs adds the "contest_participants" edge to the ContestParticipant entity by IDs.
func (cu *ContestUpdate) AddContestParticipantIDs(ids ...int64) *ContestUpdate {
	cu.mutation.AddContestParticipantIDs(ids...)
	return cu
}

// AddContestParticipants adds the "contest_participants" edges to the ContestParticipant entity.
func (cu *ContestUpdate) AddContestParticipants(c ...*ContestParticipant) *ContestUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddContestParticipantIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cu *ContestUpdate) Mutation() *ContestMutation {
	return cu.mutation
}

// ClearContestParticipants clears all "contest_participants" edges to the ContestParticipant entity.
func (cu *ContestUpdate) ClearContestParticipants() *ContestUpdate {
	cu.mutation.ClearContestParticipants()
	return cu
}

// RemoveContestParticipantIDs removes the "contest_participants" edge to ContestParticipant entities by IDs.
func (cu *ContestUpdate) RemoveContestParticipantIDs(ids ...int64) *ContestUpdate {
	cu.mutation.RemoveContestParticipantIDs(ids...)
	return cu
}

// RemoveContestParticipants removes "contest_participants" edges to ContestParticipant entities.
func (cu *ContestUpdate) RemoveContestParticipants(c ...*ContestParticipant) *ContestUpdate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveContestParticipantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContestUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContestUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContestUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContestUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ContestUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := contest.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ContestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(contest.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.AddField(contest.FieldStatus, field.TypeInt64, value)
	}
	if cu.mutation.StatusCleared() {
		_spec.ClearField(contest.FieldStatus, field.TypeInt64)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(contest.FieldName, field.TypeString, value)
	}
	if cu.mutation.NameCleared() {
		_spec.ClearField(contest.FieldName, field.TypeString)
	}
	if value, ok := cu.mutation.SignNumber(); ok {
		_spec.SetField(contest.FieldSignNumber, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedSignNumber(); ok {
		_spec.AddField(contest.FieldSignNumber, field.TypeInt64, value)
	}
	if cu.mutation.SignNumberCleared() {
		_spec.ClearField(contest.FieldSignNumber, field.TypeInt64)
	}
	if value, ok := cu.mutation.SignStartAt(); ok {
		_spec.SetField(contest.FieldSignStartAt, field.TypeTime, value)
	}
	if cu.mutation.SignStartAtCleared() {
		_spec.ClearField(contest.FieldSignStartAt, field.TypeTime)
	}
	if value, ok := cu.mutation.SignEndAt(); ok {
		_spec.SetField(contest.FieldSignEndAt, field.TypeTime, value)
	}
	if cu.mutation.SignEndAtCleared() {
		_spec.ClearField(contest.FieldSignEndAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Number(); ok {
		_spec.SetField(contest.FieldNumber, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedNumber(); ok {
		_spec.AddField(contest.FieldNumber, field.TypeInt64, value)
	}
	if cu.mutation.NumberCleared() {
		_spec.ClearField(contest.FieldNumber, field.TypeInt64)
	}
	if value, ok := cu.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
	}
	if cu.mutation.StartAtCleared() {
		_spec.ClearField(contest.FieldStartAt, field.TypeTime)
	}
	if value, ok := cu.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
	}
	if cu.mutation.EndAtCleared() {
		_spec.ClearField(contest.FieldEndAt, field.TypeTime)
	}
	if value, ok := cu.mutation.Pic(); ok {
		_spec.SetField(contest.FieldPic, field.TypeString, value)
	}
	if cu.mutation.PicCleared() {
		_spec.ClearField(contest.FieldPic, field.TypeString)
	}
	if value, ok := cu.mutation.Sponsor(); ok {
		_spec.SetField(contest.FieldSponsor, field.TypeString, value)
	}
	if cu.mutation.SponsorCleared() {
		_spec.ClearField(contest.FieldSponsor, field.TypeString)
	}
	if value, ok := cu.mutation.Fee(); ok {
		_spec.SetField(contest.FieldFee, field.TypeFloat64, value)
	}
	if value, ok := cu.mutation.AddedFee(); ok {
		_spec.AddField(contest.FieldFee, field.TypeFloat64, value)
	}
	if cu.mutation.FeeCleared() {
		_spec.ClearField(contest.FieldFee, field.TypeFloat64)
	}
	if value, ok := cu.mutation.IsCancel(); ok {
		_spec.SetField(contest.FieldIsCancel, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedIsCancel(); ok {
		_spec.AddField(contest.FieldIsCancel, field.TypeInt64, value)
	}
	if cu.mutation.IsCancelCleared() {
		_spec.ClearField(contest.FieldIsCancel, field.TypeInt64)
	}
	if value, ok := cu.mutation.CancelTime(); ok {
		_spec.SetField(contest.FieldCancelTime, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedCancelTime(); ok {
		_spec.AddField(contest.FieldCancelTime, field.TypeInt64, value)
	}
	if cu.mutation.CancelTimeCleared() {
		_spec.ClearField(contest.FieldCancelTime, field.TypeInt64)
	}
	if value, ok := cu.mutation.Detail(); ok {
		_spec.SetField(contest.FieldDetail, field.TypeString, value)
	}
	if cu.mutation.DetailCleared() {
		_spec.ClearField(contest.FieldDetail, field.TypeString)
	}
	if value, ok := cu.mutation.SignFields(); ok {
		_spec.SetField(contest.FieldSignFields, field.TypeString, value)
	}
	if cu.mutation.SignFieldsCleared() {
		_spec.ClearField(contest.FieldSignFields, field.TypeString)
	}
	if value, ok := cu.mutation.Condition(); ok {
		_spec.SetField(contest.FieldCondition, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedCondition(); ok {
		_spec.AddField(contest.FieldCondition, field.TypeInt64, value)
	}
	if cu.mutation.ConditionCleared() {
		_spec.ClearField(contest.FieldCondition, field.TypeInt64)
	}
	if cu.mutation.ContestParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedContestParticipantsIDs(); len(nodes) > 0 && !cu.mutation.ContestParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ContestParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ContestUpdateOne is the builder for updating a single Contest entity.
type ContestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContestMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ContestUpdateOne) SetUpdatedAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *ContestUpdateOne) SetStatus(i int64) *ContestUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(i)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableStatus(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetStatus(*i)
	}
	return cuo
}

// AddStatus adds i to the "status" field.
func (cuo *ContestUpdateOne) AddStatus(i int64) *ContestUpdateOne {
	cuo.mutation.AddStatus(i)
	return cuo
}

// ClearStatus clears the value of the "status" field.
func (cuo *ContestUpdateOne) ClearStatus() *ContestUpdateOne {
	cuo.mutation.ClearStatus()
	return cuo
}

// SetName sets the "name" field.
func (cuo *ContestUpdateOne) SetName(s string) *ContestUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableName(s *string) *ContestUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// ClearName clears the value of the "name" field.
func (cuo *ContestUpdateOne) ClearName() *ContestUpdateOne {
	cuo.mutation.ClearName()
	return cuo
}

// SetSignNumber sets the "sign_number" field.
func (cuo *ContestUpdateOne) SetSignNumber(i int64) *ContestUpdateOne {
	cuo.mutation.ResetSignNumber()
	cuo.mutation.SetSignNumber(i)
	return cuo
}

// SetNillableSignNumber sets the "sign_number" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableSignNumber(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetSignNumber(*i)
	}
	return cuo
}

// AddSignNumber adds i to the "sign_number" field.
func (cuo *ContestUpdateOne) AddSignNumber(i int64) *ContestUpdateOne {
	cuo.mutation.AddSignNumber(i)
	return cuo
}

// ClearSignNumber clears the value of the "sign_number" field.
func (cuo *ContestUpdateOne) ClearSignNumber() *ContestUpdateOne {
	cuo.mutation.ClearSignNumber()
	return cuo
}

// SetSignStartAt sets the "sign_start_at" field.
func (cuo *ContestUpdateOne) SetSignStartAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetSignStartAt(t)
	return cuo
}

// SetNillableSignStartAt sets the "sign_start_at" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableSignStartAt(t *time.Time) *ContestUpdateOne {
	if t != nil {
		cuo.SetSignStartAt(*t)
	}
	return cuo
}

// ClearSignStartAt clears the value of the "sign_start_at" field.
func (cuo *ContestUpdateOne) ClearSignStartAt() *ContestUpdateOne {
	cuo.mutation.ClearSignStartAt()
	return cuo
}

// SetSignEndAt sets the "sign_end_at" field.
func (cuo *ContestUpdateOne) SetSignEndAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetSignEndAt(t)
	return cuo
}

// SetNillableSignEndAt sets the "sign_end_at" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableSignEndAt(t *time.Time) *ContestUpdateOne {
	if t != nil {
		cuo.SetSignEndAt(*t)
	}
	return cuo
}

// ClearSignEndAt clears the value of the "sign_end_at" field.
func (cuo *ContestUpdateOne) ClearSignEndAt() *ContestUpdateOne {
	cuo.mutation.ClearSignEndAt()
	return cuo
}

// SetNumber sets the "number" field.
func (cuo *ContestUpdateOne) SetNumber(i int64) *ContestUpdateOne {
	cuo.mutation.ResetNumber()
	cuo.mutation.SetNumber(i)
	return cuo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableNumber(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetNumber(*i)
	}
	return cuo
}

// AddNumber adds i to the "number" field.
func (cuo *ContestUpdateOne) AddNumber(i int64) *ContestUpdateOne {
	cuo.mutation.AddNumber(i)
	return cuo
}

// ClearNumber clears the value of the "number" field.
func (cuo *ContestUpdateOne) ClearNumber() *ContestUpdateOne {
	cuo.mutation.ClearNumber()
	return cuo
}

// SetStartAt sets the "start_at" field.
func (cuo *ContestUpdateOne) SetStartAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetStartAt(t)
	return cuo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableStartAt(t *time.Time) *ContestUpdateOne {
	if t != nil {
		cuo.SetStartAt(*t)
	}
	return cuo
}

// ClearStartAt clears the value of the "start_at" field.
func (cuo *ContestUpdateOne) ClearStartAt() *ContestUpdateOne {
	cuo.mutation.ClearStartAt()
	return cuo
}

// SetEndAt sets the "end_at" field.
func (cuo *ContestUpdateOne) SetEndAt(t time.Time) *ContestUpdateOne {
	cuo.mutation.SetEndAt(t)
	return cuo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableEndAt(t *time.Time) *ContestUpdateOne {
	if t != nil {
		cuo.SetEndAt(*t)
	}
	return cuo
}

// ClearEndAt clears the value of the "end_at" field.
func (cuo *ContestUpdateOne) ClearEndAt() *ContestUpdateOne {
	cuo.mutation.ClearEndAt()
	return cuo
}

// SetPic sets the "pic" field.
func (cuo *ContestUpdateOne) SetPic(s string) *ContestUpdateOne {
	cuo.mutation.SetPic(s)
	return cuo
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillablePic(s *string) *ContestUpdateOne {
	if s != nil {
		cuo.SetPic(*s)
	}
	return cuo
}

// ClearPic clears the value of the "pic" field.
func (cuo *ContestUpdateOne) ClearPic() *ContestUpdateOne {
	cuo.mutation.ClearPic()
	return cuo
}

// SetSponsor sets the "sponsor" field.
func (cuo *ContestUpdateOne) SetSponsor(s string) *ContestUpdateOne {
	cuo.mutation.SetSponsor(s)
	return cuo
}

// SetNillableSponsor sets the "sponsor" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableSponsor(s *string) *ContestUpdateOne {
	if s != nil {
		cuo.SetSponsor(*s)
	}
	return cuo
}

// ClearSponsor clears the value of the "sponsor" field.
func (cuo *ContestUpdateOne) ClearSponsor() *ContestUpdateOne {
	cuo.mutation.ClearSponsor()
	return cuo
}

// SetFee sets the "fee" field.
func (cuo *ContestUpdateOne) SetFee(f float64) *ContestUpdateOne {
	cuo.mutation.ResetFee()
	cuo.mutation.SetFee(f)
	return cuo
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableFee(f *float64) *ContestUpdateOne {
	if f != nil {
		cuo.SetFee(*f)
	}
	return cuo
}

// AddFee adds f to the "fee" field.
func (cuo *ContestUpdateOne) AddFee(f float64) *ContestUpdateOne {
	cuo.mutation.AddFee(f)
	return cuo
}

// ClearFee clears the value of the "fee" field.
func (cuo *ContestUpdateOne) ClearFee() *ContestUpdateOne {
	cuo.mutation.ClearFee()
	return cuo
}

// SetIsCancel sets the "is_cancel" field.
func (cuo *ContestUpdateOne) SetIsCancel(i int64) *ContestUpdateOne {
	cuo.mutation.ResetIsCancel()
	cuo.mutation.SetIsCancel(i)
	return cuo
}

// SetNillableIsCancel sets the "is_cancel" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableIsCancel(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetIsCancel(*i)
	}
	return cuo
}

// AddIsCancel adds i to the "is_cancel" field.
func (cuo *ContestUpdateOne) AddIsCancel(i int64) *ContestUpdateOne {
	cuo.mutation.AddIsCancel(i)
	return cuo
}

// ClearIsCancel clears the value of the "is_cancel" field.
func (cuo *ContestUpdateOne) ClearIsCancel() *ContestUpdateOne {
	cuo.mutation.ClearIsCancel()
	return cuo
}

// SetCancelTime sets the "cancel_time" field.
func (cuo *ContestUpdateOne) SetCancelTime(i int64) *ContestUpdateOne {
	cuo.mutation.ResetCancelTime()
	cuo.mutation.SetCancelTime(i)
	return cuo
}

// SetNillableCancelTime sets the "cancel_time" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableCancelTime(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetCancelTime(*i)
	}
	return cuo
}

// AddCancelTime adds i to the "cancel_time" field.
func (cuo *ContestUpdateOne) AddCancelTime(i int64) *ContestUpdateOne {
	cuo.mutation.AddCancelTime(i)
	return cuo
}

// ClearCancelTime clears the value of the "cancel_time" field.
func (cuo *ContestUpdateOne) ClearCancelTime() *ContestUpdateOne {
	cuo.mutation.ClearCancelTime()
	return cuo
}

// SetDetail sets the "detail" field.
func (cuo *ContestUpdateOne) SetDetail(s string) *ContestUpdateOne {
	cuo.mutation.SetDetail(s)
	return cuo
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableDetail(s *string) *ContestUpdateOne {
	if s != nil {
		cuo.SetDetail(*s)
	}
	return cuo
}

// ClearDetail clears the value of the "detail" field.
func (cuo *ContestUpdateOne) ClearDetail() *ContestUpdateOne {
	cuo.mutation.ClearDetail()
	return cuo
}

// SetSignFields sets the "sign_fields" field.
func (cuo *ContestUpdateOne) SetSignFields(s string) *ContestUpdateOne {
	cuo.mutation.SetSignFields(s)
	return cuo
}

// SetNillableSignFields sets the "sign_fields" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableSignFields(s *string) *ContestUpdateOne {
	if s != nil {
		cuo.SetSignFields(*s)
	}
	return cuo
}

// ClearSignFields clears the value of the "sign_fields" field.
func (cuo *ContestUpdateOne) ClearSignFields() *ContestUpdateOne {
	cuo.mutation.ClearSignFields()
	return cuo
}

// SetCondition sets the "condition" field.
func (cuo *ContestUpdateOne) SetCondition(i int64) *ContestUpdateOne {
	cuo.mutation.ResetCondition()
	cuo.mutation.SetCondition(i)
	return cuo
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (cuo *ContestUpdateOne) SetNillableCondition(i *int64) *ContestUpdateOne {
	if i != nil {
		cuo.SetCondition(*i)
	}
	return cuo
}

// AddCondition adds i to the "condition" field.
func (cuo *ContestUpdateOne) AddCondition(i int64) *ContestUpdateOne {
	cuo.mutation.AddCondition(i)
	return cuo
}

// ClearCondition clears the value of the "condition" field.
func (cuo *ContestUpdateOne) ClearCondition() *ContestUpdateOne {
	cuo.mutation.ClearCondition()
	return cuo
}

// AddContestParticipantIDs adds the "contest_participants" edge to the ContestParticipant entity by IDs.
func (cuo *ContestUpdateOne) AddContestParticipantIDs(ids ...int64) *ContestUpdateOne {
	cuo.mutation.AddContestParticipantIDs(ids...)
	return cuo
}

// AddContestParticipants adds the "contest_participants" edges to the ContestParticipant entity.
func (cuo *ContestUpdateOne) AddContestParticipants(c ...*ContestParticipant) *ContestUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddContestParticipantIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cuo *ContestUpdateOne) Mutation() *ContestMutation {
	return cuo.mutation
}

// ClearContestParticipants clears all "contest_participants" edges to the ContestParticipant entity.
func (cuo *ContestUpdateOne) ClearContestParticipants() *ContestUpdateOne {
	cuo.mutation.ClearContestParticipants()
	return cuo
}

// RemoveContestParticipantIDs removes the "contest_participants" edge to ContestParticipant entities by IDs.
func (cuo *ContestUpdateOne) RemoveContestParticipantIDs(ids ...int64) *ContestUpdateOne {
	cuo.mutation.RemoveContestParticipantIDs(ids...)
	return cuo
}

// RemoveContestParticipants removes "contest_participants" edges to ContestParticipant entities.
func (cuo *ContestUpdateOne) RemoveContestParticipants(c ...*ContestParticipant) *ContestUpdateOne {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveContestParticipantIDs(ids...)
}

// Where appends a list predicates to the ContestUpdate builder.
func (cuo *ContestUpdateOne) Where(ps ...predicate.Contest) *ContestUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContestUpdateOne) Select(field string, fields ...string) *ContestUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Contest entity.
func (cuo *ContestUpdateOne) Save(ctx context.Context) (*Contest, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContestUpdateOne) SaveX(ctx context.Context) *Contest {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContestUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContestUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ContestUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := contest.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ContestUpdateOne) sqlSave(ctx context.Context) (_node *Contest, err error) {
	_spec := sqlgraph.NewUpdateSpec(contest.Table, contest.Columns, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Contest.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contest.FieldID)
		for _, f := range fields {
			if !contest.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(contest.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.AddField(contest.FieldStatus, field.TypeInt64, value)
	}
	if cuo.mutation.StatusCleared() {
		_spec.ClearField(contest.FieldStatus, field.TypeInt64)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(contest.FieldName, field.TypeString, value)
	}
	if cuo.mutation.NameCleared() {
		_spec.ClearField(contest.FieldName, field.TypeString)
	}
	if value, ok := cuo.mutation.SignNumber(); ok {
		_spec.SetField(contest.FieldSignNumber, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedSignNumber(); ok {
		_spec.AddField(contest.FieldSignNumber, field.TypeInt64, value)
	}
	if cuo.mutation.SignNumberCleared() {
		_spec.ClearField(contest.FieldSignNumber, field.TypeInt64)
	}
	if value, ok := cuo.mutation.SignStartAt(); ok {
		_spec.SetField(contest.FieldSignStartAt, field.TypeTime, value)
	}
	if cuo.mutation.SignStartAtCleared() {
		_spec.ClearField(contest.FieldSignStartAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.SignEndAt(); ok {
		_spec.SetField(contest.FieldSignEndAt, field.TypeTime, value)
	}
	if cuo.mutation.SignEndAtCleared() {
		_spec.ClearField(contest.FieldSignEndAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Number(); ok {
		_spec.SetField(contest.FieldNumber, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedNumber(); ok {
		_spec.AddField(contest.FieldNumber, field.TypeInt64, value)
	}
	if cuo.mutation.NumberCleared() {
		_spec.ClearField(contest.FieldNumber, field.TypeInt64)
	}
	if value, ok := cuo.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
	}
	if cuo.mutation.StartAtCleared() {
		_spec.ClearField(contest.FieldStartAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
	}
	if cuo.mutation.EndAtCleared() {
		_spec.ClearField(contest.FieldEndAt, field.TypeTime)
	}
	if value, ok := cuo.mutation.Pic(); ok {
		_spec.SetField(contest.FieldPic, field.TypeString, value)
	}
	if cuo.mutation.PicCleared() {
		_spec.ClearField(contest.FieldPic, field.TypeString)
	}
	if value, ok := cuo.mutation.Sponsor(); ok {
		_spec.SetField(contest.FieldSponsor, field.TypeString, value)
	}
	if cuo.mutation.SponsorCleared() {
		_spec.ClearField(contest.FieldSponsor, field.TypeString)
	}
	if value, ok := cuo.mutation.Fee(); ok {
		_spec.SetField(contest.FieldFee, field.TypeFloat64, value)
	}
	if value, ok := cuo.mutation.AddedFee(); ok {
		_spec.AddField(contest.FieldFee, field.TypeFloat64, value)
	}
	if cuo.mutation.FeeCleared() {
		_spec.ClearField(contest.FieldFee, field.TypeFloat64)
	}
	if value, ok := cuo.mutation.IsCancel(); ok {
		_spec.SetField(contest.FieldIsCancel, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedIsCancel(); ok {
		_spec.AddField(contest.FieldIsCancel, field.TypeInt64, value)
	}
	if cuo.mutation.IsCancelCleared() {
		_spec.ClearField(contest.FieldIsCancel, field.TypeInt64)
	}
	if value, ok := cuo.mutation.CancelTime(); ok {
		_spec.SetField(contest.FieldCancelTime, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedCancelTime(); ok {
		_spec.AddField(contest.FieldCancelTime, field.TypeInt64, value)
	}
	if cuo.mutation.CancelTimeCleared() {
		_spec.ClearField(contest.FieldCancelTime, field.TypeInt64)
	}
	if value, ok := cuo.mutation.Detail(); ok {
		_spec.SetField(contest.FieldDetail, field.TypeString, value)
	}
	if cuo.mutation.DetailCleared() {
		_spec.ClearField(contest.FieldDetail, field.TypeString)
	}
	if value, ok := cuo.mutation.SignFields(); ok {
		_spec.SetField(contest.FieldSignFields, field.TypeString, value)
	}
	if cuo.mutation.SignFieldsCleared() {
		_spec.ClearField(contest.FieldSignFields, field.TypeString)
	}
	if value, ok := cuo.mutation.Condition(); ok {
		_spec.SetField(contest.FieldCondition, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedCondition(); ok {
		_spec.AddField(contest.FieldCondition, field.TypeInt64, value)
	}
	if cuo.mutation.ConditionCleared() {
		_spec.ClearField(contest.FieldCondition, field.TypeInt64)
	}
	if cuo.mutation.ContestParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedContestParticipantsIDs(); len(nodes) > 0 && !cuo.mutation.ContestParticipantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ContestParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Contest{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contest.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
