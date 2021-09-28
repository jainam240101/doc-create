package users

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepositoryDb struct {
	Client *gorm.DB
}

func NewUserRepositoryDb(client *gorm.DB) UserRepositoryDb {
	return UserRepositoryDb{Client: client}
}

func (d UserRepositoryDb) CreateUser(u UserModel) (*UserModel, error) {
	d.Client.Save(&u)
	return &u, nil
}

func (d UserRepositoryDb) FindUserById(id string) (*UserModel, error) {
	var userModel UserModel
	if err := d.Client.Where("id=?", id).First(&userModel).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	// d.Client.First(&userModel, "id = ?", id)
	return &userModel, nil
}

func (d UserRepositoryDb) SearchUser(searchString string) ([]UserModel, error) {
	var userModel []UserModel
	if err := d.Client.Where("LOWER(name) LIKE ? OR username LIKE ?", searchString, searchString).Find(&userModel).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	return userModel, nil
}

func (db UserRepositoryDb) UpdateUser(userId string, updates UserModel) (*UserModel, error) {
	var userModel UserModel
	if err := db.Client.Model(&UserModel{}).Where("id = ?", userId).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := db.Client.Where("id=?", userId).First(&userModel).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	return &userModel, nil
}

func (db UserRepositoryDb) DeleteUser(userId string) error {
	fmt.Println("USER ID --- ", userId)
	if err := db.Client.Where("id=?", userId).Delete(&UserModel{}).Error; err != nil {
		fmt.Println("ERROR  --- ", err)
		return err
	}
	return nil
}

func (db UserRepositoryDb) CreateBookmark(id string, slug string) (*UserModel, error) {
	var user UserModel
	if err := db.Client.Where("id=?", id).First(&user).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	bookmarks := user.Bookmarks
	bookmarks = append(bookmarks, slug)
	user.Bookmarks = bookmarks
	if err := db.Client.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (db UserRepositoryDb) DeleteBookmark(id string, slug string) (*UserModel, error) {
	var user UserModel
	if err := db.Client.Where("id=?", id).First(&user).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	finalData:=[]string{}
	bookmarks := user.Bookmarks
	for i, v := range bookmarks {
    	if v == slug {
        	finalData = append(finalData[:i], finalData[i+1:]...)
        	break
    	}
	}
	user.Bookmarks = finalData
	if err := db.Client.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
