package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jainam240101/doc-create/db"
	Routes "github.com/jainam240101/doc-create/server/routes"
	"github.com/joho/godotenv"
	// cors "github.com/rs/cors/wrapper/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.Next()
			fmt.Println("Inside Options")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	godotenv.Load()
	godotenv.Load(".env")
	db.Init()
	router := gin.Default()
	router.Use(CORSMiddleware())
	Routes.UserRoutes(&router.RouterGroup, db.DB)
	Routes.AuthRoutes(&router.RouterGroup, db.DB)
	Routes.DocumentRoutes(&router.RouterGroup, db.DB)
	Routes.PageRoutes(&router.RouterGroup, db.DB)
	router.Run()
}
