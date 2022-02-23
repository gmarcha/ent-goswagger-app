package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.String("login").
			Unique(),
		field.String("firstName").
			Optional(),
		field.String("lastName").
			Optional(),
		field.String("imagePath"),
		field.Bool("calendarScope").
			Default(false),
		field.Bool("adminScope").
			Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("events", Event.Type).
			Ref("users"),
	}
}
