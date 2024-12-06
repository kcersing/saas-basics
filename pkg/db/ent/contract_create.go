// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/contract"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContractCreate is the builder for creating a Contract entity.
type ContractCreate struct {
	config
	mutation *ContractMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContractCreate) SetCreatedAt(t time.Time) *ContractCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableCreatedAt(t *time.Time) *ContractCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContractCreate) SetUpdatedAt(t time.Time) *ContractCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContractCreate) SetNillableUpdatedAt(t *time.Time) *ContractCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDeleteAt sets the "delete_at" field.
func (cc *ContractCreate) SetDeleteAt(t time.Time) *ContractCreate {
	cc.mutation.SetDeleteAt(t)
	return cc
}

// SetCreatedID sets the "created_id" field.
func (cc *ContractCreate) SetCreatedID(i int64) *ContractCreate {
	cc.mutation.SetCreatedID(i)
	return cc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (cc *ContractCreate) SetNillableCreatedID(i *int64) *ContractCreate {
	if i != nil {
		cc.SetCreatedID(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *ContractCreate) SetStatus(i int64) *ContractCreate {
	cc.mutation.SetStatus(i)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *ContractCreate) SetNillableStatus(i *int64) *ContractCreate {
	if i != nil {
		cc.SetStatus(*i)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ContractCreate) SetName(s string) *ContractCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *ContractCreate) SetNillableName(s *string) *ContractCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *ContractCreate) SetContent(s string) *ContractCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cc *ContractCreate) SetNillableContent(s *string) *ContractCreate {
	if s != nil {
		cc.SetContent(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ContractCreate) SetID(i int64) *ContractCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the ContractMutation object of the builder.
func (cc *ContractCreate) Mutation() *ContractMutation {
	return cc.mutation
}

// Save creates the Contract in the database.
func (cc *ContractCreate) Save(ctx context.Context) (*Contract, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContractCreate) SaveX(ctx context.Context) *Contract {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContractCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContractCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContractCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := contract.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := contract.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.CreatedID(); !ok {
		v := contract.DefaultCreatedID
		cc.mutation.SetCreatedID(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := contract.DefaultStatus
		cc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContractCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Contract.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Contract.updated_at"`)}
	}
	if _, ok := cc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "Contract.delete_at"`)}
	}
	if _, ok := cc.mutation.CreatedID(); !ok {
		return &ValidationError{Name: "created_id", err: errors.New(`ent: missing required field "Contract.created_id"`)}
	}
	return nil
}

func (cc *ContractCreate) sqlSave(ctx context.Context) (*Contract, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ContractCreate) createSpec() (*Contract, *sqlgraph.CreateSpec) {
	var (
		_node = &Contract{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(contract.Table, sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt64))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(contract.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(contract.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeleteAt(); ok {
		_spec.SetField(contract.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := cc.mutation.CreatedID(); ok {
		_spec.SetField(contract.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(contract.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(contract.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.SetField(contract.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	return _node, _spec
}

// ContractCreateBulk is the builder for creating many Contract entities in bulk.
type ContractCreateBulk struct {
	config
	err      error
	builders []*ContractCreate
}

// Save creates the Contract entities in the database.
func (ccb *ContractCreateBulk) Save(ctx context.Context) ([]*Contract, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contract, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContractMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ContractCreateBulk) SaveX(ctx context.Context) []*Contract {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContractCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContractCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
