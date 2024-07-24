// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/menuparam"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/role"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MenuUpdate) SetUpdatedAt(t time.Time) *MenuUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MenuUpdate) SetParentID(i int64) *MenuUpdate {
	mu.mutation.SetParentID(i)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(i *int64) *MenuUpdate {
	if i != nil {
		mu.SetParentID(*i)
	}
	return mu
}

// ClearParentID clears the value of the "parent_id" field.
func (mu *MenuUpdate) ClearParentID() *MenuUpdate {
	mu.mutation.ClearParentID()
	return mu
}

// SetPath sets the "path" field.
func (mu *MenuUpdate) SetPath(s string) *MenuUpdate {
	mu.mutation.SetPath(s)
	return mu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mu *MenuUpdate) SetNillablePath(s *string) *MenuUpdate {
	if s != nil {
		mu.SetPath(*s)
	}
	return mu
}

// ClearPath clears the value of the "path" field.
func (mu *MenuUpdate) ClearPath() *MenuUpdate {
	mu.mutation.ClearPath()
	return mu
}

// SetName sets the "name" field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetName(*s)
	}
	return mu
}

// SetOrderNo sets the "order_no" field.
func (mu *MenuUpdate) SetOrderNo(i int32) *MenuUpdate {
	mu.mutation.ResetOrderNo()
	mu.mutation.SetOrderNo(i)
	return mu
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableOrderNo(i *int32) *MenuUpdate {
	if i != nil {
		mu.SetOrderNo(*i)
	}
	return mu
}

// AddOrderNo adds i to the "order_no" field.
func (mu *MenuUpdate) AddOrderNo(i int32) *MenuUpdate {
	mu.mutation.AddOrderNo(i)
	return mu
}

// SetDisabled sets the "disabled" field.
func (mu *MenuUpdate) SetDisabled(i int32) *MenuUpdate {
	mu.mutation.ResetDisabled()
	mu.mutation.SetDisabled(i)
	return mu
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableDisabled(i *int32) *MenuUpdate {
	if i != nil {
		mu.SetDisabled(*i)
	}
	return mu
}

// AddDisabled adds i to the "disabled" field.
func (mu *MenuUpdate) AddDisabled(i int32) *MenuUpdate {
	mu.mutation.AddDisabled(i)
	return mu
}

// ClearDisabled clears the value of the "disabled" field.
func (mu *MenuUpdate) ClearDisabled() *MenuUpdate {
	mu.mutation.ClearDisabled()
	return mu
}

// SetIgnore sets the "ignore" field.
func (mu *MenuUpdate) SetIgnore(b bool) *MenuUpdate {
	mu.mutation.SetIgnore(b)
	return mu
}

// SetNillableIgnore sets the "ignore" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableIgnore(b *bool) *MenuUpdate {
	if b != nil {
		mu.SetIgnore(*b)
	}
	return mu
}

// ClearIgnore clears the value of the "ignore" field.
func (mu *MenuUpdate) ClearIgnore() *MenuUpdate {
	mu.mutation.ClearIgnore()
	return mu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (mu *MenuUpdate) AddRoleIDs(ids ...int64) *MenuUpdate {
	mu.mutation.AddRoleIDs(ids...)
	return mu
}

// AddRoles adds the "roles" edges to the Role entity.
func (mu *MenuUpdate) AddRoles(r ...*Role) *MenuUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.AddRoleIDs(ids...)
}

// SetParent sets the "parent" edge to the Menu entity.
func (mu *MenuUpdate) SetParent(m *Menu) *MenuUpdate {
	return mu.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mu *MenuUpdate) AddChildIDs(ids ...int64) *MenuUpdate {
	mu.mutation.AddChildIDs(ids...)
	return mu
}

// AddChildren adds the "children" edges to the Menu entity.
func (mu *MenuUpdate) AddChildren(m ...*Menu) *MenuUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddChildIDs(ids...)
}

