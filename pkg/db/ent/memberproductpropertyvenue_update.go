// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/memberproductproperty"
	"saas/pkg/db/ent/memberproductpropertyvenue"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberProductPropertyVenueUpdate is the builder for updating MemberProductPropertyVenue entities.
type MemberProductPropertyVenueUpdate struct {
	config
	hooks    []Hook
	mutation *MemberProductPropertyVenueMutation
}

// Where appends a list predicates to the MemberProductPropertyVenueUpdate builder.
func (mppvu *MemberProductPropertyVenueUpdate) Where(ps ...predicate.MemberProductPropertyVenue) *MemberProductPropertyVenueUpdate {
	mppvu.mutation.Where(ps...)
	return mppvu
}

// SetUpdatedAt sets the "updated_at" field.
func (mppvu *MemberProductPropertyVenueUpdate) SetUpdatedAt(t time.Time) *MemberProductPropertyVenueUpdate {
	mppvu.mutation.SetUpdatedAt(t)
	return mppvu
}

// SetVenueID sets the "venue_id" field.
func (mppvu *MemberProductPropertyVenueUpdate) SetVenueID(i int64) *MemberProductPropertyVenueUpdate {
	mppvu.mutation.ResetVenueID()
	mppvu.mutation.SetVenueID(i)
	return mppvu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mppvu *MemberProductPropertyVenueUpdate) SetNillableVenueID(i *int64) *MemberProductPropertyVenueUpdate {
	if i != nil {
		mppvu.SetVenueID(*i)
	}
	return mppvu
}

// AddVenueID adds i to the "venue_id" field.
func (mppvu *MemberProductPropertyVenueUpdate) AddVenueID(i int64) *MemberProductPropertyVenueUpdate {
	mppvu.mutation.AddVenueID(i)
	return mppvu
}

// ClearVenueID clears the value of the "venue_id" field.
func (mppvu *MemberProductPropertyVenueUpdate) ClearVenueID() *MemberProductPropertyVenueUpdate {
	mppvu.mutation.ClearVenueID()
	return mppvu
}

// SetMemberProductPropertyID sets the "member_product_property_id" field.
func (mppvu *MemberProductPropertyVenueUpdate) SetMemberProductPropertyID(i int64) *MemberProductPropertyVenueUpdate {
	mppvu.mutation.SetMemberProductPropertyID(i)
	return mppvu
}

// SetNillableMemberProductPropertyID sets the "member_product_property_id" field if the given value is not nil.
func (mppvu *MemberProductPropertyVenueUpdate) SetNillableMemberProductPropertyID(i *int64) *MemberProductPropertyVenueUpdate {
	if i != nil {
		mppvu.SetMemberProductPropertyID(*i)
	}
	return mppvu
}

// ClearMemberProductPropertyID clears the value of the "member_product_property_id" field.
func (mppvu *MemberProductPropertyVenueUpdate) ClearMemberProductPropertyID() *MemberProductPropertyVenueUpdate {
	mppvu.mutation.ClearMemberProductPropertyID()
	return mppvu
}

// SetOwnerID sets the "owner" edge to the MemberProductProperty entity by ID.
func (mppvu *MemberProductPropertyVenueUpdate) SetOwnerID(id int64) *MemberProductPropertyVenueUpdate {
	mppvu.mutation.SetOwnerID(id)
	return mppvu
}

// SetNillableOwnerID sets the "owner" edge to the MemberProductProperty entity by ID if the given value is not nil.
func (mppvu *MemberProductPropertyVenueUpdate) SetNillableOwnerID(id *int64) *MemberProductPropertyVenueUpdate {
	if id != nil {
		mppvu = mppvu.SetOwnerID(*id)
	}
	return mppvu
}

// SetOwner sets the "owner" edge to the MemberProductProperty entity.
func (mppvu *MemberProductPropertyVenueUpdate) SetOwner(m *MemberProductProperty) *MemberProductPropertyVenueUpdate {
	return mppvu.SetOwnerID(m.ID)
}

// Mutation returns the MemberProductPropertyVenueMutation object of the builder.
func (mppvu *MemberProductPropertyVenueUpdate) Mutation() *MemberProductPropertyVenueMutation {
	return mppvu.mutation
}

