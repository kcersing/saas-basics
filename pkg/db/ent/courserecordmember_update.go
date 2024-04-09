// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/courserecordmember"
	"saas/pkg/db/ent/courserecordschedule"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseRecordMemberUpdate is the builder for updating CourseRecordMember entities.
type CourseRecordMemberUpdate struct {
	config
	hooks    []Hook
	mutation *CourseRecordMemberMutation
}

// Where appends a list predicates to the CourseRecordMemberUpdate builder.
func (crmu *CourseRecordMemberUpdate) Where(ps ...predicate.CourseRecordMember) *CourseRecordMemberUpdate {
	crmu.mutation.Where(ps...)
	return crmu
}

// SetUpdatedAt sets the "updated_at" field.
func (crmu *CourseRecordMemberUpdate) SetUpdatedAt(t time.Time) *CourseRecordMemberUpdate {
	crmu.mutation.SetUpdatedAt(t)
	return crmu
}

// SetStatus sets the "status" field.
func (crmu *CourseRecordMemberUpdate) SetStatus(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.ResetStatus()
	crmu.mutation.SetStatus(i)
	return crmu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableStatus(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetStatus(*i)
	}
	return crmu
}

// AddStatus adds i to the "status" field.
func (crmu *CourseRecordMemberUpdate) AddStatus(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.AddStatus(i)
	return crmu
}

// ClearStatus clears the value of the "status" field.
func (crmu *CourseRecordMemberUpdate) ClearStatus() *CourseRecordMemberUpdate {
	crmu.mutation.ClearStatus()
	return crmu
}

// SetVenueID sets the "venue_id" field.
func (crmu *CourseRecordMemberUpdate) SetVenueID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.ResetVenueID()
	crmu.mutation.SetVenueID(i)
	return crmu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableVenueID(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetVenueID(*i)
	}
	return crmu
}

// AddVenueID adds i to the "venue_id" field.
func (crmu *CourseRecordMemberUpdate) AddVenueID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.AddVenueID(i)
	return crmu
}

// ClearVenueID clears the value of the "venue_id" field.
func (crmu *CourseRecordMemberUpdate) ClearVenueID() *CourseRecordMemberUpdate {
	crmu.mutation.ClearVenueID()
	return crmu
}

// SetMemberID sets the "member_id" field.
func (crmu *CourseRecordMemberUpdate) SetMemberID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.ResetMemberID()
	crmu.mutation.SetMemberID(i)
	return crmu
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableMemberID(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetMemberID(*i)
	}
	return crmu
}

// AddMemberID adds i to the "member_id" field.
func (crmu *CourseRecordMemberUpdate) AddMemberID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.AddMemberID(i)
	return crmu
}

// ClearMemberID clears the value of the "member_id" field.
func (crmu *CourseRecordMemberUpdate) ClearMemberID() *CourseRecordMemberUpdate {
	crmu.mutation.ClearMemberID()
	return crmu
}

// SetCourseRecordScheduleID sets the "course_record_schedule_id" field.
func (crmu *CourseRecordMemberUpdate) SetCourseRecordScheduleID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.SetCourseRecordScheduleID(i)
	return crmu
}

// SetNillableCourseRecordScheduleID sets the "course_record_schedule_id" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableCourseRecordScheduleID(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetCourseRecordScheduleID(*i)
	}
	return crmu
}

// ClearCourseRecordScheduleID clears the value of the "course_record_schedule_id" field.
func (crmu *CourseRecordMemberUpdate) ClearCourseRecordScheduleID() *CourseRecordMemberUpdate {
	crmu.mutation.ClearCourseRecordScheduleID()
	return crmu
}

