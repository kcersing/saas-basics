// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/logs"
	"saas/pkg/db/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LogsUpdate is the builder for updating Logs entities.
type LogsUpdate struct {
	config
	hooks    []Hook
	mutation *LogsMutation
}

// Where appends a list predicates to the LogsUpdate builder.
func (lu *LogsUpdate) Where(ps ...predicate.Logs) *LogsUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LogsUpdate) SetUpdatedAt(t time.Time) *LogsUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetDeleteAt sets the "delete_at" field.
func (lu *LogsUpdate) SetDeleteAt(t time.Time) *LogsUpdate {
	lu.mutation.SetDeleteAt(t)
	return lu
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableDeleteAt(t *time.Time) *LogsUpdate {
	if t != nil {
		lu.SetDeleteAt(*t)
	}
	return lu
}

// SetCreatedID sets the "created_id" field.
func (lu *LogsUpdate) SetCreatedID(i int64) *LogsUpdate {
	lu.mutation.ResetCreatedID()
	lu.mutation.SetCreatedID(i)
	return lu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableCreatedID(i *int64) *LogsUpdate {
	if i != nil {
		lu.SetCreatedID(*i)
	}
	return lu
}

// AddCreatedID adds i to the "created_id" field.
func (lu *LogsUpdate) AddCreatedID(i int64) *LogsUpdate {
	lu.mutation.AddCreatedID(i)
	return lu
}

// SetType sets the "type" field.
func (lu *LogsUpdate) SetType(s string) *LogsUpdate {
	lu.mutation.SetType(s)
	return lu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableType(s *string) *LogsUpdate {
	if s != nil {
		lu.SetType(*s)
	}
	return lu
}

// SetMethod sets the "method" field.
func (lu *LogsUpdate) SetMethod(s string) *LogsUpdate {
	lu.mutation.SetMethod(s)
	return lu
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableMethod(s *string) *LogsUpdate {
	if s != nil {
		lu.SetMethod(*s)
	}
	return lu
}

// SetAPI sets the "api" field.
func (lu *LogsUpdate) SetAPI(s string) *LogsUpdate {
	lu.mutation.SetAPI(s)
	return lu
}

// SetNillableAPI sets the "api" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableAPI(s *string) *LogsUpdate {
	if s != nil {
		lu.SetAPI(*s)
	}
	return lu
}

// SetSuccess sets the "success" field.
func (lu *LogsUpdate) SetSuccess(b bool) *LogsUpdate {
	lu.mutation.SetSuccess(b)
	return lu
}

// SetNillableSuccess sets the "success" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableSuccess(b *bool) *LogsUpdate {
	if b != nil {
		lu.SetSuccess(*b)
	}
	return lu
}

// SetReqContent sets the "req_content" field.
func (lu *LogsUpdate) SetReqContent(s string) *LogsUpdate {
	lu.mutation.SetReqContent(s)
	return lu
}

// SetNillableReqContent sets the "req_content" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableReqContent(s *string) *LogsUpdate {
	if s != nil {
		lu.SetReqContent(*s)
	}
	return lu
}

// ClearReqContent clears the value of the "req_content" field.
func (lu *LogsUpdate) ClearReqContent() *LogsUpdate {
	lu.mutation.ClearReqContent()
	return lu
}

// SetRespContent sets the "resp_content" field.
func (lu *LogsUpdate) SetRespContent(s string) *LogsUpdate {
	lu.mutation.SetRespContent(s)
	return lu
}

// SetNillableRespContent sets the "resp_content" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableRespContent(s *string) *LogsUpdate {
	if s != nil {
		lu.SetRespContent(*s)
	}
	return lu
}

// ClearRespContent clears the value of the "resp_content" field.
func (lu *LogsUpdate) ClearRespContent() *LogsUpdate {
	lu.mutation.ClearRespContent()
	return lu
}

