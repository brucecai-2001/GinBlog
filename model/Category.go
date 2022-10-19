package model

import (
	"GinBlog/utils/errmsg"
	"fmt"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20); not null" json:"name"`
}

//新增分类
func InsertCategory(cate *Category) int {
	err := db.Create(cate)
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询分类是否存在
func CheckCategory(Name string) int {

	var cate Category
	db.Select("id").Where("name = ?", Name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS

}

//查询分类列表，使用了分页，方便前端生成页面
func GetCategory(pageSize int, pageNum int) []Category {
	var cates []Category
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cates).Error
	if err != nil {
		return nil
	}
	return cates
}

//更新分类
func UpdateCategory(id int, data *Category) int {
	var cate Category
	ToUpdate := make(map[string]interface{})
	ToUpdate["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Update(ToUpdate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//删除分类
func DeleteCategory(id int) int {
	var cate Category
	err := db.Where("id = ?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
