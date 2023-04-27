package rhttp

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/jwt"
	"github.com/iguidao/redis-manager/src/middleware/model"
	v1 "github.com/iguidao/redis-manager/src/rhttp/v1"
)

// NewServer return a configured http server of gin
func NewServer() *gin.Engine {
	// 存储日志文件代码
	logpath := "./logs/" + cfg.Get_Info_String("logapipath")
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

	r.Use(static.Serve("/", static.LocalFile("website", true)))
	r.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			content, err := ioutil.ReadFile("dist/index.html")
			if (err) != nil {
				c.Writer.WriteHeader(404)
				c.Writer.WriteString("Not Found")
				return
			}
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Write((content))
			c.Writer.Flush()
		}
	})
	// vue配置
	// r.Static("/assets", "./website/assets")
	// r.LoadHTMLFiles("./website/index.html")
	home := r.Group("")
	{
		home.GET("/", v1.Home) //主页接口
	}
	base := r.Group("/redis-manager/base/v1")
	{
		base.GET("/health", v1.HealthCheck) //自检接口
	}
	login := r.Group("/redis-manager/auth/v1")
	{
		login.POST("/sign-in", v1.Login) //登陆接口

	}
	public := r.Group("/redis-manager/public/v1")
	{
		public.POST("/analysisrdb", v1.AnalysisRdb) //分析dump文件
	}
	auth := r.Group("/redis-manager/auth/v1")
	auth.Use(jwt.JWT())
	{
		auth.POST("/password", v1.ChangUserPassword) //更改用户密码
		auth.POST("/refresh", v1.Refresh)            //刷新接口
	}
	board := r.Group(model.PATHBOARD)
	board.Use(jwt.JWT())
	{
		board.GET("/desc", v1.BoardDesc) //board页面
	}
	history := r.Group(model.PATHHISTORY)
	history.Use(jwt.JWT())
	{
		history.GET("/list", v1.OpHistory) //查看历史操作记录
	}
	cfg := r.Group(model.PATHCFG)
	cfg.Use(jwt.JWT())
	{
		cfg.POST("/update", v1.CfgUpdate)          // 添加配置信息
		cfg.GET("/list", v1.CfgList)               // 获取配置信息
		cfg.DELETE("/del", v1.CfgDelete)           //删除配置
		cfg.POST("/adddefault", v1.CfgAddDefault)  //添加默认key
		cfg.GET("/listdefault", v1.CfgListDefault) //返回默认配置key
	}
	codis := r.Group(model.PATHCODIS)
	codis.Use(jwt.JWT())
	{
		codis.POST("/add", v1.CodisAdd)            //添加codis的平台地址
		codis.GET("/list", v1.CodisList)           //列出有哪些平台地址
		codis.GET("/cluster", v1.CodisClusterList) //列出该平台地址有多少个集群
		codis.GET("/group", v1.CodisGroup)         //列出该集群有多少个group
		codis.POST("/opnode", v1.CodisOpNode)      //针对codis的proxy和server节点进行操作
	}
	cloud := r.Group(model.PATHCLOUD)
	cloud.Use(jwt.JWT())
	{
		cloud.GET("/region", v1.RegionList)             //列出云的地域
		cloud.GET("/list", v1.CloudList)                //列出云的集群列表
		cloud.POST("/password", v1.ChangeCloudPassword) //修改数据库保存密码
		cloud.POST("/size", v1.ChangeSize)              //修改集群大小
		cloud.POST("/add", v1.CloudAdd)                 // 添加集群
		cloud.DELETE("/del", v1.CloudDel)               //删除集群
	}
	cluster := r.Group(model.PATHCLUSTER)
	cluster.Use(jwt.JWT())
	{
		cluster.GET("/list", v1.ClusterList)   //列出所有集群
		cluster.GET("/nodes", v1.NodeList)     // 列出集群的node
		cluster.GET("/masters", v1.MasterList) //列出master地址
		cluster.POST("/add", v1.ClusterAdd)    //添加集群
	}
	cli := r.Group(model.PATHCLI)
	cli.Use(jwt.JWT())
	{
		cli.POST("/opkey", v1.OpKey) //对key进行操作
	}
	user := r.Group(model.PATHUSER)
	user.Use(jwt.JWT())
	{
		user.POST("/add", v1.AddUser)          //新增用户接口
		user.GET("/list", v1.ListUser)         //列出所有用户
		user.DELETE("/del", v1.DelUser)        //删除用户
		user.POST("/change", v1.ChangUserType) //更改用户属性
		user.GET("/utype", v1.ListUserType)    //获取用户身份列表
	}
	rule := r.Group(model.PATHRULE)
	rule.Use(jwt.JWT())
	{
		rule.POST("/add", v1.AddRule)   //添加规则
		rule.DELETE("/del", v1.DelRule) //删除规则
		rule.GET("/list", v1.AllRule)   //查看所有规则
		rule.GET("/cfg", v1.GetRuleCfg) //查看默认配置
	}

	r.NoMethod(v1.MethodFails)
	r.NoRoute(v1.RouterNotFound)
	return r
}
