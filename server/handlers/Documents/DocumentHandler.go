package documents

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Documents"
	"github.com/jainam240101/doc-create/server/helpers"
	"github.com/jainam240101/doc-create/server/middleware"
	service "github.com/jainam240101/doc-create/server/services/Documents"
)

type DocumentHandlers struct {
	Service service.DocumentService
}

func (dh *DocumentHandlers) CreateDocument(c *gin.Context) {
	d := domain.DocumentModel{}
	if err := c.ShouldBindJSON(&d); err != nil {
		fmt.Println(err.Error())
		helpers.SendErrorResponse(c, 406, "Body Not proper")
		return
	}
	var err error
	d.OwnerId, err = middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, "Unuthorized")
		return
	}
	document, err := dh.Service.CreateDocument(d)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return

	}
	helpers.SendSuccessResponse(c, 200, document)
}

func (dh *DocumentHandlers) FindDocumentByQuery(c *gin.Context) {
	search := c.Request.URL.Query().Get("query")
	if search == "" {
		helpers.SendSuccessResponse(c, 200, []string{})
		return
	}
	data, err := dh.Service.SearchDocument(search)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) OwnedDocument(c *gin.Context) {
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, "Unuthorized")
		return
	}
	data, err := dh.Service.OwnedDocuments(userId)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) ReadAllProjectsPublishedByUser(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		helpers.SendErrorResponse(c, 406, "No User Provided")
		return
	}
	data, err := dh.Service.ReadAllProjectsPublishedByUser(userId)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) ReadDocumentUsingSlug(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "No Slug Provided")
		return
	}
	data, err := dh.Service.GetDocumentData(slug)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) UpdateDocument(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "No Slug Provided")
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, "Unuthorized")
		return
	}
	d := domain.DocumentModel{}
	if err := c.ShouldBindJSON(&d); err != nil {
		fmt.Println(err.Error())
		helpers.SendErrorResponse(c, 406, "Body Not proper")
		return
	}
	fmt.Println("UserId ------ ", userId)
	data, err := dh.Service.UpdateDocument(slug, userId, d)
	if err != nil {
		fmt.Println(err.Error())
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) DeleteDocument(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "No Slug Provided")
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, "Unuthorized")
		return
	}
	docError := dh.Service.DeleteDocument(userId, slug)
	if docError != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, "Deleted")

}
func (dh *DocumentHandlers) ForkDocument(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "No Slug Provided")
		return
	}
	userId, err := middleware.GetUserId(c)
	if err != nil {
		helpers.SendErrorResponse(c, http.StatusUnauthorized, "Unuthorized")
		return
	}
	data, err := dh.Service.ForkDocument(slug, userId)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)

}
