package users

import (
	"github.com/google/uuid"
	dto "github.com/jainam240101/doc-create/server/dto"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID         uuid.UUID `gorm:"type:char(36);primary_key"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Username   string    `json:"username"`
	ProfilePic string    `json:"profilePic"`
}

func (c UserModel) ToDto() *dto.UserResponse {
	return &dto.UserResponse{
		ID:         c.ID,
		Name:       c.Name,
		Email:      c.Email,
		Username:   c.Username,
		ProfilePic: c.ProfilePic,
	}
}

type UserRepository interface {
	CreateUser(UserModel) (*UserModel, error)
	FindUserById(string) (*UserModel, error)
	SearchUser(string) ([]UserModel, error)
	UpdateUser(string, UserModel) (*UserModel, error)
	DeleteUser(string) error
}