// ClearOwner clears the "owner" edge to the MemberProductProperty entity.
func (mppvu *MemberProductPropertyVenueUpdate) ClearOwner() *MemberProductPropertyVenueUpdate {
	mppvu.mutation.ClearOwner()
	return mppvu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mppvu *MemberProductPropertyVenueUpdate) Save(ctx context.Context) (int, error) {
	mppvu.defaults()
	return withHooks(ctx, mppvu.sqlSave, mppvu.mutation, mppvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mppvu *MemberProductPropertyVenueUpdate) SaveX(ctx context.Context) int {
	affected, err := mppvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mppvu *MemberProductPropertyVenueUpdate) Exec(ctx context.Context) error {
	_, err := mppvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mppvu *MemberProductPropertyVenueUpdate) ExecX(ctx context.Context) {
	if err := mppvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mppvu *MemberProductPropertyVenueUpdate) defaults() {
	if _, ok := mppvu.mutation.UpdatedAt(); !ok {
		v := memberproductpropertyvenue.UpdateDefaultUpdatedAt()
		mppvu.mutation.SetUpdatedAt(v)
	}
}

func (mppvu *MemberProductPropertyVenueUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(memberproductpropertyvenue.Table, memberproductpropertyvenue.Columns, sqlgraph.NewFieldSpec(memberproductpropertyvenue.FieldID, field.TypeInt64))
	if ps := mppvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mppvu.mutation.UpdatedAt(); ok {
		_spec.SetField(memberproductpropertyvenue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mppvu.mutation.VenueID(); ok {
		_spec.SetField(memberproductpropertyvenue.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := mppvu.mutation.AddedVenueID(); ok {
		_spec.AddField(memberproductpropertyvenue.FieldVenueID, field.TypeInt64, value)
	}
	if mppvu.mutation.VenueIDCleared() {
		_spec.ClearField(memberproductpropertyvenue.FieldVenueID, field.TypeInt64)
	}
	if mppvu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductpropertyvenue.OwnerTable,
			Columns: []string{memberproductpropertyvenue.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproductproperty.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mppvu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductpropertyvenue.OwnerTable,
			Columns: []string{memberproductpropertyvenue.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproductproperty.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mppvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberproductpropertyvenue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mppvu.mutation.done = true
	return n, nil
}

// MemberProductPropertyVenueUpdateOne is the builder for updating a single MemberProductPropertyVenue entity.
type MemberProductPropertyVenueUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberProductPropertyVenueMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetUpdatedAt(t time.Time) *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.SetUpdatedAt(t)
	return mppvuo
}

// SetVenueID sets the "venue_id" field.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetVenueID(i int64) *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.ResetVenueID()
	mppvuo.mutation.SetVenueID(i)
	return mppvuo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetNillableVenueID(i *int64) *MemberProductPropertyVenueUpdateOne {
	if i != nil {
		mppvuo.SetVenueID(*i)
	}
	return mppvuo
}

// AddVenueID adds i to the "venue_id" field.
func (mppvuo *MemberProductPropertyVenueUpdateOne) AddVenueID(i int64) *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.AddVenueID(i)
	return mppvuo
}

// ClearVenueID clears the value of the "venue_id" field.
func (mppvuo *MemberProductPropertyVenueUpdateOne) ClearVenueID() *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.ClearVenueID()
	return mppvuo
}

// SetMemberProductPropertyID sets the "member_product_property_id" field.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetMemberProductPropertyID(i int64) *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.SetMemberProductPropertyID(i)
	return mppvuo
}

// SetNillableMemberProductPropertyID sets the "member_product_property_id" field if the given value is not nil.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetNillableMemberProductPropertyID(i *int64) *MemberProductPropertyVenueUpdateOne {
	if i != nil {
		mppvuo.SetMemberProductPropertyID(*i)
	}
	return mppvuo
}

// ClearMemberProductPropertyID clears the value of the "member_product_property_id" field.
func (mppvuo *MemberProductPropertyVenueUpdateOne) ClearMemberProductPropertyID() *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.ClearMemberProductPropertyID()
	return mppvuo
}

// SetOwnerID sets the "owner" edge to the MemberProductProperty entity by ID.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetOwnerID(id int64) *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.SetOwnerID(id)
	return mppvuo
}

