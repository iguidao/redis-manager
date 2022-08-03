package v1

import (
	"net/http"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/mysql"

	"github.com/gin-gonic/gin"
)

func ClusterList(c *gin.Context) {
	code := hsc.SUCCESS

	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}
