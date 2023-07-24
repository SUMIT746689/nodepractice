package routes

import (
	"github.com/gofiber/fiber/v2"
	"pos/cmd/api_server/controllers/auth_controller"
	"pos/cmd/api_server/middlewares"
)

func AuthorizationRoutes() *fiber.App {
	app := fiber.New()

	router := app.Use(middleware.Protected())

	router.Post("/roles", authcontroller.CreateRole)
	//router.Patch("/users/:id", usercontroller.Update)
	//router.Delete("/users/:id", usercontroller.Delete)
	//router.Get("/users", usercontroller.Index)

	return app
}