// SetNillableOwnerID sets the "owner" edge to the MemberProductProperty entity by ID if the given value is not nil.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetNillableOwnerID(id *int64) *MemberProductPropertyVenueUpdateOne {
	if id != nil {
		mppvuo = mppvuo.SetOwnerID(*id)
	}
	return mppvuo
}

// SetOwner sets the "owner" edge to the MemberProductProperty entity.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SetOwner(m *MemberProductProperty) *MemberProductPropertyVenueUpdateOne {
	return mppvuo.SetOwnerID(m.ID)
}

// Mutation returns the MemberProductPropertyVenueMutation object of the builder.
func (mppvuo *MemberProductPropertyVenueUpdateOne) Mutation() *MemberProductPropertyVenueMutation {
	return mppvuo.mutation
}

// ClearOwner clears the "owner" edge to the MemberProductProperty entity.
func (mppvuo *MemberProductPropertyVenueUpdateOne) ClearOwner() *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.ClearOwner()
	return mppvuo
}

// Where appends a list predicates to the MemberProductPropertyVenueUpdate builder.
func (mppvuo *MemberProductPropertyVenueUpdateOne) Where(ps ...predicate.MemberProductPropertyVenue) *MemberProductPropertyVenueUpdateOne {
	mppvuo.mutation.Where(ps...)
	return mppvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (mppvuo *MemberProductPropertyVenueUpdateOne) Select(field string, fields ...string) *MemberProductPropertyVenueUpdateOne {
	mppvuo.fields = append([]string{field}, fields...)
	return mppvuo
}

// Save executes the query and returns the updated MemberProductPropertyVenue entity.
func (mppvuo *MemberProductPropertyVenueUpdateOne) Save(ctx context.Context) (*MemberProductPropertyVenue, error) {
	mppvuo.defaults()
	return withHooks(ctx, mppvuo.sqlSave, mppvuo.mutation, mppvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mppvuo *MemberProductPropertyVenueUpdateOne) SaveX(ctx context.Context) *MemberProductPropertyVenue {
	node, err := mppvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (mppvuo *MemberProductPropertyVenueUpdateOne) Exec(ctx context.Context) error {
	_, err := mppvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mppvuo *MemberProductPropertyVenueUpdateOne) ExecX(ctx context.Context) {
	if err := mppvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mppvuo *MemberProductPropertyVenueUpdateOne) defaults() {
	if _, ok := mppvuo.mutation.UpdatedAt(); !ok {
		v := memberproductpropertyvenue.UpdateDefaultUpdatedAt()
		mppvuo.mutation.SetUpdatedAt(v)
	}
}

func (mppvuo *MemberProductPropertyVenueUpdateOne) sqlSave(ctx context.Context) (_node *MemberProductPropertyVenue, err error) {
	_spec := sqlgraph.NewUpdateSpec(memberproductpropertyvenue.Table, memberproductpropertyvenue.Columns, sqlgraph.NewFieldSpec(memberproductpropertyvenue.FieldID, field.TypeInt64))
	id, ok := mppvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MemberProductPropertyVenue.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := mppvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, memberproductpropertyvenue.FieldID)
		for _, f := range fields {
			if !memberproductpropertyvenue.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != memberproductpropertyvenue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := mppvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mppvuo.mutation.UpdatedAt(); ok {
		_spec.SetField(memberproductpropertyvenue.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mppvuo.mutation.VenueID(); ok {
		_spec.SetField(memberproductpropertyvenue.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := mppvuo.mutation.AddedVenueID(); ok {
		_spec.AddField(memberproductpropertyvenue.FieldVenueID, field.TypeInt64, value)
	}
	if mppvuo.mutation.VenueIDCleared() {
		_spec.ClearField(memberproductpropertyvenue.FieldVenueID, field.TypeInt64)
	}
	if mppvuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductpropertyvenue.OwnerTable,
			Columns: []string{memberproductpropertyvenue.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproductproperty.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mppvuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   memberproductpropertyvenue.OwnerTable,
			Columns: []string{memberproductpropertyvenue.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberproductproperty.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &MemberProductPropertyVenue{config: mppvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, mppvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{memberproductpropertyvenue.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	mppvuo.mutation.done = true
	return _node, nil
}