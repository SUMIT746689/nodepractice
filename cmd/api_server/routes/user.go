package routes

import (
	usercontroller "pos/cmd/api_server/controllers/user_controller"
	middleware "pos/cmd/api_server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes() *fiber.App {
	app := fiber.New()

	router := app.Use(middleware.Protected())

	router.Post("/users", usercontroller.Create)
	router.Patch("/users", usercontroller.Update)
	router.Get("/users", usercontroller.Index)

	return app
}
