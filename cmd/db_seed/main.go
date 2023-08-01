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

	indexCompanyPermission := pkg.EntClient().Permission.Create().SetTitle("Index Company").SetValue("index_company").SaveX(context.Background())
	createCompanyPermission := pkg.EntClient().Permission.Create().SetTitle("Create Company").SetValue("create_company").SaveX(context.Background())
	updateCompanyPermission := pkg.EntClient().Permission.Create().SetTitle("Update Company").SetValue("update_company").SaveX(context.Background())
	deleteCompanyPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Company").SetValue("delete_company").SaveX(context.Background())

	createUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create User").SetValue("create_user").SaveX(context.Background())
	createAdminUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create Admin User").SetValue("create_admin_user").SaveX(context.Background())
	createCashierUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create Cashier User").SetValue("create_cashier_user").SaveX(context.Background())
	createCustomerUserPermission := pkg.EntClient().Permission.Create().SetTitle("Create Customer User").SetValue("create_customer_user").SaveX(context.Background())
	deleteUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete User").SetValue("delete_user").SaveX(context.Background())
	deleteAdminUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Admin User").SetValue("delete_admin_user").SaveX(context.Background())
	deleteCashierUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Cashier User").SetValue("delete_cashier_user").SaveX(context.Background())
	deleteCustomerUserPermission := pkg.EntClient().Permission.Create().SetTitle("Delete Customer User").SetValue("delete_customer_user").SaveX(context.Background())

	superAdminCompany := pkg.EntClient().Company.Create().SetName("Elitbuzz Technologies Ltd.").SetDomain("elibuzz.com.bd").SaveX(context.Background())
	adminCompany := pkg.EntClient().Company.Create().SetName("Admin Technologies Ltd.").SetDomain("admin.com.bd").SaveX(context.Background())

	saRole := pkg.EntClient().Role.Create().SetTitle("Superadmin").SetValue("SUPERADMIN").SaveX(context.Background())
	saRole.Update().AddPermissions(indexCompanyPermission, createCompanyPermission, updateCompanyPermission, deleteCompanyPermission, createUserPermission, createAdminUserPermission, deleteUserPermission, deleteAdminUserPermission).ExecX(context.Background())

	superadmin := pkg.EntClient().User.Create().SetFirstName("super").SetLastName("admin").
		SetUsername("superadmin").SetPassword(pkg.Hash("password")).SetRoleID(saRole.ID).SetCompanyID(superAdminCompany.ID).
		SaveX(context.Background())

	_ = superadmin

	adminRole := pkg.EntClient().Role.Create().SetTitle("admin").SetValue("ADMIN").SaveX(context.Background())
	adminRole.Update().AddPermissions(createUserPermission, createCashierUserPermission, createCustomerUserPermission, deleteUserPermission, deleteCashierUserPermission, deleteCustomerUserPermission).ExecX(context.Background())

	admin := pkg.EntClient().User.Create().SetFirstName("admin").SetLastName("admin").
		SetUsername("admin").SetPassword(pkg.Hash("password")).SetRoleID(adminRole.ID).SetCompanyID(adminCompany.ID).
		SaveX(context.Background())

	_ = admin

	cashierRole := pkg.EntClient().Role.Create().SetTitle("cashier").SetValue("CASHIER").SaveX(context.Background())
	cashierRole.Update().AddPermissions(createUserPermission, createCustomerUserPermission, deleteUserPermission, deleteCustomerUserPermission).ExecX(context.Background())

	cashier := pkg.EntClient().User.Create().SetFirstName("cashier").SetLastName("cashier").
		SetUsername("cashier").SetPassword(pkg.Hash("password")).SetRoleID(cashierRole.ID).SetCompanyID(adminCompany.ID).
		SaveX(context.Background())

	_ = cashier

	customerRole := pkg.EntClient().Role.Create().SetTitle("customer").SetValue("CUSTOMER").SaveX(context.Background())

	customer := pkg.EntClient().User.Create().SetFirstName("customer").SetLastName("customer").
		SetUsername("customer").SetPassword(pkg.Hash("password")).SetRoleID(customerRole.ID).SetCompanyID(admin.ID).
		SaveX(context.Background())

	_ = customer

}
