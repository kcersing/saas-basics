// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/venue"
	"saas/biz/dal/db/ent/venueplace"
	"saas/idl_gen/model/base"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
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

// ClearUpdatedAt clears the value of the "updated_at" field.
func (vpu *VenuePlaceUpdate) ClearUpdatedAt() *VenuePlaceUpdate {
	vpu.mutation.ClearUpdatedAt()
	return vpu
}

// SetDelete sets the "delete" field.
func (vpu *VenuePlaceUpdate) SetDelete(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetDelete()
	vpu.mutation.SetDelete(i)
	return vpu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableDelete(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetDelete(*i)
	}
	return vpu
}

// AddDelete adds i to the "delete" field.
func (vpu *VenuePlaceUpdate) AddDelete(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddDelete(i)
	return vpu
}

// ClearDelete clears the value of the "delete" field.
func (vpu *VenuePlaceUpdate) ClearDelete() *VenuePlaceUpdate {
	vpu.mutation.ClearDelete()
	return vpu
}

// SetCreatedID sets the "created_id" field.
func (vpu *VenuePlaceUpdate) SetCreatedID(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetCreatedID()
	vpu.mutation.SetCreatedID(i)
	return vpu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableCreatedID(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetCreatedID(*i)
	}
	return vpu
}

// AddCreatedID adds i to the "created_id" field.
func (vpu *VenuePlaceUpdate) AddCreatedID(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddCreatedID(i)
	return vpu
}

// ClearCreatedID clears the value of the "created_id" field.
func (vpu *VenuePlaceUpdate) ClearCreatedID() *VenuePlaceUpdate {
	vpu.mutation.ClearCreatedID()
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

// SetClassify sets the "classify" field.
func (vpu *VenuePlaceUpdate) SetClassify(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetClassify()
	vpu.mutation.SetClassify(i)
	return vpu
}

// SetNillableClassify sets the "classify" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableClassify(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetClassify(*i)
	}
	return vpu
}

// AddClassify adds i to the "classify" field.
func (vpu *VenuePlaceUpdate) AddClassify(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddClassify(i)
	return vpu
}

// ClearClassify clears the value of the "classify" field.
func (vpu *VenuePlaceUpdate) ClearClassify() *VenuePlaceUpdate {
	vpu.mutation.ClearClassify()
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

// SetNumber sets the "number" field.
func (vpu *VenuePlaceUpdate) SetNumber(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetNumber()
	vpu.mutation.SetNumber(i)
	return vpu
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableNumber(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetNumber(*i)
	}
	return vpu
}

// AddNumber adds i to the "number" field.
func (vpu *VenuePlaceUpdate) AddNumber(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddNumber(i)
	return vpu
}

// ClearNumber clears the value of the "number" field.
func (vpu *VenuePlaceUpdate) ClearNumber() *VenuePlaceUpdate {
	vpu.mutation.ClearNumber()
	return vpu
}

// SetIsShow sets the "is_show" field.
func (vpu *VenuePlaceUpdate) SetIsShow(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetIsShow()
	vpu.mutation.SetIsShow(i)
	return vpu
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableIsShow(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetIsShow(*i)
	}
	return vpu
}

// AddIsShow adds i to the "is_show" field.
func (vpu *VenuePlaceUpdate) AddIsShow(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddIsShow(i)
	return vpu
}

// ClearIsShow clears the value of the "is_show" field.
func (vpu *VenuePlaceUpdate) ClearIsShow() *VenuePlaceUpdate {
	vpu.mutation.ClearIsShow()
	return vpu
}

// SetIsAccessible sets the "is_accessible" field.
func (vpu *VenuePlaceUpdate) SetIsAccessible(i int64) *VenuePlaceUpdate {
	vpu.mutation.ResetIsAccessible()
	vpu.mutation.SetIsAccessible(i)
	return vpu
}

// SetNillableIsAccessible sets the "is_accessible" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableIsAccessible(i *int64) *VenuePlaceUpdate {
	if i != nil {
		vpu.SetIsAccessible(*i)
	}
	return vpu
}

// AddIsAccessible adds i to the "is_accessible" field.
func (vpu *VenuePlaceUpdate) AddIsAccessible(i int64) *VenuePlaceUpdate {
	vpu.mutation.AddIsAccessible(i)
	return vpu
}

// ClearIsAccessible clears the value of the "is_accessible" field.
func (vpu *VenuePlaceUpdate) ClearIsAccessible() *VenuePlaceUpdate {
	vpu.mutation.ClearIsAccessible()
	return vpu
}

