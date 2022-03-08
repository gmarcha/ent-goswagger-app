package schema

import (
	"regexp"

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
			Immutable().
			Default(uuid.New),
		field.String("login").
			Immutable().
			Match(regexp.MustCompile("^[a-z]+$")).
			Unique(),
		field.String("firstName").
			Nillable().
			Optional(),
		field.String("lastName").
			Nillable().
			Optional(),
		field.String("displayName").
			Nillable().
			Optional(),
		field.String("imagePath").
			Nillable().
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).
			Ref("users"),
		edge.From("events", Event.Type).
			Ref("users"),
	}
}
