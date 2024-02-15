// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/menu"
	"saas/pkg/db/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks    []Hook
	mutation *MenuMutation
}

// Where appends a list predicates to the MenuUpdate builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetParentID sets the "parent_id" field.
func (mu *MenuUpdate) SetParentID(i int) *MenuUpdate {
	mu.mutation.SetParentID(i)
	return mu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableParentID(i *int) *MenuUpdate {
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

// SetRouteName sets the "route_name" field.
func (mu *MenuUpdate) SetRouteName(s string) *MenuUpdate {
	mu.mutation.SetRouteName(s)
	return mu
}

// SetNillableRouteName sets the "route_name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableRouteName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetRouteName(*s)
	}
	return mu
}

// ClearRouteName clears the value of the "route_name" field.
func (mu *MenuUpdate) ClearRouteName() *MenuUpdate {
	mu.mutation.ClearRouteName()
	return mu
}

// SetRoutePath sets the "route_path" field.
func (mu *MenuUpdate) SetRoutePath(s string) *MenuUpdate {
	mu.mutation.SetRoutePath(s)
	return mu
}

// SetNillableRoutePath sets the "route_path" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableRoutePath(s *string) *MenuUpdate {
	if s != nil {
		mu.SetRoutePath(*s)
	}
	return mu
}

// ClearRoutePath clears the value of the "route_path" field.
func (mu *MenuUpdate) ClearRoutePath() *MenuUpdate {
	mu.mutation.ClearRoutePath()
	return mu
}

// SetStatus sets the "status" field.
func (mu *MenuUpdate) SetStatus(s string) *MenuUpdate {
	mu.mutation.SetStatus(s)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableStatus(s *string) *MenuUpdate {
	if s != nil {
		mu.SetStatus(*s)
	}
	return mu
}

// ClearStatus clears the value of the "status" field.
func (mu *MenuUpdate) ClearStatus() *MenuUpdate {
	mu.mutation.ClearStatus()
	return mu
}

// SetMenuName sets the "menu_name" field.
func (mu *MenuUpdate) SetMenuName(s string) *MenuUpdate {
	mu.mutation.SetMenuName(s)
	return mu
}

// SetNillableMenuName sets the "menu_name" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableMenuName(s *string) *MenuUpdate {
	if s != nil {
		mu.SetMenuName(*s)
	}
	return mu
}

// ClearMenuName clears the value of the "menu_name" field.
func (mu *MenuUpdate) ClearMenuName() *MenuUpdate {
	mu.mutation.ClearMenuName()
	return mu
}

// SetMenuType sets the "menu_type" field.
func (mu *MenuUpdate) SetMenuType(s string) *MenuUpdate {
	mu.mutation.SetMenuType(s)
	return mu
}

// SetNillableMenuType sets the "menu_type" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableMenuType(s *string) *MenuUpdate {
	if s != nil {
		mu.SetMenuType(*s)
	}
	return mu
}

// ClearMenuType clears the value of the "menu_type" field.
func (mu *MenuUpdate) ClearMenuType() *MenuUpdate {
	mu.mutation.ClearMenuType()
	return mu
}

// SetIconType sets the "icon_type" field.
func (mu *MenuUpdate) SetIconType(s string) *MenuUpdate {
	mu.mutation.SetIconType(s)
	return mu
}

// SetNillableIconType sets the "icon_type" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableIconType(s *string) *MenuUpdate {
	if s != nil {
		mu.SetIconType(*s)
	}
	return mu
}

// ClearIconType clears the value of the "icon_type" field.
func (mu *MenuUpdate) ClearIconType() *MenuUpdate {
	mu.mutation.ClearIconType()
	return mu
}

// SetIcon sets the "icon" field.
func (mu *MenuUpdate) SetIcon(s string) *MenuUpdate {
	mu.mutation.SetIcon(s)
	return mu
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableIcon(s *string) *MenuUpdate {
	if s != nil {
		mu.SetIcon(*s)
	}
	return mu
}

// ClearIcon clears the value of the "icon" field.
func (mu *MenuUpdate) ClearIcon() *MenuUpdate {
	mu.mutation.ClearIcon()
	return mu
}

// SetI18nKey sets the "i18n_key" field.
func (mu *MenuUpdate) SetI18nKey(s string) *MenuUpdate {
	mu.mutation.SetI18nKey(s)
	return mu
}

// SetNillableI18nKey sets the "i18n_key" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableI18nKey(s *string) *MenuUpdate {
	if s != nil {
		mu.SetI18nKey(*s)
	}
	return mu
}