// SetInformation sets the "information" field.
func (vpu *VenuePlaceUpdate) SetInformation(s string) *VenuePlaceUpdate {
	vpu.mutation.SetInformation(s)
	return vpu
}

// SetNillableInformation sets the "information" field if the given value is not nil.
func (vpu *VenuePlaceUpdate) SetNillableInformation(s *string) *VenuePlaceUpdate {
	if s != nil {
		vpu.SetInformation(*s)
	}
	return vpu
}

// ClearInformation clears the value of the "information" field.
func (vpu *VenuePlaceUpdate) ClearInformation() *VenuePlaceUpdate {
	vpu.mutation.ClearInformation()
	return vpu
}

// SetSeat sets the "seat" field.
func (vpu *VenuePlaceUpdate) SetSeat(b []*base.Seat) *VenuePlaceUpdate {
	vpu.mutation.SetSeat(b)
	return vpu
}

// AppendSeat appends b to the "seat" field.
func (vpu *VenuePlaceUpdate) AppendSeat(b []*base.Seat) *VenuePlaceUpdate {
	vpu.mutation.AppendSeat(b)
	return vpu
}

// ClearSeat clears the value of the "seat" field.
func (vpu *VenuePlaceUpdate) ClearSeat() *VenuePlaceUpdate {
	vpu.mutation.ClearSeat()
	return vpu
}

