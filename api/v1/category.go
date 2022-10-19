package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加分类
func AddCategory(ctx *gin.Context) {
	var data model.Category
	_ = ctx.ShouldBindJSON(&data)
	//查询是否已经使用
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		//添加到数据库
		model.InsertCategory(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.Get_Error_Msg(code),
	})
}

//查询分类 --- 列表
func GetCategory(ctx *gin.Context) {
	//接收参数
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	//查询分类
	data := model.GetCategory(pageSize, pageNum)
	//返回
	ctx.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"message": errmsg.Get_Error_Msg(errmsg.SUCCESS),
	})
}

//编辑分类
func UpdateCategory(ctx *gin.Context) {
	//取参数
	id, _ := strconv.Atoi(ctx.Param("id"))

	//解析数据
	var data model.Category
	ctx.ShouldBindJSON(&data)

	//查看分类是否存在
	code := model.CheckCategory(data.Name)
	if code == errmsg.SUCCESS {
		code = model.UpdateCategory(id, &data)
	}
	if code == errmsg.ERROR_CATENAME_USED {
		ctx.Abort()
	}

	//返回
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.Get_Error_Msg(code),
	})
}

//删除分类
func DeleteCategory(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	code := model.DeleteCategory(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"messgae": errmsg.Get_Error_Msg(code),
	})
}
