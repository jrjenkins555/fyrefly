package controllers

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func TestEndpoint(c *gin.Context) {
    fmt.Println("Test endpoint works")
}
