package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
			Immutable().Default(uuid.New),
		field.String("name"),
		field.Enum("category").
			Values("exam", "rush", "meeting", "events"),
		field.String("description").
			Optional(),
		field.Int64("tutorsRequired").
			NonNegative(),
		field.Int64("walletsReward").
			NonNegative().Optional(),
		field.Time("createdAt").
			Immutable().Default(time.Now),
		field.Time("startAt"),
		field.Time("endAt"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
