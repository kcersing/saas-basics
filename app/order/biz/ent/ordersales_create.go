// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"order/biz/ent/order"
	"order/biz/ent/ordersales"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderSalesCreate is the builder for creating a OrderSales entity.
type OrderSalesCreate struct {
	config
	mutation *OrderSalesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (osc *OrderSalesCreate) SetCreatedAt(t time.Time) *OrderSalesCreate {
	osc.mutation.SetCreatedAt(t)
	return osc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableCreatedAt(t *time.Time) *OrderSalesCreate {
	if t != nil {
		osc.SetCreatedAt(*t)
	}
	return osc
}

// SetUpdatedAt sets the "updated_at" field.
func (osc *OrderSalesCreate) SetUpdatedAt(t time.Time) *OrderSalesCreate {
	osc.mutation.SetUpdatedAt(t)
	return osc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableUpdatedAt(t *time.Time) *OrderSalesCreate {
	if t != nil {
		osc.SetUpdatedAt(*t)
	}
	return osc
}

// SetStatus sets the "status" field.
func (osc *OrderSalesCreate) SetStatus(i int64) *OrderSalesCreate {
	osc.mutation.SetStatus(i)
	return osc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableStatus(i *int64) *OrderSalesCreate {
	if i != nil {
		osc.SetStatus(*i)
	}
	return osc
}

// SetOrderID sets the "order_id" field.
func (osc *OrderSalesCreate) SetOrderID(i int64) *OrderSalesCreate {
	osc.mutation.SetOrderID(i)
	return osc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableOrderID(i *int64) *OrderSalesCreate {
	if i != nil {
		osc.SetOrderID(*i)
	}
	return osc
}

// SetMemberID sets the "member_id" field.
func (osc *OrderSalesCreate) SetMemberID(i int64) *OrderSalesCreate {
	osc.mutation.SetMemberID(i)
	return osc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableMemberID(i *int64) *OrderSalesCreate {
	if i != nil {
		osc.SetMemberID(*i)
	}
	return osc
}

// SetSalesID sets the "sales_id" field.
func (osc *OrderSalesCreate) SetSalesID(i int64) *OrderSalesCreate {
	osc.mutation.SetSalesID(i)
	return osc
}

// SetNillableSalesID sets the "sales_id" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableSalesID(i *int64) *OrderSalesCreate {
	if i != nil {
		osc.SetSalesID(*i)
	}
	return osc
}

// SetRatio sets the "ratio" field.
func (osc *OrderSalesCreate) SetRatio(i int64) *OrderSalesCreate {
	osc.mutation.SetRatio(i)
	return osc
}

// SetNillableRatio sets the "ratio" field if the given value is not nil.
func (osc *OrderSalesCreate) SetNillableRatio(i *int64) *OrderSalesCreate {
	if i != nil {
		osc.SetRatio(*i)
	}
	return osc
}

// SetID sets the "id" field.
func (osc *OrderSalesCreate) SetID(i int64) *OrderSalesCreate {
	osc.mutation.SetID(i)
	return osc
}

// SetOrder sets the "order" edge to the Order entity.
func (osc *OrderSalesCreate) SetOrder(o *Order) *OrderSalesCreate {
	return osc.SetOrderID(o.ID)
}

// Mutation returns the OrderSalesMutation object of the builder.
func (osc *OrderSalesCreate) Mutation() *OrderSalesMutation {
	return osc.mutation
}

// Save creates the OrderSales in the database.
func (osc *OrderSalesCreate) Save(ctx context.Context) (*OrderSales, error) {
	osc.defaults()
	return withHooks(ctx, osc.sqlSave, osc.mutation, osc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (osc *OrderSalesCreate) SaveX(ctx context.Context) *OrderSales {
	v, err := osc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (osc *OrderSalesCreate) Exec(ctx context.Context) error {
	_, err := osc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osc *OrderSalesCreate) ExecX(ctx context.Context) {
	if err := osc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osc *OrderSalesCreate) defaults() {
	if _, ok := osc.mutation.CreatedAt(); !ok {
		v := ordersales.DefaultCreatedAt()
		osc.mutation.SetCreatedAt(v)
	}
	if _, ok := osc.mutation.UpdatedAt(); !ok {
		v := ordersales.DefaultUpdatedAt()
		osc.mutation.SetUpdatedAt(v)
	}
	if _, ok := osc.mutation.Status(); !ok {
		v := ordersales.DefaultStatus
		osc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (osc *OrderSalesCreate) check() error {
	if _, ok := osc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderSales.created_at"`)}
	}
	if _, ok := osc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderSales.updated_at"`)}
	}
	return nil
}

func (osc *OrderSalesCreate) sqlSave(ctx context.Context) (*OrderSales, error) {
	if err := osc.check(); err != nil {
		return nil, err
	}
	_node, _spec := osc.createSpec()
	if err := sqlgraph.CreateNode(ctx, osc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	osc.mutation.id = &_node.ID
	osc.mutation.done = true
	return _node, nil
}

func (osc *OrderSalesCreate) createSpec() (*OrderSales, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderSales{config: osc.config}
		_spec = sqlgraph.NewCreateSpec(ordersales.Table, sqlgraph.NewFieldSpec(ordersales.FieldID, field.TypeInt64))
	)
	if id, ok := osc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := osc.mutation.CreatedAt(); ok {
		_spec.SetField(ordersales.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := osc.mutation.UpdatedAt(); ok {
		_spec.SetField(ordersales.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := osc.mutation.Status(); ok {
		_spec.SetField(ordersales.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := osc.mutation.MemberID(); ok {
		_spec.SetField(ordersales.FieldMemberID, field.TypeInt64, value)
		_node.MemberID = value
	}
	if value, ok := osc.mutation.SalesID(); ok {
		_spec.SetField(ordersales.FieldSalesID, field.TypeInt64, value)
		_node.SalesID = value
	}
	if value, ok := osc.mutation.Ratio(); ok {
		_spec.SetField(ordersales.FieldRatio, field.TypeInt64, value)
		_node.Ratio = value
	}
	if nodes := osc.mutation.OrderIDs(); len(nodes) > 0 {
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
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderSalesCreateBulk is the builder for creating many OrderSales entities in bulk.
type OrderSalesCreateBulk struct {
	config
	err      error
	builders []*OrderSalesCreate
}

// Save creates the OrderSales entities in the database.
func (oscb *OrderSalesCreateBulk) Save(ctx context.Context) ([]*OrderSales, error) {
	if oscb.err != nil {
		return nil, oscb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oscb.builders))
	nodes := make([]*OrderSales, len(oscb.builders))
	mutators := make([]Mutator, len(oscb.builders))
	for i := range oscb.builders {
		func(i int, root context.Context) {
			builder := oscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderSalesMutation)
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
					_, err = mutators[i+1].Mutate(root, oscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oscb *OrderSalesCreateBulk) SaveX(ctx context.Context) []*OrderSales {
	v, err := oscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oscb *OrderSalesCreateBulk) Exec(ctx context.Context) error {
	_, err := oscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oscb *OrderSalesCreateBulk) ExecX(ctx context.Context) {
	if err := oscb.Exec(ctx); err != nil {
		panic(err)
	}
}
