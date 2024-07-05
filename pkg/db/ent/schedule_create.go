// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/schedule"
	"saas/pkg/db/ent/schedulecoach"
	"saas/pkg/db/ent/schedulemember"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ScheduleCreate is the builder for creating a Schedule entity.
type ScheduleCreate struct {
	config
	mutation *ScheduleMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (sc *ScheduleCreate) SetCreatedAt(t time.Time) *ScheduleCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableCreatedAt(t *time.Time) *ScheduleCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *ScheduleCreate) SetUpdatedAt(t time.Time) *ScheduleCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableUpdatedAt(t *time.Time) *ScheduleCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetStatus sets the "status" field.
func (sc *ScheduleCreate) SetStatus(i int64) *ScheduleCreate {
	sc.mutation.SetStatus(i)
	return sc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableStatus(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetStatus(*i)
	}
	return sc
}

// SetType sets the "type" field.
func (sc *ScheduleCreate) SetType(s string) *ScheduleCreate {
	sc.mutation.SetType(s)
	return sc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableType(s *string) *ScheduleCreate {
	if s != nil {
		sc.SetType(*s)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *ScheduleCreate) SetName(s string) *ScheduleCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableName(s *string) *ScheduleCreate {
	if s != nil {
		sc.SetName(*s)
	}
	return sc
}

// SetVenueID sets the "venue_id" field.
func (sc *ScheduleCreate) SetVenueID(i int64) *ScheduleCreate {
	sc.mutation.SetVenueID(i)
	return sc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableVenueID(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetVenueID(*i)
	}
	return sc
}

// SetPropertyID sets the "property_id" field.
func (sc *ScheduleCreate) SetPropertyID(i int64) *ScheduleCreate {
	sc.mutation.SetPropertyID(i)
	return sc
}

// SetNillablePropertyID sets the "property_id" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillablePropertyID(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetPropertyID(*i)
	}
	return sc
}

// SetLength sets the "length" field.
func (sc *ScheduleCreate) SetLength(i int64) *ScheduleCreate {
	sc.mutation.SetLength(i)
	return sc
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableLength(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetLength(*i)
	}
	return sc
}

// SetPlaceID sets the "place_id" field.
func (sc *ScheduleCreate) SetPlaceID(i int64) *ScheduleCreate {
	sc.mutation.SetPlaceID(i)
	return sc
}

// SetNillablePlaceID sets the "place_id" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillablePlaceID(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetPlaceID(*i)
	}
	return sc
}

// SetNum sets the "num" field.
func (sc *ScheduleCreate) SetNum(i int64) *ScheduleCreate {
	sc.mutation.SetNum(i)
	return sc
}

// SetNillableNum sets the "num" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableNum(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetNum(*i)
	}
	return sc
}

// SetNumSurplus sets the "num_surplus" field.
func (sc *ScheduleCreate) SetNumSurplus(i int64) *ScheduleCreate {
	sc.mutation.SetNumSurplus(i)
	return sc
}

// SetNillableNumSurplus sets the "num_surplus" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableNumSurplus(i *int64) *ScheduleCreate {
	if i != nil {
		sc.SetNumSurplus(*i)
	}
	return sc
}

// SetDate sets the "date" field.
func (sc *ScheduleCreate) SetDate(s string) *ScheduleCreate {
	sc.mutation.SetDate(s)
	return sc
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableDate(s *string) *ScheduleCreate {
	if s != nil {
		sc.SetDate(*s)
	}
	return sc
}

// SetStartTime sets the "start_time" field.
func (sc *ScheduleCreate) SetStartTime(t time.Time) *ScheduleCreate {
	sc.mutation.SetStartTime(t)
	return sc
}

// SetNillableStartTime sets the "start_time" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableStartTime(t *time.Time) *ScheduleCreate {
	if t != nil {
		sc.SetStartTime(*t)
	}
	return sc
}

// SetEndTime sets the "end_time" field.
func (sc *ScheduleCreate) SetEndTime(t time.Time) *ScheduleCreate {
	sc.mutation.SetEndTime(t)
	return sc
}

// SetNillableEndTime sets the "end_time" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableEndTime(t *time.Time) *ScheduleCreate {
	if t != nil {
		sc.SetEndTime(*t)
	}
	return sc
}

// SetPrice sets the "price" field.
func (sc *ScheduleCreate) SetPrice(f float64) *ScheduleCreate {
	sc.mutation.SetPrice(f)
	return sc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillablePrice(f *float64) *ScheduleCreate {
	if f != nil {
		sc.SetPrice(*f)
	}
	return sc
}

// SetRemark sets the "remark" field.
func (sc *ScheduleCreate) SetRemark(s string) *ScheduleCreate {
	sc.mutation.SetRemark(s)
	return sc
}

// SetNillableRemark sets the "remark" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableRemark(s *string) *ScheduleCreate {
	if s != nil {
		sc.SetRemark(*s)
	}
	return sc
}

// SetVenueName sets the "venue_name" field.
func (sc *ScheduleCreate) SetVenueName(s string) *ScheduleCreate {
	sc.mutation.SetVenueName(s)
	return sc
}

// SetNillableVenueName sets the "venue_name" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillableVenueName(s *string) *ScheduleCreate {
	if s != nil {
		sc.SetVenueName(*s)
	}
	return sc
}