// SetType sets the "type" field.
func (crmu *CourseRecordMemberUpdate) SetType(s string) *CourseRecordMemberUpdate {
	crmu.mutation.SetType(s)
	return crmu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableType(s *string) *CourseRecordMemberUpdate {
	if s != nil {
		crmu.SetType(*s)
	}
	return crmu
}

// ClearType clears the value of the "type" field.
func (crmu *CourseRecordMemberUpdate) ClearType() *CourseRecordMemberUpdate {
	crmu.mutation.ClearType()
	return crmu
}

// SetStartTime sets the "start_time" field.
func (crmu *CourseRecordMemberUpdate) SetStartTime(t time.Time) *CourseRecordMemberUpdate {
	crmu.mutation.SetStartTime(t)
	return crmu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableStartTime(t *time.Time) *CourseRecordMemberUpdate {
	if t != nil {
		crmu.SetStartTime(*t)
	}
	return crmu
}

// ClearStartTime clears the value of the "start_time" field.
func (crmu *CourseRecordMemberUpdate) ClearStartTime() *CourseRecordMemberUpdate {
	crmu.mutation.ClearStartTime()
	return crmu
}

// SetEndTime sets the "end_time" field.
func (crmu *CourseRecordMemberUpdate) SetEndTime(t time.Time) *CourseRecordMemberUpdate {
	crmu.mutation.SetEndTime(t)
	return crmu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableEndTime(t *time.Time) *CourseRecordMemberUpdate {
	if t != nil {
		crmu.SetEndTime(*t)
	}
	return crmu
}

// ClearEndTime clears the value of the "end_time" field.
func (crmu *CourseRecordMemberUpdate) ClearEndTime() *CourseRecordMemberUpdate {
	crmu.mutation.ClearEndTime()
	return crmu
}

// SetSignStartTime sets the "sign_start_time" field.
func (crmu *CourseRecordMemberUpdate) SetSignStartTime(t time.Time) *CourseRecordMemberUpdate {
	crmu.mutation.SetSignStartTime(t)
	return crmu
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableSignStartTime(t *time.Time) *CourseRecordMemberUpdate {
	if t != nil {
		crmu.SetSignStartTime(*t)
	}
	return crmu
}

// ClearSignStartTime clears the value of the "sign_start_time" field.
func (crmu *CourseRecordMemberUpdate) ClearSignStartTime() *CourseRecordMemberUpdate {
	crmu.mutation.ClearSignStartTime()
	return crmu
}

// SetSignEndTime sets the "sign_end_time" field.
func (crmu *CourseRecordMemberUpdate) SetSignEndTime(t time.Time) *CourseRecordMemberUpdate {
	crmu.mutation.SetSignEndTime(t)
	return crmu
}

// SetNillableSignEndTime sets the "sign_end_time" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableSignEndTime(t *time.Time) *CourseRecordMemberUpdate {
	if t != nil {
		crmu.SetSignEndTime(*t)
	}
	return crmu
}

// ClearSignEndTime clears the value of the "sign_end_time" field.
func (crmu *CourseRecordMemberUpdate) ClearSignEndTime() *CourseRecordMemberUpdate {
	crmu.mutation.ClearSignEndTime()
	return crmu
}

// SetMemberProductID sets the "member_product_id" field.
func (crmu *CourseRecordMemberUpdate) SetMemberProductID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.ResetMemberProductID()
	crmu.mutation.SetMemberProductID(i)
	return crmu
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableMemberProductID(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetMemberProductID(*i)
	}
	return crmu
}

// AddMemberProductID adds i to the "member_product_id" field.
func (crmu *CourseRecordMemberUpdate) AddMemberProductID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.AddMemberProductID(i)
	return crmu
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (crmu *CourseRecordMemberUpdate) ClearMemberProductID() *CourseRecordMemberUpdate {
	crmu.mutation.ClearMemberProductID()
	return crmu
}

// SetMemberProductItemID sets the "member_product_item_id" field.
func (crmu *CourseRecordMemberUpdate) SetMemberProductItemID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.ResetMemberProductItemID()
	crmu.mutation.SetMemberProductItemID(i)
	return crmu
}

// SetNillableMemberProductItemID sets the "member_product_item_id" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableMemberProductItemID(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetMemberProductItemID(*i)
	}
	return crmu
}

