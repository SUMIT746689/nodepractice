package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(50),
		field.String("value").MaxLen(50),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type),
		edge.To("users", User.Type),
	}
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