// AddParamIDs adds the "params" edge to the MenuParam entity by IDs.
func (mu *MenuUpdate) AddParamIDs(ids ...int64) *MenuUpdate {
	mu.mutation.AddParamIDs(ids...)
	return mu
}

// AddParams adds the "params" edges to the MenuParam entity.
func (mu *MenuUpdate) AddParams(m ...*MenuParam) *MenuUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddParamIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (mu *MenuUpdate) ClearRoles() *MenuUpdate {
	mu.mutation.ClearRoles()
	return mu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (mu *MenuUpdate) RemoveRoleIDs(ids ...int64) *MenuUpdate {
	mu.mutation.RemoveRoleIDs(ids...)
	return mu
}

// RemoveRoles removes "roles" edges to Role entities.
func (mu *MenuUpdate) RemoveRoles(r ...*Role) *MenuUpdate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mu.RemoveRoleIDs(ids...)
}

// ClearParent clears the "parent" edge to the Menu entity.
func (mu *MenuUpdate) ClearParent() *MenuUpdate {
	mu.mutation.ClearParent()
	return mu
}

// ClearChildren clears all "children" edges to the Menu entity.
func (mu *MenuUpdate) ClearChildren() *MenuUpdate {
	mu.mutation.ClearChildren()
	return mu
}

// RemoveChildIDs removes the "children" edge to Menu entities by IDs.
func (mu *MenuUpdate) RemoveChildIDs(ids ...int64) *MenuUpdate {
	mu.mutation.RemoveChildIDs(ids...)
	return mu
}

// RemoveChildren removes "children" edges to Menu entities.
func (mu *MenuUpdate) RemoveChildren(m ...*Menu) *MenuUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveChildIDs(ids...)
}

// ClearParams clears all "params" edges to the MenuParam entity.
func (mu *MenuUpdate) ClearParams() *MenuUpdate {
	mu.mutation.ClearParams()
	return mu
}

// RemoveParamIDs removes the "params" edge to MenuParam entities by IDs.
func (mu *MenuUpdate) RemoveParamIDs(ids ...int64) *MenuUpdate {
	mu.mutation.RemoveParamIDs(ids...)
	return mu
}

