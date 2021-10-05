package routes

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Users"
	handlers "github.com/jainam240101/doc-create/server/handlers/Users"
	"github.com/jainam240101/doc-create/server/middleware"
	service "github.com/jainam240101/doc-create/server/services/Users"
	"gorm.io/gorm"
)

func UserRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	uh := handlers.Userhandlers{Service: service.NewCustomerService(domain.NewUserRepositoryDb(db))}
	route := apiRouter.Group("/users")
	{
		route.GET("/search", uh.FindUserByQuery)
		route.GET("/", uh.FindUserById)
		route.POST("/", uh.CreateUser)
		route.DELETE("/", middleware.TokenAuthMiddleware(), uh.DeleteUser)
		route.PUT("/", middleware.TokenAuthMiddleware(), uh.UpdateUser)
		route.POST("/create-bookmark", middleware.TokenAuthMiddleware(), uh.CreateBookmark)
		route.DELETE("/delete-bookmark", middleware.TokenAuthMiddleware(), uh.DeleteBookmark)
	}
}
