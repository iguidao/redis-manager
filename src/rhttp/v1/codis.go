package v1

import (
	"net/http"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/codisapi"

	"github.com/gin-gonic/gin"
)

func CodisList(c *gin.Context) {
	code := hsc.SUCCESS
	listresult := codisapi.GetList()
	if len(listresult) == 0 {
		code = hsc.NO_CONNECT_CODIS
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": listresult,
	})
}

func CodisGroup(c *gin.Context) {
	code := hsc.SUCCESS
	clustername := c.Query("cluster_name")
	listresult := codisapi.GetGroup(clustername)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": listresult,
	})
}
