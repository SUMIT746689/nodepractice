// Code generated by ent, DO NOT EDIT.

package ent

import (
	"pos/ent/permission"
	"pos/ent/role"
	"pos/ent/schema"
	"pos/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	permissionFields := schema.Permission{}.Fields()
	_ = permissionFields
	// permissionDescTitle is the schema descriptor for title field.
	permissionDescTitle := permissionFields[0].Descriptor()
	// permission.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	permission.TitleValidator = permissionDescTitle.Validators[0].(func(string) error)
	// permissionDescValue is the schema descriptor for value field.
	permissionDescValue := permissionFields[1].Descriptor()
	// permission.ValueValidator is a validator for the "value" field. It is called by the builders before save.
	permission.ValueValidator = permissionDescValue.Validators[0].(func(string) error)
	// permissionDescGroup is the schema descriptor for group field.
	permissionDescGroup := permissionFields[2].Descriptor()
	// permission.GroupValidator is a validator for the "group" field. It is called by the builders before save.
	permission.GroupValidator = permissionDescGroup.Validators[0].(func(string) error)
	roleFields := schema.Role{}.Fields()
	_ = roleFields
	// roleDescTitle is the schema descriptor for title field.
	roleDescTitle := roleFields[0].Descriptor()
	// role.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	role.TitleValidator = roleDescTitle.Validators[0].(func(string) error)
	userMixin := schema.User{}.Mixin()
	userMixinFields0 := userMixin[0].Fields()
	_ = userMixinFields0
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreateTime is the schema descriptor for create_time field.
	userDescCreateTime := userMixinFields0[0].Descriptor()
	// user.DefaultCreateTime holds the default value on creation for the create_time field.
	user.DefaultCreateTime = userDescCreateTime.Default.(func() time.Time)
	// userDescUpdateTime is the schema descriptor for update_time field.
	userDescUpdateTime := userMixinFields0[1].Descriptor()
	// user.DefaultUpdateTime holds the default value on creation for the update_time field.
	user.DefaultUpdateTime = userDescUpdateTime.Default.(func() time.Time)
	// user.UpdateDefaultUpdateTime holds the default value on update for the update_time field.
	user.UpdateDefaultUpdateTime = userDescUpdateTime.UpdateDefault.(func() time.Time)
	// userDescFirstName is the schema descriptor for first_name field.
	userDescFirstName := userFields[0].Descriptor()
	// user.FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	user.FirstNameValidator = userDescFirstName.Validators[0].(func(string) error)
	// userDescLastName is the schema descriptor for last_name field.
	userDescLastName := userFields[1].Descriptor()
	// user.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	user.LastNameValidator = userDescLastName.Validators[0].(func(string) error)
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[2].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
}
