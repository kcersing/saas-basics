// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/member"
	"saas/pkg/db/ent/memberdetails"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberDetailsCreate is the builder for creating a MemberDetails entity.
type MemberDetailsCreate struct {
	config
	mutation *MemberDetailsMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mdc *MemberDetailsCreate) SetCreatedAt(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetCreatedAt(t)
	return mdc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableCreatedAt(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetCreatedAt(*t)
	}
	return mdc
}

// SetUpdatedAt sets the "updated_at" field.
func (mdc *MemberDetailsCreate) SetUpdatedAt(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetUpdatedAt(t)
	return mdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableUpdatedAt(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetUpdatedAt(*t)
	}
	return mdc
}

// SetMemberID sets the "member_id" field.
func (mdc *MemberDetailsCreate) SetMemberID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetMemberID(i)
	return mdc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableMemberID(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetMemberID(*i)
	}
	return mdc
}

// SetNickname sets the "nickname" field.
func (mdc *MemberDetailsCreate) SetNickname(s string) *MemberDetailsCreate {
	mdc.mutation.SetNickname(s)
	return mdc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableNickname(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetNickname(*s)
	}
	return mdc
}

// SetGender sets the "gender" field.
func (mdc *MemberDetailsCreate) SetGender(i int64) *MemberDetailsCreate {
	mdc.mutation.SetGender(i)
	return mdc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableGender(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetGender(*i)
	}
	return mdc
}

// SetBirthday sets the "birthday" field.
func (mdc *MemberDetailsCreate) SetBirthday(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetBirthday(t)
	return mdc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableBirthday(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetBirthday(*t)
	}
	return mdc
}

// SetIdentityCard sets the "identity_card" field.
func (mdc *MemberDetailsCreate) SetIdentityCard(s string) *MemberDetailsCreate {
	mdc.mutation.SetIdentityCard(s)
	return mdc
}

// SetNillableIdentityCard sets the "identity_card" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableIdentityCard(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetIdentityCard(*s)
	}
	return mdc
}

// SetFaceIdentityCard sets the "face_identity_card" field.
func (mdc *MemberDetailsCreate) SetFaceIdentityCard(s string) *MemberDetailsCreate {
	mdc.mutation.SetFaceIdentityCard(s)
	return mdc
}

// SetNillableFaceIdentityCard sets the "face_identity_card" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableFaceIdentityCard(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetFaceIdentityCard(*s)
	}
	return mdc
}

// SetBackIdentityCard sets the "back_identity_card" field.
func (mdc *MemberDetailsCreate) SetBackIdentityCard(s string) *MemberDetailsCreate {
	mdc.mutation.SetBackIdentityCard(s)
	return mdc
}

// SetNillableBackIdentityCard sets the "back_identity_card" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableBackIdentityCard(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetBackIdentityCard(*s)
	}
	return mdc
}

// SetFacePic sets the "face_pic" field.
func (mdc *MemberDetailsCreate) SetFacePic(s string) *MemberDetailsCreate {
	mdc.mutation.SetFacePic(s)
	return mdc
}

// SetNillableFacePic sets the "face_pic" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableFacePic(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetFacePic(*s)
	}
	return mdc
}

// SetFaceEigenvalue sets the "face_eigenvalue" field.
func (mdc *MemberDetailsCreate) SetFaceEigenvalue(s string) *MemberDetailsCreate {
	mdc.mutation.SetFaceEigenvalue(s)
	return mdc
}

// SetNillableFaceEigenvalue sets the "face_eigenvalue" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableFaceEigenvalue(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetFaceEigenvalue(*s)
	}
	return mdc
}

// SetFacePicUpdatedTime sets the "face_pic_updated_time" field.
func (mdc *MemberDetailsCreate) SetFacePicUpdatedTime(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetFacePicUpdatedTime(t)
	return mdc
}

// SetNillableFacePicUpdatedTime sets the "face_pic_updated_time" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableFacePicUpdatedTime(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetFacePicUpdatedTime(*t)
	}
	return mdc
}

// SetMoneySum sets the "money_sum" field.
func (mdc *MemberDetailsCreate) SetMoneySum(f float64) *MemberDetailsCreate {
	mdc.mutation.SetMoneySum(f)
	return mdc
}

// SetNillableMoneySum sets the "money_sum" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableMoneySum(f *float64) *MemberDetailsCreate {
	if f != nil {
		mdc.SetMoneySum(*f)
	}
	return mdc
}

