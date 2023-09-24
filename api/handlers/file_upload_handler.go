package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func TestEndpoint(c *fiber.Ctx) error {
	fmt.Println("Test endpoint works")
	return nil
}