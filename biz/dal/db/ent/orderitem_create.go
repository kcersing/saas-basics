// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/order"
	"saas/biz/dal/db/ent/orderitem"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OrderItemCreate is the builder for creating a OrderItem entity.
type OrderItemCreate struct {
	config
	mutation *OrderItemMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oic *OrderItemCreate) SetCreatedAt(t time.Time) *OrderItemCreate {
	oic.mutation.SetCreatedAt(t)
	return oic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableCreatedAt(t *time.Time) *OrderItemCreate {
	if t != nil {
		oic.SetCreatedAt(*t)
	}
	return oic
}

// SetUpdatedAt sets the "updated_at" field.
func (oic *OrderItemCreate) SetUpdatedAt(t time.Time) *OrderItemCreate {
	oic.mutation.SetUpdatedAt(t)
	return oic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableUpdatedAt(t *time.Time) *OrderItemCreate {
	if t != nil {
		oic.SetUpdatedAt(*t)
	}
	return oic
}

// SetDelete sets the "delete" field.
func (oic *OrderItemCreate) SetDelete(i int64) *OrderItemCreate {
	oic.mutation.SetDelete(i)
	return oic
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableDelete(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetDelete(*i)
	}
	return oic
}

// SetCreatedID sets the "created_id" field.
func (oic *OrderItemCreate) SetCreatedID(i int64) *OrderItemCreate {
	oic.mutation.SetCreatedID(i)
	return oic
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableCreatedID(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetCreatedID(*i)
	}
	return oic
}

// SetOrderID sets the "order_id" field.
func (oic *OrderItemCreate) SetOrderID(i int64) *OrderItemCreate {
	oic.mutation.SetOrderID(i)
	return oic
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableOrderID(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetOrderID(*i)
	}
	return oic
}

// SetNumber sets the "number" field.
func (oic *OrderItemCreate) SetNumber(i int64) *OrderItemCreate {
	oic.mutation.SetNumber(i)
	return oic
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableNumber(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetNumber(*i)
	}
	return oic
}

// SetProductID sets the "product_id" field.
func (oic *OrderItemCreate) SetProductID(i int64) *OrderItemCreate {
	oic.mutation.SetProductID(i)
	return oic
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableProductID(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetProductID(*i)
	}
	return oic
}

// SetRelatedUserProductID sets the "related_user_product_id" field.
func (oic *OrderItemCreate) SetRelatedUserProductID(i int64) *OrderItemCreate {
	oic.mutation.SetRelatedUserProductID(i)
	return oic
}

// SetNillableRelatedUserProductID sets the "related_user_product_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableRelatedUserProductID(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetRelatedUserProductID(*i)
	}
	return oic
}

// SetContestID sets the "contest_id" field.
func (oic *OrderItemCreate) SetContestID(i int64) *OrderItemCreate {
	oic.mutation.SetContestID(i)
	return oic
}

// SetNillableContestID sets the "contest_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableContestID(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetContestID(*i)
	}
	return oic
}

// SetBootcampID sets the "bootcamp_id" field.
func (oic *OrderItemCreate) SetBootcampID(i int64) *OrderItemCreate {
	oic.mutation.SetBootcampID(i)
	return oic
}

// SetNillableBootcampID sets the "bootcamp_id" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableBootcampID(i *int64) *OrderItemCreate {
	if i != nil {
		oic.SetBootcampID(*i)
	}
	return oic
}

// SetName sets the "name" field.
func (oic *OrderItemCreate) SetName(s string) *OrderItemCreate {
	oic.mutation.SetName(s)
	return oic
}

// SetNillableName sets the "name" field if the given value is not nil.
func (oic *OrderItemCreate) SetNillableName(s *string) *OrderItemCreate {
	if s != nil {
		oic.SetName(*s)
	}
	return oic
}

// SetData sets the "data" field.
func (oic *OrderItemCreate) SetData(s []string) *OrderItemCreate {
	oic.mutation.SetData(s)
	return oic
}

// SetID sets the "id" field.
func (oic *OrderItemCreate) SetID(i int64) *OrderItemCreate {
	oic.mutation.SetID(i)
	return oic
}

// SetOrder sets the "order" edge to the Order entity.
func (oic *OrderItemCreate) SetOrder(o *Order) *OrderItemCreate {
	return oic.SetOrderID(o.ID)
}

// Mutation returns the OrderItemMutation object of the builder.
func (oic *OrderItemCreate) Mutation() *OrderItemMutation {
	return oic.mutation
}

