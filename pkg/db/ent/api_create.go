// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/api"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// APICreate is the builder for creating a API entity.
type APICreate struct {
	config
	mutation *APIMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *APICreate) SetCreatedAt(t time.Time) *APICreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *APICreate) SetNillableCreatedAt(t *time.Time) *APICreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *APICreate) SetUpdatedAt(t time.Time) *APICreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *APICreate) SetNillableUpdatedAt(t *time.Time) *APICreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDelete sets the "delete" field.
func (ac *APICreate) SetDelete(i int64) *APICreate {
	ac.mutation.SetDelete(i)
	return ac
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (ac *APICreate) SetNillableDelete(i *int64) *APICreate {
	if i != nil {
		ac.SetDelete(*i)
	}
	return ac
}

// SetCreatedID sets the "created_id" field.
func (ac *APICreate) SetCreatedID(i int64) *APICreate {
	ac.mutation.SetCreatedID(i)
	return ac
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (ac *APICreate) SetNillableCreatedID(i *int64) *APICreate {
	if i != nil {
		ac.SetCreatedID(*i)
	}
	return ac
}

// SetPath sets the "path" field.
func (ac *APICreate) SetPath(s string) *APICreate {
	ac.mutation.SetPath(s)
	return ac
}

// SetTitle sets the "title" field.
func (ac *APICreate) SetTitle(s string) *APICreate {
	ac.mutation.SetTitle(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *APICreate) SetDescription(s string) *APICreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetAPIGroup sets the "api_group" field.
func (ac *APICreate) SetAPIGroup(s string) *APICreate {
	ac.mutation.SetAPIGroup(s)
	return ac
}

// SetMethod sets the "method" field.
func (ac *APICreate) SetMethod(s string) *APICreate {
	ac.mutation.SetMethod(s)
	return ac
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (ac *APICreate) SetNillableMethod(s *string) *APICreate {
	if s != nil {
		ac.SetMethod(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *APICreate) SetID(i int64) *APICreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the APIMutation object of the builder.
func (ac *APICreate) Mutation() *APIMutation {
	return ac.mutation
}

// Save creates the API in the database.
func (ac *APICreate) Save(ctx context.Context) (*API, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *APICreate) SaveX(ctx context.Context) *API {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *APICreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *APICreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *APICreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := api.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := api.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.Delete(); !ok {
		v := api.DefaultDelete
		ac.mutation.SetDelete(v)
	}
	if _, ok := ac.mutation.CreatedID(); !ok {
		v := api.DefaultCreatedID
		ac.mutation.SetCreatedID(v)
	}
	if _, ok := ac.mutation.Method(); !ok {
		v := api.DefaultMethod
		ac.mutation.SetMethod(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *APICreate) check() error {
	if _, ok := ac.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "API.path"`)}
	}
	if _, ok := ac.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "API.title"`)}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "API.description"`)}
	}
	if _, ok := ac.mutation.APIGroup(); !ok {
		return &ValidationError{Name: "api_group", err: errors.New(`ent: missing required field "API.api_group"`)}
	}
	if _, ok := ac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "API.method"`)}
	}
	return nil
}

func (ac *APICreate) sqlSave(ctx context.Context) (*API, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *APICreate) createSpec() (*API, *sqlgraph.CreateSpec) {
	var (
		_node = &API{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(api.Table, sqlgraph.NewFieldSpec(api.FieldID, field.TypeInt64))
	)
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(api.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(api.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Delete(); ok {
		_spec.SetField(api.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := ac.mutation.CreatedID(); ok {
		_spec.SetField(api.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := ac.mutation.Path(); ok {
		_spec.SetField(api.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := ac.mutation.Title(); ok {
		_spec.SetField(api.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(api.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.APIGroup(); ok {
		_spec.SetField(api.FieldAPIGroup, field.TypeString, value)
		_node.APIGroup = value
	}
	if value, ok := ac.mutation.Method(); ok {
		_spec.SetField(api.FieldMethod, field.TypeString, value)
		_node.Method = value
	}
	return _node, _spec
}

// APICreateBulk is the builder for creating many API entities in bulk.
type APICreateBulk struct {
	config
	err      error
	builders []*APICreate
}

// Save creates the API entities in the database.
func (acb *APICreateBulk) Save(ctx context.Context) ([]*API, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*API, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*APIMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *APICreateBulk) SaveX(ctx context.Context) []*API {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *APICreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *APICreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
