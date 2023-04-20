package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/txcloud"
	"github.com/iguidao/redis-manager/src/middleware/util"
)

func CloudList(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	region := c.Query("region")
	cloud := c.Query("cloud")
	if cloud == "txredis" {
		clist := mysql.DB.GetCloudredis(cloud, region)
		if len(clist) == 0 {
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
		} else {
			result["redis_list"] = clist
		}

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
	if cloud == "txredis" {
		if !txcloud.TxCvmContent() {
			code = hsc.ERROR
		} else {
			list, ok := txcloud.TxListRegion()
			var rlist model.TxRegion
			if ok {
				err := json.Unmarshal([]byte(list), &rlist)
				if err == nil {
					result["region_list"] = rlist.Response
				} else {
					logger.Error("json tx cloud result error: ", err)
				}
			} else {
				code = hsc.ERROR
			}
		}
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
	if err != nil {
		logger.Error("Change Password Bind Json error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		ok := mysql.DB.UpdateCloudPassword(cp.instanceid, cp.Password)
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
