package handlers

import (
    "net/http"
    "link-shortener/internal/models"
    "link-shortener/internal/services"
    "github.com/gin-gonic/gin"
)

type LinkRequest struct {
    URL string `json:"url"`
}

type LinkHandler struct {
    service *services.LinkService
}

func NewLinkHandler(service *services.LinkService) *LinkHandler {
    return &LinkHandler{service: service}
}

func (h *LinkHandler) Redirect(c *gin.Context) {
    code := c.Param("code")

    link, err := h.service.GetByCode(code)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Link não encontrado"})
        return
    }

    c.Redirect(http.StatusMovedPermanently, link.URL)
}

// CreateLink cria um novo link encurtado
func (h *LinkHandler) CreateLink(c *gin.Context) {
    var request LinkRequest

    // BindJSON faz o binding do corpo JSON na variável request
    if err := c.BindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Valida que o URL foi passado
    if request.URL == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "URL é obrigatória"})
        return
    }

    // Cria o link
    link := models.Link{URL: request.URL}
    err := h.service.CreateLink(&link)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Responde com o link criado
    c.JSON(http.StatusOK, gin.H{
        "message": "Link encurtado com sucesso",
        "url":     link.URL,
        "code":    link.Code, // Ou o que você definir como código do link
    })
}

