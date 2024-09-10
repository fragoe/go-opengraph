package book

import (
	"time"
)

// Book contains Open Graph Book structure
type Book struct {
	ISBN        string     `json:"isbn,omitempty"`
	ReleaseDate *time.Time `json:"release_date,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	Authors     []string   `json:"authors,omitempty"`
}