// Save creates the OrderItem in the database.
func (oic *OrderItemCreate) Save(ctx context.Context) (*OrderItem, error) {
	oic.defaults()
	return withHooks(ctx, oic.sqlSave, oic.mutation, oic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oic *OrderItemCreate) SaveX(ctx context.Context) *OrderItem {
	v, err := oic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oic *OrderItemCreate) Exec(ctx context.Context) error {
	_, err := oic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oic *OrderItemCreate) ExecX(ctx context.Context) {
	if err := oic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oic *OrderItemCreate) defaults() {
	if _, ok := oic.mutation.CreatedAt(); !ok {
		v := orderitem.DefaultCreatedAt()
		oic.mutation.SetCreatedAt(v)
	}
	if _, ok := oic.mutation.UpdatedAt(); !ok {
		v := orderitem.DefaultUpdatedAt()
		oic.mutation.SetUpdatedAt(v)
	}
	if _, ok := oic.mutation.Delete(); !ok {
		v := orderitem.DefaultDelete
		oic.mutation.SetDelete(v)
	}
	if _, ok := oic.mutation.CreatedID(); !ok {
		v := orderitem.DefaultCreatedID
		oic.mutation.SetCreatedID(v)
	}
	if _, ok := oic.mutation.Number(); !ok {
		v := orderitem.DefaultNumber
		oic.mutation.SetNumber(v)
	}
	if _, ok := oic.mutation.RelatedUserProductID(); !ok {
		v := orderitem.DefaultRelatedUserProductID
		oic.mutation.SetRelatedUserProductID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oic *OrderItemCreate) check() error {
	return nil
}

func (oic *OrderItemCreate) sqlSave(ctx context.Context) (*OrderItem, error) {
	if err := oic.check(); err != nil {
		return nil, err
	}
	_node, _spec := oic.createSpec()
	if err := sqlgraph.CreateNode(ctx, oic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	oic.mutation.id = &_node.ID
	oic.mutation.done = true
	return _node, nil
}

func (oic *OrderItemCreate) createSpec() (*OrderItem, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderItem{config: oic.config}
		_spec = sqlgraph.NewCreateSpec(orderitem.Table, sqlgraph.NewFieldSpec(orderitem.FieldID, field.TypeInt64))
	)
	if id, ok := oic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oic.mutation.CreatedAt(); ok {
		_spec.SetField(orderitem.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oic.mutation.UpdatedAt(); ok {
		_spec.SetField(orderitem.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oic.mutation.Delete(); ok {
		_spec.SetField(orderitem.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := oic.mutation.CreatedID(); ok {
		_spec.SetField(orderitem.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := oic.mutation.Number(); ok {
		_spec.SetField(orderitem.FieldNumber, field.TypeInt64, value)
		_node.Number = value
	}
	if value, ok := oic.mutation.ProductID(); ok {
		_spec.SetField(orderitem.FieldProductID, field.TypeInt64, value)
		_node.ProductID = value
	}
	if value, ok := oic.mutation.RelatedUserProductID(); ok {
		_spec.SetField(orderitem.FieldRelatedUserProductID, field.TypeInt64, value)
		_node.RelatedUserProductID = value
	}
	if value, ok := oic.mutation.ContestID(); ok {
		_spec.SetField(orderitem.FieldContestID, field.TypeInt64, value)
		_node.ContestID = value
	}
	if value, ok := oic.mutation.BootcampID(); ok {
		_spec.SetField(orderitem.FieldBootcampID, field.TypeInt64, value)
		_node.BootcampID = value
	}
	if value, ok := oic.mutation.Name(); ok {
		_spec.SetField(orderitem.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := oic.mutation.Data(); ok {
		_spec.SetField(orderitem.FieldData, field.TypeJSON, value)
		_node.Data = value
	}
	if nodes := oic.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   orderitem.OrderTable,
			Columns: []string{orderitem.OrderColumn},
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

// OrderItemCreateBulk is the builder for creating many OrderItem entities in bulk.
type OrderItemCreateBulk struct {
	config
	err      error
	builders []*OrderItemCreate
}

// Save creates the OrderItem entities in the database.
func (oicb *OrderItemCreateBulk) Save(ctx context.Context) ([]*OrderItem, error) {
	if oicb.err != nil {
		return nil, oicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oicb.builders))
	nodes := make([]*OrderItem, len(oicb.builders))
	mutators := make([]Mutator, len(oicb.builders))
	for i := range oicb.builders {
		func(i int, root context.Context) {
			builder := oicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderItemMutation)
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
					_, err = mutators[i+1].Mutate(root, oicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, oicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) SaveX(ctx context.Context) []*OrderItem {
	v, err := oicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oicb *OrderItemCreateBulk) Exec(ctx context.Context) error {
	_, err := oicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oicb *OrderItemCreateBulk) ExecX(ctx context.Context) {
	if err := oicb.Exec(ctx); err != nil {
		panic(err)
	}
}
