package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Bool("event").Default(false),
		field.Bool("event_write").Default(false),
		field.Bool("user").Default(false),
		field.Bool("user_subscription").Default(false),
		field.Bool("user_write").Default(false),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
