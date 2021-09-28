package db

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	documents "github.com/jainam240101/doc-create/server/domain/Documents"
	pages "github.com/jainam240101/doc-create/server/domain/Pages"
	users "github.com/jainam240101/doc-create/server/domain/Users"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var RedisClient *redis.Client

var Ctx = context.Background()

func Init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	var err error
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=Jainam dbname=docCreate port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		fmt.Println("Database Error ---  ", err)
	}
	DB.AutoMigrate(&users.UserModel{}, &documents.DocumentModel{}, &pages.PageModel{})
}
