// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/contract"
	"saas/biz/dal/db/ent/dictionarydetail"
	"saas/biz/dal/db/ent/product"
	"saas/idl_gen/model/base"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCreate is the builder for creating a Product entity.
type ProductCreate struct {
	config
	mutation *ProductMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProductCreate) SetCreatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *ProductCreate) SetUpdatedAt(t time.Time) *ProductCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableUpdatedAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDelete sets the "delete" field.
func (pc *ProductCreate) SetDelete(i int64) *ProductCreate {
	pc.mutation.SetDelete(i)
	return pc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (pc *ProductCreate) SetNillableDelete(i *int64) *ProductCreate {
	if i != nil {
		pc.SetDelete(*i)
	}
	return pc
}

// SetCreatedID sets the "created_id" field.
func (pc *ProductCreate) SetCreatedID(i int64) *ProductCreate {
	pc.mutation.SetCreatedID(i)
	return pc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (pc *ProductCreate) SetNillableCreatedID(i *int64) *ProductCreate {
	if i != nil {
		pc.SetCreatedID(*i)
	}
	return pc
}

// SetStatus sets the "status" field.
func (pc *ProductCreate) SetStatus(i int64) *ProductCreate {
	pc.mutation.SetStatus(i)
	return pc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pc *ProductCreate) SetNillableStatus(i *int64) *ProductCreate {
	if i != nil {
		pc.SetStatus(*i)
	}
	return pc
}

// SetType sets the "type" field.
func (pc *ProductCreate) SetType(s string) *ProductCreate {
	pc.mutation.SetType(s)
	return pc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pc *ProductCreate) SetNillableType(s *string) *ProductCreate {
	if s != nil {
		pc.SetType(*s)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *ProductCreate) SetName(s string) *ProductCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *ProductCreate) SetNillableName(s *string) *ProductCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetStock sets the "stock" field.
func (pc *ProductCreate) SetStock(i int64) *ProductCreate {
	pc.mutation.SetStock(i)
	return pc
}

// SetNillableStock sets the "stock" field if the given value is not nil.
func (pc *ProductCreate) SetNillableStock(i *int64) *ProductCreate {
	if i != nil {
		pc.SetStock(*i)
	}
	return pc
}

// SetDeadline sets the "deadline" field.
func (pc *ProductCreate) SetDeadline(i int64) *ProductCreate {
	pc.mutation.SetDeadline(i)
	return pc
}

// SetNillableDeadline sets the "deadline" field if the given value is not nil.
func (pc *ProductCreate) SetNillableDeadline(i *int64) *ProductCreate {
	if i != nil {
		pc.SetDeadline(*i)
	}
	return pc
}

// SetDuration sets the "duration" field.
func (pc *ProductCreate) SetDuration(i int64) *ProductCreate {
	pc.mutation.SetDuration(i)
	return pc
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (pc *ProductCreate) SetNillableDuration(i *int64) *ProductCreate {
	if i != nil {
		pc.SetDuration(*i)
	}
	return pc
}

// SetLength sets the "length" field.
func (pc *ProductCreate) SetLength(i int64) *ProductCreate {
	pc.mutation.SetLength(i)
	return pc
}

// SetNillableLength sets the "length" field if the given value is not nil.
func (pc *ProductCreate) SetNillableLength(i *int64) *ProductCreate {
	if i != nil {
		pc.SetLength(*i)
	}
	return pc
}

// SetSales sets the "sales" field.
func (pc *ProductCreate) SetSales(b []*base.Sales) *ProductCreate {
	pc.mutation.SetSales(b)
	return pc
}

// SetIsSales sets the "is_sales" field.
func (pc *ProductCreate) SetIsSales(i int64) *ProductCreate {
	pc.mutation.SetIsSales(i)
	return pc
}

// SetNillableIsSales sets the "is_sales" field if the given value is not nil.
func (pc *ProductCreate) SetNillableIsSales(i *int64) *ProductCreate {
	if i != nil {
		pc.SetIsSales(*i)
	}
	return pc
}

// SetSignSalesAt sets the "sign_sales_at" field.
func (pc *ProductCreate) SetSignSalesAt(t time.Time) *ProductCreate {
	pc.mutation.SetSignSalesAt(t)
	return pc
}

// SetNillableSignSalesAt sets the "sign_sales_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableSignSalesAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetSignSalesAt(*t)
	}
	return pc
}

// SetEndSalesAt sets the "end_sales_at" field.
func (pc *ProductCreate) SetEndSalesAt(t time.Time) *ProductCreate {
	pc.mutation.SetEndSalesAt(t)
	return pc
}

// SetNillableEndSalesAt sets the "end_sales_at" field if the given value is not nil.
func (pc *ProductCreate) SetNillableEndSalesAt(t *time.Time) *ProductCreate {
	if t != nil {
		pc.SetEndSalesAt(*t)
	}
	return pc
}

// SetPic sets the "pic" field.
func (pc *ProductCreate) SetPic(s string) *ProductCreate {
	pc.mutation.SetPic(s)
	return pc
}

// SetNillablePic sets the "pic" field if the given value is not nil.
func (pc *ProductCreate) SetNillablePic(s *string) *ProductCreate {
	if s != nil {
		pc.SetPic(*s)
	}
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProductCreate) SetDescription(s string) *ProductCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProductCreate) SetNillableDescription(s *string) *ProductCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProductCreate) SetID(i int64) *ProductCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddTagIDs adds the "tag" edge to the DictionaryDetail entity by IDs.
func (pc *ProductCreate) AddTagIDs(ids ...int64) *ProductCreate {
	pc.mutation.AddTagIDs(ids...)
	return pc
}

