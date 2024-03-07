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

// ProductPropertyUpdate is the builder for updating ProductProperty entities.
type ProductPropertyUpdate struct {
	config
	hooks    []Hook
	mutation *ProductPropertyMutation
}

// Where appends a list predicates to the ProductPropertyUpdate builder.
func (ppu *ProductPropertyUpdate) Where(ps ...predicate.ProductProperty) *ProductPropertyUpdate {
	ppu.mutation.Where(ps...)
	return ppu
}

// SetUpdatedAt sets the "updated_at" field.
func (ppu *ProductPropertyUpdate) SetUpdatedAt(t time.Time) *ProductPropertyUpdate {
	ppu.mutation.SetUpdatedAt(t)
	return ppu
}

// SetType sets the "type" field.
func (ppu *ProductPropertyUpdate) SetType(s string) *ProductPropertyUpdate {
	ppu.mutation.SetType(s)
	return ppu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ppu *ProductPropertyUpdate) SetNillableType(s *string) *ProductPropertyUpdate {
	if s != nil {
		ppu.SetType(*s)
	}
	return ppu
}

// ClearType clears the value of the "type" field.
func (ppu *ProductPropertyUpdate) ClearType() *ProductPropertyUpdate {
	ppu.mutation.ClearType()
	return ppu
}

// SetSpuName sets the "spu_name" field.
func (ppu *ProductPropertyUpdate) SetSpuName(s string) *ProductPropertyUpdate {
	ppu.mutation.SetSpuName(s)
	return ppu
}

// SetNillableSpuName sets the "spu_name" field if the given value is not nil.
func (ppu *ProductPropertyUpdate) SetNillableSpuName(s *string) *ProductPropertyUpdate {
	if s != nil {
		ppu.SetSpuName(*s)
	}
	return ppu
}

// ClearSpuName clears the value of the "spu_name" field.
func (ppu *ProductPropertyUpdate) ClearSpuName() *ProductPropertyUpdate {
	ppu.mutation.ClearSpuName()
	return ppu
}

