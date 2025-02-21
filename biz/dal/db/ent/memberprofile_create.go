// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberprofile"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProfileCreate is the builder for creating a MemberProfile entity.
type MemberProfileCreate struct {
	config
	mutation *MemberProfileMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mpc *MemberProfileCreate) SetCreatedAt(t time.Time) *MemberProfileCreate {
	mpc.mutation.SetCreatedAt(t)
	return mpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableCreatedAt(t *time.Time) *MemberProfileCreate {
	if t != nil {
		mpc.SetCreatedAt(*t)
	}
	return mpc
}

// SetUpdatedAt sets the "updated_at" field.
func (mpc *MemberProfileCreate) SetUpdatedAt(t time.Time) *MemberProfileCreate {
	mpc.mutation.SetUpdatedAt(t)
	return mpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableUpdatedAt(t *time.Time) *MemberProfileCreate {
	if t != nil {
		mpc.SetUpdatedAt(*t)
	}
	return mpc
}

// SetDelete sets the "delete" field.
func (mpc *MemberProfileCreate) SetDelete(i int64) *MemberProfileCreate {
	mpc.mutation.SetDelete(i)
	return mpc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableDelete(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetDelete(*i)
	}
	return mpc
}

// SetCreatedID sets the "created_id" field.
func (mpc *MemberProfileCreate) SetCreatedID(i int64) *MemberProfileCreate {
	mpc.mutation.SetCreatedID(i)
	return mpc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableCreatedID(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetCreatedID(*i)
	}
	return mpc
}

// SetIntention sets the "intention" field.
func (mpc *MemberProfileCreate) SetIntention(i int64) *MemberProfileCreate {
	mpc.mutation.SetIntention(i)
	return mpc
}

// SetNillableIntention sets the "intention" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableIntention(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetIntention(*i)
	}
	return mpc
}

// SetSource sets the "source" field.
func (mpc *MemberProfileCreate) SetSource(i int64) *MemberProfileCreate {
	mpc.mutation.SetSource(i)
	return mpc
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableSource(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetSource(*i)
	}
	return mpc
}

// SetName sets the "name" field.
func (mpc *MemberProfileCreate) SetName(s string) *MemberProfileCreate {
	mpc.mutation.SetName(s)
	return mpc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableName(s *string) *MemberProfileCreate {
	if s != nil {
		mpc.SetName(*s)
	}
	return mpc
}

// SetMemberID sets the "member_id" field.
func (mpc *MemberProfileCreate) SetMemberID(i int64) *MemberProfileCreate {
	mpc.mutation.SetMemberID(i)
	return mpc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableMemberID(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetMemberID(*i)
	}
	return mpc
}

// SetVenueID sets the "venue_id" field.
func (mpc *MemberProfileCreate) SetVenueID(i int64) *MemberProfileCreate {
	mpc.mutation.SetVenueID(i)
	return mpc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableVenueID(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetVenueID(*i)
	}
	return mpc
}

// SetCondition sets the "condition" field.
func (mpc *MemberProfileCreate) SetCondition(i int64) *MemberProfileCreate {
	mpc.mutation.SetCondition(i)
	return mpc
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableCondition(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetCondition(*i)
	}
	return mpc
}

// SetMobileAscription sets the "mobile_ascription" field.
func (mpc *MemberProfileCreate) SetMobileAscription(i int64) *MemberProfileCreate {
	mpc.mutation.SetMobileAscription(i)
	return mpc
}

// SetNillableMobileAscription sets the "mobile_ascription" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableMobileAscription(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetMobileAscription(*i)
	}
	return mpc
}

// SetFatherName sets the "father_name" field.
func (mpc *MemberProfileCreate) SetFatherName(s string) *MemberProfileCreate {
	mpc.mutation.SetFatherName(s)
	return mpc
}

// SetNillableFatherName sets the "father_name" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableFatherName(s *string) *MemberProfileCreate {
	if s != nil {
		mpc.SetFatherName(*s)
	}
	return mpc
}

// SetMotherName sets the "mother_name" field.
func (mpc *MemberProfileCreate) SetMotherName(s string) *MemberProfileCreate {
	mpc.mutation.SetMotherName(s)
	return mpc
}

// SetNillableMotherName sets the "mother_name" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableMotherName(s *string) *MemberProfileCreate {
	if s != nil {
		mpc.SetMotherName(*s)
	}
	return mpc
}

