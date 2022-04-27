package models

// Url holds url data
type Url struct {
	Shortened string `json:"shortened,omitempty"`
	Original  string `json:"original,omitempty"`
	Count     int    `json:"count,omitempty"`
}
