package utils

import (
	"github.com/teris-io/shortid"
)

// Genereate new short URL
func Generate() (string, error) {
	return shortid.Generate()
}
