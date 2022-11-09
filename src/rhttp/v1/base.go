package v1

import (
	//"log"

	"net/http"

	"github.com/iguidao/redis-manager/src/hsc"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	code := hsc.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      true,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}
