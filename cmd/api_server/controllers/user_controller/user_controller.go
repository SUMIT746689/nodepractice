package usercontroller

import (
	"log"
	"pos/ent/role"
	"pos/internal/app"
	userrepo "pos/internal/repository/user_repo"
	"pos/pkg"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	_, roleID := pkg.GetAuthedUser(c)

	role_, err := pkg.EntClient().Role.Query().Where(role.ID(roleID)).Only(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	switch role_.Value {
	case "SUPERADMIN":
		users_, err := pkg.EntClient().Role.Query().Where(role.Value("ADMIN")).QueryUsers().WithCompany().All(c.Context())
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"users": users_,
		})
	case "ADMIN":
		users_, err := pkg.EntClient().Role.Query().Where(role.Not(role.Value("SUPERADMIN")), role.Not(role.Value("ADMIN"))).QueryUsers().WithCompany().All(c.Context())
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"users": users_,
		})
	}
	x := []string{}
	return c.JSON(fiber.Map{
		"users": x,
	})
}

func Create(c *fiber.Ctx) error {
	req := new(app.User)
	body := new(app.User)

	if err := c.BodyParser(body); err != nil {
		println(err)
	}

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	allowed, companyID := createUserGuard(c, body.RoleID, body.CompanyID)
	if !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if companyID != 0 {
		req.CompanyID = companyID
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

	allowed, canUpdateCompany := updateUserGuard(c, userID)
	if !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
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

	if canUpdateCompany {
		if req.CompanyID != 0 {
			q.SetCompanyID(req.CompanyID)
		} else {
			q.SetNillableEmail(nil)
		}
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

	if allowed := deleteUserGuard(c, userID); !allowed {
		return c.SendStatus(fiber.StatusUnauthorized)
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
	CompanyID   int    `json:"company_id" validate:""`
}
