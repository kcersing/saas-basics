// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/product"
	"saas/pkg/db/ent/productproperty"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductUpdate is the builder for updating Product entities.
type ProductUpdate struct {
	config
	hooks    []Hook
	mutation *ProductMutation
}

// Where appends a list predicates to the ProductUpdate builder.
func (pu *ProductUpdate) Where(ps ...predicate.Product) *ProductUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *ProductUpdate) SetUpdatedAt(t time.Time) *ProductUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetSn sets the "sn" field.
func (pu *ProductUpdate) SetSn(s string) *ProductUpdate {
	pu.mutation.SetSn(s)
	return pu
}

// SetNillableSn sets the "sn" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableSn(s *string) *ProductUpdate {
	if s != nil {
		pu.SetSn(*s)
	}
	return pu
}

// ClearSn clears the value of the "sn" field.
func (pu *ProductUpdate) ClearSn() *ProductUpdate {
	pu.mutation.ClearSn()
	return pu
}

// SetVenueID sets the "venue_id" field.
func (pu *ProductUpdate) SetVenueID(i int64) *ProductUpdate {
	pu.mutation.ResetVenueID()
	pu.mutation.SetVenueID(i)
	return pu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableVenueID(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetVenueID(*i)
	}
	return pu
}

// AddVenueID adds i to the "venue_id" field.
func (pu *ProductUpdate) AddVenueID(i int64) *ProductUpdate {
	pu.mutation.AddVenueID(i)
	return pu
}

// ClearVenueID clears the value of the "venue_id" field.
func (pu *ProductUpdate) ClearVenueID() *ProductUpdate {
	pu.mutation.ClearVenueID()
	return pu
}

// SetCreateID sets the "create_id" field.
func (pu *ProductUpdate) SetCreateID(i int64) *ProductUpdate {
	pu.mutation.ResetCreateID()
	pu.mutation.SetCreateID(i)
	return pu
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableCreateID(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetCreateID(*i)
	}
	return pu
}

// AddCreateID adds i to the "create_id" field.
func (pu *ProductUpdate) AddCreateID(i int64) *ProductUpdate {
	pu.mutation.AddCreateID(i)
	return pu
}

// ClearCreateID clears the value of the "create_id" field.
func (pu *ProductUpdate) ClearCreateID() *ProductUpdate {
	pu.mutation.ClearCreateID()
	return pu
}

// SetName sets the "name" field.
func (pu *ProductUpdate) SetName(s string) *ProductUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableName(s *string) *ProductUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// ClearName clears the value of the "name" field.
func (pu *ProductUpdate) ClearName() *ProductUpdate {
	pu.mutation.ClearName()
	return pu
}

// SetPic sets the "pic" field.
func (pu *ProductUpdate) SetPic(i int64) *ProductUpdate {
	pu.mutation.ResetPic()
	pu.mutation.SetPic(i)
	return pu
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (pu *ProductUpdate) SetNillablePic(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetPic(*i)
	}
	return pu
}

// AddPic adds i to the "pic" field.
func (pu *ProductUpdate) AddPic(i int64) *ProductUpdate {
	pu.mutation.AddPic(i)
	return pu
}

// ClearPic clears the value of the "pic" field.
func (pu *ProductUpdate) ClearPic() *ProductUpdate {
	pu.mutation.ClearPic()
	return pu
}

