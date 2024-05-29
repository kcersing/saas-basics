// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/order"
	"saas/pkg/db/ent/orderpay"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderPayUpdate is the builder for updating OrderPay entities.
type OrderPayUpdate struct {
	config
	hooks    []Hook
	mutation *OrderPayMutation
}

// Where appends a list predicates to the OrderPayUpdate builder.
func (opu *OrderPayUpdate) Where(ps ...predicate.OrderPay) *OrderPayUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetUpdatedAt sets the "updated_at" field.
func (opu *OrderPayUpdate) SetUpdatedAt(t time.Time) *OrderPayUpdate {
	opu.mutation.SetUpdatedAt(t)
	return opu
}

// SetOrderID sets the "order_id" field.
func (opu *OrderPayUpdate) SetOrderID(i int64) *OrderPayUpdate {
	opu.mutation.SetOrderID(i)
	return opu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableOrderID(i *int64) *OrderPayUpdate {
	if i != nil {
		opu.SetOrderID(*i)
	}
	return opu
}

// ClearOrderID clears the value of the "order_id" field.
func (opu *OrderPayUpdate) ClearOrderID() *OrderPayUpdate {
	opu.mutation.ClearOrderID()
	return opu
}

// SetPaySn sets the "pay_sn" field.
func (opu *OrderPayUpdate) SetPaySn(s string) *OrderPayUpdate {
	opu.mutation.SetPaySn(s)
	return opu
}

// SetNillablePaySn sets the "pay_sn" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePaySn(s *string) *OrderPayUpdate {
	if s != nil {
		opu.SetPaySn(*s)
	}
	return opu
}

// ClearPaySn clears the value of the "pay_sn" field.
func (opu *OrderPayUpdate) ClearPaySn() *OrderPayUpdate {
	opu.mutation.ClearPaySn()
	return opu
}

// SetRemission sets the "remission" field.
func (opu *OrderPayUpdate) SetRemission(f float64) *OrderPayUpdate {
	opu.mutation.ResetRemission()
	opu.mutation.SetRemission(f)
	return opu
}

// SetNillableRemission sets the "remission" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableRemission(f *float64) *OrderPayUpdate {
	if f != nil {
		opu.SetRemission(*f)
	}
	return opu
}

// AddRemission adds f to the "remission" field.
func (opu *OrderPayUpdate) AddRemission(f float64) *OrderPayUpdate {
	opu.mutation.AddRemission(f)
	return opu
}

// ClearRemission clears the value of the "remission" field.
func (opu *OrderPayUpdate) ClearRemission() *OrderPayUpdate {
	opu.mutation.ClearRemission()
	return opu
}

// SetPay sets the "pay" field.
func (opu *OrderPayUpdate) SetPay(f float64) *OrderPayUpdate {
	opu.mutation.ResetPay()
	opu.mutation.SetPay(f)
	return opu
}

// SetNillablePay sets the "pay" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillablePay(f *float64) *OrderPayUpdate {
	if f != nil {
		opu.SetPay(*f)
	}
	return opu
}

// AddPay adds f to the "pay" field.
func (opu *OrderPayUpdate) AddPay(f float64) *OrderPayUpdate {
	opu.mutation.AddPay(f)
	return opu
}

// ClearPay clears the value of the "pay" field.
func (opu *OrderPayUpdate) ClearPay() *OrderPayUpdate {
	opu.mutation.ClearPay()
	return opu
}

