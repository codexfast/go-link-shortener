package main

import (
    "github.com/gin-gonic/gin"
    "link-shortener/internal/routes"
)

func main() {
    r := gin.Default()
    routes.RegisterRoutes(r)
    r.Run(":8080")
}