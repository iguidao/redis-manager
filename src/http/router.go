package http

import (
	v1 "github.com/iguidao/redis-web-manager/src/http/v1"

	"github.com/gin-gonic/gin"
)

// NewServer return a configured http server of gin
func NewServer() *gin.Engine {
	// 存储日志文件代码
	gin.DisableConsoleColor()
	// f, _ := os.Create("./logs/app.log")
	// gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	r.LoadHTMLGlob("../website/html/*")
	r.Static("/static", "../website/static")

	base := r.Group("/redis-web-manager/base/v1")
	{
		base.GET("/health", v1.HealthCheck)
	}
	redis := r.Group("/redis-web-manager/v1/cluster")
	{
		redis.GET("/list", v1.ClusterList)
	}
	home := r.Group("/")
	{
		home.GET("", v1.Home)
	}
	return r
}
