// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/courserecordcoach"
	"saas/pkg/db/ent/courserecordschedule"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseRecordCoachUpdate is the builder for updating CourseRecordCoach entities.
type CourseRecordCoachUpdate struct {
	config
	hooks    []Hook
	mutation *CourseRecordCoachMutation
}

// Where appends a list predicates to the CourseRecordCoachUpdate builder.
func (crcu *CourseRecordCoachUpdate) Where(ps ...predicate.CourseRecordCoach) *CourseRecordCoachUpdate {
	crcu.mutation.Where(ps...)
	return crcu
}

// SetUpdatedAt sets the "updated_at" field.
func (crcu *CourseRecordCoachUpdate) SetUpdatedAt(t time.Time) *CourseRecordCoachUpdate {
	crcu.mutation.SetUpdatedAt(t)
	return crcu
}

// SetVenueID sets the "venue_id" field.
func (crcu *CourseRecordCoachUpdate) SetVenueID(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.ResetVenueID()
	crcu.mutation.SetVenueID(i)
	return crcu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableVenueID(i *int64) *CourseRecordCoachUpdate {
	if i != nil {
		crcu.SetVenueID(*i)
	}
	return crcu
}

// AddVenueID adds i to the "venue_id" field.
func (crcu *CourseRecordCoachUpdate) AddVenueID(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.AddVenueID(i)
	return crcu
}

// ClearVenueID clears the value of the "venue_id" field.
func (crcu *CourseRecordCoachUpdate) ClearVenueID() *CourseRecordCoachUpdate {
	crcu.mutation.ClearVenueID()
	return crcu
}

// SetCoachID sets the "coach_id" field.
func (crcu *CourseRecordCoachUpdate) SetCoachID(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.ResetCoachID()
	crcu.mutation.SetCoachID(i)
	return crcu
}

// SetNillableCoachID sets the "coach_id" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableCoachID(i *int64) *CourseRecordCoachUpdate {
	if i != nil {
		crcu.SetCoachID(*i)
	}
	return crcu
}

// AddCoachID adds i to the "coach_id" field.
func (crcu *CourseRecordCoachUpdate) AddCoachID(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.AddCoachID(i)
	return crcu
}

// ClearCoachID clears the value of the "coach_id" field.
func (crcu *CourseRecordCoachUpdate) ClearCoachID() *CourseRecordCoachUpdate {
	crcu.mutation.ClearCoachID()
	return crcu
}

// SetCourseRecordScheduleID sets the "course_record_schedule_id" field.
func (crcu *CourseRecordCoachUpdate) SetCourseRecordScheduleID(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.SetCourseRecordScheduleID(i)
	return crcu
}

// SetNillableCourseRecordScheduleID sets the "course_record_schedule_id" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableCourseRecordScheduleID(i *int64) *CourseRecordCoachUpdate {
	if i != nil {
		crcu.SetCourseRecordScheduleID(*i)
	}
	return crcu
}

// ClearCourseRecordScheduleID clears the value of the "course_record_schedule_id" field.
func (crcu *CourseRecordCoachUpdate) ClearCourseRecordScheduleID() *CourseRecordCoachUpdate {
	crcu.mutation.ClearCourseRecordScheduleID()
	return crcu
}

