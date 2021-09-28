package users

import (
	"fmt"

	"github.com/google/uuid"
	users "github.com/jainam240101/doc-create/server/domain/Users"
	"github.com/jainam240101/doc-create/server/dto"
	"golang.org/x/crypto/bcrypt"
)

type DefaultUserService struct {
	repo users.UserRepositoryDb
}

type UserService interface {
	CreateUser(users.UserModel) (*dto.UserResponse, error)
	FindUserById(string) (*dto.UserResponse, error)
	SearchUser(string) ([]dto.UserResponse, error)
	UpdateUser(users.UserModel, string) (*dto.UserResponse, error)
	DeleteUser(string) error
	CreateBookmark(string, string) (*dto.UserResponse, error)
	DeleteBookmark(string, string) (*dto.UserResponse, error)
}

func NewCustomerService(repository users.UserRepositoryDb) DefaultUserService {
	return DefaultUserService{repo: repository}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (db DefaultUserService) CreateUser(u users.UserModel) (*dto.UserResponse, error) {
	pass, err := HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = pass
	u.ID = uuid.New()
	user, err := db.repo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	return user.ToDto(), nil
}

func (db DefaultUserService) FindUserById(id string) (*dto.UserResponse, error) {
	user, err := db.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}
	return user.ToDto(), nil
}

func (db DefaultUserService) SearchUser(searchString string) ([]dto.UserResponse, error) {
	data, err := db.repo.SearchUser(string("%" + searchString + "%"))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	finalData := []dto.UserResponse{}
	for _, values := range data {
		finalData = append(finalData, *values.ToDto())
	}
	return finalData, nil
}

func (db DefaultUserService) UpdateUser(u users.UserModel, userid string) (*dto.UserResponse, error) {
	if u.Password != "" {
		pass, err := HashPassword(u.Password)
		u.Password = pass
		if err != nil {
			return nil, err
		}
	}
	dbValue, err := db.repo.UpdateUser(userid, u)
	if err != nil {
		return nil, err
	}
	return dbValue.ToDto(), nil
}

func (db DefaultUserService) DeleteUser(userId string) error {
	err := db.repo.DeleteUser(userId)
	if err != nil {
		return err
	}
	return nil
}

func (db DefaultUserService) CreateBookmark(userId string, slug string) (*dto.UserResponse, error) {
	data, err := db.repo.CreateBookmark(userId, slug)
	if err != nil {
		return nil, err
	}
	return data.ToDto(), nil
}
func (db DefaultUserService) DeleteBookmark(userId string, slug string) (*dto.UserResponse, error) {
	data, err := db.repo.DeleteBookmark(userId, slug)
	if err != nil {
		return nil, err
	}
	return data.ToDto(), nil
}
