// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/biz/dal/db/ent/dictionary"
	"saas/biz/dal/db/ent/dictionarydetail"
	"saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictionaryDetailCreate is the builder for creating a DictionaryDetail entity.
type DictionaryDetailCreate struct {
	config
	mutation *DictionaryDetailMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ddc *DictionaryDetailCreate) SetCreatedAt(t time.Time) *DictionaryDetailCreate {
	ddc.mutation.SetCreatedAt(t)
	return ddc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableCreatedAt(t *time.Time) *DictionaryDetailCreate {
	if t != nil {
		ddc.SetCreatedAt(*t)
	}
	return ddc
}

// SetUpdatedAt sets the "updated_at" field.
func (ddc *DictionaryDetailCreate) SetUpdatedAt(t time.Time) *DictionaryDetailCreate {
	ddc.mutation.SetUpdatedAt(t)
	return ddc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableUpdatedAt(t *time.Time) *DictionaryDetailCreate {
	if t != nil {
		ddc.SetUpdatedAt(*t)
	}
	return ddc
}

// SetDelete sets the "delete" field.
func (ddc *DictionaryDetailCreate) SetDelete(i int64) *DictionaryDetailCreate {
	ddc.mutation.SetDelete(i)
	return ddc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableDelete(i *int64) *DictionaryDetailCreate {
	if i != nil {
		ddc.SetDelete(*i)
	}
	return ddc
}

// SetCreatedID sets the "created_id" field.
func (ddc *DictionaryDetailCreate) SetCreatedID(i int64) *DictionaryDetailCreate {
	ddc.mutation.SetCreatedID(i)
	return ddc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableCreatedID(i *int64) *DictionaryDetailCreate {
	if i != nil {
		ddc.SetCreatedID(*i)
	}
	return ddc
}

// SetStatus sets the "status" field.
func (ddc *DictionaryDetailCreate) SetStatus(i int64) *DictionaryDetailCreate {
	ddc.mutation.SetStatus(i)
	return ddc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableStatus(i *int64) *DictionaryDetailCreate {
	if i != nil {
		ddc.SetStatus(*i)
	}
	return ddc
}

// SetTitle sets the "title" field.
func (ddc *DictionaryDetailCreate) SetTitle(s string) *DictionaryDetailCreate {
	ddc.mutation.SetTitle(s)
	return ddc
}

// SetKey sets the "key" field.
func (ddc *DictionaryDetailCreate) SetKey(s string) *DictionaryDetailCreate {
	ddc.mutation.SetKey(s)
	return ddc
}

// SetValue sets the "value" field.
func (ddc *DictionaryDetailCreate) SetValue(s string) *DictionaryDetailCreate {
	ddc.mutation.SetValue(s)
	return ddc
}

// SetDictionaryID sets the "dictionary_id" field.
func (ddc *DictionaryDetailCreate) SetDictionaryID(i int64) *DictionaryDetailCreate {
	ddc.mutation.SetDictionaryID(i)
	return ddc
}

// SetNillableDictionaryID sets the "dictionary_id" field if the given value is not nil.
func (ddc *DictionaryDetailCreate) SetNillableDictionaryID(i *int64) *DictionaryDetailCreate {
	if i != nil {
		ddc.SetDictionaryID(*i)
	}
	return ddc
}

// SetID sets the "id" field.
func (ddc *DictionaryDetailCreate) SetID(i int64) *DictionaryDetailCreate {
	ddc.mutation.SetID(i)
	return ddc
}

// SetDictionary sets the "dictionary" edge to the Dictionary entity.
func (ddc *DictionaryDetailCreate) SetDictionary(d *Dictionary) *DictionaryDetailCreate {
	return ddc.SetDictionaryID(d.ID)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (ddc *DictionaryDetailCreate) AddUserIDs(ids ...int64) *DictionaryDetailCreate {
	ddc.mutation.AddUserIDs(ids...)
	return ddc
}

// AddUsers adds the "users" edges to the User entity.
func (ddc *DictionaryDetailCreate) AddUsers(u ...*User) *DictionaryDetailCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ddc.AddUserIDs(ids...)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (ddc *DictionaryDetailCreate) AddProductIDs(ids ...int64) *DictionaryDetailCreate {
	ddc.mutation.AddProductIDs(ids...)
	return ddc
}

// AddProducts adds the "products" edges to the Product entity.
func (ddc *DictionaryDetailCreate) AddProducts(p ...*Product) *DictionaryDetailCreate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ddc.AddProductIDs(ids...)
}

// Mutation returns the DictionaryDetailMutation object of the builder.
func (ddc *DictionaryDetailCreate) Mutation() *DictionaryDetailMutation {
	return ddc.mutation
}

// Save creates the DictionaryDetail in the database.
func (ddc *DictionaryDetailCreate) Save(ctx context.Context) (*DictionaryDetail, error) {
	ddc.defaults()
	return withHooks(ctx, ddc.sqlSave, ddc.mutation, ddc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ddc *DictionaryDetailCreate) SaveX(ctx context.Context) *DictionaryDetail {
	v, err := ddc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddc *DictionaryDetailCreate) Exec(ctx context.Context) error {
	_, err := ddc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddc *DictionaryDetailCreate) ExecX(ctx context.Context) {
	if err := ddc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddc *DictionaryDetailCreate) defaults() {
	if _, ok := ddc.mutation.CreatedAt(); !ok {
		v := dictionarydetail.DefaultCreatedAt()
		ddc.mutation.SetCreatedAt(v)
	}
	if _, ok := ddc.mutation.UpdatedAt(); !ok {
		v := dictionarydetail.DefaultUpdatedAt()
		ddc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ddc.mutation.Delete(); !ok {
		v := dictionarydetail.DefaultDelete
		ddc.mutation.SetDelete(v)
	}
	if _, ok := ddc.mutation.CreatedID(); !ok {
		v := dictionarydetail.DefaultCreatedID
		ddc.mutation.SetCreatedID(v)
	}
	if _, ok := ddc.mutation.Status(); !ok {
		v := dictionarydetail.DefaultStatus
		ddc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddc *DictionaryDetailCreate) check() error {
	if _, ok := ddc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "DictionaryDetail.title"`)}
	}
	if _, ok := ddc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "DictionaryDetail.key"`)}
	}
	if _, ok := ddc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "DictionaryDetail.value"`)}
	}
	return nil
}

func (ddc *DictionaryDetailCreate) sqlSave(ctx context.Context) (*DictionaryDetail, error) {
	if err := ddc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ddc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ddc.mutation.id = &_node.ID
	ddc.mutation.done = true
	return _node, nil
}

func (ddc *DictionaryDetailCreate) createSpec() (*DictionaryDetail, *sqlgraph.CreateSpec) {
	var (
		_node = &DictionaryDetail{config: ddc.config}
		_spec = sqlgraph.NewCreateSpec(dictionarydetail.Table, sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeInt64))
	)
	_spec.Schema = ddc.schemaConfig.DictionaryDetail
	if id, ok := ddc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ddc.mutation.CreatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ddc.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionarydetail.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ddc.mutation.Delete(); ok {
		_spec.SetField(dictionarydetail.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := ddc.mutation.CreatedID(); ok {
		_spec.SetField(dictionarydetail.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := ddc.mutation.Status(); ok {
		_spec.SetField(dictionarydetail.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := ddc.mutation.Title(); ok {
		_spec.SetField(dictionarydetail.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := ddc.mutation.Key(); ok {
		_spec.SetField(dictionarydetail.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := ddc.mutation.Value(); ok {
		_spec.SetField(dictionarydetail.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := ddc.mutation.DictionaryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   dictionarydetail.DictionaryTable,
			Columns: []string{dictionarydetail.DictionaryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = ddc.schemaConfig.DictionaryDetail
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DictionaryID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ddc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dictionarydetail.UsersTable,
			Columns: dictionarydetail.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = ddc.schemaConfig.UserTags
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ddc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   dictionarydetail.ProductsTable,
			Columns: dictionarydetail.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = ddc.schemaConfig.ProductTags
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DictionaryDetailCreateBulk is the builder for creating many DictionaryDetail entities in bulk.
type DictionaryDetailCreateBulk struct {
	config
	err      error
	builders []*DictionaryDetailCreate
}

// Save creates the DictionaryDetail entities in the database.
func (ddcb *DictionaryDetailCreateBulk) Save(ctx context.Context) ([]*DictionaryDetail, error) {
	if ddcb.err != nil {
		return nil, ddcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ddcb.builders))
	nodes := make([]*DictionaryDetail, len(ddcb.builders))
	mutators := make([]Mutator, len(ddcb.builders))
	for i := range ddcb.builders {
		func(i int, root context.Context) {
			builder := ddcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DictionaryDetailMutation)
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
					_, err = mutators[i+1].Mutate(root, ddcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ddcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddcb *DictionaryDetailCreateBulk) SaveX(ctx context.Context) []*DictionaryDetail {
	v, err := ddcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ddcb *DictionaryDetailCreateBulk) Exec(ctx context.Context) error {
	_, err := ddcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddcb *DictionaryDetailCreateBulk) ExecX(ctx context.Context) {
	if err := ddcb.Exec(ctx); err != nil {
		panic(err)
	}
}
