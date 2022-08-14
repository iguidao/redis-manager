package v1

import (
	"encoding/json"
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
	username, _ := c.Get("UserName")
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Query Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory(username.(string), "QueryKey", string(jsonBody))
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
	var result map[string]interface{}
	code := hsc.SUCCESS
	username, _ := c.Get("UserName")
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory(username.(string), "BigKey", string(jsonBody))
		serverip := codisapi.GetSlave(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.BigKey()
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func HotKey(c *gin.Context) {
	var cliquery CliQuery
	var result map[string]int
	code := hsc.SUCCESS
	username, _ := c.Get("UserName")
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory(username.(string), "HotKey", string(jsonBody))
		serverip := codisapi.GetMaster(cliquery.ClusterName, cliquery.GroupName)
		result = opredis.HotKey(serverip)
	}
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
	username, _ := c.Get("UserName")
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory(username.(string), "AllKey", string(jsonBody))
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
	username, _ := c.Get("UserName")
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory(username.(string), "SlowKey", string(jsonBody))
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