// SetVenue sets the "venue" edge to the Venue entity.
func (vpu *VenuePlaceUpdate) SetVenue(v *Venue) *VenuePlaceUpdate {
	return vpu.SetVenueID(v.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (vpu *VenuePlaceUpdate) AddProductIDs(ids ...int64) *VenuePlaceUpdate {
	vpu.mutation.AddProductIDs(ids...)
	return vpu
}

// AddProducts adds the "products" edges to the Product entity.
func (vpu *VenuePlaceUpdate) AddProducts(p ...*Product) *VenuePlaceUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vpu.AddProductIDs(ids...)
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

// ClearProducts clears all "products" edges to the Product entity.
func (vpu *VenuePlaceUpdate) ClearProducts() *VenuePlaceUpdate {
	vpu.mutation.ClearProducts()
	return vpu
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (vpu *VenuePlaceUpdate) RemoveProductIDs(ids ...int64) *VenuePlaceUpdate {
	vpu.mutation.RemoveProductIDs(ids...)
	return vpu
}

// RemoveProducts removes "products" edges to Product entities.
func (vpu *VenuePlaceUpdate) RemoveProducts(p ...*Product) *VenuePlaceUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vpu.RemoveProductIDs(ids...)
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
	if _, ok := vpu.mutation.UpdatedAt(); !ok && !vpu.mutation.UpdatedAtCleared() {
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
	if vpu.mutation.CreatedAtCleared() {
		_spec.ClearField(venueplace.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := vpu.mutation.UpdatedAt(); ok {
		_spec.SetField(venueplace.FieldUpdatedAt, field.TypeTime, value)
	}
	if vpu.mutation.UpdatedAtCleared() {
		_spec.ClearField(venueplace.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := vpu.mutation.Delete(); ok {
		_spec.SetField(venueplace.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedDelete(); ok {
		_spec.AddField(venueplace.FieldDelete, field.TypeInt64, value)
	}
	if vpu.mutation.DeleteCleared() {
		_spec.ClearField(venueplace.FieldDelete, field.TypeInt64)
	}
	if value, ok := vpu.mutation.CreatedID(); ok {
		_spec.SetField(venueplace.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedCreatedID(); ok {
		_spec.AddField(venueplace.FieldCreatedID, field.TypeInt64, value)
	}
	if vpu.mutation.CreatedIDCleared() {
		_spec.ClearField(venueplace.FieldCreatedID, field.TypeInt64)
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
	if value, ok := vpu.mutation.Name(); ok {
		_spec.SetField(venueplace.FieldName, field.TypeString, value)
	}
	if vpu.mutation.NameCleared() {
		_spec.ClearField(venueplace.FieldName, field.TypeString)
	}
	if value, ok := vpu.mutation.Classify(); ok {
		_spec.SetField(venueplace.FieldClassify, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedClassify(); ok {
		_spec.AddField(venueplace.FieldClassify, field.TypeInt64, value)
	}
	if vpu.mutation.ClassifyCleared() {
		_spec.ClearField(venueplace.FieldClassify, field.TypeInt64)
	}
	if value, ok := vpu.mutation.Pic(); ok {
		_spec.SetField(venueplace.FieldPic, field.TypeString, value)
	}
	if vpu.mutation.PicCleared() {
		_spec.ClearField(venueplace.FieldPic, field.TypeString)
	}
	if value, ok := vpu.mutation.Number(); ok {
		_spec.SetField(venueplace.FieldNumber, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedNumber(); ok {
		_spec.AddField(venueplace.FieldNumber, field.TypeInt64, value)
	}
	if vpu.mutation.NumberCleared() {
		_spec.ClearField(venueplace.FieldNumber, field.TypeInt64)
	}
	if value, ok := vpu.mutation.IsShow(); ok {
		_spec.SetField(venueplace.FieldIsShow, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedIsShow(); ok {
		_spec.AddField(venueplace.FieldIsShow, field.TypeInt64, value)
	}
	if vpu.mutation.IsShowCleared() {
		_spec.ClearField(venueplace.FieldIsShow, field.TypeInt64)
	}
	if value, ok := vpu.mutation.IsAccessible(); ok {
		_spec.SetField(venueplace.FieldIsAccessible, field.TypeInt64, value)
	}
	if value, ok := vpu.mutation.AddedIsAccessible(); ok {
		_spec.AddField(venueplace.FieldIsAccessible, field.TypeInt64, value)
	}
	if vpu.mutation.IsAccessibleCleared() {
		_spec.ClearField(venueplace.FieldIsAccessible, field.TypeInt64)
	}
	if value, ok := vpu.mutation.Information(); ok {
		_spec.SetField(venueplace.FieldInformation, field.TypeString, value)
	}
	if vpu.mutation.InformationCleared() {
		_spec.ClearField(venueplace.FieldInformation, field.TypeString)
	}
	if value, ok := vpu.mutation.Seat(); ok {
		_spec.SetField(venueplace.FieldSeat, field.TypeJSON, value)
	}
	if value, ok := vpu.mutation.AppendedSeat(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, venueplace.FieldSeat, value)
		})
	}
	if vpu.mutation.SeatCleared() {
		_spec.ClearField(venueplace.FieldSeat, field.TypeJSON)
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
	if vpu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   venueplace.ProductsTable,
			Columns: venueplace.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vpu.mutation.RemovedProductsIDs(); len(nodes) > 0 && !vpu.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   venueplace.ProductsTable,
			Columns: venueplace.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vpu.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   venueplace.ProductsTable,
			Columns: venueplace.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
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

// ClearUpdatedAt clears the value of the "updated_at" field.
func (vpuo *VenuePlaceUpdateOne) ClearUpdatedAt() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearUpdatedAt()
	return vpuo
}

// SetDelete sets the "delete" field.
func (vpuo *VenuePlaceUpdateOne) SetDelete(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetDelete()
	vpuo.mutation.SetDelete(i)
	return vpuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableDelete(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetDelete(*i)
	}
	return vpuo
}

// AddDelete adds i to the "delete" field.
func (vpuo *VenuePlaceUpdateOne) AddDelete(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddDelete(i)
	return vpuo
}

// ClearDelete clears the value of the "delete" field.
func (vpuo *VenuePlaceUpdateOne) ClearDelete() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearDelete()
	return vpuo
}

// SetCreatedID sets the "created_id" field.
func (vpuo *VenuePlaceUpdateOne) SetCreatedID(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetCreatedID()
	vpuo.mutation.SetCreatedID(i)
	return vpuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableCreatedID(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetCreatedID(*i)
	}
	return vpuo
}

// AddCreatedID adds i to the "created_id" field.
func (vpuo *VenuePlaceUpdateOne) AddCreatedID(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddCreatedID(i)
	return vpuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (vpuo *VenuePlaceUpdateOne) ClearCreatedID() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearCreatedID()
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

// SetClassify sets the "classify" field.
func (vpuo *VenuePlaceUpdateOne) SetClassify(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetClassify()
	vpuo.mutation.SetClassify(i)
	return vpuo
}

// SetNillableClassify sets the "classify" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableClassify(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetClassify(*i)
	}
	return vpuo
}

// AddClassify adds i to the "classify" field.
func (vpuo *VenuePlaceUpdateOne) AddClassify(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddClassify(i)
	return vpuo
}

// ClearClassify clears the value of the "classify" field.
func (vpuo *VenuePlaceUpdateOne) ClearClassify() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearClassify()
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

// SetNumber sets the "number" field.
func (vpuo *VenuePlaceUpdateOne) SetNumber(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetNumber()
	vpuo.mutation.SetNumber(i)
	return vpuo
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableNumber(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetNumber(*i)
	}
	return vpuo
}

// AddNumber adds i to the "number" field.
func (vpuo *VenuePlaceUpdateOne) AddNumber(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddNumber(i)
	return vpuo
}

// ClearNumber clears the value of the "number" field.
func (vpuo *VenuePlaceUpdateOne) ClearNumber() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearNumber()
	return vpuo
}

// SetIsShow sets the "is_show" field.
func (vpuo *VenuePlaceUpdateOne) SetIsShow(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetIsShow()
	vpuo.mutation.SetIsShow(i)
	return vpuo
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableIsShow(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetIsShow(*i)
	}
	return vpuo
}

// AddIsShow adds i to the "is_show" field.
func (vpuo *VenuePlaceUpdateOne) AddIsShow(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddIsShow(i)
	return vpuo
}

// ClearIsShow clears the value of the "is_show" field.
func (vpuo *VenuePlaceUpdateOne) ClearIsShow() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearIsShow()
	return vpuo
}

// SetIsAccessible sets the "is_accessible" field.
func (vpuo *VenuePlaceUpdateOne) SetIsAccessible(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.ResetIsAccessible()
	vpuo.mutation.SetIsAccessible(i)
	return vpuo
}

// SetNillableIsAccessible sets the "is_accessible" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableIsAccessible(i *int64) *VenuePlaceUpdateOne {
	if i != nil {
		vpuo.SetIsAccessible(*i)
	}
	return vpuo
}

// AddIsAccessible adds i to the "is_accessible" field.
func (vpuo *VenuePlaceUpdateOne) AddIsAccessible(i int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddIsAccessible(i)
	return vpuo
}

// ClearIsAccessible clears the value of the "is_accessible" field.
func (vpuo *VenuePlaceUpdateOne) ClearIsAccessible() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearIsAccessible()
	return vpuo
}

// SetInformation sets the "information" field.
func (vpuo *VenuePlaceUpdateOne) SetInformation(s string) *VenuePlaceUpdateOne {
	vpuo.mutation.SetInformation(s)
	return vpuo
}

// SetNillableInformation sets the "information" field if the given value is not nil.
func (vpuo *VenuePlaceUpdateOne) SetNillableInformation(s *string) *VenuePlaceUpdateOne {
	if s != nil {
		vpuo.SetInformation(*s)
	}
	return vpuo
}

// ClearInformation clears the value of the "information" field.
func (vpuo *VenuePlaceUpdateOne) ClearInformation() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearInformation()
	return vpuo
}

// SetSeat sets the "seat" field.
func (vpuo *VenuePlaceUpdateOne) SetSeat(b []*base.Seat) *VenuePlaceUpdateOne {
	vpuo.mutation.SetSeat(b)
	return vpuo
}

// AppendSeat appends b to the "seat" field.
func (vpuo *VenuePlaceUpdateOne) AppendSeat(b []*base.Seat) *VenuePlaceUpdateOne {
	vpuo.mutation.AppendSeat(b)
	return vpuo
}

// ClearSeat clears the value of the "seat" field.
func (vpuo *VenuePlaceUpdateOne) ClearSeat() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearSeat()
	return vpuo
}

// SetVenue sets the "venue" edge to the Venue entity.
func (vpuo *VenuePlaceUpdateOne) SetVenue(v *Venue) *VenuePlaceUpdateOne {
	return vpuo.SetVenueID(v.ID)
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (vpuo *VenuePlaceUpdateOne) AddProductIDs(ids ...int64) *VenuePlaceUpdateOne {
	vpuo.mutation.AddProductIDs(ids...)
	return vpuo
}

// AddProducts adds the "products" edges to the Product entity.
func (vpuo *VenuePlaceUpdateOne) AddProducts(p ...*Product) *VenuePlaceUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vpuo.AddProductIDs(ids...)
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

// ClearProducts clears all "products" edges to the Product entity.
func (vpuo *VenuePlaceUpdateOne) ClearProducts() *VenuePlaceUpdateOne {
	vpuo.mutation.ClearProducts()
	return vpuo
}

// RemoveProductIDs removes the "products" edge to Product entities by IDs.
func (vpuo *VenuePlaceUpdateOne) RemoveProductIDs(ids ...int64) *VenuePlaceUpdateOne {
	vpuo.mutation.RemoveProductIDs(ids...)
	return vpuo
}

// RemoveProducts removes "products" edges to Product entities.
func (vpuo *VenuePlaceUpdateOne) RemoveProducts(p ...*Product) *VenuePlaceUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return vpuo.RemoveProductIDs(ids...)
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
	if _, ok := vpuo.mutation.UpdatedAt(); !ok && !vpuo.mutation.UpdatedAtCleared() {
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
	if vpuo.mutation.CreatedAtCleared() {
		_spec.ClearField(venueplace.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := vpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(venueplace.FieldUpdatedAt, field.TypeTime, value)
	}
	if vpuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(venueplace.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := vpuo.mutation.Delete(); ok {
		_spec.SetField(venueplace.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedDelete(); ok {
		_spec.AddField(venueplace.FieldDelete, field.TypeInt64, value)
	}
	if vpuo.mutation.DeleteCleared() {
		_spec.ClearField(venueplace.FieldDelete, field.TypeInt64)
	}
	if value, ok := vpuo.mutation.CreatedID(); ok {
		_spec.SetField(venueplace.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(venueplace.FieldCreatedID, field.TypeInt64, value)
	}
	if vpuo.mutation.CreatedIDCleared() {
		_spec.ClearField(venueplace.FieldCreatedID, field.TypeInt64)
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
	if value, ok := vpuo.mutation.Name(); ok {
		_spec.SetField(venueplace.FieldName, field.TypeString, value)
	}
	if vpuo.mutation.NameCleared() {
		_spec.ClearField(venueplace.FieldName, field.TypeString)
	}
	if value, ok := vpuo.mutation.Classify(); ok {
		_spec.SetField(venueplace.FieldClassify, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedClassify(); ok {
		_spec.AddField(venueplace.FieldClassify, field.TypeInt64, value)
	}
	if vpuo.mutation.ClassifyCleared() {
		_spec.ClearField(venueplace.FieldClassify, field.TypeInt64)
	}
	if value, ok := vpuo.mutation.Pic(); ok {
		_spec.SetField(venueplace.FieldPic, field.TypeString, value)
	}
	if vpuo.mutation.PicCleared() {
		_spec.ClearField(venueplace.FieldPic, field.TypeString)
	}
	if value, ok := vpuo.mutation.Number(); ok {
		_spec.SetField(venueplace.FieldNumber, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedNumber(); ok {
		_spec.AddField(venueplace.FieldNumber, field.TypeInt64, value)
	}
	if vpuo.mutation.NumberCleared() {
		_spec.ClearField(venueplace.FieldNumber, field.TypeInt64)
	}
	if value, ok := vpuo.mutation.IsShow(); ok {
		_spec.SetField(venueplace.FieldIsShow, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedIsShow(); ok {
		_spec.AddField(venueplace.FieldIsShow, field.TypeInt64, value)
	}
	if vpuo.mutation.IsShowCleared() {
		_spec.ClearField(venueplace.FieldIsShow, field.TypeInt64)
	}
	if value, ok := vpuo.mutation.IsAccessible(); ok {
		_spec.SetField(venueplace.FieldIsAccessible, field.TypeInt64, value)
	}
	if value, ok := vpuo.mutation.AddedIsAccessible(); ok {
		_spec.AddField(venueplace.FieldIsAccessible, field.TypeInt64, value)
	}
	if vpuo.mutation.IsAccessibleCleared() {
		_spec.ClearField(venueplace.FieldIsAccessible, field.TypeInt64)
	}
	if value, ok := vpuo.mutation.Information(); ok {
		_spec.SetField(venueplace.FieldInformation, field.TypeString, value)
	}
	if vpuo.mutation.InformationCleared() {
		_spec.ClearField(venueplace.FieldInformation, field.TypeString)
	}
	if value, ok := vpuo.mutation.Seat(); ok {
		_spec.SetField(venueplace.FieldSeat, field.TypeJSON, value)
	}
	if value, ok := vpuo.mutation.AppendedSeat(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, venueplace.FieldSeat, value)
		})
	}
	if vpuo.mutation.SeatCleared() {
		_spec.ClearField(venueplace.FieldSeat, field.TypeJSON)
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
	if vpuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   venueplace.ProductsTable,
			Columns: venueplace.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vpuo.mutation.RemovedProductsIDs(); len(nodes) > 0 && !vpuo.mutation.ProductsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   venueplace.ProductsTable,
			Columns: venueplace.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vpuo.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   venueplace.ProductsTable,
			Columns: venueplace.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
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
