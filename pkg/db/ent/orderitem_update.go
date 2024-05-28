// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/orderitem"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderItemUpdate is the builder for updating OrderItem entities.
type OrderItemUpdate struct {
	config
	hooks    []Hook
	mutation *OrderItemMutation
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiu *OrderItemUpdate) Where(ps ...predicate.OrderItem) *OrderItemUpdate {
	oiu.mutation.Where(ps...)
	return oiu
}

// SetUpdatedAt sets the "updated_at" field.
func (oiu *OrderItemUpdate) SetUpdatedAt(t time.Time) *OrderItemUpdate {
	oiu.mutation.SetUpdatedAt(t)
	return oiu
}

// SetOrderID sets the "order_id" field.
func (oiu *OrderItemUpdate) SetOrderID(i int64) *OrderItemUpdate {
	oiu.mutation.SetOrderID(i)
	return oiu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableOrderID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetOrderID(*i)
	}
	return oiu
}

// ClearOrderID clears the value of the "order_id" field.
func (oiu *OrderItemUpdate) ClearOrderID() *OrderItemUpdate {
	oiu.mutation.ClearOrderID()
	return oiu
}

// SetProductID sets the "product_id" field.
func (oiu *OrderItemUpdate) SetProductID(i int64) *OrderItemUpdate {
	oiu.mutation.ResetProductID()
	oiu.mutation.SetProductID(i)
	return oiu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableProductID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetProductID(*i)
	}
	return oiu
}

// AddProductID adds i to the "product_id" field.
func (oiu *OrderItemUpdate) AddProductID(i int64) *OrderItemUpdate {
	oiu.mutation.AddProductID(i)
	return oiu
}

// ClearProductID clears the value of the "product_id" field.
func (oiu *OrderItemUpdate) ClearProductID() *OrderItemUpdate {
	oiu.mutation.ClearProductID()
	return oiu
}

// SetRelatedUserProductID sets the "related_user_product_id" field.
func (oiu *OrderItemUpdate) SetRelatedUserProductID(i int64) *OrderItemUpdate {
	oiu.mutation.ResetRelatedUserProductID()
	oiu.mutation.SetRelatedUserProductID(i)
	return oiu
}

// SetNillableRelatedUserProductID sets the "related_user_product_id" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableRelatedUserProductID(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetRelatedUserProductID(*i)
	}
	return oiu
}

// AddRelatedUserProductID adds i to the "related_user_product_id" field.
func (oiu *OrderItemUpdate) AddRelatedUserProductID(i int64) *OrderItemUpdate {
	oiu.mutation.AddRelatedUserProductID(i)
	return oiu
}

// ClearRelatedUserProductID clears the value of the "related_user_product_id" field.
func (oiu *OrderItemUpdate) ClearRelatedUserProductID() *OrderItemUpdate {
	oiu.mutation.ClearRelatedUserProductID()
	return oiu
}

// SetData sets the "data" field.
func (oiu *OrderItemUpdate) SetData(i int64) *OrderItemUpdate {
	oiu.mutation.ResetData()
	oiu.mutation.SetData(i)
	return oiu
}

// SetNillableData sets the "data" field if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableData(i *int64) *OrderItemUpdate {
	if i != nil {
		oiu.SetData(*i)
	}
	return oiu
}

// AddData adds i to the "data" field.
func (oiu *OrderItemUpdate) AddData(i int64) *OrderItemUpdate {
	oiu.mutation.AddData(i)
	return oiu
}

// ClearData clears the value of the "data" field.
func (oiu *OrderItemUpdate) ClearData() *OrderItemUpdate {
	oiu.mutation.ClearData()
	return oiu
}

// SetAufkID sets the "aufk" edge to the Order entity by ID.
func (oiu *OrderItemUpdate) SetAufkID(id int64) *OrderItemUpdate {
	oiu.mutation.SetAufkID(id)
	return oiu
}

// SetNillableAufkID sets the "aufk" edge to the Order entity by ID if the given value is not nil.
func (oiu *OrderItemUpdate) SetNillableAufkID(id *int64) *OrderItemUpdate {
	if id != nil {
		oiu = oiu.SetAufkID(*id)
	}
	return oiu
}

// SetAufk sets the "aufk" edge to the Order entity.
func (oiu *OrderItemUpdate) SetAufk(o *Order) *OrderItemUpdate {
	return oiu.SetAufkID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiu *OrderItemUpdate) Mutation() *OrderItemMutation {
	return oiu.mutation
}

