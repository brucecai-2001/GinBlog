package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	//查询用户名是否已经使用
	code := model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		//添加到数据库
		model.InsertUser(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.Get_Error_Msg(code),
	})
}

//查询用户 --- 列表
func GetUsers(ctx *gin.Context) {
	//接收参数
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	//查询用户
	data := model.GetUsers(pageSize, pageNum)
	//返回
	ctx.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"message": errmsg.Get_Error_Msg(errmsg.SUCCESS),
	})
}

//编辑用户
func UpdateUser(ctx *gin.Context) {
	//取参数
	id, _ := strconv.Atoi(ctx.Param("id"))

	//解析数据
	var data model.User
	ctx.ShouldBindJSON(&data)

	//查看用户是否存在
	code := model.CheckUser(data.UserName)
	if code == errmsg.SUCCESS {
		code = model.UpdateUser(id, &data)
	}
	if code == errmsg.ERROR_UserName_Used {
		ctx.Abort()
	}

	//返回
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.Get_Error_Msg(code),
	})
}

//删除用户
func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	code := model.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"messgae": errmsg.Get_Error_Msg(code),
	})
}
