package documents

import (
	"fmt"

	"gorm.io/gorm"
)

type DocumentRepositoryDb struct {
	Client *gorm.DB
}

func NewRepositoryDb(dbClient *gorm.DB) DocumentRepositoryDb {
	return DocumentRepositoryDb{Client: dbClient}
}

func (db DocumentRepositoryDb) CreateDocument(d DocumentModel) (*DocumentModel, error) {
	if err := db.Client.Save(&d).Error; err != nil {
		return nil, err
	}
	return &d, nil
}
func (db DocumentRepositoryDb) ReadDocument() (*DocumentModel, error) {
	var d DocumentModel
	if err := db.Client.First(&d).Error; err != nil {
		return nil, err
	}
	return &d, nil
}

func (db DocumentRepositoryDb) SearchPublsihedDocument(query string) ([]DocumentModel, error) {
	var documentModel []DocumentModel
	if err := db.Client.Where("LOWER(name) LIKE ? and status= ?", query, "Published").Find(&documentModel).Error; err != nil {
		return nil, err
	}
	return documentModel, nil
}

func (db DocumentRepositoryDb) OwnedDocuments(userid string) ([]DocumentModel, error) {
	var documentModel []DocumentModel
	if err := db.Client.Where("owner_id=?", userid).Find(&documentModel).Error; err != nil {
		return nil, err
	}
	return documentModel, nil
}
func (db DocumentRepositoryDb) ReadSpecificProjectUsingSlug(slug string) (*DocumentModel, error) {
	var documentModel DocumentModel
	if err := db.Client.Where("slug=? AND status=?", slug, "Published").Find(&documentModel).Error; err != nil {
		return nil, err
	}
	return &documentModel, nil
}

func (db DocumentRepositoryDb) UpdateDocument(id string, slug string, userId string, updates DocumentModel) (*DocumentModel, error) {
	fmt.Println("Update --- ", updates.Name)
	fmt.Println("Update --- ", updates.Slug)
	var documentModel DocumentModel
	if err := db.Client.Model(&DocumentModel{}).Where("owner_id = ? AND slug= ? AND status= ?", userId, slug, "ongoing").Updates(updates).Error; err != nil {
		fmt.Println("Error --- ", err)
		return nil, err
	}
	if err := db.Client.Where("id = ?", id).First(&documentModel).Error; err != nil {
		return nil, err
	}
	return &documentModel, nil
}

func (db DocumentRepositoryDb) DeleteDocument(userId string, slug string) error {
	fmt.Println("Userid ---", userId)
	fmt.Println("Slug ---", slug)
	if err := db.Client.Where("owner_id=? AND slug=?", userId, slug).Delete(&DocumentModel{}).Error; err != nil {
		fmt.Println("Error --- ", err)
		return err
	}
	return nil
}
