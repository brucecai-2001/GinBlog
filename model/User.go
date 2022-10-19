package model

import (
	"GinBlog/utils/errmsg"
	"encoding/base64"
	"fmt"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20); not null" json:"username"`
	PassWord string `gorm:"type:varchar(20); not null" json:"password"`
	Role     int    `gorm:"type:int " json:"role"`
}

//新增用户
func InsertUser(data *User) int {
	//加密用户密码
	data.PassWord = ScryptPw(data.PassWord)
	err := db.Create(data)
	if err != nil {
		fmt.Println(err)
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询用户是否存在
func CheckUser(username string) int {

	var user User
	db.Select("id").Where("user_name = ?", username).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_UserName_Used
	}
	return errmsg.SUCCESS

}

//查询用户列表，使用了分页，方便前端生成页面
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err := db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
	if err != nil {
		return nil
	}
	return users
}

//更新用户
func UpdateUser(id int, data *User) int {
	var user User
	ToUpdate := make(map[string]interface{})
	ToUpdate["user_name"] = data.UserName
	ToUpdate["role"] = data.Role
	err := db.Model(&user).Where("id = ?", id).Update(ToUpdate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 34, 23, 46, 54, 32, 13, 13}
	HashedPw, _ := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	fpw := base64.StdEncoding.EncodeToString(HashedPw)
	return fpw
}

//删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//登录验证
func CheckLogin(username, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_User_Not_Existed
	}
	if ScryptPw(password) != user.PassWord {
		return errmsg.ERROR_Password_WRONG
	}
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_PERMISSION
	}
	return errmsg.SUCCESS
}
