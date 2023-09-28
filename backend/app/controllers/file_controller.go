package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	// "io/ioutil"
)

func TestEndpoint(c *fiber.Ctx) error {
	fmt.Println("Test endpoint works")
	return nil
}

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fileContent, err := file.Open()
	defer fileContent.Close()

	// fileBytes, err := ioutil.ReadAll(fileContent)
	// fileString := string(fileBytes)
	
	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
	})
	return nil
}
