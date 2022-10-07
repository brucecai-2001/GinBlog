package routers

import (
	"GinBlog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("hello", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	r.Run(utils.HttpPort)
}
