package main

import (
	"go-web/redis"
	"go-web/routers"
	"go-web/sql"
)

func main() {

	_ = sql.ConnectDB()

	_ = redis.ConnectRedis()

	ginRouter := routers.InitGinRouter()
	ginRouter.Run(":8080")

}
