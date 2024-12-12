// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/schedulecoach"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleCoachDelete is the builder for deleting a ScheduleCoach entity.
type ScheduleCoachDelete struct {
	config
	hooks    []Hook
	mutation *ScheduleCoachMutation
}

// Where appends a list predicates to the ScheduleCoachDelete builder.
func (scd *ScheduleCoachDelete) Where(ps ...predicate.ScheduleCoach) *ScheduleCoachDelete {
	scd.mutation.Where(ps...)
	return scd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (scd *ScheduleCoachDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, scd.sqlExec, scd.mutation, scd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (scd *ScheduleCoachDelete) ExecX(ctx context.Context) int {
	n, err := scd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (scd *ScheduleCoachDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(schedulecoach.Table, sqlgraph.NewFieldSpec(schedulecoach.FieldID, field.TypeInt64))
	if ps := scd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, scd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	scd.mutation.done = true
	return affected, err
}

// ScheduleCoachDeleteOne is the builder for deleting a single ScheduleCoach entity.
type ScheduleCoachDeleteOne struct {
	scd *ScheduleCoachDelete
}

// Where appends a list predicates to the ScheduleCoachDelete builder.
func (scdo *ScheduleCoachDeleteOne) Where(ps ...predicate.ScheduleCoach) *ScheduleCoachDeleteOne {
	scdo.scd.mutation.Where(ps...)
	return scdo
}

// Exec executes the deletion query.
func (scdo *ScheduleCoachDeleteOne) Exec(ctx context.Context) error {
	n, err := scdo.scd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{schedulecoach.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scdo *ScheduleCoachDeleteOne) ExecX(ctx context.Context) {
	if err := scdo.Exec(ctx); err != nil {
		panic(err)
	}
}