// SetGender sets the "gender" field.
func (mpc *MemberProfileCreate) SetGender(i int64) *MemberProfileCreate {
	mpc.mutation.SetGender(i)
	return mpc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableGender(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetGender(*i)
	}
	return mpc
}

// SetBirthday sets the "birthday" field.
func (mpc *MemberProfileCreate) SetBirthday(t time.Time) *MemberProfileCreate {
	mpc.mutation.SetBirthday(t)
	return mpc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableBirthday(t *time.Time) *MemberProfileCreate {
	if t != nil {
		mpc.SetBirthday(*t)
	}
	return mpc
}

// SetGrade sets the "grade" field.
func (mpc *MemberProfileCreate) SetGrade(i int64) *MemberProfileCreate {
	mpc.mutation.SetGrade(i)
	return mpc
}

// SetNillableGrade sets the "grade" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableGrade(i *int64) *MemberProfileCreate {
	if i != nil {
		mpc.SetGrade(*i)
	}
	return mpc
}

// SetEmail sets the "email" field.
func (mpc *MemberProfileCreate) SetEmail(s string) *MemberProfileCreate {
	mpc.mutation.SetEmail(s)
	return mpc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableEmail(s *string) *MemberProfileCreate {
	if s != nil {
		mpc.SetEmail(*s)
	}
	return mpc
}

// SetWecom sets the "wecom" field.
func (mpc *MemberProfileCreate) SetWecom(s string) *MemberProfileCreate {
	mpc.mutation.SetWecom(s)
	return mpc
}

// SetNillableWecom sets the "wecom" field if the given value is not nil.
func (mpc *MemberProfileCreate) SetNillableWecom(s *string) *MemberProfileCreate {
	if s != nil {
		mpc.SetWecom(*s)
	}
	return mpc
}

// SetID sets the "id" field.
func (mpc *MemberProfileCreate) SetID(i int64) *MemberProfileCreate {
	mpc.mutation.SetID(i)
	return mpc
}

// SetMember sets the "member" edge to the Member entity.
func (mpc *MemberProfileCreate) SetMember(m *Member) *MemberProfileCreate {
	return mpc.SetMemberID(m.ID)
}

// Mutation returns the MemberProfileMutation object of the builder.
func (mpc *MemberProfileCreate) Mutation() *MemberProfileMutation {
	return mpc.mutation
}

// Save creates the MemberProfile in the database.
func (mpc *MemberProfileCreate) Save(ctx context.Context) (*MemberProfile, error) {
	mpc.defaults()
	return withHooks(ctx, mpc.sqlSave, mpc.mutation, mpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mpc *MemberProfileCreate) SaveX(ctx context.Context) *MemberProfile {
	v, err := mpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpc *MemberProfileCreate) Exec(ctx context.Context) error {
	_, err := mpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpc *MemberProfileCreate) ExecX(ctx context.Context) {
	if err := mpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpc *MemberProfileCreate) defaults() {
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		v := memberprofile.DefaultCreatedAt()
		mpc.mutation.SetCreatedAt(v)
	}
	if _, ok := mpc.mutation.UpdatedAt(); !ok {
		v := memberprofile.DefaultUpdatedAt()
		mpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mpc.mutation.Delete(); !ok {
		v := memberprofile.DefaultDelete
		mpc.mutation.SetDelete(v)
	}
	if _, ok := mpc.mutation.CreatedID(); !ok {
		v := memberprofile.DefaultCreatedID
		mpc.mutation.SetCreatedID(v)
	}
	if _, ok := mpc.mutation.Intention(); !ok {
		v := memberprofile.DefaultIntention
		mpc.mutation.SetIntention(v)
	}
	if _, ok := mpc.mutation.Source(); !ok {
		v := memberprofile.DefaultSource
		mpc.mutation.SetSource(v)
	}
	if _, ok := mpc.mutation.Condition(); !ok {
		v := memberprofile.DefaultCondition
		mpc.mutation.SetCondition(v)
	}
	if _, ok := mpc.mutation.MobileAscription(); !ok {
		v := memberprofile.DefaultMobileAscription
		mpc.mutation.SetMobileAscription(v)
	}
	if _, ok := mpc.mutation.Gender(); !ok {
		v := memberprofile.DefaultGender
		mpc.mutation.SetGender(v)
	}
	if _, ok := mpc.mutation.Grade(); !ok {
		v := memberprofile.DefaultGrade
		mpc.mutation.SetGrade(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpc *MemberProfileCreate) check() error {
	return nil
}

func (mpc *MemberProfileCreate) sqlSave(ctx context.Context) (*MemberProfile, error) {
	if err := mpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mpc.mutation.id = &_node.ID
	mpc.mutation.done = true
	return _node, nil
}

func (mpc *MemberProfileCreate) createSpec() (*MemberProfile, *sqlgraph.CreateSpec) {
	var (
		_node = &MemberProfile{config: mpc.config}
		_spec = sqlgraph.NewCreateSpec(memberprofile.Table, sqlgraph.NewFieldSpec(memberprofile.FieldID, field.TypeInt64))
	)
	_spec.Schema = mpc.schemaConfig.MemberProfile
	if id, ok := mpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mpc.mutation.CreatedAt(); ok {
		_spec.SetField(memberprofile.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mpc.mutation.UpdatedAt(); ok {
		_spec.SetField(memberprofile.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mpc.mutation.Delete(); ok {
		_spec.SetField(memberprofile.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := mpc.mutation.CreatedID(); ok {
		_spec.SetField(memberprofile.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := mpc.mutation.Intention(); ok {
		_spec.SetField(memberprofile.FieldIntention, field.TypeInt64, value)
		_node.Intention = value
	}
	if value, ok := mpc.mutation.Source(); ok {
		_spec.SetField(memberprofile.FieldSource, field.TypeInt64, value)
		_node.Source = value
	}
	if value, ok := mpc.mutation.Name(); ok {
		_spec.SetField(memberprofile.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mpc.mutation.VenueID(); ok {
		_spec.SetField(memberprofile.FieldVenueID, field.TypeInt64, value)
		_node.VenueID = value
	}
	if value, ok := mpc.mutation.Condition(); ok {
		_spec.SetField(memberprofile.FieldCondition, field.TypeInt64, value)
		_node.Condition = value
	}
	if value, ok := mpc.mutation.MobileAscription(); ok {
		_spec.SetField(memberprofile.FieldMobileAscription, field.TypeInt64, value)
		_node.MobileAscription = value
	}
	if value, ok := mpc.mutation.FatherName(); ok {
		_spec.SetField(memberprofile.FieldFatherName, field.TypeString, value)
		_node.FatherName = value
	}
	if value, ok := mpc.mutation.MotherName(); ok {
		_spec.SetField(memberprofile.FieldMotherName, field.TypeString, value)
		_node.MotherName = value
	}
	if value, ok := mpc.mutation.Gender(); ok {
		_spec.SetField(memberprofile.FieldGender, field.TypeInt64, value)
		_node.Gender = value
	}
	if value, ok := mpc.mutation.Birthday(); ok {
		_spec.SetField(memberprofile.FieldBirthday, field.TypeTime, value)
		_node.Birthday = value
	}
	if value, ok := mpc.mutation.Grade(); ok {
		_spec.SetField(memberprofile.FieldGrade, field.TypeInt64, value)
		_node.Grade = value
	}
	if value, ok := mpc.mutation.Email(); ok {
		_spec.SetField(memberprofile.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := mpc.mutation.Wecom(); ok {
		_spec.SetField(memberprofile.FieldWecom, field.TypeString, value)
		_node.Wecom = value
	}
	if nodes := mpc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberprofile.MemberTable,
			Columns: []string{memberprofile.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = mpc.schemaConfig.MemberProfile
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberProfileCreateBulk is the builder for creating many MemberProfile entities in bulk.
type MemberProfileCreateBulk struct {
	config
	err      error
	builders []*MemberProfileCreate
}

// Save creates the MemberProfile entities in the database.
func (mpcb *MemberProfileCreateBulk) Save(ctx context.Context) ([]*MemberProfile, error) {
	if mpcb.err != nil {
		return nil, mpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mpcb.builders))
	nodes := make([]*MemberProfile, len(mpcb.builders))
	mutators := make([]Mutator, len(mpcb.builders))
	for i := range mpcb.builders {
		func(i int, root context.Context) {
			builder := mpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, mpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mpcb *MemberProfileCreateBulk) SaveX(ctx context.Context) []*MemberProfile {
	v, err := mpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpcb *MemberProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := mpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpcb *MemberProfileCreateBulk) ExecX(ctx context.Context) {
	if err := mpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
