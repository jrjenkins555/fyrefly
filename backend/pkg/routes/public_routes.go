package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/jrjenkins555/fyrefly/app/controllers"
)

func PublicRoutes(app *fiber.App) {
    // file upload routes
    route := app.Group("/api/v1")
    route.Get("/test", controllers.TestEndpoint)
    route.Post("/upload", controllers.Upload)
}
