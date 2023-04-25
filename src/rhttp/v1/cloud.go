package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/alicloud"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/txcloud"
	"github.com/iguidao/redis-manager/src/middleware/util"
)

func CloudAdd(c *gin.Context) {
	var result interface{}
	code := hsc.ERROR
	// username, _ := c.Get("UserId")
	// urlinfo := c.Request.URL
	// jsonBody, _ := json.Marshal(shardcfg)
	// method := c.Request.Method
	// go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func CloudDel(c *gin.Context) {
	code := hsc.SUCCESS
	instanceid := c.Query("instanceid")
	result := false
	if instanceid != "" {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(instanceid)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		if !mysql.DB.DelCloud(instanceid) {
			result = false
			code = hsc.SERVER_ERROR
		}
	} else {
		result = false
		code = hsc.INVALID_PARAMS
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func CloudList(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	region := c.Query("region")
	cloud := c.Query("cloud")
	clist := mysql.DB.GetCloudredis(cloud, region)
	if len(clist) == 0 {
		switch cloud {
		case "txredis":
			if !txcloud.TxRedisContent(region) {
				code = hsc.ERROR
			} else {
				list, ok := txcloud.TxListRedis()
				var rlist model.TxL
				if ok {
					err := json.Unmarshal([]byte(list), &rlist)
					if err == nil {
						go util.TxWriteRedis(cloud, rlist)
						code = hsc.WARN_BACKGROUND
					} else {
						logger.Error("json tx cloud result error: ", err)
						code = hsc.ERROR_BACKGROUND
					}
				} else {
					code = hsc.ERROR
				}
			}
		case "aliredis":
			if !alicloud.AliRedisContent() {
				code = hsc.ERROR
			} else {
				list, ok := alicloud.AliListRedis(region)
				var rlist model.AliRedis
				if ok {
					err := json.Unmarshal([]byte(list), &rlist)
					if err == nil {
						go util.AliWriteRedis(cloud, rlist)
						code = hsc.WARN_BACKGROUND
					} else {
						logger.Error("json tx cloud result error: ", err)
						code = hsc.ERROR_BACKGROUND
					}
				} else {
					code = hsc.ERROR
				}
			}
		default:
			result["WARN"] = "暂时不支持该云操作"
		}
	} else {
		result["redis_list"] = clist
	}

	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func RegionList(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	cloud := c.Query("cloud")
	switch cloud {
	case "txredis":
		if !txcloud.TxCvmContent() {
			code = hsc.ERROR_CLOUD_CONNECT
		} else {
			list, ok := txcloud.TxListRegion()
			var rlist model.TxRegion
			if ok {
				err := json.Unmarshal([]byte(list), &rlist)
				if err == nil {
					result["region_list"] = rlist.Response.RegionSet
				} else {
					logger.Error("json tx cloud result error: ", err)
				}
			} else {
				code = hsc.ERROR_CLOUD_GET
			}
		}
	case "aliredis":
		if !alicloud.AliRedisContent() {
			code = hsc.ERROR_CLOUD_CONNECT
		} else {
			list, ok := alicloud.AliListRegion()
			var rlist model.AliRegion
			if ok {
				err := json.Unmarshal([]byte(list), &rlist)
				if err == nil {
					result["region_list"] = rlist.RegionIds.KVStoreRegion
				} else {
					logger.Error("json ali cloud result error: ", err)
				}
			} else {
				code = hsc.ERROR_CLOUD_GET
			}
		}
	default:
		code = hsc.WARN_NOT_FOUND_CLOUD
		result["WARN"] = "暂时不支持该云操作"
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func ChangeCloudPassword(c *gin.Context) {
	var cp CloudPassword
	var result interface{}
	code := hsc.ERROR
	err := c.BindJSON(&cp)
	if err != nil || cp.Instanceid == "" || cp.Password == "" || cp.Cloud == "" {
		logger.Error("Change Password Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(cp)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		ok := mysql.DB.UpdateCloudPassword(cp.Cloud, cp.Instanceid, cp.Password)
		if ok {
			code = hsc.SUCCESS
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func ChangeSize(c *gin.Context) {
	var shardcfg TxShardCfg
	var result interface{}
	code := hsc.ERROR
	err := c.BindJSON(&shardcfg)
	if err != nil {
		logger.Error("Change Password Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(shardcfg)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		if shardcfg.Cloud == "txredis" {

		}
		// ok := mysql.DB.UpdateCloudPassword(cp.Instanceid, cp.Password)
		// if ok {
		// 	code = hsc.SUCCESS
		// }
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
