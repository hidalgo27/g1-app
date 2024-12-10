package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hidalgo27/app-g1/cmd/server"

	//"github.com/go-playground/validator/v10"
	"github.com/hidalgo27/app-g1/api/controllers"
	"github.com/hidalgo27/app-g1/config"
)

func GoFiberRoutes(app *fiber.App) {
	// Definicion de los routes
	api := app.Group("/api/v1")
	config.Use("/inquire", api, controllers.NewInquireController())

	config.Use("/package", api, controllers.NewPackageController())

	config.Use("/auth", api, controllers.NewAuthController())

	// Rutas protegidas
	users := api.Group("/users", server.JWTMiddleware())
	config.Use("", users, controllers.NewUserController())

	protected := api.Group("/protected", server.JWTMiddleware())
	config.Use("", protected, controllers.NewProtectedController())

	//Rutas protegidas por roles
	adminRoutes := api.Group("/admin", server.JWTMiddleware(), server.RoleMiddleware("SuperAdmin"))
	config.Use("/dashboard", adminRoutes, controllers.NewAdminController())

	// Rutas protegidas por permisos
	editorRoutes := api.Group("/sales", server.JWTMiddleware(), server.PermissionMiddleware("view_dashboard"))
	config.Use("/dashboard", editorRoutes, controllers.NewEditorController())
}
