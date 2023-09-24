package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jrjenkins555/fyrefly/api/handlers"
)

func PublicRoutes(app *fiber.App) {
	// file upload routes
	route := app.Group("/api/v1")
	route.Get("/test", handlers.TestEndpoint)
}
