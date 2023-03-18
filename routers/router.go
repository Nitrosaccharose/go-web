package routers

import (
	"github.com/gin-gonic/gin"
	"go-web/controller"
)

func InitGinRouter() *gin.Engine {
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

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

	router.GET("/api/auth/kv/get/key", controller.GETKVTByKey)
	router.POST("/api/auth/kv/post", controller.POSTKVT)
}
