package main

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
	"github.com/jrjenkins555/fyrefly/pkg/routes"
	"github.com/jrjenkins555/fyrefly/pkg/middleware"
	"github.com/jrjenkins555/fyrefly/app/service"
)

func main() {
    app := fiber.New()

	// middleware
	middleware.FiberMiddleware(app)

	// setup public routes
	routes.PublicRoutes(app)

	service.CallOpenAI()

	fmt.Println("running on port 8080...")
    app.Listen(":8080")
}