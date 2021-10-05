package routes

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Pages"
	handlers "github.com/jainam240101/doc-create/server/handlers/Page"
	"github.com/jainam240101/doc-create/server/middleware"
	service "github.com/jainam240101/doc-create/server/services/Page"
	"gorm.io/gorm"
)

func PageRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	ph := handlers.PageHandlers{Service: service.NewPageService(domain.NewPageRepositoryDb(db))}
	route := apiRouter.Group("/page")
	{
		route.POST("/create-page", middleware.TokenAuthMiddleware(), ph.CreatePage) //Protected
		route.GET("/getTitles/:documentId", ph.GetTitlesOfADocument)
		route.GET("/:slug", ph.GetDataofPage)
		route.GET("ongoing/:slug", middleware.TokenAuthMiddleware(), ph.EditingPage)
		route.PUT("/:slug", middleware.TokenAuthMiddleware(), ph.UpdatePage)                     //Protected
		route.DELETE("/:slug", middleware.TokenAuthMiddleware(), ph.DeletePage)                  //Protected
		route.POST("change-order/:documentId", middleware.TokenAuthMiddleware(), ph.ChangeOrder) //Protected
		route.POST("/fork", middleware.TokenAuthMiddleware(), ph.ForkPage)                       //Protected
	}
}