// SetDescription sets the "description" field.
func (pu *ProductUpdate) SetDescription(s string) *ProductUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableDescription(s *string) *ProductUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *ProductUpdate) ClearDescription() *ProductUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetPrice sets the "price" field.
func (pu *ProductUpdate) SetPrice(s string) *ProductUpdate {
	pu.mutation.SetPrice(s)
	return pu
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (pu *ProductUpdate) SetNillablePrice(s *string) *ProductUpdate {
	if s != nil {
		pu.SetPrice(*s)
	}
	return pu
}

// ClearPrice clears the value of the "price" field.
func (pu *ProductUpdate) ClearPrice() *ProductUpdate {
	pu.mutation.ClearPrice()
	return pu
}

// SetStock sets the "stock" field.
func (pu *ProductUpdate) SetStock(i int64) *ProductUpdate {
	pu.mutation.ResetStock()
	pu.mutation.SetStock(i)
	return pu
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableStock(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetStock(*i)
	}
	return pu
}

// AddStock adds i to the "stock" field.
func (pu *ProductUpdate) AddStock(i int64) *ProductUpdate {
	pu.mutation.AddStock(i)
	return pu
}

// ClearStock clears the value of the "stock" field.
func (pu *ProductUpdate) ClearStock() *ProductUpdate {
	pu.mutation.ClearStock()
	return pu
}

// SetStatus sets the "status" field.
func (pu *ProductUpdate) SetStatus(i int64) *ProductUpdate {
	pu.mutation.ResetStatus()
	pu.mutation.SetStatus(i)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *ProductUpdate) SetNillableStatus(i *int64) *ProductUpdate {
	if i != nil {
		pu.SetStatus(*i)
	}
	return pu
}

// AddStatus adds i to the "status" field.
func (pu *ProductUpdate) AddStatus(i int64) *ProductUpdate {
	pu.mutation.AddStatus(i)
	return pu
}

// ClearStatus clears the value of the "status" field.
func (pu *ProductUpdate) ClearStatus() *ProductUpdate {
	pu.mutation.ClearStatus()
	return pu
}

// AddPropertyIDs adds the "propertys" edge to the ProductProperty entity by IDs.
func (pu *ProductUpdate) AddPropertyIDs(ids ...int64) *ProductUpdate {
	pu.mutation.AddPropertyIDs(ids...)
	return pu
}

// AddPropertys adds the "propertys" edges to the ProductProperty entity.
func (pu *ProductUpdate) AddPropertys(p ...*ProductProperty) *ProductUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPropertyIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pu *ProductUpdate) Mutation() *ProductMutation {
	return pu.mutation
}

// ClearPropertys clears all "propertys" edges to the ProductProperty entity.
func (pu *ProductUpdate) ClearPropertys() *ProductUpdate {
	pu.mutation.ClearPropertys()
	return pu
}

// RemovePropertyIDs removes the "propertys" edge to ProductProperty entities by IDs.
func (pu *ProductUpdate) RemovePropertyIDs(ids ...int64) *ProductUpdate {
	pu.mutation.RemovePropertyIDs(ids...)
	return pu
}

