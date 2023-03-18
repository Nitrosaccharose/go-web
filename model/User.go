package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName     string `gorm:"type:varchar(32); not null;" json:"user_name" binding:"required"`
	UserAccount  string `gorm:"type:varchar(16); not null;" json:"user_account" binding:"required"`
	UserPassword string `gorm:"type:varchar(32); not null;" json:"user_password" binding:"required"`
}
