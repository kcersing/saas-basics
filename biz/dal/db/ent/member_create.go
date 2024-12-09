// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"saas/biz/dal/db/ent/contestparticipant"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/membercontract"
	"saas/biz/dal/db/ent/memberdetails"
	"saas/biz/dal/db/ent/membernote"
	"saas/biz/dal/db/ent/order"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberCreate is the builder for creating a Member entity.
type MemberCreate struct {
	config
	mutation *MemberMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (mc *MemberCreate) SetCreatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCreatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *MemberCreate) SetUpdatedAt(t time.Time) *MemberCreate {
	mc.mutation.SetUpdatedAt(t)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *MemberCreate) SetNillableUpdatedAt(t *time.Time) *MemberCreate {
	if t != nil {
		mc.SetUpdatedAt(*t)
	}
	return mc
}

// SetDelete sets the "delete" field.
func (mc *MemberCreate) SetDelete(i int64) *MemberCreate {
	mc.mutation.SetDelete(i)
	return mc
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (mc *MemberCreate) SetNillableDelete(i *int64) *MemberCreate {
	if i != nil {
		mc.SetDelete(*i)
	}
	return mc
}

// SetCreatedID sets the "created_id" field.
func (mc *MemberCreate) SetCreatedID(i int64) *MemberCreate {
	mc.mutation.SetCreatedID(i)
	return mc
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCreatedID(i *int64) *MemberCreate {
	if i != nil {
		mc.SetCreatedID(*i)
	}
	return mc
}

// SetStatus sets the "status" field.
func (mc *MemberCreate) SetStatus(i int64) *MemberCreate {
	mc.mutation.SetStatus(i)
	return mc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mc *MemberCreate) SetNillableStatus(i *int64) *MemberCreate {
	if i != nil {
		mc.SetStatus(*i)
	}
	return mc
}

// SetPassword sets the "password" field.
func (mc *MemberCreate) SetPassword(s string) *MemberCreate {
	mc.mutation.SetPassword(s)
	return mc
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (mc *MemberCreate) SetNillablePassword(s *string) *MemberCreate {
	if s != nil {
		mc.SetPassword(*s)
	}
	return mc
}

// SetName sets the "name" field.
func (mc *MemberCreate) SetName(s string) *MemberCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (mc *MemberCreate) SetNillableName(s *string) *MemberCreate {
	if s != nil {
		mc.SetName(*s)
	}
	return mc
}

// SetNickname sets the "nickname" field.
func (mc *MemberCreate) SetNickname(s string) *MemberCreate {
	mc.mutation.SetNickname(s)
	return mc
}

// SetNillableNickname sets the "nickname" field if the given value is not nil.
func (mc *MemberCreate) SetNillableNickname(s *string) *MemberCreate {
	if s != nil {
		mc.SetNickname(*s)
	}
	return mc
}

// SetMobile sets the "mobile" field.
func (mc *MemberCreate) SetMobile(s string) *MemberCreate {
	mc.mutation.SetMobile(s)
	return mc
}

// SetNillableMobile sets the "mobile" field if the given value is not nil.
func (mc *MemberCreate) SetNillableMobile(s *string) *MemberCreate {
	if s != nil {
		mc.SetMobile(*s)
	}
	return mc
}

// SetAvatar sets the "avatar" field.
func (mc *MemberCreate) SetAvatar(s string) *MemberCreate {
	mc.mutation.SetAvatar(s)
	return mc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (mc *MemberCreate) SetNillableAvatar(s *string) *MemberCreate {
	if s != nil {
		mc.SetAvatar(*s)
	}
	return mc
}

// SetCondition sets the "condition" field.
func (mc *MemberCreate) SetCondition(i int64) *MemberCreate {
	mc.mutation.SetCondition(i)
	return mc
}

// SetNillableCondition sets the "condition" field if the given value is not nil.
func (mc *MemberCreate) SetNillableCondition(i *int64) *MemberCreate {
	if i != nil {
		mc.SetCondition(*i)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MemberCreate) SetID(i int64) *MemberCreate {
	mc.mutation.SetID(i)
	return mc
}

// AddMemberDetailIDs adds the "member_details" edge to the MemberDetails entity by IDs.
func (mc *MemberCreate) AddMemberDetailIDs(ids ...int64) *MemberCreate {
	mc.mutation.AddMemberDetailIDs(ids...)
	return mc
}

// AddMemberDetails adds the "member_details" edges to the MemberDetails entity.
func (mc *MemberCreate) AddMemberDetails(m ...*MemberDetails) *MemberCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMemberDetailIDs(ids...)
}

// AddMemberNoteIDs adds the "member_notes" edge to the MemberNote entity by IDs.
func (mc *MemberCreate) AddMemberNoteIDs(ids ...int64) *MemberCreate {
	mc.mutation.AddMemberNoteIDs(ids...)
	return mc
}

// AddMemberNotes adds the "member_notes" edges to the MemberNote entity.
func (mc *MemberCreate) AddMemberNotes(m ...*MemberNote) *MemberCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMemberNoteIDs(ids...)
}

// AddMemberOrderIDs adds the "member_orders" edge to the Order entity by IDs.
func (mc *MemberCreate) AddMemberOrderIDs(ids ...int64) *MemberCreate {
	mc.mutation.AddMemberOrderIDs(ids...)
	return mc
}

// AddMemberOrders adds the "member_orders" edges to the Order entity.
func (mc *MemberCreate) AddMemberOrders(o ...*Order) *MemberCreate {
	ids := make([]int64, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mc.AddMemberOrderIDs(ids...)
}

// AddMemberEntryIDs adds the "member_entry" edge to the EntryLogs entity by IDs.
func (mc *MemberCreate) AddMemberEntryIDs(ids ...int64) *MemberCreate {
	mc.mutation.AddMemberEntryIDs(ids...)
	return mc
}

// AddMemberEntry adds the "member_entry" edges to the EntryLogs entity.
func (mc *MemberCreate) AddMemberEntry(e ...*EntryLogs) *MemberCreate {
	ids := make([]int64, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return mc.AddMemberEntryIDs(ids...)
}

// AddMemberContentIDs adds the "member_contents" edge to the MemberContract entity by IDs.
func (mc *MemberCreate) AddMemberContentIDs(ids ...int64) *MemberCreate {
	mc.mutation.AddMemberContentIDs(ids...)
	return mc
}

// AddMemberContents adds the "member_contents" edges to the MemberContract entity.
func (mc *MemberCreate) AddMemberContents(m ...*MemberContract) *MemberCreate {
	ids := make([]int64, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mc.AddMemberContentIDs(ids...)
}

// AddParticipantIDs adds the "participants" edge to the ContestParticipant entity by IDs.
func (mc *MemberCreate) AddParticipantIDs(ids ...int64) *MemberCreate {
	mc.mutation.AddParticipantIDs(ids...)
	return mc
}

// AddParticipants adds the "participants" edges to the ContestParticipant entity.
func (mc *MemberCreate) AddParticipants(c ...*ContestParticipant) *MemberCreate {
	ids := make([]int64, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mc.AddParticipantIDs(ids...)
}

// Mutation returns the MemberMutation object of the builder.
func (mc *MemberCreate) Mutation() *MemberMutation {
	return mc.mutation
}

// Save creates the Member in the database.
func (mc *MemberCreate) Save(ctx context.Context) (*Member, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MemberCreate) SaveX(ctx context.Context) *Member {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MemberCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MemberCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MemberCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := member.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := member.DefaultUpdatedAt()
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.Delete(); !ok {
		v := member.DefaultDelete
		mc.mutation.SetDelete(v)
	}
	if _, ok := mc.mutation.CreatedID(); !ok {
		v := member.DefaultCreatedID
		mc.mutation.SetCreatedID(v)
	}
	if _, ok := mc.mutation.Status(); !ok {
		v := member.DefaultStatus
		mc.mutation.SetStatus(v)
	}
	if _, ok := mc.mutation.Avatar(); !ok {
		v := member.DefaultAvatar
		mc.mutation.SetAvatar(v)
	}
	if _, ok := mc.mutation.Condition(); !ok {
		v := member.DefaultCondition
		mc.mutation.SetCondition(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MemberCreate) check() error {
	return nil
}

func (mc *MemberCreate) sqlSave(ctx context.Context) (*Member, error) {
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

func (mc *MemberCreate) createSpec() (*Member, *sqlgraph.CreateSpec) {
	var (
		_node = &Member{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(member.Table, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Delete(); ok {
		_spec.SetField(member.FieldDelete, field.TypeInt64, value)
		_node.Delete = value
	}
	if value, ok := mc.mutation.CreatedID(); ok {
		_spec.SetField(member.FieldCreatedID, field.TypeInt64, value)
		_node.CreatedID = value
	}
	if value, ok := mc.mutation.Status(); ok {
		_spec.SetField(member.FieldStatus, field.TypeInt64, value)
		_node.Status = value
	}
	if value, ok := mc.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(member.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.Nickname(); ok {
		_spec.SetField(member.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := mc.mutation.Mobile(); ok {
		_spec.SetField(member.FieldMobile, field.TypeString, value)
		_node.Mobile = value
	}
	if value, ok := mc.mutation.Avatar(); ok {
		_spec.SetField(member.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := mc.mutation.Condition(); ok {
		_spec.SetField(member.FieldCondition, field.TypeInt64, value)
		_node.Condition = value
	}
	if nodes := mc.mutation.MemberDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberDetailsTable,
			Columns: []string{member.MemberDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(memberdetails.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MemberNotesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberNotesTable,
			Columns: []string{member.MemberNotesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membernote.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MemberOrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberOrdersTable,
			Columns: []string{member.MemberOrdersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(order.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MemberEntryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberEntryTable,
			Columns: []string{member.MemberEntryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MemberContentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   member.MemberContentsTable,
			Columns: []string{member.MemberContentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(membercontract.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   member.ParticipantsTable,
			Columns: member.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contestparticipant.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MemberCreateBulk is the builder for creating many Member entities in bulk.
type MemberCreateBulk struct {
	config
	err      error
	builders []*MemberCreate
}

// Save creates the Member entities in the database.
func (mcb *MemberCreateBulk) Save(ctx context.Context) ([]*Member, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Member, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MemberMutation)
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
func (mcb *MemberCreateBulk) SaveX(ctx context.Context) []*Member {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MemberCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MemberCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
