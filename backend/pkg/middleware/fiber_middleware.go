package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func GinMiddleware(r *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 3600

	r.Use(cors.New(config))
}
