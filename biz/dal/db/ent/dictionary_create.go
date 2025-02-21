// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/biz/dal/db/ent/dictionary"
	"saas/biz/dal/db/ent/dictionarydetail"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DictionaryCreate is the builder for creating a Dictionary entity.
type DictionaryCreate struct {
	config
	mutation *DictionaryMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (dc *DictionaryCreate) SetCreatedAt(t time.Time) *DictionaryCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableCreatedAt(t *time.Time) *DictionaryCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// SetUpdatedAt sets the "updated_at" field.
func (dc *DictionaryCreate) SetUpdatedAt(t time.Time) *DictionaryCreate {
	dc.mutation.SetUpdatedAt(t)
	return dc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableUpdatedAt(t *time.Time) *DictionaryCreate {
	if t != nil {
		dc.SetUpdatedAt(*t)
	}
	return dc
}

// SetDelete sets the "delete" field.
func (dc *DictionaryCreate) SetDelete(i int64) *DictionaryCreate {
	dc.mutation.SetDelete(i)
	return dc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableDelete(i *int64) *DictionaryCreate {
	if i != nil {
		dc.SetDelete(*i)
	}
	return dc
}

// SetCreatedID sets the "created_id" field.
func (dc *DictionaryCreate) SetCreatedID(i int64) *DictionaryCreate {
	dc.mutation.SetCreatedID(i)
	return dc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableCreatedID(i *int64) *DictionaryCreate {
	if i != nil {
		dc.SetCreatedID(*i)
	}
	return dc
}

// SetStatus sets the "status" field.
func (dc *DictionaryCreate) SetStatus(i int64) *DictionaryCreate {
	dc.mutation.SetStatus(i)
	return dc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (dc *DictionaryCreate) SetNillableStatus(i *int64) *DictionaryCreate {
	if i != nil {
		dc.SetStatus(*i)
	}
	return dc
}

// SetTitle sets the "title" field.
func (dc *DictionaryCreate) SetTitle(s string) *DictionaryCreate {
	dc.mutation.SetTitle(s)
	return dc
}

// SetName sets the "name" field.
func (dc *DictionaryCreate) SetName(s string) *DictionaryCreate {
	dc.mutation.SetName(s)
	return dc
}

// SetDescription sets the "description" field.
func (dc *DictionaryCreate) SetDescription(s string) *DictionaryCreate {
	dc.mutation.SetDescription(s)
	return dc
}

// SetID sets the "id" field.
func (dc *DictionaryCreate) SetID(i int64) *DictionaryCreate {
	dc.mutation.SetID(i)
	return dc
}

// AddDictionaryDetailIDs adds the "dictionary_details" edge to the DictionaryDetail entity by IDs.
func (dc *DictionaryCreate) AddDictionaryDetailIDs(ids ...int64) *DictionaryCreate {
	dc.mutation.AddDictionaryDetailIDs(ids...)
	return dc
}

// AddDictionaryDetails adds the "dictionary_details" edges to the DictionaryDetail entity.
func (dc *DictionaryCreate) AddDictionaryDetails(d ...*DictionaryDetail) *DictionaryCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDictionaryDetailIDs(ids...)
}

// Mutation returns the DictionaryMutation object of the builder.
func (dc *DictionaryCreate) Mutation() *DictionaryMutation {
	return dc.mutation
}

// Save creates the Dictionary in the database.
func (dc *DictionaryCreate) Save(ctx context.Context) (*Dictionary, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DictionaryCreate) SaveX(ctx context.Context) *Dictionary {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DictionaryCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DictionaryCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DictionaryCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := dictionary.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
	if _, ok := dc.mutation.UpdatedAt(); !ok {
		v := dictionary.DefaultUpdatedAt()
		dc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dc.mutation.Delete(); !ok {
		v := dictionary.DefaultDelete
		dc.mutation.SetDelete(v)
	}
	if _, ok := dc.mutation.CreatedID(); !ok {
		v := dictionary.DefaultCreatedID
		dc.mutation.SetCreatedID(v)
	}
	if _, ok := dc.mutation.Status(); !ok {
		v := dictionary.DefaultStatus
		dc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DictionaryCreate) check() error {
	if _, ok := dc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Dictionary.title"`)}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Dictionary.name"`)}
	}
	if _, ok := dc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Dictionary.description"`)}
	}
	return nil
}

func (dc *DictionaryCreate) sqlSave(ctx context.Context) (*Dictionary, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DictionaryCreate) createSpec() (*Dictionary, *sqlgraph.CreateSpec) {
	var (
		_node = &Dictionary{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(dictionary.Table, sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeInt64))
	)
	_spec.Schema = dc.schemaConfig.Dictionary
	if id, ok := dc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(dictionary.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := dc.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionary.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := dc.mutation.Delete(); ok {
		_spec.SetField(dictionary.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := dc.mutation.CreatedID(); ok {
		_spec.SetField(dictionary.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(dictionary.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.Title(); ok {
		_spec.SetField(dictionary.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.SetField(dictionary.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := dc.mutation.Description(); ok {
		_spec.SetField(dictionary.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := dc.mutation.DictionaryDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = dc.schemaConfig.DictionaryDetail
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// DictionaryCreateBulk is the builder for creating many Dictionary entities in bulk.
type DictionaryCreateBulk struct {
	config
	err      error
	builders []*DictionaryCreate
}

// Save creates the Dictionary entities in the database.
func (dcb *DictionaryCreateBulk) Save(ctx context.Context) ([]*Dictionary, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Dictionary, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DictionaryMutation)
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
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DictionaryCreateBulk) SaveX(ctx context.Context) []*Dictionary {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DictionaryCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DictionaryCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
