package main

import (
	"GinBlog/model"
	"GinBlog/routers"
	_ "GinBlog/utils"
)

func main() {
	//初始化数据库
	model.InitDB()
	//初始化路由，启动服务器
	routers.InitRouter()
}
