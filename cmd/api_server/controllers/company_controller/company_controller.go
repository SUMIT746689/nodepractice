package usercontroller

import (
	"pos/internal/domain"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	// _, roleID := pkg.GetAuthedUser(c)

	// if err != nil {
	// 	return c.SendStatus(fiber.StatusInternalServerError)
	// }

	if allowed := indexCompanyGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	users_, err := pkg.EntClient().Company.Query().All(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"users": users_,
	})
}

func Create(c *fiber.Ctx) error {
	req := new(domain.Company)

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if allowed := createCompanyGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	name := req.Name
	domain := req.Domain

	u, err := pkg.EntClient().Company.Create().SetName(name).SetDomain(domain).Save(c.Context())

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(u)
}

func Update(c *fiber.Ctx) error {
	req := new(domain.UpdateCompanyRequest)
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if allowed := updateCompanyGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	q := pkg.EntClient().Company.UpdateOneID(userID)

	if req.Name != "" {
		q.SetName(req.Name)
	}

	if req.Domain != "" {
		q.SetDomain(req.Domain)
	}

	_, err = q.Save(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if allowed := deleteCompanyGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = pkg.EntClient().Company.DeleteOneID(userID).Exec(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
