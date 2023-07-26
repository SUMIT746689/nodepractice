package usercontroller

import (
	"log"
	"pos/internal/app"
	userrepo "pos/internal/repository/user_repo"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	users_, err := pkg.EntClient().User.Query().Limit(10).All(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"users": users_,
	})
}

func Create(c *fiber.Ctx) error {
	req := new(app.User)

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	u, err := userrepo.Save(c.Context(), req)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(u)
}

func Update(c *fiber.Ctx) error {
	req := new(updateUserRequest)
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	err = pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	q := pkg.EntClient().User.UpdateOneID(userID).SetFirstName(req.FirstName).SetLastName(req.LastName).SetUsername(req.Username)

	if req.PhoneNumber != "" {
		q.SetPhoneNumber(req.PhoneNumber)
	} else {
		q.SetNillablePhoneNumber(nil)
	}

	if req.Email != "" {
		q.SetEmail(req.Email)
	} else {
		q.SetNillableEmail(nil)
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

	err = pkg.EntClient().User.DeleteOneID(userID).Exec(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

type updateUserRequest struct {
	FirstName   string `json:"first_name" validate:"required,max=50"`
	LastName    string `json:"last_name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"omitempty,numeric"`
	Email       string `json:"email" validate:"omitempty,email"`
}
