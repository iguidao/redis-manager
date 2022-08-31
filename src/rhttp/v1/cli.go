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
	code := hsc.NO_CONNECT_CODIS

	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Query Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("QueryKey", string(jsonBody))
		proxylist := codisapi.GetProxy(cliquery.ClusterName)
		for _, v := range proxylist {
			if opredis.ConnectRedis(v) {
				result = opredis.QueryKey(cliquery.KeyName)
				code = hsc.SUCCESS
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
	code := hsc.NO_CONNECT_CODIS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("BigKey", string(jsonBody))
		serverip := codisapi.GetSlave(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.BigKey()
			code = hsc.SUCCESS
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
	code := hsc.NO_CONNECT_CODIS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("HotKey", string(jsonBody))
		serverip := codisapi.GetMaster(cliquery.ClusterName, cliquery.GroupName)
		result = opredis.HotKey(serverip)
		code = hsc.SUCCESS
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
	code := hsc.NO_CONNECT_CODIS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("AllKey", string(jsonBody))
		serverip := codisapi.GetSlave(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.AllKey()
			code = hsc.SUCCESS
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
	code := hsc.NO_CONNECT_CODIS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Slow Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("SlowKey", string(jsonBody))
		serverip := codisapi.GetMaster(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.SlowKey()
			code = hsc.SUCCESS
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func DelKey(c *gin.Context) {
	var cliquery CliQuery
	var result string
	code := hsc.NO_CONNECT_CODIS
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Query Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("DelKey", string(jsonBody))
		serverip := codisapi.GetSlave(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip) {
			result = opredis.DeleteKey(cliquery.KeyName)
			code = hsc.SUCCESS
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