// AddMemberProductItemID adds i to the "member_product_item_id" field.
func (crmu *CourseRecordMemberUpdate) AddMemberProductItemID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.AddMemberProductItemID(i)
	return crmu
}

// ClearMemberProductItemID clears the value of the "member_product_item_id" field.
func (crmu *CourseRecordMemberUpdate) ClearMemberProductItemID() *CourseRecordMemberUpdate {
	crmu.mutation.ClearMemberProductItemID()
	return crmu
}

// SetCoachID sets the "coach_id" field.
func (crmu *CourseRecordMemberUpdate) SetCoachID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.ResetCoachID()
	crmu.mutation.SetCoachID(i)
	return crmu
}

// SetNillableCoachID sets the "coach_id" field if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableCoachID(i *int64) *CourseRecordMemberUpdate {
	if i != nil {
		crmu.SetCoachID(*i)
	}
	return crmu
}

// AddCoachID adds i to the "coach_id" field.
func (crmu *CourseRecordMemberUpdate) AddCoachID(i int64) *CourseRecordMemberUpdate {
	crmu.mutation.AddCoachID(i)
	return crmu
}

// ClearCoachID clears the value of the "coach_id" field.
func (crmu *CourseRecordMemberUpdate) ClearCoachID() *CourseRecordMemberUpdate {
	crmu.mutation.ClearCoachID()
	return crmu
}

// SetScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID.
func (crmu *CourseRecordMemberUpdate) SetScheduleID(id int64) *CourseRecordMemberUpdate {
	crmu.mutation.SetScheduleID(id)
	return crmu
}

// SetNillableScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID if the given value is not nil.
func (crmu *CourseRecordMemberUpdate) SetNillableScheduleID(id *int64) *CourseRecordMemberUpdate {
	if id != nil {
		crmu = crmu.SetScheduleID(*id)
	}
	return crmu
}

// SetSchedule sets the "schedule" edge to the CourseRecordSchedule entity.
func (crmu *CourseRecordMemberUpdate) SetSchedule(c *CourseRecordSchedule) *CourseRecordMemberUpdate {
	return crmu.SetScheduleID(c.ID)
}

// Mutation returns the CourseRecordMemberMutation object of the builder.
func (crmu *CourseRecordMemberUpdate) Mutation() *CourseRecordMemberMutation {
	return crmu.mutation
}

