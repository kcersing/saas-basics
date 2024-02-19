// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/menuparam"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuParamCreate is the builder for creating a MenuParam entity.
type MenuParamCreate struct {
	config
	mutation *MenuParamMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mpc *MenuParamCreate) SetCreatedAt(t time.Time) *MenuParamCreate {
	mpc.mutation.SetCreatedAt(t)
	return mpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mpc *MenuParamCreate) SetNillableCreatedAt(t *time.Time) *MenuParamCreate {
	if t != nil {
		mpc.SetCreatedAt(*t)
	}
	return mpc
}

// SetUpdatedAt sets the "updated_at" field.
func (mpc *MenuParamCreate) SetUpdatedAt(t time.Time) *MenuParamCreate {
	mpc.mutation.SetUpdatedAt(t)
	return mpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mpc *MenuParamCreate) SetNillableUpdatedAt(t *time.Time) *MenuParamCreate {
	if t != nil {
		mpc.SetUpdatedAt(*t)
	}
	return mpc
}

// SetType sets the "type" field.
func (mpc *MenuParamCreate) SetType(s string) *MenuParamCreate {
	mpc.mutation.SetType(s)
	return mpc
}

// SetKey sets the "key" field.
func (mpc *MenuParamCreate) SetKey(s string) *MenuParamCreate {
	mpc.mutation.SetKey(s)
	return mpc
}

// SetValue sets the "value" field.
func (mpc *MenuParamCreate) SetValue(s string) *MenuParamCreate {
	mpc.mutation.SetValue(s)
	return mpc
}

// SetID sets the "id" field.
func (mpc *MenuParamCreate) SetID(u uint64) *MenuParamCreate {
	mpc.mutation.SetID(u)
	return mpc
}

// SetMenusID sets the "menus" edge to the Menu entity by ID.
func (mpc *MenuParamCreate) SetMenusID(id uint64) *MenuParamCreate {
	mpc.mutation.SetMenusID(id)
	return mpc
}

// SetNillableMenusID sets the "menus" edge to the Menu entity by ID if the given value is not nil.
func (mpc *MenuParamCreate) SetNillableMenusID(id *uint64) *MenuParamCreate {
	if id != nil {
		mpc = mpc.SetMenusID(*id)
	}
	return mpc
}

// SetMenus sets the "menus" edge to the Menu entity.
func (mpc *MenuParamCreate) SetMenus(m *Menu) *MenuParamCreate {
	return mpc.SetMenusID(m.ID)
}

// Mutation returns the MenuParamMutation object of the builder.
func (mpc *MenuParamCreate) Mutation() *MenuParamMutation {
	return mpc.mutation
}

// Save creates the MenuParam in the database.
func (mpc *MenuParamCreate) Save(ctx context.Context) (*MenuParam, error) {
	mpc.defaults()
	return withHooks(ctx, mpc.sqlSave, mpc.mutation, mpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mpc *MenuParamCreate) SaveX(ctx context.Context) *MenuParam {
	v, err := mpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpc *MenuParamCreate) Exec(ctx context.Context) error {
	_, err := mpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpc *MenuParamCreate) ExecX(ctx context.Context) {
	if err := mpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpc *MenuParamCreate) defaults() {
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		v := menuparam.DefaultCreatedAt()
		mpc.mutation.SetCreatedAt(v)
	}
	if _, ok := mpc.mutation.UpdatedAt(); !ok {
		v := menuparam.DefaultUpdatedAt()
		mpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpc *MenuParamCreate) check() error {
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MenuParam.created_at"`)}
	}
	if _, ok := mpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MenuParam.updated_at"`)}
	}
	if _, ok := mpc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "MenuParam.type"`)}
	}
	if _, ok := mpc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "MenuParam.key"`)}
	}
	if _, ok := mpc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "MenuParam.value"`)}
	}
	return nil
}

func (mpc *MenuParamCreate) sqlSave(ctx context.Context) (*MenuParam, error) {
	if err := mpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	mpc.mutation.id = &_node.ID
	mpc.mutation.done = true
	return _node, nil
}

func (mpc *MenuParamCreate) createSpec() (*MenuParam, *sqlgraph.CreateSpec) {
	var (
		_node = &MenuParam{config: mpc.config}
		_spec = sqlgraph.NewCreateSpec(menuparam.Table, sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeUint64))
	)
	if id, ok := mpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mpc.mutation.CreatedAt(); ok {
		_spec.SetField(menuparam.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mpc.mutation.UpdatedAt(); ok {
		_spec.SetField(menuparam.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mpc.mutation.GetType(); ok {
		_spec.SetField(menuparam.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mpc.mutation.Key(); ok {
		_spec.SetField(menuparam.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := mpc.mutation.Value(); ok {
		_spec.SetField(menuparam.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := mpc.mutation.MenusIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menuparam.MenusTable,
			Columns: []string{menuparam.MenusColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.menu_params = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MenuParamCreateBulk is the builder for creating many MenuParam entities in bulk.
type MenuParamCreateBulk struct {
	config
	err      error
	builders []*MenuParamCreate
}

// Save creates the MenuParam entities in the database.
func (mpcb *MenuParamCreateBulk) Save(ctx context.Context) ([]*MenuParam, error) {
	if mpcb.err != nil {
		return nil, mpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mpcb.builders))
	nodes := make([]*MenuParam, len(mpcb.builders))
	mutators := make([]Mutator, len(mpcb.builders))
	for i := range mpcb.builders {
		func(i int, root context.Context) {
			builder := mpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuParamMutation)
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
					_, err = mutators[i+1].Mutate(root, mpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mpcb.driver, spec); err != nil {
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
					nodes[i].ID = uint64(id)
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
		if _, err := mutators[0].Mutate(ctx, mpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mpcb *MenuParamCreateBulk) SaveX(ctx context.Context) []*MenuParam {
	v, err := mpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpcb *MenuParamCreateBulk) Exec(ctx context.Context) error {
	_, err := mpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpcb *MenuParamCreateBulk) ExecX(ctx context.Context) {
	if err := mpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
