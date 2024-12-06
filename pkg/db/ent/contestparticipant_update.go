// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/contest"
	"saas/pkg/db/ent/contestparticipant"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ContestParticipantUpdate is the builder for updating ContestParticipant entities.
type ContestParticipantUpdate struct {
	config
	hooks    []Hook
	mutation *ContestParticipantMutation
}

// Where appends a list predicates to the ContestParticipantUpdate builder.
func (cpu *ContestParticipantUpdate) Where(ps ...predicate.ContestParticipant) *ContestParticipantUpdate {
	cpu.mutation.Where(ps...)
	return cpu
}

// SetUpdatedAt sets the "updated_at" field.
func (cpu *ContestParticipantUpdate) SetUpdatedAt(t time.Time) *ContestParticipantUpdate {
	cpu.mutation.SetUpdatedAt(t)
	return cpu
}

// SetDeleteAt sets the "delete_at" field.
func (cpu *ContestParticipantUpdate) SetDeleteAt(t time.Time) *ContestParticipantUpdate {
	cpu.mutation.SetDeleteAt(t)
	return cpu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableDeleteAt(t *time.Time) *ContestParticipantUpdate {
	if t != nil {
		cpu.SetDeleteAt(*t)
	}
	return cpu
}

// SetCreatedID sets the "created_id" field.
func (cpu *ContestParticipantUpdate) SetCreatedID(i int64) *ContestParticipantUpdate {
	cpu.mutation.ResetCreatedID()
	cpu.mutation.SetCreatedID(i)
	return cpu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableCreatedID(i *int64) *ContestParticipantUpdate {
	if i != nil {
		cpu.SetCreatedID(*i)
	}
	return cpu
}

// AddCreatedID adds i to the "created_id" field.
func (cpu *ContestParticipantUpdate) AddCreatedID(i int64) *ContestParticipantUpdate {
	cpu.mutation.AddCreatedID(i)
	return cpu
}

// SetStatus sets the "status" field.
func (cpu *ContestParticipantUpdate) SetStatus(i int64) *ContestParticipantUpdate {
	cpu.mutation.ResetStatus()
	cpu.mutation.SetStatus(i)
	return cpu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableStatus(i *int64) *ContestParticipantUpdate {
	if i != nil {
		cpu.SetStatus(*i)
	}
	return cpu
}

// AddStatus adds i to the "status" field.
func (cpu *ContestParticipantUpdate) AddStatus(i int64) *ContestParticipantUpdate {
	cpu.mutation.AddStatus(i)
	return cpu
}

// ClearStatus clears the value of the "status" field.
func (cpu *ContestParticipantUpdate) ClearStatus() *ContestParticipantUpdate {
	cpu.mutation.ClearStatus()
	return cpu
}

// SetContestID sets the "contest_id" field.
func (cpu *ContestParticipantUpdate) SetContestID(i int64) *ContestParticipantUpdate {
	cpu.mutation.SetContestID(i)
	return cpu
}

// SetNillableContestID sets the "contest_id" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableContestID(i *int64) *ContestParticipantUpdate {
	if i != nil {
		cpu.SetContestID(*i)
	}
	return cpu
}

// ClearContestID clears the value of the "contest_id" field.
func (cpu *ContestParticipantUpdate) ClearContestID() *ContestParticipantUpdate {
	cpu.mutation.ClearContestID()
	return cpu
}

// SetName sets the "name" field.
func (cpu *ContestParticipantUpdate) SetName(s string) *ContestParticipantUpdate {
	cpu.mutation.SetName(s)
	return cpu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableName(s *string) *ContestParticipantUpdate {
	if s != nil {
		cpu.SetName(*s)
	}
	return cpu
}

// ClearName clears the value of the "name" field.
func (cpu *ContestParticipantUpdate) ClearName() *ContestParticipantUpdate {
	cpu.mutation.ClearName()
	return cpu
}

// SetMobile sets the "mobile" field.
func (cpu *ContestParticipantUpdate) SetMobile(s string) *ContestParticipantUpdate {
	cpu.mutation.SetMobile(s)
	return cpu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableMobile(s *string) *ContestParticipantUpdate {
	if s != nil {
		cpu.SetMobile(*s)
	}
	return cpu
}

// ClearMobile clears the value of the "mobile" field.
func (cpu *ContestParticipantUpdate) ClearMobile() *ContestParticipantUpdate {
	cpu.mutation.ClearMobile()
	return cpu
}

// SetFields sets the "fields" field.
func (cpu *ContestParticipantUpdate) SetFields(s string) *ContestParticipantUpdate {
	cpu.mutation.SetFields(s)
	return cpu
}

// SetNillableFields sets the "fields" field if the given value is not nil.
func (cpu *ContestParticipantUpdate) SetNillableFields(s *string) *ContestParticipantUpdate {
	if s != nil {
		cpu.SetFields(*s)
	}
	return cpu
}

// ClearFields clears the value of the "fields" field.
func (cpu *ContestParticipantUpdate) ClearFields() *ContestParticipantUpdate {
	cpu.mutation.ClearFields()
	return cpu
}

// SetContest sets the "contest" edge to the Contest entity.
func (cpu *ContestParticipantUpdate) SetContest(c *Contest) *ContestParticipantUpdate {
	return cpu.SetContestID(c.ID)
}

// Mutation returns the ContestParticipantMutation object of the builder.
func (cpu *ContestParticipantUpdate) Mutation() *ContestParticipantMutation {
	return cpu.mutation
}

// ClearContest clears the "contest" edge to the Contest entity.
func (cpu *ContestParticipantUpdate) ClearContest() *ContestParticipantUpdate {
	cpu.mutation.ClearContest()
	return cpu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cpu *ContestParticipantUpdate) Save(ctx context.Context) (int, error) {
	cpu.defaults()
	return withHooks(ctx, cpu.sqlSave, cpu.mutation, cpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpu *ContestParticipantUpdate) SaveX(ctx context.Context) int {
	affected, err := cpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cpu *ContestParticipantUpdate) Exec(ctx context.Context) error {
	_, err := cpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpu *ContestParticipantUpdate) ExecX(ctx context.Context) {
	if err := cpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpu *ContestParticipantUpdate) defaults() {
	if _, ok := cpu.mutation.UpdatedAt(); !ok {
		v := contestparticipant.UpdateDefaultUpdatedAt()
		cpu.mutation.SetUpdatedAt(v)
	}
}

func (cpu *ContestParticipantUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(contestparticipant.Table, contestparticipant.Columns, sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64))
	if ps := cpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpu.mutation.UpdatedAt(); ok {
		_spec.SetField(contestparticipant.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.DeleteAt(); ok {
		_spec.SetField(contestparticipant.FieldDeleteAt, field.TypeTime, value)
	}
	if value, ok := cpu.mutation.CreatedID(); ok {
		_spec.SetField(contestparticipant.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := cpu.mutation.AddedCreatedID(); ok {
		_spec.AddField(contestparticipant.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := cpu.mutation.Status(); ok {
		_spec.SetField(contestparticipant.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := cpu.mutation.AddedStatus(); ok {
		_spec.AddField(contestparticipant.FieldStatus, field.TypeInt64, value)
	}
	if cpu.mutation.StatusCleared() {
		_spec.ClearField(contestparticipant.FieldStatus, field.TypeInt64)
	}
	if value, ok := cpu.mutation.Name(); ok {
		_spec.SetField(contestparticipant.FieldName, field.TypeString, value)
	}
	if cpu.mutation.NameCleared() {
		_spec.ClearField(contestparticipant.FieldName, field.TypeString)
	}
	if value, ok := cpu.mutation.Mobile(); ok {
		_spec.SetField(contestparticipant.FieldMobile, field.TypeString, value)
	}
	if cpu.mutation.MobileCleared() {
		_spec.ClearField(contestparticipant.FieldMobile, field.TypeString)
	}
	if value, ok := cpu.mutation.GetFields(); ok {
		_spec.SetField(contestparticipant.FieldFields, field.TypeString, value)
	}
	if cpu.mutation.FieldsCleared() {
		_spec.ClearField(contestparticipant.FieldFields, field.TypeString)
	}
	if cpu.mutation.ContestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contestparticipant.ContestTable,
			Columns: []string{contestparticipant.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpu.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contestparticipant.ContestTable,
			Columns: []string{contestparticipant.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contestparticipant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cpu.mutation.done = true
	return n, nil
}

// ContestParticipantUpdateOne is the builder for updating a single ContestParticipant entity.
type ContestParticipantUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContestParticipantMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cpuo *ContestParticipantUpdateOne) SetUpdatedAt(t time.Time) *ContestParticipantUpdateOne {
	cpuo.mutation.SetUpdatedAt(t)
	return cpuo
}

// SetDeleteAt sets the "delete_at" field.
func (cpuo *ContestParticipantUpdateOne) SetDeleteAt(t time.Time) *ContestParticipantUpdateOne {
	cpuo.mutation.SetDeleteAt(t)
	return cpuo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableDeleteAt(t *time.Time) *ContestParticipantUpdateOne {
	if t != nil {
		cpuo.SetDeleteAt(*t)
	}
	return cpuo
}

// SetCreatedID sets the "created_id" field.
func (cpuo *ContestParticipantUpdateOne) SetCreatedID(i int64) *ContestParticipantUpdateOne {
	cpuo.mutation.ResetCreatedID()
	cpuo.mutation.SetCreatedID(i)
	return cpuo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableCreatedID(i *int64) *ContestParticipantUpdateOne {
	if i != nil {
		cpuo.SetCreatedID(*i)
	}
	return cpuo
}

// AddCreatedID adds i to the "created_id" field.
func (cpuo *ContestParticipantUpdateOne) AddCreatedID(i int64) *ContestParticipantUpdateOne {
	cpuo.mutation.AddCreatedID(i)
	return cpuo
}

// SetStatus sets the "status" field.
func (cpuo *ContestParticipantUpdateOne) SetStatus(i int64) *ContestParticipantUpdateOne {
	cpuo.mutation.ResetStatus()
	cpuo.mutation.SetStatus(i)
	return cpuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableStatus(i *int64) *ContestParticipantUpdateOne {
	if i != nil {
		cpuo.SetStatus(*i)
	}
	return cpuo
}

// AddStatus adds i to the "status" field.
func (cpuo *ContestParticipantUpdateOne) AddStatus(i int64) *ContestParticipantUpdateOne {
	cpuo.mutation.AddStatus(i)
	return cpuo
}

// ClearStatus clears the value of the "status" field.
func (cpuo *ContestParticipantUpdateOne) ClearStatus() *ContestParticipantUpdateOne {
	cpuo.mutation.ClearStatus()
	return cpuo
}

// SetContestID sets the "contest_id" field.
func (cpuo *ContestParticipantUpdateOne) SetContestID(i int64) *ContestParticipantUpdateOne {
	cpuo.mutation.SetContestID(i)
	return cpuo
}

// SetNillableContestID sets the "contest_id" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableContestID(i *int64) *ContestParticipantUpdateOne {
	if i != nil {
		cpuo.SetContestID(*i)
	}
	return cpuo
}

// ClearContestID clears the value of the "contest_id" field.
func (cpuo *ContestParticipantUpdateOne) ClearContestID() *ContestParticipantUpdateOne {
	cpuo.mutation.ClearContestID()
	return cpuo
}

// SetName sets the "name" field.
func (cpuo *ContestParticipantUpdateOne) SetName(s string) *ContestParticipantUpdateOne {
	cpuo.mutation.SetName(s)
	return cpuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableName(s *string) *ContestParticipantUpdateOne {
	if s != nil {
		cpuo.SetName(*s)
	}
	return cpuo
}

// ClearName clears the value of the "name" field.
func (cpuo *ContestParticipantUpdateOne) ClearName() *ContestParticipantUpdateOne {
	cpuo.mutation.ClearName()
	return cpuo
}

// SetMobile sets the "mobile" field.
func (cpuo *ContestParticipantUpdateOne) SetMobile(s string) *ContestParticipantUpdateOne {
	cpuo.mutation.SetMobile(s)
	return cpuo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableMobile(s *string) *ContestParticipantUpdateOne {
	if s != nil {
		cpuo.SetMobile(*s)
	}
	return cpuo
}

// ClearMobile clears the value of the "mobile" field.
func (cpuo *ContestParticipantUpdateOne) ClearMobile() *ContestParticipantUpdateOne {
	cpuo.mutation.ClearMobile()
	return cpuo
}

// SetFields sets the "fields" field.
func (cpuo *ContestParticipantUpdateOne) SetFields(s string) *ContestParticipantUpdateOne {
	cpuo.mutation.SetFields(s)
	return cpuo
}

// SetNillableFields sets the "fields" field if the given value is not nil.
func (cpuo *ContestParticipantUpdateOne) SetNillableFields(s *string) *ContestParticipantUpdateOne {
	if s != nil {
		cpuo.SetFields(*s)
	}
	return cpuo
}

// ClearFields clears the value of the "fields" field.
func (cpuo *ContestParticipantUpdateOne) ClearFields() *ContestParticipantUpdateOne {
	cpuo.mutation.ClearFields()
	return cpuo
}

// SetContest sets the "contest" edge to the Contest entity.
func (cpuo *ContestParticipantUpdateOne) SetContest(c *Contest) *ContestParticipantUpdateOne {
	return cpuo.SetContestID(c.ID)
}

// Mutation returns the ContestParticipantMutation object of the builder.
func (cpuo *ContestParticipantUpdateOne) Mutation() *ContestParticipantMutation {
	return cpuo.mutation
}

// ClearContest clears the "contest" edge to the Contest entity.
func (cpuo *ContestParticipantUpdateOne) ClearContest() *ContestParticipantUpdateOne {
	cpuo.mutation.ClearContest()
	return cpuo
}

// Where appends a list predicates to the ContestParticipantUpdate builder.
func (cpuo *ContestParticipantUpdateOne) Where(ps ...predicate.ContestParticipant) *ContestParticipantUpdateOne {
	cpuo.mutation.Where(ps...)
	return cpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cpuo *ContestParticipantUpdateOne) Select(field string, fields ...string) *ContestParticipantUpdateOne {
	cpuo.fields = append([]string{field}, fields...)
	return cpuo
}

// Save executes the query and returns the updated ContestParticipant entity.
func (cpuo *ContestParticipantUpdateOne) Save(ctx context.Context) (*ContestParticipant, error) {
	cpuo.defaults()
	return withHooks(ctx, cpuo.sqlSave, cpuo.mutation, cpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cpuo *ContestParticipantUpdateOne) SaveX(ctx context.Context) *ContestParticipant {
	node, err := cpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cpuo *ContestParticipantUpdateOne) Exec(ctx context.Context) error {
	_, err := cpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cpuo *ContestParticipantUpdateOne) ExecX(ctx context.Context) {
	if err := cpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cpuo *ContestParticipantUpdateOne) defaults() {
	if _, ok := cpuo.mutation.UpdatedAt(); !ok {
		v := contestparticipant.UpdateDefaultUpdatedAt()
		cpuo.mutation.SetUpdatedAt(v)
	}
}

func (cpuo *ContestParticipantUpdateOne) sqlSave(ctx context.Context) (_node *ContestParticipant, err error) {
	_spec := sqlgraph.NewUpdateSpec(contestparticipant.Table, contestparticipant.Columns, sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64))
	id, ok := cpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ContestParticipant.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contestparticipant.FieldID)
		for _, f := range fields {
			if !contestparticipant.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != contestparticipant.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cpuo.mutation.UpdatedAt(); ok {
		_spec.SetField(contestparticipant.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.DeleteAt(); ok {
		_spec.SetField(contestparticipant.FieldDeleteAt, field.TypeTime, value)
	}
	if value, ok := cpuo.mutation.CreatedID(); ok {
		_spec.SetField(contestparticipant.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := cpuo.mutation.AddedCreatedID(); ok {
		_spec.AddField(contestparticipant.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := cpuo.mutation.Status(); ok {
		_spec.SetField(contestparticipant.FieldStatus, field.TypeInt64, value)
	}
	if value, ok := cpuo.mutation.AddedStatus(); ok {
		_spec.AddField(contestparticipant.FieldStatus, field.TypeInt64, value)
	}
	if cpuo.mutation.StatusCleared() {
		_spec.ClearField(contestparticipant.FieldStatus, field.TypeInt64)
	}
	if value, ok := cpuo.mutation.Name(); ok {
		_spec.SetField(contestparticipant.FieldName, field.TypeString, value)
	}
	if cpuo.mutation.NameCleared() {
		_spec.ClearField(contestparticipant.FieldName, field.TypeString)
	}
	if value, ok := cpuo.mutation.Mobile(); ok {
		_spec.SetField(contestparticipant.FieldMobile, field.TypeString, value)
	}
	if cpuo.mutation.MobileCleared() {
		_spec.ClearField(contestparticipant.FieldMobile, field.TypeString)
	}
	if value, ok := cpuo.mutation.GetFields(); ok {
		_spec.SetField(contestparticipant.FieldFields, field.TypeString, value)
	}
	if cpuo.mutation.FieldsCleared() {
		_spec.ClearField(contestparticipant.FieldFields, field.TypeString)
	}
	if cpuo.mutation.ContestCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contestparticipant.ContestTable,
			Columns: []string{contestparticipant.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cpuo.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contestparticipant.ContestTable,
			Columns: []string{contestparticipant.ContestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ContestParticipant{config: cpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{contestparticipant.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cpuo.mutation.done = true
	return _node, nil
}