// ClearSchedule clears the "schedule" edge to the CourseRecordSchedule entity.
func (crmu *CourseRecordMemberUpdate) ClearSchedule() *CourseRecordMemberUpdate {
	crmu.mutation.ClearSchedule()
	return crmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (crmu *CourseRecordMemberUpdate) Save(ctx context.Context) (int, error) {
	crmu.defaults()
	return withHooks(ctx, crmu.sqlSave, crmu.mutation, crmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (crmu *CourseRecordMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := crmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (crmu *CourseRecordMemberUpdate) Exec(ctx context.Context) error {
	_, err := crmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crmu *CourseRecordMemberUpdate) ExecX(ctx context.Context) {
	if err := crmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crmu *CourseRecordMemberUpdate) defaults() {
	if _, ok := crmu.mutation.UpdatedAt(); !ok {
		v := courserecordmember.UpdateDefaultUpdatedAt()
		crmu.mutation.SetUpdatedAt(v)
	}
}

func (crmu *CourseRecordMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(courserecordmember.Table, courserecordmember.Columns, sqlgraph.NewFieldSpec(courserecordmember.FieldID, field.TypeInt64))
	if ps := crmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := crmu.mutation.UpdatedAt(); ok {
		_spec.SetField(courserecordmember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := crmu.mutation.Status(); ok {
		_spec.SetField(courserecordmember.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := crmu.mutation.AddedStatus(); ok {
		_spec.AddField(courserecordmember.FieldStatus, field.TypeInt64, value)
	}
	if crmu.mutation.StatusCleared() {
		_spec.ClearField(courserecordmember.FieldStatus, field.TypeInt64)
	}
	if value, ok := crmu.mutation.VenueID(); ok {
		_spec.SetField(courserecordmember.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := crmu.mutation.AddedVenueID(); ok {
		_spec.AddField(courserecordmember.FieldVenueID, field.TypeInt64, value)
	}
	if crmu.mutation.VenueIDCleared() {
		_spec.ClearField(courserecordmember.FieldVenueID, field.TypeInt64)
	}
	if value, ok := crmu.mutation.MemberID(); ok {
		_spec.SetField(courserecordmember.FieldMemberID, field.TypeInt64, value)
	}
	if value, ok := crmu.mutation.AddedMemberID(); ok {
		_spec.AddField(courserecordmember.FieldMemberID, field.TypeInt64, value)
	}
	if crmu.mutation.MemberIDCleared() {
		_spec.ClearField(courserecordmember.FieldMemberID, field.TypeInt64)
	}
	if value, ok := crmu.mutation.GetType(); ok {
		_spec.SetField(courserecordmember.FieldType, field.TypeString, value)
	}
	if crmu.mutation.TypeCleared() {
		_spec.ClearField(courserecordmember.FieldType, field.TypeString)
	}
	if value, ok := crmu.mutation.StartTime(); ok {
		_spec.SetField(courserecordmember.FieldStartTime, field.TypeTime, value)
	}
	if crmu.mutation.StartTimeCleared() {
		_spec.ClearField(courserecordmember.FieldStartTime, field.TypeTime)
	}
	if value, ok := crmu.mutation.EndTime(); ok {
		_spec.SetField(courserecordmember.FieldEndTime, field.TypeTime, value)
	}
	if crmu.mutation.EndTimeCleared() {
		_spec.ClearField(courserecordmember.FieldEndTime, field.TypeTime)
	}
	if value, ok := crmu.mutation.SignStartTime(); ok {
		_spec.SetField(courserecordmember.FieldSignStartTime, field.TypeTime, value)
	}
	if crmu.mutation.SignStartTimeCleared() {
		_spec.ClearField(courserecordmember.FieldSignStartTime, field.TypeTime)
	}
	if value, ok := crmu.mutation.SignEndTime(); ok {
		_spec.SetField(courserecordmember.FieldSignEndTime, field.TypeTime, value)
	}
	if crmu.mutation.SignEndTimeCleared() {
		_spec.ClearField(courserecordmember.FieldSignEndTime, field.TypeTime)
	}
	if value, ok := crmu.mutation.MemberProductID(); ok {
		_spec.SetField(courserecordmember.FieldMemberProductID, field.TypeInt64, value)
	}
	if value, ok := crmu.mutation.AddedMemberProductID(); ok {
		_spec.AddField(courserecordmember.FieldMemberProductID, field.TypeInt64, value)
	}
	if crmu.mutation.MemberProductIDCleared() {
		_spec.ClearField(courserecordmember.FieldMemberProductID, field.TypeInt64)
	}
	if value, ok := crmu.mutation.MemberProductItemID(); ok {
		_spec.SetField(courserecordmember.FieldMemberProductItemID, field.TypeInt64, value)
	}
	if value, ok := crmu.mutation.AddedMemberProductItemID(); ok {
		_spec.AddField(courserecordmember.FieldMemberProductItemID, field.TypeInt64, value)
	}
	if crmu.mutation.MemberProductItemIDCleared() {
		_spec.ClearField(courserecordmember.FieldMemberProductItemID, field.TypeInt64)
	}
	if value, ok := crmu.mutation.CoachID(); ok {
		_spec.SetField(courserecordmember.FieldCoachID, field.TypeInt64, value)
	}
	if value, ok := crmu.mutation.AddedCoachID(); ok {
		_spec.AddField(courserecordmember.FieldCoachID, field.TypeInt64, value)
	}
	if crmu.mutation.CoachIDCleared() {
		_spec.ClearField(courserecordmember.FieldCoachID, field.TypeInt64)
	}
	if crmu.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordmember.ScheduleTable,
			Columns: []string{courserecordmember.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := crmu.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordmember.ScheduleTable,
			Columns: []string{courserecordmember.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, crmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courserecordmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	crmu.mutation.done = true
	return n, nil
}

// CourseRecordMemberUpdateOne is the builder for updating a single CourseRecordMember entity.
type CourseRecordMemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseRecordMemberMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (crmuo *CourseRecordMemberUpdateOne) SetUpdatedAt(t time.Time) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetUpdatedAt(t)
	return crmuo
}

// SetStatus sets the "status" field.
func (crmuo *CourseRecordMemberUpdateOne) SetStatus(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.ResetStatus()
	crmuo.mutation.SetStatus(i)
	return crmuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableStatus(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetStatus(*i)
	}
	return crmuo
}

// AddStatus adds i to the "status" field.
func (crmuo *CourseRecordMemberUpdateOne) AddStatus(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.AddStatus(i)
	return crmuo
}

// ClearStatus clears the value of the "status" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearStatus() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearStatus()
	return crmuo
}

// SetVenueID sets the "venue_id" field.
func (crmuo *CourseRecordMemberUpdateOne) SetVenueID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.ResetVenueID()
	crmuo.mutation.SetVenueID(i)
	return crmuo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableVenueID(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetVenueID(*i)
	}
	return crmuo
}

// AddVenueID adds i to the "venue_id" field.
func (crmuo *CourseRecordMemberUpdateOne) AddVenueID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.AddVenueID(i)
	return crmuo
}

// ClearVenueID clears the value of the "venue_id" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearVenueID() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearVenueID()
	return crmuo
}

// SetMemberID sets the "member_id" field.
func (crmuo *CourseRecordMemberUpdateOne) SetMemberID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.ResetMemberID()
	crmuo.mutation.SetMemberID(i)
	return crmuo
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableMemberID(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetMemberID(*i)
	}
	return crmuo
}

// AddMemberID adds i to the "member_id" field.
func (crmuo *CourseRecordMemberUpdateOne) AddMemberID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.AddMemberID(i)
	return crmuo
}

// ClearMemberID clears the value of the "member_id" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearMemberID() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearMemberID()
	return crmuo
}

// SetCourseRecordScheduleID sets the "course_record_schedule_id" field.
func (crmuo *CourseRecordMemberUpdateOne) SetCourseRecordScheduleID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetCourseRecordScheduleID(i)
	return crmuo
}

// SetNillableCourseRecordScheduleID sets the "course_record_schedule_id" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableCourseRecordScheduleID(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetCourseRecordScheduleID(*i)
	}
	return crmuo
}

// ClearCourseRecordScheduleID clears the value of the "course_record_schedule_id" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearCourseRecordScheduleID() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearCourseRecordScheduleID()
	return crmuo
}

