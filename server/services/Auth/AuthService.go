package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	redis "github.com/jainam240101/doc-create/db"
	AuthModel "github.com/jainam240101/doc-create/server/domain/Auth"
)

type AuthService interface {
	CreateToken(string, string) (*AuthModel.AuthResponse, error)
	DeleteAuth(string) (int64, error)
}

type DefaultAuthService struct {
	repo AuthModel.AuthRepositoryDb
}

func NewAuthServie(repo AuthModel.AuthRepositoryDb) DefaultAuthService {
	return DefaultAuthService{repo: repo}
}

func (db DefaultAuthService) CreateToken(email string, password string) (*AuthModel.AuthResponse, error) {
	var err error
	u, err := db.repo.CheckCredentials(email, password)
	if err != nil {
		return nil, err
	}

	td := AuthModel.TokenDetails{}
	// td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AtExpires = time.Now().Add(time.Hour * 15).Unix()
	td.AccessUuid = uuid.New().String()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.New().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = u.ID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = u.ID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	err = CreateAuth(u.ID.String(), &td)
	if err != nil {
		return nil, err
	}
	td.UserResponse = *u.ToDto()
	return td.ToDto(), nil

}

func CreateAuth(userid string, td *AuthModel.TokenDetails) error {
	fmt.Println("ACCESS UUID --- ", td.AccessUuid)
	err := redis.RedisClient.Set(redis.Ctx, td.AccessUuid, userid, 0).Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("REFRESH UUID --- ", td.RefreshUuid)
	errRefresh := redis.RedisClient.Set(redis.Ctx, td.RefreshUuid, userid, 0).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func (db DefaultAuthService) DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := redis.RedisClient.Del(redis.Ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
