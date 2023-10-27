package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jrjenkins555/fyrefly/pkg/routes"
	"github.com/jrjenkins555/fyrefly/pkg/middleware"
)

func main() {
    app := gin.Default()

	// middleware
	middleware.GinMiddleware(app)

	// setup public routes
	routes.PublicRoutes(app)

	fmt.Println("running on port 8080...")
    app.Run(":8080")
}
