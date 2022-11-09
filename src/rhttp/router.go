package rhttp

import (
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/jwt"
	v1 "github.com/iguidao/redis-manager/src/rhttp/v1"
)

// NewServer return a configured http server of gin
func NewServer() *gin.Engine {
	// 存储日志文件代码
	logpath := cfg.Get_Info_String("logapipath")
	gin.DisableConsoleColor()
	f, _ := os.Create(logpath)
	gin.DefaultWriter = io.MultiWriter(f)
	r := gin.Default()

	// 跨域信息
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTION"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Static("/static", "./website")
	r.LoadHTMLFiles("./website/index.html")
	home := r.Group("")
	{
		home.GET("/", v1.Home)
	}

	base := r.Group("/redis-manager/base/v1")
	{
		base.GET("/health", v1.HealthCheck)
	}
	user := r.Group("/redis-manager/user/v1")
	{
		user.POST("/sign-in", v1.Login)

	}
	auth := r.Group("/redis-manager/auth/v1")
	auth.Use(jwt.JWT())
	{
		auth.POST("/sign-up", v1.Register)
		auth.POST("/refresh", v1.Refresh)
	}
	history := r.Group("/redis-manager/ophistory/v1")
	{
		history.GET("/list", v1.OpHistory) //查看历史操作记录
	}

	cli := r.Group("/redis-manager/cli/v1")
	{
		cli.POST("/opkey", v1.OpKey)             //对key进行操作
		cli.POST("/analysisrdb", v1.AnalysisRdb) //分析dump文件
	}
	codis := r.Group("/redis-manager/codis/v1")
	{
		codis.POST("/add", v1.CodisAdd)            //添加codis的平台地址
		codis.GET("/list", v1.CodisList)           // 列出有哪些平台地址
		codis.GET("/cluster", v1.CodisClusterList) //列出该平台地址有多少个集群
		codis.GET("/group", v1.CodisGroup)         //列出该集群有多少个group
	}
	redis := r.Group("/redis-manager/redis/v1")
	{
		redis.GET("/list", v1.RedisList)
		redis.POST("/add", v1.RedisAdd)
	}
	return r
}
