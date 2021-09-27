package documents

import (
	"fmt"

	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Documents"
	"github.com/jainam240101/doc-create/server/helpers"
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
	d.OwnerId = "1234"
	document, err := dh.Service.CreateDocument(d)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return

	}
	helpers.SendSuccessResponse(c, 200, document)
}
func (dh *DocumentHandlers) ReadDocument(c *gin.Context) {
	data, err := dh.Service.ReadDocument()
	if err != nil {
		fmt.Println(err)
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) FindDocumentByQuery(c *gin.Context) {
	search := c.Request.URL.Query().Get("query")
	if search == "" {
		helpers.SendErrorResponse(c, 406, "No Search Provided")
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
	userId := "231"
	data, err := dh.Service.OwnedDocuments(userId)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) ReadDocumentUsingSlug(c *gin.Context) {
	slug := c.Param("slug")
	data, err := dh.Service.GetDocumentData(slug)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (dh *DocumentHandlers) UpdateDocument(c *gin.Context) {
	slug := c.Param("slug")
	userId := "1234"
	d := domain.DocumentModel{}
	if err := c.ShouldBindJSON(&d); err != nil {
		fmt.Println(err.Error())
		helpers.SendErrorResponse(c, 406, "Body Not proper")
		return
	}
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
	userId := "1234"
	err := dh.Service.DeleteDocument(userId, slug)
	if err != nil {
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, "Deleted")

}