// SetDuration sets the "duration" field.
func (ppu *ProductPropertyUpdate) SetDuration(i int64) *ProductPropertyUpdate {
	ppu.mutation.ResetDuration()
	ppu.mutation.SetDuration(i)
	return ppu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (ppu *ProductPropertyUpdate) SetNillableDuration(i *int64) *ProductPropertyUpdate {
	if i != nil {
		ppu.SetDuration(*i)
	}
	return ppu
}

// AddDuration adds i to the "duration" field.
func (ppu *ProductPropertyUpdate) AddDuration(i int64) *ProductPropertyUpdate {
	ppu.mutation.AddDuration(i)
	return ppu
}

// ClearDuration clears the value of the "duration" field.
func (ppu *ProductPropertyUpdate) ClearDuration() *ProductPropertyUpdate {
	ppu.mutation.ClearDuration()
	return ppu
}

// SetLength sets the "length" field.
func (ppu *ProductPropertyUpdate) SetLength(i int64) *ProductPropertyUpdate {
	ppu.mutation.ResetLength()
	ppu.mutation.SetLength(i)
	return ppu
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (ppu *ProductPropertyUpdate) SetNillableLength(i *int64) *ProductPropertyUpdate {
	if i != nil {
		ppu.SetLength(*i)
	}
	return ppu
}

// AddLength adds i to the "length" field.
func (ppu *ProductPropertyUpdate) AddLength(i int64) *ProductPropertyUpdate {
	ppu.mutation.AddLength(i)
	return ppu
}

// ClearLength clears the value of the "length" field.
func (ppu *ProductPropertyUpdate) ClearLength() *ProductPropertyUpdate {
	ppu.mutation.ClearLength()
	return ppu
}

// SetCount sets the "count" field.
func (ppu *ProductPropertyUpdate) SetCount(i int64) *ProductPropertyUpdate {
	ppu.mutation.ResetCount()
	ppu.mutation.SetCount(i)
	return ppu
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (ppu *ProductPropertyUpdate) SetNillableCount(i *int64) *ProductPropertyUpdate {
	if i != nil {
		ppu.SetCount(*i)
	}
	return ppu
}

// AddCount adds i to the "count" field.
func (ppu *ProductPropertyUpdate) AddCount(i int64) *ProductPropertyUpdate {
	ppu.mutation.AddCount(i)
	return ppu
}

// ClearCount clears the value of the "count" field.
func (ppu *ProductPropertyUpdate) ClearCount() *ProductPropertyUpdate {
	ppu.mutation.ClearCount()
	return ppu
}

// SetSpuPrice sets the "spu_price" field.
func (ppu *ProductPropertyUpdate) SetSpuPrice(f float64) *ProductPropertyUpdate {
	ppu.mutation.ResetSpuPrice()
	ppu.mutation.SetSpuPrice(f)
	return ppu
}

// SetNillableSpuPrice sets the "spu_price" field if the given value is not nil.
func (ppu *ProductPropertyUpdate) SetNillableSpuPrice(f *float64) *ProductPropertyUpdate {
	if f != nil {
		ppu.SetSpuPrice(*f)
	}
	return ppu
}

// AddSpuPrice adds f to the "spu_price" field.
func (ppu *ProductPropertyUpdate) AddSpuPrice(f float64) *ProductPropertyUpdate {
	ppu.mutation.AddSpuPrice(f)
	return ppu
}

// ClearSpuPrice clears the value of the "spu_price" field.
func (ppu *ProductPropertyUpdate) ClearSpuPrice() *ProductPropertyUpdate {
	ppu.mutation.ClearSpuPrice()
	return ppu
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (ppu *ProductPropertyUpdate) AddProductIDs(ids ...int64) *ProductPropertyUpdate {
	ppu.mutation.AddProductIDs(ids...)
	return ppu
}

// AddProduct adds the "product" edges to the Product entity.
func (ppu *ProductPropertyUpdate) AddProduct(p ...*Product) *ProductPropertyUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.AddProductIDs(ids...)
}

// Mutation returns the ProductPropertyMutation object of the builder.
func (ppu *ProductPropertyUpdate) Mutation() *ProductPropertyMutation {
	return ppu.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (ppu *ProductPropertyUpdate) ClearProduct() *ProductPropertyUpdate {
	ppu.mutation.ClearProduct()
	return ppu
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (ppu *ProductPropertyUpdate) RemoveProductIDs(ids ...int64) *ProductPropertyUpdate {
	ppu.mutation.RemoveProductIDs(ids...)
	return ppu
}

// RemoveProduct removes "product" edges to Product entities.
func (ppu *ProductPropertyUpdate) RemoveProduct(p ...*Product) *ProductPropertyUpdate {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppu.RemoveProductIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ppu *ProductPropertyUpdate) Save(ctx context.Context) (int, error) {
	ppu.defaults()
	return withHooks(ctx, ppu.sqlSave, ppu.mutation, ppu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppu *ProductPropertyUpdate) SaveX(ctx context.Context) int {
	affected, err := ppu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ppu *ProductPropertyUpdate) Exec(ctx context.Context) error {
	_, err := ppu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppu *ProductPropertyUpdate) ExecX(ctx context.Context) {
	if err := ppu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppu *ProductPropertyUpdate) defaults() {
	if _, ok := ppu.mutation.UpdatedAt(); !ok {
		v := productproperty.UpdateDefaultUpdatedAt()
		ppu.mutation.SetUpdatedAt(v)
	}
}

func (ppu *ProductPropertyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(productproperty.Table, productproperty.Columns, sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64))
	if ps := ppu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppu.mutation.UpdatedAt(); ok {
		_spec.SetField(productproperty.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ppu.mutation.GetType(); ok {
		_spec.SetField(productproperty.FieldType, field.TypeString, value)
	}
	if ppu.mutation.TypeCleared() {
		_spec.ClearField(productproperty.FieldType, field.TypeString)
	}
	if value, ok := ppu.mutation.SpuName(); ok {
		_spec.SetField(productproperty.FieldSpuName, field.TypeString, value)
	}
	if ppu.mutation.SpuNameCleared() {
		_spec.ClearField(productproperty.FieldSpuName, field.TypeString)
	}
	if value, ok := ppu.mutation.Duration(); ok {
		_spec.SetField(productproperty.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.AddedDuration(); ok {
		_spec.AddField(productproperty.FieldDuration, field.TypeInt64, value)
	}
	if ppu.mutation.DurationCleared() {
		_spec.ClearField(productproperty.FieldDuration, field.TypeInt64)
	}
	if value, ok := ppu.mutation.Length(); ok {
		_spec.SetField(productproperty.FieldLength, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.AddedLength(); ok {
		_spec.AddField(productproperty.FieldLength, field.TypeInt64, value)
	}
	if ppu.mutation.LengthCleared() {
		_spec.ClearField(productproperty.FieldLength, field.TypeInt64)
	}
	if value, ok := ppu.mutation.Count(); ok {
		_spec.SetField(productproperty.FieldCount, field.TypeInt64, value)
	}
	if value, ok := ppu.mutation.AddedCount(); ok {
		_spec.AddField(productproperty.FieldCount, field.TypeInt64, value)
	}
	if ppu.mutation.CountCleared() {
		_spec.ClearField(productproperty.FieldCount, field.TypeInt64)
	}
	if value, ok := ppu.mutation.SpuPrice(); ok {
		_spec.SetField(productproperty.FieldSpuPrice, field.TypeFloat64, value)
	}
	if value, ok := ppu.mutation.AddedSpuPrice(); ok {
		_spec.AddField(productproperty.FieldSpuPrice, field.TypeFloat64, value)
	}
	if ppu.mutation.SpuPriceCleared() {
		_spec.ClearField(productproperty.FieldSpuPrice, field.TypeFloat64)
	}
	if ppu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productproperty.ProductTable,
			Columns: productproperty.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppu.mutation.RemovedProductIDs(); len(nodes) > 0 && !ppu.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productproperty.ProductTable,
			Columns: productproperty.ProductPrimaryKey,
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
	if nodes := ppu.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productproperty.ProductTable,
			Columns: productproperty.ProductPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, ppu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ppu.mutation.done = true
	return n, nil
}

// ProductPropertyUpdateOne is the builder for updating a single ProductProperty entity.
type ProductPropertyUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductPropertyMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ppuo *ProductPropertyUpdateOne) SetUpdatedAt(t time.Time) *ProductPropertyUpdateOne {
	ppuo.mutation.SetUpdatedAt(t)
	return ppuo
}

// SetType sets the "type" field.
func (ppuo *ProductPropertyUpdateOne) SetType(s string) *ProductPropertyUpdateOne {
	ppuo.mutation.SetType(s)
	return ppuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ppuo *ProductPropertyUpdateOne) SetNillableType(s *string) *ProductPropertyUpdateOne {
	if s != nil {
		ppuo.SetType(*s)
	}
	return ppuo
}

// ClearType clears the value of the "type" field.
func (ppuo *ProductPropertyUpdateOne) ClearType() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearType()
	return ppuo
}

// SetSpuName sets the "spu_name" field.
func (ppuo *ProductPropertyUpdateOne) SetSpuName(s string) *ProductPropertyUpdateOne {
	ppuo.mutation.SetSpuName(s)
	return ppuo
}

// SetNillableSpuName sets the "spu_name" field if the given value is not nil.
func (ppuo *ProductPropertyUpdateOne) SetNillableSpuName(s *string) *ProductPropertyUpdateOne {
	if s != nil {
		ppuo.SetSpuName(*s)
	}
	return ppuo
}

// ClearSpuName clears the value of the "spu_name" field.
func (ppuo *ProductPropertyUpdateOne) ClearSpuName() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearSpuName()
	return ppuo
}

// SetDuration sets the "duration" field.
func (ppuo *ProductPropertyUpdateOne) SetDuration(i int64) *ProductPropertyUpdateOne {
	ppuo.mutation.ResetDuration()
	ppuo.mutation.SetDuration(i)
	return ppuo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (ppuo *ProductPropertyUpdateOne) SetNillableDuration(i *int64) *ProductPropertyUpdateOne {
	if i != nil {
		ppuo.SetDuration(*i)
	}
	return ppuo
}

// AddDuration adds i to the "duration" field.
func (ppuo *ProductPropertyUpdateOne) AddDuration(i int64) *ProductPropertyUpdateOne {
	ppuo.mutation.AddDuration(i)
	return ppuo
}

// ClearDuration clears the value of the "duration" field.
func (ppuo *ProductPropertyUpdateOne) ClearDuration() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearDuration()
	return ppuo
}

// SetLength sets the "length" field.
func (ppuo *ProductPropertyUpdateOne) SetLength(i int64) *ProductPropertyUpdateOne {
	ppuo.mutation.ResetLength()
	ppuo.mutation.SetLength(i)
	return ppuo
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (ppuo *ProductPropertyUpdateOne) SetNillableLength(i *int64) *ProductPropertyUpdateOne {
	if i != nil {
		ppuo.SetLength(*i)
	}
	return ppuo
}

// AddLength adds i to the "length" field.
func (ppuo *ProductPropertyUpdateOne) AddLength(i int64) *ProductPropertyUpdateOne {
	ppuo.mutation.AddLength(i)
	return ppuo
}

// ClearLength clears the value of the "length" field.
func (ppuo *ProductPropertyUpdateOne) ClearLength() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearLength()
	return ppuo
}

// SetCount sets the "count" field.
func (ppuo *ProductPropertyUpdateOne) SetCount(i int64) *ProductPropertyUpdateOne {
	ppuo.mutation.ResetCount()
	ppuo.mutation.SetCount(i)
	return ppuo
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (ppuo *ProductPropertyUpdateOne) SetNillableCount(i *int64) *ProductPropertyUpdateOne {
	if i != nil {
		ppuo.SetCount(*i)
	}
	return ppuo
}

// AddCount adds i to the "count" field.
func (ppuo *ProductPropertyUpdateOne) AddCount(i int64) *ProductPropertyUpdateOne {
	ppuo.mutation.AddCount(i)
	return ppuo
}

// ClearCount clears the value of the "count" field.
func (ppuo *ProductPropertyUpdateOne) ClearCount() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearCount()
	return ppuo
}

// SetSpuPrice sets the "spu_price" field.
func (ppuo *ProductPropertyUpdateOne) SetSpuPrice(f float64) *ProductPropertyUpdateOne {
	ppuo.mutation.ResetSpuPrice()
	ppuo.mutation.SetSpuPrice(f)
	return ppuo
}

// SetNillableSpuPrice sets the "spu_price" field if the given value is not nil.
func (ppuo *ProductPropertyUpdateOne) SetNillableSpuPrice(f *float64) *ProductPropertyUpdateOne {
	if f != nil {
		ppuo.SetSpuPrice(*f)
	}
	return ppuo
}

// AddSpuPrice adds f to the "spu_price" field.
func (ppuo *ProductPropertyUpdateOne) AddSpuPrice(f float64) *ProductPropertyUpdateOne {
	ppuo.mutation.AddSpuPrice(f)
	return ppuo
}

// ClearSpuPrice clears the value of the "spu_price" field.
func (ppuo *ProductPropertyUpdateOne) ClearSpuPrice() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearSpuPrice()
	return ppuo
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (ppuo *ProductPropertyUpdateOne) AddProductIDs(ids ...int64) *ProductPropertyUpdateOne {
	ppuo.mutation.AddProductIDs(ids...)
	return ppuo
}

// AddProduct adds the "product" edges to the Product entity.
func (ppuo *ProductPropertyUpdateOne) AddProduct(p ...*Product) *ProductPropertyUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.AddProductIDs(ids...)
}

// Mutation returns the ProductPropertyMutation object of the builder.
func (ppuo *ProductPropertyUpdateOne) Mutation() *ProductPropertyMutation {
	return ppuo.mutation
}

// ClearProduct clears all "product" edges to the Product entity.
func (ppuo *ProductPropertyUpdateOne) ClearProduct() *ProductPropertyUpdateOne {
	ppuo.mutation.ClearProduct()
	return ppuo
}

// RemoveProductIDs removes the "product" edge to Product entities by IDs.
func (ppuo *ProductPropertyUpdateOne) RemoveProductIDs(ids ...int64) *ProductPropertyUpdateOne {
	ppuo.mutation.RemoveProductIDs(ids...)
	return ppuo
}

// RemoveProduct removes "product" edges to Product entities.
func (ppuo *ProductPropertyUpdateOne) RemoveProduct(p ...*Product) *ProductPropertyUpdateOne {
	ids := make([]int64, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ppuo.RemoveProductIDs(ids...)
}

// Where appends a list predicates to the ProductPropertyUpdate builder.
func (ppuo *ProductPropertyUpdateOne) Where(ps ...predicate.ProductProperty) *ProductPropertyUpdateOne {
	ppuo.mutation.Where(ps...)
	return ppuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ppuo *ProductPropertyUpdateOne) Select(field string, fields ...string) *ProductPropertyUpdateOne {
	ppuo.fields = append([]string{field}, fields...)
	return ppuo
}

// Save executes the query and returns the updated ProductProperty entity.
func (ppuo *ProductPropertyUpdateOne) Save(ctx context.Context) (*ProductProperty, error) {
	ppuo.defaults()
	return withHooks(ctx, ppuo.sqlSave, ppuo.mutation, ppuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ppuo *ProductPropertyUpdateOne) SaveX(ctx context.Context) *ProductProperty {
	node, err := ppuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ppuo *ProductPropertyUpdateOne) Exec(ctx context.Context) error {
	_, err := ppuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppuo *ProductPropertyUpdateOne) ExecX(ctx context.Context) {
	if err := ppuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppuo *ProductPropertyUpdateOne) defaults() {
	if _, ok := ppuo.mutation.UpdatedAt(); !ok {
		v := productproperty.UpdateDefaultUpdatedAt()
		ppuo.mutation.SetUpdatedAt(v)
	}
}

func (ppuo *ProductPropertyUpdateOne) sqlSave(ctx context.Context) (_node *ProductProperty, err error) {
	_spec := sqlgraph.NewUpdateSpec(productproperty.Table, productproperty.Columns, sqlgraph.NewFieldSpec(productproperty.FieldID, field.TypeInt64))
	id, ok := ppuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductProperty.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ppuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productproperty.FieldID)
		for _, f := range fields {
			if !productproperty.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productproperty.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ppuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ppuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productproperty.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ppuo.mutation.GetType(); ok {
		_spec.SetField(productproperty.FieldType, field.TypeString, value)
	}
	if ppuo.mutation.TypeCleared() {
		_spec.ClearField(productproperty.FieldType, field.TypeString)
	}
	if value, ok := ppuo.mutation.SpuName(); ok {
		_spec.SetField(productproperty.FieldSpuName, field.TypeString, value)
	}
	if ppuo.mutation.SpuNameCleared() {
		_spec.ClearField(productproperty.FieldSpuName, field.TypeString)
	}
	if value, ok := ppuo.mutation.Duration(); ok {
		_spec.SetField(productproperty.FieldDuration, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.AddedDuration(); ok {
		_spec.AddField(productproperty.FieldDuration, field.TypeInt64, value)
	}
	if ppuo.mutation.DurationCleared() {
		_spec.ClearField(productproperty.FieldDuration, field.TypeInt64)
	}
	if value, ok := ppuo.mutation.Length(); ok {
		_spec.SetField(productproperty.FieldLength, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.AddedLength(); ok {
		_spec.AddField(productproperty.FieldLength, field.TypeInt64, value)
	}
	if ppuo.mutation.LengthCleared() {
		_spec.ClearField(productproperty.FieldLength, field.TypeInt64)
	}
	if value, ok := ppuo.mutation.Count(); ok {
		_spec.SetField(productproperty.FieldCount, field.TypeInt64, value)
	}
	if value, ok := ppuo.mutation.AddedCount(); ok {
		_spec.AddField(productproperty.FieldCount, field.TypeInt64, value)
	}
	if ppuo.mutation.CountCleared() {
		_spec.ClearField(productproperty.FieldCount, field.TypeInt64)
	}
	if value, ok := ppuo.mutation.SpuPrice(); ok {
		_spec.SetField(productproperty.FieldSpuPrice, field.TypeFloat64, value)
	}
	if value, ok := ppuo.mutation.AddedSpuPrice(); ok {
		_spec.AddField(productproperty.FieldSpuPrice, field.TypeFloat64, value)
	}
	if ppuo.mutation.SpuPriceCleared() {
		_spec.ClearField(productproperty.FieldSpuPrice, field.TypeFloat64)
	}
	if ppuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productproperty.ProductTable,
			Columns: productproperty.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ppuo.mutation.RemovedProductIDs(); len(nodes) > 0 && !ppuo.mutation.ProductCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productproperty.ProductTable,
			Columns: productproperty.ProductPrimaryKey,
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
	if nodes := ppuo.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   productproperty.ProductTable,
			Columns: productproperty.ProductPrimaryKey,
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
	_node = &ProductProperty{config: ppuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ppuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productproperty.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ppuo.mutation.done = true
	return _node, nil
}