// RemovePropertys removes "propertys" edges to ProductProperty entities.
func (pu *ProductUpdate) RemovePropertys(p ...*ProductProperty) *ProductUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePropertyIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProductUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProductUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProductUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProductUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProductUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *ProductUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Sn(); ok {
		_spec.SetField(product.FieldSn, field.TypeString, value)
	}
	if pu.mutation.SnCleared() {
		_spec.ClearField(product.FieldSn, field.TypeString)
	}
	if value, ok := pu.mutation.VenueID(); ok {
		_spec.SetField(product.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedVenueID(); ok {
		_spec.AddField(product.FieldVenueID, field.TypeInt64, value)
	}
	if pu.mutation.VenueIDCleared() {
		_spec.ClearField(product.FieldVenueID, field.TypeInt64)
	}
	if value, ok := pu.mutation.CreateID(); ok {
		_spec.SetField(product.FieldCreateID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedCreateID(); ok {
		_spec.AddField(product.FieldCreateID, field.TypeInt64, value)
	}
	if pu.mutation.CreateIDCleared() {
		_spec.ClearField(product.FieldCreateID, field.TypeInt64)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if pu.mutation.NameCleared() {
		_spec.ClearField(product.FieldName, field.TypeString)
	}
	if value, ok := pu.mutation.Pic(); ok {
		_spec.SetField(product.FieldPic, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedPic(); ok {
		_spec.AddField(product.FieldPic, field.TypeInt64, value)
	}
	if pu.mutation.PicCleared() {
		_spec.ClearField(product.FieldPic, field.TypeInt64)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(product.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeString, value)
	}
	if pu.mutation.PriceCleared() {
		_spec.ClearField(product.FieldPrice, field.TypeString)
	}
	if value, ok := pu.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedStock(); ok {
		_spec.AddField(product.FieldStock, field.TypeInt64, value)
	}
	if pu.mutation.StockCleared() {
		_spec.ClearField(product.FieldStock, field.TypeInt64)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedStatus(); ok {
		_spec.AddField(product.FieldStatus, field.TypeInt64, value)
	}
	if pu.mutation.StatusCleared() {
		_spec.ClearField(product.FieldStatus, field.TypeInt64)
	}
	if pu.mutation.PropertysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.PropertysTable,
			Columns: product.PropertysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPropertysIDs(); len(nodes) > 0 && !pu.mutation.PropertysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.PropertysTable,
			Columns: product.PropertysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PropertysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.PropertysTable,
			Columns: product.PropertysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProductUpdateOne is the builder for updating a single Product entity.
type ProductUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *ProductUpdateOne) SetUpdatedAt(t time.Time) *ProductUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetSn sets the "sn" field.
func (puo *ProductUpdateOne) SetSn(s string) *ProductUpdateOne {
	puo.mutation.SetSn(s)
	return puo
}

// SetNillableSn sets the "sn" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableSn(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetSn(*s)
	}
	return puo
}

// ClearSn clears the value of the "sn" field.
func (puo *ProductUpdateOne) ClearSn() *ProductUpdateOne {
	puo.mutation.ClearSn()
	return puo
}

// SetVenueID sets the "venue_id" field.
func (puo *ProductUpdateOne) SetVenueID(i int64) *ProductUpdateOne {
	puo.mutation.ResetVenueID()
	puo.mutation.SetVenueID(i)
	return puo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableVenueID(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetVenueID(*i)
	}
	return puo
}

// AddVenueID adds i to the "venue_id" field.
func (puo *ProductUpdateOne) AddVenueID(i int64) *ProductUpdateOne {
	puo.mutation.AddVenueID(i)
	return puo
}

// ClearVenueID clears the value of the "venue_id" field.
func (puo *ProductUpdateOne) ClearVenueID() *ProductUpdateOne {
	puo.mutation.ClearVenueID()
	return puo
}

// SetCreateID sets the "create_id" field.
func (puo *ProductUpdateOne) SetCreateID(i int64) *ProductUpdateOne {
	puo.mutation.ResetCreateID()
	puo.mutation.SetCreateID(i)
	return puo
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableCreateID(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetCreateID(*i)
	}
	return puo
}

// AddCreateID adds i to the "create_id" field.
func (puo *ProductUpdateOne) AddCreateID(i int64) *ProductUpdateOne {
	puo.mutation.AddCreateID(i)
	return puo
}

// ClearCreateID clears the value of the "create_id" field.
func (puo *ProductUpdateOne) ClearCreateID() *ProductUpdateOne {
	puo.mutation.ClearCreateID()
	return puo
}

// SetName sets the "name" field.
func (puo *ProductUpdateOne) SetName(s string) *ProductUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableName(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// ClearName clears the value of the "name" field.
func (puo *ProductUpdateOne) ClearName() *ProductUpdateOne {
	puo.mutation.ClearName()
	return puo
}

// SetPic sets the "pic" field.
func (puo *ProductUpdateOne) SetPic(i int64) *ProductUpdateOne {
	puo.mutation.ResetPic()
	puo.mutation.SetPic(i)
	return puo
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillablePic(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetPic(*i)
	}
	return puo
}

// AddPic adds i to the "pic" field.
func (puo *ProductUpdateOne) AddPic(i int64) *ProductUpdateOne {
	puo.mutation.AddPic(i)
	return puo
}

// ClearPic clears the value of the "pic" field.
func (puo *ProductUpdateOne) ClearPic() *ProductUpdateOne {
	puo.mutation.ClearPic()
	return puo
}

// SetDescription sets the "description" field.
func (puo *ProductUpdateOne) SetDescription(s string) *ProductUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableDescription(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *ProductUpdateOne) ClearDescription() *ProductUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetPrice sets the "price" field.
func (puo *ProductUpdateOne) SetPrice(s string) *ProductUpdateOne {
	puo.mutation.SetPrice(s)
	return puo
}

// SetNillablePrice sets the "price" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillablePrice(s *string) *ProductUpdateOne {
	if s != nil {
		puo.SetPrice(*s)
	}
	return puo
}

// ClearPrice clears the value of the "price" field.
func (puo *ProductUpdateOne) ClearPrice() *ProductUpdateOne {
	puo.mutation.ClearPrice()
	return puo
}

// SetStock sets the "stock" field.
func (puo *ProductUpdateOne) SetStock(i int64) *ProductUpdateOne {
	puo.mutation.ResetStock()
	puo.mutation.SetStock(i)
	return puo
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableStock(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetStock(*i)
	}
	return puo
}

// AddStock adds i to the "stock" field.
func (puo *ProductUpdateOne) AddStock(i int64) *ProductUpdateOne {
	puo.mutation.AddStock(i)
	return puo
}

// ClearStock clears the value of the "stock" field.
func (puo *ProductUpdateOne) ClearStock() *ProductUpdateOne {
	puo.mutation.ClearStock()
	return puo
}

// SetStatus sets the "status" field.
func (puo *ProductUpdateOne) SetStatus(i int64) *ProductUpdateOne {
	puo.mutation.ResetStatus()
	puo.mutation.SetStatus(i)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *ProductUpdateOne) SetNillableStatus(i *int64) *ProductUpdateOne {
	if i != nil {
		puo.SetStatus(*i)
	}
	return puo
}

// AddStatus adds i to the "status" field.
func (puo *ProductUpdateOne) AddStatus(i int64) *ProductUpdateOne {
	puo.mutation.AddStatus(i)
	return puo
}

// ClearStatus clears the value of the "status" field.
func (puo *ProductUpdateOne) ClearStatus() *ProductUpdateOne {
	puo.mutation.ClearStatus()
	return puo
}

// AddPropertyIDs adds the "propertys" edge to the ProductProperty entity by IDs.
func (puo *ProductUpdateOne) AddPropertyIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.AddPropertyIDs(ids...)
	return puo
}

// AddPropertys adds the "propertys" edges to the ProductProperty entity.
func (puo *ProductUpdateOne) AddPropertys(p ...*ProductProperty) *ProductUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPropertyIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (puo *ProductUpdateOne) Mutation() *ProductMutation {
	return puo.mutation
}

// ClearPropertys clears all "propertys" edges to the ProductProperty entity.
func (puo *ProductUpdateOne) ClearPropertys() *ProductUpdateOne {
	puo.mutation.ClearPropertys()
	return puo
}

// RemovePropertyIDs removes the "propertys" edge to ProductProperty entities by IDs.
func (puo *ProductUpdateOne) RemovePropertyIDs(ids ...int64) *ProductUpdateOne {
	puo.mutation.RemovePropertyIDs(ids...)
	return puo
}

// RemovePropertys removes "propertys" edges to ProductProperty entities.
func (puo *ProductUpdateOne) RemovePropertys(p ...*ProductProperty) *ProductUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePropertyIDs(ids...)
}

// Where appends a list predicates to the ProductUpdate builder.
func (puo *ProductUpdateOne) Where(ps ...predicate.Product) *ProductUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProductUpdateOne) Select(field string, fields ...string) *ProductUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Product entity.
func (puo *ProductUpdateOne) Save(ctx context.Context) (*Product, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProductUpdateOne) SaveX(ctx context.Context) *Product {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProductUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProductUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProductUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := product.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *ProductUpdateOne) sqlSave(ctx context.Context) (_node *Product, err error) {
	_spec := sqlgraph.NewUpdateSpec(product.Table, product.Columns, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Product.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, product.FieldID)
		for _, f := range fields {
			if !product.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != product.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Sn(); ok {
		_spec.SetField(product.FieldSn, field.TypeString, value)
	}
	if puo.mutation.SnCleared() {
		_spec.ClearField(product.FieldSn, field.TypeString)
	}
	if value, ok := puo.mutation.VenueID(); ok {
		_spec.SetField(product.FieldVenueID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedVenueID(); ok {
		_spec.AddField(product.FieldVenueID, field.TypeInt64, value)
	}
	if puo.mutation.VenueIDCleared() {
		_spec.ClearField(product.FieldVenueID, field.TypeInt64)
	}
	if value, ok := puo.mutation.CreateID(); ok {
		_spec.SetField(product.FieldCreateID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedCreateID(); ok {
		_spec.AddField(product.FieldCreateID, field.TypeInt64, value)
	}
	if puo.mutation.CreateIDCleared() {
		_spec.ClearField(product.FieldCreateID, field.TypeInt64)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
	}
	if puo.mutation.NameCleared() {
		_spec.ClearField(product.FieldName, field.TypeString)
	}
	if value, ok := puo.mutation.Pic(); ok {
		_spec.SetField(product.FieldPic, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedPic(); ok {
		_spec.AddField(product.FieldPic, field.TypeInt64, value)
	}
	if puo.mutation.PicCleared() {
		_spec.ClearField(product.FieldPic, field.TypeInt64)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(product.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Price(); ok {
		_spec.SetField(product.FieldPrice, field.TypeString, value)
	}
	if puo.mutation.PriceCleared() {
		_spec.ClearField(product.FieldPrice, field.TypeString)
	}
	if value, ok := puo.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedStock(); ok {
		_spec.AddField(product.FieldStock, field.TypeInt64, value)
	}
	if puo.mutation.StockCleared() {
		_spec.ClearField(product.FieldStock, field.TypeInt64)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedStatus(); ok {
		_spec.AddField(product.FieldStatus, field.TypeInt64, value)
	}
	if puo.mutation.StatusCleared() {
		_spec.ClearField(product.FieldStatus, field.TypeInt64)
	}
	if puo.mutation.PropertysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.PropertysTable,
			Columns: product.PropertysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPropertysIDs(); len(nodes) > 0 && !puo.mutation.PropertysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.PropertysTable,
			Columns: product.PropertysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PropertysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.PropertysTable,
			Columns: product.PropertysPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Product{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{product.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}