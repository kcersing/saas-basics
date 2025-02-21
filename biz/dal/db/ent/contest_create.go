// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/contest"
	"saas/biz/dal/db/ent/contestparticipant"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContestCreate is the builder for creating a Contest entity.
type ContestCreate struct {
	config
	mutation *ContestMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *ContestCreate) SetCreatedAt(t time.Time) *ContestCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableCreatedAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ContestCreate) SetUpdatedAt(t time.Time) *ContestCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableUpdatedAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDelete sets the "delete" field.
func (cc *ContestCreate) SetDelete(i int64) *ContestCreate {
	cc.mutation.SetDelete(i)
	return cc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (cc *ContestCreate) SetNillableDelete(i *int64) *ContestCreate {
	if i != nil {
		cc.SetDelete(*i)
	}
	return cc
}

// SetCreatedID sets the "created_id" field.
func (cc *ContestCreate) SetCreatedID(i int64) *ContestCreate {
	cc.mutation.SetCreatedID(i)
	return cc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (cc *ContestCreate) SetNillableCreatedID(i *int64) *ContestCreate {
	if i != nil {
		cc.SetCreatedID(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *ContestCreate) SetStatus(i int64) *ContestCreate {
	cc.mutation.SetStatus(i)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *ContestCreate) SetNillableStatus(i *int64) *ContestCreate {
	if i != nil {
		cc.SetStatus(*i)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *ContestCreate) SetName(s string) *ContestCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *ContestCreate) SetNillableName(s *string) *ContestCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetSignNumber sets the "sign_number" field.
func (cc *ContestCreate) SetSignNumber(i int64) *ContestCreate {
	cc.mutation.SetSignNumber(i)
	return cc
}

// SetNillableSignNumber sets the "sign_number" field if the given value is not nil.
func (cc *ContestCreate) SetNillableSignNumber(i *int64) *ContestCreate {
	if i != nil {
		cc.SetSignNumber(*i)
	}
	return cc
}

// SetSignStartAt sets the "sign_start_at" field.
func (cc *ContestCreate) SetSignStartAt(t time.Time) *ContestCreate {
	cc.mutation.SetSignStartAt(t)
	return cc
}

// SetNillableSignStartAt sets the "sign_start_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableSignStartAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetSignStartAt(*t)
	}
	return cc
}

// SetSignEndAt sets the "sign_end_at" field.
func (cc *ContestCreate) SetSignEndAt(t time.Time) *ContestCreate {
	cc.mutation.SetSignEndAt(t)
	return cc
}

// SetNillableSignEndAt sets the "sign_end_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableSignEndAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetSignEndAt(*t)
	}
	return cc
}

// SetNumber sets the "number" field.
func (cc *ContestCreate) SetNumber(i int64) *ContestCreate {
	cc.mutation.SetNumber(i)
	return cc
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cc *ContestCreate) SetNillableNumber(i *int64) *ContestCreate {
	if i != nil {
		cc.SetNumber(*i)
	}
	return cc
}

// SetStartAt sets the "start_at" field.
func (cc *ContestCreate) SetStartAt(t time.Time) *ContestCreate {
	cc.mutation.SetStartAt(t)
	return cc
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableStartAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetStartAt(*t)
	}
	return cc
}

// SetEndAt sets the "end_at" field.
func (cc *ContestCreate) SetEndAt(t time.Time) *ContestCreate {
	cc.mutation.SetEndAt(t)
	return cc
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cc *ContestCreate) SetNillableEndAt(t *time.Time) *ContestCreate {
	if t != nil {
		cc.SetEndAt(*t)
	}
	return cc
}

