package schema

import (
	"entgo.io/ent"
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
		field.Enum("role").Values("SUPERADMIN", "ADMIN", "CASHIER", "CUSTOMER"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
