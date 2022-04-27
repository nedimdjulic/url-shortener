package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DB holds database connection
type DB struct {
	*sql.DB
}

// Config holds db config data
type Config struct {
	Host     string
	Port     string
	User     string
	Name     string
	Password string
}

// Open opens connection to db and creates DB
func Open(driver string, config *Config) (*DB, error) {

	dbInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name)

	sqlDB, err := sql.Open(driver, dbInfo)
	if err != nil {
		return nil, err
	}

	return &DB{sqlDB}, nil
}

// Init will create new `url` table if it doesn't exits
func (d *DB) Init() error {
	var exists bool

	err := d.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'url')").Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		_, err = d.Exec("CREATE TABLE url (shortened VARCHAR(255) PRIMARY KEY, original VARCHAR(255), count INT)")
		if err != nil {
			return err
		}
	}

	return nil
}