// ClearAufk clears the "aufk" edge to the Order entity.
func (oiu *OrderItemUpdate) ClearAufk() *OrderItemUpdate {
	oiu.mutation.ClearAufk()
	return oiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (oiu *OrderItemUpdate) Save(ctx context.Context) (int, error) {
	oiu.defaults()
	return withHooks(ctx, oiu.sqlSave, oiu.mutation, oiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiu *OrderItemUpdate) SaveX(ctx context.Context) int {
	affected, err := oiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (oiu *OrderItemUpdate) Exec(ctx context.Context) error {
	_, err := oiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiu *OrderItemUpdate) ExecX(ctx context.Context) {
	if err := oiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oiu *OrderItemUpdate) defaults() {
	if _, ok := oiu.mutation.UpdatedAt(); !ok {
		v := orderitem.UpdateDefaultUpdatedAt()
		oiu.mutation.SetUpdatedAt(v)
	}
}

func (oiu *OrderItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt64))
	if ps := oiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oiu.mutation.ProductID(); ok {
		_spec.SetField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedProductID(); ok {
		_spec.AddField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if oiu.mutation.ProductIDCleared() {
		_spec.ClearField(orderitem.FieldProductID, field.TypeInt64)
	}
	if value, ok := oiu.mutation.RelatedUserProductID(); ok {
		_spec.SetField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedRelatedUserProductID(); ok {
		_spec.AddField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if oiu.mutation.RelatedUserProductIDCleared() {
		_spec.ClearField(orderitem.FieldRelatedUserProductID, field.TypeInt64)
	}
	if value, ok := oiu.mutation.Data(); ok {
		_spec.SetField(orderitem.FieldData, field.TypeInt64, value)
	}
	if value, ok := oiu.mutation.AddedData(); ok {
		_spec.AddField(orderitem.FieldData, field.TypeInt64, value)
	}
	if oiu.mutation.DataCleared() {
		_spec.ClearField(orderitem.FieldData, field.TypeInt64)
	}
	if oiu.mutation.AufkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.AufkTable,
			Columns: []string{orderitem.AufkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiu.mutation.AufkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.AufkTable,
			Columns: []string{orderitem.AufkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, oiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	oiu.mutation.done = true
	return n, nil
}

// OrderItemUpdateOne is the builder for updating a single OrderItem entity.
type OrderItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderItemMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (oiuo *OrderItemUpdateOne) SetUpdatedAt(t time.Time) *OrderItemUpdateOne {
	oiuo.mutation.SetUpdatedAt(t)
	return oiuo
}

// SetOrderID sets the "order_id" field.
func (oiuo *OrderItemUpdateOne) SetOrderID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.SetOrderID(i)
	return oiuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableOrderID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetOrderID(*i)
	}
	return oiuo
}

// ClearOrderID clears the value of the "order_id" field.
func (oiuo *OrderItemUpdateOne) ClearOrderID() *OrderItemUpdateOne {
	oiuo.mutation.ClearOrderID()
	return oiuo
}

// SetProductID sets the "product_id" field.
func (oiuo *OrderItemUpdateOne) SetProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetProductID()
	oiuo.mutation.SetProductID(i)
	return oiuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableProductID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetProductID(*i)
	}
	return oiuo
}

// AddProductID adds i to the "product_id" field.
func (oiuo *OrderItemUpdateOne) AddProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddProductID(i)
	return oiuo
}

// ClearProductID clears the value of the "product_id" field.
func (oiuo *OrderItemUpdateOne) ClearProductID() *OrderItemUpdateOne {
	oiuo.mutation.ClearProductID()
	return oiuo
}

// SetRelatedUserProductID sets the "related_user_product_id" field.
func (oiuo *OrderItemUpdateOne) SetRelatedUserProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetRelatedUserProductID()
	oiuo.mutation.SetRelatedUserProductID(i)
	return oiuo
}

// SetNillableRelatedUserProductID sets the "related_user_product_id" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableRelatedUserProductID(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetRelatedUserProductID(*i)
	}
	return oiuo
}

// AddRelatedUserProductID adds i to the "related_user_product_id" field.
func (oiuo *OrderItemUpdateOne) AddRelatedUserProductID(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddRelatedUserProductID(i)
	return oiuo
}

// ClearRelatedUserProductID clears the value of the "related_user_product_id" field.
func (oiuo *OrderItemUpdateOne) ClearRelatedUserProductID() *OrderItemUpdateOne {
	oiuo.mutation.ClearRelatedUserProductID()
	return oiuo
}

// SetData sets the "data" field.
func (oiuo *OrderItemUpdateOne) SetData(i int64) *OrderItemUpdateOne {
	oiuo.mutation.ResetData()
	oiuo.mutation.SetData(i)
	return oiuo
}

