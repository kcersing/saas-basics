// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/pkg/db/ent/courserecordschedule"
	"saas/pkg/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseRecordScheduleDelete is the builder for deleting a CourseRecordSchedule entity.
type CourseRecordScheduleDelete struct {
	config
	hooks    []Hook
	mutation *CourseRecordScheduleMutation
}

// Where appends a list predicates to the CourseRecordScheduleDelete builder.
func (crsd *CourseRecordScheduleDelete) Where(ps ...predicate.CourseRecordSchedule) *CourseRecordScheduleDelete {
	crsd.mutation.Where(ps...)
	return crsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crsd *CourseRecordScheduleDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crsd.sqlExec, crsd.mutation, crsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crsd *CourseRecordScheduleDelete) ExecX(ctx context.Context) int {
	n, err := crsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crsd *CourseRecordScheduleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(courserecordschedule.Table, sqlgraph.NewFieldSpec(courserecordschedule.FieldID, field.TypeInt64))
	if ps := crsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crsd.mutation.done = true
	return affected, err
}

// CourseRecordScheduleDeleteOne is the builder for deleting a single CourseRecordSchedule entity.
type CourseRecordScheduleDeleteOne struct {
	crsd *CourseRecordScheduleDelete
}

// Where appends a list predicates to the CourseRecordScheduleDelete builder.
func (crsdo *CourseRecordScheduleDeleteOne) Where(ps ...predicate.CourseRecordSchedule) *CourseRecordScheduleDeleteOne {
	crsdo.crsd.mutation.Where(ps...)
	return crsdo
}

// Exec executes the deletion query.
func (crsdo *CourseRecordScheduleDeleteOne) Exec(ctx context.Context) error {
	n, err := crsdo.crsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{courserecordschedule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crsdo *CourseRecordScheduleDeleteOne) ExecX(ctx context.Context) {
	if err := crsdo.Exec(ctx); err != nil {
		panic(err)
	}
}
