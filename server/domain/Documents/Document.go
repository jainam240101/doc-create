package documents

import (
	"github.com/google/uuid"
	dto "github.com/jainam240101/doc-create/server/dto"
	"gorm.io/gorm"
)

type DocumentModel struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:char(36);primary_key"`
	OwnerId     string
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Status      string `json:"status" gorm:"default:ongoing"`
	Description string `json:"description" gorm:"type:text"`
	Font        string `json:"font"`
}

func (d DocumentModel) ToDto() *dto.DocumentResponse {
	return &dto.DocumentResponse{
		ID:          d.ID,
		Name:        d.Name,
		Slug:        d.Slug,
		Status:      d.Status,
		Description: d.Description,
		Font:        d.Font,
	}
}

type DocumentRepository interface {
	CreateDocument(DocumentModel) (*DocumentModel, error)
	SearchPublsihedDocument(string) ([]DocumentModel, error)
	OwnedDocuments(string) ([]DocumentModel, error)
	ReadDocument() (*DocumentModel, error)
}
