package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EventType holds the schema definition for the EventType entity.
type EventType struct {
	ent.Schema
}

// Fields of the EventType.
func (EventType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New),
		field.String("name").
			Unique(),
		field.String("color"),
	}
}

// Edges of the EventType.
func (EventType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type),
	}
}
