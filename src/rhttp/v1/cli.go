package v1

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/cosop"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/opredis"
	"github.com/iguidao/redis-manager/src/middleware/tools"
)

func OpKey(c *gin.Context) {
	locaktime := time.Duration(cfg.Get_Info_Int("locktime")) * time.Second
	var cliquery CliQuery
	var result interface{}
	var ok bool
	code := hsc.NO_CONNECT_CODIS
	urlinfo := c.Request.URL
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Op Key Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else if !opredis.LockCheck(cliquery.CacheOp+"-"+cliquery.CacheType+"-"+cliquery.ClusterName+"-"+cliquery.KeyName, locaktime) {
		logger.Error(cliquery.CacheOp + "-" + cliquery.CacheType + "-" + cliquery.ClusterName + "-" + cliquery.KeyName + " Key click repeatedly")
		code = hsc.CLICK_REPEATEDLY
		result = "别点了，太多人操作了，该操作1次只能1个人！！！"
	} else {
		jsonBody, _ := json.Marshal(cliquery)
		go mysql.DB.AddHistory(urlinfo.Path, string(jsonBody))
		if cliquery.CacheType == "codis" {
			result, ok = CodisOp(cliquery)
			if ok {
				code = hsc.SUCCESS
			}
		}
		go opredis.LockRm(cliquery.CacheOp + "-" + cliquery.CacheType + "-" + cliquery.ClusterName + "-" + cliquery.KeyName)
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}

func CodisOp(cliquery CliQuery) (interface{}, bool) {
	switch cliquery.CacheOp {
	case "query":
		proxylist := codisapi.GetProxy(cliquery.CodisUrl, cliquery.ClusterName)
		for _, v := range proxylist {
			if opredis.ConnectRedis(v, "") {
				result := opredis.QueryKey(cliquery.KeyName)
				return result, true
			}
		}
		return nil, false
	case "hot":
		serverip := codisapi.GetMaster(cliquery.CodisUrl, cliquery.ClusterName, cliquery.GroupName)
		result := opredis.HotKey(serverip)
		return result, true
	case "all":
		serverip := codisapi.GetSlave(cliquery.CodisUrl, cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip, "") {
			result := opredis.AllKey()
			return result, true
		}
		return nil, false
	case "slow":
		serverip := codisapi.GetMaster(cliquery.CodisUrl, cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(serverip, "") {
			result := opredis.SlowKey()
			return result, true
		}
		return nil, false
	case "del":
		proxylist := codisapi.GetProxy(cliquery.CodisUrl, cliquery.ClusterName)
		for _, v := range proxylist {
			if opredis.ConnectRedis(v, "") {
				result := opredis.DeleteKey(cliquery.KeyName)
				return result, true
			}
		}
		return nil, false
	case "big":
		result := make(map[string]interface{})
		serverip := codisapi.GetSlave(cliquery.CodisUrl, cliquery.ClusterName, cliquery.GroupName)
		if opredis.ConnectRedis(cfg.Get_Info_String("REDIS"), cfg.Get_Info_String("redispw")) {
			clickkeyname := "Click-Bigkey-" + cliquery.CacheType + "-" + cliquery.ClusterName + "-" + cliquery.GroupName
			tips, ok := tools.BigKeyClick(cliquery.ClusterName, cliquery.GroupName, clickkeyname)
			result["友情提示"] = "大key分析执行出现了问题，请找sre服务台！！！"
			switch ok {
			case 0:
				result["友情提示"] = tips
			case 1:
				result["友情提示"] = tips
				opredis.ExpireKey(clickkeyname, cfg.Get_Info_Int("biglocktime"))
				if opredis.ConnectRedis(serverip, "") {
					opredis.RedisSave(serverip)
				}
			case 2:
				result["友情提示"] = tips
			case 3:
				result["友情提示"] = tips
				if opredis.ExistsKey("bigkey-" + serverip) {
					keyvalue, ok := opredis.GetStringKey("bigkey-" + serverip)
					if ok {
						result["Top-Key"] = tools.JsonToMap(keyvalue)
					}
				}
			}
			return result, true
		}
		return nil, false
	default:
		return "没有找到这个查询key的方式: " + cliquery.CacheOp, false
	}

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
			if opredis.ConnectRedis(cfg.Get_Info_String("REDIS"), cfg.Get_Info_String("redispw")) {
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
