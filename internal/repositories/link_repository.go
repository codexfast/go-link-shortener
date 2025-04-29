package repositories

import (
    "link-shortener/internal/models"
    "errors"
)

type LinkRepository struct {
    storage map[string]models.Link
}

func NewLinkRepository() *LinkRepository {
    return &LinkRepository{storage: make(map[string]models.Link)}
}

func (r *LinkRepository) Save(link models.Link) error {
    r.storage[link.Code] = link
    return nil
}

func (r *LinkRepository) FindByCode(code string) (models.Link, error) {
    for _, link := range r.storage {
        if link.Code == code {
            return link, nil
        }
    }
    return models.Link{}, errors.New("link não encontrado")
}
