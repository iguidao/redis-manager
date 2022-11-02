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
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}
func CodisList(c *gin.Context) {
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCodis()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
}
func CodisClusterList(c *gin.Context) {
	var listresult []string
	code := hsc.SUCCESS
	var codisinfo CodisInfo
	err := c.BindJSON(&codisinfo)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Codis get error: ", err)
	} else {
		listresult = codisapi.GeClusterList(codisinfo.Curl)
		if len(listresult) == 0 {
			code = hsc.NO_CONNECT_CODIS
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": listresult,
	})
}

func CodisGroup(c *gin.Context) {
	var listresult []string
	code := hsc.SUCCESS
	var codisinfo CodisInfo
	err := c.BindJSON(&codisinfo)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Codis get error: ", err)
	} else {
		listresult = codisapi.GetGroup(codisinfo.Curl, codisinfo.ClusterName)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": listresult,
	})
}
