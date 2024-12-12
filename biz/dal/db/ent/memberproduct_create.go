// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/membercontract"
	"saas/biz/dal/db/ent/memberproduct"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProductCreate is the builder for creating a MemberProduct entity.
type MemberProductCreate struct {
	config
	mutation *MemberProductMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mpc *MemberProductCreate) SetCreatedAt(t time.Time) *MemberProductCreate {
	mpc.mutation.SetCreatedAt(t)
	return mpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableCreatedAt(t *time.Time) *MemberProductCreate {
	if t != nil {
		mpc.SetCreatedAt(*t)
	}
	return mpc
}

// SetUpdatedAt sets the "updated_at" field.
func (mpc *MemberProductCreate) SetUpdatedAt(t time.Time) *MemberProductCreate {
	mpc.mutation.SetUpdatedAt(t)
	return mpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableUpdatedAt(t *time.Time) *MemberProductCreate {
	if t != nil {
		mpc.SetUpdatedAt(*t)
	}
	return mpc
}

// SetDelete sets the "delete" field.
func (mpc *MemberProductCreate) SetDelete(i int64) *MemberProductCreate {
	mpc.mutation.SetDelete(i)
	return mpc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableDelete(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetDelete(*i)
	}
	return mpc
}

// SetCreatedID sets the "created_id" field.
func (mpc *MemberProductCreate) SetCreatedID(i int64) *MemberProductCreate {
	mpc.mutation.SetCreatedID(i)
	return mpc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableCreatedID(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetCreatedID(*i)
	}
	return mpc
}

// SetStatus sets the "status" field.
func (mpc *MemberProductCreate) SetStatus(i int64) *MemberProductCreate {
	mpc.mutation.SetStatus(i)
	return mpc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableStatus(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetStatus(*i)
	}
	return mpc
}

// SetSn sets the "sn" field.
func (mpc *MemberProductCreate) SetSn(s string) *MemberProductCreate {
	mpc.mutation.SetSn(s)
	return mpc
}

// SetNillableSn sets the "sn" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableSn(s *string) *MemberProductCreate {
	if s != nil {
		mpc.SetSn(*s)
	}
	return mpc
}

// SetType sets the "type" field.
func (mpc *MemberProductCreate) SetType(s string) *MemberProductCreate {
	mpc.mutation.SetType(s)
	return mpc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableType(s *string) *MemberProductCreate {
	if s != nil {
		mpc.SetType(*s)
	}
	return mpc
}

// SetMemberID sets the "member_id" field.
func (mpc *MemberProductCreate) SetMemberID(i int64) *MemberProductCreate {
	mpc.mutation.SetMemberID(i)
	return mpc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableMemberID(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetMemberID(*i)
	}
	return mpc
}

// SetProductID sets the "product_id" field.
func (mpc *MemberProductCreate) SetProductID(i int64) *MemberProductCreate {
	mpc.mutation.SetProductID(i)
	return mpc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableProductID(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetProductID(*i)
	}
	return mpc
}

// SetVenueID sets the "venue_id" field.
func (mpc *MemberProductCreate) SetVenueID(i int64) *MemberProductCreate {
	mpc.mutation.SetVenueID(i)
	return mpc
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableVenueID(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetVenueID(*i)
	}
	return mpc
}

// SetOrderID sets the "order_id" field.
func (mpc *MemberProductCreate) SetOrderID(i int64) *MemberProductCreate {
	mpc.mutation.SetOrderID(i)
	return mpc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableOrderID(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetOrderID(*i)
	}
	return mpc
}

// SetName sets the "name" field.
func (mpc *MemberProductCreate) SetName(s string) *MemberProductCreate {
	mpc.mutation.SetName(s)
	return mpc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableName(s *string) *MemberProductCreate {
	if s != nil {
		mpc.SetName(*s)
	}
	return mpc
}

// SetPrice sets the "price" field.
func (mpc *MemberProductCreate) SetPrice(f float64) *MemberProductCreate {
	mpc.mutation.SetPrice(f)
	return mpc
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillablePrice(f *float64) *MemberProductCreate {
	if f != nil {
		mpc.SetPrice(*f)
	}
	return mpc
}

// SetFee sets the "fee" field.
func (mpc *MemberProductCreate) SetFee(f float64) *MemberProductCreate {
	mpc.mutation.SetFee(f)
	return mpc
}

// SetNillableFee sets the "fee" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableFee(f *float64) *MemberProductCreate {
	if f != nil {
		mpc.SetFee(*f)
	}
	return mpc
}

// SetDuration sets the "duration" field.
func (mpc *MemberProductCreate) SetDuration(i int64) *MemberProductCreate {
	mpc.mutation.SetDuration(i)
	return mpc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableDuration(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetDuration(*i)
	}
	return mpc
}

// SetLength sets the "length" field.
func (mpc *MemberProductCreate) SetLength(i int64) *MemberProductCreate {
	mpc.mutation.SetLength(i)
	return mpc
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableLength(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetLength(*i)
	}
	return mpc
}

// SetCount sets the "count" field.
func (mpc *MemberProductCreate) SetCount(i int64) *MemberProductCreate {
	mpc.mutation.SetCount(i)
	return mpc
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableCount(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetCount(*i)
	}
	return mpc
}

// SetCountSurplus sets the "count_surplus" field.
func (mpc *MemberProductCreate) SetCountSurplus(i int64) *MemberProductCreate {
	mpc.mutation.SetCountSurplus(i)
	return mpc
}

// SetNillableCountSurplus sets the "count_surplus" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableCountSurplus(i *int64) *MemberProductCreate {
	if i != nil {
		mpc.SetCountSurplus(*i)
	}
	return mpc
}

// SetValidityAt sets the "validity_at" field.
func (mpc *MemberProductCreate) SetValidityAt(t time.Time) *MemberProductCreate {
	mpc.mutation.SetValidityAt(t)
	return mpc
}

// SetNillableValidityAt sets the "validity_at" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableValidityAt(t *time.Time) *MemberProductCreate {
	if t != nil {
		mpc.SetValidityAt(*t)
	}
	return mpc
}

// SetCancelAt sets the "cancel_at" field.
func (mpc *MemberProductCreate) SetCancelAt(t time.Time) *MemberProductCreate {
	mpc.mutation.SetCancelAt(t)
	return mpc
}

// SetNillableCancelAt sets the "cancel_at" field if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableCancelAt(t *time.Time) *MemberProductCreate {
	if t != nil {
		mpc.SetCancelAt(*t)
	}
	return mpc
}

