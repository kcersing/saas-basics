// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"saas/biz/dal/db/ent/internal"
	"saas/biz/dal/db/ent/membertoken"
	"saas/biz/dal/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberTokenDelete is the builder for deleting a MemberToken entity.
type MemberTokenDelete struct {
	config
	hooks    []Hook
	mutation *MemberTokenMutation
}

// Where appends a list predicates to the MemberTokenDelete builder.
func (mtd *MemberTokenDelete) Where(ps ...predicate.MemberToken) *MemberTokenDelete {
	mtd.mutation.Where(ps...)
	return mtd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mtd *MemberTokenDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, mtd.sqlExec, mtd.mutation, mtd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (mtd *MemberTokenDelete) ExecX(ctx context.Context) int {
	n, err := mtd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mtd *MemberTokenDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(membertoken.Table, sqlgraph.NewFieldSpec(membertoken.FieldID, field.TypeInt64))
	_spec.Node.Schema = mtd.schemaConfig.MemberToken
	ctx = internal.NewSchemaConfigContext(ctx, mtd.schemaConfig)
	if ps := mtd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mtd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	mtd.mutation.done = true
	return affected, err
}

// MemberTokenDeleteOne is the builder for deleting a single MemberToken entity.
type MemberTokenDeleteOne struct {
	mtd *MemberTokenDelete
}

// Where appends a list predicates to the MemberTokenDelete builder.
func (mtdo *MemberTokenDeleteOne) Where(ps ...predicate.MemberToken) *MemberTokenDeleteOne {
	mtdo.mtd.mutation.Where(ps...)
	return mtdo
}

// Exec executes the deletion query.
func (mtdo *MemberTokenDeleteOne) Exec(ctx context.Context) error {
	n, err := mtdo.mtd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{membertoken.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mtdo *MemberTokenDeleteOne) ExecX(ctx context.Context) {
	if err := mtdo.Exec(ctx); err != nil {
		panic(err)
	}
}
