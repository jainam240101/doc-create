package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Users"
	"github.com/jainam240101/doc-create/server/helpers"
	"github.com/jainam240101/doc-create/server/middleware"
	service "github.com/jainam240101/doc-create/server/services/Users"
)

type Userhandlers struct {
	Service service.UserService
}

func (uh *Userhandlers) CreateUser(c *gin.Context) {
	u := domain.UserModel{}
	if err := c.BindJSON(&u); err != nil {
		helpers.SendErrorResponse(c, 406, "Body has Parameters missing")
		return
	}
	user, err := uh.Service.CreateUser(u)
	if err != nil {
		helpers.SendErrorResponse(c, 406, "Failed to create User ")
		return
	}
	helpers.SendSuccessResponse(c, 200, user)
}

func (uh *Userhandlers) FindUserById(c *gin.Context) {
	username := c.Request.URL.Query().Get("username")
	if username == "" {
		helpers.SendSuccessResponse(c, 200, []string{})
		return
	}
	user, err := uh.Service.FindUserByUsername(username)
	if err != nil {
		helpers.SendErrorResponse(c, 404, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, user)
}

func (uh *Userhandlers) FindUserByQuery(c *gin.Context) {
	search := c.Request.URL.Query().Get("search")
	if search == "" {
		helpers.SendSuccessResponse(c, 200, []string{})
		return
	}
	data, err := uh.Service.SearchUser(search)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}
func (uh *Userhandlers) UpdateUser(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	u := domain.UserModel{}
	if err := c.BindJSON(&u); err != nil {
		helpers.SendErrorResponse(c, 406, "Body has Parameters missing")
		return
	}
	data, err := uh.Service.UpdateUser(u, userId)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (uh *Userhandlers) DeleteUser(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	deleteErr := uh.Service.DeleteUser(userId)
	if deleteErr != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, deleteErr.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, "User Deleted")
}

//For using Bookmarks
type body struct {
	ProjectId string `json:"projectId"`
}

func (uh *Userhandlers) CreateBookmark(c *gin.Context) {
	var b body
	if err := c.BindJSON(&b); err != nil {
		helpers.SendErrorResponse(c, 406, "Body has Parameters missing")
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	data, err := uh.Service.CreateBookmark(userId, b.ProjectId)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}
func (uh *Userhandlers) DeleteBookmark(c *gin.Context) {
	var b body
	if err := c.BindJSON(&b); err != nil {
		helpers.SendErrorResponse(c, 406, "Body has Parameters missing")
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	data, err := uh.Service.DeleteBookmark(userId, b.ProjectId)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}
