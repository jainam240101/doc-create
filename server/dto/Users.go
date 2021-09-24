package dto

import "github.com/google/uuid"

type UserResponse struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	ProfilePic string    `json:"profilePic"`
}