// SetType sets the "type" field.
func (crmuo *CourseRecordMemberUpdateOne) SetType(s string) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetType(s)
	return crmuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableType(s *string) *CourseRecordMemberUpdateOne {
	if s != nil {
		crmuo.SetType(*s)
	}
	return crmuo
}

// ClearType clears the value of the "type" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearType() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearType()
	return crmuo
}

// SetStartTime sets the "start_time" field.
func (crmuo *CourseRecordMemberUpdateOne) SetStartTime(t time.Time) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetStartTime(t)
	return crmuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableStartTime(t *time.Time) *CourseRecordMemberUpdateOne {
	if t != nil {
		crmuo.SetStartTime(*t)
	}
	return crmuo
}

// ClearStartTime clears the value of the "start_time" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearStartTime() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearStartTime()
	return crmuo
}

// SetEndTime sets the "end_time" field.
func (crmuo *CourseRecordMemberUpdateOne) SetEndTime(t time.Time) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetEndTime(t)
	return crmuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableEndTime(t *time.Time) *CourseRecordMemberUpdateOne {
	if t != nil {
		crmuo.SetEndTime(*t)
	}
	return crmuo
}

// ClearEndTime clears the value of the "end_time" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearEndTime() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearEndTime()
	return crmuo
}

