package auth

import (
	"fmt"

	userModel "github.com/jainam240101/doc-create/server/domain/Users"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepositoryDb struct {
	Client *gorm.DB
}

func NewUserRepositoryDb(client *gorm.DB) AuthRepositoryDb {
	return AuthRepositoryDb{Client: client}
}

func (d AuthRepositoryDb) CheckCredentials(email string, password string) (*userModel.UserModel, error) {
	var u userModel.UserModel
	if err := d.Client.Where("email=?", email).First(&u).Error; err != nil {
		fmt.Println("Error --- ", err.Error())
		return nil, err
	}
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return nil, err
	}
	return &u, nil
}
