package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type User struct {
	gorm.Model
	UserName     string
	UserPassword string
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/ginsql?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	_ = db.AutoMigrate(&User{})

	router := gin.Default()

	router.GET("/hello", func(context *gin.Context) {
		context.String(200, "Hello World!")
	})
	_ = router.Run(":8080")
}
