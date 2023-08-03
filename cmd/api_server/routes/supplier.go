package routes

import (
	suppliercontroller "pos/cmd/api_server/controllers/supplier_controller"
	middleware "pos/cmd/api_server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func SupplierRoutes() *fiber.App {
	app := fiber.New()

	router := app.Use(middleware.Protected())

	router.Post("/suppliers", suppliercontroller.Create)
	router.Patch("/suppliers/:id", suppliercontroller.Update)
	router.Delete("/suppliers/:id", suppliercontroller.Delete)
	router.Get("/suppliers", suppliercontroller.Index)

	return app
}