// SetID sets the "id" field.
func (mpc *MemberProductCreate) SetID(i int64) *MemberProductCreate {
	mpc.mutation.SetID(i)
	return mpc
}

// SetMembersID sets the "members" edge to the Member entity by ID.
func (mpc *MemberProductCreate) SetMembersID(id int64) *MemberProductCreate {
	mpc.mutation.SetMembersID(id)
	return mpc
}

// SetNillableMembersID sets the "members" edge to the Member entity by ID if the given value is not nil.
func (mpc *MemberProductCreate) SetNillableMembersID(id *int64) *MemberProductCreate {
	if id != nil {
		mpc = mpc.SetMembersID(*id)
	}
	return mpc
}

// SetMembers sets the "members" edge to the Member entity.
func (mpc *MemberProductCreate) SetMembers(m *Member) *MemberProductCreate {
	return mpc.SetMembersID(m.ID)
}

// AddMemberProductEntryIDs adds the "member_product_entry" edge to the EntryLogs entity by IDs.
func (mpc *MemberProductCreate) AddMemberProductEntryIDs(ids ...int64) *MemberProductCreate {
	mpc.mutation.AddMemberProductEntryIDs(ids...)
	return mpc
}

// AddMemberProductEntry adds the "member_product_entry" edges to the EntryLogs entity.
func (mpc *MemberProductCreate) AddMemberProductEntry(e ...*EntryLogs) *MemberProductCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mpc.AddMemberProductEntryIDs(ids...)
}

// AddMemberProductContentIDs adds the "member_product_contents" edge to the MemberContract entity by IDs.
func (mpc *MemberProductCreate) AddMemberProductContentIDs(ids ...int64) *MemberProductCreate {
	mpc.mutation.AddMemberProductContentIDs(ids...)
	return mpc
}

// AddMemberProductContents adds the "member_product_contents" edges to the MemberContract entity.
func (mpc *MemberProductCreate) AddMemberProductContents(m ...*MemberContract) *MemberProductCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mpc.AddMemberProductContentIDs(ids...)
}

// Mutation returns the MemberProductMutation object of the builder.
func (mpc *MemberProductCreate) Mutation() *MemberProductMutation {
	return mpc.mutation
}

