// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/event"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/eventtype"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/predicate"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/user"
	"github.com/google/uuid"
)

// EventUpdate is the builder for updating Event entities.
type EventUpdate struct {
	config
	hooks    []Hook
	mutation *EventMutation
}

// Where appends a list predicates to the EventUpdate builder.
func (eu *EventUpdate) Where(ps ...predicate.Event) *EventUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetName sets the "name" field.
func (eu *EventUpdate) SetName(s string) *EventUpdate {
	eu.mutation.SetName(s)
	return eu
}

// SetDescription sets the "description" field.
func (eu *EventUpdate) SetDescription(s string) *EventUpdate {
	eu.mutation.SetDescription(s)
	return eu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (eu *EventUpdate) SetNillableDescription(s *string) *EventUpdate {
	if s != nil {
		eu.SetDescription(*s)
	}
	return eu
}

// ClearDescription clears the value of the "description" field.
func (eu *EventUpdate) ClearDescription() *EventUpdate {
	eu.mutation.ClearDescription()
	return eu
}

// SetTutorsRequired sets the "tutorsRequired" field.
func (eu *EventUpdate) SetTutorsRequired(i int64) *EventUpdate {
	eu.mutation.ResetTutorsRequired()
	eu.mutation.SetTutorsRequired(i)
	return eu
}

// SetNillableTutorsRequired sets the "tutorsRequired" field if the given value is not nil.
func (eu *EventUpdate) SetNillableTutorsRequired(i *int64) *EventUpdate {
	if i != nil {
		eu.SetTutorsRequired(*i)
	}
	return eu
}

// AddTutorsRequired adds i to the "tutorsRequired" field.
func (eu *EventUpdate) AddTutorsRequired(i int64) *EventUpdate {
	eu.mutation.AddTutorsRequired(i)
	return eu
}

// ClearTutorsRequired clears the value of the "tutorsRequired" field.
func (eu *EventUpdate) ClearTutorsRequired() *EventUpdate {
	eu.mutation.ClearTutorsRequired()
	return eu
}

// SetWalletsReward sets the "walletsReward" field.
func (eu *EventUpdate) SetWalletsReward(i int64) *EventUpdate {
	eu.mutation.ResetWalletsReward()
	eu.mutation.SetWalletsReward(i)
	return eu
}

// SetNillableWalletsReward sets the "walletsReward" field if the given value is not nil.
func (eu *EventUpdate) SetNillableWalletsReward(i *int64) *EventUpdate {
	if i != nil {
		eu.SetWalletsReward(*i)
	}
	return eu
}

// AddWalletsReward adds i to the "walletsReward" field.
func (eu *EventUpdate) AddWalletsReward(i int64) *EventUpdate {
	eu.mutation.AddWalletsReward(i)
	return eu
}

// ClearWalletsReward clears the value of the "walletsReward" field.
func (eu *EventUpdate) ClearWalletsReward() *EventUpdate {
	eu.mutation.ClearWalletsReward()
	return eu
}

// SetStartAt sets the "startAt" field.
func (eu *EventUpdate) SetStartAt(t time.Time) *EventUpdate {
	eu.mutation.SetStartAt(t)
	return eu
}

// SetEndAt sets the "endAt" field.
func (eu *EventUpdate) SetEndAt(t time.Time) *EventUpdate {
	eu.mutation.SetEndAt(t)
	return eu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (eu *EventUpdate) AddUserIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.AddUserIDs(ids...)
	return eu
}

// AddUsers adds the "users" edges to the User entity.
func (eu *EventUpdate) AddUsers(u ...*User) *EventUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.AddUserIDs(ids...)
}

// SetCategoryID sets the "category" edge to the EventType entity by ID.
func (eu *EventUpdate) SetCategoryID(id uuid.UUID) *EventUpdate {
	eu.mutation.SetCategoryID(id)
	return eu
}

// SetNillableCategoryID sets the "category" edge to the EventType entity by ID if the given value is not nil.
func (eu *EventUpdate) SetNillableCategoryID(id *uuid.UUID) *EventUpdate {
	if id != nil {
		eu = eu.SetCategoryID(*id)
	}
	return eu
}

// SetCategory sets the "category" edge to the EventType entity.
func (eu *EventUpdate) SetCategory(e *EventType) *EventUpdate {
	return eu.SetCategoryID(e.ID)
}

