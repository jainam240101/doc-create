package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	redis "github.com/jainam240101/doc-create/db"
	"github.com/jainam240101/doc-create/server/helpers"
)

type AccessDetails struct {
	AccessUuid string
	UserId     string
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ExtractTokenMetadata(ctoken string) (*AccessDetails, error) {
	token, err := VerifyToken(ctoken)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId := fmt.Sprintf("%v", claims["user_id"])
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func FetchAuth(authD *AccessDetails) (string, error) {
	userid, err := redis.RedisClient.Get(redis.Ctx, authD.AccessUuid).Result()
	if err != nil {
		return "", err
	}
	return userid, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r.Header["Authorization"][0])
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			helpers.SendErrorResponse(c, http.StatusUnauthorized, gin.H{
				"error": "JWT Expired",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

func GetUserId(c *gin.Context) (string, error) {
	tokenAuth, err := ExtractTokenMetadata(c.Request.Header["Authorization"][0])
	if err != nil {
		return "", err
	}
	userId, err := FetchAuth(tokenAuth)
	if err != nil {
		return "", err
	}
	return userId, nil
}