// RemoveParams removes "params" edges to MenuParam entities.
func (mu *MenuUpdate) RemoveParams(m ...*MenuParam) *MenuUpdate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveParamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MenuUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := menu.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (mu *MenuUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdate {
	mu.modifiers = append(mu.modifiers, modifiers...)
	return mu
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
	}
	if mu.mutation.PathCleared() {
		_spec.ClearField(menu.FieldPath, field.TypeString)
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if value, ok := mu.mutation.OrderNo(); ok {
		_spec.SetField(menu.FieldOrderNo, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedOrderNo(); ok {
		_spec.AddField(menu.FieldOrderNo, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.Disabled(); ok {
		_spec.SetField(menu.FieldDisabled, field.TypeInt32, value)
	}
	if value, ok := mu.mutation.AddedDisabled(); ok {
		_spec.AddField(menu.FieldDisabled, field.TypeInt32, value)
	}
	if mu.mutation.DisabledCleared() {
		_spec.ClearField(menu.FieldDisabled, field.TypeInt32)
	}
	if value, ok := mu.mutation.Ignore(); ok {
		_spec.SetField(menu.FieldIgnore, field.TypeBool, value)
	}
	if mu.mutation.IgnoreCleared() {
		_spec.ClearField(menu.FieldIgnore, field.TypeBool)
	}
	if mu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !mu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !mu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ParamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedParamsIDs(); len(nodes) > 0 && !mu.mutation.ParamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ParamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(mu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *MenuMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MenuUpdateOne) SetUpdatedAt(t time.Time) *MenuUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetParentID sets the "parent_id" field.
func (muo *MenuUpdateOne) SetParentID(i int64) *MenuUpdateOne {
	muo.mutation.SetParentID(i)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(i *int64) *MenuUpdateOne {
	if i != nil {
		muo.SetParentID(*i)
	}
	return muo
}

// ClearParentID clears the value of the "parent_id" field.
func (muo *MenuUpdateOne) ClearParentID() *MenuUpdateOne {
	muo.mutation.ClearParentID()
	return muo
}

// SetPath sets the "path" field.
func (muo *MenuUpdateOne) SetPath(s string) *MenuUpdateOne {
	muo.mutation.SetPath(s)
	return muo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillablePath(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetPath(*s)
	}
	return muo
}

// ClearPath clears the value of the "path" field.
func (muo *MenuUpdateOne) ClearPath() *MenuUpdateOne {
	muo.mutation.ClearPath()
	return muo
}

// SetName sets the "name" field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetName(*s)
	}
	return muo
}

// SetOrderNo sets the "order_no" field.
func (muo *MenuUpdateOne) SetOrderNo(i int32) *MenuUpdateOne {
	muo.mutation.ResetOrderNo()
	muo.mutation.SetOrderNo(i)
	return muo
}

// SetNillableOrderNo sets the "order_no" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableOrderNo(i *int32) *MenuUpdateOne {
	if i != nil {
		muo.SetOrderNo(*i)
	}
	return muo
}

// AddOrderNo adds i to the "order_no" field.
func (muo *MenuUpdateOne) AddOrderNo(i int32) *MenuUpdateOne {
	muo.mutation.AddOrderNo(i)
	return muo
}

// SetDisabled sets the "disabled" field.
func (muo *MenuUpdateOne) SetDisabled(i int32) *MenuUpdateOne {
	muo.mutation.ResetDisabled()
	muo.mutation.SetDisabled(i)
	return muo
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableDisabled(i *int32) *MenuUpdateOne {
	if i != nil {
		muo.SetDisabled(*i)
	}
	return muo
}

// AddDisabled adds i to the "disabled" field.
func (muo *MenuUpdateOne) AddDisabled(i int32) *MenuUpdateOne {
	muo.mutation.AddDisabled(i)
	return muo
}

// ClearDisabled clears the value of the "disabled" field.
func (muo *MenuUpdateOne) ClearDisabled() *MenuUpdateOne {
	muo.mutation.ClearDisabled()
	return muo
}

// SetIgnore sets the "ignore" field.
func (muo *MenuUpdateOne) SetIgnore(b bool) *MenuUpdateOne {
	muo.mutation.SetIgnore(b)
	return muo
}

// SetNillableIgnore sets the "ignore" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableIgnore(b *bool) *MenuUpdateOne {
	if b != nil {
		muo.SetIgnore(*b)
	}
	return muo
}

// ClearIgnore clears the value of the "ignore" field.
func (muo *MenuUpdateOne) ClearIgnore() *MenuUpdateOne {
	muo.mutation.ClearIgnore()
	return muo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (muo *MenuUpdateOne) AddRoleIDs(ids ...int64) *MenuUpdateOne {
	muo.mutation.AddRoleIDs(ids...)
	return muo
}

// AddRoles adds the "roles" edges to the Role entity.
func (muo *MenuUpdateOne) AddRoles(r ...*Role) *MenuUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.AddRoleIDs(ids...)
}

// SetParent sets the "parent" edge to the Menu entity.
func (muo *MenuUpdateOne) SetParent(m *Menu) *MenuUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (muo *MenuUpdateOne) AddChildIDs(ids ...int64) *MenuUpdateOne {
	muo.mutation.AddChildIDs(ids...)
	return muo
}

// AddChildren adds the "children" edges to the Menu entity.
func (muo *MenuUpdateOne) AddChildren(m ...*Menu) *MenuUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddChildIDs(ids...)
}

// AddParamIDs adds the "params" edge to the MenuParam entity by IDs.
func (muo *MenuUpdateOne) AddParamIDs(ids ...int64) *MenuUpdateOne {
	muo.mutation.AddParamIDs(ids...)
	return muo
}

// AddParams adds the "params" edges to the MenuParam entity.
func (muo *MenuUpdateOne) AddParams(m ...*MenuParam) *MenuUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddParamIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (muo *MenuUpdateOne) ClearRoles() *MenuUpdateOne {
	muo.mutation.ClearRoles()
	return muo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (muo *MenuUpdateOne) RemoveRoleIDs(ids ...int64) *MenuUpdateOne {
	muo.mutation.RemoveRoleIDs(ids...)
	return muo
}

// RemoveRoles removes "roles" edges to Role entities.
func (muo *MenuUpdateOne) RemoveRoles(r ...*Role) *MenuUpdateOne {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return muo.RemoveRoleIDs(ids...)
}

// ClearParent clears the "parent" edge to the Menu entity.
func (muo *MenuUpdateOne) ClearParent() *MenuUpdateOne {
	muo.mutation.ClearParent()
	return muo
}

// ClearChildren clears all "children" edges to the Menu entity.
func (muo *MenuUpdateOne) ClearChildren() *MenuUpdateOne {
	muo.mutation.ClearChildren()
	return muo
}

// RemoveChildIDs removes the "children" edge to Menu entities by IDs.
func (muo *MenuUpdateOne) RemoveChildIDs(ids ...int64) *MenuUpdateOne {
	muo.mutation.RemoveChildIDs(ids...)
	return muo
}

// RemoveChildren removes "children" edges to Menu entities.
func (muo *MenuUpdateOne) RemoveChildren(m ...*Menu) *MenuUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveChildIDs(ids...)
}

// ClearParams clears all "params" edges to the MenuParam entity.
func (muo *MenuUpdateOne) ClearParams() *MenuUpdateOne {
	muo.mutation.ClearParams()
	return muo
}

// RemoveParamIDs removes the "params" edge to MenuParam entities by IDs.
func (muo *MenuUpdateOne) RemoveParamIDs(ids ...int64) *MenuUpdateOne {
	muo.mutation.RemoveParamIDs(ids...)
	return muo
}

// RemoveParams removes "params" edges to MenuParam entities.
func (muo *MenuUpdateOne) RemoveParams(m ...*MenuParam) *MenuUpdateOne {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveParamIDs(ids...)
}

// Where appends a list predicates to the MenuUpdate builder.
func (muo *MenuUpdateOne) Where(ps ...predicate.Menu) *MenuUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MenuUpdateOne) Select(field string, fields ...string) *MenuUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Menu entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MenuUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := menu.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (muo *MenuUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *MenuUpdateOne {
	muo.modifiers = append(muo.modifiers, modifiers...)
	return muo
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Menu.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, menu.FieldID)
		for _, f := range fields {
			if !menu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != menu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
	}
	if muo.mutation.PathCleared() {
		_spec.ClearField(menu.FieldPath, field.TypeString)
	}
	if value, ok := muo.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
	}
	if value, ok := muo.mutation.OrderNo(); ok {
		_spec.SetField(menu.FieldOrderNo, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedOrderNo(); ok {
		_spec.AddField(menu.FieldOrderNo, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.Disabled(); ok {
		_spec.SetField(menu.FieldDisabled, field.TypeInt32, value)
	}
	if value, ok := muo.mutation.AddedDisabled(); ok {
		_spec.AddField(menu.FieldDisabled, field.TypeInt32, value)
	}
	if muo.mutation.DisabledCleared() {
		_spec.ClearField(menu.FieldDisabled, field.TypeInt32)
	}
	if value, ok := muo.mutation.Ignore(); ok {
		_spec.SetField(menu.FieldIgnore, field.TypeBool, value)
	}
	if muo.mutation.IgnoreCleared() {
		_spec.ClearField(menu.FieldIgnore, field.TypeBool)
	}
	if muo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !muo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.RolesTable,
			Columns: menu.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !muo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ChildrenTable,
			Columns: []string{menu.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ParamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedParamsIDs(); len(nodes) > 0 && !muo.mutation.ParamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ParamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   menu.ParamsTable,
			Columns: []string{menu.ParamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menuparam.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(muo.modifiers...)
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
