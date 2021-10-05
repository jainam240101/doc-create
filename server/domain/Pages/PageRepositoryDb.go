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

func (db PageRepositoryDb) EditingPage(slug string, userId string) (*PageModel, error) {
	var data PageModel
	if err := db.Client.Where("slug=? and owner_id=?", slug, userId).Find(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (db PageRepositoryDb) UpdatePage(slug string, p PageModel, id string) (*PageModel, error) {
	var page PageModel
	result := db.Client.Model(&PageModel{}).Where("slug = ? AND owner_id=?", slug, p.OwnerId).Updates(p)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("you Do not have the permission to update ")
	}
	if err := db.Client.Where("id = ?", id).Find(&page).Error; err != nil {
		return nil, err
	}
	return &page, nil
}

func (db PageRepositoryDb) DeletePage(slug string, userId string) error {
	result := db.Client.Where("slug=? AND owner_id=?", slug, userId).Delete(&PageModel{})
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("you do not have the permission to delete")
	}
	return nil
}

func (db PageRepositoryDb) ChangeOrder(documentId string, data []string, userId string) ([]PageModel, error) {
	for i, value := range data {
		result := db.Client.Model(&PageModel{}).Where("slug = ? and document_id=? and owner_id=?", value, documentId, userId).Updates(PageModel{OrderNo: i + 1})
		if result.Error != nil || result.RowsAffected == 0 {
			fmt.Println("Error -- ", result.Error)
			return nil, errors.New("you Do not have the permission to change the order")
		}
	}
	var finalData []PageModel
	if err := db.Client.Where("document_id = ?", documentId).Order("order_no asc").Find(&finalData).Error; err != nil {
		return nil, err
	}
	return finalData, nil
}
