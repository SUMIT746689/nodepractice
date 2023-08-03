package schema

import (
	"pos/internal/domain"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Vendor holds the schema definition for the Vendor entity.
type Vendor struct {
	ent.Schema
}

// Fields of the Vendor.
func (Vendor) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("address"),
		field.String("email").Optional(),
		field.JSON("representative", domain.Representative{}),
	}
}

// Edges of the Vendor.
func (Vendor) Edges() []ent.Edge {
	return nil
}