// ClearI18nKey clears the value of the "i18n_key" field.
func (mu *MenuUpdate) ClearI18nKey() *MenuUpdate {
	mu.mutation.ClearI18nKey()
	return mu
}

// SetLevel sets the "level" field.
func (mu *MenuUpdate) SetLevel(s string) *MenuUpdate {
	mu.mutation.SetLevel(s)
	return mu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (mu *MenuUpdate) SetNillableLevel(s *string) *MenuUpdate {
	if s != nil {
		mu.SetLevel(*s)
	}
	return mu
}

// ClearLevel clears the value of the "level" field.
func (mu *MenuUpdate) ClearLevel() *MenuUpdate {
	mu.mutation.ClearLevel()
	return mu
}

// SetParent sets the "parent" edge to the Menu entity.
func (mu *MenuUpdate) SetParent(m *Menu) *MenuUpdate {
	return mu.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (mu *MenuUpdate) AddChildIDs(ids ...int) *MenuUpdate {
	mu.mutation.AddChildIDs(ids...)
	return mu
}

// AddChildren adds the "children" edges to the Menu entity.
func (mu *MenuUpdate) AddChildren(m ...*Menu) *MenuUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddChildIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
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
func (mu *MenuUpdate) RemoveChildIDs(ids ...int) *MenuUpdate {
	mu.mutation.RemoveChildIDs(ids...)
	return mu
}

// RemoveChildren removes "children" edges to Menu entities.
func (mu *MenuUpdate) RemoveChildren(m ...*Menu) *MenuUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
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

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.RouteName(); ok {
		_spec.SetField(menu.FieldRouteName, field.TypeString, value)
	}
	if mu.mutation.RouteNameCleared() {
		_spec.ClearField(menu.FieldRouteName, field.TypeString)
	}
	if value, ok := mu.mutation.RoutePath(); ok {
		_spec.SetField(menu.FieldRoutePath, field.TypeString, value)
	}
	if mu.mutation.RoutePathCleared() {
		_spec.ClearField(menu.FieldRoutePath, field.TypeString)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeString, value)
	}
	if mu.mutation.StatusCleared() {
		_spec.ClearField(menu.FieldStatus, field.TypeString)
	}
	if value, ok := mu.mutation.MenuName(); ok {
		_spec.SetField(menu.FieldMenuName, field.TypeString, value)
	}
	if mu.mutation.MenuNameCleared() {
		_spec.ClearField(menu.FieldMenuName, field.TypeString)
	}
	if value, ok := mu.mutation.MenuType(); ok {
		_spec.SetField(menu.FieldMenuType, field.TypeString, value)
	}
	if mu.mutation.MenuTypeCleared() {
		_spec.ClearField(menu.FieldMenuType, field.TypeString)
	}
	if value, ok := mu.mutation.IconType(); ok {
		_spec.SetField(menu.FieldIconType, field.TypeString, value)
	}
	if mu.mutation.IconTypeCleared() {
		_spec.ClearField(menu.FieldIconType, field.TypeString)
	}
	if value, ok := mu.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
	}
	if mu.mutation.IconCleared() {
		_spec.ClearField(menu.FieldIcon, field.TypeString)
	}
	if value, ok := mu.mutation.I18nKey(); ok {
		_spec.SetField(menu.FieldI18nKey, field.TypeString, value)
	}
	if mu.mutation.I18nKeyCleared() {
		_spec.ClearField(menu.FieldI18nKey, field.TypeString)
	}
	if value, ok := mu.mutation.Level(); ok {
		_spec.SetField(menu.FieldLevel, field.TypeString, value)
	}
	if mu.mutation.LevelCleared() {
		_spec.ClearField(menu.FieldLevel, field.TypeString)
	}
	if mu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
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
	fields   []string
	hooks    []Hook
	mutation *MenuMutation
}

