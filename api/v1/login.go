package v1

import (
	"GinBlog/middleware"
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"net/http"

	"github.com/gin-gonic/gin"
)

//登录
func Login(ctx *gin.Context) {
	var data model.User
	var token string
	ctx.ShouldBindJSON(&data)
	code := model.CheckLogin(data.UserName, data.PassWord)
	if code == errmsg.SUCCESS {
		//登录成功生成Token
		token, _ = middleware.SetToken(data.UserName, data.PassWord)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.Get_Error_Msg(code),
		"token":   token,
	})
}