// SetType sets the "type" field.
func (crcu *CourseRecordCoachUpdate) SetType(s string) *CourseRecordCoachUpdate {
	crcu.mutation.SetType(s)
	return crcu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableType(s *string) *CourseRecordCoachUpdate {
	if s != nil {
		crcu.SetType(*s)
	}
	return crcu
}

// ClearType clears the value of the "type" field.
func (crcu *CourseRecordCoachUpdate) ClearType() *CourseRecordCoachUpdate {
	crcu.mutation.ClearType()
	return crcu
}

// SetStartTime sets the "start_time" field.
func (crcu *CourseRecordCoachUpdate) SetStartTime(t time.Time) *CourseRecordCoachUpdate {
	crcu.mutation.SetStartTime(t)
	return crcu
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableStartTime(t *time.Time) *CourseRecordCoachUpdate {
	if t != nil {
		crcu.SetStartTime(*t)
	}
	return crcu
}

// ClearStartTime clears the value of the "start_time" field.
func (crcu *CourseRecordCoachUpdate) ClearStartTime() *CourseRecordCoachUpdate {
	crcu.mutation.ClearStartTime()
	return crcu
}

// SetEndTime sets the "end_time" field.
func (crcu *CourseRecordCoachUpdate) SetEndTime(t time.Time) *CourseRecordCoachUpdate {
	crcu.mutation.SetEndTime(t)
	return crcu
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableEndTime(t *time.Time) *CourseRecordCoachUpdate {
	if t != nil {
		crcu.SetEndTime(*t)
	}
	return crcu
}

// ClearEndTime clears the value of the "end_time" field.
func (crcu *CourseRecordCoachUpdate) ClearEndTime() *CourseRecordCoachUpdate {
	crcu.mutation.ClearEndTime()
	return crcu
}

// SetSignStartTime sets the "sign_start_time" field.
func (crcu *CourseRecordCoachUpdate) SetSignStartTime(t time.Time) *CourseRecordCoachUpdate {
	crcu.mutation.SetSignStartTime(t)
	return crcu
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableSignStartTime(t *time.Time) *CourseRecordCoachUpdate {
	if t != nil {
		crcu.SetSignStartTime(*t)
	}
	return crcu
}

// ClearSignStartTime clears the value of the "sign_start_time" field.
func (crcu *CourseRecordCoachUpdate) ClearSignStartTime() *CourseRecordCoachUpdate {
	crcu.mutation.ClearSignStartTime()
	return crcu
}

// SetSignNdTime sets the "sign_nd_time" field.
func (crcu *CourseRecordCoachUpdate) SetSignNdTime(t time.Time) *CourseRecordCoachUpdate {
	crcu.mutation.SetSignNdTime(t)
	return crcu
}

// SetNillableSignNdTime sets the "sign_nd_time" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableSignNdTime(t *time.Time) *CourseRecordCoachUpdate {
	if t != nil {
		crcu.SetSignNdTime(*t)
	}
	return crcu
}

// ClearSignNdTime clears the value of the "sign_nd_time" field.
func (crcu *CourseRecordCoachUpdate) ClearSignNdTime() *CourseRecordCoachUpdate {
	crcu.mutation.ClearSignNdTime()
	return crcu
}

// SetStatus sets the "status" field.
func (crcu *CourseRecordCoachUpdate) SetStatus(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.ResetStatus()
	crcu.mutation.SetStatus(i)
	return crcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableStatus(i *int64) *CourseRecordCoachUpdate {
	if i != nil {
		crcu.SetStatus(*i)
	}
	return crcu
}

// AddStatus adds i to the "status" field.
func (crcu *CourseRecordCoachUpdate) AddStatus(i int64) *CourseRecordCoachUpdate {
	crcu.mutation.AddStatus(i)
	return crcu
}

// ClearStatus clears the value of the "status" field.
func (crcu *CourseRecordCoachUpdate) ClearStatus() *CourseRecordCoachUpdate {
	crcu.mutation.ClearStatus()
	return crcu
}

// SetScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID.
func (crcu *CourseRecordCoachUpdate) SetScheduleID(id int64) *CourseRecordCoachUpdate {
	crcu.mutation.SetScheduleID(id)
	return crcu
}

// SetNillableScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID if the given value is not nil.
func (crcu *CourseRecordCoachUpdate) SetNillableScheduleID(id *int64) *CourseRecordCoachUpdate {
	if id != nil {
		crcu = crcu.SetScheduleID(*id)
	}
	return crcu
}

// SetSchedule sets the "schedule" edge to the CourseRecordSchedule entity.
func (crcu *CourseRecordCoachUpdate) SetSchedule(c *CourseRecordSchedule) *CourseRecordCoachUpdate {
	return crcu.SetScheduleID(c.ID)
}

// Mutation returns the CourseRecordCoachMutation object of the builder.
func (crcu *CourseRecordCoachUpdate) Mutation() *CourseRecordCoachMutation {
	return crcu.mutation
}

// ClearSchedule clears the "schedule" edge to the CourseRecordSchedule entity.
func (crcu *CourseRecordCoachUpdate) ClearSchedule() *CourseRecordCoachUpdate {
	crcu.mutation.ClearSchedule()
	return crcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (crcu *CourseRecordCoachUpdate) Save(ctx context.Context) (int, error) {
	crcu.defaults()
	return withHooks(ctx, crcu.sqlSave, crcu.mutation, crcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (crcu *CourseRecordCoachUpdate) SaveX(ctx context.Context) int {
	affected, err := crcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (crcu *CourseRecordCoachUpdate) Exec(ctx context.Context) error {
	_, err := crcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcu *CourseRecordCoachUpdate) ExecX(ctx context.Context) {
	if err := crcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crcu *CourseRecordCoachUpdate) defaults() {
	if _, ok := crcu.mutation.UpdatedAt(); !ok {
		v := courserecordcoach.UpdateDefaultUpdatedAt()
		crcu.mutation.SetUpdatedAt(v)
	}
}

func (crcu *CourseRecordCoachUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(courserecordcoach.Table, courserecordcoach.Columns, sqlgraph.NewFieldSpec(courserecordcoach.FieldID, field.TypeInt64))
	if ps := crcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := crcu.mutation.UpdatedAt(); ok {
		_spec.SetField(courserecordcoach.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := crcu.mutation.VenueID(); ok {
		_spec.SetField(courserecordcoach.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := crcu.mutation.AddedVenueID(); ok {
		_spec.AddField(courserecordcoach.FieldVenueID, field.TypeInt64, value)
	}
	if crcu.mutation.VenueIDCleared() {
		_spec.ClearField(courserecordcoach.FieldVenueID, field.TypeInt64)
	}
	if value, ok := crcu.mutation.CoachID(); ok {
		_spec.SetField(courserecordcoach.FieldCoachID, field.TypeInt64, value)
	}
	if value, ok := crcu.mutation.AddedCoachID(); ok {
		_spec.AddField(courserecordcoach.FieldCoachID, field.TypeInt64, value)
	}
	if crcu.mutation.CoachIDCleared() {
		_spec.ClearField(courserecordcoach.FieldCoachID, field.TypeInt64)
	}
	if value, ok := crcu.mutation.GetType(); ok {
		_spec.SetField(courserecordcoach.FieldType, field.TypeString, value)
	}
	if crcu.mutation.TypeCleared() {
		_spec.ClearField(courserecordcoach.FieldType, field.TypeString)
	}
	if value, ok := crcu.mutation.StartTime(); ok {
		_spec.SetField(courserecordcoach.FieldStartTime, field.TypeTime, value)
	}
	if crcu.mutation.StartTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldStartTime, field.TypeTime)
	}
	if value, ok := crcu.mutation.EndTime(); ok {
		_spec.SetField(courserecordcoach.FieldEndTime, field.TypeTime, value)
	}
	if crcu.mutation.EndTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldEndTime, field.TypeTime)
	}
	if value, ok := crcu.mutation.SignStartTime(); ok {
		_spec.SetField(courserecordcoach.FieldSignStartTime, field.TypeTime, value)
	}
	if crcu.mutation.SignStartTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldSignStartTime, field.TypeTime)
	}
	if value, ok := crcu.mutation.SignNdTime(); ok {
		_spec.SetField(courserecordcoach.FieldSignNdTime, field.TypeTime, value)
	}
	if crcu.mutation.SignNdTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldSignNdTime, field.TypeTime)
	}
	if value, ok := crcu.mutation.Status(); ok {
		_spec.SetField(courserecordcoach.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := crcu.mutation.AddedStatus(); ok {
		_spec.AddField(courserecordcoach.FieldStatus, field.TypeInt64, value)
	}
	if crcu.mutation.StatusCleared() {
		_spec.ClearField(courserecordcoach.FieldStatus, field.TypeInt64)
	}
	if crcu.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordcoach.ScheduleTable,
			Columns: []string{courserecordcoach.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := crcu.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordcoach.ScheduleTable,
			Columns: []string{courserecordcoach.ScheduleColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, crcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courserecordcoach.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	crcu.mutation.done = true
	return n, nil
}

// CourseRecordCoachUpdateOne is the builder for updating a single CourseRecordCoach entity.
type CourseRecordCoachUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseRecordCoachMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (crcuo *CourseRecordCoachUpdateOne) SetUpdatedAt(t time.Time) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetUpdatedAt(t)
	return crcuo
}

// SetVenueID sets the "venue_id" field.
func (crcuo *CourseRecordCoachUpdateOne) SetVenueID(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.ResetVenueID()
	crcuo.mutation.SetVenueID(i)
	return crcuo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableVenueID(i *int64) *CourseRecordCoachUpdateOne {
	if i != nil {
		crcuo.SetVenueID(*i)
	}
	return crcuo
}

// AddVenueID adds i to the "venue_id" field.
func (crcuo *CourseRecordCoachUpdateOne) AddVenueID(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.AddVenueID(i)
	return crcuo
}

// ClearVenueID clears the value of the "venue_id" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearVenueID() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearVenueID()
	return crcuo
}

// SetCoachID sets the "coach_id" field.
func (crcuo *CourseRecordCoachUpdateOne) SetCoachID(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.ResetCoachID()
	crcuo.mutation.SetCoachID(i)
	return crcuo
}

// SetNillableCoachID sets the "coach_id" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableCoachID(i *int64) *CourseRecordCoachUpdateOne {
	if i != nil {
		crcuo.SetCoachID(*i)
	}
	return crcuo
}

// AddCoachID adds i to the "coach_id" field.
func (crcuo *CourseRecordCoachUpdateOne) AddCoachID(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.AddCoachID(i)
	return crcuo
}

// ClearCoachID clears the value of the "coach_id" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearCoachID() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearCoachID()
	return crcuo
}

// SetCourseRecordScheduleID sets the "course_record_schedule_id" field.
func (crcuo *CourseRecordCoachUpdateOne) SetCourseRecordScheduleID(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetCourseRecordScheduleID(i)
	return crcuo
}

// SetNillableCourseRecordScheduleID sets the "course_record_schedule_id" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableCourseRecordScheduleID(i *int64) *CourseRecordCoachUpdateOne {
	if i != nil {
		crcuo.SetCourseRecordScheduleID(*i)
	}
	return crcuo
}

// ClearCourseRecordScheduleID clears the value of the "course_record_schedule_id" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearCourseRecordScheduleID() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearCourseRecordScheduleID()
	return crcuo
}

// SetType sets the "type" field.
func (crcuo *CourseRecordCoachUpdateOne) SetType(s string) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetType(s)
	return crcuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableType(s *string) *CourseRecordCoachUpdateOne {
	if s != nil {
		crcuo.SetType(*s)
	}
	return crcuo
}

// ClearType clears the value of the "type" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearType() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearType()
	return crcuo
}

// SetStartTime sets the "start_time" field.
func (crcuo *CourseRecordCoachUpdateOne) SetStartTime(t time.Time) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetStartTime(t)
	return crcuo
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableStartTime(t *time.Time) *CourseRecordCoachUpdateOne {
	if t != nil {
		crcuo.SetStartTime(*t)
	}
	return crcuo
}

// ClearStartTime clears the value of the "start_time" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearStartTime() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearStartTime()
	return crcuo
}

// SetEndTime sets the "end_time" field.
func (crcuo *CourseRecordCoachUpdateOne) SetEndTime(t time.Time) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetEndTime(t)
	return crcuo
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableEndTime(t *time.Time) *CourseRecordCoachUpdateOne {
	if t != nil {
		crcuo.SetEndTime(*t)
	}
	return crcuo
}

// ClearEndTime clears the value of the "end_time" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearEndTime() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearEndTime()
	return crcuo
}

