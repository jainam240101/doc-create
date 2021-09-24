package users

import (
	"github.com/gin-gonic/gin"
	domain "github.com/jainam240101/doc-create/server/domain/Auth"
	handlers "github.com/jainam240101/doc-create/server/handlers/Auth"
	service "github.com/jainam240101/doc-create/server/services/Auth"
	"gorm.io/gorm"
)

func AuthRoutes(apiRouter *gin.RouterGroup, db *gorm.DB) {
	rh := handlers.Authhandlers{Service: service.NewAuthServie(domain.NewUserRepositoryDb(db))}
	route := apiRouter.Group("/auth")
	{
		route.POST("/login", rh.Login)
		route.POST("/logout", rh.Logout)
	}
}
