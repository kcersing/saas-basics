// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/product"
	"saas/biz/dal/db/ent/productcourses"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProductCoursesUpdate is the builder for updating ProductCourses entities.
type ProductCoursesUpdate struct {
	config
	hooks    []Hook
	mutation *ProductCoursesMutation
}

// Where appends a list predicates to the ProductCoursesUpdate builder.
func (pcu *ProductCoursesUpdate) Where(ps ...predicate.ProductCourses) *ProductCoursesUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetUpdatedAt sets the "updated_at" field.
func (pcu *ProductCoursesUpdate) SetUpdatedAt(t time.Time) *ProductCoursesUpdate {
	pcu.mutation.SetUpdatedAt(t)
	return pcu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pcu *ProductCoursesUpdate) ClearUpdatedAt() *ProductCoursesUpdate {
	pcu.mutation.ClearUpdatedAt()
	return pcu
}

// SetDelete sets the "delete" field.
func (pcu *ProductCoursesUpdate) SetDelete(i int64) *ProductCoursesUpdate {
	pcu.mutation.ResetDelete()
	pcu.mutation.SetDelete(i)
	return pcu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableDelete(i *int64) *ProductCoursesUpdate {
	if i != nil {
		pcu.SetDelete(*i)
	}
	return pcu
}

// AddDelete adds i to the "delete" field.
func (pcu *ProductCoursesUpdate) AddDelete(i int64) *ProductCoursesUpdate {
	pcu.mutation.AddDelete(i)
	return pcu
}

// ClearDelete clears the value of the "delete" field.
func (pcu *ProductCoursesUpdate) ClearDelete() *ProductCoursesUpdate {
	pcu.mutation.ClearDelete()
	return pcu
}

// SetCreatedID sets the "created_id" field.
func (pcu *ProductCoursesUpdate) SetCreatedID(i int64) *ProductCoursesUpdate {
	pcu.mutation.ResetCreatedID()
	pcu.mutation.SetCreatedID(i)
	return pcu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableCreatedID(i *int64) *ProductCoursesUpdate {
	if i != nil {
		pcu.SetCreatedID(*i)
	}
	return pcu
}

// AddCreatedID adds i to the "created_id" field.
func (pcu *ProductCoursesUpdate) AddCreatedID(i int64) *ProductCoursesUpdate {
	pcu.mutation.AddCreatedID(i)
	return pcu
}

// ClearCreatedID clears the value of the "created_id" field.
func (pcu *ProductCoursesUpdate) ClearCreatedID() *ProductCoursesUpdate {
	pcu.mutation.ClearCreatedID()
	return pcu
}

// SetStatus sets the "status" field.
func (pcu *ProductCoursesUpdate) SetStatus(i int64) *ProductCoursesUpdate {
	pcu.mutation.ResetStatus()
	pcu.mutation.SetStatus(i)
	return pcu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableStatus(i *int64) *ProductCoursesUpdate {
	if i != nil {
		pcu.SetStatus(*i)
	}
	return pcu
}

// AddStatus adds i to the "status" field.
func (pcu *ProductCoursesUpdate) AddStatus(i int64) *ProductCoursesUpdate {
	pcu.mutation.AddStatus(i)
	return pcu
}

// ClearStatus clears the value of the "status" field.
func (pcu *ProductCoursesUpdate) ClearStatus() *ProductCoursesUpdate {
	pcu.mutation.ClearStatus()
	return pcu
}

// SetType sets the "type" field.
func (pcu *ProductCoursesUpdate) SetType(s string) *ProductCoursesUpdate {
	pcu.mutation.SetType(s)
	return pcu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableType(s *string) *ProductCoursesUpdate {
	if s != nil {
		pcu.SetType(*s)
	}
	return pcu
}

// ClearType clears the value of the "type" field.
func (pcu *ProductCoursesUpdate) ClearType() *ProductCoursesUpdate {
	pcu.mutation.ClearType()
	return pcu
}

// SetName sets the "name" field.
func (pcu *ProductCoursesUpdate) SetName(s string) *ProductCoursesUpdate {
	pcu.mutation.SetName(s)
	return pcu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableName(s *string) *ProductCoursesUpdate {
	if s != nil {
		pcu.SetName(*s)
	}
	return pcu
}

// ClearName clears the value of the "name" field.
func (pcu *ProductCoursesUpdate) ClearName() *ProductCoursesUpdate {
	pcu.mutation.ClearName()
	return pcu
}

// SetProductID sets the "product_id" field.
func (pcu *ProductCoursesUpdate) SetProductID(i int64) *ProductCoursesUpdate {
	pcu.mutation.SetProductID(i)
	return pcu
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableProductID(i *int64) *ProductCoursesUpdate {
	if i != nil {
		pcu.SetProductID(*i)
	}
	return pcu
}

// ClearProductID clears the value of the "product_id" field.
func (pcu *ProductCoursesUpdate) ClearProductID() *ProductCoursesUpdate {
	pcu.mutation.ClearProductID()
	return pcu
}

// SetCoursesID sets the "courses_id" field.
func (pcu *ProductCoursesUpdate) SetCoursesID(i int64) *ProductCoursesUpdate {
	pcu.mutation.ResetCoursesID()
	pcu.mutation.SetCoursesID(i)
	return pcu
}

// SetNillableCoursesID sets the "courses_id" field if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableCoursesID(i *int64) *ProductCoursesUpdate {
	if i != nil {
		pcu.SetCoursesID(*i)
	}
	return pcu
}

// AddCoursesID adds i to the "courses_id" field.
func (pcu *ProductCoursesUpdate) AddCoursesID(i int64) *ProductCoursesUpdate {
	pcu.mutation.AddCoursesID(i)
	return pcu
}

// ClearCoursesID clears the value of the "courses_id" field.
func (pcu *ProductCoursesUpdate) ClearCoursesID() *ProductCoursesUpdate {
	pcu.mutation.ClearCoursesID()
	return pcu
}

// SetNodeCID sets the "nodeC" edge to the Product entity by ID.
func (pcu *ProductCoursesUpdate) SetNodeCID(id int64) *ProductCoursesUpdate {
	pcu.mutation.SetNodeCID(id)
	return pcu
}

// SetNillableNodeCID sets the "nodeC" edge to the Product entity by ID if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableNodeCID(id *int64) *ProductCoursesUpdate {
	if id != nil {
		pcu = pcu.SetNodeCID(*id)
	}
	return pcu
}

// SetNodeC sets the "nodeC" edge to the Product entity.
func (pcu *ProductCoursesUpdate) SetNodeC(p *Product) *ProductCoursesUpdate {
	return pcu.SetNodeCID(p.ID)
}

// SetNodeLID sets the "nodeL" edge to the Product entity by ID.
func (pcu *ProductCoursesUpdate) SetNodeLID(id int64) *ProductCoursesUpdate {
	pcu.mutation.SetNodeLID(id)
	return pcu
}

// SetNillableNodeLID sets the "nodeL" edge to the Product entity by ID if the given value is not nil.
func (pcu *ProductCoursesUpdate) SetNillableNodeLID(id *int64) *ProductCoursesUpdate {
	if id != nil {
		pcu = pcu.SetNodeLID(*id)
	}
	return pcu
}

// SetNodeL sets the "nodeL" edge to the Product entity.
func (pcu *ProductCoursesUpdate) SetNodeL(p *Product) *ProductCoursesUpdate {
	return pcu.SetNodeLID(p.ID)
}

// Mutation returns the ProductCoursesMutation object of the builder.
func (pcu *ProductCoursesUpdate) Mutation() *ProductCoursesMutation {
	return pcu.mutation
}

// ClearNodeC clears the "nodeC" edge to the Product entity.
func (pcu *ProductCoursesUpdate) ClearNodeC() *ProductCoursesUpdate {
	pcu.mutation.ClearNodeC()
	return pcu
}

// ClearNodeL clears the "nodeL" edge to the Product entity.
func (pcu *ProductCoursesUpdate) ClearNodeL() *ProductCoursesUpdate {
	pcu.mutation.ClearNodeL()
	return pcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProductCoursesUpdate) Save(ctx context.Context) (int, error) {
	pcu.defaults()
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProductCoursesUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProductCoursesUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProductCoursesUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcu *ProductCoursesUpdate) defaults() {
	if _, ok := pcu.mutation.UpdatedAt(); !ok && !pcu.mutation.UpdatedAtCleared() {
		v := productcourses.UpdateDefaultUpdatedAt()
		pcu.mutation.SetUpdatedAt(v)
	}
}

func (pcu *ProductCoursesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(productcourses.Table, productcourses.Columns, sqlgraph.NewFieldSpec(productcourses.FieldID, field.TypeInt64))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pcu.mutation.CreatedAtCleared() {
		_spec.ClearField(productcourses.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := pcu.mutation.UpdatedAt(); ok {
		_spec.SetField(productcourses.FieldUpdatedAt, field.TypeTime, value)
	}
	if pcu.mutation.UpdatedAtCleared() {
		_spec.ClearField(productcourses.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pcu.mutation.Delete(); ok {
		_spec.SetField(productcourses.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := pcu.mutation.AddedDelete(); ok {
		_spec.AddField(productcourses.FieldDelete, field.TypeInt64, value)
	}
	if pcu.mutation.DeleteCleared() {
		_spec.ClearField(productcourses.FieldDelete, field.TypeInt64)
	}
	if value, ok := pcu.mutation.CreatedID(); ok {
		_spec.SetField(productcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := pcu.mutation.AddedCreatedID(); ok {
		_spec.AddField(productcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if pcu.mutation.CreatedIDCleared() {
		_spec.ClearField(productcourses.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := pcu.mutation.Status(); ok {
		_spec.SetField(productcourses.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := pcu.mutation.AddedStatus(); ok {
		_spec.AddField(productcourses.FieldStatus, field.TypeInt64, value)
	}
	if pcu.mutation.StatusCleared() {
		_spec.ClearField(productcourses.FieldStatus, field.TypeInt64)
	}
	if value, ok := pcu.mutation.GetType(); ok {
		_spec.SetField(productcourses.FieldType, field.TypeString, value)
	}
	if pcu.mutation.TypeCleared() {
		_spec.ClearField(productcourses.FieldType, field.TypeString)
	}
	if value, ok := pcu.mutation.Name(); ok {
		_spec.SetField(productcourses.FieldName, field.TypeString, value)
	}
	if pcu.mutation.NameCleared() {
		_spec.ClearField(productcourses.FieldName, field.TypeString)
	}
	if value, ok := pcu.mutation.CoursesID(); ok {
		_spec.SetField(productcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if value, ok := pcu.mutation.AddedCoursesID(); ok {
		_spec.AddField(productcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if pcu.mutation.CoursesIDCleared() {
		_spec.ClearField(productcourses.FieldCoursesID, field.TypeInt64)
	}
	if pcu.mutation.NodeCCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeCTable,
			Columns: []string{productcourses.NodeCColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.NodeCIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeCTable,
			Columns: []string{productcourses.NodeCColumn},
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
	if pcu.mutation.NodeLCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeLTable,
			Columns: []string{productcourses.NodeLColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.NodeLIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeLTable,
			Columns: []string{productcourses.NodeLColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcourses.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// ProductCoursesUpdateOne is the builder for updating a single ProductCourses entity.
type ProductCoursesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProductCoursesMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (pcuo *ProductCoursesUpdateOne) SetUpdatedAt(t time.Time) *ProductCoursesUpdateOne {
	pcuo.mutation.SetUpdatedAt(t)
	return pcuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pcuo *ProductCoursesUpdateOne) ClearUpdatedAt() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearUpdatedAt()
	return pcuo
}

// SetDelete sets the "delete" field.
func (pcuo *ProductCoursesUpdateOne) SetDelete(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.ResetDelete()
	pcuo.mutation.SetDelete(i)
	return pcuo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableDelete(i *int64) *ProductCoursesUpdateOne {
	if i != nil {
		pcuo.SetDelete(*i)
	}
	return pcuo
}

// AddDelete adds i to the "delete" field.
func (pcuo *ProductCoursesUpdateOne) AddDelete(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.AddDelete(i)
	return pcuo
}

// ClearDelete clears the value of the "delete" field.
func (pcuo *ProductCoursesUpdateOne) ClearDelete() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearDelete()
	return pcuo
}

// SetCreatedID sets the "created_id" field.
func (pcuo *ProductCoursesUpdateOne) SetCreatedID(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.ResetCreatedID()
	pcuo.mutation.SetCreatedID(i)
	return pcuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableCreatedID(i *int64) *ProductCoursesUpdateOne {
	if i != nil {
		pcuo.SetCreatedID(*i)
	}
	return pcuo
}

// AddCreatedID adds i to the "created_id" field.
func (pcuo *ProductCoursesUpdateOne) AddCreatedID(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.AddCreatedID(i)
	return pcuo
}

// ClearCreatedID clears the value of the "created_id" field.
func (pcuo *ProductCoursesUpdateOne) ClearCreatedID() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearCreatedID()
	return pcuo
}

// SetStatus sets the "status" field.
func (pcuo *ProductCoursesUpdateOne) SetStatus(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.ResetStatus()
	pcuo.mutation.SetStatus(i)
	return pcuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableStatus(i *int64) *ProductCoursesUpdateOne {
	if i != nil {
		pcuo.SetStatus(*i)
	}
	return pcuo
}

// AddStatus adds i to the "status" field.
func (pcuo *ProductCoursesUpdateOne) AddStatus(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.AddStatus(i)
	return pcuo
}

// ClearStatus clears the value of the "status" field.
func (pcuo *ProductCoursesUpdateOne) ClearStatus() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearStatus()
	return pcuo
}

// SetType sets the "type" field.
func (pcuo *ProductCoursesUpdateOne) SetType(s string) *ProductCoursesUpdateOne {
	pcuo.mutation.SetType(s)
	return pcuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableType(s *string) *ProductCoursesUpdateOne {
	if s != nil {
		pcuo.SetType(*s)
	}
	return pcuo
}

// ClearType clears the value of the "type" field.
func (pcuo *ProductCoursesUpdateOne) ClearType() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearType()
	return pcuo
}

// SetName sets the "name" field.
func (pcuo *ProductCoursesUpdateOne) SetName(s string) *ProductCoursesUpdateOne {
	pcuo.mutation.SetName(s)
	return pcuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableName(s *string) *ProductCoursesUpdateOne {
	if s != nil {
		pcuo.SetName(*s)
	}
	return pcuo
}

// ClearName clears the value of the "name" field.
func (pcuo *ProductCoursesUpdateOne) ClearName() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearName()
	return pcuo
}

// SetProductID sets the "product_id" field.
func (pcuo *ProductCoursesUpdateOne) SetProductID(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.SetProductID(i)
	return pcuo
}

// SetNillableProductID sets the "product_id" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableProductID(i *int64) *ProductCoursesUpdateOne {
	if i != nil {
		pcuo.SetProductID(*i)
	}
	return pcuo
}

// ClearProductID clears the value of the "product_id" field.
func (pcuo *ProductCoursesUpdateOne) ClearProductID() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearProductID()
	return pcuo
}

// SetCoursesID sets the "courses_id" field.
func (pcuo *ProductCoursesUpdateOne) SetCoursesID(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.ResetCoursesID()
	pcuo.mutation.SetCoursesID(i)
	return pcuo
}

// SetNillableCoursesID sets the "courses_id" field if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableCoursesID(i *int64) *ProductCoursesUpdateOne {
	if i != nil {
		pcuo.SetCoursesID(*i)
	}
	return pcuo
}

// AddCoursesID adds i to the "courses_id" field.
func (pcuo *ProductCoursesUpdateOne) AddCoursesID(i int64) *ProductCoursesUpdateOne {
	pcuo.mutation.AddCoursesID(i)
	return pcuo
}

// ClearCoursesID clears the value of the "courses_id" field.
func (pcuo *ProductCoursesUpdateOne) ClearCoursesID() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearCoursesID()
	return pcuo
}

// SetNodeCID sets the "nodeC" edge to the Product entity by ID.
func (pcuo *ProductCoursesUpdateOne) SetNodeCID(id int64) *ProductCoursesUpdateOne {
	pcuo.mutation.SetNodeCID(id)
	return pcuo
}

// SetNillableNodeCID sets the "nodeC" edge to the Product entity by ID if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableNodeCID(id *int64) *ProductCoursesUpdateOne {
	if id != nil {
		pcuo = pcuo.SetNodeCID(*id)
	}
	return pcuo
}

// SetNodeC sets the "nodeC" edge to the Product entity.
func (pcuo *ProductCoursesUpdateOne) SetNodeC(p *Product) *ProductCoursesUpdateOne {
	return pcuo.SetNodeCID(p.ID)
}

// SetNodeLID sets the "nodeL" edge to the Product entity by ID.
func (pcuo *ProductCoursesUpdateOne) SetNodeLID(id int64) *ProductCoursesUpdateOne {
	pcuo.mutation.SetNodeLID(id)
	return pcuo
}

// SetNillableNodeLID sets the "nodeL" edge to the Product entity by ID if the given value is not nil.
func (pcuo *ProductCoursesUpdateOne) SetNillableNodeLID(id *int64) *ProductCoursesUpdateOne {
	if id != nil {
		pcuo = pcuo.SetNodeLID(*id)
	}
	return pcuo
}

// SetNodeL sets the "nodeL" edge to the Product entity.
func (pcuo *ProductCoursesUpdateOne) SetNodeL(p *Product) *ProductCoursesUpdateOne {
	return pcuo.SetNodeLID(p.ID)
}

// Mutation returns the ProductCoursesMutation object of the builder.
func (pcuo *ProductCoursesUpdateOne) Mutation() *ProductCoursesMutation {
	return pcuo.mutation
}

// ClearNodeC clears the "nodeC" edge to the Product entity.
func (pcuo *ProductCoursesUpdateOne) ClearNodeC() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearNodeC()
	return pcuo
}

// ClearNodeL clears the "nodeL" edge to the Product entity.
func (pcuo *ProductCoursesUpdateOne) ClearNodeL() *ProductCoursesUpdateOne {
	pcuo.mutation.ClearNodeL()
	return pcuo
}

// Where appends a list predicates to the ProductCoursesUpdate builder.
func (pcuo *ProductCoursesUpdateOne) Where(ps ...predicate.ProductCourses) *ProductCoursesUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProductCoursesUpdateOne) Select(field string, fields ...string) *ProductCoursesUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProductCourses entity.
func (pcuo *ProductCoursesUpdateOne) Save(ctx context.Context) (*ProductCourses, error) {
	pcuo.defaults()
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProductCoursesUpdateOne) SaveX(ctx context.Context) *ProductCourses {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProductCoursesUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProductCoursesUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcuo *ProductCoursesUpdateOne) defaults() {
	if _, ok := pcuo.mutation.UpdatedAt(); !ok && !pcuo.mutation.UpdatedAtCleared() {
		v := productcourses.UpdateDefaultUpdatedAt()
		pcuo.mutation.SetUpdatedAt(v)
	}
}

func (pcuo *ProductCoursesUpdateOne) sqlSave(ctx context.Context) (_node *ProductCourses, err error) {
	_spec := sqlgraph.NewUpdateSpec(productcourses.Table, productcourses.Columns, sqlgraph.NewFieldSpec(productcourses.FieldID, field.TypeInt64))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProductCourses.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, productcourses.FieldID)
		for _, f := range fields {
			if !productcourses.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != productcourses.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pcuo.mutation.CreatedAtCleared() {
		_spec.ClearField(productcourses.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := pcuo.mutation.UpdatedAt(); ok {
		_spec.SetField(productcourses.FieldUpdatedAt, field.TypeTime, value)
	}
	if pcuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(productcourses.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pcuo.mutation.Delete(); ok {
		_spec.SetField(productcourses.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := pcuo.mutation.AddedDelete(); ok {
		_spec.AddField(productcourses.FieldDelete, field.TypeInt64, value)
	}
	if pcuo.mutation.DeleteCleared() {
		_spec.ClearField(productcourses.FieldDelete, field.TypeInt64)
	}
	if value, ok := pcuo.mutation.CreatedID(); ok {
		_spec.SetField(productcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := pcuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(productcourses.FieldCreatedID, field.TypeInt64, value)
	}
	if pcuo.mutation.CreatedIDCleared() {
		_spec.ClearField(productcourses.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := pcuo.mutation.Status(); ok {
		_spec.SetField(productcourses.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := pcuo.mutation.AddedStatus(); ok {
		_spec.AddField(productcourses.FieldStatus, field.TypeInt64, value)
	}
	if pcuo.mutation.StatusCleared() {
		_spec.ClearField(productcourses.FieldStatus, field.TypeInt64)
	}
	if value, ok := pcuo.mutation.GetType(); ok {
		_spec.SetField(productcourses.FieldType, field.TypeString, value)
	}
	if pcuo.mutation.TypeCleared() {
		_spec.ClearField(productcourses.FieldType, field.TypeString)
	}
	if value, ok := pcuo.mutation.Name(); ok {
		_spec.SetField(productcourses.FieldName, field.TypeString, value)
	}
	if pcuo.mutation.NameCleared() {
		_spec.ClearField(productcourses.FieldName, field.TypeString)
	}
	if value, ok := pcuo.mutation.CoursesID(); ok {
		_spec.SetField(productcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if value, ok := pcuo.mutation.AddedCoursesID(); ok {
		_spec.AddField(productcourses.FieldCoursesID, field.TypeInt64, value)
	}
	if pcuo.mutation.CoursesIDCleared() {
		_spec.ClearField(productcourses.FieldCoursesID, field.TypeInt64)
	}
	if pcuo.mutation.NodeCCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeCTable,
			Columns: []string{productcourses.NodeCColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.NodeCIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeCTable,
			Columns: []string{productcourses.NodeCColumn},
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
	if pcuo.mutation.NodeLCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeLTable,
			Columns: []string{productcourses.NodeLColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(product.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.NodeLIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   productcourses.NodeLTable,
			Columns: []string{productcourses.NodeLColumn},
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
	_node = &ProductCourses{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{productcourses.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}