// SetProductID sets the "product_id" field.
func (mdc *MemberDetailsCreate) SetProductID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetProductID(i)
	return mdc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableProductID(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetProductID(*i)
	}
	return mdc
}

// SetProductVenue sets the "product_venue" field.
func (mdc *MemberDetailsCreate) SetProductVenue(i int64) *MemberDetailsCreate {
	mdc.mutation.SetProductVenue(i)
	return mdc
}

// SetNillableProductVenue sets the "product_venue" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableProductVenue(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetProductVenue(*i)
	}
	return mdc
}

// SetEntrySum sets the "entry_sum" field.
func (mdc *MemberDetailsCreate) SetEntrySum(i int64) *MemberDetailsCreate {
	mdc.mutation.SetEntrySum(i)
	return mdc
}

// SetNillableEntrySum sets the "entry_sum" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableEntrySum(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetEntrySum(*i)
	}
	return mdc
}

// SetEntryLastTime sets the "entry_last_time" field.
func (mdc *MemberDetailsCreate) SetEntryLastTime(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetEntryLastTime(t)
	return mdc
}

// SetNillableEntryLastTime sets the "entry_last_time" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableEntryLastTime(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetEntryLastTime(*t)
	}
	return mdc
}

// SetEntryDeadlineTime sets the "entry_deadline_time" field.
func (mdc *MemberDetailsCreate) SetEntryDeadlineTime(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetEntryDeadlineTime(t)
	return mdc
}

// SetNillableEntryDeadlineTime sets the "entry_deadline_time" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableEntryDeadlineTime(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetEntryDeadlineTime(*t)
	}
	return mdc
}

// SetClassLastTime sets the "class_last_time" field.
func (mdc *MemberDetailsCreate) SetClassLastTime(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetClassLastTime(t)
	return mdc
}

// SetNillableClassLastTime sets the "class_last_time" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableClassLastTime(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetClassLastTime(*t)
	}
	return mdc
}

// SetRelationUID sets the "relation_uid" field.
func (mdc *MemberDetailsCreate) SetRelationUID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetRelationUID(i)
	return mdc
}

// SetNillableRelationUID sets the "relation_uid" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableRelationUID(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetRelationUID(*i)
	}
	return mdc
}

// SetRelationMid sets the "relation_mid" field.
func (mdc *MemberDetailsCreate) SetRelationMid(i int64) *MemberDetailsCreate {
	mdc.mutation.SetRelationMid(i)
	return mdc
}

// SetNillableRelationMid sets the "relation_mid" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableRelationMid(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetRelationMid(*i)
	}
	return mdc
}

// SetID sets the "id" field.
func (mdc *MemberDetailsCreate) SetID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetID(i)
	return mdc
}

// SetInfoID sets the "info" edge to the Member entity by ID.
func (mdc *MemberDetailsCreate) SetInfoID(id int64) *MemberDetailsCreate {
	mdc.mutation.SetInfoID(id)
	return mdc
}

// SetNillableInfoID sets the "info" edge to the Member entity by ID if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableInfoID(id *int64) *MemberDetailsCreate {
	if id != nil {
		mdc = mdc.SetInfoID(*id)
	}
	return mdc
}

// SetInfo sets the "info" edge to the Member entity.
func (mdc *MemberDetailsCreate) SetInfo(m *Member) *MemberDetailsCreate {
	return mdc.SetInfoID(m.ID)
}

// Mutation returns the MemberDetailsMutation object of the builder.
func (mdc *MemberDetailsCreate) Mutation() *MemberDetailsMutation {
	return mdc.mutation
}

