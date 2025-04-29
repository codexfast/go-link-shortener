package services

import (
    "errors"
    "link-shortener/internal/models"
    "link-shortener/internal/repositories"
    "math/rand"
    "time"
)

type LinkService struct {
    repo *repositories.LinkRepository
}

func generateRandomCode(n int) string {
    // Definindo o conjunto de caracteres válidos para o código
    letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

    // Inicializando o rand com um valor baseado no tempo
    rand.Seed(time.Now().UnixNano())

    // Gerando um código aleatório de 'n' caracteres
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func NewLinkService(repo *repositories.LinkRepository) *LinkService {
    return &LinkService{repo: repo}
}

func (s *LinkService) CreateLink(link *models.Link) error {
    if link.URL == "" {
        return errors.New("URL é obrigatória")
    }

    link.Code = generateRandomCode(10)

    return s.repo.Save(*link) // salva o valor em si
}

func (s *LinkService) GetByCode(code string) (models.Link, error) {
    return s.repo.FindByCode(code)
}