package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func BoardDesc(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	result["aliredis"] = mysql.DB.GetCloudNumber("aliredis")
	result["txredis"] = mysql.DB.GetCloudNumber("txredis")
	result["codis"] = mysql.DB.GetCodisNumber()
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
