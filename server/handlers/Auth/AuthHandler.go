package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	User "github.com/jainam240101/doc-create/server/domain/Auth"
	"github.com/jainam240101/doc-create/server/helpers"
	"github.com/jainam240101/doc-create/server/middleware"
	AuthService "github.com/jainam240101/doc-create/server/services/Auth"
)

type Authhandlers struct {
	Service AuthService.AuthService
}

func (uh *Authhandlers) Login(c *gin.Context) {
	var u User.User
	if err := c.ShouldBindJSON(&u); err != nil {
		helpers.SendErrorResponse(c, 406, "Invalid JSON")
		return
	}
	ts, err := uh.Service.CreateToken(u.Email, u.Password)
	if err != nil {
		helpers.SendErrorResponse(c, 406, "Error in Creating Token")
		return
	}
	// saveErr := AuthService.CreateAuth("a92dheu1231", ts)
	// if saveErr != nil {
	// 	helpers.SendErrorResponse(c, 406, saveErr.Error())
	// 	return
	// }
	helpers.SendSuccessResponse(c, 200, ts)
}

func (uh *Authhandlers) Logout(c *gin.Context) {
	au, err := middleware.ExtractTokenMetadata(c.Request.Header["Authorization"][0])
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized1")
		return
	}
	deleted, delErr := uh.Service.DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, "Unauthroized")
		return
	}
	helpers.SendSuccessResponse(c, 200, "Successfully Logged out")
}
