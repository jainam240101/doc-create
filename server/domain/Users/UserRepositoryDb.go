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
	if err := d.Client.Where("name LIKE ? OR username LIKE ?", searchString, searchString).Find(&userModel).Error; err != nil {
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
