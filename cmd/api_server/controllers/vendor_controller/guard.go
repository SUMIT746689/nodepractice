package vendorcontroller

import (
	"pos/internal/domain"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func indexGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.INDEX_VENDOR)
	return allowed
}

func createGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.CREATE_VENDOR)
	return allowed
}

func updateGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.UPDATE_VENDOR)
	return allowed
}

func deleteGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.DELETE_VENDOR)
	return allowed
}
