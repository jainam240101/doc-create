package documents

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	documents "github.com/jainam240101/doc-create/server/domain/Documents"
	"github.com/jainam240101/doc-create/server/dto"
	"github.com/jainam240101/doc-create/server/helpers"
)

type DefaultDocumentService struct {
	repo documents.DocumentRepositoryDb
}

type DocumentService interface {
	CreateDocument(documents.DocumentModel) (*dto.DocumentResponse, error)
	ReadDocument() (*dto.DocumentResponse, error)
	SearchDocument(string) ([]dto.DocumentResponse, error)
	OwnedDocuments(string) ([]dto.DocumentResponse, error)
	ReadAllProjectsPublishedByUser(string) ([]dto.DocumentResponse, error)
	GetDocumentData(string) (*dto.DocumentResponse, error)
	UpdateDocument(string, string, documents.DocumentModel) (*dto.DocumentResponse, error)
	DeleteDocument(string, string) error
	ForkDocument(string, string) (*dto.DocumentResponse, error)
}

func NewDocumentService(repo documents.DocumentRepositoryDb) DefaultDocumentService {
	return DefaultDocumentService{repo: repo}
}

func (s DefaultDocumentService) CreateDocument(d documents.DocumentModel) (*dto.DocumentResponse, error) {
	fmt.Println("Page ", d)
	d.ID = uuid.New()
	d.Slug = helpers.CreateSlug(d.Name) + "--" + d.ID.String()
	document, err := s.repo.CreateDocument(d)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return document.ToDto(), nil
}

func (s DefaultDocumentService) ReadDocument() (*dto.DocumentResponse, error) {
	document, err := s.repo.ReadDocument()
	if err != nil {
		fmt.Println(err)
	}
	return document.ToDto(), nil
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
func (db DefaultDocumentService) ReadAllProjectsPublishedByUser(userid string) ([]dto.DocumentResponse, error) {
	data, err := db.repo.ReadAllProjectsPublishedByUser(userid)
	if err != nil {
		return nil, err
	}
	finalData := []dto.DocumentResponse{}
	for _, values := range data {
		finalData = append(finalData, *values.ToDto())
	}
	return finalData, nil
}

func (db DefaultDocumentService) GetDocumentData(slug string) (*dto.DocumentResponse, error) {
	data, err := db.repo.ReadSpecificProjectUsingSlug(slug)
	if data.Name == "" {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}
	return data.ToDto(), nil
}

func (db DefaultDocumentService) UpdateDocument(slug string, userId string, d documents.DocumentModel) (*dto.DocumentResponse, error) {
	id := strings.SplitAfter(slug, "--")[1]
	if d.Name != "" {
		d.Slug = helpers.CreateSlug(d.Name) + "--" + id
	}
	newValue, err := db.repo.UpdateDocument(id, slug, userId, d)
	if err != nil {
		return nil, err
	}
	return newValue.ToDto(), nil
}

func (db DefaultDocumentService) DeleteDocument(userId string, slug string) error {
	err := db.repo.DeleteDocument(userId, slug)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultDocumentService) ForkDocument(slug string, ownerId string) (*dto.DocumentResponse, error) {
	data, err := s.repo.ReadSpecificProjectUsingSlug(slug)
	if data.Name == "" {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}
	data.Name = data.Name + "-Forked"
	data.ID = uuid.New()
	data.Status = "ongoing"
	data.Slug = helpers.CreateSlug(data.Name) + "--" + data.ID.String()
	data.OwnerId = ownerId
	document, err := s.repo.CreateDocument(*data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return document.ToDto(), nil
}
