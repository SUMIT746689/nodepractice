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
	createCashierUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create Cashier User").SetValue("create_cashier_user").SaveX(context.Background())
	createCustomerUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create Customer User").SetValue("create_customer_user").SaveX(context.Background())
	deleteUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete User").SetValue("delete_user").SaveX(context.Background())
	deleteAdminUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Admin User").SetValue("delete_admin_user").SaveX(context.Background())
	deleteCashierUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Cashier User").SetValue("delete_cashier_user").SaveX(context.Background())
	deleteCustomerUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Customer User").SetValue("delete_customer_user").SaveX(context.Background())

	// superAdminCompany := pkg.EntClient().Company.Create().SetName("Elitbuzz Technologies Ltd.")

	saRole := pkg.EntClient().Role.Create().SetTitle("Superadmin").SetValue("SUPERADMIN").SaveX(context.Background())
	saRole.Update().AddPermissions(createUserPermission, createAdminUserPermission, deleteUserPermission, deleteAdminUserPermission).ExecX(context.Background())
	// saRole.Update().AddPermissions(createUserPermission, createAdminUserPermission, deleteUserPermission, deleteAdminUserPermission).AddCompanys(superAdminCompany).ExecX(context.Background())

	superadmin := pkg.EntClient().User.Create().SetFirstName("super").SetLastName("admin").
		SetUsername("superadmin").SetPassword(pkg.Hash("password")).SetRoleID(saRole.ID).
		SaveX(context.Background())

	_ = superadmin

	adminRole := pkg.EntClient().Role.Create().SetTitle("admin").SetValue("ADMIN").SaveX(context.Background())
	adminRole.Update().AddPermissions(createUserPermission, createCashierUserPermission, createCustomerUserPermission, deleteUserPermission, deleteCashierUserPermission, deleteCustomerUserPermission).ExecX(context.Background())

	admin := pkg.EntClient().User.Create().SetFirstName("admin").SetLastName("admin").
		SetUsername("admin").SetPassword(pkg.Hash("password")).SetRoleID(adminRole.ID).
		SaveX(context.Background())

	_ = admin

	cashierRole := pkg.EntClient().Role.Create().SetTitle("cashier").SetValue("CASHIER").SaveX(context.Background())
	cashierRole.Update().AddPermissions(createUserPermission, createCashierUserPermission, createCustomerUserPermission, deleteUserPermission, deleteCashierUserPermission, deleteCustomerUserPermission).ExecX(context.Background())

	cashier := pkg.EntClient().User.Create().SetFirstName("cashier").SetLastName("cashier").
		SetUsername("cashier").SetPassword(pkg.Hash("password")).SetRoleID(cashierRole.ID).
		SaveX(context.Background())

	_ = cashier

	customerRole := pkg.EntClient().Role.Create().SetTitle("customer").SetValue("CUSTOMER").SaveX(context.Background())

	customer := pkg.EntClient().User.Create().SetFirstName("customer").SetLastName("customer").
		SetUsername("customer").SetPassword(pkg.Hash("password")).SetRoleID(customerRole.ID).
		SaveX(context.Background())

	_ = customer

}
