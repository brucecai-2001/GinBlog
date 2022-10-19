package v1

import (
	"GinBlog/model"
	"GinBlog/utils/errmsg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//添加文章
func AddArt(ctx *gin.Context) {
	var data model.Article
	_ = ctx.ShouldBindJSON(&data)

	//添加到数据库
	code := model.InsertArticle(&data)

	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.Get_Error_Msg(code),
	})
}

//查询文章 --- 单个
func GetArt(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, code := model.GetArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   data,
		"msg":    errmsg.Get_Error_Msg(code),
	})

}

//查询所有文章列表
func GetArts(ctx *gin.Context) {
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
	data, code := model.GetArts(pageSize, pageNum)
	//返回
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.Get_Error_Msg(errmsg.SUCCESS),
	})
}

//查询单个目录下的文章
func Get_Cat_Arts(ctx *gin.Context) {
	//接收参数
	pageSize, _ := strconv.Atoi(ctx.Query("pagesize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pagenum"))
	id, _ := strconv.Atoi(ctx.Param("id"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}

	data, code := model.Get_Cat_Arts(id, pageSize, pageNum)

	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.Get_Error_Msg(errmsg.SUCCESS),
	})
}

//编辑文章
func UpdateArt(ctx *gin.Context) {
	//取参数
	id, _ := strconv.Atoi(ctx.Param("id"))

	//解析数据
	var data model.Article
	ctx.ShouldBindJSON(&data)

	code := model.UpdateArt(id, &data)

	//返回
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.Get_Error_Msg(code),
	})
}

//删除文章
func DeleteArt(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	code := model.DeleteArt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"status":  code,
		"messgae": errmsg.Get_Error_Msg(code),
	})
}