// SetIP sets the "ip" field.
func (lu *LogsUpdate) SetIP(s string) *LogsUpdate {
	lu.mutation.SetIP(s)
	return lu
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableIP(s *string) *LogsUpdate {
	if s != nil {
		lu.SetIP(*s)
	}
	return lu
}

// ClearIP clears the value of the "ip" field.
func (lu *LogsUpdate) ClearIP() *LogsUpdate {
	lu.mutation.ClearIP()
	return lu
}

// SetUserAgent sets the "user_agent" field.
func (lu *LogsUpdate) SetUserAgent(s string) *LogsUpdate {
	lu.mutation.SetUserAgent(s)
	return lu
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableUserAgent(s *string) *LogsUpdate {
	if s != nil {
		lu.SetUserAgent(*s)
	}
	return lu
}

// ClearUserAgent clears the value of the "user_agent" field.
func (lu *LogsUpdate) ClearUserAgent() *LogsUpdate {
	lu.mutation.ClearUserAgent()
	return lu
}

// SetOperator sets the "operator" field.
func (lu *LogsUpdate) SetOperator(s string) *LogsUpdate {
	lu.mutation.SetOperator(s)
	return lu
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableOperator(s *string) *LogsUpdate {
	if s != nil {
		lu.SetOperator(*s)
	}
	return lu
}

// ClearOperator clears the value of the "operator" field.
func (lu *LogsUpdate) ClearOperator() *LogsUpdate {
	lu.mutation.ClearOperator()
	return lu
}

// SetTime sets the "time" field.
func (lu *LogsUpdate) SetTime(i int) *LogsUpdate {
	lu.mutation.ResetTime()
	lu.mutation.SetTime(i)
	return lu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (lu *LogsUpdate) SetNillableTime(i *int) *LogsUpdate {
	if i != nil {
		lu.SetTime(*i)
	}
	return lu
}

// AddTime adds i to the "time" field.
func (lu *LogsUpdate) AddTime(i int) *LogsUpdate {
	lu.mutation.AddTime(i)
	return lu
}

// ClearTime clears the value of the "time" field.
func (lu *LogsUpdate) ClearTime() *LogsUpdate {
	lu.mutation.ClearTime()
	return lu
}

// Mutation returns the LogsMutation object of the builder.
func (lu *LogsUpdate) Mutation() *LogsMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LogsUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LogsUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LogsUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LogsUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LogsUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := logs.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
}

func (lu *LogsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(logs.Table, logs.Columns, sqlgraph.NewFieldSpec(logs.FieldID, field.TypeInt64))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(logs.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.DeleteAt(); ok {
		_spec.SetField(logs.FieldDeleteAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.CreatedID(); ok {
		_spec.SetField(logs.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := lu.mutation.AddedCreatedID(); ok {
		_spec.AddField(logs.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := lu.mutation.GetType(); ok {
		_spec.SetField(logs.FieldType, field.TypeString, value)
	}
	if value, ok := lu.mutation.Method(); ok {
		_spec.SetField(logs.FieldMethod, field.TypeString, value)
	}
	if value, ok := lu.mutation.API(); ok {
		_spec.SetField(logs.FieldAPI, field.TypeString, value)
	}
	if value, ok := lu.mutation.Success(); ok {
		_spec.SetField(logs.FieldSuccess, field.TypeBool, value)
	}
	if value, ok := lu.mutation.ReqContent(); ok {
		_spec.SetField(logs.FieldReqContent, field.TypeString, value)
	}
	if lu.mutation.ReqContentCleared() {
		_spec.ClearField(logs.FieldReqContent, field.TypeString)
	}
	if value, ok := lu.mutation.RespContent(); ok {
		_spec.SetField(logs.FieldRespContent, field.TypeString, value)
	}
	if lu.mutation.RespContentCleared() {
		_spec.ClearField(logs.FieldRespContent, field.TypeString)
	}
	if value, ok := lu.mutation.IP(); ok {
		_spec.SetField(logs.FieldIP, field.TypeString, value)
	}
	if lu.mutation.IPCleared() {
		_spec.ClearField(logs.FieldIP, field.TypeString)
	}
	if value, ok := lu.mutation.UserAgent(); ok {
		_spec.SetField(logs.FieldUserAgent, field.TypeString, value)
	}
	if lu.mutation.UserAgentCleared() {
		_spec.ClearField(logs.FieldUserAgent, field.TypeString)
	}
	if value, ok := lu.mutation.Operator(); ok {
		_spec.SetField(logs.FieldOperator, field.TypeString, value)
	}
	if lu.mutation.OperatorCleared() {
		_spec.ClearField(logs.FieldOperator, field.TypeString)
	}
	if value, ok := lu.mutation.Time(); ok {
		_spec.SetField(logs.FieldTime, field.TypeInt, value)
	}
	if value, ok := lu.mutation.AddedTime(); ok {
		_spec.AddField(logs.FieldTime, field.TypeInt, value)
	}
	if lu.mutation.TimeCleared() {
		_spec.ClearField(logs.FieldTime, field.TypeInt)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LogsUpdateOne is the builder for updating a single Logs entity.
type LogsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LogsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LogsUpdateOne) SetUpdatedAt(t time.Time) *LogsUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetDeleteAt sets the "delete_at" field.
func (luo *LogsUpdateOne) SetDeleteAt(t time.Time) *LogsUpdateOne {
	luo.mutation.SetDeleteAt(t)
	return luo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableDeleteAt(t *time.Time) *LogsUpdateOne {
	if t != nil {
		luo.SetDeleteAt(*t)
	}
	return luo
}

// SetCreatedID sets the "created_id" field.
func (luo *LogsUpdateOne) SetCreatedID(i int64) *LogsUpdateOne {
	luo.mutation.ResetCreatedID()
	luo.mutation.SetCreatedID(i)
	return luo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableCreatedID(i *int64) *LogsUpdateOne {
	if i != nil {
		luo.SetCreatedID(*i)
	}
	return luo
}

// AddCreatedID adds i to the "created_id" field.
func (luo *LogsUpdateOne) AddCreatedID(i int64) *LogsUpdateOne {
	luo.mutation.AddCreatedID(i)
	return luo
}

// SetType sets the "type" field.
func (luo *LogsUpdateOne) SetType(s string) *LogsUpdateOne {
	luo.mutation.SetType(s)
	return luo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableType(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetType(*s)
	}
	return luo
}

// SetMethod sets the "method" field.
func (luo *LogsUpdateOne) SetMethod(s string) *LogsUpdateOne {
	luo.mutation.SetMethod(s)
	return luo
}

// SetNillableMethod sets the "method" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableMethod(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetMethod(*s)
	}
	return luo
}

// SetAPI sets the "api" field.
func (luo *LogsUpdateOne) SetAPI(s string) *LogsUpdateOne {
	luo.mutation.SetAPI(s)
	return luo
}

// SetNillableAPI sets the "api" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableAPI(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetAPI(*s)
	}
	return luo
}

// SetSuccess sets the "success" field.
func (luo *LogsUpdateOne) SetSuccess(b bool) *LogsUpdateOne {
	luo.mutation.SetSuccess(b)
	return luo
}

// SetNillableSuccess sets the "success" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableSuccess(b *bool) *LogsUpdateOne {
	if b != nil {
		luo.SetSuccess(*b)
	}
	return luo
}

// SetReqContent sets the "req_content" field.
func (luo *LogsUpdateOne) SetReqContent(s string) *LogsUpdateOne {
	luo.mutation.SetReqContent(s)
	return luo
}

// SetNillableReqContent sets the "req_content" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableReqContent(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetReqContent(*s)
	}
	return luo
}

// ClearReqContent clears the value of the "req_content" field.
func (luo *LogsUpdateOne) ClearReqContent() *LogsUpdateOne {
	luo.mutation.ClearReqContent()
	return luo
}

// SetRespContent sets the "resp_content" field.
func (luo *LogsUpdateOne) SetRespContent(s string) *LogsUpdateOne {
	luo.mutation.SetRespContent(s)
	return luo
}

// SetNillableRespContent sets the "resp_content" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableRespContent(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetRespContent(*s)
	}
	return luo
}

// ClearRespContent clears the value of the "resp_content" field.
func (luo *LogsUpdateOne) ClearRespContent() *LogsUpdateOne {
	luo.mutation.ClearRespContent()
	return luo
}

// SetIP sets the "ip" field.
func (luo *LogsUpdateOne) SetIP(s string) *LogsUpdateOne {
	luo.mutation.SetIP(s)
	return luo
}

// SetNillableIP sets the "ip" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableIP(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetIP(*s)
	}
	return luo
}

// ClearIP clears the value of the "ip" field.
func (luo *LogsUpdateOne) ClearIP() *LogsUpdateOne {
	luo.mutation.ClearIP()
	return luo
}

// SetUserAgent sets the "user_agent" field.
func (luo *LogsUpdateOne) SetUserAgent(s string) *LogsUpdateOne {
	luo.mutation.SetUserAgent(s)
	return luo
}

// SetNillableUserAgent sets the "user_agent" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableUserAgent(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetUserAgent(*s)
	}
	return luo
}

