package pages

import (
	"github.com/google/uuid"
	dto "github.com/jainam240101/doc-create/server/dto"
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
	Content     string `json:"content" gorm:"text"`
}

func (d PageModel) ToDto() *dto.PageResponse {
	return &dto.PageResponse{
		Name:        d.Name,
		Slug:        d.Slug,
		Description: d.Description,
		Content:     d.Content,
	}
}

type PageRepository interface {
	CreatePage(PageModel) (*PageModel, error)
	GetTitlesofDocument(string) ([]PageModel, error)
	GetDataofPage(string) (*PageModel, error)
	UpdatePage(PageModel) (*PageModel, error)
	DeletePage() error
}
