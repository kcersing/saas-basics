// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"saas/biz/dal/db/ent/entrylogs"
	"saas/biz/dal/db/ent/member"
	"saas/biz/dal/db/ent/predicate"
	"saas/biz/dal/db/ent/user"
	"saas/biz/dal/db/ent/venue"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// EntryLogsUpdate is the builder for updating EntryLogs entities.
type EntryLogsUpdate struct {
	config
	hooks    []Hook
	mutation *EntryLogsMutation
}

// Where appends a list predicates to the EntryLogsUpdate builder.
func (elu *EntryLogsUpdate) Where(ps ...predicate.EntryLogs) *EntryLogsUpdate {
	elu.mutation.Where(ps...)
	return elu
}

// SetUpdatedAt sets the "updated_at" field.
func (elu *EntryLogsUpdate) SetUpdatedAt(t time.Time) *EntryLogsUpdate {
	elu.mutation.SetUpdatedAt(t)
	return elu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (elu *EntryLogsUpdate) ClearUpdatedAt() *EntryLogsUpdate {
	elu.mutation.ClearUpdatedAt()
	return elu
}

// SetDelete sets the "delete" field.
func (elu *EntryLogsUpdate) SetDelete(i int64) *EntryLogsUpdate {
	elu.mutation.ResetDelete()
	elu.mutation.SetDelete(i)
	return elu
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableDelete(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetDelete(*i)
	}
	return elu
}

// AddDelete adds i to the "delete" field.
func (elu *EntryLogsUpdate) AddDelete(i int64) *EntryLogsUpdate {
	elu.mutation.AddDelete(i)
	return elu
}

// ClearDelete clears the value of the "delete" field.
func (elu *EntryLogsUpdate) ClearDelete() *EntryLogsUpdate {
	elu.mutation.ClearDelete()
	return elu
}

// SetCreatedID sets the "created_id" field.
func (elu *EntryLogsUpdate) SetCreatedID(i int64) *EntryLogsUpdate {
	elu.mutation.ResetCreatedID()
	elu.mutation.SetCreatedID(i)
	return elu
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableCreatedID(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetCreatedID(*i)
	}
	return elu
}

// AddCreatedID adds i to the "created_id" field.
func (elu *EntryLogsUpdate) AddCreatedID(i int64) *EntryLogsUpdate {
	elu.mutation.AddCreatedID(i)
	return elu
}

// ClearCreatedID clears the value of the "created_id" field.
func (elu *EntryLogsUpdate) ClearCreatedID() *EntryLogsUpdate {
	elu.mutation.ClearCreatedID()
	return elu
}

// SetMemberID sets the "member_id" field.
func (elu *EntryLogsUpdate) SetMemberID(i int64) *EntryLogsUpdate {
	elu.mutation.SetMemberID(i)
	return elu
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableMemberID(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetMemberID(*i)
	}
	return elu
}

// ClearMemberID clears the value of the "member_id" field.
func (elu *EntryLogsUpdate) ClearMemberID() *EntryLogsUpdate {
	elu.mutation.ClearMemberID()
	return elu
}

// SetUserID sets the "user_id" field.
func (elu *EntryLogsUpdate) SetUserID(i int64) *EntryLogsUpdate {
	elu.mutation.SetUserID(i)
	return elu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableUserID(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetUserID(*i)
	}
	return elu
}

// ClearUserID clears the value of the "user_id" field.
func (elu *EntryLogsUpdate) ClearUserID() *EntryLogsUpdate {
	elu.mutation.ClearUserID()
	return elu
}

// SetVenueID sets the "venue_id" field.
func (elu *EntryLogsUpdate) SetVenueID(i int64) *EntryLogsUpdate {
	elu.mutation.SetVenueID(i)
	return elu
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableVenueID(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetVenueID(*i)
	}
	return elu
}

// ClearVenueID clears the value of the "venue_id" field.
func (elu *EntryLogsUpdate) ClearVenueID() *EntryLogsUpdate {
	elu.mutation.ClearVenueID()
	return elu
}

// SetMemberProductID sets the "member_product_id" field.
func (elu *EntryLogsUpdate) SetMemberProductID(i int64) *EntryLogsUpdate {
	elu.mutation.ResetMemberProductID()
	elu.mutation.SetMemberProductID(i)
	return elu
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableMemberProductID(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetMemberProductID(*i)
	}
	return elu
}

// AddMemberProductID adds i to the "member_product_id" field.
func (elu *EntryLogsUpdate) AddMemberProductID(i int64) *EntryLogsUpdate {
	elu.mutation.AddMemberProductID(i)
	return elu
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (elu *EntryLogsUpdate) ClearMemberProductID() *EntryLogsUpdate {
	elu.mutation.ClearMemberProductID()
	return elu
}

// SetMemberPropertyID sets the "member_property_id" field.
func (elu *EntryLogsUpdate) SetMemberPropertyID(i int64) *EntryLogsUpdate {
	elu.mutation.ResetMemberPropertyID()
	elu.mutation.SetMemberPropertyID(i)
	return elu
}

// SetNillableMemberPropertyID sets the "member_property_id" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableMemberPropertyID(i *int64) *EntryLogsUpdate {
	if i != nil {
		elu.SetMemberPropertyID(*i)
	}
	return elu
}

// AddMemberPropertyID adds i to the "member_property_id" field.
func (elu *EntryLogsUpdate) AddMemberPropertyID(i int64) *EntryLogsUpdate {
	elu.mutation.AddMemberPropertyID(i)
	return elu
}

// ClearMemberPropertyID clears the value of the "member_property_id" field.
func (elu *EntryLogsUpdate) ClearMemberPropertyID() *EntryLogsUpdate {
	elu.mutation.ClearMemberPropertyID()
	return elu
}

// SetEntryTime sets the "entry_time" field.
func (elu *EntryLogsUpdate) SetEntryTime(t time.Time) *EntryLogsUpdate {
	elu.mutation.SetEntryTime(t)
	return elu
}

// SetNillableEntryTime sets the "entry_time" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableEntryTime(t *time.Time) *EntryLogsUpdate {
	if t != nil {
		elu.SetEntryTime(*t)
	}
	return elu
}

