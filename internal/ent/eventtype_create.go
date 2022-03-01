// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/event"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/eventtype"
	"github.com/google/uuid"
)

// EventTypeCreate is the builder for creating a EventType entity.
type EventTypeCreate struct {
	config
	mutation *EventTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (etc *EventTypeCreate) SetName(s string) *EventTypeCreate {
	etc.mutation.SetName(s)
	return etc
}

// SetColor sets the "color" field.
func (etc *EventTypeCreate) SetColor(s string) *EventTypeCreate {
	etc.mutation.SetColor(s)
	return etc
}

// SetID sets the "id" field.
func (etc *EventTypeCreate) SetID(u uuid.UUID) *EventTypeCreate {
	etc.mutation.SetID(u)
	return etc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (etc *EventTypeCreate) SetNillableID(u *uuid.UUID) *EventTypeCreate {
	if u != nil {
		etc.SetID(*u)
	}
	return etc
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (etc *EventTypeCreate) AddEventIDs(ids ...uuid.UUID) *EventTypeCreate {
	etc.mutation.AddEventIDs(ids...)
	return etc
}

// AddEvents adds the "events" edges to the Event entity.
func (etc *EventTypeCreate) AddEvents(e ...*Event) *EventTypeCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return etc.AddEventIDs(ids...)
}

// Mutation returns the EventTypeMutation object of the builder.
func (etc *EventTypeCreate) Mutation() *EventTypeMutation {
	return etc.mutation
}

// Save creates the EventType in the database.
func (etc *EventTypeCreate) Save(ctx context.Context) (*EventType, error) {
	var (
		err  error
		node *EventType
	)
	etc.defaults()
	if len(etc.hooks) == 0 {
		if err = etc.check(); err != nil {
			return nil, err
		}
		node, err = etc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = etc.check(); err != nil {
				return nil, err
			}
			etc.mutation = mutation
			if node, err = etc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(etc.hooks) - 1; i >= 0; i-- {
			if etc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (etc *EventTypeCreate) SaveX(ctx context.Context) *EventType {
	v, err := etc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etc *EventTypeCreate) Exec(ctx context.Context) error {
	_, err := etc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etc *EventTypeCreate) ExecX(ctx context.Context) {
	if err := etc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (etc *EventTypeCreate) defaults() {
	if _, ok := etc.mutation.ID(); !ok {
		v := eventtype.DefaultID()
		etc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (etc *EventTypeCreate) check() error {
	if _, ok := etc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "EventType.name"`)}
	}
	if _, ok := etc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "EventType.color"`)}
	}
	return nil
}

func (etc *EventTypeCreate) sqlSave(ctx context.Context) (*EventType, error) {
	_node, _spec := etc.createSpec()
	if err := sqlgraph.CreateNode(ctx, etc.driver, _spec); err != nil {
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

func (etc *EventTypeCreate) createSpec() (*EventType, *sqlgraph.CreateSpec) {
	var (
		_node = &EventType{config: etc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: eventtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventtype.FieldID,
			},
		}
	)
	if id, ok := etc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := etc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventtype.FieldName,
		})
		_node.Name = value
	}
	if value, ok := etc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: eventtype.FieldColor,
		})
		_node.Color = value
	}
	if nodes := etc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   eventtype.EventsTable,
			Columns: []string{eventtype.EventsColumn},
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

// EventTypeCreateBulk is the builder for creating many EventType entities in bulk.
type EventTypeCreateBulk struct {
	config
	builders []*EventTypeCreate
}

// Save creates the EventType entities in the database.
func (etcb *EventTypeCreateBulk) Save(ctx context.Context) ([]*EventType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(etcb.builders))
	nodes := make([]*EventType, len(etcb.builders))
	mutators := make([]Mutator, len(etcb.builders))
	for i := range etcb.builders {
		func(i int, root context.Context) {
			builder := etcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EventTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, etcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, etcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, etcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (etcb *EventTypeCreateBulk) SaveX(ctx context.Context) []*EventType {
	v, err := etcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (etcb *EventTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := etcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (etcb *EventTypeCreateBulk) ExecX(ctx context.Context) {
	if err := etcb.Exec(ctx); err != nil {
		panic(err)
	}
}
