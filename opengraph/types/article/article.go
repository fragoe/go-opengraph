package article

import (
	"time"
)

// Article contain Open Graph Article structure
type Article struct {
	PublishedTime  *time.Time `json:"published_time,omitempty"`
	ModifiedTime   *time.Time `json:"modified_time,omitempty"`
	ExpirationTime *time.Time `json:"expiration_time,omitempty"`
	Section        string     `json:"section,omitempty"`
	Tags           []string   `json:"tags,omitempty"`
	Authors        []string   `json:"authors,omitempty"`
}
