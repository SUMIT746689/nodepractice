package usercontroller

import (
	"pos/internal/app"
	"pos/internal/domain"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func createCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	body := new(app.User)

	if err := c.BodyParser(body); err != nil {
		return false
	}

	allowed := pkg.UserHasPermission(userID, domain.CREATE_COMPANY)
	return allowed

}
func updateCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	body := new(app.User)

	if err := c.BodyParser(body); err != nil {
		return false
	}

	allowed := pkg.UserHasPermission(userID, domain.UPDATE_COMPANY)
	return allowed

}

func deleteCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)

	allowed := pkg.UserHasPermission(userID, domain.DELETE_COMPANY)
	if !allowed {
		return false
	}

	return allowed
}
