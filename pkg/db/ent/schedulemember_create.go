// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/schedule"
	"saas/pkg/db/ent/schedulemember"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleMemberCreate is the builder for creating a ScheduleMember entity.
type ScheduleMemberCreate struct {
	config
	mutation *ScheduleMemberMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (smc *ScheduleMemberCreate) SetCreatedAt(t time.Time) *ScheduleMemberCreate {
	smc.mutation.SetCreatedAt(t)
	return smc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableCreatedAt(t *time.Time) *ScheduleMemberCreate {
	if t != nil {
		smc.SetCreatedAt(*t)
	}
	return smc
}

// SetUpdatedAt sets the "updated_at" field.
func (smc *ScheduleMemberCreate) SetUpdatedAt(t time.Time) *ScheduleMemberCreate {
	smc.mutation.SetUpdatedAt(t)
	return smc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableUpdatedAt(t *time.Time) *ScheduleMemberCreate {
	if t != nil {
		smc.SetUpdatedAt(*t)
	}
	return smc
}

// SetStatus sets the "status" field.
func (smc *ScheduleMemberCreate) SetStatus(i int64) *ScheduleMemberCreate {
	smc.mutation.SetStatus(i)
	return smc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableStatus(i *int64) *ScheduleMemberCreate {
	if i != nil {
		smc.SetStatus(*i)
	}
	return smc
}

// SetVenueID sets the "venue_id" field.
func (smc *ScheduleMemberCreate) SetVenueID(i int64) *ScheduleMemberCreate {
	smc.mutation.SetVenueID(i)
	return smc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableVenueID(i *int64) *ScheduleMemberCreate {
	if i != nil {
		smc.SetVenueID(*i)
	}
	return smc
}

// SetScheduleID sets the "schedule_id" field.
func (smc *ScheduleMemberCreate) SetScheduleID(i int64) *ScheduleMemberCreate {
	smc.mutation.SetScheduleID(i)
	return smc
}

// SetNillableScheduleID sets the "schedule_id" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableScheduleID(i *int64) *ScheduleMemberCreate {
	if i != nil {
		smc.SetScheduleID(*i)
	}
	return smc
}

// SetMemberID sets the "member_id" field.
func (smc *ScheduleMemberCreate) SetMemberID(i int64) *ScheduleMemberCreate {
	smc.mutation.SetMemberID(i)
	return smc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableMemberID(i *int64) *ScheduleMemberCreate {
	if i != nil {
		smc.SetMemberID(*i)
	}
	return smc
}

// SetMemberProductID sets the "member_product_id" field.
func (smc *ScheduleMemberCreate) SetMemberProductID(i int64) *ScheduleMemberCreate {
	smc.mutation.SetMemberProductID(i)
	return smc
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableMemberProductID(i *int64) *ScheduleMemberCreate {
	if i != nil {
		smc.SetMemberProductID(*i)
	}
	return smc
}

// SetMemberProductPropertyID sets the "member_product_property_id" field.
func (smc *ScheduleMemberCreate) SetMemberProductPropertyID(i int64) *ScheduleMemberCreate {
	smc.mutation.SetMemberProductPropertyID(i)
	return smc
}

// SetNillableMemberProductPropertyID sets the "member_product_property_id" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableMemberProductPropertyID(i *int64) *ScheduleMemberCreate {
	if i != nil {
		smc.SetMemberProductPropertyID(*i)
	}
	return smc
}

// SetType sets the "type" field.
func (smc *ScheduleMemberCreate) SetType(s string) *ScheduleMemberCreate {
	smc.mutation.SetType(s)
	return smc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableType(s *string) *ScheduleMemberCreate {
	if s != nil {
		smc.SetType(*s)
	}
	return smc
}

// SetStartTime sets the "start_time" field.
func (smc *ScheduleMemberCreate) SetStartTime(t time.Time) *ScheduleMemberCreate {
	smc.mutation.SetStartTime(t)
	return smc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableStartTime(t *time.Time) *ScheduleMemberCreate {
	if t != nil {
		smc.SetStartTime(*t)
	}
	return smc
}

// SetEndTime sets the "end_time" field.
func (smc *ScheduleMemberCreate) SetEndTime(t time.Time) *ScheduleMemberCreate {
	smc.mutation.SetEndTime(t)
	return smc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableEndTime(t *time.Time) *ScheduleMemberCreate {
	if t != nil {
		smc.SetEndTime(*t)
	}
	return smc
}

// SetSignStartTime sets the "sign_start_time" field.
func (smc *ScheduleMemberCreate) SetSignStartTime(t time.Time) *ScheduleMemberCreate {
	smc.mutation.SetSignStartTime(t)
	return smc
}

// SetNillableSignStartTime sets the "sign_start_time" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableSignStartTime(t *time.Time) *ScheduleMemberCreate {
	if t != nil {
		smc.SetSignStartTime(*t)
	}
	return smc
}

// SetSignEndTime sets the "sign_end_time" field.
func (smc *ScheduleMemberCreate) SetSignEndTime(t time.Time) *ScheduleMemberCreate {
	smc.mutation.SetSignEndTime(t)
	return smc
}

// SetNillableSignEndTime sets the "sign_end_time" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableSignEndTime(t *time.Time) *ScheduleMemberCreate {
	if t != nil {
		smc.SetSignEndTime(*t)
	}
	return smc
}

// SetMemberName sets the "member_name" field.
func (smc *ScheduleMemberCreate) SetMemberName(s string) *ScheduleMemberCreate {
	smc.mutation.SetMemberName(s)
	return smc
}

// SetNillableMemberName sets the "member_name" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableMemberName(s *string) *ScheduleMemberCreate {
	if s != nil {
		smc.SetMemberName(*s)
	}
	return smc
}

// SetMemberProductName sets the "member_product_name" field.
func (smc *ScheduleMemberCreate) SetMemberProductName(s string) *ScheduleMemberCreate {
	smc.mutation.SetMemberProductName(s)
	return smc
}

// SetNillableMemberProductName sets the "member_product_name" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableMemberProductName(s *string) *ScheduleMemberCreate {
	if s != nil {
		smc.SetMemberProductName(*s)
	}
	return smc
}

// SetMemberProductPropertyName sets the "member_product_property_name" field.
func (smc *ScheduleMemberCreate) SetMemberProductPropertyName(s string) *ScheduleMemberCreate {
	smc.mutation.SetMemberProductPropertyName(s)
	return smc
}

// SetNillableMemberProductPropertyName sets the "member_product_property_name" field if the given value is not nil.
func (smc *ScheduleMemberCreate) SetNillableMemberProductPropertyName(s *string) *ScheduleMemberCreate {
	if s != nil {
		smc.SetMemberProductPropertyName(*s)
	}
	return smc
}

// SetID sets the "id" field.
func (smc *ScheduleMemberCreate) SetID(i int64) *ScheduleMemberCreate {
	smc.mutation.SetID(i)
	return smc
}

// SetSchedule sets the "schedule" edge to the Schedule entity.
func (smc *ScheduleMemberCreate) SetSchedule(s *Schedule) *ScheduleMemberCreate {
	return smc.SetScheduleID(s.ID)
}

// Mutation returns the ScheduleMemberMutation object of the builder.
func (smc *ScheduleMemberCreate) Mutation() *ScheduleMemberMutation {
	return smc.mutation
}

// Save creates the ScheduleMember in the database.
func (smc *ScheduleMemberCreate) Save(ctx context.Context) (*ScheduleMember, error) {
	smc.defaults()
	return withHooks(ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *ScheduleMemberCreate) SaveX(ctx context.Context) *ScheduleMember {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *ScheduleMemberCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *ScheduleMemberCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smc *ScheduleMemberCreate) defaults() {
	if _, ok := smc.mutation.CreatedAt(); !ok {
		v := schedulemember.DefaultCreatedAt()
		smc.mutation.SetCreatedAt(v)
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		v := schedulemember.DefaultUpdatedAt()
		smc.mutation.SetUpdatedAt(v)
	}
	if _, ok := smc.mutation.Status(); !ok {
		v := schedulemember.DefaultStatus
		smc.mutation.SetStatus(v)
	}
	if _, ok := smc.mutation.StartTime(); !ok {
		v := schedulemember.DefaultStartTime()
		smc.mutation.SetStartTime(v)
	}
	if _, ok := smc.mutation.EndTime(); !ok {
		v := schedulemember.DefaultEndTime()
		smc.mutation.SetEndTime(v)
	}
	if _, ok := smc.mutation.SignStartTime(); !ok {
		v := schedulemember.DefaultSignStartTime()
		smc.mutation.SetSignStartTime(v)
	}
	if _, ok := smc.mutation.SignEndTime(); !ok {
		v := schedulemember.DefaultSignEndTime()
		smc.mutation.SetSignEndTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *ScheduleMemberCreate) check() error {
	if _, ok := smc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ScheduleMember.created_at"`)}
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ScheduleMember.updated_at"`)}
	}
	return nil
}

func (smc *ScheduleMemberCreate) sqlSave(ctx context.Context) (*ScheduleMember, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *ScheduleMemberCreate) createSpec() (*ScheduleMember, *sqlgraph.CreateSpec) {
	var (
		_node = &ScheduleMember{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(schedulemember.Table, sqlgraph.NewFieldSpec(schedulemember.FieldID, field.TypeInt64))
	)
	if id, ok := smc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := smc.mutation.CreatedAt(); ok {
		_spec.SetField(schedulemember.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := smc.mutation.UpdatedAt(); ok {
		_spec.SetField(schedulemember.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := smc.mutation.Status(); ok {
		_spec.SetField(schedulemember.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := smc.mutation.VenueID(); ok {
		_spec.SetField(schedulemember.FieldVenueID, field.TypeInt64, value)
		_node.VenueID = value
	}
	if value, ok := smc.mutation.MemberID(); ok {
		_spec.SetField(schedulemember.FieldMemberID, field.TypeInt64, value)
		_node.MemberID = value
	}
	if value, ok := smc.mutation.MemberProductID(); ok {
		_spec.SetField(schedulemember.FieldMemberProductID, field.TypeInt64, value)
		_node.MemberProductID = value
	}
	if value, ok := smc.mutation.MemberProductPropertyID(); ok {
		_spec.SetField(schedulemember.FieldMemberProductPropertyID, field.TypeInt64, value)
		_node.MemberProductPropertyID = value
	}
	if value, ok := smc.mutation.GetType(); ok {
		_spec.SetField(schedulemember.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := smc.mutation.StartTime(); ok {
		_spec.SetField(schedulemember.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := smc.mutation.EndTime(); ok {
		_spec.SetField(schedulemember.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := smc.mutation.SignStartTime(); ok {
		_spec.SetField(schedulemember.FieldSignStartTime, field.TypeTime, value)
		_node.SignStartTime = value
	}
	if value, ok := smc.mutation.SignEndTime(); ok {
		_spec.SetField(schedulemember.FieldSignEndTime, field.TypeTime, value)
		_node.SignEndTime = value
	}
	if value, ok := smc.mutation.MemberName(); ok {
		_spec.SetField(schedulemember.FieldMemberName, field.TypeString, value)
		_node.MemberName = value
	}
	if value, ok := smc.mutation.MemberProductName(); ok {
		_spec.SetField(schedulemember.FieldMemberProductName, field.TypeString, value)
		_node.MemberProductName = value
	}
	if value, ok := smc.mutation.MemberProductPropertyName(); ok {
		_spec.SetField(schedulemember.FieldMemberProductPropertyName, field.TypeString, value)
		_node.MemberProductPropertyName = value
	}
	if nodes := smc.mutation.ScheduleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   schedulemember.ScheduleTable,
			Columns: []string{schedulemember.ScheduleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ScheduleID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScheduleMemberCreateBulk is the builder for creating many ScheduleMember entities in bulk.
type ScheduleMemberCreateBulk struct {
	config
	err      error
	builders []*ScheduleMemberCreate
}

// Save creates the ScheduleMember entities in the database.
func (smcb *ScheduleMemberCreateBulk) Save(ctx context.Context) ([]*ScheduleMember, error) {
	if smcb.err != nil {
		return nil, smcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*ScheduleMember, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScheduleMemberMutation)
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
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *ScheduleMemberCreateBulk) SaveX(ctx context.Context) []*ScheduleMember {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *ScheduleMemberCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *ScheduleMemberCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
