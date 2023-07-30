package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").MaxLen(50),
		field.String("domain").MaxLen(50).Optional(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
