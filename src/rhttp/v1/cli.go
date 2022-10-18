package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/cosop"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/opredis"
	"github.com/iguidao/redis-manager/src/middleware/tools"
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

func AnalysisRdb(c *gin.Context) {
	var clirdb CliRdb
	var result string
	code := hsc.SUCCESS
	err := c.BindJSON(&clirdb)
	if err != nil {
		logger.Error("Rdb Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		if cosop.CosGet(clirdb.RdbName, "/tmp/"+clirdb.RdbName) {
			if opredis.ConnectRedis(cfg.Get_Info("REDIS")) {
				go func() {
					opredis.Analysis("/tmp/"+clirdb.RdbName, "bigkey-"+clirdb.ServerIp)
				}()
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func NewBigkey(c *gin.Context) {
	locaktime := time.Duration(cfg.Get_Info_Int("locktime")) * time.Second
	var cliquery CliQuery
	var serverip string
	result := make(map[string]interface{})
	code := hsc.SUCCESS
	username, _ := c.Get("UserName")
	err := c.BindJSON(&cliquery)
	if err != nil || username == nil {
		logger.Error("Big Key Bind Json error: ", err, "or username: ", username)
		code = hsc.INVALID_PARAMS
	} else if !opredis.LockCheck("BigKey-"+cliquery.ClusterName+"-"+cliquery.GroupName, locaktime) {
		logger.Error("Big Key click repeatedly")
		result["友情提示"] = "别点了，太多人操作了，该操作1次只能1个人！！！"
		code = hsc.CLICK_REPEATEDLY
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory("BigKey", string(jsonBody))
		serverip = codisapi.GetSlave(cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(cfg.Get_Info("REDIS")) {
			clickkeyname := "Click-Bigkey-" + cliquery.ClusterName + "-" + cliquery.GroupName
			tips, ok := tools.BigKeyClick(cliquery.ClusterName, cliquery.GroupName, clickkeyname)
			result["友情提示"] = "大key分析执行出现了问题，请找sre服务台！！！"
			switch ok {
			case 0:
				result["友情提示"] = tips
			case 1:
				if opredis.RedisSave(serverip) {
					result["友情提示"] = tips
					opredis.ExpireKey(clickkeyname, cfg.Get_Info_Int("biglocktime"))
					// result["result"] = "查大key命令已经后台执行，请等待，如若超过10min未查出结果，请找SRE服务台..."
				}
			case 2:
				result["友情提示"] = tips
			case 3:
				result["友情提示"] = tips
				if opredis.ExistsKey("bigkey-" + serverip) {
					keyvalue, ok := opredis.GetStringKey("bigkey-" + serverip)
					if ok {
						result["Top-Key"] = tools.JsonToMap(keyvalue)
						// result["友情提示"] = "据听说1分钟点击5次查询，可以生成最新的大key分析数据"
					}
				}
			}
			// }
		}
		go opredis.LockRm("BigKey-" + cliquery.ClusterName + "-" + cliquery.GroupName)
	}
	// resultmap["友情提示"] = "据听说1分钟点击5次查询，可以生成最新的大key分析数据"
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
