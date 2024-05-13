// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/memberproduct"
	"saas/pkg/db/ent/memberproductproperty"
	"saas/pkg/db/ent/venue"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProductPropertyCreate is the builder for creating a MemberProductProperty entity.
type MemberProductPropertyCreate struct {
	config
	mutation *MemberProductPropertyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mppc *MemberProductPropertyCreate) SetCreatedAt(t time.Time) *MemberProductPropertyCreate {
	mppc.mutation.SetCreatedAt(t)
	return mppc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableCreatedAt(t *time.Time) *MemberProductPropertyCreate {
	if t != nil {
		mppc.SetCreatedAt(*t)
	}
	return mppc
}

// SetUpdatedAt sets the "updated_at" field.
func (mppc *MemberProductPropertyCreate) SetUpdatedAt(t time.Time) *MemberProductPropertyCreate {
	mppc.mutation.SetUpdatedAt(t)
	return mppc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableUpdatedAt(t *time.Time) *MemberProductPropertyCreate {
	if t != nil {
		mppc.SetUpdatedAt(*t)
	}
	return mppc
}

// SetStatus sets the "status" field.
func (mppc *MemberProductPropertyCreate) SetStatus(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetStatus(i)
	return mppc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableStatus(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetStatus(*i)
	}
	return mppc
}

// SetMemberID sets the "member_id" field.
func (mppc *MemberProductPropertyCreate) SetMemberID(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetMemberID(i)
	return mppc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableMemberID(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetMemberID(*i)
	}
	return mppc
}

// SetMemberProductID sets the "member_product_id" field.
func (mppc *MemberProductPropertyCreate) SetMemberProductID(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetMemberProductID(i)
	return mppc
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableMemberProductID(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetMemberProductID(*i)
	}
	return mppc
}

// SetPropertyID sets the "property_id" field.
func (mppc *MemberProductPropertyCreate) SetPropertyID(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetPropertyID(i)
	return mppc
}

// SetNillablePropertyID sets the "property_id" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillablePropertyID(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetPropertyID(*i)
	}
	return mppc
}

