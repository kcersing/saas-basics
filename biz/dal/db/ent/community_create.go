// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/community"
	"saas/biz/dal/db/ent/communityparticipant"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommunityCreate is the builder for creating a Community entity.
type CommunityCreate struct {
	config
	mutation *CommunityMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommunityCreate) SetCreatedAt(t time.Time) *CommunityCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableCreatedAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommunityCreate) SetUpdatedAt(t time.Time) *CommunityCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableUpdatedAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetDelete sets the "delete" field.
func (cc *CommunityCreate) SetDelete(i int64) *CommunityCreate {
	cc.mutation.SetDelete(i)
	return cc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableDelete(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetDelete(*i)
	}
	return cc
}

// SetCreatedID sets the "created_id" field.
func (cc *CommunityCreate) SetCreatedID(i int64) *CommunityCreate {
	cc.mutation.SetCreatedID(i)
	return cc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableCreatedID(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetCreatedID(*i)
	}
	return cc
}

// SetStatus sets the "status" field.
func (cc *CommunityCreate) SetStatus(i int64) *CommunityCreate {
	cc.mutation.SetStatus(i)
	return cc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableStatus(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetStatus(*i)
	}
	return cc
}

// SetName sets the "name" field.
func (cc *CommunityCreate) SetName(s string) *CommunityCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableName(s *string) *CommunityCreate {
	if s != nil {
		cc.SetName(*s)
	}
	return cc
}

// SetSignNumber sets the "sign_number" field.
func (cc *CommunityCreate) SetSignNumber(i int64) *CommunityCreate {
	cc.mutation.SetSignNumber(i)
	return cc
}

// SetNillableSignNumber sets the "sign_number" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableSignNumber(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetSignNumber(*i)
	}
	return cc
}

// SetSignStartAt sets the "sign_start_at" field.
func (cc *CommunityCreate) SetSignStartAt(t time.Time) *CommunityCreate {
	cc.mutation.SetSignStartAt(t)
	return cc
}

// SetNillableSignStartAt sets the "sign_start_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableSignStartAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetSignStartAt(*t)
	}
	return cc
}

// SetSignEndAt sets the "sign_end_at" field.
func (cc *CommunityCreate) SetSignEndAt(t time.Time) *CommunityCreate {
	cc.mutation.SetSignEndAt(t)
	return cc
}

// SetNillableSignEndAt sets the "sign_end_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableSignEndAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetSignEndAt(*t)
	}
	return cc
}

// SetNumber sets the "number" field.
func (cc *CommunityCreate) SetNumber(i int64) *CommunityCreate {
	cc.mutation.SetNumber(i)
	return cc
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableNumber(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetNumber(*i)
	}
	return cc
}

// SetStartAt sets the "start_at" field.
func (cc *CommunityCreate) SetStartAt(t time.Time) *CommunityCreate {
	cc.mutation.SetStartAt(t)
	return cc
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableStartAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetStartAt(*t)
	}
	return cc
}

// SetEndAt sets the "end_at" field.
func (cc *CommunityCreate) SetEndAt(t time.Time) *CommunityCreate {
	cc.mutation.SetEndAt(t)
	return cc
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableEndAt(t *time.Time) *CommunityCreate {
	if t != nil {
		cc.SetEndAt(*t)
	}
	return cc
}

// SetPic sets the "pic" field.
func (cc *CommunityCreate) SetPic(s string) *CommunityCreate {
	cc.mutation.SetPic(s)
	return cc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (cc *CommunityCreate) SetNillablePic(s *string) *CommunityCreate {
	if s != nil {
		cc.SetPic(*s)
	}
	return cc
}

// SetSponsor sets the "sponsor" field.
func (cc *CommunityCreate) SetSponsor(s string) *CommunityCreate {
	cc.mutation.SetSponsor(s)
	return cc
}

// SetNillableSponsor sets the "sponsor" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableSponsor(s *string) *CommunityCreate {
	if s != nil {
		cc.SetSponsor(*s)
	}
	return cc
}

// SetFee sets the "fee" field.
func (cc *CommunityCreate) SetFee(f float64) *CommunityCreate {
	cc.mutation.SetFee(f)
	return cc
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableFee(f *float64) *CommunityCreate {
	if f != nil {
		cc.SetFee(*f)
	}
	return cc
}

// SetIsFee sets the "is_fee" field.
func (cc *CommunityCreate) SetIsFee(i int64) *CommunityCreate {
	cc.mutation.SetIsFee(i)
	return cc
}

// SetNillableIsFee sets the "is_fee" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableIsFee(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetIsFee(*i)
	}
	return cc
}

