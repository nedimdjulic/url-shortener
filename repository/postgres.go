package repository

import (
	"database/sql"
	"log"

	"github.com/nedimdjulic/url-shortener/models"
)

// UrlDB represents url postgres implementation
type UrlDB struct {
	client *sql.DB
}

// NewUrlDB returns and instance of UrlDB
func NewUrlDB(c *sql.DB) *UrlDB {
	return &UrlDB{client: c}
}

// PGRepo interface --
type PGRepo interface {
	CreateShortURL(url models.Url) error
	RetrieveByShortURL(shortURL string) (*models.Url, error)
	UpdateCount(shortURL string, count int) error
	RetrieveCount(shortURL string) (int, error)
	Delete(shortURL string) error
}

// CreateShortURL inserts new URL into db
func (u *UrlDB) CreateShortURL(url models.Url) error {
	_, err := u.client.Exec("INSERT INTO url(shortened, original, count) VALUES($1, $2, $3)", url.Shortened, url.Original, url.Count)
	if err != nil {
		return err
	}

	return nil
}

// RetrieveByShortURL fetches URL data from database, provided shortURL
func (u *UrlDB) RetrieveByShortURL(shortURL string) (*models.Url, error) {
	var url models.Url

	err := u.client.QueryRow("SELECT original, count FROM url WHERE shortened=$1", shortURL).Scan(&url.Original, &url.Count)
	if err != nil {
		log.Println(err)
	}

	return &url, err
}

// UpdateCount increments redirection counter
func (u *UrlDB) UpdateCount(shortURL string, count int) error {
	_, err := u.client.Exec("UPDATE url SET count=$1 WHERE shortened=$2", count, shortURL)

	return err
}

// RetrieveCount fetches count of url redirections
func (u *UrlDB) RetrieveCount(shortURL string) (int, error) {
	var count int

	err := u.client.QueryRow("SELECT count FROM url WHERE shortened=$1", shortURL).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Delete removes URL entry from database
func (u *UrlDB) Delete(shortURL string) error {
	_, err := u.client.Exec("DELETE * FROM url WHERE shortened=$1", shortURL)

	return err
}
