package model

import (
	"GinBlog/utils/errmsg"
	"fmt"

	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title    string   `gorm:"type:varchar(100); not null" json:"title"`
	Category Category `gorm:"foreginkey:Cid"`
	//对应的category id
	Cid     int    `gorm:"type:int; not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200); not null" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

//新增文章
func InsertArticle(article *Article) int {
	err := db.Create(article)
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个文章
func GetArt(id int) (Article, int) {
	var art Article
	err := db.Preload("Category").Where("id = ?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR
	}
	return art, errmsg.SUCCESS
}

//查询分类下所有文章
func Get_Cat_Arts(id int, pageSize int, pageNum int) ([]Article, int) {
	var cate_art_list []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("id = ?", id).Find(&cate_art_list).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXISTED
	}
	return cate_art_list, errmsg.SUCCESS
}

//查询所有文章，使用了分页，方便前端生成页面
func GetArts(pageSize int, pageNum int) ([]Article, int) {
	var lists []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&lists).Error
	if err != nil {
		return nil, errmsg.ERROR
	}
	return lists, errmsg.SUCCESS
}

//更新文章
func UpdateArt(id int, data *Article) int {
	var art Article
	ToUpdate := make(map[string]interface{})
	ToUpdate["Title"] = data.Title
	ToUpdate["Cid"] = data.Cid
	ToUpdate["Desc"] = data.Desc
	ToUpdate["Content"] = data.Content
	ToUpdate["Img"] = data.Img
	err := db.Model(&art).Where("id = ?", id).Update(ToUpdate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除文章
func DeleteArt(id int) int {
	var art Article
	err := db.Where("id = ?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
