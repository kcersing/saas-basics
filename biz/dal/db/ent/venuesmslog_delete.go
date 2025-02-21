// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/biz/dal/db/ent/internal"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/venuesmslog"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VenueSmsLogDelete is the builder for deleting a VenueSmsLog entity.
type VenueSmsLogDelete struct {
	config
	hooks    []Hook
	mutation *VenueSmsLogMutation
}

// Where appends a list predicates to the VenueSmsLogDelete builder.
func (vsld *VenueSmsLogDelete) Where(ps ...predicate.VenueSmsLog) *VenueSmsLogDelete {
	vsld.mutation.Where(ps...)
	return vsld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vsld *VenueSmsLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vsld.sqlExec, vsld.mutation, vsld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vsld *VenueSmsLogDelete) ExecX(ctx context.Context) int {
	n, err := vsld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vsld *VenueSmsLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(venuesmslog.Table, sqlgraph.NewFieldSpec(venuesmslog.FieldID, field.TypeInt64))
	_spec.Node.Schema = vsld.schemaConfig.VenueSmsLog
	ctx = internal.NewSchemaConfigContext(ctx, vsld.schemaConfig)
	if ps := vsld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vsld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vsld.mutation.done = true
	return affected, err
}

// VenueSmsLogDeleteOne is the builder for deleting a single VenueSmsLog entity.
type VenueSmsLogDeleteOne struct {
	vsld *VenueSmsLogDelete
}

// Where appends a list predicates to the VenueSmsLogDelete builder.
func (vsldo *VenueSmsLogDeleteOne) Where(ps ...predicate.VenueSmsLog) *VenueSmsLogDeleteOne {
	vsldo.vsld.mutation.Where(ps...)
	return vsldo
}

// Exec executes the deletion query.
func (vsldo *VenueSmsLogDeleteOne) Exec(ctx context.Context) error {
	n, err := vsldo.vsld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{venuesmslog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vsldo *VenueSmsLogDeleteOne) ExecX(ctx context.Context) {
	if err := vsldo.Exec(ctx); err != nil {
		panic(err)
	}
}
