// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/banner"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BannerCreate is the builder for creating a Banner entity.
type BannerCreate struct {
	config
	mutation *BannerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BannerCreate) SetCreatedAt(t time.Time) *BannerCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BannerCreate) SetNillableCreatedAt(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BannerCreate) SetUpdatedAt(t time.Time) *BannerCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BannerCreate) SetNillableUpdatedAt(t *time.Time) *BannerCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetDeleteAt sets the "delete_at" field.
func (bc *BannerCreate) SetDeleteAt(t time.Time) *BannerCreate {
	bc.mutation.SetDeleteAt(t)
	return bc
}

// SetCreatedID sets the "created_id" field.
func (bc *BannerCreate) SetCreatedID(i int64) *BannerCreate {
	bc.mutation.SetCreatedID(i)
	return bc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (bc *BannerCreate) SetNillableCreatedID(i *int64) *BannerCreate {
	if i != nil {
		bc.SetCreatedID(*i)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BannerCreate) SetStatus(i int64) *BannerCreate {
	bc.mutation.SetStatus(i)
	return bc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bc *BannerCreate) SetNillableStatus(i *int64) *BannerCreate {
	if i != nil {
		bc.SetStatus(*i)
	}
	return bc
}

// SetName sets the "name" field.
func (bc *BannerCreate) SetName(s string) *BannerCreate {
	bc.mutation.SetName(s)
	return bc
}

// SetPic sets the "pic" field.
func (bc *BannerCreate) SetPic(s string) *BannerCreate {
	bc.mutation.SetPic(s)
	return bc
}

// SetLink sets the "link" field.
func (bc *BannerCreate) SetLink(s string) *BannerCreate {
	bc.mutation.SetLink(s)
	return bc
}

// SetIsShow sets the "is_show" field.
func (bc *BannerCreate) SetIsShow(i int64) *BannerCreate {
	bc.mutation.SetIsShow(i)
	return bc
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (bc *BannerCreate) SetNillableIsShow(i *int64) *BannerCreate {
	if i != nil {
		bc.SetIsShow(*i)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BannerCreate) SetID(i int64) *BannerCreate {
	bc.mutation.SetID(i)
	return bc
}

// Mutation returns the BannerMutation object of the builder.
func (bc *BannerCreate) Mutation() *BannerMutation {
	return bc.mutation
}

// Save creates the Banner in the database.
func (bc *BannerCreate) Save(ctx context.Context) (*Banner, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BannerCreate) SaveX(ctx context.Context) *Banner {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BannerCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BannerCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BannerCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := banner.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := banner.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.CreatedID(); !ok {
		v := banner.DefaultCreatedID
		bc.mutation.SetCreatedID(v)
	}
	if _, ok := bc.mutation.Status(); !ok {
		v := banner.DefaultStatus
		bc.mutation.SetStatus(v)
	}
	if _, ok := bc.mutation.IsShow(); !ok {
		v := banner.DefaultIsShow
		bc.mutation.SetIsShow(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BannerCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Banner.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Banner.updated_at"`)}
	}
	if _, ok := bc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "Banner.delete_at"`)}
	}
	if _, ok := bc.mutation.CreatedID(); !ok {
		return &ValidationError{Name: "created_id", err: errors.New(`ent: missing required field "Banner.created_id"`)}
	}
	if _, ok := bc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Banner.name"`)}
	}
	if _, ok := bc.mutation.Pic(); !ok {
		return &ValidationError{Name: "pic", err: errors.New(`ent: missing required field "Banner.pic"`)}
	}
	if _, ok := bc.mutation.Link(); !ok {
		return &ValidationError{Name: "link", err: errors.New(`ent: missing required field "Banner.link"`)}
	}
	return nil
}

func (bc *BannerCreate) sqlSave(ctx context.Context) (*Banner, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BannerCreate) createSpec() (*Banner, *sqlgraph.CreateSpec) {
	var (
		_node = &Banner{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(banner.Table, sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt64))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(banner.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(banner.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.DeleteAt(); ok {
		_spec.SetField(banner.FieldDeleteAt, field.TypeTime, value)
		_node.DeleteAt = value
	}
	if value, ok := bc.mutation.CreatedID(); ok {
		_spec.SetField(banner.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(banner.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.Name(); ok {
		_spec.SetField(banner.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := bc.mutation.Pic(); ok {
		_spec.SetField(banner.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := bc.mutation.Link(); ok {
		_spec.SetField(banner.FieldLink, field.TypeString, value)
		_node.Link = value
	}
	if value, ok := bc.mutation.IsShow(); ok {
		_spec.SetField(banner.FieldIsShow, field.TypeInt64, value)
		_node.IsShow = value
	}
	return _node, _spec
}

// BannerCreateBulk is the builder for creating many Banner entities in bulk.
type BannerCreateBulk struct {
	config
	err      error
	builders []*BannerCreate
}

// Save creates the Banner entities in the database.
func (bcb *BannerCreateBulk) Save(ctx context.Context) ([]*Banner, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Banner, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BannerMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BannerCreateBulk) SaveX(ctx context.Context) []*Banner {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BannerCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BannerCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