// ClearEntryTime clears the value of the "entry_time" field.
func (elu *EntryLogsUpdate) ClearEntryTime() *EntryLogsUpdate {
	elu.mutation.ClearEntryTime()
	return elu
}

// SetLeavingTime sets the "leaving_time" field.
func (elu *EntryLogsUpdate) SetLeavingTime(t time.Time) *EntryLogsUpdate {
	elu.mutation.SetLeavingTime(t)
	return elu
}

// SetNillableLeavingTime sets the "leaving_time" field if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableLeavingTime(t *time.Time) *EntryLogsUpdate {
	if t != nil {
		elu.SetLeavingTime(*t)
	}
	return elu
}

// ClearLeavingTime clears the value of the "leaving_time" field.
func (elu *EntryLogsUpdate) ClearLeavingTime() *EntryLogsUpdate {
	elu.mutation.ClearLeavingTime()
	return elu
}

// SetVenuesID sets the "venues" edge to the Venue entity by ID.
func (elu *EntryLogsUpdate) SetVenuesID(id int64) *EntryLogsUpdate {
	elu.mutation.SetVenuesID(id)
	return elu
}

// SetNillableVenuesID sets the "venues" edge to the Venue entity by ID if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableVenuesID(id *int64) *EntryLogsUpdate {
	if id != nil {
		elu = elu.SetVenuesID(*id)
	}
	return elu
}

// SetVenues sets the "venues" edge to the Venue entity.
func (elu *EntryLogsUpdate) SetVenues(v *Venue) *EntryLogsUpdate {
	return elu.SetVenuesID(v.ID)
}

// SetMembersID sets the "members" edge to the Member entity by ID.
func (elu *EntryLogsUpdate) SetMembersID(id int64) *EntryLogsUpdate {
	elu.mutation.SetMembersID(id)
	return elu
}

// SetNillableMembersID sets the "members" edge to the Member entity by ID if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableMembersID(id *int64) *EntryLogsUpdate {
	if id != nil {
		elu = elu.SetMembersID(*id)
	}
	return elu
}

