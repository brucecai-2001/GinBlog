package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	log_path := "log/blog.log"
	src, err := os.Open(log_path)
	if err != nil {
		fmt.Println(err)
	}
	logger := logrus.New()

	logger.Out = src
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()
		stopTime := time.Since(startTime)
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := ctx.Writer.Status()
		clientIP := ctx.ClientIP()
		//客户端信息
		userAgent := ctx.Request.UserAgent()
		datasize := ctx.Writer.Size()
		method := ctx.Request.Method
		path := ctx.Request.RequestURI

		entry := logger.WithFields(logrus.Fields{
			"host_name": hostName,
			"status":    statusCode,
			"time_cost": stopTime,
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
			"data_size": datasize,
			"agent":     userAgent,
		})

		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Info()
		}
	}
}
