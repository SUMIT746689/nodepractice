package usercontroller

import (
	"context"
	"pos/ent/role"
	"pos/ent/user"
	"pos/internal/domain"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func createUserGuard(c *fiber.Ctx, roleID int, companyID int) (bool, int) {
	userID, _ := pkg.GetAuthedUser(c)

	// superadmin and body.companyID

	allowed := pkg.UserHasPermission(userID, domain.CREATE_USER)
	if !allowed {
		return false, 0
	}

	isAllowed := verifyCreateUserRolePermission(roleID, userID)

	if !isAllowed {
		return isAllowed, 0
	}

	user_, _ := pkg.EntClient().User.Query().Where(user.ID(userID)).WithRole().WithCompany().First(c.Context())

	if user_.Edges.Role.Title != "SUPERADMIN" {
		return isAllowed, user_.Edges.Company.ID
	}
	return isAllowed, 0
	// allowed = pkg.UserHasPermission(userID, domain.CREATE_ADMIN_USER)
	// if !allowed {
	// 	return false
	// }

	// return true

}

func updateUserGuard(c *fiber.Ctx, updateUserID int) (bool, bool) {
	userID, _ := pkg.GetAuthedUser(c)

	// allowed := pkg.UserHasPermission(userID, domain.DELETE_USER)
	// if !allowed {
	// 	return false
	// }
	canUpdateCompany := false
	user_, _ := pkg.EntClient().User.Query().Where(user.ID(userID)).WithRole().First(c.Context())
	if user_.Edges.Role.Title == "SUPERADMIN" {
		canUpdateCompany = true
	}

	role, err := pkg.EntClient().User.Query().Where(user.ID(updateUserID)).QueryRole().First(c.Context())
	if err != nil {
		return false, canUpdateCompany
	}

	switch role.Value {
	case "SUPERADMIN":
		// allowed := pkg.UserHasPermission(userID, domain.DELETE_SUPER_ADMIN_USER)
		return false, canUpdateCompany
	case "ADMIN":
		var requiredPermissions = []domain.USER_PERMISSION{domain.UPDATE_USER, domain.UPDATE_ADMIN_USER}
		allowed := pkg.UserHasAllPermissions(userID, requiredPermissions)
		println("allowed", allowed)
		return allowed, canUpdateCompany
	case "CASHIER":
		var requiredPermissions = []domain.USER_PERMISSION{domain.UPDATE_USER, domain.UPDATE_CASHIER_USER}
		allowed := pkg.UserHasAllPermissions(userID, requiredPermissions)
		return allowed, canUpdateCompany
	case "CUSTOMER":
		var requiredPermissions = []domain.USER_PERMISSION{domain.UPDATE_USER, domain.UPDATE_CUSTOMER_USER}
		allowed := pkg.UserHasAllPermissions(userID, requiredPermissions)
		return allowed, canUpdateCompany
	}
	return false, canUpdateCompany
}

func deleteUserGuard(c *fiber.Ctx, deleteUserID int) bool {
	userID, _ := pkg.GetAuthedUser(c)

	// allowed := pkg.UserHasPermission(userID, domain.DELETE_USER)
	// if !allowed {
	// 	return false
	// }

	role, err := pkg.EntClient().User.Query().Where(user.ID(deleteUserID)).QueryRole().First(c.Context())
	if err != nil {
		return false
	}

	switch role.Value {
	case "SUPERADMIN":
		// allowed := pkg.UserHasPermission(userID, domain.DELETE_SUPER_ADMIN_USER)
		return false
	case "ADMIN":
		var requiredPermissions = []domain.USER_PERMISSION{domain.DELETE_USER, domain.DELETE_ADMIN_USER}
		allowed := pkg.UserHasAllPermissions(userID, requiredPermissions)
		println("allowed", allowed)
		return allowed
	case "CASHIER":
		var requiredPermissions = []domain.USER_PERMISSION{domain.DELETE_USER, domain.DELETE_CASHIER_USER}
		allowed := pkg.UserHasAllPermissions(userID, requiredPermissions)
		return allowed
	case "CUSTOMER":
		var requiredPermissions = []domain.USER_PERMISSION{domain.DELETE_USER, domain.DELETE_CUSTOMER_USER}
		allowed := pkg.UserHasAllPermissions(userID, requiredPermissions)
		return allowed
	}
	return false
}

func verifyCreateUserRolePermission(createRoleID int, userID int) bool {
	role, err := pkg.EntClient().Role.Query().Where(role.ID(createRoleID)).First(context.Background())
	println("CREATE ROLE", role)
	if err != nil {
		return false
	}
	println("CREATE ROLE", role.Value)
	switch role.Value {
	case "SUPERADMIN":
		allowed := pkg.UserHasPermission(userID, domain.CREATE_ADMIN_USER)
		return allowed
	case "ADMIN":
		allowed := pkg.UserHasPermission(userID, domain.CREATE_ADMIN_USER)
		return allowed
	case "CASHIER":
		allowed := pkg.UserHasPermission(userID, domain.CREATE_CASHIER_USER)
		return allowed
	case "CUSTOMER":
		allowed := pkg.UserHasPermission(userID, domain.CREATE_CUSTOMER_USER)
		return allowed

	}
	return false
}
