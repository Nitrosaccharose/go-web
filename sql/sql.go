package sql

import (
	"go-web/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	// ------ gorm初始化 ------
	dsn := "root:1234@tcp(127.0.0.1:3306)/ginsql?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	setAutoMigrate(db)

	return db
}

func setAutoMigrate(db *gorm.DB) {
	_ = db.AutoMigrate(&model.User{})
}

func GetDB() *gorm.DB {
	return db
}
