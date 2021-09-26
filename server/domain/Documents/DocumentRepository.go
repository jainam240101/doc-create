package documents

import "gorm.io/gorm"

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
