package users

import (
	"github.com/google/uuid"
	dto "github.com/jainam240101/doc-create/server/dto"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	ID         uuid.UUID      `gorm:"type:char(36);primary_key"`
	Name       string         `json:"name"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Username   string         `json:"username"`
	ProfilePic string         `json:"profilePic"`
	Bookmarks  pq.StringArray `gorm:"type:text[]" json:"bookmarks"`
}

func (c UserModel) ToDto() *dto.UserResponse {
	return &dto.UserResponse{
		ID:         c.ID,
		Name:       c.Name,
		Email:      c.Email,
		Username:   c.Username,
		ProfilePic: c.ProfilePic,
		Bookmarks:  c.Bookmarks,
	}
}

type UserRepository interface {
	CreateUser(UserModel) (*UserModel, error)
	FindUserById(string) (*UserModel, error)
	SearchUser(string) ([]UserModel, error)
	UpdateUser(string, UserModel) (*UserModel, error)
	DeleteUser(string) error
	CreateBookmark(string, string) (*UserModel, error)
	DeleteBookmark(string, string) (*UserModel, error)
}
