// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/venue"
	"saas/pkg/db/ent/venueplace"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VenuePlaceUpdate is the builder for updating VenuePlace entities.
type VenuePlaceUpdate struct {
	config
	hooks    []Hook
	mutation *VenuePlaceMutation
}

// Where appends a list predicates to the VenuePlaceUpdate builder.
func (vpu *VenuePlaceUpdate) Where(ps ...predicate.VenuePlace) *VenuePlaceUpdate {
	vpu.mutation.Where(ps...)
	return vpu
}

// SetUpdatedAt sets the "updated_at" field.
func (vpu *VenuePlaceUpdate) SetUpdatedAt(t time.Time) *VenuePlaceUpdate {
	vpu.mutation.SetUpdatedAt(t)
	return vpu
}

// SetName sets the "name" field.
func (vpu *VenuePlaceUpdate) SetName(s string) *VenuePlaceUpdate {
	vpu.mutation.SetName(s)
	return vpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableName(s *string) *VenuePlaceUpdate {
	if s != nil {
		vpu.SetName(*s)
	}
	return vpu
}

// ClearName clears the value of the "name" field.
func (vpu *VenuePlaceUpdate) ClearName() *VenuePlaceUpdate {
	vpu.mutation.ClearName()
	return vpu
}

// SetPic sets the "pic" field.
func (vpu *VenuePlaceUpdate) SetPic(s string) *VenuePlaceUpdate {
	vpu.mutation.SetPic(s)
	return vpu
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillablePic(s *string) *VenuePlaceUpdate {
	if s != nil {
		vpu.SetPic(*s)
	}
	return vpu
}

// ClearPic clears the value of the "pic" field.
func (vpu *VenuePlaceUpdate) ClearPic() *VenuePlaceUpdate {
	vpu.mutation.ClearPic()
	return vpu
}

// SetVenueID sets the "venue_id" field.
func (vpu *VenuePlaceUpdate) SetVenueID(i int64) *VenuePlaceUpdate {
	vpu.mutation.SetVenueID(i)
	return vpu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableVenueID(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetVenueID(*i)
	}
	return vpu
}

// ClearVenueID clears the value of the "venue_id" field.
func (vpu *VenuePlaceUpdate) ClearVenueID() *VenuePlaceUpdate {
	vpu.mutation.ClearVenueID()
	return vpu
}

// SetStatus sets the "status" field.
func (vpu *VenuePlaceUpdate) SetStatus(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetStatus()
	vpu.mutation.SetStatus(i)
	return vpu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableStatus(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetStatus(*i)
	}
	return vpu
}

// AddStatus adds i to the "status" field.
func (vpu *VenuePlaceUpdate) AddStatus(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddStatus(i)
	return vpu
}

// ClearStatus clears the value of the "status" field.
func (vpu *VenuePlaceUpdate) ClearStatus() *VenuePlaceUpdate {
	vpu.mutation.ClearStatus()
	return vpu
}

// SetVenue sets the "venue" edge to the Venue entity.
func (vpu *VenuePlaceUpdate) SetVenue(v *Venue) *VenuePlaceUpdate {
	return vpu.SetVenueID(v.ID)
}

// Mutation returns the VenuePlaceMutation object of the builder.
func (vpu *VenuePlaceUpdate) Mutation() *VenuePlaceMutation {
	return vpu.mutation
}

// ClearVenue clears the "venue" edge to the Venue entity.
func (vpu *VenuePlaceUpdate) ClearVenue() *VenuePlaceUpdate {
	vpu.mutation.ClearVenue()
	return vpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vpu *VenuePlaceUpdate) Save(ctx context.Context) (int, error) {
	vpu.defaults()
	return withHooks(ctx, vpu.sqlSave, vpu.mutation, vpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vpu *VenuePlaceUpdate) SaveX(ctx context.Context) int {
	affected, err := vpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vpu *VenuePlaceUpdate) Exec(ctx context.Context) error {
	_, err := vpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vpu *VenuePlaceUpdate) ExecX(ctx context.Context) {
	if err := vpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vpu *VenuePlaceUpdate) defaults() {
	if _, ok := vpu.mutation.UpdatedAt(); !ok {
		v := venueplace.UpdateDefaultUpdatedAt()
		vpu.mutation.SetUpdatedAt(v)
	}
}

func (vpu *VenuePlaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(venueplace.Table, venueplace.Columns, sqlgraph.NewFieldSpec(venueplace.FieldID, field.TypeInt64))
	if ps := vpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vpu.mutation.UpdatedAt(); ok {
		_spec.SetField(venueplace.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vpu.mutation.Name(); ok {
		_spec.SetField(venueplace.FieldName, field.TypeString, value)
	}
	if vpu.mutation.NameCleared() {
		_spec.ClearField(venueplace.FieldName, field.TypeString)
	}
	if value, ok := vpu.mutation.Pic(); ok {
		_spec.SetField(venueplace.FieldPic, field.TypeString, value)
	}
	if vpu.mutation.PicCleared() {
		_spec.ClearField(venueplace.FieldPic, field.TypeString)
	}
	if value, ok := vpu.mutation.Status(); ok {
		_spec.SetField(venueplace.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedStatus(); ok {
		_spec.AddField(venueplace.FieldStatus, field.TypeInt64, value)
	}
	if vpu.mutation.StatusCleared() {
		_spec.ClearField(venueplace.FieldStatus, field.TypeInt64)
	}
	if vpu.mutation.VenueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   venueplace.VenueTable,
			Columns: []string{venueplace.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vpu.mutation.VenueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   venueplace.VenueTable,
			Columns: []string{venueplace.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{venueplace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vpu.mutation.done = true
	return n, nil
}

// VenuePlaceUpdateOne is the builder for updating a single VenuePlace entity.
type VenuePlaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VenuePlaceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (vpuo *VenuePlaceUpdateOne) SetUpdatedAt(t time.Time) *VenuePlaceUpdateOne {
	vpuo.mutation.SetUpdatedAt(t)
	return vpuo
}

// SetName sets the "name" field.
func (vpuo *VenuePlaceUpdateOne) SetName(s string) *VenuePlaceUpdateOne {
	vpuo.mutation.SetName(s)
	return vpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableName(s *string) *VenuePlaceUpdateOne {
	if s != nil {
		vpuo.SetName(*s)
	}
	return vpuo
}

// ClearName clears the value of the "name" field.
func (vpuo *VenuePlaceUpdateOne) ClearName() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearName()
	return vpuo
}

// SetPic sets the "pic" field.
func (vpuo *VenuePlaceUpdateOne) SetPic(s string) *VenuePlaceUpdateOne {
	vpuo.mutation.SetPic(s)
	return vpuo
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillablePic(s *string) *VenuePlaceUpdateOne {
	if s != nil {
		vpuo.SetPic(*s)
	}
	return vpuo
}

// ClearPic clears the value of the "pic" field.
func (vpuo *VenuePlaceUpdateOne) ClearPic() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearPic()
	return vpuo
}

// SetVenueID sets the "venue_id" field.
func (vpuo *VenuePlaceUpdateOne) SetVenueID(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.SetVenueID(i)
	return vpuo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableVenueID(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetVenueID(*i)
	}
	return vpuo
}

// ClearVenueID clears the value of the "venue_id" field.
func (vpuo *VenuePlaceUpdateOne) ClearVenueID() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearVenueID()
	return vpuo
}

// SetStatus sets the "status" field.
func (vpuo *VenuePlaceUpdateOne) SetStatus(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetStatus()
	vpuo.mutation.SetStatus(i)
	return vpuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableStatus(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetStatus(*i)
	}
	return vpuo
}

// AddStatus adds i to the "status" field.
func (vpuo *VenuePlaceUpdateOne) AddStatus(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddStatus(i)
	return vpuo
}

// ClearStatus clears the value of the "status" field.
func (vpuo *VenuePlaceUpdateOne) ClearStatus() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearStatus()
	return vpuo
}

// SetVenue sets the "venue" edge to the Venue entity.
func (vpuo *VenuePlaceUpdateOne) SetVenue(v *Venue) *VenuePlaceUpdateOne {
	return vpuo.SetVenueID(v.ID)
}

// Mutation returns the VenuePlaceMutation object of the builder.
func (vpuo *VenuePlaceUpdateOne) Mutation() *VenuePlaceMutation {
	return vpuo.mutation
}

// ClearVenue clears the "venue" edge to the Venue entity.
func (vpuo *VenuePlaceUpdateOne) ClearVenue() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearVenue()
	return vpuo
}

// Where appends a list predicates to the VenuePlaceUpdate builder.
func (vpuo *VenuePlaceUpdateOne) Where(ps ...predicate.VenuePlace) *VenuePlaceUpdateOne {
	vpuo.mutation.Where(ps...)
	return vpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vpuo *VenuePlaceUpdateOne) Select(field string, fields ...string) *VenuePlaceUpdateOne {
	vpuo.fields = append([]string{field}, fields...)
	return vpuo
}

// Save executes the query and returns the updated VenuePlace entity.
func (vpuo *VenuePlaceUpdateOne) Save(ctx context.Context) (*VenuePlace, error) {
	vpuo.defaults()
	return withHooks(ctx, vpuo.sqlSave, vpuo.mutation, vpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vpuo *VenuePlaceUpdateOne) SaveX(ctx context.Context) *VenuePlace {
	node, err := vpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vpuo *VenuePlaceUpdateOne) Exec(ctx context.Context) error {
	_, err := vpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vpuo *VenuePlaceUpdateOne) ExecX(ctx context.Context) {
	if err := vpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vpuo *VenuePlaceUpdateOne) defaults() {
	if _, ok := vpuo.mutation.UpdatedAt(); !ok {
		v := venueplace.UpdateDefaultUpdatedAt()
		vpuo.mutation.SetUpdatedAt(v)
	}
}

func (vpuo *VenuePlaceUpdateOne) sqlSave(ctx context.Context) (_node *VenuePlace, err error) {
	_spec := sqlgraph.NewUpdateSpec(venueplace.Table, venueplace.Columns, sqlgraph.NewFieldSpec(venueplace.FieldID, field.TypeInt64))
	id, ok := vpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VenuePlace.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, venueplace.FieldID)
		for _, f := range fields {
			if !venueplace.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != venueplace.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(venueplace.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := vpuo.mutation.Name(); ok {
		_spec.SetField(venueplace.FieldName, field.TypeString, value)
	}
	if vpuo.mutation.NameCleared() {
		_spec.ClearField(venueplace.FieldName, field.TypeString)
	}
	if value, ok := vpuo.mutation.Pic(); ok {
		_spec.SetField(venueplace.FieldPic, field.TypeString, value)
	}
	if vpuo.mutation.PicCleared() {
		_spec.ClearField(venueplace.FieldPic, field.TypeString)
	}
	if value, ok := vpuo.mutation.Status(); ok {
		_spec.SetField(venueplace.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedStatus(); ok {
		_spec.AddField(venueplace.FieldStatus, field.TypeInt64, value)
	}
	if vpuo.mutation.StatusCleared() {
		_spec.ClearField(venueplace.FieldStatus, field.TypeInt64)
	}
	if vpuo.mutation.VenueCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   venueplace.VenueTable,
			Columns: []string{venueplace.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vpuo.mutation.VenueIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   venueplace.VenueTable,
			Columns: []string{venueplace.VenueColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &VenuePlace{config: vpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{venueplace.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vpuo.mutation.done = true
	return _node, nil
}
