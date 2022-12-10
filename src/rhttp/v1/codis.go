package v1

import (
	"net/http"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"

	"github.com/gin-gonic/gin"
)

func CodisAdd(c *gin.Context) {
	code := hsc.SUCCESS
	var codisinfo CodisInfo
	err := c.BindJSON(&codisinfo)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Codis add error: ", err)
	}

	result, ok := mysql.DB.AddCodis(codisinfo.Curl, codisinfo.Cname)
	if !ok {
		code = hsc.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func CodisList(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	codislist := mysql.DB.GetAllCodis()
	result["lists"] = codislist
	result["total"] = len(codislist)
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func CodisClusterList(c *gin.Context) {
	var listresult []string
	code := hsc.INVALID_PARAMS
	curl := c.Query("curl")
	if curl != "" {
		code = hsc.SUCCESS
		listresult = codisapi.GeClusterList(curl)
		if len(listresult) == 0 {
			code = hsc.NO_CONNECT_CODIS
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      listresult,
	})
}

func CodisGroup(c *gin.Context) {
	var listresult []string
	code := hsc.INVALID_PARAMS
	curl := c.Query("curl")
	clustername := c.Query("cluster_name")
	if curl != "" || clustername != "" {
		code = hsc.SUCCESS
		listresult = codisapi.GetGroup(curl, clustername)
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      listresult,
	})
}
