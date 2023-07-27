package usercontroller

import (
	"context"
	"pos/ent/role"
	"pos/ent/user"
	"pos/internal/app"
	"pos/internal/domain"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func createUserGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	body := new(app.User)

	if err := c.BodyParser(body); err != nil {
		return false
	}

	allowed := pkg.UserHasPermission(userID, domain.CREATE_USER)
	if !allowed {
		return false
	}

	isAllowed := verifyCreateUserRolePermission(body.RoleID, userID)

	return isAllowed

	// allowed = pkg.UserHasPermission(userID, domain.CREATE_ADMIN_USER)
	// if !allowed {
	// 	return false
	// }

	// return true

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
