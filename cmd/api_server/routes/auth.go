package routes

import (
	authcontroller "pos/cmd/api_server/controllers/auth_controller"
	middleware "pos/cmd/api_server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes() *fiber.App {
	app := fiber.New()

	app.Post("/register", authcontroller.Register)
	app.Post("/login", authcontroller.Login)

	protected := app.Group("", middleware.Protected())
	protected.Get("/me", authcontroller.Me)
	protected.Post("/logout", func(c *fiber.Ctx) error {
		return nil
	})

	return app
}
