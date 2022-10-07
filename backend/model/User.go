package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20); not null" json:"username"`
	PassWord string `gorm:"type:varchar(20); not null" json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}