// ClearUserAgent clears the value of the "user_agent" field.
func (luo *LogsUpdateOne) ClearUserAgent() *LogsUpdateOne {
	luo.mutation.ClearUserAgent()
	return luo
}

// SetOperator sets the "operator" field.
func (luo *LogsUpdateOne) SetOperator(s string) *LogsUpdateOne {
	luo.mutation.SetOperator(s)
	return luo
}

// SetNillableOperator sets the "operator" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableOperator(s *string) *LogsUpdateOne {
	if s != nil {
		luo.SetOperator(*s)
	}
	return luo
}

// ClearOperator clears the value of the "operator" field.
func (luo *LogsUpdateOne) ClearOperator() *LogsUpdateOne {
	luo.mutation.ClearOperator()
	return luo
}

// SetTime sets the "time" field.
func (luo *LogsUpdateOne) SetTime(i int) *LogsUpdateOne {
	luo.mutation.ResetTime()
	luo.mutation.SetTime(i)
	return luo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (luo *LogsUpdateOne) SetNillableTime(i *int) *LogsUpdateOne {
	if i != nil {
		luo.SetTime(*i)
	}
	return luo
}

// AddTime adds i to the "time" field.
func (luo *LogsUpdateOne) AddTime(i int) *LogsUpdateOne {
	luo.mutation.AddTime(i)
	return luo
}

// ClearTime clears the value of the "time" field.
func (luo *LogsUpdateOne) ClearTime() *LogsUpdateOne {
	luo.mutation.ClearTime()
	return luo
}

// Mutation returns the LogsMutation object of the builder.
func (luo *LogsUpdateOne) Mutation() *LogsMutation {
	return luo.mutation
}

// Where appends a list predicates to the LogsUpdate builder.
func (luo *LogsUpdateOne) Where(ps ...predicate.Logs) *LogsUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LogsUpdateOne) Select(field string, fields ...string) *LogsUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Logs entity.
func (luo *LogsUpdateOne) Save(ctx context.Context) (*Logs, error) {
	luo.defaults()
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LogsUpdateOne) SaveX(ctx context.Context) *Logs {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LogsUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LogsUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LogsUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := logs.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
}

