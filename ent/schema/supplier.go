package schema

import (
	"pos/internal/domain"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Supplier holds the schema definition for the Supplier entity.
type Supplier struct {
	ent.Schema
}

// Fields of the Supplier.
func (Supplier) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("address"),
		field.String("email").Optional(),
		field.JSON("representative", domain.Representative{}),
	}
}

// Edges of the Supplier.
func (Supplier) Edges() []ent.Edge {
	return nil
}
