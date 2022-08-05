package rhttp

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/cfg"
	v1 "github.com/iguidao/redis-manager/src/rhttp/v1"
)

// NewServer return a configured http server of gin
func NewServer() *gin.Engine {
	// 存储日志文件代码
	logpath := cfg.Get_Local("logapipath")
	gin.DisableConsoleColor()
	f, _ := os.Create(logpath)
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	// r.LoadHTMLGlob("../website/html/*")
	// r.Static("/static", "../website/static")

	base := r.Group("/redis-manager/base/v1")
	{
		base.GET("/health", v1.HealthCheck)
	}
	cluster := r.Group("/redis-manager/cluster/v1")
	{
		cluster.GET("/list", v1.ClusterList)
		cluster.POST("/add", v1.ClusterAdd)
	}
	// home := r.Group("/")
	// {
	// 	home.GET("", v1.Home)
	// }
	return r
}