// SetIsShow sets the "is_show" field.
func (cc *CommunityCreate) SetIsShow(i int64) *CommunityCreate {
	cc.mutation.SetIsShow(i)
	return cc
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableIsShow(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetIsShow(*i)
	}
	return cc
}

// SetIsCancel sets the "is_cancel" field.
func (cc *CommunityCreate) SetIsCancel(i int64) *CommunityCreate {
	cc.mutation.SetIsCancel(i)
	return cc
}

// SetNillableIsCancel sets the "is_cancel" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableIsCancel(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetIsCancel(*i)
	}
	return cc
}

// SetCancelTime sets the "cancel_time" field.
func (cc *CommunityCreate) SetCancelTime(i int64) *CommunityCreate {
	cc.mutation.SetCancelTime(i)
	return cc
}

// SetNillableCancelTime sets the "cancel_time" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableCancelTime(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetCancelTime(*i)
	}
	return cc
}

// SetDetail sets the "detail" field.
func (cc *CommunityCreate) SetDetail(s string) *CommunityCreate {
	cc.mutation.SetDetail(s)
	return cc
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableDetail(s *string) *CommunityCreate {
	if s != nil {
		cc.SetDetail(*s)
	}
	return cc
}

// SetSignFields sets the "sign_fields" field.
func (cc *CommunityCreate) SetSignFields(s string) *CommunityCreate {
	cc.mutation.SetSignFields(s)
	return cc
}

// SetNillableSignFields sets the "sign_fields" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableSignFields(s *string) *CommunityCreate {
	if s != nil {
		cc.SetSignFields(*s)
	}
	return cc
}

