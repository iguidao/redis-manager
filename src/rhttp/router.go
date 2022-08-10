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
	history := r.Group("/redis-manager/ophistory/v1")
	{
		history.GET("/list", v1.OpHistory)
	}
	cli := r.Group("/redis-manager/cli/v1")
	{
		cli.GET("/querykey", v1.QueryKey) //查key
		cli.GET("/bigkey", v1.BigKey)     //大key
		cli.GET("/hotkey", v1.HotKey)     //热key
		cli.GET("/allkey", v1.AllKey)     //所有key
		cli.GET("/slowkey", v1.SlowKey)   //慢key
	}
	codis := r.Group("/redis-manager/codis/v1")
	{
		codis.GET("/list", v1.CodisList)
		codis.GET("/group", v1.CodisGroup)
	}
	// home := r.Group("/")
	// {
	// 	home.GET("", v1.Home)
	// }
	return r
}
