package routers

import (
	"GinBlog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		//用户模块的路由接口
		v1.POST("user/add", v1.AddUser)
		//分类模块的路由接口
		//文章模块的路由接口
	}

	r.Run(utils.HttpPort)
}
