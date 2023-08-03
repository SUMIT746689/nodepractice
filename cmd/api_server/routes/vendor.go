package routes

import (
	vendorcontroller "pos/cmd/api_server/controllers/vendor_controller"
	middleware "pos/cmd/api_server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func VenderRoutes() *fiber.App {
	app := fiber.New()

	router := app.Use(middleware.Protected())

	router.Post("/vendors", vendorcontroller.Create)
	router.Patch("/vendors/:id", vendorcontroller.Update)
	router.Delete("/vendors/:id", vendorcontroller.Delete)
	router.Get("/vendors", vendorcontroller.Index)

	return app
}