// Mutation returns the EventMutation object of the builder.
func (eu *EventUpdate) Mutation() *EventMutation {
	return eu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (eu *EventUpdate) ClearUsers() *EventUpdate {
	eu.mutation.ClearUsers()
	return eu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (eu *EventUpdate) RemoveUserIDs(ids ...uuid.UUID) *EventUpdate {
	eu.mutation.RemoveUserIDs(ids...)
	return eu
}

// RemoveUsers removes "users" edges to User entities.
func (eu *EventUpdate) RemoveUsers(u ...*User) *EventUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return eu.RemoveUserIDs(ids...)
}

// ClearCategory clears the "category" edge to the EventType entity.
func (eu *EventUpdate) ClearCategory() *EventUpdate {
	eu.mutation.ClearCategory()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EventUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EventUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EventUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EventUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EventUpdate) check() error {
	if v, ok := eu.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if v, ok := eu.mutation.TutorsRequired(); ok {
		if err := event.TutorsRequiredValidator(v); err != nil {
			return &ValidationError{Name: "tutorsRequired", err: fmt.Errorf(`ent: validator failed for field "Event.tutorsRequired": %w`, err)}
		}
	}
	if v, ok := eu.mutation.WalletsReward(); ok {
		if err := event.WalletsRewardValidator(v); err != nil {
			return &ValidationError{Name: "walletsReward", err: fmt.Errorf(`ent: validator failed for field "Event.walletsReward": %w`, err)}
		}
	}
	return nil
}

func (eu *EventUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := eu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
	}
	if eu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: event.FieldDescription,
		})
	}
	if value, ok := eu.mutation.TutorsRequired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTutorsRequired,
		})
	}
	if value, ok := eu.mutation.AddedTutorsRequired(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTutorsRequired,
		})
	}
	if eu.mutation.TutorsRequiredCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: event.FieldTutorsRequired,
		})
	}
	if value, ok := eu.mutation.WalletsReward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldWalletsReward,
		})
	}
	if value, ok := eu.mutation.AddedWalletsReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldWalletsReward,
		})
	}
	if eu.mutation.WalletsRewardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: event.FieldWalletsReward,
		})
	}
	if value, ok := eu.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStartAt,
		})
	}
	if value, ok := eu.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEndAt,
		})
	}
	if eu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !eu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.CategoryTable,
			Columns: []string{event.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.CategoryTable,
			Columns: []string{event.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EventUpdateOne is the builder for updating a single Event entity.
type EventUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EventMutation
}

// SetName sets the "name" field.
func (euo *EventUpdateOne) SetName(s string) *EventUpdateOne {
	euo.mutation.SetName(s)
	return euo
}

// SetDescription sets the "description" field.
func (euo *EventUpdateOne) SetDescription(s string) *EventUpdateOne {
	euo.mutation.SetDescription(s)
	return euo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableDescription(s *string) *EventUpdateOne {
	if s != nil {
		euo.SetDescription(*s)
	}
	return euo
}

// ClearDescription clears the value of the "description" field.
func (euo *EventUpdateOne) ClearDescription() *EventUpdateOne {
	euo.mutation.ClearDescription()
	return euo
}

// SetTutorsRequired sets the "tutorsRequired" field.
func (euo *EventUpdateOne) SetTutorsRequired(i int64) *EventUpdateOne {
	euo.mutation.ResetTutorsRequired()
	euo.mutation.SetTutorsRequired(i)
	return euo
}

// SetNillableTutorsRequired sets the "tutorsRequired" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableTutorsRequired(i *int64) *EventUpdateOne {
	if i != nil {
		euo.SetTutorsRequired(*i)
	}
	return euo
}

// AddTutorsRequired adds i to the "tutorsRequired" field.
func (euo *EventUpdateOne) AddTutorsRequired(i int64) *EventUpdateOne {
	euo.mutation.AddTutorsRequired(i)
	return euo
}

// ClearTutorsRequired clears the value of the "tutorsRequired" field.
func (euo *EventUpdateOne) ClearTutorsRequired() *EventUpdateOne {
	euo.mutation.ClearTutorsRequired()
	return euo
}

// SetWalletsReward sets the "walletsReward" field.
func (euo *EventUpdateOne) SetWalletsReward(i int64) *EventUpdateOne {
	euo.mutation.ResetWalletsReward()
	euo.mutation.SetWalletsReward(i)
	return euo
}

// SetNillableWalletsReward sets the "walletsReward" field if the given value is not nil.
func (euo *EventUpdateOne) SetNillableWalletsReward(i *int64) *EventUpdateOne {
	if i != nil {
		euo.SetWalletsReward(*i)
	}
	return euo
}

// AddWalletsReward adds i to the "walletsReward" field.
func (euo *EventUpdateOne) AddWalletsReward(i int64) *EventUpdateOne {
	euo.mutation.AddWalletsReward(i)
	return euo
}

// ClearWalletsReward clears the value of the "walletsReward" field.
func (euo *EventUpdateOne) ClearWalletsReward() *EventUpdateOne {
	euo.mutation.ClearWalletsReward()
	return euo
}

// SetStartAt sets the "startAt" field.
func (euo *EventUpdateOne) SetStartAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetStartAt(t)
	return euo
}

