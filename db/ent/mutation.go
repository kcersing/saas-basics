// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas_basics/db/ent/predicate"
	"saas_basics/db/ent/user"
	"sync"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeUser = "User"
)

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	password      *string
	gender        *int
	addgender     *int
	age           *int
	addage        *int
	introduce     *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id int) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ClearName clears the value of the "name" field.
func (m *UserMutation) ClearName() {
	m.name = nil
	m.clearedFields[user.FieldName] = struct{}{}
}

// NameCleared returns if the "name" field was cleared in this mutation.
func (m *UserMutation) NameCleared() bool {
	_, ok := m.clearedFields[user.FieldName]
	return ok
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
	delete(m.clearedFields, user.FieldName)
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ClearPassword clears the value of the "password" field.
func (m *UserMutation) ClearPassword() {
	m.password = nil
	m.clearedFields[user.FieldPassword] = struct{}{}
}

// PasswordCleared returns if the "password" field was cleared in this mutation.
func (m *UserMutation) PasswordCleared() bool {
	_, ok := m.clearedFields[user.FieldPassword]
	return ok
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
	delete(m.clearedFields, user.FieldPassword)
}

// SetGender sets the "gender" field.
func (m *UserMutation) SetGender(i int) {
	m.gender = &i
	m.addgender = nil
}

// Gender returns the value of the "gender" field in the mutation.
func (m *UserMutation) Gender() (r int, exists bool) {
	v := m.gender
	if v == nil {
		return
	}
	return *v, true
}

// OldGender returns the old "gender" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldGender(ctx context.Context) (v *int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGender is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGender requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGender: %w", err)
	}
	return oldValue.Gender, nil
}

// AddGender adds i to the "gender" field.
func (m *UserMutation) AddGender(i int) {
	if m.addgender != nil {
		*m.addgender += i
	} else {
		m.addgender = &i
	}
}

// AddedGender returns the value that was added to the "gender" field in this mutation.
func (m *UserMutation) AddedGender() (r int, exists bool) {
	v := m.addgender
	if v == nil {
		return
	}
	return *v, true
}

// ClearGender clears the value of the "gender" field.
func (m *UserMutation) ClearGender() {
	m.gender = nil
	m.addgender = nil
	m.clearedFields[user.FieldGender] = struct{}{}
}

// GenderCleared returns if the "gender" field was cleared in this mutation.
func (m *UserMutation) GenderCleared() bool {
	_, ok := m.clearedFields[user.FieldGender]
	return ok
}

// ResetGender resets all changes to the "gender" field.
func (m *UserMutation) ResetGender() {
	m.gender = nil
	m.addgender = nil
	delete(m.clearedFields, user.FieldGender)
}

// SetAge sets the "age" field.
func (m *UserMutation) SetAge(i int) {
	m.age = &i
	m.addage = nil
}

// Age returns the value of the "age" field in the mutation.
func (m *UserMutation) Age() (r int, exists bool) {
	v := m.age
	if v == nil {
		return
	}
	return *v, true
}

// OldAge returns the old "age" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAge(ctx context.Context) (v *int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAge is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAge requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAge: %w", err)
	}
	return oldValue.Age, nil
}

// AddAge adds i to the "age" field.
func (m *UserMutation) AddAge(i int) {
	if m.addage != nil {
		*m.addage += i
	} else {
		m.addage = &i
	}
}

// AddedAge returns the value that was added to the "age" field in this mutation.
func (m *UserMutation) AddedAge() (r int, exists bool) {
	v := m.addage
	if v == nil {
		return
	}
	return *v, true
}

// ClearAge clears the value of the "age" field.
func (m *UserMutation) ClearAge() {
	m.age = nil
	m.addage = nil
	m.clearedFields[user.FieldAge] = struct{}{}
}

// AgeCleared returns if the "age" field was cleared in this mutation.
func (m *UserMutation) AgeCleared() bool {
	_, ok := m.clearedFields[user.FieldAge]
	return ok
}

// ResetAge resets all changes to the "age" field.
func (m *UserMutation) ResetAge() {
	m.age = nil
	m.addage = nil
	delete(m.clearedFields, user.FieldAge)
}

// SetIntroduce sets the "introduce" field.
func (m *UserMutation) SetIntroduce(s string) {
	m.introduce = &s
}

// Introduce returns the value of the "introduce" field in the mutation.
func (m *UserMutation) Introduce() (r string, exists bool) {
	v := m.introduce
	if v == nil {
		return
	}
	return *v, true
}

// OldIntroduce returns the old "introduce" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldIntroduce(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIntroduce is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIntroduce requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIntroduce: %w", err)
	}
	return oldValue.Introduce, nil
}

// ClearIntroduce clears the value of the "introduce" field.
func (m *UserMutation) ClearIntroduce() {
	m.introduce = nil
	m.clearedFields[user.FieldIntroduce] = struct{}{}
}

// IntroduceCleared returns if the "introduce" field was cleared in this mutation.
func (m *UserMutation) IntroduceCleared() bool {
	_, ok := m.clearedFields[user.FieldIntroduce]
	return ok
}

// ResetIntroduce resets all changes to the "introduce" field.
func (m *UserMutation) ResetIntroduce() {
	m.introduce = nil
	delete(m.clearedFields, user.FieldIntroduce)
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.gender != nil {
		fields = append(fields, user.FieldGender)
	}
	if m.age != nil {
		fields = append(fields, user.FieldAge)
	}
	if m.introduce != nil {
		fields = append(fields, user.FieldIntroduce)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldPassword:
		return m.Password()
	case user.FieldGender:
		return m.Gender()
	case user.FieldAge:
		return m.Age()
	case user.FieldIntroduce:
		return m.Introduce()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldGender:
		return m.OldGender(ctx)
	case user.FieldAge:
		return m.OldAge(ctx)
	case user.FieldIntroduce:
		return m.OldIntroduce(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldGender:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGender(v)
		return nil
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAge(v)
		return nil
	case user.FieldIntroduce:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIntroduce(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addgender != nil {
		fields = append(fields, user.FieldGender)
	}
	if m.addage != nil {
		fields = append(fields, user.FieldAge)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldGender:
		return m.AddedGender()
	case user.FieldAge:
		return m.AddedAge()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldGender:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddGender(v)
		return nil
	case user.FieldAge:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAge(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldName) {
		fields = append(fields, user.FieldName)
	}
	if m.FieldCleared(user.FieldPassword) {
		fields = append(fields, user.FieldPassword)
	}
	if m.FieldCleared(user.FieldGender) {
		fields = append(fields, user.FieldGender)
	}
	if m.FieldCleared(user.FieldAge) {
		fields = append(fields, user.FieldAge)
	}
	if m.FieldCleared(user.FieldIntroduce) {
		fields = append(fields, user.FieldIntroduce)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldName:
		m.ClearName()
		return nil
	case user.FieldPassword:
		m.ClearPassword()
		return nil
	case user.FieldGender:
		m.ClearGender()
		return nil
	case user.FieldAge:
		m.ClearAge()
		return nil
	case user.FieldIntroduce:
		m.ClearIntroduce()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldGender:
		m.ResetGender()
		return nil
	case user.FieldAge:
		m.ResetAge()
		return nil
	case user.FieldIntroduce:
		m.ResetIntroduce()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown User edge %s", name)
}
