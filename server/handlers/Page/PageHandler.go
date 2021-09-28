package page

import (
	"fmt"

	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Pages"
	"github.com/jainam240101/doc-create/server/helpers"
	service "github.com/jainam240101/doc-create/server/services/Page"
)

type PageHandlers struct {
	Service service.PageService
}

func (ph *PageHandlers) CreatePage(c *gin.Context) {
	p := domain.PageModel{}
	if err := c.BindJSON(&p); err != nil {
		helpers.SendErrorResponse(c, 406, "Body has parameters missing")
		return
	}
	p.OwnerId = "1234"
	page, err := ph.Service.CreatePage(p)
	if err != nil {
		fmt.Println("Error --- ", err)
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, page)
}

func (ph *PageHandlers) GetTitlesOfADocument(c *gin.Context) {
	documentId := c.Param("documentId")
	if documentId == "" {
		helpers.SendErrorResponse(c, 406, "Document Id not mentioned")
		return
	}
	data, err := ph.Service.GetTitlesofDocument(documentId)
	if err != nil {
		fmt.Println("Error --- ", err)
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (ph *PageHandlers) GetDataofPage(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "Slug Id not mentioned")
		return
	}
	data, err := ph.Service.GetDataofPage(slug)
	if err != nil {
		fmt.Println("Error --- ", err)
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (ph *PageHandlers) UpdatePage(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "Slug Id not mentioned")
		return
	}
	p := domain.PageModel{}
	if err := c.BindJSON(&p); err != nil {
		helpers.SendErrorResponse(c, 406, "Body has parameters missing")
		return
	}
	data, err := ph.Service.UpdatePage(slug, p)
	if err != nil {
		fmt.Println("Error --- ", err)
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, data)
}

func (ph *PageHandlers) DeletePage(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		helpers.SendErrorResponse(c, 406, "Slug Id not mentioned")
		return
	}
	err := ph.Service.DeletePage(slug)
	if err != nil {
		fmt.Println("Error --- ", err)
		helpers.SendErrorResponse(c, 406, err.Error())
		return
	}
	helpers.SendSuccessResponse(c, 200, "Deleted")
}
