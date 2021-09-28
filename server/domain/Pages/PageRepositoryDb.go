package pages

import (
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
	if err := db.Client.Where("document_id = ?", documentId).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (db PageRepositoryDb) GetDataofPage(slug string) (*PageModel, error) {
	var data PageModel
	if err := db.Client.Where("slug = ?", slug).Find(&data).Error; err != nil {
		return nil, err
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


