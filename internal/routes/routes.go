package routes

import (
    "github.com/gin-gonic/gin"
    "link-shortener/internal/handlers"
    "link-shortener/internal/services"
    "link-shortener/internal/repositories"
)

func RegisterRoutes(r *gin.Engine) {

    repo := repositories.NewLinkRepository()
    service := services.NewLinkService(repo)
    handler := handlers.NewLinkHandler(service)
    
    shortenerRoutes := r.Group("/shortener")
    {
        shortenerRoutes.POST("/", handler.CreateLink)
    }

    r.GET("/:code", handler.Redirect)
}