// SetParentID sets the "parent_id" field.
func (muo *MenuUpdateOne) SetParentID(i int) *MenuUpdateOne {
	muo.mutation.SetParentID(i)
	return muo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableParentID(i *int) *MenuUpdateOne {
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

// SetRouteName sets the "route_name" field.
func (muo *MenuUpdateOne) SetRouteName(s string) *MenuUpdateOne {
	muo.mutation.SetRouteName(s)
	return muo
}

// SetNillableRouteName sets the "route_name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableRouteName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetRouteName(*s)
	}
	return muo
}

// ClearRouteName clears the value of the "route_name" field.
func (muo *MenuUpdateOne) ClearRouteName() *MenuUpdateOne {
	muo.mutation.ClearRouteName()
	return muo
}

// SetRoutePath sets the "route_path" field.
func (muo *MenuUpdateOne) SetRoutePath(s string) *MenuUpdateOne {
	muo.mutation.SetRoutePath(s)
	return muo
}

// SetNillableRoutePath sets the "route_path" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableRoutePath(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetRoutePath(*s)
	}
	return muo
}

// ClearRoutePath clears the value of the "route_path" field.
func (muo *MenuUpdateOne) ClearRoutePath() *MenuUpdateOne {
	muo.mutation.ClearRoutePath()
	return muo
}

// SetStatus sets the "status" field.
func (muo *MenuUpdateOne) SetStatus(s string) *MenuUpdateOne {
	muo.mutation.SetStatus(s)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableStatus(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetStatus(*s)
	}
	return muo
}

// ClearStatus clears the value of the "status" field.
func (muo *MenuUpdateOne) ClearStatus() *MenuUpdateOne {
	muo.mutation.ClearStatus()
	return muo
}

// SetMenuName sets the "menu_name" field.
func (muo *MenuUpdateOne) SetMenuName(s string) *MenuUpdateOne {
	muo.mutation.SetMenuName(s)
	return muo
}

// SetNillableMenuName sets the "menu_name" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableMenuName(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetMenuName(*s)
	}
	return muo
}

// ClearMenuName clears the value of the "menu_name" field.
func (muo *MenuUpdateOne) ClearMenuName() *MenuUpdateOne {
	muo.mutation.ClearMenuName()
	return muo
}

// SetMenuType sets the "menu_type" field.
func (muo *MenuUpdateOne) SetMenuType(s string) *MenuUpdateOne {
	muo.mutation.SetMenuType(s)
	return muo
}

// SetNillableMenuType sets the "menu_type" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableMenuType(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetMenuType(*s)
	}
	return muo
}

// ClearMenuType clears the value of the "menu_type" field.
func (muo *MenuUpdateOne) ClearMenuType() *MenuUpdateOne {
	muo.mutation.ClearMenuType()
	return muo
}

// SetIconType sets the "icon_type" field.
func (muo *MenuUpdateOne) SetIconType(s string) *MenuUpdateOne {
	muo.mutation.SetIconType(s)
	return muo
}

// SetNillableIconType sets the "icon_type" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableIconType(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetIconType(*s)
	}
	return muo
}

// ClearIconType clears the value of the "icon_type" field.
func (muo *MenuUpdateOne) ClearIconType() *MenuUpdateOne {
	muo.mutation.ClearIconType()
	return muo
}

// SetIcon sets the "icon" field.
func (muo *MenuUpdateOne) SetIcon(s string) *MenuUpdateOne {
	muo.mutation.SetIcon(s)
	return muo
}

// SetNillableIcon sets the "icon" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableIcon(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetIcon(*s)
	}
	return muo
}

// ClearIcon clears the value of the "icon" field.
func (muo *MenuUpdateOne) ClearIcon() *MenuUpdateOne {
	muo.mutation.ClearIcon()
	return muo
}

// SetI18nKey sets the "i18n_key" field.
func (muo *MenuUpdateOne) SetI18nKey(s string) *MenuUpdateOne {
	muo.mutation.SetI18nKey(s)
	return muo
}

// SetNillableI18nKey sets the "i18n_key" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableI18nKey(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetI18nKey(*s)
	}
	return muo
}