// SetMembers sets the "members" edge to the Member entity.
func (elu *EntryLogsUpdate) SetMembers(m *Member) *EntryLogsUpdate {
	return elu.SetMembersID(m.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (elu *EntryLogsUpdate) SetUsersID(id int64) *EntryLogsUpdate {
	elu.mutation.SetUsersID(id)
	return elu
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (elu *EntryLogsUpdate) SetNillableUsersID(id *int64) *EntryLogsUpdate {
	if id != nil {
		elu = elu.SetUsersID(*id)
	}
	return elu
}

// SetUsers sets the "users" edge to the User entity.
func (elu *EntryLogsUpdate) SetUsers(u *User) *EntryLogsUpdate {
	return elu.SetUsersID(u.ID)
}

// Mutation returns the EntryLogsMutation object of the builder.
func (elu *EntryLogsUpdate) Mutation() *EntryLogsMutation {
	return elu.mutation
}

// ClearVenues clears the "venues" edge to the Venue entity.
func (elu *EntryLogsUpdate) ClearVenues() *EntryLogsUpdate {
	elu.mutation.ClearVenues()
	return elu
}

// ClearMembers clears the "members" edge to the Member entity.
func (elu *EntryLogsUpdate) ClearMembers() *EntryLogsUpdate {
	elu.mutation.ClearMembers()
	return elu
}

// ClearUsers clears the "users" edge to the User entity.
func (elu *EntryLogsUpdate) ClearUsers() *EntryLogsUpdate {
	elu.mutation.ClearUsers()
	return elu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (elu *EntryLogsUpdate) Save(ctx context.Context) (int, error) {
	elu.defaults()
	return withHooks(ctx, elu.sqlSave, elu.mutation, elu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (elu *EntryLogsUpdate) SaveX(ctx context.Context) int {
	affected, err := elu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (elu *EntryLogsUpdate) Exec(ctx context.Context) error {
	_, err := elu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elu *EntryLogsUpdate) ExecX(ctx context.Context) {
	if err := elu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (elu *EntryLogsUpdate) defaults() {
	if _, ok := elu.mutation.UpdatedAt(); !ok && !elu.mutation.UpdatedAtCleared() {
		v := entrylogs.UpdateDefaultUpdatedAt()
		elu.mutation.SetUpdatedAt(v)
	}
}

func (elu *EntryLogsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(entrylogs.Table, entrylogs.Columns, sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64))
	if ps := elu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if elu.mutation.CreatedAtCleared() {
		_spec.ClearField(entrylogs.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := elu.mutation.UpdatedAt(); ok {
		_spec.SetField(entrylogs.FieldUpdatedAt, field.TypeTime, value)
	}
	if elu.mutation.UpdatedAtCleared() {
		_spec.ClearField(entrylogs.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := elu.mutation.Delete(); ok {
		_spec.SetField(entrylogs.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := elu.mutation.AddedDelete(); ok {
		_spec.AddField(entrylogs.FieldDelete, field.TypeInt64, value)
	}
	if elu.mutation.DeleteCleared() {
		_spec.ClearField(entrylogs.FieldDelete, field.TypeInt64)
	}
	if value, ok := elu.mutation.CreatedID(); ok {
		_spec.SetField(entrylogs.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := elu.mutation.AddedCreatedID(); ok {
		_spec.AddField(entrylogs.FieldCreatedID, field.TypeInt64, value)
	}
	if elu.mutation.CreatedIDCleared() {
		_spec.ClearField(entrylogs.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := elu.mutation.MemberProductID(); ok {
		_spec.SetField(entrylogs.FieldMemberProductID, field.TypeInt64, value)
	}
	if value, ok := elu.mutation.AddedMemberProductID(); ok {
		_spec.AddField(entrylogs.FieldMemberProductID, field.TypeInt64, value)
	}
	if elu.mutation.MemberProductIDCleared() {
		_spec.ClearField(entrylogs.FieldMemberProductID, field.TypeInt64)
	}
	if value, ok := elu.mutation.MemberPropertyID(); ok {
		_spec.SetField(entrylogs.FieldMemberPropertyID, field.TypeInt64, value)
	}
	if value, ok := elu.mutation.AddedMemberPropertyID(); ok {
		_spec.AddField(entrylogs.FieldMemberPropertyID, field.TypeInt64, value)
	}
	if elu.mutation.MemberPropertyIDCleared() {
		_spec.ClearField(entrylogs.FieldMemberPropertyID, field.TypeInt64)
	}
	if value, ok := elu.mutation.EntryTime(); ok {
		_spec.SetField(entrylogs.FieldEntryTime, field.TypeTime, value)
	}
	if elu.mutation.EntryTimeCleared() {
		_spec.ClearField(entrylogs.FieldEntryTime, field.TypeTime)
	}
	if value, ok := elu.mutation.LeavingTime(); ok {
		_spec.SetField(entrylogs.FieldLeavingTime, field.TypeTime, value)
	}
	if elu.mutation.LeavingTimeCleared() {
		_spec.ClearField(entrylogs.FieldLeavingTime, field.TypeTime)
	}
	if elu.mutation.VenuesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.VenuesTable,
			Columns: []string{entrylogs.VenuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := elu.mutation.VenuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.VenuesTable,
			Columns: []string{entrylogs.VenuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if elu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.MembersTable,
			Columns: []string{entrylogs.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := elu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.MembersTable,
			Columns: []string{entrylogs.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if elu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.UsersTable,
			Columns: []string{entrylogs.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := elu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.UsersTable,
			Columns: []string{entrylogs.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, elu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entrylogs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	elu.mutation.done = true
	return n, nil
}

// EntryLogsUpdateOne is the builder for updating a single EntryLogs entity.
type EntryLogsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntryLogsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (eluo *EntryLogsUpdateOne) SetUpdatedAt(t time.Time) *EntryLogsUpdateOne {
	eluo.mutation.SetUpdatedAt(t)
	return eluo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (eluo *EntryLogsUpdateOne) ClearUpdatedAt() *EntryLogsUpdateOne {
	eluo.mutation.ClearUpdatedAt()
	return eluo
}

// SetDelete sets the "delete" field.
func (eluo *EntryLogsUpdateOne) SetDelete(i int64) *EntryLogsUpdateOne {
	eluo.mutation.ResetDelete()
	eluo.mutation.SetDelete(i)
	return eluo
}

// SetNillableDelete sets the "delete" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableDelete(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetDelete(*i)
	}
	return eluo
}

// AddDelete adds i to the "delete" field.
func (eluo *EntryLogsUpdateOne) AddDelete(i int64) *EntryLogsUpdateOne {
	eluo.mutation.AddDelete(i)
	return eluo
}

// ClearDelete clears the value of the "delete" field.
func (eluo *EntryLogsUpdateOne) ClearDelete() *EntryLogsUpdateOne {
	eluo.mutation.ClearDelete()
	return eluo
}

// SetCreatedID sets the "created_id" field.
func (eluo *EntryLogsUpdateOne) SetCreatedID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.ResetCreatedID()
	eluo.mutation.SetCreatedID(i)
	return eluo
}

// SetNillableCreatedID sets the "created_id" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableCreatedID(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetCreatedID(*i)
	}
	return eluo
}

// AddCreatedID adds i to the "created_id" field.
func (eluo *EntryLogsUpdateOne) AddCreatedID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.AddCreatedID(i)
	return eluo
}

// ClearCreatedID clears the value of the "created_id" field.
func (eluo *EntryLogsUpdateOne) ClearCreatedID() *EntryLogsUpdateOne {
	eluo.mutation.ClearCreatedID()
	return eluo
}

// SetMemberID sets the "member_id" field.
func (eluo *EntryLogsUpdateOne) SetMemberID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.SetMemberID(i)
	return eluo
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableMemberID(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetMemberID(*i)
	}
	return eluo
}

// ClearMemberID clears the value of the "member_id" field.
func (eluo *EntryLogsUpdateOne) ClearMemberID() *EntryLogsUpdateOne {
	eluo.mutation.ClearMemberID()
	return eluo
}

// SetUserID sets the "user_id" field.
func (eluo *EntryLogsUpdateOne) SetUserID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.SetUserID(i)
	return eluo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableUserID(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetUserID(*i)
	}
	return eluo
}

// ClearUserID clears the value of the "user_id" field.
func (eluo *EntryLogsUpdateOne) ClearUserID() *EntryLogsUpdateOne {
	eluo.mutation.ClearUserID()
	return eluo
}

// SetVenueID sets the "venue_id" field.
func (eluo *EntryLogsUpdateOne) SetVenueID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.SetVenueID(i)
	return eluo
}

// SetNillableVenueID sets the "venue_id" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableVenueID(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetVenueID(*i)
	}
	return eluo
}

// ClearVenueID clears the value of the "venue_id" field.
func (eluo *EntryLogsUpdateOne) ClearVenueID() *EntryLogsUpdateOne {
	eluo.mutation.ClearVenueID()
	return eluo
}

// SetMemberProductID sets the "member_product_id" field.
func (eluo *EntryLogsUpdateOne) SetMemberProductID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.ResetMemberProductID()
	eluo.mutation.SetMemberProductID(i)
	return eluo
}

// SetNillableMemberProductID sets the "member_product_id" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableMemberProductID(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetMemberProductID(*i)
	}
	return eluo
}

// AddMemberProductID adds i to the "member_product_id" field.
func (eluo *EntryLogsUpdateOne) AddMemberProductID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.AddMemberProductID(i)
	return eluo
}

// ClearMemberProductID clears the value of the "member_product_id" field.
func (eluo *EntryLogsUpdateOne) ClearMemberProductID() *EntryLogsUpdateOne {
	eluo.mutation.ClearMemberProductID()
	return eluo
}

// SetMemberPropertyID sets the "member_property_id" field.
func (eluo *EntryLogsUpdateOne) SetMemberPropertyID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.ResetMemberPropertyID()
	eluo.mutation.SetMemberPropertyID(i)
	return eluo
}

// SetNillableMemberPropertyID sets the "member_property_id" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableMemberPropertyID(i *int64) *EntryLogsUpdateOne {
	if i != nil {
		eluo.SetMemberPropertyID(*i)
	}
	return eluo
}

// AddMemberPropertyID adds i to the "member_property_id" field.
func (eluo *EntryLogsUpdateOne) AddMemberPropertyID(i int64) *EntryLogsUpdateOne {
	eluo.mutation.AddMemberPropertyID(i)
	return eluo
}

// ClearMemberPropertyID clears the value of the "member_property_id" field.
func (eluo *EntryLogsUpdateOne) ClearMemberPropertyID() *EntryLogsUpdateOne {
	eluo.mutation.ClearMemberPropertyID()
	return eluo
}

// SetEntryTime sets the "entry_time" field.
func (eluo *EntryLogsUpdateOne) SetEntryTime(t time.Time) *EntryLogsUpdateOne {
	eluo.mutation.SetEntryTime(t)
	return eluo
}

// SetNillableEntryTime sets the "entry_time" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableEntryTime(t *time.Time) *EntryLogsUpdateOne {
	if t != nil {
		eluo.SetEntryTime(*t)
	}
	return eluo
}

// ClearEntryTime clears the value of the "entry_time" field.
func (eluo *EntryLogsUpdateOne) ClearEntryTime() *EntryLogsUpdateOne {
	eluo.mutation.ClearEntryTime()
	return eluo
}

// SetLeavingTime sets the "leaving_time" field.
func (eluo *EntryLogsUpdateOne) SetLeavingTime(t time.Time) *EntryLogsUpdateOne {
	eluo.mutation.SetLeavingTime(t)
	return eluo
}

// SetNillableLeavingTime sets the "leaving_time" field if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableLeavingTime(t *time.Time) *EntryLogsUpdateOne {
	if t != nil {
		eluo.SetLeavingTime(*t)
	}
	return eluo
}

// ClearLeavingTime clears the value of the "leaving_time" field.
func (eluo *EntryLogsUpdateOne) ClearLeavingTime() *EntryLogsUpdateOne {
	eluo.mutation.ClearLeavingTime()
	return eluo
}

// SetVenuesID sets the "venues" edge to the Venue entity by ID.
func (eluo *EntryLogsUpdateOne) SetVenuesID(id int64) *EntryLogsUpdateOne {
	eluo.mutation.SetVenuesID(id)
	return eluo
}

// SetNillableVenuesID sets the "venues" edge to the Venue entity by ID if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableVenuesID(id *int64) *EntryLogsUpdateOne {
	if id != nil {
		eluo = eluo.SetVenuesID(*id)
	}
	return eluo
}

// SetVenues sets the "venues" edge to the Venue entity.
func (eluo *EntryLogsUpdateOne) SetVenues(v *Venue) *EntryLogsUpdateOne {
	return eluo.SetVenuesID(v.ID)
}

// SetMembersID sets the "members" edge to the Member entity by ID.
func (eluo *EntryLogsUpdateOne) SetMembersID(id int64) *EntryLogsUpdateOne {
	eluo.mutation.SetMembersID(id)
	return eluo
}

// SetNillableMembersID sets the "members" edge to the Member entity by ID if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableMembersID(id *int64) *EntryLogsUpdateOne {
	if id != nil {
		eluo = eluo.SetMembersID(*id)
	}
	return eluo
}

// SetMembers sets the "members" edge to the Member entity.
func (eluo *EntryLogsUpdateOne) SetMembers(m *Member) *EntryLogsUpdateOne {
	return eluo.SetMembersID(m.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (eluo *EntryLogsUpdateOne) SetUsersID(id int64) *EntryLogsUpdateOne {
	eluo.mutation.SetUsersID(id)
	return eluo
}

// SetNillableUsersID sets the "users" edge to the User entity by ID if the given value is not nil.
func (eluo *EntryLogsUpdateOne) SetNillableUsersID(id *int64) *EntryLogsUpdateOne {
	if id != nil {
		eluo = eluo.SetUsersID(*id)
	}
	return eluo
}

// SetUsers sets the "users" edge to the User entity.
func (eluo *EntryLogsUpdateOne) SetUsers(u *User) *EntryLogsUpdateOne {
	return eluo.SetUsersID(u.ID)
}

// Mutation returns the EntryLogsMutation object of the builder.
func (eluo *EntryLogsUpdateOne) Mutation() *EntryLogsMutation {
	return eluo.mutation
}

// ClearVenues clears the "venues" edge to the Venue entity.
func (eluo *EntryLogsUpdateOne) ClearVenues() *EntryLogsUpdateOne {
	eluo.mutation.ClearVenues()
	return eluo
}

// ClearMembers clears the "members" edge to the Member entity.
func (eluo *EntryLogsUpdateOne) ClearMembers() *EntryLogsUpdateOne {
	eluo.mutation.ClearMembers()
	return eluo
}

// ClearUsers clears the "users" edge to the User entity.
func (eluo *EntryLogsUpdateOne) ClearUsers() *EntryLogsUpdateOne {
	eluo.mutation.ClearUsers()
	return eluo
}

// Where appends a list predicates to the EntryLogsUpdate builder.
func (eluo *EntryLogsUpdateOne) Where(ps ...predicate.EntryLogs) *EntryLogsUpdateOne {
	eluo.mutation.Where(ps...)
	return eluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eluo *EntryLogsUpdateOne) Select(field string, fields ...string) *EntryLogsUpdateOne {
	eluo.fields = append([]string{field}, fields...)
	return eluo
}

// Save executes the query and returns the updated EntryLogs entity.
func (eluo *EntryLogsUpdateOne) Save(ctx context.Context) (*EntryLogs, error) {
	eluo.defaults()
	return withHooks(ctx, eluo.sqlSave, eluo.mutation, eluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eluo *EntryLogsUpdateOne) SaveX(ctx context.Context) *EntryLogs {
	node, err := eluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eluo *EntryLogsUpdateOne) Exec(ctx context.Context) error {
	_, err := eluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eluo *EntryLogsUpdateOne) ExecX(ctx context.Context) {
	if err := eluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eluo *EntryLogsUpdateOne) defaults() {
	if _, ok := eluo.mutation.UpdatedAt(); !ok && !eluo.mutation.UpdatedAtCleared() {
		v := entrylogs.UpdateDefaultUpdatedAt()
		eluo.mutation.SetUpdatedAt(v)
	}
}

func (eluo *EntryLogsUpdateOne) sqlSave(ctx context.Context) (_node *EntryLogs, err error) {
	_spec := sqlgraph.NewUpdateSpec(entrylogs.Table, entrylogs.Columns, sqlgraph.NewFieldSpec(entrylogs.FieldID, field.TypeInt64))
	id, ok := eluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EntryLogs.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entrylogs.FieldID)
		for _, f := range fields {
			if !entrylogs.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entrylogs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if eluo.mutation.CreatedAtCleared() {
		_spec.ClearField(entrylogs.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := eluo.mutation.UpdatedAt(); ok {
		_spec.SetField(entrylogs.FieldUpdatedAt, field.TypeTime, value)
	}
	if eluo.mutation.UpdatedAtCleared() {
		_spec.ClearField(entrylogs.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := eluo.mutation.Delete(); ok {
		_spec.SetField(entrylogs.FieldDelete, field.TypeInt64, value)
	}
	if value, ok := eluo.mutation.AddedDelete(); ok {
		_spec.AddField(entrylogs.FieldDelete, field.TypeInt64, value)
	}
	if eluo.mutation.DeleteCleared() {
		_spec.ClearField(entrylogs.FieldDelete, field.TypeInt64)
	}
	if value, ok := eluo.mutation.CreatedID(); ok {
		_spec.SetField(entrylogs.FieldCreatedID, field.TypeInt64, value)
	}
	if value, ok := eluo.mutation.AddedCreatedID(); ok {
		_spec.AddField(entrylogs.FieldCreatedID, field.TypeInt64, value)
	}
	if eluo.mutation.CreatedIDCleared() {
		_spec.ClearField(entrylogs.FieldCreatedID, field.TypeInt64)
	}
	if value, ok := eluo.mutation.MemberProductID(); ok {
		_spec.SetField(entrylogs.FieldMemberProductID, field.TypeInt64, value)
	}
	if value, ok := eluo.mutation.AddedMemberProductID(); ok {
		_spec.AddField(entrylogs.FieldMemberProductID, field.TypeInt64, value)
	}
	if eluo.mutation.MemberProductIDCleared() {
		_spec.ClearField(entrylogs.FieldMemberProductID, field.TypeInt64)
	}
	if value, ok := eluo.mutation.MemberPropertyID(); ok {
		_spec.SetField(entrylogs.FieldMemberPropertyID, field.TypeInt64, value)
	}
	if value, ok := eluo.mutation.AddedMemberPropertyID(); ok {
		_spec.AddField(entrylogs.FieldMemberPropertyID, field.TypeInt64, value)
	}
	if eluo.mutation.MemberPropertyIDCleared() {
		_spec.ClearField(entrylogs.FieldMemberPropertyID, field.TypeInt64)
	}
	if value, ok := eluo.mutation.EntryTime(); ok {
		_spec.SetField(entrylogs.FieldEntryTime, field.TypeTime, value)
	}
	if eluo.mutation.EntryTimeCleared() {
		_spec.ClearField(entrylogs.FieldEntryTime, field.TypeTime)
	}
	if value, ok := eluo.mutation.LeavingTime(); ok {
		_spec.SetField(entrylogs.FieldLeavingTime, field.TypeTime, value)
	}
	if eluo.mutation.LeavingTimeCleared() {
		_spec.ClearField(entrylogs.FieldLeavingTime, field.TypeTime)
	}
	if eluo.mutation.VenuesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.VenuesTable,
			Columns: []string{entrylogs.VenuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eluo.mutation.VenuesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.VenuesTable,
			Columns: []string{entrylogs.VenuesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(venue.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eluo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.MembersTable,
			Columns: []string{entrylogs.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eluo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.MembersTable,
			Columns: []string{entrylogs.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eluo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.UsersTable,
			Columns: []string{entrylogs.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eluo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   entrylogs.UsersTable,
			Columns: []string{entrylogs.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EntryLogs{config: eluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entrylogs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	eluo.mutation.done = true
	return _node, nil
}
