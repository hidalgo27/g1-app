package routes

import (
	"github.com/gofiber/fiber/v2"
	//"github.com/go-playground/validator/v10"
	"github.com/hidalgo27/app-g1/api/controllers"
	"github.com/hidalgo27/app-g1/config"
)

func GoFiberRoutes(app *fiber.App) {
	// Definicion de los routes
	api := app.Group("/api/v1")
	config.Use("/inquire", api, controllers.NewInquireController())
	config.Use("/package", api, controllers.NewPackageController())

}