// SetCondition sets the "condition" field.
func (cc *CommunityCreate) SetCondition(i int64) *CommunityCreate {
	cc.mutation.SetCondition(i)
	return cc
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (cc *CommunityCreate) SetNillableCondition(i *int64) *CommunityCreate {
	if i != nil {
		cc.SetCondition(*i)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommunityCreate) SetID(i int64) *CommunityCreate {
	cc.mutation.SetID(i)
	return cc
}

// AddCommunityParticipantIDs adds the "community_participants" edge to the CommunityParticipant entity by IDs.
func (cc *CommunityCreate) AddCommunityParticipantIDs(ids ...int64) *CommunityCreate {
	cc.mutation.AddCommunityParticipantIDs(ids...)
	return cc
}

// AddCommunityParticipants adds the "community_participants" edges to the CommunityParticipant entity.
func (cc *CommunityCreate) AddCommunityParticipants(c ...*CommunityParticipant) *CommunityCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cc.AddCommunityParticipantIDs(ids...)
}

// Mutation returns the CommunityMutation object of the builder.
func (cc *CommunityCreate) Mutation() *CommunityMutation {
	return cc.mutation
}

// Save creates the Community in the database.
func (cc *CommunityCreate) Save(ctx context.Context) (*Community, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommunityCreate) SaveX(ctx context.Context) *Community {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommunityCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommunityCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommunityCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := community.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := community.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.Delete(); !ok {
		v := community.DefaultDelete
		cc.mutation.SetDelete(v)
	}
	if _, ok := cc.mutation.CreatedID(); !ok {
		v := community.DefaultCreatedID
		cc.mutation.SetCreatedID(v)
	}
	if _, ok := cc.mutation.Status(); !ok {
		v := community.DefaultStatus
		cc.mutation.SetStatus(v)
	}
	if _, ok := cc.mutation.IsFee(); !ok {
		v := community.DefaultIsFee
		cc.mutation.SetIsFee(v)
	}
	if _, ok := cc.mutation.IsShow(); !ok {
		v := community.DefaultIsShow
		cc.mutation.SetIsShow(v)
	}
	if _, ok := cc.mutation.IsCancel(); !ok {
		v := community.DefaultIsCancel
		cc.mutation.SetIsCancel(v)
	}
	if _, ok := cc.mutation.CancelTime(); !ok {
		v := community.DefaultCancelTime
		cc.mutation.SetCancelTime(v)
	}
	if _, ok := cc.mutation.Condition(); !ok {
		v := community.DefaultCondition
		cc.mutation.SetCondition(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommunityCreate) check() error {
	return nil
}

func (cc *CommunityCreate) sqlSave(ctx context.Context) (*Community, error) {
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

func (cc *CommunityCreate) createSpec() (*Community, *sqlgraph.CreateSpec) {
	var (
		_node = &Community{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(community.Table, sqlgraph.NewFieldSpec(community.FieldID, field.TypeInt64))
	)
	_spec.Schema = cc.schemaConfig.Community
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(community.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(community.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Delete(); ok {
		_spec.SetField(community.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := cc.mutation.CreatedID(); ok {
		_spec.SetField(community.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := cc.mutation.Status(); ok {
		_spec.SetField(community.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(community.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.SignNumber(); ok {
		_spec.SetField(community.FieldSignNumber, field.TypeInt64, value)
		_node.SignNumber = value
	}
	if value, ok := cc.mutation.SignStartAt(); ok {
		_spec.SetField(community.FieldSignStartAt, field.TypeTime, value)
		_node.SignStartAt = value
	}
	if value, ok := cc.mutation.SignEndAt(); ok {
		_spec.SetField(community.FieldSignEndAt, field.TypeTime, value)
		_node.SignEndAt = value
	}
	if value, ok := cc.mutation.Number(); ok {
		_spec.SetField(community.FieldNumber, field.TypeInt64, value)
		_node.Number = value
	}
	if value, ok := cc.mutation.StartAt(); ok {
		_spec.SetField(community.FieldStartAt, field.TypeTime, value)
		_node.StartAt = value
	}
	if value, ok := cc.mutation.EndAt(); ok {
		_spec.SetField(community.FieldEndAt, field.TypeTime, value)
		_node.EndAt = value
	}
	if value, ok := cc.mutation.Pic(); ok {
		_spec.SetField(community.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := cc.mutation.Sponsor(); ok {
		_spec.SetField(community.FieldSponsor, field.TypeString, value)
		_node.Sponsor = value
	}
	if value, ok := cc.mutation.Fee(); ok {
		_spec.SetField(community.FieldFee, field.TypeFloat64, value)
		_node.Fee = value
	}
	if value, ok := cc.mutation.IsFee(); ok {
		_spec.SetField(community.FieldIsFee, field.TypeInt64, value)
		_node.IsFee = value
	}
	if value, ok := cc.mutation.IsShow(); ok {
		_spec.SetField(community.FieldIsShow, field.TypeInt64, value)
		_node.IsShow = value
	}
	if value, ok := cc.mutation.IsCancel(); ok {
		_spec.SetField(community.FieldIsCancel, field.TypeInt64, value)
		_node.IsCancel = value
	}
	if value, ok := cc.mutation.CancelTime(); ok {
		_spec.SetField(community.FieldCancelTime, field.TypeInt64, value)
		_node.CancelTime = value
	}
	if value, ok := cc.mutation.Detail(); ok {
		_spec.SetField(community.FieldDetail, field.TypeString, value)
		_node.Detail = value
	}
	if value, ok := cc.mutation.SignFields(); ok {
		_spec.SetField(community.FieldSignFields, field.TypeString, value)
		_node.SignFields = value
	}
	if value, ok := cc.mutation.Condition(); ok {
		_spec.SetField(community.FieldCondition, field.TypeInt64, value)
		_node.Condition = value
	}
	if nodes := cc.mutation.CommunityParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   community.CommunityParticipantsTable,
			Columns: []string{community.CommunityParticipantsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(communityparticipant.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = cc.schemaConfig.CommunityParticipant
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommunityCreateBulk is the builder for creating many Community entities in bulk.
type CommunityCreateBulk struct {
	config
	err      error
	builders []*CommunityCreate
}

// Save creates the Community entities in the database.
func (ccb *CommunityCreateBulk) Save(ctx context.Context) ([]*Community, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Community, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommunityMutation)
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
func (ccb *CommunityCreateBulk) SaveX(ctx context.Context) []*Community {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommunityCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommunityCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