// ClearI18nKey clears the value of the "i18n_key" field.
func (muo *MenuUpdateOne) ClearI18nKey() *MenuUpdateOne {
	muo.mutation.ClearI18nKey()
	return muo
}

// SetLevel sets the "level" field.
func (muo *MenuUpdateOne) SetLevel(s string) *MenuUpdateOne {
	muo.mutation.SetLevel(s)
	return muo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableLevel(s *string) *MenuUpdateOne {
	if s != nil {
		muo.SetLevel(*s)
	}
	return muo
}

// ClearLevel clears the value of the "level" field.
func (muo *MenuUpdateOne) ClearLevel() *MenuUpdateOne {
	muo.mutation.ClearLevel()
	return muo
}

// SetParent sets the "parent" edge to the Menu entity.
func (muo *MenuUpdateOne) SetParent(m *Menu) *MenuUpdateOne {
	return muo.SetParentID(m.ID)
}

// AddChildIDs adds the "children" edge to the Menu entity by IDs.
func (muo *MenuUpdateOne) AddChildIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.AddChildIDs(ids...)
	return muo
}

// AddChildren adds the "children" edges to the Menu entity.
func (muo *MenuUpdateOne) AddChildren(m ...*Menu) *MenuUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddChildIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
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
func (muo *MenuUpdateOne) RemoveChildIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.RemoveChildIDs(ids...)
	return muo
}

// RemoveChildren removes "children" edges to Menu entities.
func (muo *MenuUpdateOne) RemoveChildren(m ...*Menu) *MenuUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveChildIDs(ids...)
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

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := sqlgraph.NewUpdateSpec(menu.Table, menu.Columns, sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt))
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
	if value, ok := muo.mutation.RouteName(); ok {
		_spec.SetField(menu.FieldRouteName, field.TypeString, value)
	}
	if muo.mutation.RouteNameCleared() {
		_spec.ClearField(menu.FieldRouteName, field.TypeString)
	}
	if value, ok := muo.mutation.RoutePath(); ok {
		_spec.SetField(menu.FieldRoutePath, field.TypeString, value)
	}
	if muo.mutation.RoutePathCleared() {
		_spec.ClearField(menu.FieldRoutePath, field.TypeString)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(menu.FieldStatus, field.TypeString, value)
	}
	if muo.mutation.StatusCleared() {
		_spec.ClearField(menu.FieldStatus, field.TypeString)
	}
	if value, ok := muo.mutation.MenuName(); ok {
		_spec.SetField(menu.FieldMenuName, field.TypeString, value)
	}
	if muo.mutation.MenuNameCleared() {
		_spec.ClearField(menu.FieldMenuName, field.TypeString)
	}
	if value, ok := muo.mutation.MenuType(); ok {
		_spec.SetField(menu.FieldMenuType, field.TypeString, value)
	}
	if muo.mutation.MenuTypeCleared() {
		_spec.ClearField(menu.FieldMenuType, field.TypeString)
	}
	if value, ok := muo.mutation.IconType(); ok {
		_spec.SetField(menu.FieldIconType, field.TypeString, value)
	}
	if muo.mutation.IconTypeCleared() {
		_spec.ClearField(menu.FieldIconType, field.TypeString)
	}
	if value, ok := muo.mutation.Icon(); ok {
		_spec.SetField(menu.FieldIcon, field.TypeString, value)
	}
	if muo.mutation.IconCleared() {
		_spec.ClearField(menu.FieldIcon, field.TypeString)
	}
	if value, ok := muo.mutation.I18nKey(); ok {
		_spec.SetField(menu.FieldI18nKey, field.TypeString, value)
	}
	if muo.mutation.I18nKeyCleared() {
		_spec.ClearField(menu.FieldI18nKey, field.TypeString)
	}
	if value, ok := muo.mutation.Level(); ok {
		_spec.SetField(menu.FieldLevel, field.TypeString, value)
	}
	if muo.mutation.LevelCleared() {
		_spec.ClearField(menu.FieldLevel, field.TypeString)
	}
	if muo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   menu.ParentTable,
			Columns: []string{menu.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
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
				IDSpec: sqlgraph.NewFieldSpec(menu.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
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