package suppliercontroller

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

	suppliers_, err := pkg.EntClient().Supplier.Query().All(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"suppliers": suppliers_,
	})
}

func Create(c *fiber.Ctx) error {
	req := new(domain.Supplier)

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	if allowed := createGuard(c); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	name := req.Name
	address := req.Address
	email := req.Email
	representative := req.Representative

	u := pkg.EntClient().Supplier.Create().SetName(name).SetEmail(email).SetRepresentative(domain.Representative(*representative))

	if req.Address != "" {
		u.SetAddress(address)
	}

	_, err = u.Save(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(u)
}

func Update(c *fiber.Ctx) error {
	req := new(domain.UpdateSupplier)
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

	q := pkg.EntClient().Supplier.UpdateOneID(userID)

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

	err = pkg.EntClient().Supplier.DeleteOneID(userID).Exec(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
