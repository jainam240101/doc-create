package routes

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Pages"
	handlers "github.com/jainam240101/doc-create/server/handlers/Page"
	service "github.com/jainam240101/doc-create/server/services/Page"
	"gorm.io/gorm"
)

func PageRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	ph := handlers.PageHandlers{Service: service.NewPageService(domain.NewPageRepositoryDb(db))}
	route := apiRouter.Group("/page")
	{
		route.POST("/create-page", ph.CreatePage) //Protected
		route.GET("/getTitles/:documentId", ph.GetTitlesOfADocument)
		route.GET("/:slug", ph.GetDataofPage)
		route.PUT("/:slug", ph.UpdatePage)                     //Protected
		route.DELETE("/:slug", ph.DeletePage)                  //Protected
		route.POST("change-order/:documentId", ph.ChangeOrder) //Protected
		route.POST("/fork", ph.ForkPage)                       //Protected
	}
}
