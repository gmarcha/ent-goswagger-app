package schema

import "entgo.io/ent"

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return nil
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
