package page

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	pages "github.com/jainam240101/doc-create/server/domain/Pages"
	"github.com/jainam240101/doc-create/server/dto"
	"github.com/jainam240101/doc-create/server/helpers"
)

type DefaultPageService struct {
	repo pages.PageRepositoryDb
}

type Title struct {
	Title string `json:"title"`
	Slug  string `json:"slug"`
}
type PageService interface {
	CreatePage(pages.PageModel) (*dto.PageResponse, error)
	GetTitlesofDocument(string) ([]Title, error)
	GetDataofPage(string) (*dto.PageResponse, error)
	UpdatePage(string, pages.PageModel) (*dto.PageResponse, error)
	EditingPage(string, string) (*dto.PageResponse, error)
	DeletePage(string, string) error
	ChangeOrder(string, []string, string) ([]Title, error)
	ForkPages(string, string, string) error
}

func NewPageService(repo pages.PageRepositoryDb) DefaultPageService {
	return DefaultPageService{repo: repo}
}

func (db DefaultPageService) CreatePage(page pages.PageModel) (*dto.PageResponse, error) {
	page.ID = uuid.New()
	page.Slug = helpers.CreateSlug(page.Name) + "--" + page.ID.String()
	data, err := db.repo.CreateNewPage(page)
	if err != nil {
		return nil, err
	}
	return data.ToDto(), err
}
func (db DefaultPageService) GetTitlesofDocument(documentId string) ([]Title, error) {

	data, err := db.repo.GetTitlesofDocument(documentId)
	var titles []Title
	if err != nil {
		fmt.Println("Error --- ", err)
		return nil, err
	}
	for _, v := range data {
		titles = append(titles, Title{
			Title: v.Name,
			Slug:  v.Slug,
		})
	}
	return titles, nil
}

func (db DefaultPageService) GetDataofPage(slug string) (*dto.PageResponse, error) {
	data, err := db.repo.GetDataofPage(slug)
	if err != nil {
		return nil, err
	}
	return data.ToDto(), nil
}

func (db DefaultPageService) UpdatePage(slug string, page pages.PageModel) (*dto.PageResponse, error) {
	id := strings.SplitAfter(slug, "--")[1]
	fmt.Println("Service ID ---- ", id)

	if page.Name != "" {
		page.Slug = helpers.CreateSlug(page.Name) + "--" + id
	}
	newValue, err := db.repo.UpdatePage(slug, page, id)
	if err != nil {
		return nil, err
	}
	return newValue.ToDto(), nil
}

func (db DefaultPageService) DeletePage(userId string, slug string) error {
	err := db.repo.DeletePage(slug, userId)
	if err != nil {
		return err
	}
	return nil
}

func (db DefaultPageService) ChangeOrder(documentId string, value []string, userId string) ([]Title, error) {
	data, err := db.repo.ChangeOrder(documentId, value, userId)
	if err != nil {
		return nil, err
	}
	var titles []Title
	if err != nil {
		fmt.Println("Error --- ", err)
		return nil, err
	}
	for _, v := range data {
		titles = append(titles, Title{
			Title: v.Name,
			Slug:  v.Slug,
		})
	}
	return titles, nil
}

func (db DefaultPageService) ForkPages(documentId string, ownerId string, newDocumentId string) error {
	data, err := db.repo.GetTitlesofDocument(documentId)
	if err != nil {
		fmt.Println("Error --- ", err)
		return err
	}
	for _, value := range data {
		value.ID = uuid.New()
		value.OwnerId = ownerId
		value.Slug = helpers.CreateSlug(value.Name) + "--" + value.ID.String()
		value.DocumentId = newDocumentId
		_, err := db.repo.CreateNewPage(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (db DefaultPageService) EditingPage(slug string, userId string) (*dto.PageResponse, error) {
	data, err := db.repo.EditingPage(slug, userId)
	if err != nil {
		return nil, err
	}
	return data.ToDto(), nil
}
