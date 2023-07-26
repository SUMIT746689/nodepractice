package usercontroller

import (
	"github.com/gofiber/fiber/v2"
	"pos/internal/domain"
	"pos/pkg"
)

func createUserGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.CREATE_USER)
	if !allowed {
		return false
	}

	allowed = pkg.UserHasPermission(userID, domain.CREATE_ADMIN_USER)
	if !allowed {
		return false
	}

	return true

}
