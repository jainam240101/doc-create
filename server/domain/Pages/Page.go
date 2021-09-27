package pages

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PageModel struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	OwnerId     string
	DocumentId  string `json:"documentId"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Content     string `gorm:"text"`
}
