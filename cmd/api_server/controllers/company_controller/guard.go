package companycontroller

import (
	"pos/internal/domain"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func indexCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.INDEX_COMPANY)
	return allowed
}

func createCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.CREATE_COMPANY)
	return allowed
}

func updateCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.UPDATE_COMPANY)
	return allowed
}

func deleteCompanyGuard(c *fiber.Ctx) bool {
	userID, _ := pkg.GetAuthedUser(c)
	allowed := pkg.UserHasPermission(userID, domain.DELETE_COMPANY)
	return allowed
}
