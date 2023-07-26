package main

import (
	"context"
	"pos/pkg"
	"pos/pkg/config"
)

func main() {
	if err := pkg.LoadENV(".env"); err != nil {
		panic(err)
	}

	config.InitAuthConfig()

	pkg.InitEnt()

	createUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create User").SetValue("create_user").SaveX(context.Background())
	createAdminUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create Admin User").SetValue("create_admin_user").SaveX(context.Background())

	saRole := pkg.EntClient().Role.Create().SetTitle("Superadmin").SetValue("SUPERADMIN").SaveX(context.Background())
	saRole.Update().AddPermissions(createUserPermission, createAdminUserPermission).ExecX(context.Background())

	superadmin := pkg.EntClient().User.Create().SetFirstName("super").SetLastName("admin").
		SetUsername("superadmin").SetPassword(pkg.Hash("password")).SetRoleID(saRole.ID).
		SaveX(context.Background())

	_ = superadmin
}
