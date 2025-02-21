// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/internal"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntryLogsDelete is the builder for deleting a EntryLogs entity.
type EntryLogsDelete struct {
	config
	hooks    []Hook
	mutation *EntryLogsMutation
}

// Where appends a list predicates to the EntryLogsDelete builder.
func (eld *EntryLogsDelete) Where(ps ...predicate.EntryLogs) *EntryLogsDelete {
	eld.mutation.Where(ps...)
	return eld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eld *EntryLogsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, eld.sqlExec, eld.mutation, eld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (eld *EntryLogsDelete) ExecX(ctx context.Context) int {
	n, err := eld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eld *EntryLogsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(entrylogs.Table, sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64))
	_spec.Node.Schema = eld.schemaConfig.EntryLogs
	ctx = internal.NewSchemaConfigContext(ctx, eld.schemaConfig)
	if ps := eld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	eld.mutation.done = true
	return affected, err
}

// EntryLogsDeleteOne is the builder for deleting a single EntryLogs entity.
type EntryLogsDeleteOne struct {
	eld *EntryLogsDelete
}

// Where appends a list predicates to the EntryLogsDelete builder.
func (eldo *EntryLogsDeleteOne) Where(ps ...predicate.EntryLogs) *EntryLogsDeleteOne {
	eldo.eld.mutation.Where(ps...)
	return eldo
}

// Exec executes the deletion query.
func (eldo *EntryLogsDeleteOne) Exec(ctx context.Context) error {
	n, err := eldo.eld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{entrylogs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eldo *EntryLogsDeleteOne) ExecX(ctx context.Context) {
	if err := eldo.Exec(ctx); err != nil {
		panic(err)
	}
}
