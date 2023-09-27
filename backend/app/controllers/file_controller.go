package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func TestEndpoint(c *fiber.Ctx) error {
	fmt.Println("Test endpoint works")
	return nil
}

func Upload(c *fiber.Ctx) error {
	fmt.Println("gettig to func")
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	contentType := file.Header.Get("Content-Type")

	// Print information about the uploaded file
	fmt.Println("File Name:", file.Filename)
	fmt.Println("File Size:", file.Size)
	fmt.Println("File Type:", contentType)
	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
	})
	return nil
}