// SetPic sets the "pic" field.
func (cc *ContestCreate) SetPic(s string) *ContestCreate {
	cc.mutation.SetPic(s)
	return cc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (cc *ContestCreate) SetNillablePic(s *string) *ContestCreate {
	if s != nil {
		cc.SetPic(*s)
	}
	return cc
}

// SetSponsor sets the "sponsor" field.
func (cc *ContestCreate) SetSponsor(s string) *ContestCreate {
	cc.mutation.SetSponsor(s)
	return cc
}

// SetNillableSponsor sets the "sponsor" field if the given value is not nil.
func (cc *ContestCreate) SetNillableSponsor(s *string) *ContestCreate {
	if s != nil {
		cc.SetSponsor(*s)
	}
	return cc
}

// SetFee sets the "fee" field.
func (cc *ContestCreate) SetFee(f float64) *ContestCreate {
	cc.mutation.SetFee(f)
	return cc
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (cc *ContestCreate) SetNillableFee(f *float64) *ContestCreate {
	if f != nil {
		cc.SetFee(*f)
	}
	return cc
}

// SetIsFee sets the "is_fee" field.
func (cc *ContestCreate) SetIsFee(i int64) *ContestCreate {
	cc.mutation.SetIsFee(i)
	return cc
}

// SetNillableIsFee sets the "is_fee" field if the given value is not nil.
func (cc *ContestCreate) SetNillableIsFee(i *int64) *ContestCreate {
	if i != nil {
		cc.SetIsFee(*i)
	}
	return cc
}

// SetIsShow sets the "is_show" field.
func (cc *ContestCreate) SetIsShow(i int64) *ContestCreate {
	cc.mutation.SetIsShow(i)
	return cc
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (cc *ContestCreate) SetNillableIsShow(i *int64) *ContestCreate {
	if i != nil {
		cc.SetIsShow(*i)
	}
	return cc
}

// SetIsCancel sets the "is_cancel" field.
func (cc *ContestCreate) SetIsCancel(i int64) *ContestCreate {
	cc.mutation.SetIsCancel(i)
	return cc
}

// SetNillableIsCancel sets the "is_cancel" field if the given value is not nil.
func (cc *ContestCreate) SetNillableIsCancel(i *int64) *ContestCreate {
	if i != nil {
		cc.SetIsCancel(*i)
	}
	return cc
}

// SetCancelTime sets the "cancel_time" field.
func (cc *ContestCreate) SetCancelTime(i int64) *ContestCreate {
	cc.mutation.SetCancelTime(i)
	return cc
}

// SetNillableCancelTime sets the "cancel_time" field if the given value is not nil.
func (cc *ContestCreate) SetNillableCancelTime(i *int64) *ContestCreate {
	if i != nil {
		cc.SetCancelTime(*i)
	}
	return cc
}

// SetDetail sets the "detail" field.
func (cc *ContestCreate) SetDetail(s string) *ContestCreate {
	cc.mutation.SetDetail(s)
	return cc
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (cc *ContestCreate) SetNillableDetail(s *string) *ContestCreate {
	if s != nil {
		cc.SetDetail(*s)
	}
	return cc
}

// SetSignFields sets the "sign_fields" field.
func (cc *ContestCreate) SetSignFields(s string) *ContestCreate {
	cc.mutation.SetSignFields(s)
	return cc
}

// SetNillableSignFields sets the "sign_fields" field if the given value is not nil.
func (cc *ContestCreate) SetNillableSignFields(s *string) *ContestCreate {
	if s != nil {
		cc.SetSignFields(*s)
	}
	return cc
}

// SetCondition sets the "condition" field.
func (cc *ContestCreate) SetCondition(i int64) *ContestCreate {
	cc.mutation.SetCondition(i)
	return cc
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (cc *ContestCreate) SetNillableCondition(i *int64) *ContestCreate {
	if i != nil {
		cc.SetCondition(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ContestCreate) SetID(i int64) *ContestCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddContestParticipantIDs adds the "contest_participants" edge to the ContestParticipant entity by IDs.
func (cc *ContestCreate) AddContestParticipantIDs(ids ...int64) *ContestCreate {
	cc.mutation.AddContestParticipantIDs(ids...)
	return cc
}

// AddContestParticipants adds the "contest_participants" edges to the ContestParticipant entity.
func (cc *ContestCreate) AddContestParticipants(c ...*ContestParticipant) *ContestCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddContestParticipantIDs(ids...)
}

// Mutation returns the ContestMutation object of the builder.
func (cc *ContestCreate) Mutation() *ContestMutation {
	return cc.mutation
}

// Save creates the Contest in the database.
func (cc *ContestCreate) Save(ctx context.Context) (*Contest, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ContestCreate) SaveX(ctx context.Context) *Contest {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ContestCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ContestCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ContestCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := contest.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := contest.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Delete(); !ok {
		v := contest.DefaultDelete
		cc.mutation.SetDelete(v)
	}
	if _, ok := cc.mutation.CreatedID(); !ok {
		v := contest.DefaultCreatedID
		cc.mutation.SetCreatedID(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := contest.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.IsFee(); !ok {
		v := contest.DefaultIsFee
		cc.mutation.SetIsFee(v)
	}
	if _, ok := cc.mutation.IsShow(); !ok {
		v := contest.DefaultIsShow
		cc.mutation.SetIsShow(v)
	}
	if _, ok := cc.mutation.IsCancel(); !ok {
		v := contest.DefaultIsCancel
		cc.mutation.SetIsCancel(v)
	}
	if _, ok := cc.mutation.CancelTime(); !ok {
		v := contest.DefaultCancelTime
		cc.mutation.SetCancelTime(v)
	}
	if _, ok := cc.mutation.Condition(); !ok {
		v := contest.DefaultCondition
		cc.mutation.SetCondition(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ContestCreate) check() error {
	return nil
}

func (cc *ContestCreate) sqlSave(ctx context.Context) (*Contest, error) {
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

func (cc *ContestCreate) createSpec() (*Contest, *sqlgraph.CreateSpec) {
	var (
		_node = &Contest{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(contest.Table, sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64))
	)
	_spec.Schema = cc.schemaConfig.Contest
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(contest.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(contest.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Delete(); ok {
		_spec.SetField(contest.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := cc.mutation.CreatedID(); ok {
		_spec.SetField(contest.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(contest.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(contest.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.SignNumber(); ok {
		_spec.SetField(contest.FieldSignNumber, field.TypeInt64, value)
		_node.SignNumber = value
	}
	if value, ok := cc.mutation.SignStartAt(); ok {
		_spec.SetField(contest.FieldSignStartAt, field.TypeTime, value)
		_node.SignStartAt = value
	}
	if value, ok := cc.mutation.SignEndAt(); ok {
		_spec.SetField(contest.FieldSignEndAt, field.TypeTime, value)
		_node.SignEndAt = value
	}
	if value, ok := cc.mutation.Number(); ok {
		_spec.SetField(contest.FieldNumber, field.TypeInt64, value)
		_node.Number = value
	}
	if value, ok := cc.mutation.StartAt(); ok {
		_spec.SetField(contest.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := cc.mutation.EndAt(); ok {
		_spec.SetField(contest.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := cc.mutation.Pic(); ok {
		_spec.SetField(contest.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := cc.mutation.Sponsor(); ok {
		_spec.SetField(contest.FieldSponsor, field.TypeString, value)
		_node.Sponsor = value
	}
	if value, ok := cc.mutation.Fee(); ok {
		_spec.SetField(contest.FieldFee, field.TypeFloat64, value)
		_node.Fee = value
	}
	if value, ok := cc.mutation.IsFee(); ok {
		_spec.SetField(contest.FieldIsFee, field.TypeInt64, value)
		_node.IsFee = value
	}
	if value, ok := cc.mutation.IsShow(); ok {
		_spec.SetField(contest.FieldIsShow, field.TypeInt64, value)
		_node.IsShow = value
	}
	if value, ok := cc.mutation.IsCancel(); ok {
		_spec.SetField(contest.FieldIsCancel, field.TypeInt64, value)
		_node.IsCancel = value
	}
	if value, ok := cc.mutation.CancelTime(); ok {
		_spec.SetField(contest.FieldCancelTime, field.TypeInt64, value)
		_node.CancelTime = value
	}
	if value, ok := cc.mutation.Detail(); ok {
		_spec.SetField(contest.FieldDetail, field.TypeString, value)
		_node.Detail = value
	}
	if value, ok := cc.mutation.SignFields(); ok {
		_spec.SetField(contest.FieldSignFields, field.TypeString, value)
		_node.SignFields = value
	}
	if value, ok := cc.mutation.Condition(); ok {
		_spec.SetField(contest.FieldCondition, field.TypeInt64, value)
		_node.Condition = value
	}
	if nodes := cc.mutation.ContestParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   contest.ContestParticipantsTable,
			Columns: []string{contest.ContestParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = cc.schemaConfig.ContestParticipant
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ContestCreateBulk is the builder for creating many Contest entities in bulk.
type ContestCreateBulk struct {
	config
	err      error
	builders []*ContestCreate
}

// Save creates the Contest entities in the database.
func (ccb *ContestCreateBulk) Save(ctx context.Context) ([]*Contest, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Contest, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContestMutation)
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
func (ccb *ContestCreateBulk) SaveX(ctx context.Context) []*Contest {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ContestCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ContestCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
