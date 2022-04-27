package service

import (
	"github.com/nedimdjulic/url-shortener/models"
	repo "github.com/nedimdjulic/url-shortener/repository"
	"github.com/nedimdjulic/url-shortener/utils"
)

// New URL service
func New(r repo.PGRepo) *URL {
	return &URL{db: r}
}

// URL contains URL service data
type URL struct {
	db repo.PGRepo
}

// Service interface
type Service interface {
	Create(url *models.Url) (*models.Url, error)
	RetrieveFullLink(urlKey string) (string, error)
	RetrieveCount(shortURL string) (int, error)
	DeleteURL(urlKey string) error
}

// Create generates new short URL and saves it
func (u *URL) Create(url *models.Url) (*models.Url, error) {

	shortLink, err := utils.Generate()
	if err != nil {
		return nil, err
	}

	url.Shortened = shortLink

	if err = u.db.CreateShortURL(*url); err != nil {
		return nil, err
	}

	return url, nil
}

// RetrieveFullLink fetches original URL and increments the counter
func (u *URL) RetrieveFullLink(urlKey string) (string, error) {
	url, err := u.db.RetrieveByShortURL(urlKey)
	if err != nil {
		return "", err
	}

	url.Count++

	if err = u.db.UpdateCount(urlKey, url.Count); err != nil {
		return "", err
	}

	return url.Original, nil
}

// GetCount --
func (u *URL) RetrieveCount(urlKey string) (int, error) {
	return u.db.RetrieveCount(urlKey)
}

// DeleteURL --
func (u *URL) DeleteURL(urlKey string) error {
	return u.db.Delete(urlKey)
}