// SetSignStartTime sets the "sign_start_time" field.
func (crmuo *CourseRecordMemberUpdateOne) SetSignStartTime(t time.Time) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetSignStartTime(t)
	return crmuo
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableSignStartTime(t *time.Time) *CourseRecordMemberUpdateOne {
	if t != nil {
		crmuo.SetSignStartTime(*t)
	}
	return crmuo
}

// ClearSignStartTime clears the value of the "sign_start_time" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearSignStartTime() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearSignStartTime()
	return crmuo
}

// SetSignEndTime sets the "sign_end_time" field.
func (crmuo *CourseRecordMemberUpdateOne) SetSignEndTime(t time.Time) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetSignEndTime(t)
	return crmuo
}

// SetNillableSignEndTime sets the "sign_end_time" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableSignEndTime(t *time.Time) *CourseRecordMemberUpdateOne {
	if t != nil {
		crmuo.SetSignEndTime(*t)
	}
	return crmuo
}

// ClearSignEndTime clears the value of the "sign_end_time" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearSignEndTime() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearSignEndTime()
	return crmuo
}

// SetMemberProductID sets the "member_product_id" field.
func (crmuo *CourseRecordMemberUpdateOne) SetMemberProductID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.ResetMemberProductID()
	crmuo.mutation.SetMemberProductID(i)
	return crmuo
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableMemberProductID(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetMemberProductID(*i)
	}
	return crmuo
}

// AddMemberProductID adds i to the "member_product_id" field.
func (crmuo *CourseRecordMemberUpdateOne) AddMemberProductID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.AddMemberProductID(i)
	return crmuo
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearMemberProductID() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearMemberProductID()
	return crmuo
}

// SetMemberProductItemID sets the "member_product_item_id" field.
func (crmuo *CourseRecordMemberUpdateOne) SetMemberProductItemID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.ResetMemberProductItemID()
	crmuo.mutation.SetMemberProductItemID(i)
	return crmuo
}

// SetNillableMemberProductItemID sets the "member_product_item_id" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableMemberProductItemID(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetMemberProductItemID(*i)
	}
	return crmuo
}

// AddMemberProductItemID adds i to the "member_product_item_id" field.
func (crmuo *CourseRecordMemberUpdateOne) AddMemberProductItemID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.AddMemberProductItemID(i)
	return crmuo
}

// ClearMemberProductItemID clears the value of the "member_product_item_id" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearMemberProductItemID() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearMemberProductItemID()
	return crmuo
}

// SetCoachID sets the "coach_id" field.
func (crmuo *CourseRecordMemberUpdateOne) SetCoachID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.ResetCoachID()
	crmuo.mutation.SetCoachID(i)
	return crmuo
}

// SetNillableCoachID sets the "coach_id" field if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableCoachID(i *int64) *CourseRecordMemberUpdateOne {
	if i != nil {
		crmuo.SetCoachID(*i)
	}
	return crmuo
}

// AddCoachID adds i to the "coach_id" field.
func (crmuo *CourseRecordMemberUpdateOne) AddCoachID(i int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.AddCoachID(i)
	return crmuo
}

// ClearCoachID clears the value of the "coach_id" field.
func (crmuo *CourseRecordMemberUpdateOne) ClearCoachID() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearCoachID()
	return crmuo
}

// SetScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID.
func (crmuo *CourseRecordMemberUpdateOne) SetScheduleID(id int64) *CourseRecordMemberUpdateOne {
	crmuo.mutation.SetScheduleID(id)
	return crmuo
}

// SetNillableScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID if the given value is not nil.
func (crmuo *CourseRecordMemberUpdateOne) SetNillableScheduleID(id *int64) *CourseRecordMemberUpdateOne {
	if id != nil {
		crmuo = crmuo.SetScheduleID(*id)
	}
	return crmuo
}

// SetSchedule sets the "schedule" edge to the CourseRecordSchedule entity.
func (crmuo *CourseRecordMemberUpdateOne) SetSchedule(c *CourseRecordSchedule) *CourseRecordMemberUpdateOne {
	return crmuo.SetScheduleID(c.ID)
}

// Mutation returns the CourseRecordMemberMutation object of the builder.
func (crmuo *CourseRecordMemberUpdateOne) Mutation() *CourseRecordMemberMutation {
	return crmuo.mutation
}

// ClearSchedule clears the "schedule" edge to the CourseRecordSchedule entity.
func (crmuo *CourseRecordMemberUpdateOne) ClearSchedule() *CourseRecordMemberUpdateOne {
	crmuo.mutation.ClearSchedule()
	return crmuo
}

// Where appends a list predicates to the CourseRecordMemberUpdate builder.
func (crmuo *CourseRecordMemberUpdateOne) Where(ps ...predicate.CourseRecordMember) *CourseRecordMemberUpdateOne {
	crmuo.mutation.Where(ps...)
	return crmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (crmuo *CourseRecordMemberUpdateOne) Select(field string, fields ...string) *CourseRecordMemberUpdateOne {
	crmuo.fields = append([]string{field}, fields...)
	return crmuo
}

// Save executes the query and returns the updated CourseRecordMember entity.
func (crmuo *CourseRecordMemberUpdateOne) Save(ctx context.Context) (*CourseRecordMember, error) {
	crmuo.defaults()
	return withHooks(ctx, crmuo.sqlSave, crmuo.mutation, crmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (crmuo *CourseRecordMemberUpdateOne) SaveX(ctx context.Context) *CourseRecordMember {
	node, err := crmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (crmuo *CourseRecordMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := crmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crmuo *CourseRecordMemberUpdateOne) ExecX(ctx context.Context) {
	if err := crmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crmuo *CourseRecordMemberUpdateOne) defaults() {
	if _, ok := crmuo.mutation.UpdatedAt(); !ok {
		v := courserecordmember.UpdateDefaultUpdatedAt()
		crmuo.mutation.SetUpdatedAt(v)
	}
}

func (crmuo *CourseRecordMemberUpdateOne) sqlSave(ctx context.Context) (_node *CourseRecordMember, err error) {
	_spec := sqlgraph.NewUpdateSpec(courserecordmember.Table, courserecordmember.Columns, sqlgraph.NewFieldSpec(courserecordmember.FieldID, field.TypeInt64))
	id, ok := crmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CourseRecordMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := crmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, courserecordmember.FieldID)
		for _, f := range fields {
			if !courserecordmember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != courserecordmember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := crmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := crmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(courserecordmember.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := crmuo.mutation.Status(); ok {
		_spec.SetField(courserecordmember.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := crmuo.mutation.AddedStatus(); ok {
		_spec.AddField(courserecordmember.FieldStatus, field.TypeInt64, value)
	}
	if crmuo.mutation.StatusCleared() {
		_spec.ClearField(courserecordmember.FieldStatus, field.TypeInt64)
	}
	if value, ok := crmuo.mutation.VenueID(); ok {
		_spec.SetField(courserecordmember.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := crmuo.mutation.AddedVenueID(); ok {
		_spec.AddField(courserecordmember.FieldVenueID, field.TypeInt64, value)
	}
	if crmuo.mutation.VenueIDCleared() {
		_spec.ClearField(courserecordmember.FieldVenueID, field.TypeInt64)
	}
	if value, ok := crmuo.mutation.MemberID(); ok {
		_spec.SetField(courserecordmember.FieldMemberID, field.TypeInt64, value)
	}
	if value, ok := crmuo.mutation.AddedMemberID(); ok {
		_spec.AddField(courserecordmember.FieldMemberID, field.TypeInt64, value)
	}
	if crmuo.mutation.MemberIDCleared() {
		_spec.ClearField(courserecordmember.FieldMemberID, field.TypeInt64)
	}
	if value, ok := crmuo.mutation.GetType(); ok {
		_spec.SetField(courserecordmember.FieldType, field.TypeString, value)
	}
	if crmuo.mutation.TypeCleared() {
		_spec.ClearField(courserecordmember.FieldType, field.TypeString)
	}
	if value, ok := crmuo.mutation.StartTime(); ok {
		_spec.SetField(courserecordmember.FieldStartTime, field.TypeTime, value)
	}
	if crmuo.mutation.StartTimeCleared() {
		_spec.ClearField(courserecordmember.FieldStartTime, field.TypeTime)
	}
	if value, ok := crmuo.mutation.EndTime(); ok {
		_spec.SetField(courserecordmember.FieldEndTime, field.TypeTime, value)
	}
	if crmuo.mutation.EndTimeCleared() {
		_spec.ClearField(courserecordmember.FieldEndTime, field.TypeTime)
	}
	if value, ok := crmuo.mutation.SignStartTime(); ok {
		_spec.SetField(courserecordmember.FieldSignStartTime, field.TypeTime, value)
	}
	if crmuo.mutation.SignStartTimeCleared() {
		_spec.ClearField(courserecordmember.FieldSignStartTime, field.TypeTime)
	}
	if value, ok := crmuo.mutation.SignEndTime(); ok {
		_spec.SetField(courserecordmember.FieldSignEndTime, field.TypeTime, value)
	}
	if crmuo.mutation.SignEndTimeCleared() {
		_spec.ClearField(courserecordmember.FieldSignEndTime, field.TypeTime)
	}
	if value, ok := crmuo.mutation.MemberProductID(); ok {
		_spec.SetField(courserecordmember.FieldMemberProductID, field.TypeInt64, value)
	}
	if value, ok := crmuo.mutation.AddedMemberProductID(); ok {
		_spec.AddField(courserecordmember.FieldMemberProductID, field.TypeInt64, value)
	}
	if crmuo.mutation.MemberProductIDCleared() {
		_spec.ClearField(courserecordmember.FieldMemberProductID, field.TypeInt64)
	}
	if value, ok := crmuo.mutation.MemberProductItemID(); ok {
		_spec.SetField(courserecordmember.FieldMemberProductItemID, field.TypeInt64, value)
	}
	if value, ok := crmuo.mutation.AddedMemberProductItemID(); ok {
		_spec.AddField(courserecordmember.FieldMemberProductItemID, field.TypeInt64, value)
	}
	if crmuo.mutation.MemberProductItemIDCleared() {
		_spec.ClearField(courserecordmember.FieldMemberProductItemID, field.TypeInt64)
	}
	if value, ok := crmuo.mutation.CoachID(); ok {
		_spec.SetField(courserecordmember.FieldCoachID, field.TypeInt64, value)
	}
	if value, ok := crmuo.mutation.AddedCoachID(); ok {
		_spec.AddField(courserecordmember.FieldCoachID, field.TypeInt64, value)
	}
	if crmuo.mutation.CoachIDCleared() {
		_spec.ClearField(courserecordmember.FieldCoachID, field.TypeInt64)
	}
	if crmuo.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordmember.ScheduleTable,
			Columns: []string{courserecordmember.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := crmuo.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordmember.ScheduleTable,
			Columns: []string{courserecordmember.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CourseRecordMember{config: crmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, crmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courserecordmember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	crmuo.mutation.done = true
	return _node, nil
}
