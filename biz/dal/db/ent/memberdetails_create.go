// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/memberdetails"
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

// SetDelete sets the "delete" field.
func (mdc *MemberDetailsCreate) SetDelete(i int64) *MemberDetailsCreate {
	mdc.mutation.SetDelete(i)
	return mdc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableDelete(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetDelete(*i)
	}
	return mdc
}

// SetCreatedID sets the "created_id" field.
func (mdc *MemberDetailsCreate) SetCreatedID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetCreatedID(i)
	return mdc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableCreatedID(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetCreatedID(*i)
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

// SetVenueID sets the "venue_id" field.
func (mdc *MemberDetailsCreate) SetVenueID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetVenueID(i)
	return mdc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableVenueID(i *int64) *MemberDetailsCreate {
	if i != nil {
		mdc.SetVenueID(*i)
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

// SetProductName sets the "product_name" field.
func (mdc *MemberDetailsCreate) SetProductName(s string) *MemberDetailsCreate {
	mdc.mutation.SetProductName(s)
	return mdc
}

// SetNillableProductName sets the "product_name" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableProductName(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetProductName(*s)
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

// SetRelationUname sets the "relation_uname" field.
func (mdc *MemberDetailsCreate) SetRelationUname(s string) *MemberDetailsCreate {
	mdc.mutation.SetRelationUname(s)
	return mdc
}

// SetNillableRelationUname sets the "relation_uname" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableRelationUname(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetRelationUname(*s)
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

// SetRelationMame sets the "relation_mame" field.
func (mdc *MemberDetailsCreate) SetRelationMame(s string) *MemberDetailsCreate {
	mdc.mutation.SetRelationMame(s)
	return mdc
}

// SetNillableRelationMame sets the "relation_mame" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableRelationMame(s *string) *MemberDetailsCreate {
	if s != nil {
		mdc.SetRelationMame(*s)
	}
	return mdc
}

// SetFirstTime sets the "first_time" field.
func (mdc *MemberDetailsCreate) SetFirstTime(t time.Time) *MemberDetailsCreate {
	mdc.mutation.SetFirstTime(t)
	return mdc
}

// SetNillableFirstTime sets the "first_time" field if the given value is not nil.
func (mdc *MemberDetailsCreate) SetNillableFirstTime(t *time.Time) *MemberDetailsCreate {
	if t != nil {
		mdc.SetFirstTime(*t)
	}
	return mdc
}

// SetID sets the "id" field.
func (mdc *MemberDetailsCreate) SetID(i int64) *MemberDetailsCreate {
	mdc.mutation.SetID(i)
	return mdc
}

// SetMember sets the "member" edge to the Member entity.
func (mdc *MemberDetailsCreate) SetMember(m *Member) *MemberDetailsCreate {
	return mdc.SetMemberID(m.ID)
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
	if _, ok := mdc.mutation.Delete(); !ok {
		v := memberdetails.DefaultDelete
		mdc.mutation.SetDelete(v)
	}
	if _, ok := mdc.mutation.CreatedID(); !ok {
		v := memberdetails.DefaultCreatedID
		mdc.mutation.SetCreatedID(v)
	}
	if _, ok := mdc.mutation.MoneySum(); !ok {
		v := memberdetails.DefaultMoneySum
		mdc.mutation.SetMoneySum(v)
	}
	if _, ok := mdc.mutation.ProductID(); !ok {
		v := memberdetails.DefaultProductID
		mdc.mutation.SetProductID(v)
	}
	if _, ok := mdc.mutation.EntrySum(); !ok {
		v := memberdetails.DefaultEntrySum
		mdc.mutation.SetEntrySum(v)
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
	_spec.Schema = mdc.schemaConfig.MemberDetails
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
	if value, ok := mdc.mutation.Delete(); ok {
		_spec.SetField(memberdetails.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := mdc.mutation.CreatedID(); ok {
		_spec.SetField(memberdetails.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := mdc.mutation.VenueID(); ok {
		_spec.SetField(memberdetails.FieldVenueID, field.TypeInt64, value)
		_node.VenueID = value
	}
	if value, ok := mdc.mutation.MoneySum(); ok {
		_spec.SetField(memberdetails.FieldMoneySum, field.TypeFloat64, value)
		_node.MoneySum = value
	}
	if value, ok := mdc.mutation.ProductID(); ok {
		_spec.SetField(memberdetails.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := mdc.mutation.ProductName(); ok {
		_spec.SetField(memberdetails.FieldProductName, field.TypeString, value)
		_node.ProductName = value
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
	if value, ok := mdc.mutation.RelationUname(); ok {
		_spec.SetField(memberdetails.FieldRelationUname, field.TypeString, value)
		_node.RelationUname = value
	}
	if value, ok := mdc.mutation.RelationMid(); ok {
		_spec.SetField(memberdetails.FieldRelationMid, field.TypeInt64, value)
		_node.RelationMid = value
	}
	if value, ok := mdc.mutation.RelationMame(); ok {
		_spec.SetField(memberdetails.FieldRelationMame, field.TypeString, value)
		_node.RelationMame = value
	}
	if value, ok := mdc.mutation.FirstTime(); ok {
		_spec.SetField(memberdetails.FieldFirstTime, field.TypeTime, value)
		_node.FirstTime = value
	}
	if nodes := mdc.mutation.MemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberdetails.MemberTable,
			Columns: []string{memberdetails.MemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		edge.Schema = mdc.schemaConfig.MemberDetails
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