// SetEndAt sets the "endAt" field.
func (euo *EventUpdateOne) SetEndAt(t time.Time) *EventUpdateOne {
	euo.mutation.SetEndAt(t)
	return euo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (euo *EventUpdateOne) AddUserIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.AddUserIDs(ids...)
	return euo
}

// AddUsers adds the "users" edges to the User entity.
func (euo *EventUpdateOne) AddUsers(u ...*User) *EventUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.AddUserIDs(ids...)
}

// SetCategoryID sets the "category" edge to the EventType entity by ID.
func (euo *EventUpdateOne) SetCategoryID(id uuid.UUID) *EventUpdateOne {
	euo.mutation.SetCategoryID(id)
	return euo
}

// SetNillableCategoryID sets the "category" edge to the EventType entity by ID if the given value is not nil.
func (euo *EventUpdateOne) SetNillableCategoryID(id *uuid.UUID) *EventUpdateOne {
	if id != nil {
		euo = euo.SetCategoryID(*id)
	}
	return euo
}

// SetCategory sets the "category" edge to the EventType entity.
func (euo *EventUpdateOne) SetCategory(e *EventType) *EventUpdateOne {
	return euo.SetCategoryID(e.ID)
}

// Mutation returns the EventMutation object of the builder.
func (euo *EventUpdateOne) Mutation() *EventMutation {
	return euo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (euo *EventUpdateOne) ClearUsers() *EventUpdateOne {
	euo.mutation.ClearUsers()
	return euo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (euo *EventUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *EventUpdateOne {
	euo.mutation.RemoveUserIDs(ids...)
	return euo
}

// RemoveUsers removes "users" edges to User entities.
func (euo *EventUpdateOne) RemoveUsers(u ...*User) *EventUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return euo.RemoveUserIDs(ids...)
}

// ClearCategory clears the "category" edge to the EventType entity.
func (euo *EventUpdateOne) ClearCategory() *EventUpdateOne {
	euo.mutation.ClearCategory()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EventUpdateOne) Select(field string, fields ...string) *EventUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Event entity.
func (euo *EventUpdateOne) Save(ctx context.Context) (*Event, error) {
	var (
		err  error
		node *Event
	)
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EventUpdateOne) SaveX(ctx context.Context) *Event {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EventUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EventUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EventUpdateOne) check() error {
	if v, ok := euo.mutation.Name(); ok {
		if err := event.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Event.name": %w`, err)}
		}
	}
	if v, ok := euo.mutation.TutorsRequired(); ok {
		if err := event.TutorsRequiredValidator(v); err != nil {
			return &ValidationError{Name: "tutorsRequired", err: fmt.Errorf(`ent: validator failed for field "Event.tutorsRequired": %w`, err)}
		}
	}
	if v, ok := euo.mutation.WalletsReward(); ok {
		if err := event.WalletsRewardValidator(v); err != nil {
			return &ValidationError{Name: "walletsReward", err: fmt.Errorf(`ent: validator failed for field "Event.walletsReward": %w`, err)}
		}
	}
	return nil
}

func (euo *EventUpdateOne) sqlSave(ctx context.Context) (_node *Event, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   event.Table,
			Columns: event.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: event.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Event.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, event.FieldID)
		for _, f := range fields {
			if !event.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != event.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldName,
		})
	}
	if value, ok := euo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: event.FieldDescription,
		})
	}
	if euo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: event.FieldDescription,
		})
	}
	if value, ok := euo.mutation.TutorsRequired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTutorsRequired,
		})
	}
	if value, ok := euo.mutation.AddedTutorsRequired(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldTutorsRequired,
		})
	}
	if euo.mutation.TutorsRequiredCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: event.FieldTutorsRequired,
		})
	}
	if value, ok := euo.mutation.WalletsReward(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldWalletsReward,
		})
	}
	if value, ok := euo.mutation.AddedWalletsReward(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: event.FieldWalletsReward,
		})
	}
	if euo.mutation.WalletsRewardCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Column: event.FieldWalletsReward,
		})
	}
	if value, ok := euo.mutation.StartAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldStartAt,
		})
	}
	if value, ok := euo.mutation.EndAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: event.FieldEndAt,
		})
	}
	if euo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !euo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   event.UsersTable,
			Columns: event.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.CategoryTable,
			Columns: []string{event.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventtype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   event.CategoryTable,
			Columns: []string{event.CategoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: eventtype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Event{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{event.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
