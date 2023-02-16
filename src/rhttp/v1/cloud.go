package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/txcloud"
)

func CloudList(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	if !txcloud.TxContent() {
		code = hsc.ERROR
	} else {
		list, ok := txcloud.TxListRedis()
		var rlist model.TxL
		if ok {
			err := json.Unmarshal([]byte(list), &rlist)
			if err == nil {
				result["redis_list"] = rlist
			} else {
				logger.Error("json tx cloud result error: ", err)
			}
		} else {
			code = hsc.ERROR
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
