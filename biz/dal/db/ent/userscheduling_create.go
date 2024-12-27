// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/userscheduling"
	"saas/idl_gen/model/base"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserSchedulingCreate is the builder for creating a UserScheduling entity.
type UserSchedulingCreate struct {
	config
	mutation *UserSchedulingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (usc *UserSchedulingCreate) SetCreatedAt(t time.Time) *UserSchedulingCreate {
	usc.mutation.SetCreatedAt(t)
	return usc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableCreatedAt(t *time.Time) *UserSchedulingCreate {
	if t != nil {
		usc.SetCreatedAt(*t)
	}
	return usc
}

// SetUpdatedAt sets the "updated_at" field.
func (usc *UserSchedulingCreate) SetUpdatedAt(t time.Time) *UserSchedulingCreate {
	usc.mutation.SetUpdatedAt(t)
	return usc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableUpdatedAt(t *time.Time) *UserSchedulingCreate {
	if t != nil {
		usc.SetUpdatedAt(*t)
	}
	return usc
}

// SetDelete sets the "delete" field.
func (usc *UserSchedulingCreate) SetDelete(i int64) *UserSchedulingCreate {
	usc.mutation.SetDelete(i)
	return usc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableDelete(i *int64) *UserSchedulingCreate {
	if i != nil {
		usc.SetDelete(*i)
	}
	return usc
}

// SetCreatedID sets the "created_id" field.
func (usc *UserSchedulingCreate) SetCreatedID(i int64) *UserSchedulingCreate {
	usc.mutation.SetCreatedID(i)
	return usc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableCreatedID(i *int64) *UserSchedulingCreate {
	if i != nil {
		usc.SetCreatedID(*i)
	}
	return usc
}

// SetStatus sets the "status" field.
func (usc *UserSchedulingCreate) SetStatus(i int64) *UserSchedulingCreate {
	usc.mutation.SetStatus(i)
	return usc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableStatus(i *int64) *UserSchedulingCreate {
	if i != nil {
		usc.SetStatus(*i)
	}
	return usc
}

// SetDate sets the "date" field.
func (usc *UserSchedulingCreate) SetDate(t time.Time) *UserSchedulingCreate {
	usc.mutation.SetDate(t)
	return usc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableDate(t *time.Time) *UserSchedulingCreate {
	if t != nil {
		usc.SetDate(*t)
	}
	return usc
}

// SetPeriod sets the "period" field.
func (usc *UserSchedulingCreate) SetPeriod(bsd base.UserSchedulingDate) *UserSchedulingCreate {
	usc.mutation.SetPeriod(bsd)
	return usc
}

// SetNillablePeriod sets the "period" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillablePeriod(bsd *base.UserSchedulingDate) *UserSchedulingCreate {
	if bsd != nil {
		usc.SetPeriod(*bsd)
	}
	return usc
}

// SetUserID sets the "user_id" field.
func (usc *UserSchedulingCreate) SetUserID(i int64) *UserSchedulingCreate {
	usc.mutation.SetUserID(i)
	return usc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableUserID(i *int64) *UserSchedulingCreate {
	if i != nil {
		usc.SetUserID(*i)
	}
	return usc
}

// SetID sets the "id" field.
func (usc *UserSchedulingCreate) SetID(i int64) *UserSchedulingCreate {
	usc.mutation.SetID(i)
	return usc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (usc *UserSchedulingCreate) SetUsersID(id int64) *UserSchedulingCreate {
	usc.mutation.SetUsersID(id)
	return usc
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (usc *UserSchedulingCreate) SetNillableUsersID(id *int64) *UserSchedulingCreate {
	if id != nil {
		usc = usc.SetUsersID(*id)
	}
	return usc
}

// SetUsers sets the "users" edge to the User entity.
func (usc *UserSchedulingCreate) SetUsers(u *User) *UserSchedulingCreate {
	return usc.SetUsersID(u.ID)
}

// Mutation returns the UserSchedulingMutation object of the builder.
func (usc *UserSchedulingCreate) Mutation() *UserSchedulingMutation {
	return usc.mutation
}

// Save creates the UserScheduling in the database.
func (usc *UserSchedulingCreate) Save(ctx context.Context) (*UserScheduling, error) {
	usc.defaults()
	return withHooks(ctx, usc.sqlSave, usc.mutation, usc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (usc *UserSchedulingCreate) SaveX(ctx context.Context) *UserScheduling {
	v, err := usc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (usc *UserSchedulingCreate) Exec(ctx context.Context) error {
	_, err := usc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (usc *UserSchedulingCreate) ExecX(ctx context.Context) {
	if err := usc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (usc *UserSchedulingCreate) defaults() {
	if _, ok := usc.mutation.CreatedAt(); !ok {
		v := userscheduling.DefaultCreatedAt()
		usc.mutation.SetCreatedAt(v)
	}
	if _, ok := usc.mutation.UpdatedAt(); !ok {
		v := userscheduling.DefaultUpdatedAt()
		usc.mutation.SetUpdatedAt(v)
	}
	if _, ok := usc.mutation.Delete(); !ok {
		v := userscheduling.DefaultDelete
		usc.mutation.SetDelete(v)
	}
	if _, ok := usc.mutation.CreatedID(); !ok {
		v := userscheduling.DefaultCreatedID
		usc.mutation.SetCreatedID(v)
	}
	if _, ok := usc.mutation.Status(); !ok {
		v := userscheduling.DefaultStatus
		usc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (usc *UserSchedulingCreate) check() error {
	return nil
}

func (usc *UserSchedulingCreate) sqlSave(ctx context.Context) (*UserScheduling, error) {
	if err := usc.check(); err != nil {
		return nil, err
	}
	_node, _spec := usc.createSpec()
	if err := sqlgraph.CreateNode(ctx, usc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	usc.mutation.id = &_node.ID
	usc.mutation.done = true
	return _node, nil
}

func (usc *UserSchedulingCreate) createSpec() (*UserScheduling, *sqlgraph.CreateSpec) {
	var (
		_node = &UserScheduling{config: usc.config}
		_spec = sqlgraph.NewCreateSpec(userscheduling.Table, sqlgraph.NewFieldSpec(userscheduling.FieldID, field.TypeInt64))
	)
	if id, ok := usc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := usc.mutation.CreatedAt(); ok {
		_spec.SetField(userscheduling.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := usc.mutation.UpdatedAt(); ok {
		_spec.SetField(userscheduling.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := usc.mutation.Delete(); ok {
		_spec.SetField(userscheduling.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := usc.mutation.CreatedID(); ok {
		_spec.SetField(userscheduling.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := usc.mutation.Status(); ok {
		_spec.SetField(userscheduling.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := usc.mutation.Date(); ok {
		_spec.SetField(userscheduling.FieldDate, field.TypeTime, value)
		_node.Date = value
	}
	if value, ok := usc.mutation.Period(); ok {
		_spec.SetField(userscheduling.FieldPeriod, field.TypeJSON, value)
		_node.Period = value
	}
	if nodes := usc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userscheduling.UsersTable,
			Columns: []string{userscheduling.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserSchedulingCreateBulk is the builder for creating many UserScheduling entities in bulk.
type UserSchedulingCreateBulk struct {
	config
	err      error
	builders []*UserSchedulingCreate
}

// Save creates the UserScheduling entities in the database.
func (uscb *UserSchedulingCreateBulk) Save(ctx context.Context) ([]*UserScheduling, error) {
	if uscb.err != nil {
		return nil, uscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uscb.builders))
	nodes := make([]*UserScheduling, len(uscb.builders))
	mutators := make([]Mutator, len(uscb.builders))
	for i := range uscb.builders {
		func(i int, root context.Context) {
			builder := uscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserSchedulingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uscb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uscb *UserSchedulingCreateBulk) SaveX(ctx context.Context) []*UserScheduling {
	v, err := uscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uscb *UserSchedulingCreateBulk) Exec(ctx context.Context) error {
	_, err := uscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uscb *UserSchedulingCreateBulk) ExecX(ctx context.Context) {
	if err := uscb.Exec(ctx); err != nil {
		panic(err)
	}
}
