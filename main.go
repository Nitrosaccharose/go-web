package main

import (
	"go-web/model"
	"go-web/routers"
	"go-web/sql"
)

func main() {

	db := sql.ConnectDB()

	// ------ 自动迁移model ------
	_ = db.AutoMigrate(&model.User{})

	ginRouter := routers.InitGinRouter()
	_ = ginRouter.Run(":8080")

}
