package routes

import (
	companycontroller "pos/cmd/api_server/controllers/company_controller"
	middleware "pos/cmd/api_server/middlewares"

	"github.com/gofiber/fiber/v2"
)

func CompanyRoutes() *fiber.App {
	app := fiber.New()

	router := app.Use(middleware.Protected())

	router.Post("/companies", companycontroller.Create)
	router.Patch("/companies/:id", companycontroller.Update)
	router.Delete("/companies/:id", companycontroller.Delete)
	router.Get("/companies", companycontroller.Index)

	return app
}
