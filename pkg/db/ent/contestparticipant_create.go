// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/pkg/db/ent/contest"
	"saas/pkg/db/ent/contestparticipant"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContestParticipantCreate is the builder for creating a ContestParticipant entity.
type ContestParticipantCreate struct {
	config
	mutation *ContestParticipantMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cpc *ContestParticipantCreate) SetCreatedAt(t time.Time) *ContestParticipantCreate {
	cpc.mutation.SetCreatedAt(t)
	return cpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableCreatedAt(t *time.Time) *ContestParticipantCreate {
	if t != nil {
		cpc.SetCreatedAt(*t)
	}
	return cpc
}

// SetUpdatedAt sets the "updated_at" field.
func (cpc *ContestParticipantCreate) SetUpdatedAt(t time.Time) *ContestParticipantCreate {
	cpc.mutation.SetUpdatedAt(t)
	return cpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableUpdatedAt(t *time.Time) *ContestParticipantCreate {
	if t != nil {
		cpc.SetUpdatedAt(*t)
	}
	return cpc
}

// SetDelete sets the "delete" field.
func (cpc *ContestParticipantCreate) SetDelete(i int64) *ContestParticipantCreate {
	cpc.mutation.SetDelete(i)
	return cpc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableDelete(i *int64) *ContestParticipantCreate {
	if i != nil {
		cpc.SetDelete(*i)
	}
	return cpc
}

// SetCreatedID sets the "created_id" field.
func (cpc *ContestParticipantCreate) SetCreatedID(i int64) *ContestParticipantCreate {
	cpc.mutation.SetCreatedID(i)
	return cpc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableCreatedID(i *int64) *ContestParticipantCreate {
	if i != nil {
		cpc.SetCreatedID(*i)
	}
	return cpc
}

// SetStatus sets the "status" field.
func (cpc *ContestParticipantCreate) SetStatus(i int64) *ContestParticipantCreate {
	cpc.mutation.SetStatus(i)
	return cpc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableStatus(i *int64) *ContestParticipantCreate {
	if i != nil {
		cpc.SetStatus(*i)
	}
	return cpc
}

// SetContestID sets the "contest_id" field.
func (cpc *ContestParticipantCreate) SetContestID(i int64) *ContestParticipantCreate {
	cpc.mutation.SetContestID(i)
	return cpc
}

// SetNillableContestID sets the "contest_id" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableContestID(i *int64) *ContestParticipantCreate {
	if i != nil {
		cpc.SetContestID(*i)
	}
	return cpc
}

// SetName sets the "name" field.
func (cpc *ContestParticipantCreate) SetName(s string) *ContestParticipantCreate {
	cpc.mutation.SetName(s)
	return cpc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableName(s *string) *ContestParticipantCreate {
	if s != nil {
		cpc.SetName(*s)
	}
	return cpc
}

// SetMobile sets the "mobile" field.
func (cpc *ContestParticipantCreate) SetMobile(s string) *ContestParticipantCreate {
	cpc.mutation.SetMobile(s)
	return cpc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableMobile(s *string) *ContestParticipantCreate {
	if s != nil {
		cpc.SetMobile(*s)
	}
	return cpc
}

// SetFields sets the "fields" field.
func (cpc *ContestParticipantCreate) SetFields(s string) *ContestParticipantCreate {
	cpc.mutation.SetFields(s)
	return cpc
}

// SetNillableFields sets the "fields" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableFields(s *string) *ContestParticipantCreate {
	if s != nil {
		cpc.SetFields(*s)
	}
	return cpc
}

// SetOrderID sets the "order_id" field.
func (cpc *ContestParticipantCreate) SetOrderID(i int64) *ContestParticipantCreate {
	cpc.mutation.SetOrderID(i)
	return cpc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableOrderID(i *int64) *ContestParticipantCreate {
	if i != nil {
		cpc.SetOrderID(*i)
	}
	return cpc
}

// SetOrderSn sets the "order_sn" field.
func (cpc *ContestParticipantCreate) SetOrderSn(s string) *ContestParticipantCreate {
	cpc.mutation.SetOrderSn(s)
	return cpc
}

// SetNillableOrderSn sets the "order_sn" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableOrderSn(s *string) *ContestParticipantCreate {
	if s != nil {
		cpc.SetOrderSn(*s)
	}
	return cpc
}

// SetFee sets the "fee" field.
func (cpc *ContestParticipantCreate) SetFee(f float64) *ContestParticipantCreate {
	cpc.mutation.SetFee(f)
	return cpc
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (cpc *ContestParticipantCreate) SetNillableFee(f *float64) *ContestParticipantCreate {
	if f != nil {
		cpc.SetFee(*f)
	}
	return cpc
}

// SetID sets the "id" field.
func (cpc *ContestParticipantCreate) SetID(i int64) *ContestParticipantCreate {
	cpc.mutation.SetID(i)
	return cpc
}

// SetContest sets the "contest" edge to the Contest entity.
func (cpc *ContestParticipantCreate) SetContest(c *Contest) *ContestParticipantCreate {
	return cpc.SetContestID(c.ID)
}

// Mutation returns the ContestParticipantMutation object of the builder.
func (cpc *ContestParticipantCreate) Mutation() *ContestParticipantMutation {
	return cpc.mutation
}

// Save creates the ContestParticipant in the database.
func (cpc *ContestParticipantCreate) Save(ctx context.Context) (*ContestParticipant, error) {
	cpc.defaults()
	return withHooks(ctx, cpc.sqlSave, cpc.mutation, cpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cpc *ContestParticipantCreate) SaveX(ctx context.Context) *ContestParticipant {
	v, err := cpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpc *ContestParticipantCreate) Exec(ctx context.Context) error {
	_, err := cpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpc *ContestParticipantCreate) ExecX(ctx context.Context) {
	if err := cpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpc *ContestParticipantCreate) defaults() {
	if _, ok := cpc.mutation.CreatedAt(); !ok {
		v := contestparticipant.DefaultCreatedAt()
		cpc.mutation.SetCreatedAt(v)
	}
	if _, ok := cpc.mutation.UpdatedAt(); !ok {
		v := contestparticipant.DefaultUpdatedAt()
		cpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cpc.mutation.Delete(); !ok {
		v := contestparticipant.DefaultDelete
		cpc.mutation.SetDelete(v)
	}
	if _, ok := cpc.mutation.CreatedID(); !ok {
		v := contestparticipant.DefaultCreatedID
		cpc.mutation.SetCreatedID(v)
	}
	if _, ok := cpc.mutation.Status(); !ok {
		v := contestparticipant.DefaultStatus
		cpc.mutation.SetStatus(v)
	}
	if _, ok := cpc.mutation.OrderID(); !ok {
		v := contestparticipant.DefaultOrderID
		cpc.mutation.SetOrderID(v)
	}
	if _, ok := cpc.mutation.OrderSn(); !ok {
		v := contestparticipant.DefaultOrderSn
		cpc.mutation.SetOrderSn(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cpc *ContestParticipantCreate) check() error {
	return nil
}

func (cpc *ContestParticipantCreate) sqlSave(ctx context.Context) (*ContestParticipant, error) {
	if err := cpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	cpc.mutation.id = &_node.ID
	cpc.mutation.done = true
	return _node, nil
}

func (cpc *ContestParticipantCreate) createSpec() (*ContestParticipant, *sqlgraph.CreateSpec) {
	var (
		_node = &ContestParticipant{config: cpc.config}
		_spec = sqlgraph.NewCreateSpec(contestparticipant.Table, sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64))
	)
	if id, ok := cpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cpc.mutation.CreatedAt(); ok {
		_spec.SetField(contestparticipant.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cpc.mutation.UpdatedAt(); ok {
		_spec.SetField(contestparticipant.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cpc.mutation.Delete(); ok {
		_spec.SetField(contestparticipant.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := cpc.mutation.CreatedID(); ok {
		_spec.SetField(contestparticipant.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := cpc.mutation.Status(); ok {
		_spec.SetField(contestparticipant.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := cpc.mutation.Name(); ok {
		_spec.SetField(contestparticipant.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cpc.mutation.Mobile(); ok {
		_spec.SetField(contestparticipant.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := cpc.mutation.GetFields(); ok {
		_spec.SetField(contestparticipant.FieldFields, field.TypeString, value)
		_node.Fields = value
	}
	if value, ok := cpc.mutation.OrderID(); ok {
		_spec.SetField(contestparticipant.FieldOrderID, field.TypeInt64, value)
		_node.OrderID = value
	}
	if value, ok := cpc.mutation.OrderSn(); ok {
		_spec.SetField(contestparticipant.FieldOrderSn, field.TypeString, value)
		_node.OrderSn = value
	}
	if value, ok := cpc.mutation.Fee(); ok {
		_spec.SetField(contestparticipant.FieldFee, field.TypeFloat64, value)
		_node.Fee = value
	}
	if nodes := cpc.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contestparticipant.ContestTable,
			Columns: []string{contestparticipant.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ContestID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContestParticipantCreateBulk is the builder for creating many ContestParticipant entities in bulk.
type ContestParticipantCreateBulk struct {
	config
	err      error
	builders []*ContestParticipantCreate
}

// Save creates the ContestParticipant entities in the database.
func (cpcb *ContestParticipantCreateBulk) Save(ctx context.Context) ([]*ContestParticipant, error) {
	if cpcb.err != nil {
		return nil, cpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cpcb.builders))
	nodes := make([]*ContestParticipant, len(cpcb.builders))
	mutators := make([]Mutator, len(cpcb.builders))
	for i := range cpcb.builders {
		func(i int, root context.Context) {
			builder := cpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContestParticipantMutation)
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
					_, err = mutators[i+1].Mutate(root, cpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, cpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cpcb *ContestParticipantCreateBulk) SaveX(ctx context.Context) []*ContestParticipant {
	v, err := cpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cpcb *ContestParticipantCreateBulk) Exec(ctx context.Context) error {
	_, err := cpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpcb *ContestParticipantCreateBulk) ExecX(ctx context.Context) {
	if err := cpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
