package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-web/model"
	"go-web/redis"
	"net/http"
	"time"
)

func GETKVTByKey(c *gin.Context) {
	redisClient := redis.GetRedis()
	key := c.Query("key")
	value, err := redisClient.Get(context.Background(), key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Redis Get Error",
			"data": gin.H{},
			"code": http.StatusInternalServerError,
		})
		return
	}
	ttl, err := redisClient.TTL(context.Background(), key).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "Redis TTL Error",
			"data": gin.H{},
			"code": http.StatusInternalServerError,
		})
		return
	}
	kv := model.KVObject{
		Key:   key,
		Value: value,
		Time:  int(ttl.Seconds()),
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"data": kv,
		"code": http.StatusOK,
	})
}

func POSTKVT(c *gin.Context) {
	var datakv model.KVObject
	err := c.ShouldBindJSON(&datakv)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "Failure",
			"data": gin.H{},
			"code": http.StatusBadRequest,
		})
	} else {
		redisClient := redis.GetRedis()
		err := redisClient.Set(context.Background(), datakv.Key, datakv.Value, time.Duration(datakv.Time)*time.Second).Err()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg":  "Redis Set Error",
				"data": gin.H{},
				"code": http.StatusInternalServerError,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Success",
				"data": gin.H{"key": datakv.Key, "value": datakv.Value, "time": datakv.Time},
				"code": http.StatusOK,
			})
		}
	}
}
