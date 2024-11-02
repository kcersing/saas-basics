// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"order/biz/dal/mysql/ent/order"
	"order/biz/dal/mysql/ent/ordersales"
	"order/biz/dal/mysql/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderSalesUpdate is the builder for updating OrderSales entities.
type OrderSalesUpdate struct {
	config
	hooks    []Hook
	mutation *OrderSalesMutation
}

// Where appends a list predicates to the OrderSalesUpdate builder.
func (osu *OrderSalesUpdate) Where(ps ...predicate.OrderSales) *OrderSalesUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetUpdatedAt sets the "updated_at" field.
func (osu *OrderSalesUpdate) SetUpdatedAt(t time.Time) *OrderSalesUpdate {
	osu.mutation.SetUpdatedAt(t)
	return osu
}

// SetStatus sets the "status" field.
func (osu *OrderSalesUpdate) SetStatus(i int64) *OrderSalesUpdate {
	osu.mutation.ResetStatus()
	osu.mutation.SetStatus(i)
	return osu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (osu *OrderSalesUpdate) SetNillableStatus(i *int64) *OrderSalesUpdate {
	if i != nil {
		osu.SetStatus(*i)
	}
	return osu
}

// AddStatus adds i to the "status" field.
func (osu *OrderSalesUpdate) AddStatus(i int64) *OrderSalesUpdate {
	osu.mutation.AddStatus(i)
	return osu
}

// ClearStatus clears the value of the "status" field.
func (osu *OrderSalesUpdate) ClearStatus() *OrderSalesUpdate {
	osu.mutation.ClearStatus()
	return osu
}

// SetOrderID sets the "order_id" field.
func (osu *OrderSalesUpdate) SetOrderID(i int64) *OrderSalesUpdate {
	osu.mutation.SetOrderID(i)
	return osu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osu *OrderSalesUpdate) SetNillableOrderID(i *int64) *OrderSalesUpdate {
	if i != nil {
		osu.SetOrderID(*i)
	}
	return osu
}

// ClearOrderID clears the value of the "order_id" field.
func (osu *OrderSalesUpdate) ClearOrderID() *OrderSalesUpdate {
	osu.mutation.ClearOrderID()
	return osu
}

// SetMemberID sets the "member_id" field.
func (osu *OrderSalesUpdate) SetMemberID(i int64) *OrderSalesUpdate {
	osu.mutation.ResetMemberID()
	osu.mutation.SetMemberID(i)
	return osu
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (osu *OrderSalesUpdate) SetNillableMemberID(i *int64) *OrderSalesUpdate {
	if i != nil {
		osu.SetMemberID(*i)
	}
	return osu
}

// AddMemberID adds i to the "member_id" field.
func (osu *OrderSalesUpdate) AddMemberID(i int64) *OrderSalesUpdate {
	osu.mutation.AddMemberID(i)
	return osu
}

// ClearMemberID clears the value of the "member_id" field.
func (osu *OrderSalesUpdate) ClearMemberID() *OrderSalesUpdate {
	osu.mutation.ClearMemberID()
	return osu
}

// SetSalesID sets the "sales_id" field.
func (osu *OrderSalesUpdate) SetSalesID(i int64) *OrderSalesUpdate {
	osu.mutation.ResetSalesID()
	osu.mutation.SetSalesID(i)
	return osu
}

// SetNillableSalesID sets the "sales_id" field if the given value is not nil.
func (osu *OrderSalesUpdate) SetNillableSalesID(i *int64) *OrderSalesUpdate {
	if i != nil {
		osu.SetSalesID(*i)
	}
	return osu
}

// AddSalesID adds i to the "sales_id" field.
func (osu *OrderSalesUpdate) AddSalesID(i int64) *OrderSalesUpdate {
	osu.mutation.AddSalesID(i)
	return osu
}

// ClearSalesID clears the value of the "sales_id" field.
func (osu *OrderSalesUpdate) ClearSalesID() *OrderSalesUpdate {
	osu.mutation.ClearSalesID()
	return osu
}

// SetRatio sets the "ratio" field.
func (osu *OrderSalesUpdate) SetRatio(i int64) *OrderSalesUpdate {
	osu.mutation.ResetRatio()
	osu.mutation.SetRatio(i)
	return osu
}

// SetNillableRatio sets the "ratio" field if the given value is not nil.
func (osu *OrderSalesUpdate) SetNillableRatio(i *int64) *OrderSalesUpdate {
	if i != nil {
		osu.SetRatio(*i)
	}
	return osu
}

// AddRatio adds i to the "ratio" field.
func (osu *OrderSalesUpdate) AddRatio(i int64) *OrderSalesUpdate {
	osu.mutation.AddRatio(i)
	return osu
}

// ClearRatio clears the value of the "ratio" field.
func (osu *OrderSalesUpdate) ClearRatio() *OrderSalesUpdate {
	osu.mutation.ClearRatio()
	return osu
}

// SetOrder sets the "order" edge to the Order entity.
func (osu *OrderSalesUpdate) SetOrder(o *Order) *OrderSalesUpdate {
	return osu.SetOrderID(o.ID)
}

// Mutation returns the OrderSalesMutation object of the builder.
func (osu *OrderSalesUpdate) Mutation() *OrderSalesMutation {
	return osu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (osu *OrderSalesUpdate) ClearOrder() *OrderSalesUpdate {
	osu.mutation.ClearOrder()
	return osu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OrderSalesUpdate) Save(ctx context.Context) (int, error) {
	osu.defaults()
	return withHooks(ctx, osu.sqlSave, osu.mutation, osu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OrderSalesUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OrderSalesUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OrderSalesUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osu *OrderSalesUpdate) defaults() {
	if _, ok := osu.mutation.UpdatedAt(); !ok {
		v := ordersales.UpdateDefaultUpdatedAt()
		osu.mutation.SetUpdatedAt(v)
	}
}

func (osu *OrderSalesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ordersales.Table, ordersales.Columns, sqlgraph.NewFieldSpec(ordersales.FieldID, field.TypeInt64))
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.UpdatedAt(); ok {
		_spec.SetField(ordersales.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := osu.mutation.Status(); ok {
		_spec.SetField(ordersales.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := osu.mutation.AddedStatus(); ok {
		_spec.AddField(ordersales.FieldStatus, field.TypeInt64, value)
	}
	if osu.mutation.StatusCleared() {
		_spec.ClearField(ordersales.FieldStatus, field.TypeInt64)
	}
	if value, ok := osu.mutation.MemberID(); ok {
		_spec.SetField(ordersales.FieldMemberID, field.TypeInt64, value)
	}
	if value, ok := osu.mutation.AddedMemberID(); ok {
		_spec.AddField(ordersales.FieldMemberID, field.TypeInt64, value)
	}
	if osu.mutation.MemberIDCleared() {
		_spec.ClearField(ordersales.FieldMemberID, field.TypeInt64)
	}
	if value, ok := osu.mutation.SalesID(); ok {
		_spec.SetField(ordersales.FieldSalesID, field.TypeInt64, value)
	}
	if value, ok := osu.mutation.AddedSalesID(); ok {
		_spec.AddField(ordersales.FieldSalesID, field.TypeInt64, value)
	}
	if osu.mutation.SalesIDCleared() {
		_spec.ClearField(ordersales.FieldSalesID, field.TypeInt64)
	}
	if value, ok := osu.mutation.Ratio(); ok {
		_spec.SetField(ordersales.FieldRatio, field.TypeInt64, value)
	}
	if value, ok := osu.mutation.AddedRatio(); ok {
		_spec.AddField(ordersales.FieldRatio, field.TypeInt64, value)
	}
	if osu.mutation.RatioCleared() {
		_spec.ClearField(ordersales.FieldRatio, field.TypeInt64)
	}
	if osu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersales.OrderTable,
			Columns: []string{ordersales.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersales.OrderTable,
			Columns: []string{ordersales.OrderColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordersales.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	osu.mutation.done = true
	return n, nil
}

// OrderSalesUpdateOne is the builder for updating a single OrderSales entity.
type OrderSalesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderSalesMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (osuo *OrderSalesUpdateOne) SetUpdatedAt(t time.Time) *OrderSalesUpdateOne {
	osuo.mutation.SetUpdatedAt(t)
	return osuo
}

// SetStatus sets the "status" field.
func (osuo *OrderSalesUpdateOne) SetStatus(i int64) *OrderSalesUpdateOne {
	osuo.mutation.ResetStatus()
	osuo.mutation.SetStatus(i)
	return osuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (osuo *OrderSalesUpdateOne) SetNillableStatus(i *int64) *OrderSalesUpdateOne {
	if i != nil {
		osuo.SetStatus(*i)
	}
	return osuo
}

// AddStatus adds i to the "status" field.
func (osuo *OrderSalesUpdateOne) AddStatus(i int64) *OrderSalesUpdateOne {
	osuo.mutation.AddStatus(i)
	return osuo
}

// ClearStatus clears the value of the "status" field.
func (osuo *OrderSalesUpdateOne) ClearStatus() *OrderSalesUpdateOne {
	osuo.mutation.ClearStatus()
	return osuo
}

// SetOrderID sets the "order_id" field.
func (osuo *OrderSalesUpdateOne) SetOrderID(i int64) *OrderSalesUpdateOne {
	osuo.mutation.SetOrderID(i)
	return osuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osuo *OrderSalesUpdateOne) SetNillableOrderID(i *int64) *OrderSalesUpdateOne {
	if i != nil {
		osuo.SetOrderID(*i)
	}
	return osuo
}

// ClearOrderID clears the value of the "order_id" field.
func (osuo *OrderSalesUpdateOne) ClearOrderID() *OrderSalesUpdateOne {
	osuo.mutation.ClearOrderID()
	return osuo
}

// SetMemberID sets the "member_id" field.
func (osuo *OrderSalesUpdateOne) SetMemberID(i int64) *OrderSalesUpdateOne {
	osuo.mutation.ResetMemberID()
	osuo.mutation.SetMemberID(i)
	return osuo
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (osuo *OrderSalesUpdateOne) SetNillableMemberID(i *int64) *OrderSalesUpdateOne {
	if i != nil {
		osuo.SetMemberID(*i)
	}
	return osuo
}

// AddMemberID adds i to the "member_id" field.
func (osuo *OrderSalesUpdateOne) AddMemberID(i int64) *OrderSalesUpdateOne {
	osuo.mutation.AddMemberID(i)
	return osuo
}

// ClearMemberID clears the value of the "member_id" field.
func (osuo *OrderSalesUpdateOne) ClearMemberID() *OrderSalesUpdateOne {
	osuo.mutation.ClearMemberID()
	return osuo
}

// SetSalesID sets the "sales_id" field.
func (osuo *OrderSalesUpdateOne) SetSalesID(i int64) *OrderSalesUpdateOne {
	osuo.mutation.ResetSalesID()
	osuo.mutation.SetSalesID(i)
	return osuo
}

// SetNillableSalesID sets the "sales_id" field if the given value is not nil.
func (osuo *OrderSalesUpdateOne) SetNillableSalesID(i *int64) *OrderSalesUpdateOne {
	if i != nil {
		osuo.SetSalesID(*i)
	}
	return osuo
}

// AddSalesID adds i to the "sales_id" field.
func (osuo *OrderSalesUpdateOne) AddSalesID(i int64) *OrderSalesUpdateOne {
	osuo.mutation.AddSalesID(i)
	return osuo
}

// ClearSalesID clears the value of the "sales_id" field.
func (osuo *OrderSalesUpdateOne) ClearSalesID() *OrderSalesUpdateOne {
	osuo.mutation.ClearSalesID()
	return osuo
}

// SetRatio sets the "ratio" field.
func (osuo *OrderSalesUpdateOne) SetRatio(i int64) *OrderSalesUpdateOne {
	osuo.mutation.ResetRatio()
	osuo.mutation.SetRatio(i)
	return osuo
}

// SetNillableRatio sets the "ratio" field if the given value is not nil.
func (osuo *OrderSalesUpdateOne) SetNillableRatio(i *int64) *OrderSalesUpdateOne {
	if i != nil {
		osuo.SetRatio(*i)
	}
	return osuo
}

// AddRatio adds i to the "ratio" field.
func (osuo *OrderSalesUpdateOne) AddRatio(i int64) *OrderSalesUpdateOne {
	osuo.mutation.AddRatio(i)
	return osuo
}

// ClearRatio clears the value of the "ratio" field.
func (osuo *OrderSalesUpdateOne) ClearRatio() *OrderSalesUpdateOne {
	osuo.mutation.ClearRatio()
	return osuo
}

// SetOrder sets the "order" edge to the Order entity.
func (osuo *OrderSalesUpdateOne) SetOrder(o *Order) *OrderSalesUpdateOne {
	return osuo.SetOrderID(o.ID)
}

// Mutation returns the OrderSalesMutation object of the builder.
func (osuo *OrderSalesUpdateOne) Mutation() *OrderSalesMutation {
	return osuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (osuo *OrderSalesUpdateOne) ClearOrder() *OrderSalesUpdateOne {
	osuo.mutation.ClearOrder()
	return osuo
}

// Where appends a list predicates to the OrderSalesUpdate builder.
func (osuo *OrderSalesUpdateOne) Where(ps ...predicate.OrderSales) *OrderSalesUpdateOne {
	osuo.mutation.Where(ps...)
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OrderSalesUpdateOne) Select(field string, fields ...string) *OrderSalesUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OrderSales entity.
func (osuo *OrderSalesUpdateOne) Save(ctx context.Context) (*OrderSales, error) {
	osuo.defaults()
	return withHooks(ctx, osuo.sqlSave, osuo.mutation, osuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OrderSalesUpdateOne) SaveX(ctx context.Context) *OrderSales {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OrderSalesUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OrderSalesUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osuo *OrderSalesUpdateOne) defaults() {
	if _, ok := osuo.mutation.UpdatedAt(); !ok {
		v := ordersales.UpdateDefaultUpdatedAt()
		osuo.mutation.SetUpdatedAt(v)
	}
}

func (osuo *OrderSalesUpdateOne) sqlSave(ctx context.Context) (_node *OrderSales, err error) {
	_spec := sqlgraph.NewUpdateSpec(ordersales.Table, ordersales.Columns, sqlgraph.NewFieldSpec(ordersales.FieldID, field.TypeInt64))
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrderSales.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ordersales.FieldID)
		for _, f := range fields {
			if !ordersales.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ordersales.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.UpdatedAt(); ok {
		_spec.SetField(ordersales.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := osuo.mutation.Status(); ok {
		_spec.SetField(ordersales.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := osuo.mutation.AddedStatus(); ok {
		_spec.AddField(ordersales.FieldStatus, field.TypeInt64, value)
	}
	if osuo.mutation.StatusCleared() {
		_spec.ClearField(ordersales.FieldStatus, field.TypeInt64)
	}
	if value, ok := osuo.mutation.MemberID(); ok {
		_spec.SetField(ordersales.FieldMemberID, field.TypeInt64, value)
	}
	if value, ok := osuo.mutation.AddedMemberID(); ok {
		_spec.AddField(ordersales.FieldMemberID, field.TypeInt64, value)
	}
	if osuo.mutation.MemberIDCleared() {
		_spec.ClearField(ordersales.FieldMemberID, field.TypeInt64)
	}
	if value, ok := osuo.mutation.SalesID(); ok {
		_spec.SetField(ordersales.FieldSalesID, field.TypeInt64, value)
	}
	if value, ok := osuo.mutation.AddedSalesID(); ok {
		_spec.AddField(ordersales.FieldSalesID, field.TypeInt64, value)
	}
	if osuo.mutation.SalesIDCleared() {
		_spec.ClearField(ordersales.FieldSalesID, field.TypeInt64)
	}
	if value, ok := osuo.mutation.Ratio(); ok {
		_spec.SetField(ordersales.FieldRatio, field.TypeInt64, value)
	}
	if value, ok := osuo.mutation.AddedRatio(); ok {
		_spec.AddField(ordersales.FieldRatio, field.TypeInt64, value)
	}
	if osuo.mutation.RatioCleared() {
		_spec.ClearField(ordersales.FieldRatio, field.TypeInt64)
	}
	if osuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersales.OrderTable,
			Columns: []string{ordersales.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ordersales.OrderTable,
			Columns: []string{ordersales.OrderColumn},
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
	_node = &OrderSales{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ordersales.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	osuo.mutation.done = true
	return _node, nil
}