// SetPlaceName sets the "place_name" field.
func (sc *ScheduleCreate) SetPlaceName(s string) *ScheduleCreate {
	sc.mutation.SetPlaceName(s)
	return sc
}

// SetNillablePlaceName sets the "place_name" field if the given value is not nil.
func (sc *ScheduleCreate) SetNillablePlaceName(s *string) *ScheduleCreate {
	if s != nil {
		sc.SetPlaceName(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *ScheduleCreate) SetID(i int64) *ScheduleCreate {
	sc.mutation.SetID(i)
	return sc
}

// AddMemberIDs adds the "members" edge to the ScheduleMember entity by IDs.
func (sc *ScheduleCreate) AddMemberIDs(ids ...int64) *ScheduleCreate {
	sc.mutation.AddMemberIDs(ids...)
	return sc
}

// AddMembers adds the "members" edges to the ScheduleMember entity.
func (sc *ScheduleCreate) AddMembers(s ...*ScheduleMember) *ScheduleCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddMemberIDs(ids...)
}

// AddCoachIDs adds the "coachs" edge to the ScheduleCoach entity by IDs.
func (sc *ScheduleCreate) AddCoachIDs(ids ...int64) *ScheduleCreate {
	sc.mutation.AddCoachIDs(ids...)
	return sc
}

// AddCoachs adds the "coachs" edges to the ScheduleCoach entity.
func (sc *ScheduleCreate) AddCoachs(s ...*ScheduleCoach) *ScheduleCreate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddCoachIDs(ids...)
}

// Mutation returns the ScheduleMutation object of the builder.
func (sc *ScheduleCreate) Mutation() *ScheduleMutation {
	return sc.mutation
}

// Save creates the Schedule in the database.
func (sc *ScheduleCreate) Save(ctx context.Context) (*Schedule, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScheduleCreate) SaveX(ctx context.Context) *Schedule {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ScheduleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ScheduleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ScheduleCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := schedule.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := schedule.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.Status(); !ok {
		v := schedule.DefaultStatus
		sc.mutation.SetStatus(v)
	}
	if _, ok := sc.mutation.Price(); !ok {
		v := schedule.DefaultPrice
		sc.mutation.SetPrice(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ScheduleCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Schedule.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Schedule.updated_at"`)}
	}
	return nil
}

func (sc *ScheduleCreate) sqlSave(ctx context.Context) (*Schedule, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *ScheduleCreate) createSpec() (*Schedule, *sqlgraph.CreateSpec) {
	var (
		_node = &Schedule{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(schedule.Table, sqlgraph.NewFieldSpec(schedule.FieldID, field.TypeInt64))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(schedule.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(schedule.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.SetField(schedule.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.SetField(schedule.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(schedule.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.VenueID(); ok {
		_spec.SetField(schedule.FieldVenueID, field.TypeInt64, value)
		_node.VenueID = value
	}
	if value, ok := sc.mutation.PropertyID(); ok {
		_spec.SetField(schedule.FieldPropertyID, field.TypeInt64, value)
		_node.PropertyID = value
	}
	if value, ok := sc.mutation.Length(); ok {
		_spec.SetField(schedule.FieldLength, field.TypeInt64, value)
		_node.Length = value
	}
	if value, ok := sc.mutation.PlaceID(); ok {
		_spec.SetField(schedule.FieldPlaceID, field.TypeInt64, value)
		_node.PlaceID = value
	}
	if value, ok := sc.mutation.Num(); ok {
		_spec.SetField(schedule.FieldNum, field.TypeInt64, value)
		_node.Num = value
	}
	if value, ok := sc.mutation.NumSurplus(); ok {
		_spec.SetField(schedule.FieldNumSurplus, field.TypeInt64, value)
		_node.NumSurplus = value
	}
	if value, ok := sc.mutation.Date(); ok {
		_spec.SetField(schedule.FieldDate, field.TypeString, value)
		_node.Date = value
	}
	if value, ok := sc.mutation.StartTime(); ok {
		_spec.SetField(schedule.FieldStartTime, field.TypeTime, value)
		_node.StartTime = value
	}
	if value, ok := sc.mutation.EndTime(); ok {
		_spec.SetField(schedule.FieldEndTime, field.TypeTime, value)
		_node.EndTime = value
	}
	if value, ok := sc.mutation.Price(); ok {
		_spec.SetField(schedule.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := sc.mutation.Remark(); ok {
		_spec.SetField(schedule.FieldRemark, field.TypeString, value)
		_node.Remark = value
	}
	if value, ok := sc.mutation.VenueName(); ok {
		_spec.SetField(schedule.FieldVenueName, field.TypeString, value)
		_node.VenueName = value
	}
	if value, ok := sc.mutation.PlaceName(); ok {
		_spec.SetField(schedule.FieldPlaceName, field.TypeString, value)
		_node.PlaceName = value
	}
	if nodes := sc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.MembersTable,
			Columns: []string{schedule.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedulemember.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.CoachsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   schedule.CoachsTable,
			Columns: []string{schedule.CoachsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(schedulecoach.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ScheduleCreateBulk is the builder for creating many Schedule entities in bulk.
type ScheduleCreateBulk struct {
	config
	err      error
	builders []*ScheduleCreate
}

// Save creates the Schedule entities in the database.
func (scb *ScheduleCreateBulk) Save(ctx context.Context) ([]*Schedule, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Schedule, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScheduleMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ScheduleCreateBulk) SaveX(ctx context.Context) []*Schedule {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ScheduleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ScheduleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