// Save creates the MemberDetails in the database.
func (mdc *MemberDetailsCreate) Save(ctx context.Context) (*MemberDetails, error) {
	mdc.defaults()
	return withHooks(ctx, mdc.sqlSave, mdc.mutation, mdc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mdc *MemberDetailsCreate) SaveX(ctx context.Context) *MemberDetails {
	v, err := mdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mdc *MemberDetailsCreate) Exec(ctx context.Context) error {
	_, err := mdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdc *MemberDetailsCreate) ExecX(ctx context.Context) {
	if err := mdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mdc *MemberDetailsCreate) defaults() {
	if _, ok := mdc.mutation.CreatedAt(); !ok {
		v := memberdetails.DefaultCreatedAt()
		mdc.mutation.SetCreatedAt(v)
	}
	if _, ok := mdc.mutation.UpdatedAt(); !ok {
		v := memberdetails.DefaultUpdatedAt()
		mdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mdc.mutation.Gender(); !ok {
		v := memberdetails.DefaultGender
		mdc.mutation.SetGender(v)
	}
	if _, ok := mdc.mutation.FaceIdentityCard(); !ok {
		v := memberdetails.DefaultFaceIdentityCard
		mdc.mutation.SetFaceIdentityCard(v)
	}
	if _, ok := mdc.mutation.BackIdentityCard(); !ok {
		v := memberdetails.DefaultBackIdentityCard
		mdc.mutation.SetBackIdentityCard(v)
	}
	if _, ok := mdc.mutation.FacePic(); !ok {
		v := memberdetails.DefaultFacePic
		mdc.mutation.SetFacePic(v)
	}
	if _, ok := mdc.mutation.FaceEigenvalue(); !ok {
		v := memberdetails.DefaultFaceEigenvalue
		mdc.mutation.SetFaceEigenvalue(v)
	}
	if _, ok := mdc.mutation.FacePicUpdatedTime(); !ok {
		v := memberdetails.DefaultFacePicUpdatedTime()
		mdc.mutation.SetFacePicUpdatedTime(v)
	}
	if _, ok := mdc.mutation.MoneySum(); !ok {
		v := memberdetails.DefaultMoneySum
		mdc.mutation.SetMoneySum(v)
	}
	if _, ok := mdc.mutation.ProductID(); !ok {
		v := memberdetails.DefaultProductID
		mdc.mutation.SetProductID(v)
	}
	if _, ok := mdc.mutation.ProductVenue(); !ok {
		v := memberdetails.DefaultProductVenue
		mdc.mutation.SetProductVenue(v)
	}
	if _, ok := mdc.mutation.EntrySum(); !ok {
		v := memberdetails.DefaultEntrySum
		mdc.mutation.SetEntrySum(v)
	}
	if _, ok := mdc.mutation.EntryLastTime(); !ok {
		v := memberdetails.DefaultEntryLastTime()
		mdc.mutation.SetEntryLastTime(v)
	}
	if _, ok := mdc.mutation.EntryDeadlineTime(); !ok {
		v := memberdetails.DefaultEntryDeadlineTime()
		mdc.mutation.SetEntryDeadlineTime(v)
	}
	if _, ok := mdc.mutation.ClassLastTime(); !ok {
		v := memberdetails.DefaultClassLastTime()
		mdc.mutation.SetClassLastTime(v)
	}
	if _, ok := mdc.mutation.RelationUID(); !ok {
		v := memberdetails.DefaultRelationUID
		mdc.mutation.SetRelationUID(v)
	}
	if _, ok := mdc.mutation.RelationMid(); !ok {
		v := memberdetails.DefaultRelationMid
		mdc.mutation.SetRelationMid(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mdc *MemberDetailsCreate) check() error {
	if _, ok := mdc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "MemberDetails.created_at"`)}
	}
	if _, ok := mdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "MemberDetails.updated_at"`)}
	}
	if _, ok := mdc.mutation.FacePicUpdatedTime(); !ok {
		return &ValidationError{Name: "face_pic_updated_time", err: errors.New(`ent: missing required field "MemberDetails.face_pic_updated_time"`)}
	}
	return nil
}

func (mdc *MemberDetailsCreate) sqlSave(ctx context.Context) (*MemberDetails, error) {
	if err := mdc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mdc.mutation.id = &_node.ID
	mdc.mutation.done = true
	return _node, nil
}

func (mdc *MemberDetailsCreate) createSpec() (*MemberDetails, *sqlgraph.CreateSpec) {
	var (
		_node = &MemberDetails{config: mdc.config}
		_spec = sqlgraph.NewCreateSpec(memberdetails.Table, sqlgraph.NewFieldSpec(memberdetails.FieldID, field.TypeInt64))
	)
	if id, ok := mdc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mdc.mutation.CreatedAt(); ok {
		_spec.SetField(memberdetails.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mdc.mutation.UpdatedAt(); ok {
		_spec.SetField(memberdetails.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mdc.mutation.Nickname(); ok {
		_spec.SetField(memberdetails.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := mdc.mutation.Gender(); ok {
		_spec.SetField(memberdetails.FieldGender, field.TypeInt64, value)
		_node.Gender = value
	}
	if value, ok := mdc.mutation.Birthday(); ok {
		_spec.SetField(memberdetails.FieldBirthday, field.TypeTime, value)
		_node.Birthday = value
	}
	if value, ok := mdc.mutation.IdentityCard(); ok {
		_spec.SetField(memberdetails.FieldIdentityCard, field.TypeString, value)
		_node.IdentityCard = value
	}
	if value, ok := mdc.mutation.FaceIdentityCard(); ok {
		_spec.SetField(memberdetails.FieldFaceIdentityCard, field.TypeString, value)
		_node.FaceIdentityCard = value
	}
	if value, ok := mdc.mutation.BackIdentityCard(); ok {
		_spec.SetField(memberdetails.FieldBackIdentityCard, field.TypeString, value)
		_node.BackIdentityCard = value
	}
	if value, ok := mdc.mutation.FacePic(); ok {
		_spec.SetField(memberdetails.FieldFacePic, field.TypeString, value)
		_node.FacePic = value
	}
	if value, ok := mdc.mutation.FaceEigenvalue(); ok {
		_spec.SetField(memberdetails.FieldFaceEigenvalue, field.TypeString, value)
		_node.FaceEigenvalue = value
	}
	if value, ok := mdc.mutation.FacePicUpdatedTime(); ok {
		_spec.SetField(memberdetails.FieldFacePicUpdatedTime, field.TypeTime, value)
		_node.FacePicUpdatedTime = value
	}
	if value, ok := mdc.mutation.MoneySum(); ok {
		_spec.SetField(memberdetails.FieldMoneySum, field.TypeFloat64, value)
		_node.MoneySum = value
	}
	if value, ok := mdc.mutation.ProductID(); ok {
		_spec.SetField(memberdetails.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := mdc.mutation.ProductVenue(); ok {
		_spec.SetField(memberdetails.FieldProductVenue, field.TypeInt64, value)
		_node.ProductVenue = value
	}
	if value, ok := mdc.mutation.EntrySum(); ok {
		_spec.SetField(memberdetails.FieldEntrySum, field.TypeInt64, value)
		_node.EntrySum = value
	}
	if value, ok := mdc.mutation.EntryLastTime(); ok {
		_spec.SetField(memberdetails.FieldEntryLastTime, field.TypeTime, value)
		_node.EntryLastTime = value
	}
	if value, ok := mdc.mutation.EntryDeadlineTime(); ok {
		_spec.SetField(memberdetails.FieldEntryDeadlineTime, field.TypeTime, value)
		_node.EntryDeadlineTime = value
	}
	if value, ok := mdc.mutation.ClassLastTime(); ok {
		_spec.SetField(memberdetails.FieldClassLastTime, field.TypeTime, value)
		_node.ClassLastTime = value
	}
	if value, ok := mdc.mutation.RelationUID(); ok {
		_spec.SetField(memberdetails.FieldRelationUID, field.TypeInt64, value)
		_node.RelationUID = value
	}
	if value, ok := mdc.mutation.RelationMid(); ok {
		_spec.SetField(memberdetails.FieldRelationMid, field.TypeInt64, value)
		_node.RelationMid = value
	}
	if nodes := mdc.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberdetails.InfoTable,
			Columns: []string{memberdetails.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberDetailsCreateBulk is the builder for creating many MemberDetails entities in bulk.
type MemberDetailsCreateBulk struct {
	config
	err      error
	builders []*MemberDetailsCreate
}

// Save creates the MemberDetails entities in the database.
func (mdcb *MemberDetailsCreateBulk) Save(ctx context.Context) ([]*MemberDetails, error) {
	if mdcb.err != nil {
		return nil, mdcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mdcb.builders))
	nodes := make([]*MemberDetails, len(mdcb.builders))
	mutators := make([]Mutator, len(mdcb.builders))
	for i := range mdcb.builders {
		func(i int, root context.Context) {
			builder := mdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberDetailsMutation)
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
					_, err = mutators[i+1].Mutate(root, mdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mdcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mdcb *MemberDetailsCreateBulk) SaveX(ctx context.Context) []*MemberDetails {
	v, err := mdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mdcb *MemberDetailsCreateBulk) Exec(ctx context.Context) error {
	_, err := mdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mdcb *MemberDetailsCreateBulk) ExecX(ctx context.Context) {
	if err := mdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
