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

	// vue配置
	// r.Static("/assets", "./website/assets")
	// r.LoadHTMLFiles("./website/index.html")
	// home := r.Group("")
	// {
	// 	home.GET("/", v1.Home) //主页接口
	// }

	base := r.Group("/redis-manager/base/v1")
	{
		base.GET("/health", v1.HealthCheck) //自检接口
	}
	user := r.Group("/redis-manager/user/v1")
	{
		user.POST("/sign-in", v1.Login) //登陆接口

	}
	auth := r.Group("/redis-manager/auth/v1")
	// auth.Use(jwt.JWT())
	{
		auth.POST("/sign-up", v1.Register) //注册接口
		auth.POST("/refresh", v1.Refresh)  //刷新接口
	}
	history := r.Group("/redis-manager/ophistory/v1")
	history.Use(jwt.JWT())
	{
		history.GET("/list", v1.OpHistory) //查看历史操作记录
	}
	cfg := r.Group("/redis-manager/cfg/v1")
	cfg.Use(jwt.JWT())
	{
		cfg.POST("/update", v1.CfgUpdate)          // 添加配置信息
		cfg.GET("/list", v1.CfgList)               // 获取配置信息
		cfg.DELETE("/del", v1.CfgDelete)           //删除配置
		cfg.POST("/adddefault", v1.CfgAddDefault)  //添加默认key
		cfg.GET("/listdefault", v1.CfgListDefault) //返回默认配置key
	}
	codis := r.Group("/redis-manager/codis/v1")
	codis.Use(jwt.JWT())
	{
		codis.POST("/add", v1.CodisAdd)            //添加codis的平台地址
		codis.GET("/list", v1.CodisList)           //列出有哪些平台地址
		codis.GET("/cluster", v1.CodisClusterList) //列出该平台地址有多少个集群
		codis.GET("/group", v1.CodisGroup)         //列出该集群有多少个group
		codis.POST("/opnode", v1.CodisOpNode)      //针对codis的proxy和server节点进行操作
	}
	cli := r.Group("/redis-manager/cli/v1")
	cli.Use(jwt.JWT())
	{
		cli.POST("/opkey", v1.OpKey)             //对key进行操作
		cli.POST("/analysisrdb", v1.AnalysisRdb) //分析dump文件
	}

	redis := r.Group("/redis-manager/redis/v1")
	redis.Use(jwt.JWT())
	{
		redis.GET("/list", v1.RedisList)
		redis.POST("/add", v1.RedisAdd)
	}
	cloud := r.Group("/redis-manager/cloud/v1")
	cloud.Use(jwt.JWT())
	{
		cloud.GET("/list", v1.CloudList)
		cloud.GET("/region", v1.RegionList)
		cloud.POST("/password", v1.ChangeCloudPassword)
	}
	rootrule := r.Group("/permission-internal/v1")
	rootrule.Use(jwt.JWT())
	{
		rootrule.POST("/rule/add", v1.AddRule)
		rootrule.DELETE("/rule/del", v1.DelRule)
		rootrule.PUT("/rule/update", v1.UpdateRule)
		rootrule.GET("/rule/all", v1.AllRule)
	}
	checkrule := r.Group("/permission-internal/v1")
	{
		checkrule.POST("/rule/check", v1.CheckRule)
	}
	authcheck := r.Group("/permission-internal/authcheck/v1")
	authcheck.Use(jwt.JWT())
	{
		authcheck.GET("/test", v1.AuthCheck)
	}
	r.NoMethod(v1.MethodFails)
	r.NoRoute(v1.RouterNotFound)
	return r
}
