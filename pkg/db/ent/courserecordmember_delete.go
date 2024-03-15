// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/pkg/db/ent/courserecordmember"
	"saas/pkg/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseRecordMemberDelete is the builder for deleting a CourseRecordMember entity.
type CourseRecordMemberDelete struct {
	config
	hooks    []Hook
	mutation *CourseRecordMemberMutation
}

// Where appends a list predicates to the CourseRecordMemberDelete builder.
func (crmd *CourseRecordMemberDelete) Where(ps ...predicate.CourseRecordMember) *CourseRecordMemberDelete {
	crmd.mutation.Where(ps...)
	return crmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (crmd *CourseRecordMemberDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, crmd.sqlExec, crmd.mutation, crmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (crmd *CourseRecordMemberDelete) ExecX(ctx context.Context) int {
	n, err := crmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (crmd *CourseRecordMemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(courserecordmember.Table, sqlgraph.NewFieldSpec(courserecordmember.FieldID, field.TypeInt64))
	if ps := crmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, crmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	crmd.mutation.done = true
	return affected, err
}

// CourseRecordMemberDeleteOne is the builder for deleting a single CourseRecordMember entity.
type CourseRecordMemberDeleteOne struct {
	crmd *CourseRecordMemberDelete
}

// Where appends a list predicates to the CourseRecordMemberDelete builder.
func (crmdo *CourseRecordMemberDeleteOne) Where(ps ...predicate.CourseRecordMember) *CourseRecordMemberDeleteOne {
	crmdo.crmd.mutation.Where(ps...)
	return crmdo
}

// Exec executes the deletion query.
func (crmdo *CourseRecordMemberDeleteOne) Exec(ctx context.Context) error {
	n, err := crmdo.crmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{courserecordmember.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (crmdo *CourseRecordMemberDeleteOne) ExecX(ctx context.Context) {
	if err := crmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
