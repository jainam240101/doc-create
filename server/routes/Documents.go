package routes

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Documents"
	handlers "github.com/jainam240101/doc-create/server/handlers/Documents"
	"github.com/jainam240101/doc-create/server/middleware"
	service "github.com/jainam240101/doc-create/server/services/Documents"
	"gorm.io/gorm"
)

func DocumentRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	dh := handlers.DocumentHandlers{Service: service.NewDocumentService(domain.NewRepositoryDb(db))}
	route := apiRouter.Group("/documents")
	{
		route.POST("/create-document", middleware.TokenAuthMiddleware(), dh.CreateDocument) //Protected
		route.GET("/search", dh.FindDocumentByQuery)
		route.GET("/published-documents/:userId", dh.ReadAllProjectsPublishedByUser)
		route.GET("/my-documents", middleware.TokenAuthMiddleware(), dh.OwnedDocument) //Protected
		route.GET("/document/:slug", dh.ReadDocumentUsingSlug)
		route.PUT("/:slug", middleware.TokenAuthMiddleware(), dh.UpdateDocument)     //Protected
		route.DELETE("/:slug", middleware.TokenAuthMiddleware(), dh.DeleteDocument)  //Protected
		route.POST("/fork/:slug", middleware.TokenAuthMiddleware(), dh.ForkDocument) //Protected
	}
}
