package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hidalgo27/app-g1/api/routes"
	"github.com/hidalgo27/app-g1/cmd/database"
	"github.com/hidalgo27/app-g1/cmd/server"
)

func main() {
	app := fiber.New()
	database.InitDB()

	app.Use(server.ContextServerMiddleware())
	app.Use(cors.New())

	// Interceptor de respuestas
	app.Use(server.ResponseInterceptorMiddleware())
	// Routes
	routes.GoFiberRoutes(app)
	// Iniciar servidor
	server.RunServer(app)

}
