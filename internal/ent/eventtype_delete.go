// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/eventtype"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/predicate"
)

// EventTypeDelete is the builder for deleting a EventType entity.
type EventTypeDelete struct {
	config
	hooks    []Hook
	mutation *EventTypeMutation
}

// Where appends a list predicates to the EventTypeDelete builder.
func (etd *EventTypeDelete) Where(ps ...predicate.EventType) *EventTypeDelete {
	etd.mutation.Where(ps...)
	return etd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (etd *EventTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(etd.hooks) == 0 {
		affected, err = etd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EventTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			etd.mutation = mutation
			affected, err = etd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(etd.hooks) - 1; i >= 0; i-- {
			if etd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = etd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, etd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (etd *EventTypeDelete) ExecX(ctx context.Context) int {
	n, err := etd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (etd *EventTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: eventtype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: eventtype.FieldID,
			},
		},
	}
	if ps := etd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, etd.driver, _spec)
}

// EventTypeDeleteOne is the builder for deleting a single EventType entity.
type EventTypeDeleteOne struct {
	etd *EventTypeDelete
}

// Exec executes the deletion query.
func (etdo *EventTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := etdo.etd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{eventtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (etdo *EventTypeDeleteOne) ExecX(ctx context.Context) {
	etdo.etd.ExecX(ctx)
}