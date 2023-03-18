package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/controller"
)

func InitGinRouter() *gin.Engine {
	router := gin.Default()
	CollectRouter(router)
	return router
}
func CollectRouter(router *gin.Engine) {
	router.GET("/api/auth/user/get/all", controller.GETUserAll)
	router.GET("/api/auth/user/get/username", controller.GETUserByUserName)
	router.GET("/api/auth/user/get/useraccount", controller.GETUserByUserAccount)
	router.POST("/api/auth/user/post", controller.POSTUser)
	router.PUT("/api/auth/user/put", controller.PUTUser)
	router.DELETE("/api/auth/user/delete", controller.DELETEUser)
}
