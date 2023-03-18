package controller

import (
	"github.com/gin-gonic/gin"
	"go-web/model"
	"go-web/sql"
	"net/http"
)

func GETUserAll(c *gin.Context) {
	db := sql.GetDB()
	var users []model.User
	db.Find(&users)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "Success",
		"data": users,
		"code": http.StatusOK,
	})
}

func GETUserByUserName(c *gin.Context) {
	userName := c.Query("user_name")
	db := sql.GetDB()
	var dataListUser []model.User
	db.Where("user_name = ?", userName).Find(&dataListUser)
	if len(dataListUser) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "No found",
			"data": []string{},
			"code": http.StatusBadRequest,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Success",
			"data": dataListUser,
			"code": http.StatusOK,
		})
	}
}

func GETUserByUserAccount(c *gin.Context) {
	userAccount := c.Query("user_account")
	db := sql.GetDB()
	var dataListUser []model.User
	db.Where("user_account = ?", userAccount).Find(&dataListUser)

	if len(dataListUser) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "No found",
			"data": []string{},
			"code": http.StatusBadRequest,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Success",
			"data": dataListUser,
			"code": http.StatusOK,
		})
	}
}

func POSTUser(c *gin.Context) {
	var dataUser []model.User
	err := c.ShouldBindJSON(&dataUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "Failure",
			"data": gin.H{},
			"code": http.StatusBadRequest,
		})
	} else {
		db := sql.GetDB()
		db.Create(&dataUser)
		c.JSON(http.StatusOK, gin.H{
			"msg":  "Success",
			"data": dataUser,
			"code": http.StatusOK,
		})
	}
}

func PUTUser(c *gin.Context) {
	var updateUser model.User
	err := c.ShouldBindJSON(&updateUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "Failure",
			"data": gin.H{},
			"code": http.StatusBadRequest,
		})
	} else {
		db := sql.GetDB()
		var user model.User
		db.Where("id = ?", updateUser.ID).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"msg":  "User not found",
				"data": gin.H{},
				"code": http.StatusNotFound,
			})
		} else {
			db.Model(&user).Updates(&updateUser)
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Success",
				"data": user,
				"code": http.StatusOK,
			})
		}
	}
}

func DELETEUser(c *gin.Context) {
	var deleteUser model.User
	err := c.ShouldBindJSON(&deleteUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "Failure",
			"data": gin.H{},
			"code": http.StatusBadRequest,
		})
	} else {
		db := sql.GetDB()
		var user model.User
		db.Where("id = ?", deleteUser.ID).First(&user)
		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"msg":  "User not found",
				"data": gin.H{},
				"code": http.StatusNotFound,
			})
		} else {
			db.Delete(&user)
			c.JSON(http.StatusOK, gin.H{
				"msg":  "Success",
				"data": user,
				"code": http.StatusOK,
			})
		}
	}
}