func (luo *LogsUpdateOne) sqlSave(ctx context.Context) (_node *Logs, err error) {
	_spec := sqlgraph.NewUpdateSpec(logs.Table, logs.Columns, sqlgraph.NewFieldSpec(logs.FieldID, field.TypeInt64))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Logs.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logs.FieldID)
		for _, f := range fields {
			if !logs.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != logs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(logs.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.DeleteAt(); ok {
		_spec.SetField(logs.FieldDeleteAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.CreatedID(); ok {
		_spec.SetField(logs.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := luo.mutation.AddedCreatedID(); ok {
		_spec.AddField(logs.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := luo.mutation.GetType(); ok {
		_spec.SetField(logs.FieldType, field.TypeString, value)
	}
	if value, ok := luo.mutation.Method(); ok {
		_spec.SetField(logs.FieldMethod, field.TypeString, value)
	}
	if value, ok := luo.mutation.API(); ok {
		_spec.SetField(logs.FieldAPI, field.TypeString, value)
	}
	if value, ok := luo.mutation.Success(); ok {
		_spec.SetField(logs.FieldSuccess, field.TypeBool, value)
	}
	if value, ok := luo.mutation.ReqContent(); ok {
		_spec.SetField(logs.FieldReqContent, field.TypeString, value)
	}
	if luo.mutation.ReqContentCleared() {
		_spec.ClearField(logs.FieldReqContent, field.TypeString)
	}
	if value, ok := luo.mutation.RespContent(); ok {
		_spec.SetField(logs.FieldRespContent, field.TypeString, value)
	}
	if luo.mutation.RespContentCleared() {
		_spec.ClearField(logs.FieldRespContent, field.TypeString)
	}
	if value, ok := luo.mutation.IP(); ok {
		_spec.SetField(logs.FieldIP, field.TypeString, value)
	}
	if luo.mutation.IPCleared() {
		_spec.ClearField(logs.FieldIP, field.TypeString)
	}
	if value, ok := luo.mutation.UserAgent(); ok {
		_spec.SetField(logs.FieldUserAgent, field.TypeString, value)
	}
	if luo.mutation.UserAgentCleared() {
		_spec.ClearField(logs.FieldUserAgent, field.TypeString)
	}
	if value, ok := luo.mutation.Operator(); ok {
		_spec.SetField(logs.FieldOperator, field.TypeString, value)
	}
	if luo.mutation.OperatorCleared() {
		_spec.ClearField(logs.FieldOperator, field.TypeString)
	}
	if value, ok := luo.mutation.Time(); ok {
		_spec.SetField(logs.FieldTime, field.TypeInt, value)
	}
	if value, ok := luo.mutation.AddedTime(); ok {
		_spec.AddField(logs.FieldTime, field.TypeInt, value)
	}
	if luo.mutation.TimeCleared() {
		_spec.ClearField(logs.FieldTime, field.TypeInt)
	}
	_node = &Logs{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{logs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
