package main

import (
	"pos/cmd/api_server/routes"
	"pos/pkg"
	"pos/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.InitAuthConfig()

	pkg.InitEnt()

	app := fiber.New()
	router := app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	v1 := api.Group("/v1")

	v1.Mount("", routes.AuthRoutes())
	v1.Mount("", routes.UserRoutes())
	v1.Mount("", routes.AuthorizationRoutes())
	v1.Mount("", routes.CompanyRoutes())
	v1.Mount("", routes.SupplierRoutes())
	v1.Mount("", routes.VenderRoutes())

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}
