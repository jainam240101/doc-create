package routes

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Documents"
	handlers "github.com/jainam240101/doc-create/server/handlers/Documents"
	service "github.com/jainam240101/doc-create/server/services/Documents"
	"gorm.io/gorm"
)

func DocumentRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	dh := handlers.DocumentHandlers{Service: service.NewDocumentService(domain.NewRepositoryDb(db))}
	route := apiRouter.Group("/documents")
	{
		route.POST("/create-document", dh.CreateDocument)
		route.GET("/get-document", dh.ReadDocument)
		route.GET("/search", dh.FindDocumentByQuery)
		route.GET("/my-documents", dh.OwnedDocument)
	}
}