// SetNote sets the "note" field.
func (opu *OrderPayUpdate) SetNote(s string) *OrderPayUpdate {
	opu.mutation.SetNote(s)
	return opu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableNote(s *string) *OrderPayUpdate {
	if s != nil {
		opu.SetNote(*s)
	}
	return opu
}

// ClearNote clears the value of the "note" field.
func (opu *OrderPayUpdate) ClearNote() *OrderPayUpdate {
	opu.mutation.ClearNote()
	return opu
}

// SetCreateID sets the "create_id" field.
func (opu *OrderPayUpdate) SetCreateID(i int64) *OrderPayUpdate {
	opu.mutation.ResetCreateID()
	opu.mutation.SetCreateID(i)
	return opu
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (opu *OrderPayUpdate) SetNillableCreateID(i *int64) *OrderPayUpdate {
	if i != nil {
		opu.SetCreateID(*i)
	}
	return opu
}

// AddCreateID adds i to the "create_id" field.
func (opu *OrderPayUpdate) AddCreateID(i int64) *OrderPayUpdate {
	opu.mutation.AddCreateID(i)
	return opu
}

// ClearCreateID clears the value of the "create_id" field.
func (opu *OrderPayUpdate) ClearCreateID() *OrderPayUpdate {
	opu.mutation.ClearCreateID()
	return opu
}

// SetOrder sets the "order" edge to the Order entity.
func (opu *OrderPayUpdate) SetOrder(o *Order) *OrderPayUpdate {
	return opu.SetOrderID(o.ID)
}

// Mutation returns the OrderPayMutation object of the builder.
func (opu *OrderPayUpdate) Mutation() *OrderPayMutation {
	return opu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (opu *OrderPayUpdate) ClearOrder() *OrderPayUpdate {
	opu.mutation.ClearOrder()
	return opu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OrderPayUpdate) Save(ctx context.Context) (int, error) {
	opu.defaults()
	return withHooks(ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OrderPayUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OrderPayUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OrderPayUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OrderPayUpdate) defaults() {
	if _, ok := opu.mutation.UpdatedAt(); !ok {
		v := orderpay.UpdateDefaultUpdatedAt()
		opu.mutation.SetUpdatedAt(v)
	}
}

func (opu *OrderPayUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderpay.Table, orderpay.Columns, sqlgraph.NewFieldSpec(orderpay.FieldID, field.TypeInt64))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opu.mutation.UpdatedAt(); ok {
		_spec.SetField(orderpay.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := opu.mutation.PaySn(); ok {
		_spec.SetField(orderpay.FieldPaySn, field.TypeString, value)
	}
	if opu.mutation.PaySnCleared() {
		_spec.ClearField(orderpay.FieldPaySn, field.TypeString)
	}
	if value, ok := opu.mutation.Remission(); ok {
		_spec.SetField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if value, ok := opu.mutation.AddedRemission(); ok {
		_spec.AddField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if opu.mutation.RemissionCleared() {
		_spec.ClearField(orderpay.FieldRemission, field.TypeFloat64)
	}
	if value, ok := opu.mutation.Pay(); ok {
		_spec.SetField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if value, ok := opu.mutation.AddedPay(); ok {
		_spec.AddField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if opu.mutation.PayCleared() {
		_spec.ClearField(orderpay.FieldPay, field.TypeFloat64)
	}
	if value, ok := opu.mutation.Note(); ok {
		_spec.SetField(orderpay.FieldNote, field.TypeString, value)
	}
	if opu.mutation.NoteCleared() {
		_spec.ClearField(orderpay.FieldNote, field.TypeString)
	}
	if value, ok := opu.mutation.CreateID(); ok {
		_spec.SetField(orderpay.FieldCreateID, field.TypeInt64, value)
	}
	if value, ok := opu.mutation.AddedCreateID(); ok {
		_spec.AddField(orderpay.FieldCreateID, field.TypeInt64, value)
	}
	if opu.mutation.CreateIDCleared() {
		_spec.ClearField(orderpay.FieldCreateID, field.TypeInt64)
	}
	if opu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpay.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OrderPayUpdateOne is the builder for updating a single OrderPay entity.
type OrderPayUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderPayMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (opuo *OrderPayUpdateOne) SetUpdatedAt(t time.Time) *OrderPayUpdateOne {
	opuo.mutation.SetUpdatedAt(t)
	return opuo
}

// SetOrderID sets the "order_id" field.
func (opuo *OrderPayUpdateOne) SetOrderID(i int64) *OrderPayUpdateOne {
	opuo.mutation.SetOrderID(i)
	return opuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableOrderID(i *int64) *OrderPayUpdateOne {
	if i != nil {
		opuo.SetOrderID(*i)
	}
	return opuo
}

// ClearOrderID clears the value of the "order_id" field.
func (opuo *OrderPayUpdateOne) ClearOrderID() *OrderPayUpdateOne {
	opuo.mutation.ClearOrderID()
	return opuo
}

// SetPaySn sets the "pay_sn" field.
func (opuo *OrderPayUpdateOne) SetPaySn(s string) *OrderPayUpdateOne {
	opuo.mutation.SetPaySn(s)
	return opuo
}

// SetNillablePaySn sets the "pay_sn" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePaySn(s *string) *OrderPayUpdateOne {
	if s != nil {
		opuo.SetPaySn(*s)
	}
	return opuo
}

// ClearPaySn clears the value of the "pay_sn" field.
func (opuo *OrderPayUpdateOne) ClearPaySn() *OrderPayUpdateOne {
	opuo.mutation.ClearPaySn()
	return opuo
}

// SetRemission sets the "remission" field.
func (opuo *OrderPayUpdateOne) SetRemission(f float64) *OrderPayUpdateOne {
	opuo.mutation.ResetRemission()
	opuo.mutation.SetRemission(f)
	return opuo
}

// SetNillableRemission sets the "remission" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableRemission(f *float64) *OrderPayUpdateOne {
	if f != nil {
		opuo.SetRemission(*f)
	}
	return opuo
}

// AddRemission adds f to the "remission" field.
func (opuo *OrderPayUpdateOne) AddRemission(f float64) *OrderPayUpdateOne {
	opuo.mutation.AddRemission(f)
	return opuo
}

// ClearRemission clears the value of the "remission" field.
func (opuo *OrderPayUpdateOne) ClearRemission() *OrderPayUpdateOne {
	opuo.mutation.ClearRemission()
	return opuo
}

// SetPay sets the "pay" field.
func (opuo *OrderPayUpdateOne) SetPay(f float64) *OrderPayUpdateOne {
	opuo.mutation.ResetPay()
	opuo.mutation.SetPay(f)
	return opuo
}

// SetNillablePay sets the "pay" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillablePay(f *float64) *OrderPayUpdateOne {
	if f != nil {
		opuo.SetPay(*f)
	}
	return opuo
}

// AddPay adds f to the "pay" field.
func (opuo *OrderPayUpdateOne) AddPay(f float64) *OrderPayUpdateOne {
	opuo.mutation.AddPay(f)
	return opuo
}

// ClearPay clears the value of the "pay" field.
func (opuo *OrderPayUpdateOne) ClearPay() *OrderPayUpdateOne {
	opuo.mutation.ClearPay()
	return opuo
}

// SetNote sets the "note" field.
func (opuo *OrderPayUpdateOne) SetNote(s string) *OrderPayUpdateOne {
	opuo.mutation.SetNote(s)
	return opuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableNote(s *string) *OrderPayUpdateOne {
	if s != nil {
		opuo.SetNote(*s)
	}
	return opuo
}

// ClearNote clears the value of the "note" field.
func (opuo *OrderPayUpdateOne) ClearNote() *OrderPayUpdateOne {
	opuo.mutation.ClearNote()
	return opuo
}

// SetCreateID sets the "create_id" field.
func (opuo *OrderPayUpdateOne) SetCreateID(i int64) *OrderPayUpdateOne {
	opuo.mutation.ResetCreateID()
	opuo.mutation.SetCreateID(i)
	return opuo
}

// SetNillableCreateID sets the "create_id" field if the given value is not nil.
func (opuo *OrderPayUpdateOne) SetNillableCreateID(i *int64) *OrderPayUpdateOne {
	if i != nil {
		opuo.SetCreateID(*i)
	}
	return opuo
}

// AddCreateID adds i to the "create_id" field.
func (opuo *OrderPayUpdateOne) AddCreateID(i int64) *OrderPayUpdateOne {
	opuo.mutation.AddCreateID(i)
	return opuo
}

// ClearCreateID clears the value of the "create_id" field.
func (opuo *OrderPayUpdateOne) ClearCreateID() *OrderPayUpdateOne {
	opuo.mutation.ClearCreateID()
	return opuo
}

// SetOrder sets the "order" edge to the Order entity.
func (opuo *OrderPayUpdateOne) SetOrder(o *Order) *OrderPayUpdateOne {
	return opuo.SetOrderID(o.ID)
}

// Mutation returns the OrderPayMutation object of the builder.
func (opuo *OrderPayUpdateOne) Mutation() *OrderPayMutation {
	return opuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (opuo *OrderPayUpdateOne) ClearOrder() *OrderPayUpdateOne {
	opuo.mutation.ClearOrder()
	return opuo
}

// Where appends a list predicates to the OrderPayUpdate builder.
func (opuo *OrderPayUpdateOne) Where(ps ...predicate.OrderPay) *OrderPayUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OrderPayUpdateOne) Select(field string, fields ...string) *OrderPayUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OrderPay entity.
func (opuo *OrderPayUpdateOne) Save(ctx context.Context) (*OrderPay, error) {
	opuo.defaults()
	return withHooks(ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OrderPayUpdateOne) SaveX(ctx context.Context) *OrderPay {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OrderPayUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OrderPayUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OrderPayUpdateOne) defaults() {
	if _, ok := opuo.mutation.UpdatedAt(); !ok {
		v := orderpay.UpdateDefaultUpdatedAt()
		opuo.mutation.SetUpdatedAt(v)
	}
}

func (opuo *OrderPayUpdateOne) sqlSave(ctx context.Context) (_node *OrderPay, err error) {
	_spec := sqlgraph.NewUpdateSpec(orderpay.Table, orderpay.Columns, sqlgraph.NewFieldSpec(orderpay.FieldID, field.TypeInt64))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderPay.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderpay.FieldID)
		for _, f := range fields {
			if !orderpay.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != orderpay.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opuo.mutation.UpdatedAt(); ok {
		_spec.SetField(orderpay.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := opuo.mutation.PaySn(); ok {
		_spec.SetField(orderpay.FieldPaySn, field.TypeString, value)
	}
	if opuo.mutation.PaySnCleared() {
		_spec.ClearField(orderpay.FieldPaySn, field.TypeString)
	}
	if value, ok := opuo.mutation.Remission(); ok {
		_spec.SetField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if value, ok := opuo.mutation.AddedRemission(); ok {
		_spec.AddField(orderpay.FieldRemission, field.TypeFloat64, value)
	}
	if opuo.mutation.RemissionCleared() {
		_spec.ClearField(orderpay.FieldRemission, field.TypeFloat64)
	}
	if value, ok := opuo.mutation.Pay(); ok {
		_spec.SetField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if value, ok := opuo.mutation.AddedPay(); ok {
		_spec.AddField(orderpay.FieldPay, field.TypeFloat64, value)
	}
	if opuo.mutation.PayCleared() {
		_spec.ClearField(orderpay.FieldPay, field.TypeFloat64)
	}
	if value, ok := opuo.mutation.Note(); ok {
		_spec.SetField(orderpay.FieldNote, field.TypeString, value)
	}
	if opuo.mutation.NoteCleared() {
		_spec.ClearField(orderpay.FieldNote, field.TypeString)
	}
	if value, ok := opuo.mutation.CreateID(); ok {
		_spec.SetField(orderpay.FieldCreateID, field.TypeInt64, value)
	}
	if value, ok := opuo.mutation.AddedCreateID(); ok {
		_spec.AddField(orderpay.FieldCreateID, field.TypeInt64, value)
	}
	if opuo.mutation.CreateIDCleared() {
		_spec.ClearField(orderpay.FieldCreateID, field.TypeInt64)
	}
	if opuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := opuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderpay.OrderTable,
			Columns: []string{orderpay.OrderColumn},
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
	_node = &OrderPay{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{orderpay.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}
