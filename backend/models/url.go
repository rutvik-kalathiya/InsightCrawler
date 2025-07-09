package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Url struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	URL           string         `json:"url"`
	HtmlVersion   string         `json:"html_version"`
	Title         string         `json:"title"`
	H1            int            `json:"h1"`
	H2            int            `json:"h2"`
	H3            int            `json:"h3"`
	H4            int            `json:"h4"`
	H5            int            `json:"h5"`
	H6            int            `json:"h6"`
	InternalLinks int            `json:"internal_links"`
	ExternalLinks int            `json:"external_links"`
	BrokenLinks   datatypes.JSON `json:"broken_links"` // Array of { url, status }
	HasLoginForm  bool           `json:"has_login_form"`
	Status        string         `json:"status"` // queued, running, done, error
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
