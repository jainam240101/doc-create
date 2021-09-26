package documents

import (
	"fmt"

	"github.com/google/uuid"
	documents "github.com/jainam240101/doc-create/server/domain/Documents"
	"github.com/jainam240101/doc-create/server/dto"
)

type DefaultDocumentService struct {
	repo documents.DocumentRepositoryDb
}

type DocumentService interface {
	CreateDocument(documents.DocumentModel) (*dto.DocumentResponse, error)
	ReadDocument() (*documents.DocumentModel, error)
	SearchDocument(string) ([]dto.DocumentResponse, error)
	OwnedDocuments(string) ([]dto.DocumentResponse, error)
}

func NewDocumentService(repo documents.DocumentRepositoryDb) DefaultDocumentService {
	return DefaultDocumentService{repo: repo}
}

func (s DefaultDocumentService) CreateDocument(d documents.DocumentModel) (*dto.DocumentResponse, error) {
	fmt.Println("Page ", d)
	d.ID = uuid.New()
	d.Slug = d.Slug + "-" + d.ID.String()
	document, err := s.repo.CreateDocument(d)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return document.ToDto(), nil
}

func (s DefaultDocumentService) ReadDocument() (*documents.DocumentModel, error) {
	document, err := s.repo.ReadDocument()
	if err != nil {
		fmt.Println(err)
	}
	return document, nil
}

func (db DefaultDocumentService) SearchDocument(searchString string) ([]dto.DocumentResponse, error) {
	data, err := db.repo.SearchPublsihedDocument(string("%" + searchString + "%"))
	if err != nil {
		return nil, err
	}
	finalData := []dto.DocumentResponse{}
	for _, values := range data {
		finalData = append(finalData, *values.ToDto())
	}
	return finalData, nil
}
func (db DefaultDocumentService) OwnedDocuments(userid string) ([]dto.DocumentResponse, error) {
	data, err := db.repo.OwnedDocuments(userid)
	if err != nil {
		return nil, err
	}
	finalData := []dto.DocumentResponse{}
	for _, values := range data {
		finalData = append(finalData, *values.ToDto())
	}
	return finalData, nil
}
