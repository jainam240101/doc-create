package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jainam240101/doc-create/db"
	Routes "github.com/jainam240101/doc-create/server/routes"
)

func main() {
	db.Init()
	router := gin.Default()
	Routes.UserRoutes(&router.RouterGroup, db.DB)
	Routes.AuthRoutes(&router.RouterGroup, db.DB)
	Routes.DocumentRoutes(&router.RouterGroup, db.DB)
	router.Run()

}
