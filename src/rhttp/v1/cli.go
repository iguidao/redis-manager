package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/opredis"
)

func QueryKey(c *gin.Context) {
	var cliquery CliQuery
	var result opredis.QueryResult
	code := hsc.SUCCESS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Query Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		proxylist := codisapi.GetProxy(cliquery.ClusterName)
		for _, v := range proxylist {
			if opredis.ConnectRedis(v) {
				result = opredis.QueryKey(cliquery.KeyName)
				break
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func BigKey(c *gin.Context) {
	var cliquery CliQuery
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	code := hsc.SUCCESS
	result := mysql.DB.GetAllHistory()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func HotKey(c *gin.Context) {
	var cliquery CliQuery
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	code := hsc.SUCCESS
	result := mysql.DB.GetAllHistory()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func AllKey(c *gin.Context) {
	var cliquery CliQuery
	var result []string
	code := hsc.SUCCESS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		serverip := codisapi.GetSlave(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.AllKey()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func SlowKey(c *gin.Context) {
	var cliquery CliQuery
	var result []redis.SlowLog
	code := hsc.SUCCESS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		serverip := codisapi.GetMaster(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.SlowKey()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}
