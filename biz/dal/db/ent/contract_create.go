// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/contract"
	"saas/biz/dal/db/ent/product"
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

// SetDelete sets the "delete" field.
func (cc *ContractCreate) SetDelete(i int64) *ContractCreate {
	cc.mutation.SetDelete(i)
	return cc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (cc *ContractCreate) SetNillableDelete(i *int64) *ContractCreate {
	if i != nil {
		cc.SetDelete(*i)
	}
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

// SetVenueID sets the "venue_id" field.
func (cc *ContractCreate) SetVenueID(i int64) *ContractCreate {
	cc.mutation.SetVenueID(i)
	return cc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (cc *ContractCreate) SetNillableVenueID(i *int64) *ContractCreate {
	if i != nil {
		cc.SetVenueID(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ContractCreate) SetID(i int64) *ContractCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (cc *ContractCreate) AddProductIDs(ids ...int64) *ContractCreate {
	cc.mutation.AddProductIDs(ids...)
	return cc
}

// AddProducts adds the "products" edges to the Product entity.
func (cc *ContractCreate) AddProducts(p ...*Product) *ContractCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProductIDs(ids...)
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
	if _, ok := cc.mutation.Delete(); !ok {
		v := contract.DefaultDelete
		cc.mutation.SetDelete(v)
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
	_spec.Schema = cc.schemaConfig.Contract
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
	if value, ok := cc.mutation.Delete(); ok {
		_spec.SetField(contract.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
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
	if value, ok := cc.mutation.VenueID(); ok {
		_spec.SetField(contract.FieldVenueID, field.TypeInt64, value)
		_node.VenueID = value
	}
	if nodes := cc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   contract.ProductsTable,
			Columns: contract.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = cc.schemaConfig.ProductContracts
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
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