// SetSignStartTime sets the "sign_start_time" field.
func (crcuo *CourseRecordCoachUpdateOne) SetSignStartTime(t time.Time) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetSignStartTime(t)
	return crcuo
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableSignStartTime(t *time.Time) *CourseRecordCoachUpdateOne {
	if t != nil {
		crcuo.SetSignStartTime(*t)
	}
	return crcuo
}

// ClearSignStartTime clears the value of the "sign_start_time" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearSignStartTime() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearSignStartTime()
	return crcuo
}

// SetSignNdTime sets the "sign_nd_time" field.
func (crcuo *CourseRecordCoachUpdateOne) SetSignNdTime(t time.Time) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetSignNdTime(t)
	return crcuo
}

// SetNillableSignNdTime sets the "sign_nd_time" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableSignNdTime(t *time.Time) *CourseRecordCoachUpdateOne {
	if t != nil {
		crcuo.SetSignNdTime(*t)
	}
	return crcuo
}

// ClearSignNdTime clears the value of the "sign_nd_time" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearSignNdTime() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearSignNdTime()
	return crcuo
}

// SetStatus sets the "status" field.
func (crcuo *CourseRecordCoachUpdateOne) SetStatus(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.ResetStatus()
	crcuo.mutation.SetStatus(i)
	return crcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableStatus(i *int64) *CourseRecordCoachUpdateOne {
	if i != nil {
		crcuo.SetStatus(*i)
	}
	return crcuo
}

// AddStatus adds i to the "status" field.
func (crcuo *CourseRecordCoachUpdateOne) AddStatus(i int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.AddStatus(i)
	return crcuo
}

// ClearStatus clears the value of the "status" field.
func (crcuo *CourseRecordCoachUpdateOne) ClearStatus() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearStatus()
	return crcuo
}

// SetScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID.
func (crcuo *CourseRecordCoachUpdateOne) SetScheduleID(id int64) *CourseRecordCoachUpdateOne {
	crcuo.mutation.SetScheduleID(id)
	return crcuo
}

// SetNillableScheduleID sets the "schedule" edge to the CourseRecordSchedule entity by ID if the given value is not nil.
func (crcuo *CourseRecordCoachUpdateOne) SetNillableScheduleID(id *int64) *CourseRecordCoachUpdateOne {
	if id != nil {
		crcuo = crcuo.SetScheduleID(*id)
	}
	return crcuo
}

// SetSchedule sets the "schedule" edge to the CourseRecordSchedule entity.
func (crcuo *CourseRecordCoachUpdateOne) SetSchedule(c *CourseRecordSchedule) *CourseRecordCoachUpdateOne {
	return crcuo.SetScheduleID(c.ID)
}

// Mutation returns the CourseRecordCoachMutation object of the builder.
func (crcuo *CourseRecordCoachUpdateOne) Mutation() *CourseRecordCoachMutation {
	return crcuo.mutation
}

// ClearSchedule clears the "schedule" edge to the CourseRecordSchedule entity.
func (crcuo *CourseRecordCoachUpdateOne) ClearSchedule() *CourseRecordCoachUpdateOne {
	crcuo.mutation.ClearSchedule()
	return crcuo
}

