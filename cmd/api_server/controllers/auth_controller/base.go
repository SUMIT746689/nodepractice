package authcontroller

import (
	"log"
	"os"
	"pos/ent"
	"pos/ent/permission"
	"pos/ent/user"
	"pos/internal/app"
	userrepo "pos/internal/repository/user_repo"
	"pos/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	// "github.com/golang-jwt/jwt/v4"
	"github.com/golang-jwt/jwt/v5"
)

func Register(c *fiber.Ctx) error {
	req := new(app.User)

	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	_, err = userrepo.Save(c.Context(), req)
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}

func Login(c *fiber.Ctx) error {
	req := new(loginRequest)
	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(err.Error())
	}

	identity := req.Identity
	pass := req.Password

	exists, err := pkg.EntClient().User.Query().Where(user.Username(identity)).Exist(c.Context())
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if !exists {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	user_, err := pkg.EntClient().User.Query().Where(user.Username(identity)).Only(c.Context())
	if err != nil {
		log.Println(err)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	if pkg.Hash(pass) != user_.Password {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["identity"] = identity
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t, "env": os.Getenv("APP_ENV")})
}

func Me(c *fiber.Ctx) error {
	usr := c.Locals("user").(*jwt.Token)
	claim := usr.Claims.(jwt.MapClaims)
	identity := claim["identity"].(string)

	_user, err := pkg.EntClient().User.Query().Where(user.Username(identity)).WithPermissions(func(query *ent.PermissionQuery) {
		query.Select(permission.FieldValue)
	}).WithRole(func(query *ent.RoleQuery) {
		query.WithPermissions(func(query *ent.PermissionQuery) {
			query.Select(permission.FieldValue)
		})
	}).Only(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"user": _user})
}

type loginRequest struct {
	Identity string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func CreateRole(c *fiber.Ctx) error {
	type CreateRoleRequest struct {
		Title string `json:"title" validate:"required"`
	}
	req := new(CreateRoleRequest)
	err := pkg.BindNValidate(c, req)
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	_role, err := pkg.EntClient().Role.Create().SetTitle(req.Title).Save(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.JSON(_role)
}

func IndexRole(c *fiber.Ctx) error {
	all, err := pkg.EntClient().Role.Query().All(c.Context())
	if err != nil {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}

	return c.JSON(all)
}
