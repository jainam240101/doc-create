package page

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	pages "github.com/jainam240101/doc-create/server/domain/Pages"
	"github.com/jainam240101/doc-create/server/dto"
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
	DeletePage(string) error
	ChangeOrder(string,[]string)
}

func NewPageService(repo pages.PageRepositoryDb) DefaultPageService {
	return DefaultPageService{repo: repo}
}

func (db DefaultPageService) CreatePage(page pages.PageModel) (*dto.PageResponse, error) {
	page.ID = uuid.New()
	page.Slug = page.Slug + "--" + page.ID.String()
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
	if page.Name != "" {
		page.Slug = page.Slug + "--" + id
	}
	newValue, err := db.repo.UpdatePage(slug, page)
	if err != nil {
		return nil, err
	}
	return newValue.ToDto(), nil
}

func (db DefaultPageService) DeletePage(slug string) error {
	err := db.repo.DeletePage(slug)
	if err != nil {
		return err
	}
	return nil
}

func (db DefaultPageService) ChangeOrder(documentId string,value []string) {
	db.repo.ChangeOrder(documentId,value)
	// if err != nil {
	// 	return err
	// }
	// return nil
}
