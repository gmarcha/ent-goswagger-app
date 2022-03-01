package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().Default(uuid.New),
		field.String("name").Unique(),
		field.String("event").Sensitive().Default("false"),
		field.String("event_write").Sensitive().Default("false"),
		field.String("user").Sensitive().Default("false"),
		field.String("user_subscription").Sensitive().Default("false"),
		field.String("user_write").Sensitive().Default("false"),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
