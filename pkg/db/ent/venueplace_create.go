// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/venue"
	"saas/pkg/db/ent/venueplace"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VenuePlaceCreate is the builder for creating a VenuePlace entity.
type VenuePlaceCreate struct {
	config
	mutation *VenuePlaceMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (vpc *VenuePlaceCreate) SetCreatedAt(t time.Time) *VenuePlaceCreate {
	vpc.mutation.SetCreatedAt(t)
	return vpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vpc *VenuePlaceCreate) SetNillableCreatedAt(t *time.Time) *VenuePlaceCreate {
	if t != nil {
		vpc.SetCreatedAt(*t)
	}
	return vpc
}

// SetUpdatedAt sets the "updated_at" field.
func (vpc *VenuePlaceCreate) SetUpdatedAt(t time.Time) *VenuePlaceCreate {
	vpc.mutation.SetUpdatedAt(t)
	return vpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (vpc *VenuePlaceCreate) SetNillableUpdatedAt(t *time.Time) *VenuePlaceCreate {
	if t != nil {
		vpc.SetUpdatedAt(*t)
	}
	return vpc
}

// SetName sets the "name" field.
func (vpc *VenuePlaceCreate) SetName(s string) *VenuePlaceCreate {
	vpc.mutation.SetName(s)
	return vpc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vpc *VenuePlaceCreate) SetNillableName(s *string) *VenuePlaceCreate {
	if s != nil {
		vpc.SetName(*s)
	}
	return vpc
}

// SetPic sets the "pic" field.
func (vpc *VenuePlaceCreate) SetPic(s string) *VenuePlaceCreate {
	vpc.mutation.SetPic(s)
	return vpc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (vpc *VenuePlaceCreate) SetNillablePic(s *string) *VenuePlaceCreate {
	if s != nil {
		vpc.SetPic(*s)
	}
	return vpc
}

// SetVenueID sets the "venue_id" field.
func (vpc *VenuePlaceCreate) SetVenueID(i int64) *VenuePlaceCreate {
	vpc.mutation.SetVenueID(i)
	return vpc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (vpc *VenuePlaceCreate) SetNillableVenueID(i *int64) *VenuePlaceCreate {
	if i != nil {
		vpc.SetVenueID(*i)
	}
	return vpc
}

// SetStatus sets the "status" field.
func (vpc *VenuePlaceCreate) SetStatus(i int64) *VenuePlaceCreate {
	vpc.mutation.SetStatus(i)
	return vpc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (vpc *VenuePlaceCreate) SetNillableStatus(i *int64) *VenuePlaceCreate {
	if i != nil {
		vpc.SetStatus(*i)
	}
	return vpc
}

// SetID sets the "id" field.
func (vpc *VenuePlaceCreate) SetID(i int64) *VenuePlaceCreate {
	vpc.mutation.SetID(i)
	return vpc
}

// SetVenue sets the "venue" edge to the Venue entity.
func (vpc *VenuePlaceCreate) SetVenue(v *Venue) *VenuePlaceCreate {
	return vpc.SetVenueID(v.ID)
}

// Mutation returns the VenuePlaceMutation object of the builder.
func (vpc *VenuePlaceCreate) Mutation() *VenuePlaceMutation {
	return vpc.mutation
}

// Save creates the VenuePlace in the database.
func (vpc *VenuePlaceCreate) Save(ctx context.Context) (*VenuePlace, error) {
	vpc.defaults()
	return withHooks(ctx, vpc.sqlSave, vpc.mutation, vpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vpc *VenuePlaceCreate) SaveX(ctx context.Context) *VenuePlace {
	v, err := vpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vpc *VenuePlaceCreate) Exec(ctx context.Context) error {
	_, err := vpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vpc *VenuePlaceCreate) ExecX(ctx context.Context) {
	if err := vpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vpc *VenuePlaceCreate) defaults() {
	if _, ok := vpc.mutation.CreatedAt(); !ok {
		v := venueplace.DefaultCreatedAt()
		vpc.mutation.SetCreatedAt(v)
	}
	if _, ok := vpc.mutation.UpdatedAt(); !ok {
		v := venueplace.DefaultUpdatedAt()
		vpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := vpc.mutation.Status(); !ok {
		v := venueplace.DefaultStatus
		vpc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vpc *VenuePlaceCreate) check() error {
	if _, ok := vpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "VenuePlace.created_at"`)}
	}
	if _, ok := vpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "VenuePlace.updated_at"`)}
	}
	return nil
}

func (vpc *VenuePlaceCreate) sqlSave(ctx context.Context) (*VenuePlace, error) {
	if err := vpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	vpc.mutation.id = &_node.ID
	vpc.mutation.done = true
	return _node, nil
}

func (vpc *VenuePlaceCreate) createSpec() (*VenuePlace, *sqlgraph.CreateSpec) {
	var (
		_node = &VenuePlace{config: vpc.config}
		_spec = sqlgraph.NewCreateSpec(venueplace.Table, sqlgraph.NewFieldSpec(venueplace.FieldID, field.TypeInt64))
	)
	if id, ok := vpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := vpc.mutation.CreatedAt(); ok {
		_spec.SetField(venueplace.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := vpc.mutation.UpdatedAt(); ok {
		_spec.SetField(venueplace.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := vpc.mutation.Name(); ok {
		_spec.SetField(venueplace.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vpc.mutation.Pic(); ok {
		_spec.SetField(venueplace.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := vpc.mutation.Status(); ok {
		_spec.SetField(venueplace.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if nodes := vpc.mutation.VenueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   venueplace.VenueTable,
			Columns: []string{venueplace.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VenueID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VenuePlaceCreateBulk is the builder for creating many VenuePlace entities in bulk.
type VenuePlaceCreateBulk struct {
	config
	err      error
	builders []*VenuePlaceCreate
}

// Save creates the VenuePlace entities in the database.
func (vpcb *VenuePlaceCreateBulk) Save(ctx context.Context) ([]*VenuePlace, error) {
	if vpcb.err != nil {
		return nil, vpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(vpcb.builders))
	nodes := make([]*VenuePlace, len(vpcb.builders))
	mutators := make([]Mutator, len(vpcb.builders))
	for i := range vpcb.builders {
		func(i int, root context.Context) {
			builder := vpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VenuePlaceMutation)
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
					_, err = mutators[i+1].Mutate(root, vpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vpcb *VenuePlaceCreateBulk) SaveX(ctx context.Context) []*VenuePlace {
	v, err := vpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vpcb *VenuePlaceCreateBulk) Exec(ctx context.Context) error {
	_, err := vpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vpcb *VenuePlaceCreateBulk) ExecX(ctx context.Context) {
	if err := vpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
