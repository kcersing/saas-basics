// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/menuparam"
	"saas/pkg/db/ent/role"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuCreate is the builder for creating a Menu entity.
type MenuCreate struct {
	config
	mutation *MenuMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MenuCreate) SetCreatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MenuCreate) SetUpdatedAt(t time.Time) *MenuCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MenuCreate) SetNillableUpdatedAt(t *time.Time) *MenuCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDelete sets the "delete" field.
func (mc *MenuCreate) SetDelete(i int64) *MenuCreate {
	mc.mutation.SetDelete(i)
	return mc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDelete(i *int64) *MenuCreate {
	if i != nil {
		mc.SetDelete(*i)
	}
	return mc
}

// SetCreatedID sets the "created_id" field.
func (mc *MenuCreate) SetCreatedID(i int64) *MenuCreate {
	mc.mutation.SetCreatedID(i)
	return mc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableCreatedID(i *int64) *MenuCreate {
	if i != nil {
		mc.SetCreatedID(*i)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MenuCreate) SetStatus(i int64) *MenuCreate {
	mc.mutation.SetStatus(i)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MenuCreate) SetNillableStatus(i *int64) *MenuCreate {
	if i != nil {
		mc.SetStatus(*i)
	}
	return mc
}

// SetParentID sets the "parent_id" field.
func (mc *MenuCreate) SetParentID(i int64) *MenuCreate {
	mc.mutation.SetParentID(i)
	return mc
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mc *MenuCreate) SetNillableParentID(i *int64) *MenuCreate {
	if i != nil {
		mc.SetParentID(*i)
	}
	return mc
}

// SetPath sets the "path" field.
func (mc *MenuCreate) SetPath(s string) *MenuCreate {
	mc.mutation.SetPath(s)
	return mc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (mc *MenuCreate) SetNillablePath(s *string) *MenuCreate {
	if s != nil {
		mc.SetPath(*s)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MenuCreate) SetName(s string) *MenuCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mc *MenuCreate) SetNillableName(s *string) *MenuCreate {
	if s != nil {
		mc.SetName(*s)
	}
	return mc
}

// SetSort sets the "sort" field.
func (mc *MenuCreate) SetSort(i int64) *MenuCreate {
	mc.mutation.SetSort(i)
	return mc
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (mc *MenuCreate) SetNillableSort(i *int64) *MenuCreate {
	if i != nil {
		mc.SetSort(*i)
	}
	return mc
}

// SetDisabled sets the "disabled" field.
func (mc *MenuCreate) SetDisabled(i int64) *MenuCreate {
	mc.mutation.SetDisabled(i)
	return mc
}

// SetNillableDisabled sets the "disabled" field if the given value is not nil.
func (mc *MenuCreate) SetNillableDisabled(i *int64) *MenuCreate {
	if i != nil {
		mc.SetDisabled(*i)
	}
	return mc
}

// SetIgnore sets the "ignore" field.
func (mc *MenuCreate) SetIgnore(b bool) *MenuCreate {
	mc.mutation.SetIgnore(b)
	return mc
}

// SetNillableIgnore sets the "ignore" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIgnore(b *bool) *MenuCreate {
	if b != nil {
		mc.SetIgnore(*b)
	}
	return mc
}

// SetLevel sets the "level" field.
func (mc *MenuCreate) SetLevel(i int64) *MenuCreate {
	mc.mutation.SetLevel(i)
	return mc
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (mc *MenuCreate) SetNillableLevel(i *int64) *MenuCreate {
	if i != nil {
		mc.SetLevel(*i)
	}
	return mc
}

// SetMenuType sets the "menu_type" field.
func (mc *MenuCreate) SetMenuType(i int64) *MenuCreate {
	mc.mutation.SetMenuType(i)
	return mc
}

// SetNillableMenuType sets the "menu_type" field if the given value is not nil.
func (mc *MenuCreate) SetNillableMenuType(i *int64) *MenuCreate {
	if i != nil {
		mc.SetMenuType(*i)
	}
	return mc
}

// SetRedirect sets the "redirect" field.
func (mc *MenuCreate) SetRedirect(s string) *MenuCreate {
	mc.mutation.SetRedirect(s)
	return mc
}

// SetNillableRedirect sets the "redirect" field if the given value is not nil.
func (mc *MenuCreate) SetNillableRedirect(s *string) *MenuCreate {
	if s != nil {
		mc.SetRedirect(*s)
	}
	return mc
}

// SetComponent sets the "component" field.
func (mc *MenuCreate) SetComponent(s string) *MenuCreate {
	mc.mutation.SetComponent(s)
	return mc
}

// SetNillableComponent sets the "component" field if the given value is not nil.
func (mc *MenuCreate) SetNillableComponent(s *string) *MenuCreate {
	if s != nil {
		mc.SetComponent(*s)
	}
	return mc
}

// SetURL sets the "url" field.
func (mc *MenuCreate) SetURL(s string) *MenuCreate {
	mc.mutation.SetURL(s)
	return mc
}

// SetNillableURL sets the "url" field if the given value is not nil.
func (mc *MenuCreate) SetNillableURL(s *string) *MenuCreate {
	if s != nil {
		mc.SetURL(*s)
	}
	return mc
}

// SetHidden sets the "hidden" field.
func (mc *MenuCreate) SetHidden(b bool) *MenuCreate {
	mc.mutation.SetHidden(b)
	return mc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (mc *MenuCreate) SetNillableHidden(b *bool) *MenuCreate {
	if b != nil {
		mc.SetHidden(*b)
	}
	return mc
}

// SetTitle sets the "title" field.
func (mc *MenuCreate) SetTitle(s string) *MenuCreate {
	mc.mutation.SetTitle(s)
	return mc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (mc *MenuCreate) SetNillableTitle(s *string) *MenuCreate {
	if s != nil {
		mc.SetTitle(*s)
	}
	return mc
}

// SetIcon sets the "icon" field.
func (mc *MenuCreate) SetIcon(s string) *MenuCreate {
	mc.mutation.SetIcon(s)
	return mc
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mc *MenuCreate) SetNillableIcon(s *string) *MenuCreate {
	if s != nil {
		mc.SetIcon(*s)
	}
	return mc
}

// SetActiveMenu sets the "active_menu" field.
func (mc *MenuCreate) SetActiveMenu(s string) *MenuCreate {
	mc.mutation.SetActiveMenu(s)
	return mc
}

// SetNillableActiveMenu sets the "active_menu" field if the given value is not nil.
func (mc *MenuCreate) SetNillableActiveMenu(s *string) *MenuCreate {
	if s != nil {
		mc.SetActiveMenu(*s)
	}
	return mc
}

// SetAffix sets the "affix" field.
func (mc *MenuCreate) SetAffix(b bool) *MenuCreate {
	mc.mutation.SetAffix(b)
	return mc
}

// SetNillableAffix sets the "affix" field if the given value is not nil.
func (mc *MenuCreate) SetNillableAffix(b *bool) *MenuCreate {
	if b != nil {
		mc.SetAffix(*b)
	}
	return mc
}

// SetNoCache sets the "no_cache" field.
func (mc *MenuCreate) SetNoCache(b bool) *MenuCreate {
	mc.mutation.SetNoCache(b)
	return mc
}

// SetNillableNoCache sets the "no_cache" field if the given value is not nil.
func (mc *MenuCreate) SetNillableNoCache(b *bool) *MenuCreate {
	if b != nil {
		mc.SetNoCache(*b)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MenuCreate) SetID(i int64) *MenuCreate {
	mc.mutation.SetID(i)
	return mc
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (mc *MenuCreate) AddRoleIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddRoleIDs(ids...)
	return mc
}

// AddRoles adds the "roles" edges to the Role entity.
func (mc *MenuCreate) AddRoles(r ...*Role) *MenuCreate {
	ids := make([]int64, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return mc.AddRoleIDs(ids...)
}

// SetParent sets the "parent" edge to the Menu entity.
func (mc *MenuCreate) SetParent(m *Menu) *MenuCreate {
	return mc.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mc *MenuCreate) AddChildIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddChildIDs(ids...)
	return mc
}

// AddChildren adds the "children" edges to the Menu entity.
func (mc *MenuCreate) AddChildren(m ...*Menu) *MenuCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddChildIDs(ids...)
}

// AddParamIDs adds the "params" edge to the MenuParam entity by IDs.
func (mc *MenuCreate) AddParamIDs(ids ...int64) *MenuCreate {
	mc.mutation.AddParamIDs(ids...)
	return mc
}

// AddParams adds the "params" edges to the MenuParam entity.
func (mc *MenuCreate) AddParams(m ...*MenuParam) *MenuCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddParamIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mc *MenuCreate) Mutation() *MenuMutation {
	return mc.mutation
}

// Save creates the Menu in the database.
func (mc *MenuCreate) Save(ctx context.Context) (*Menu, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MenuCreate) SaveX(ctx context.Context) *Menu {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MenuCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MenuCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MenuCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := menu.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := menu.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Delete(); !ok {
		v := menu.DefaultDelete
		mc.mutation.SetDelete(v)
	}
	if _, ok := mc.mutation.CreatedID(); !ok {
		v := menu.DefaultCreatedID
		mc.mutation.SetCreatedID(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := menu.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.Path(); !ok {
		v := menu.DefaultPath
		mc.mutation.SetPath(v)
	}
	if _, ok := mc.mutation.Sort(); !ok {
		v := menu.DefaultSort
		mc.mutation.SetSort(v)
	}
	if _, ok := mc.mutation.Disabled(); !ok {
		v := menu.DefaultDisabled
		mc.mutation.SetDisabled(v)
	}
	if _, ok := mc.mutation.Ignore(); !ok {
		v := menu.DefaultIgnore
		mc.mutation.SetIgnore(v)
	}
	if _, ok := mc.mutation.Redirect(); !ok {
		v := menu.DefaultRedirect
		mc.mutation.SetRedirect(v)
	}
	if _, ok := mc.mutation.Component(); !ok {
		v := menu.DefaultComponent
		mc.mutation.SetComponent(v)
	}
	if _, ok := mc.mutation.URL(); !ok {
		v := menu.DefaultURL
		mc.mutation.SetURL(v)
	}
	if _, ok := mc.mutation.Hidden(); !ok {
		v := menu.DefaultHidden
		mc.mutation.SetHidden(v)
	}
	if _, ok := mc.mutation.ActiveMenu(); !ok {
		v := menu.DefaultActiveMenu
		mc.mutation.SetActiveMenu(v)
	}
	if _, ok := mc.mutation.Affix(); !ok {
		v := menu.DefaultAffix
		mc.mutation.SetAffix(v)
	}
	if _, ok := mc.mutation.NoCache(); !ok {
		v := menu.DefaultNoCache
		mc.mutation.SetNoCache(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MenuCreate) check() error {
	return nil
}

func (mc *MenuCreate) sqlSave(ctx context.Context) (*Menu, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MenuCreate) createSpec() (*Menu, *sqlgraph.CreateSpec) {
	var (
		_node = &Menu{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(menu.Table, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt64))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(menu.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(menu.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Delete(); ok {
		_spec.SetField(menu.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := mc.mutation.CreatedID(); ok {
		_spec.SetField(menu.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.Path(); ok {
		_spec.SetField(menu.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(menu.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Sort(); ok {
		_spec.SetField(menu.FieldSort, field.TypeInt64, value)
		_node.Sort = value
	}
	if value, ok := mc.mutation.Disabled(); ok {
		_spec.SetField(menu.FieldDisabled, field.TypeInt64, value)
		_node.Disabled = value
	}
	if value, ok := mc.mutation.Ignore(); ok {
		_spec.SetField(menu.FieldIgnore, field.TypeBool, value)
		_node.Ignore = value
	}
	if value, ok := mc.mutation.Level(); ok {
		_spec.SetField(menu.FieldLevel, field.TypeInt64, value)
		_node.Level = value
	}
	if value, ok := mc.mutation.MenuType(); ok {
		_spec.SetField(menu.FieldMenuType, field.TypeInt64, value)
		_node.MenuType = value
	}
	if value, ok := mc.mutation.Redirect(); ok {
		_spec.SetField(menu.FieldRedirect, field.TypeString, value)
		_node.Redirect = value
	}
	if value, ok := mc.mutation.Component(); ok {
		_spec.SetField(menu.FieldComponent, field.TypeString, value)
		_node.Component = value
	}
	if value, ok := mc.mutation.URL(); ok {
		_spec.SetField(menu.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := mc.mutation.Hidden(); ok {
		_spec.SetField(menu.FieldHidden, field.TypeBool, value)
		_node.Hidden = value
	}
	if value, ok := mc.mutation.Title(); ok {
		_spec.SetField(menu.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := mc.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
		_node.Icon = value
	}
	if value, ok := mc.mutation.ActiveMenu(); ok {
		_spec.SetField(menu.FieldActiveMenu, field.TypeString, value)
		_node.ActiveMenu = value
	}
	if value, ok := mc.mutation.Affix(); ok {
		_spec.SetField(menu.FieldAffix, field.TypeBool, value)
		_node.Affix = value
	}
	if value, ok := mc.mutation.NoCache(); ok {
		_spec.SetField(menu.FieldNoCache, field.TypeBool, value)
		_node.NoCache = value
	}
	if nodes := mc.mutation.RolesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParentIDs(); len(nodes) > 0 {
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
		_node.ParentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParamsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MenuCreateBulk is the builder for creating many Menu entities in bulk.
type MenuCreateBulk struct {
	config
	err      error
	builders []*MenuCreate
}

// Save creates the Menu entities in the database.
func (mcb *MenuCreateBulk) Save(ctx context.Context) ([]*Menu, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Menu, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MenuMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MenuCreateBulk) SaveX(ctx context.Context) []*Menu {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MenuCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MenuCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
