package routers

import (
	v1 "GinBlog/api/v1"
	"GinBlog/middleware"
	"GinBlog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	engine := gin.New()

	//注册中间件
	engine.Use(gin.Recovery())
	engine.Use(middleware.Logger())

	auth := engine.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		//用户模块的路由接口
		auth.PUT("user/:id", v1.UpdateUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.UpdateCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块的路由接口
		auth.POST("article/add", v1.AddArt)
		auth.PUT("article/:id", v1.UpdateArt)
		auth.DELETE("article/:id", v1.DeleteArt)
	}
	public := engine.Group("api/v1")
	{
		public.POST("user/add", v1.AddCategory)
		public.GET("user/GetUsers", v1.GetUsers)
		public.GET("category", v1.GetCategory)
		public.GET("article/GetArts", v1.GetArts)
		public.GET("article/GetArt", v1.GetArt)
		public.GET("article/category/:id", v1.Get_Cat_Arts)
		public.POST("login", v1.Login)
	}
	engine.Run(utils.HttpPort)
}
