package auth

import (
		dto "github.com/jainam240101/doc-create/server/dto"
		// users "github.com/jainam240101/doc-create/server/domain/users"
)

type TokenDetails struct {
	dto.UserResponse
	AccessToken  string 
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AuthResponse struct {
	dto.UserResponse
	AccessToken  string `json:"access_token"` 
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func(a *TokenDetails)ToDto() *AuthResponse{
	return &AuthResponse{
		UserResponse: a.UserResponse,
		AccessToken: a.AccessToken,
		RefreshToken: a.RefreshToken,
	}
}

type AuthRepository interface {
	CheckCredentials(email string,password string) (*TokenDetails, error)
}