// Save creates the MemberProduct in the database.
func (mpc *MemberProductCreate) Save(ctx context.Context) (*MemberProduct, error) {
	mpc.defaults()
	return withHooks(ctx, mpc.sqlSave, mpc.mutation, mpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mpc *MemberProductCreate) SaveX(ctx context.Context) *MemberProduct {
	v, err := mpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpc *MemberProductCreate) Exec(ctx context.Context) error {
	_, err := mpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpc *MemberProductCreate) ExecX(ctx context.Context) {
	if err := mpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mpc *MemberProductCreate) defaults() {
	if _, ok := mpc.mutation.CreatedAt(); !ok {
		v := memberproduct.DefaultCreatedAt()
		mpc.mutation.SetCreatedAt(v)
	}
	if _, ok := mpc.mutation.UpdatedAt(); !ok {
		v := memberproduct.DefaultUpdatedAt()
		mpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mpc.mutation.Delete(); !ok {
		v := memberproduct.DefaultDelete
		mpc.mutation.SetDelete(v)
	}
	if _, ok := mpc.mutation.CreatedID(); !ok {
		v := memberproduct.DefaultCreatedID
		mpc.mutation.SetCreatedID(v)
	}
	if _, ok := mpc.mutation.Status(); !ok {
		v := memberproduct.DefaultStatus
		mpc.mutation.SetStatus(v)
	}
	if _, ok := mpc.mutation.Count(); !ok {
		v := memberproduct.DefaultCount
		mpc.mutation.SetCount(v)
	}
	if _, ok := mpc.mutation.CountSurplus(); !ok {
		v := memberproduct.DefaultCountSurplus
		mpc.mutation.SetCountSurplus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mpc *MemberProductCreate) check() error {
	return nil
}

func (mpc *MemberProductCreate) sqlSave(ctx context.Context) (*MemberProduct, error) {
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

func (mpc *MemberProductCreate) createSpec() (*MemberProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &MemberProduct{config: mpc.config}
		_spec = sqlgraph.NewCreateSpec(memberproduct.Table, sqlgraph.NewFieldSpec(memberproduct.FieldID, field.TypeInt64))
	)
	if id, ok := mpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mpc.mutation.CreatedAt(); ok {
		_spec.SetField(memberproduct.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mpc.mutation.UpdatedAt(); ok {
		_spec.SetField(memberproduct.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mpc.mutation.Delete(); ok {
		_spec.SetField(memberproduct.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := mpc.mutation.CreatedID(); ok {
		_spec.SetField(memberproduct.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := mpc.mutation.Status(); ok {
		_spec.SetField(memberproduct.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := mpc.mutation.Sn(); ok {
		_spec.SetField(memberproduct.FieldSn, field.TypeString, value)
		_node.Sn = value
	}
	if value, ok := mpc.mutation.GetType(); ok {
		_spec.SetField(memberproduct.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := mpc.mutation.ProductID(); ok {
		_spec.SetField(memberproduct.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := mpc.mutation.VenueID(); ok {
		_spec.SetField(memberproduct.FieldVenueID, field.TypeInt64, value)
		_node.VenueID = value
	}
	if value, ok := mpc.mutation.OrderID(); ok {
		_spec.SetField(memberproduct.FieldOrderID, field.TypeInt64, value)
		_node.OrderID = value
	}
	if value, ok := mpc.mutation.Name(); ok {
		_spec.SetField(memberproduct.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mpc.mutation.Price(); ok {
		_spec.SetField(memberproduct.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := mpc.mutation.Fee(); ok {
		_spec.SetField(memberproduct.FieldFee, field.TypeFloat64, value)
		_node.Fee = value
	}
	if value, ok := mpc.mutation.Duration(); ok {
		_spec.SetField(memberproduct.FieldDuration, field.TypeInt64, value)
		_node.Duration = value
	}
	if value, ok := mpc.mutation.Length(); ok {
		_spec.SetField(memberproduct.FieldLength, field.TypeInt64, value)
		_node.Length = value
	}
	if value, ok := mpc.mutation.Count(); ok {
		_spec.SetField(memberproduct.FieldCount, field.TypeInt64, value)
		_node.Count = value
	}
	if value, ok := mpc.mutation.CountSurplus(); ok {
		_spec.SetField(memberproduct.FieldCountSurplus, field.TypeInt64, value)
		_node.CountSurplus = value
	}
	if value, ok := mpc.mutation.ValidityAt(); ok {
		_spec.SetField(memberproduct.FieldValidityAt, field.TypeTime, value)
		_node.ValidityAt = value
	}
	if value, ok := mpc.mutation.CancelAt(); ok {
		_spec.SetField(memberproduct.FieldCancelAt, field.TypeTime, value)
		_node.CancelAt = value
	}
	if nodes := mpc.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproduct.MembersTable,
			Columns: []string{memberproduct.MembersColumn},
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
	if nodes := mpc.mutation.MemberProductEntryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   memberproduct.MemberProductEntryTable,
			Columns: []string{memberproduct.MemberProductEntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mpc.mutation.MemberProductContentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   memberproduct.MemberProductContentsTable,
			Columns: []string{memberproduct.MemberProductContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberProductCreateBulk is the builder for creating many MemberProduct entities in bulk.
type MemberProductCreateBulk struct {
	config
	err      error
	builders []*MemberProductCreate
}

// Save creates the MemberProduct entities in the database.
func (mpcb *MemberProductCreateBulk) Save(ctx context.Context) ([]*MemberProduct, error) {
	if mpcb.err != nil {
		return nil, mpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mpcb.builders))
	nodes := make([]*MemberProduct, len(mpcb.builders))
	mutators := make([]Mutator, len(mpcb.builders))
	for i := range mpcb.builders {
		func(i int, root context.Context) {
			builder := mpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberProductMutation)
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
func (mpcb *MemberProductCreateBulk) SaveX(ctx context.Context) []*MemberProduct {
	v, err := mpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mpcb *MemberProductCreateBulk) Exec(ctx context.Context) error {
	_, err := mpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mpcb *MemberProductCreateBulk) ExecX(ctx context.Context) {
	if err := mpcb.Exec(ctx); err != nil {
		panic(err)
	}
}