// SetType sets the "type" field.
func (mppc *MemberProductPropertyCreate) SetType(s string) *MemberProductPropertyCreate {
	mppc.mutation.SetType(s)
	return mppc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableType(s *string) *MemberProductPropertyCreate {
	if s != nil {
		mppc.SetType(*s)
	}
	return mppc
}

// SetName sets the "name" field.
func (mppc *MemberProductPropertyCreate) SetName(s string) *MemberProductPropertyCreate {
	mppc.mutation.SetName(s)
	return mppc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableName(s *string) *MemberProductPropertyCreate {
	if s != nil {
		mppc.SetName(*s)
	}
	return mppc
}

// SetDuration sets the "duration" field.
func (mppc *MemberProductPropertyCreate) SetDuration(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetDuration(i)
	return mppc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableDuration(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetDuration(*i)
	}
	return mppc
}

// SetLength sets the "length" field.
func (mppc *MemberProductPropertyCreate) SetLength(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetLength(i)
	return mppc
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableLength(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetLength(*i)
	}
	return mppc
}

// SetCount sets the "count" field.
func (mppc *MemberProductPropertyCreate) SetCount(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetCount(i)
	return mppc
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableCount(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetCount(*i)
	}
	return mppc
}

// SetCountSurplus sets the "count_surplus" field.
func (mppc *MemberProductPropertyCreate) SetCountSurplus(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetCountSurplus(i)
	return mppc
}

// SetNillableCountSurplus sets the "count_surplus" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableCountSurplus(i *int64) *MemberProductPropertyCreate {
	if i != nil {
		mppc.SetCountSurplus(*i)
	}
	return mppc
}

// SetPrice sets the "price" field.
func (mppc *MemberProductPropertyCreate) SetPrice(f float64) *MemberProductPropertyCreate {
	mppc.mutation.SetPrice(f)
	return mppc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillablePrice(f *float64) *MemberProductPropertyCreate {
	if f != nil {
		mppc.SetPrice(*f)
	}
	return mppc
}

// SetID sets the "id" field.
func (mppc *MemberProductPropertyCreate) SetID(i int64) *MemberProductPropertyCreate {
	mppc.mutation.SetID(i)
	return mppc
}

// SetOwnerID sets the "owner" edge to the MemberProduct entity by ID.
func (mppc *MemberProductPropertyCreate) SetOwnerID(id int64) *MemberProductPropertyCreate {
	mppc.mutation.SetOwnerID(id)
	return mppc
}

// SetNillableOwnerID sets the "owner" edge to the MemberProduct entity by ID if the given value is not nil.
func (mppc *MemberProductPropertyCreate) SetNillableOwnerID(id *int64) *MemberProductPropertyCreate {
	if id != nil {
		mppc = mppc.SetOwnerID(*id)
	}
	return mppc
}

// SetOwner sets the "owner" edge to the MemberProduct entity.
func (mppc *MemberProductPropertyCreate) SetOwner(m *MemberProduct) *MemberProductPropertyCreate {
	return mppc.SetOwnerID(m.ID)
}

// AddVenueIDs adds the "venues" edge to the Venue entity by IDs.
func (mppc *MemberProductPropertyCreate) AddVenueIDs(ids ...int64) *MemberProductPropertyCreate {
	mppc.mutation.AddVenueIDs(ids...)
	return mppc
}

// AddVenues adds the "venues" edges to the Venue entity.
func (mppc *MemberProductPropertyCreate) AddVenues(v ...*Venue) *MemberProductPropertyCreate {
	ids := make([]int64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return mppc.AddVenueIDs(ids...)
}

// Mutation returns the MemberProductPropertyMutation object of the builder.
func (mppc *MemberProductPropertyCreate) Mutation() *MemberProductPropertyMutation {
	return mppc.mutation
}

// Save creates the MemberProductProperty in the database.
func (mppc *MemberProductPropertyCreate) Save(ctx context.Context) (*MemberProductProperty, error) {
	mppc.defaults()
	return withHooks(ctx, mppc.sqlSave, mppc.mutation, mppc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mppc *MemberProductPropertyCreate) SaveX(ctx context.Context) *MemberProductProperty {
	v, err := mppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mppc *MemberProductPropertyCreate) Exec(ctx context.Context) error {
	_, err := mppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mppc *MemberProductPropertyCreate) ExecX(ctx context.Context) {
	if err := mppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mppc *MemberProductPropertyCreate) defaults() {
	if _, ok := mppc.mutation.CreatedAt(); !ok {
		v := memberproductproperty.DefaultCreatedAt()
		mppc.mutation.SetCreatedAt(v)
	}
	if _, ok := mppc.mutation.UpdatedAt(); !ok {
		v := memberproductproperty.DefaultUpdatedAt()
		mppc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mppc.mutation.Status(); !ok {
		v := memberproductproperty.DefaultStatus
		mppc.mutation.SetStatus(v)
	}
	if _, ok := mppc.mutation.Count(); !ok {
		v := memberproductproperty.DefaultCount
		mppc.mutation.SetCount(v)
	}
	if _, ok := mppc.mutation.CountSurplus(); !ok {
		v := memberproductproperty.DefaultCountSurplus
		mppc.mutation.SetCountSurplus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mppc *MemberProductPropertyCreate) check() error {
	if _, ok := mppc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MemberProductProperty.created_at"`)}
	}
	if _, ok := mppc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MemberProductProperty.updated_at"`)}
	}
	return nil
}

func (mppc *MemberProductPropertyCreate) sqlSave(ctx context.Context) (*MemberProductProperty, error) {
	if err := mppc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mppc.mutation.id = &_node.ID
	mppc.mutation.done = true
	return _node, nil
}

func (mppc *MemberProductPropertyCreate) createSpec() (*MemberProductProperty, *sqlgraph.CreateSpec) {
	var (
		_node = &MemberProductProperty{config: mppc.config}
		_spec = sqlgraph.NewCreateSpec(memberproductproperty.Table, sqlgraph.NewFieldSpec(memberproductproperty.FieldID, field.TypeInt64))
	)
	if id, ok := mppc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mppc.mutation.CreatedAt(); ok {
		_spec.SetField(memberproductproperty.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mppc.mutation.UpdatedAt(); ok {
		_spec.SetField(memberproductproperty.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mppc.mutation.Status(); ok {
		_spec.SetField(memberproductproperty.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := mppc.mutation.MemberID(); ok {
		_spec.SetField(memberproductproperty.FieldMemberID, field.TypeInt64, value)
		_node.MemberID = value
	}
	if value, ok := mppc.mutation.PropertyID(); ok {
		_spec.SetField(memberproductproperty.FieldPropertyID, field.TypeInt64, value)
		_node.PropertyID = value
	}
	if value, ok := mppc.mutation.GetType(); ok {
		_spec.SetField(memberproductproperty.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mppc.mutation.Name(); ok {
		_spec.SetField(memberproductproperty.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mppc.mutation.Duration(); ok {
		_spec.SetField(memberproductproperty.FieldDuration, field.TypeInt64, value)
		_node.Duration = value
	}
	if value, ok := mppc.mutation.Length(); ok {
		_spec.SetField(memberproductproperty.FieldLength, field.TypeInt64, value)
		_node.Length = value
	}
	if value, ok := mppc.mutation.Count(); ok {
		_spec.SetField(memberproductproperty.FieldCount, field.TypeInt64, value)
		_node.Count = value
	}
	if value, ok := mppc.mutation.CountSurplus(); ok {
		_spec.SetField(memberproductproperty.FieldCountSurplus, field.TypeInt64, value)
		_node.CountSurplus = value
	}
	if value, ok := mppc.mutation.Price(); ok {
		_spec.SetField(memberproductproperty.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if nodes := mppc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductproperty.OwnerTable,
			Columns: []string{memberproductproperty.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mppc.mutation.VenuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   memberproductproperty.VenuesTable,
			Columns: memberproductproperty.VenuesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberProductPropertyCreateBulk is the builder for creating many MemberProductProperty entities in bulk.
type MemberProductPropertyCreateBulk struct {
	config
	err      error
	builders []*MemberProductPropertyCreate
}

// Save creates the MemberProductProperty entities in the database.
func (mppcb *MemberProductPropertyCreateBulk) Save(ctx context.Context) ([]*MemberProductProperty, error) {
	if mppcb.err != nil {
		return nil, mppcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mppcb.builders))
	nodes := make([]*MemberProductProperty, len(mppcb.builders))
	mutators := make([]Mutator, len(mppcb.builders))
	for i := range mppcb.builders {
		func(i int, root context.Context) {
			builder := mppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberProductPropertyMutation)
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
					_, err = mutators[i+1].Mutate(root, mppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mppcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mppcb *MemberProductPropertyCreateBulk) SaveX(ctx context.Context) []*MemberProductProperty {
	v, err := mppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mppcb *MemberProductPropertyCreateBulk) Exec(ctx context.Context) error {
	_, err := mppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mppcb *MemberProductPropertyCreateBulk) ExecX(ctx context.Context) {
	if err := mppcb.Exec(ctx); err != nil {
		panic(err)
	}
}
