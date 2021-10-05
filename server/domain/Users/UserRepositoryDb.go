package users

import (
	"errors"
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

func (d UserRepositoryDb) FindUserById(username string) (*UserModel, error) {
	var userModel UserModel
	if err := d.Client.Where("username=?", username).First(&userModel).Error; err != nil {
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
	result := db.Client.Model(&UserModel{}).Where("id = ?", userId).Updates(updates)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("you do not have the permission")
	}
	if err := db.Client.Where("id=?", userId).First(&userModel).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	return &userModel, nil
}

func (db UserRepositoryDb) DeleteUser(userId string) error {
	result := db.Client.Where("id=?", userId).Delete(&UserModel{})
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("you do not have the permission")
	}
	return nil
}