// SetNillableData sets the "data" field if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableData(i *int64) *OrderItemUpdateOne {
	if i != nil {
		oiuo.SetData(*i)
	}
	return oiuo
}

// AddData adds i to the "data" field.
func (oiuo *OrderItemUpdateOne) AddData(i int64) *OrderItemUpdateOne {
	oiuo.mutation.AddData(i)
	return oiuo
}

// ClearData clears the value of the "data" field.
func (oiuo *OrderItemUpdateOne) ClearData() *OrderItemUpdateOne {
	oiuo.mutation.ClearData()
	return oiuo
}

// SetAufkID sets the "aufk" edge to the Order entity by ID.
func (oiuo *OrderItemUpdateOne) SetAufkID(id int64) *OrderItemUpdateOne {
	oiuo.mutation.SetAufkID(id)
	return oiuo
}

// SetNillableAufkID sets the "aufk" edge to the Order entity by ID if the given value is not nil.
func (oiuo *OrderItemUpdateOne) SetNillableAufkID(id *int64) *OrderItemUpdateOne {
	if id != nil {
		oiuo = oiuo.SetAufkID(*id)
	}
	return oiuo
}

// SetAufk sets the "aufk" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) SetAufk(o *Order) *OrderItemUpdateOne {
	return oiuo.SetAufkID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oiuo *OrderItemUpdateOne) Mutation() *OrderItemMutation {
	return oiuo.mutation
}

// ClearAufk clears the "aufk" edge to the Order entity.
func (oiuo *OrderItemUpdateOne) ClearAufk() *OrderItemUpdateOne {
	oiuo.mutation.ClearAufk()
	return oiuo
}

// Where appends a list predicates to the OrderItemUpdate builder.
func (oiuo *OrderItemUpdateOne) Where(ps ...predicate.OrderItem) *OrderItemUpdateOne {
	oiuo.mutation.Where(ps...)
	return oiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (oiuo *OrderItemUpdateOne) Select(field string, fields ...string) *OrderItemUpdateOne {
	oiuo.fields = append([]string{field}, fields...)
	return oiuo
}

// Save executes the query and returns the updated OrderItem entity.
func (oiuo *OrderItemUpdateOne) Save(ctx context.Context) (*OrderItem, error) {
	oiuo.defaults()
	return withHooks(ctx, oiuo.sqlSave, oiuo.mutation, oiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) SaveX(ctx context.Context) *OrderItem {
	node, err := oiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (oiuo *OrderItemUpdateOne) Exec(ctx context.Context) error {
	_, err := oiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oiuo *OrderItemUpdateOne) ExecX(ctx context.Context) {
	if err := oiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oiuo *OrderItemUpdateOne) defaults() {
	if _, ok := oiuo.mutation.UpdatedAt(); !ok {
		v := orderitem.UpdateDefaultUpdatedAt()
		oiuo.mutation.SetUpdatedAt(v)
	}
}

func (oiuo *OrderItemUpdateOne) sqlSave(ctx context.Context) (_node *OrderItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderitem.Table, orderitem.Columns, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt64))
	id, ok := oiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := oiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderitem.FieldID)
		for _, f := range fields {
			if !orderitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := oiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := oiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := oiuo.mutation.ProductID(); ok {
		_spec.SetField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedProductID(); ok {
		_spec.AddField(orderitem.FieldProductID, field.TypeInt64, value)
	}
	if oiuo.mutation.ProductIDCleared() {
		_spec.ClearField(orderitem.FieldProductID, field.TypeInt64)
	}
	if value, ok := oiuo.mutation.RelatedUserProductID(); ok {
		_spec.SetField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedRelatedUserProductID(); ok {
		_spec.AddField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
	}
	if oiuo.mutation.RelatedUserProductIDCleared() {
		_spec.ClearField(orderitem.FieldRelatedUserProductID, field.TypeInt64)
	}
	if value, ok := oiuo.mutation.Data(); ok {
		_spec.SetField(orderitem.FieldData, field.TypeInt64, value)
	}
	if value, ok := oiuo.mutation.AddedData(); ok {
		_spec.AddField(orderitem.FieldData, field.TypeInt64, value)
	}
	if oiuo.mutation.DataCleared() {
		_spec.ClearField(orderitem.FieldData, field.TypeInt64)
	}
	if oiuo.mutation.AufkCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.AufkTable,
			Columns: []string{orderitem.AufkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := oiuo.mutation.AufkIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.AufkTable,
			Columns: []string{orderitem.AufkColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &OrderItem{config: oiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, oiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	oiuo.mutation.done = true
	return _node, nil
}
