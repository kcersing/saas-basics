// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/pkg/db/ent/predicate"
	"saas/pkg/db/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rs/xid"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetAccountID sets the "account_id" field.
func (uu *UserUpdate) SetAccountID(x xid.ID) *UserUpdate {
	uu.mutation.SetAccountID(x)
	return uu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAccountID(x *xid.ID) *UserUpdate {
	if x != nil {
		uu.SetAccountID(*x)
	}
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUsername(s *string) *UserUpdate {
	if s != nil {
		uu.SetUsername(*s)
	}
	return uu
}

// ClearUsername clears the value of the "username" field.
func (uu *UserUpdate) ClearUsername() *UserUpdate {
	uu.mutation.ClearUsername()
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of the "password" field.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.mutation.ClearPassword()
	return uu
}

// SetMobile sets the "mobile" field.
func (uu *UserUpdate) SetMobile(s string) *UserUpdate {
	uu.mutation.SetMobile(s)
	return uu
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uu *UserUpdate) SetNillableMobile(s *string) *UserUpdate {
	if s != nil {
		uu.SetMobile(*s)
	}
	return uu
}

// ClearMobile clears the value of the "mobile" field.
func (uu *UserUpdate) ClearMobile() *UserUpdate {
	uu.mutation.ClearMobile()
	return uu
}

// SetGender sets the "gender" field.
func (uu *UserUpdate) SetGender(s string) *UserUpdate {
	uu.mutation.SetGender(s)
	return uu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uu *UserUpdate) SetNillableGender(s *string) *UserUpdate {
	if s != nil {
		uu.SetGender(*s)
	}
	return uu
}

// ClearGender clears the value of the "gender" field.
func (uu *UserUpdate) ClearGender() *UserUpdate {
	uu.mutation.ClearGender()
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(s string) *UserUpdate {
	uu.mutation.SetAge(s)
	return uu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge(s *string) *UserUpdate {
	if s != nil {
		uu.SetAge(*s)
	}
	return uu
}

// ClearAge clears the value of the "age" field.
func (uu *UserUpdate) ClearAge() *UserUpdate {
	uu.mutation.ClearAge()
	return uu
}

// SetIntroduce sets the "introduce" field.
func (uu *UserUpdate) SetIntroduce(s string) *UserUpdate {
	uu.mutation.SetIntroduce(s)
	return uu
}

// SetNillableIntroduce sets the "introduce" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIntroduce(s *string) *UserUpdate {
	if s != nil {
		uu.SetIntroduce(*s)
	}
	return uu
}

// ClearIntroduce clears the value of the "introduce" field.
func (uu *UserUpdate) ClearIntroduce() *UserUpdate {
	uu.mutation.ClearIntroduce()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.AccountID(); ok {
		_spec.SetField(user.FieldAccountID, field.TypeString, value)
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if uu.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uu.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uu.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uu.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
	}
	if uu.mutation.GenderCleared() {
		_spec.ClearField(user.FieldGender, field.TypeString)
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeString, value)
	}
	if uu.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeString)
	}
	if value, ok := uu.mutation.Introduce(); ok {
		_spec.SetField(user.FieldIntroduce, field.TypeString, value)
	}
	if uu.mutation.IntroduceCleared() {
		_spec.ClearField(user.FieldIntroduce, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAccountID sets the "account_id" field.
func (uuo *UserUpdateOne) SetAccountID(x xid.ID) *UserUpdateOne {
	uuo.mutation.SetAccountID(x)
	return uuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAccountID(x *xid.ID) *UserUpdateOne {
	if x != nil {
		uuo.SetAccountID(*x)
	}
	return uuo
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetNillableUsername sets the "username" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUsername(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetUsername(*s)
	}
	return uuo
}

// ClearUsername clears the value of the "username" field.
func (uuo *UserUpdateOne) ClearUsername() *UserUpdateOne {
	uuo.mutation.ClearUsername()
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of the "password" field.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.mutation.ClearPassword()
	return uuo
}

// SetMobile sets the "mobile" field.
func (uuo *UserUpdateOne) SetMobile(s string) *UserUpdateOne {
	uuo.mutation.SetMobile(s)
	return uuo
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableMobile(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetMobile(*s)
	}
	return uuo
}

// ClearMobile clears the value of the "mobile" field.
func (uuo *UserUpdateOne) ClearMobile() *UserUpdateOne {
	uuo.mutation.ClearMobile()
	return uuo
}

// SetGender sets the "gender" field.
func (uuo *UserUpdateOne) SetGender(s string) *UserUpdateOne {
	uuo.mutation.SetGender(s)
	return uuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableGender(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetGender(*s)
	}
	return uuo
}

// ClearGender clears the value of the "gender" field.
func (uuo *UserUpdateOne) ClearGender() *UserUpdateOne {
	uuo.mutation.ClearGender()
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(s string) *UserUpdateOne {
	uuo.mutation.SetAge(s)
	return uuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAge(*s)
	}
	return uuo
}

// ClearAge clears the value of the "age" field.
func (uuo *UserUpdateOne) ClearAge() *UserUpdateOne {
	uuo.mutation.ClearAge()
	return uuo
}

// SetIntroduce sets the "introduce" field.
func (uuo *UserUpdateOne) SetIntroduce(s string) *UserUpdateOne {
	uuo.mutation.SetIntroduce(s)
	return uuo
}

// SetNillableIntroduce sets the "introduce" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIntroduce(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIntroduce(*s)
	}
	return uuo
}

// ClearIntroduce clears the value of the "introduce" field.
func (uuo *UserUpdateOne) ClearIntroduce() *UserUpdateOne {
	uuo.mutation.ClearIntroduce()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.AccountID(); ok {
		_spec.SetField(user.FieldAccountID, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if uuo.mutation.UsernameCleared() {
		_spec.ClearField(user.FieldUsername, field.TypeString)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if uuo.mutation.PasswordCleared() {
		_spec.ClearField(user.FieldPassword, field.TypeString)
	}
	if value, ok := uuo.mutation.Mobile(); ok {
		_spec.SetField(user.FieldMobile, field.TypeString, value)
	}
	if uuo.mutation.MobileCleared() {
		_spec.ClearField(user.FieldMobile, field.TypeString)
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.SetField(user.FieldGender, field.TypeString, value)
	}
	if uuo.mutation.GenderCleared() {
		_spec.ClearField(user.FieldGender, field.TypeString)
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeString, value)
	}
	if uuo.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeString)
	}
	if value, ok := uuo.mutation.Introduce(); ok {
		_spec.SetField(user.FieldIntroduce, field.TypeString, value)
	}
	if uuo.mutation.IntroduceCleared() {
		_spec.ClearField(user.FieldIntroduce, field.TypeString)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
