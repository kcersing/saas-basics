// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"member/biz/dal/mysql/ent/membernote"
	"member/biz/dal/mysql/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberNoteDelete is the builder for deleting a MemberNote entity.
type MemberNoteDelete struct {
	config
	hooks    []Hook
	mutation *MemberNoteMutation
}

// Where appends a list predicates to the MemberNoteDelete builder.
func (mnd *MemberNoteDelete) Where(ps ...predicate.MemberNote) *MemberNoteDelete {
	mnd.mutation.Where(ps...)
	return mnd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mnd *MemberNoteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mnd.sqlExec, mnd.mutation, mnd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mnd *MemberNoteDelete) ExecX(ctx context.Context) int {
	n, err := mnd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mnd *MemberNoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(membernote.Table, sqlgraph.NewFieldSpec(membernote.FieldID, field.TypeInt64))
	if ps := mnd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mnd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mnd.mutation.done = true
	return affected, err
}

// MemberNoteDeleteOne is the builder for deleting a single MemberNote entity.
type MemberNoteDeleteOne struct {
	mnd *MemberNoteDelete
}

// Where appends a list predicates to the MemberNoteDelete builder.
func (mndo *MemberNoteDeleteOne) Where(ps ...predicate.MemberNote) *MemberNoteDeleteOne {
	mndo.mnd.mutation.Where(ps...)
	return mndo
}

// Exec executes the deletion query.
func (mndo *MemberNoteDeleteOne) Exec(ctx context.Context) error {
	n, err := mndo.mnd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{membernote.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mndo *MemberNoteDeleteOne) ExecX(ctx context.Context) {
	if err := mndo.Exec(ctx); err != nil {
		panic(err)
	}
}
