// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/productcourses"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCoursesCreate is the builder for creating a ProductCourses entity.
type ProductCoursesCreate struct {
	config
	mutation *ProductCoursesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pcc *ProductCoursesCreate) SetCreatedAt(t time.Time) *ProductCoursesCreate {
	pcc.mutation.SetCreatedAt(t)
	return pcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableCreatedAt(t *time.Time) *ProductCoursesCreate {
	if t != nil {
		pcc.SetCreatedAt(*t)
	}
	return pcc
}

// SetUpdatedAt sets the "updated_at" field.
func (pcc *ProductCoursesCreate) SetUpdatedAt(t time.Time) *ProductCoursesCreate {
	pcc.mutation.SetUpdatedAt(t)
	return pcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableUpdatedAt(t *time.Time) *ProductCoursesCreate {
	if t != nil {
		pcc.SetUpdatedAt(*t)
	}
	return pcc
}

// SetDelete sets the "delete" field.
func (pcc *ProductCoursesCreate) SetDelete(i int64) *ProductCoursesCreate {
	pcc.mutation.SetDelete(i)
	return pcc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableDelete(i *int64) *ProductCoursesCreate {
	if i != nil {
		pcc.SetDelete(*i)
	}
	return pcc
}

// SetCreatedID sets the "created_id" field.
func (pcc *ProductCoursesCreate) SetCreatedID(i int64) *ProductCoursesCreate {
	pcc.mutation.SetCreatedID(i)
	return pcc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableCreatedID(i *int64) *ProductCoursesCreate {
	if i != nil {
		pcc.SetCreatedID(*i)
	}
	return pcc
}

// SetStatus sets the "status" field.
func (pcc *ProductCoursesCreate) SetStatus(i int64) *ProductCoursesCreate {
	pcc.mutation.SetStatus(i)
	return pcc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableStatus(i *int64) *ProductCoursesCreate {
	if i != nil {
		pcc.SetStatus(*i)
	}
	return pcc
}

// SetType sets the "type" field.
func (pcc *ProductCoursesCreate) SetType(s string) *ProductCoursesCreate {
	pcc.mutation.SetType(s)
	return pcc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableType(s *string) *ProductCoursesCreate {
	if s != nil {
		pcc.SetType(*s)
	}
	return pcc
}

// SetName sets the "name" field.
func (pcc *ProductCoursesCreate) SetName(s string) *ProductCoursesCreate {
	pcc.mutation.SetName(s)
	return pcc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableName(s *string) *ProductCoursesCreate {
	if s != nil {
		pcc.SetName(*s)
	}
	return pcc
}

// SetNumber sets the "number" field.
func (pcc *ProductCoursesCreate) SetNumber(i int64) *ProductCoursesCreate {
	pcc.mutation.SetNumber(i)
	return pcc
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableNumber(i *int64) *ProductCoursesCreate {
	if i != nil {
		pcc.SetNumber(*i)
	}
	return pcc
}

// SetProductID sets the "product_id" field.
func (pcc *ProductCoursesCreate) SetProductID(i int64) *ProductCoursesCreate {
	pcc.mutation.SetProductID(i)
	return pcc
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableProductID(i *int64) *ProductCoursesCreate {
	if i != nil {
		pcc.SetProductID(*i)
	}
	return pcc
}

// SetCoursesID sets the "courses_id" field.
func (pcc *ProductCoursesCreate) SetCoursesID(i int64) *ProductCoursesCreate {
	pcc.mutation.SetCoursesID(i)
	return pcc
}

// SetNillableCoursesID sets the "courses_id" field if the given value is not nil.
func (pcc *ProductCoursesCreate) SetNillableCoursesID(i *int64) *ProductCoursesCreate {
	if i != nil {
		pcc.SetCoursesID(*i)
	}
	return pcc
}

// SetID sets the "id" field.
func (pcc *ProductCoursesCreate) SetID(i int64) *ProductCoursesCreate {
	pcc.mutation.SetID(i)
	return pcc
}

// SetProduct sets the "product" edge to the Product entity.
func (pcc *ProductCoursesCreate) SetProduct(p *Product) *ProductCoursesCreate {
	return pcc.SetProductID(p.ID)
}

// Mutation returns the ProductCoursesMutation object of the builder.
func (pcc *ProductCoursesCreate) Mutation() *ProductCoursesMutation {
	return pcc.mutation
}

// Save creates the ProductCourses in the database.
func (pcc *ProductCoursesCreate) Save(ctx context.Context) (*ProductCourses, error) {
	pcc.defaults()
	return withHooks(ctx, pcc.sqlSave, pcc.mutation, pcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *ProductCoursesCreate) SaveX(ctx context.Context) *ProductCourses {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *ProductCoursesCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *ProductCoursesCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *ProductCoursesCreate) defaults() {
	if _, ok := pcc.mutation.CreatedAt(); !ok {
		v := productcourses.DefaultCreatedAt()
		pcc.mutation.SetCreatedAt(v)
	}
	if _, ok := pcc.mutation.UpdatedAt(); !ok {
		v := productcourses.DefaultUpdatedAt()
		pcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pcc.mutation.Delete(); !ok {
		v := productcourses.DefaultDelete
		pcc.mutation.SetDelete(v)
	}
	if _, ok := pcc.mutation.CreatedID(); !ok {
		v := productcourses.DefaultCreatedID
		pcc.mutation.SetCreatedID(v)
	}
	if _, ok := pcc.mutation.Status(); !ok {
		v := productcourses.DefaultStatus
		pcc.mutation.SetStatus(v)
	}
	if _, ok := pcc.mutation.GetType(); !ok {
		v := productcourses.DefaultType
		pcc.mutation.SetType(v)
	}
	if _, ok := pcc.mutation.Name(); !ok {
		v := productcourses.DefaultName
		pcc.mutation.SetName(v)
	}
	if _, ok := pcc.mutation.Number(); !ok {
		v := productcourses.DefaultNumber
		pcc.mutation.SetNumber(v)
	}
	if _, ok := pcc.mutation.ProductID(); !ok {
		v := productcourses.DefaultProductID
		pcc.mutation.SetProductID(v)
	}
	if _, ok := pcc.mutation.CoursesID(); !ok {
		v := productcourses.DefaultCoursesID
		pcc.mutation.SetCoursesID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *ProductCoursesCreate) check() error {
	return nil
}

func (pcc *ProductCoursesCreate) sqlSave(ctx context.Context) (*ProductCourses, error) {
	if err := pcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	pcc.mutation.id = &_node.ID
	pcc.mutation.done = true
	return _node, nil
}

func (pcc *ProductCoursesCreate) createSpec() (*ProductCourses, *sqlgraph.CreateSpec) {
	var (
		_node = &ProductCourses{config: pcc.config}
		_spec = sqlgraph.NewCreateSpec(productcourses.Table, sqlgraph.NewFieldSpec(productcourses.FieldID, field.TypeInt64))
	)
	if id, ok := pcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pcc.mutation.CreatedAt(); ok {
		_spec.SetField(productcourses.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pcc.mutation.UpdatedAt(); ok {
		_spec.SetField(productcourses.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pcc.mutation.Delete(); ok {
		_spec.SetField(productcourses.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := pcc.mutation.CreatedID(); ok {
		_spec.SetField(productcourses.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := pcc.mutation.Status(); ok {
		_spec.SetField(productcourses.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := pcc.mutation.GetType(); ok {
		_spec.SetField(productcourses.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := pcc.mutation.Name(); ok {
		_spec.SetField(productcourses.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pcc.mutation.Number(); ok {
		_spec.SetField(productcourses.FieldNumber, field.TypeInt64, value)
		_node.Number = value
	}
	if value, ok := pcc.mutation.CoursesID(); ok {
		_spec.SetField(productcourses.FieldCoursesID, field.TypeInt64, value)
		_node.CoursesID = value
	}
	if nodes := pcc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.ProductTable,
			Columns: []string{productcourses.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ProductID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCoursesCreateBulk is the builder for creating many ProductCourses entities in bulk.
type ProductCoursesCreateBulk struct {
	config
	err      error
	builders []*ProductCoursesCreate
}

// Save creates the ProductCourses entities in the database.
func (pccb *ProductCoursesCreateBulk) Save(ctx context.Context) ([]*ProductCourses, error) {
	if pccb.err != nil {
		return nil, pccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*ProductCourses, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductCoursesMutation)
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
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *ProductCoursesCreateBulk) SaveX(ctx context.Context) []*ProductCourses {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *ProductCoursesCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *ProductCoursesCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
