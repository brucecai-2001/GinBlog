package model

import (
	"GinBlog/utils"
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

func InitDB() {
	//连接数据库
	db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		"password",
		utils.DbHost,
		utils.Dbport,
		utils.DbName,
	))
	if err != nil {
		fmt.Print("连接数据库失败", err)
	}

	//设置连接池最大闲置连接数
	db.DB().SetMaxIdleConns(10)
	//设置数据库最大连接数
	db.DB().SetMaxOpenConns(10)
	//设置连接最大复用时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//模型迁移
	db.SingularTable(true)
	db.AutoMigrate(&User{}, &Category{}, &Article{})
}
