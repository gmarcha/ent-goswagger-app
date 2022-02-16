// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gamarcha/ent-goswagger-app/internal/ent/event"
	"github.com/gamarcha/ent-goswagger-app/internal/ent/user"
	"github.com/google/uuid"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetLogin sets the "login" field.
func (uc *UserCreate) SetLogin(s string) *UserCreate {
	uc.mutation.SetLogin(s)
	return uc
}

// SetFirstName sets the "firstName" field.
func (uc *UserCreate) SetFirstName(s string) *UserCreate {
	uc.mutation.SetFirstName(s)
	return uc
}

// SetNillableFirstName sets the "firstName" field if the given value is not nil.
func (uc *UserCreate) SetNillableFirstName(s *string) *UserCreate {
	if s != nil {
		uc.SetFirstName(*s)
	}
	return uc
}

// SetLastName sets the "lastName" field.
func (uc *UserCreate) SetLastName(s string) *UserCreate {
	uc.mutation.SetLastName(s)
	return uc
}

// SetNillableLastName sets the "lastName" field if the given value is not nil.
func (uc *UserCreate) SetNillableLastName(s *string) *UserCreate {
	if s != nil {
		uc.SetLastName(*s)
	}
	return uc
}

// SetHoursDone sets the "hoursDone" field.
func (uc *UserCreate) SetHoursDone(i int64) *UserCreate {
	uc.mutation.SetHoursDone(i)
	return uc
}

// SetTutorScope sets the "tutorScope" field.
func (uc *UserCreate) SetTutorScope(b bool) *UserCreate {
	uc.mutation.SetTutorScope(b)
	return uc
}

// SetNillableTutorScope sets the "tutorScope" field if the given value is not nil.
func (uc *UserCreate) SetNillableTutorScope(b *bool) *UserCreate {
	if b != nil {
		uc.SetTutorScope(*b)
	}
	return uc
}

// SetAdminScope sets the "adminScope" field.
func (uc *UserCreate) SetAdminScope(b bool) *UserCreate {
	uc.mutation.SetAdminScope(b)
	return uc
}

// SetNillableAdminScope sets the "adminScope" field if the given value is not nil.
func (uc *UserCreate) SetNillableAdminScope(b *bool) *UserCreate {
	if b != nil {
		uc.SetAdminScope(*b)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(u uuid.UUID) *UserCreate {
	uc.mutation.SetID(u)
	return uc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (uc *UserCreate) SetNillableID(u *uuid.UUID) *UserCreate {
	if u != nil {
		uc.SetID(*u)
	}
	return uc
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (uc *UserCreate) AddEventIDs(ids ...uuid.UUID) *UserCreate {
	uc.mutation.AddEventIDs(ids...)
	return uc
}

// AddEvents adds the "events" edges to the Event entity.
func (uc *UserCreate) AddEvents(e ...*Event) *UserCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return uc.AddEventIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() {
	if _, ok := uc.mutation.TutorScope(); !ok {
		v := user.DefaultTutorScope
		uc.mutation.SetTutorScope(v)
	}
	if _, ok := uc.mutation.AdminScope(); !ok {
		v := user.DefaultAdminScope
		uc.mutation.SetAdminScope(v)
	}
	if _, ok := uc.mutation.ID(); !ok {
		v := user.DefaultID()
		uc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.Login(); !ok {
		return &ValidationError{Name: "login", err: errors.New(`ent: missing required field "User.login"`)}
	}
	if _, ok := uc.mutation.HoursDone(); !ok {
		return &ValidationError{Name: "hoursDone", err: errors.New(`ent: missing required field "User.hoursDone"`)}
	}
	if v, ok := uc.mutation.HoursDone(); ok {
		if err := user.HoursDoneValidator(v); err != nil {
			return &ValidationError{Name: "hoursDone", err: fmt.Errorf(`ent: validator failed for field "User.hoursDone": %w`, err)}
		}
	}
	if _, ok := uc.mutation.TutorScope(); !ok {
		return &ValidationError{Name: "tutorScope", err: errors.New(`ent: missing required field "User.tutorScope"`)}
	}
	if _, ok := uc.mutation.AdminScope(); !ok {
		return &ValidationError{Name: "adminScope", err: errors.New(`ent: missing required field "User.adminScope"`)}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: user.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := uc.mutation.Login(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLogin,
		})
		_node.Login = value
	}
	if value, ok := uc.mutation.FirstName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldFirstName,
		})
		_node.FirstName = value
	}
	if value, ok := uc.mutation.LastName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldLastName,
		})
		_node.LastName = value
	}
	if value, ok := uc.mutation.HoursDone(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldHoursDone,
		})
		_node.HoursDone = value
	}
	if value, ok := uc.mutation.TutorScope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldTutorScope,
		})
		_node.TutorScope = value
	}
	if value, ok := uc.mutation.AdminScope(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: user.FieldAdminScope,
		})
		_node.AdminScope = value
	}
	if nodes := uc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.EventsTable,
			Columns: user.EventsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: event.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
