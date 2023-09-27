package middleware

import (
    "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
)

func FiberMiddleware(a *fiber.App) error {
	// Allow cross port comms
    a.Use(
        cors.New(cors.Config{
            AllowOrigins:     "http://localhost:3000",
            AllowMethods:     "GET, POST, PUT, DELETE",
            AllowHeaders:     "Origin, Content-Type, Accept",
            ExposeHeaders:    "Content-Length",
            AllowCredentials: true,
            MaxAge:           3600,
        }),
        logger.New(),
    )

    return nil
}
