package vendorcontroller

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

	if allowed := indexGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	vendors_, err := pkg.EntClient().Vendor.Query().All(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"vendors": vendors_,
	})
}

func Create(c *fiber.Ctx) error {
	req := new(domain.Vendor)

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if allowed := createGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	name := req.Name
	email := req.Email

	u, err := pkg.EntClient().Vendor.Create().SetName(name).SetEmail(email).Save(c.Context())

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(u)
}

func Update(c *fiber.Ctx) error {
	req := new(domain.UpdateVendor)
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if allowed := updateGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	q := pkg.EntClient().Vendor.UpdateOneID(userID)

	if req.Name != "" {
		q.SetName(req.Name)
	}
	if req.Name != "" {
		q.SetName(req.Name)
	}
	if req.Email != "" {
		q.SetEmail(req.Email)
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

	if allowed := deleteGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	err = pkg.EntClient().Vendor.DeleteOneID(userID).Exec(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
