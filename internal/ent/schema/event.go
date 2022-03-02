package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/gmarcha/ent-goswagger-app/internal/ent"
	"github.com/gmarcha/ent-goswagger-app/internal/ent/hook"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("name").
			MinLen(1),
		field.String("description").
			Optional(),
		field.Int64("tutorsRequired").
			NonNegative().
			Optional().
			Nillable(),
		field.Int64("walletsReward").
			NonNegative().
			Optional().
			Nillable(),
		field.Time("createdAt").
			Immutable().
			Default(time.Now),
		field.Time("startAt"),
		field.Time("endAt"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.From("category", EventType.Type).
			Unique().
			Ref("events"),
	}
}

func (Event) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.EventFunc(func(ctx context.Context, m *gen.EventMutation) (ent.Value, error) {
					startAt, start := m.StartAt()
					endAt, end := m.EndAt()
					if start && end && startAt.Before(endAt) {
						return nil, fmt.Errorf("invalid date")
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
