package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").MaxLen(50),
		field.String("last_name").MaxLen(50),
		field.String("username").MaxLen(50).Unique(),
		field.String("password").Sensitive(),
		field.String("phone_number").Optional(),
		field.String("email").Optional(),
		// field.Enum("role_id").Values("SUPERADMIN", "ADMIN", "CASHIER", "CUSTOMER"),
		field.Int("role_id"),
		field.Enum("has_permission").Values("NULL", "ROLE", "USER").Default("NULL"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type),
		// edge.To("roles", Role.Type).Field("role_id"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
