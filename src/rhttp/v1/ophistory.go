package v1

import (
	"net/http"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/mysql"

	"github.com/gin-gonic/gin"
)

func OpHistory(c *gin.Context) {
	code := hsc.SUCCESS
	result := mysql.DB.GetAllHistory()
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}
