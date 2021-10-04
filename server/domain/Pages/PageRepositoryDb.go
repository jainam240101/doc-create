package pages

import (
	"errors"
	"fmt"

	Document "github.com/jainam240101/doc-create/server/domain/Documents"
	"gorm.io/gorm"
)

type PageRepositoryDb struct {
	Client *gorm.DB
}

func NewPageRepositoryDb(Client *gorm.DB) PageRepositoryDb {
	return PageRepositoryDb{Client: Client}
}

func (db PageRepositoryDb) CreateNewPage(p PageModel) (*PageModel, error) {
	if err := db.Client.Save(&p).Error; err != nil {
		return nil, err
	}
	return &p, nil
}
func (db PageRepositoryDb) GetTitlesofDocument(documentId string) ([]PageModel, error) {
	var data []PageModel
	if err := db.Client.Where("document_id = ?", documentId).Order("order_no asc").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (db PageRepositoryDb) GetDataofPage(slug string) (*PageModel, error) {
	var doc Document.DocumentModel
	var data PageModel
	if err := db.Client.Where("slug=?", slug).Find(&data).Error; err != nil {
		return nil, err
	}
	if err := db.Client.Where("id=?", data.DocumentId).Find(&doc).Error; err != nil {
		return nil, err
	}
	if doc.Status != "Published" {
		return nil, errors.New("page not published")

	}
	return &data, nil
}

func (db PageRepositoryDb) UpdatePage(slug string, p PageModel) (*PageModel, error) {
	var page PageModel
	if err := db.Client.Model(&PageModel{}).Where("slug = ?", slug).Updates(p).Error; err != nil {
		return nil, err
	}
	if err := db.Client.Where("slug = ?", slug).Find(&page).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (db PageRepositoryDb) DeletePage(slug string) error {
	if err := db.Client.Where(" slug=?", slug).Delete(&PageModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (db PageRepositoryDb) ChangeOrder(documentId string, data []string) ([]PageModel, error) {
	for i, value := range data {
		fmt.Println("Index -- ", i)
		fmt.Println("Value -- ", value)
		if err := db.Client.Model(&PageModel{}).Where("slug = ? and document_id=?", value, documentId).Updates(PageModel{OrderNo: i + 1}).Error; err != nil {
			fmt.Println("Error -- ", err)
			// return nil, err
		}
	}
	var finalData []PageModel
	if err := db.Client.Where("document_id = ?", documentId).Order("order_no asc").Find(&finalData).Error; err != nil {
		return nil, err
	}
	return finalData, nil
}
