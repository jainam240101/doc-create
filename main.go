package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jainam240101/doc-create/db"
	Routes "github.com/jainam240101/doc-create/server/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	godotenv.Load(".env")
	db.Init()
	router := gin.Default()
	Routes.UserRoutes(&router.RouterGroup, db.DB)
	Routes.AuthRoutes(&router.RouterGroup, db.DB)
	Routes.DocumentRoutes(&router.RouterGroup, db.DB)
	Routes.PageRoutes(&router.RouterGroup, db.DB)
	router.Run()

}