// Where appends a list predicates to the CourseRecordCoachUpdate builder.
func (crcuo *CourseRecordCoachUpdateOne) Where(ps ...predicate.CourseRecordCoach) *CourseRecordCoachUpdateOne {
	crcuo.mutation.Where(ps...)
	return crcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (crcuo *CourseRecordCoachUpdateOne) Select(field string, fields ...string) *CourseRecordCoachUpdateOne {
	crcuo.fields = append([]string{field}, fields...)
	return crcuo
}

// Save executes the query and returns the updated CourseRecordCoach entity.
func (crcuo *CourseRecordCoachUpdateOne) Save(ctx context.Context) (*CourseRecordCoach, error) {
	crcuo.defaults()
	return withHooks(ctx, crcuo.sqlSave, crcuo.mutation, crcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (crcuo *CourseRecordCoachUpdateOne) SaveX(ctx context.Context) *CourseRecordCoach {
	node, err := crcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (crcuo *CourseRecordCoachUpdateOne) Exec(ctx context.Context) error {
	_, err := crcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcuo *CourseRecordCoachUpdateOne) ExecX(ctx context.Context) {
	if err := crcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crcuo *CourseRecordCoachUpdateOne) defaults() {
	if _, ok := crcuo.mutation.UpdatedAt(); !ok {
		v := courserecordcoach.UpdateDefaultUpdatedAt()
		crcuo.mutation.SetUpdatedAt(v)
	}
}

func (crcuo *CourseRecordCoachUpdateOne) sqlSave(ctx context.Context) (_node *CourseRecordCoach, err error) {
	_spec := sqlgraph.NewUpdateSpec(courserecordcoach.Table, courserecordcoach.Columns, sqlgraph.NewFieldSpec(courserecordcoach.FieldID, field.TypeInt64))
	id, ok := crcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CourseRecordCoach.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := crcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, courserecordcoach.FieldID)
		for _, f := range fields {
			if !courserecordcoach.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != courserecordcoach.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := crcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := crcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(courserecordcoach.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := crcuo.mutation.VenueID(); ok {
		_spec.SetField(courserecordcoach.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := crcuo.mutation.AddedVenueID(); ok {
		_spec.AddField(courserecordcoach.FieldVenueID, field.TypeInt64, value)
	}
	if crcuo.mutation.VenueIDCleared() {
		_spec.ClearField(courserecordcoach.FieldVenueID, field.TypeInt64)
	}
	if value, ok := crcuo.mutation.CoachID(); ok {
		_spec.SetField(courserecordcoach.FieldCoachID, field.TypeInt64, value)
	}
	if value, ok := crcuo.mutation.AddedCoachID(); ok {
		_spec.AddField(courserecordcoach.FieldCoachID, field.TypeInt64, value)
	}
	if crcuo.mutation.CoachIDCleared() {
		_spec.ClearField(courserecordcoach.FieldCoachID, field.TypeInt64)
	}
	if value, ok := crcuo.mutation.GetType(); ok {
		_spec.SetField(courserecordcoach.FieldType, field.TypeString, value)
	}
	if crcuo.mutation.TypeCleared() {
		_spec.ClearField(courserecordcoach.FieldType, field.TypeString)
	}
	if value, ok := crcuo.mutation.StartTime(); ok {
		_spec.SetField(courserecordcoach.FieldStartTime, field.TypeTime, value)
	}
	if crcuo.mutation.StartTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldStartTime, field.TypeTime)
	}
	if value, ok := crcuo.mutation.EndTime(); ok {
		_spec.SetField(courserecordcoach.FieldEndTime, field.TypeTime, value)
	}
	if crcuo.mutation.EndTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldEndTime, field.TypeTime)
	}
	if value, ok := crcuo.mutation.SignStartTime(); ok {
		_spec.SetField(courserecordcoach.FieldSignStartTime, field.TypeTime, value)
	}
	if crcuo.mutation.SignStartTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldSignStartTime, field.TypeTime)
	}
	if value, ok := crcuo.mutation.SignNdTime(); ok {
		_spec.SetField(courserecordcoach.FieldSignNdTime, field.TypeTime, value)
	}
	if crcuo.mutation.SignNdTimeCleared() {
		_spec.ClearField(courserecordcoach.FieldSignNdTime, field.TypeTime)
	}
	if value, ok := crcuo.mutation.Status(); ok {
		_spec.SetField(courserecordcoach.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := crcuo.mutation.AddedStatus(); ok {
		_spec.AddField(courserecordcoach.FieldStatus, field.TypeInt64, value)
	}
	if crcuo.mutation.StatusCleared() {
		_spec.ClearField(courserecordcoach.FieldStatus, field.TypeInt64)
	}
	if crcuo.mutation.ScheduleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordcoach.ScheduleTable,
			Columns: []string{courserecordcoach.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := crcuo.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   courserecordcoach.ScheduleTable,
			Columns: []string{courserecordcoach.ScheduleColumn},
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
	_node = &CourseRecordCoach{config: crcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, crcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{courserecordcoach.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	crcuo.mutation.done = true
	return _node, nil
}