// AddTag adds the "tag" edges to the DictionaryDetail entity.
func (pc *ProductCreate) AddTag(d ...*DictionaryDetail) *ProductCreate {
	ids := make([]int64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddTagIDs(ids...)
}

// AddContractIDs adds the "contracts" edge to the Contract entity by IDs.
func (pc *ProductCreate) AddContractIDs(ids ...int64) *ProductCreate {
	pc.mutation.AddContractIDs(ids...)
	return pc
}

// AddContracts adds the "contracts" edges to the Contract entity.
func (pc *ProductCreate) AddContracts(c ...*Contract) *ProductCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddContractIDs(ids...)
}

// Mutation returns the ProductMutation object of the builder.
func (pc *ProductCreate) Mutation() *ProductMutation {
	return pc.mutation
}

// Save creates the Product in the database.
func (pc *ProductCreate) Save(ctx context.Context) (*Product, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProductCreate) SaveX(ctx context.Context) *Product {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProductCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProductCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProductCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := product.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := product.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Delete(); !ok {
		v := product.DefaultDelete
		pc.mutation.SetDelete(v)
	}
	if _, ok := pc.mutation.CreatedID(); !ok {
		v := product.DefaultCreatedID
		pc.mutation.SetCreatedID(v)
	}
	if _, ok := pc.mutation.Status(); !ok {
		v := product.DefaultStatus
		pc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProductCreate) check() error {
	return nil
}

func (pc *ProductCreate) sqlSave(ctx context.Context) (*Product, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProductCreate) createSpec() (*Product, *sqlgraph.CreateSpec) {
	var (
		_node = &Product{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(product.Table, sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(product.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(product.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Delete(); ok {
		_spec.SetField(product.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := pc.mutation.CreatedID(); ok {
		_spec.SetField(product.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(product.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := pc.mutation.GetType(); ok {
		_spec.SetField(product.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(product.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Stock(); ok {
		_spec.SetField(product.FieldStock, field.TypeInt64, value)
		_node.Stock = value
	}
	if value, ok := pc.mutation.Deadline(); ok {
		_spec.SetField(product.FieldDeadline, field.TypeInt64, value)
		_node.Deadline = value
	}
	if value, ok := pc.mutation.Duration(); ok {
		_spec.SetField(product.FieldDuration, field.TypeInt64, value)
		_node.Duration = value
	}
	if value, ok := pc.mutation.Length(); ok {
		_spec.SetField(product.FieldLength, field.TypeInt64, value)
		_node.Length = value
	}
	if value, ok := pc.mutation.Sales(); ok {
		_spec.SetField(product.FieldSales, field.TypeJSON, value)
		_node.Sales = value
	}
	if value, ok := pc.mutation.IsSales(); ok {
		_spec.SetField(product.FieldIsSales, field.TypeInt64, value)
		_node.IsSales = value
	}
	if value, ok := pc.mutation.SignSalesAt(); ok {
		_spec.SetField(product.FieldSignSalesAt, field.TypeTime, value)
		_node.SignSalesAt = value
	}
	if value, ok := pc.mutation.EndSalesAt(); ok {
		_spec.SetField(product.FieldEndSalesAt, field.TypeTime, value)
		_node.EndSalesAt = value
	}
	if value, ok := pc.mutation.Pic(); ok {
		_spec.SetField(product.FieldPic, field.TypeString, value)
		_node.Pic = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(product.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := pc.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   product.TagTable,
			Columns: []string{product.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ContractsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   product.ContractsTable,
			Columns: product.ContractsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contract.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProductCreateBulk is the builder for creating many Product entities in bulk.
type ProductCreateBulk struct {
	config
	err      error
	builders []*ProductCreate
}

// Save creates the Product entities in the database.
func (pcb *ProductCreateBulk) Save(ctx context.Context) ([]*Product, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Product, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProductMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProductCreateBulk) SaveX(ctx context.Context) []*Product {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProductCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProductCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
