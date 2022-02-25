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
			Match(regexp.MustCompile("^[a-z]+$")).
			Unique(),
		field.String("firstName").
			Match(regexp.MustCompile("^[A-Z][a-z]*$")).
			Optional(),
		field.String("lastName").
			Match(regexp.MustCompile("^[A-Z][a-z]*$")).
			Optional(),
		field.String("imagePath").
			// Todo: apply regexp (depending on if it is a path or an uri)
			Optional(),
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
