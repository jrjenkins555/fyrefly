package main

import (
	"fmt"
    "github.com/gofiber/fiber/v2"
	"github.com/jrjenkins555/fyrefly/pkg/routes"
)

func main() {
    app := fiber.New()

	// setup public routes
	routes.PublicRoutes(app)

	fmt.Println("running on port 3000...")
    app.Listen(":3000")
}