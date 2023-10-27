package routes

import (
    "github.com/gin-gonic/gin"
    "github.com/jrjenkins555/fyrefly/app/controllers"
)

func PublicRoutes(r *gin.Engine) {
    route := r.Group("/api/v1")

    route.GET("/test", controllers.TestEndpoint)
}
