package v1

import (
	"encoding/json"
	"net/http"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/opredis"
	"github.com/iguidao/redis-manager/src/middleware/tools"

	"github.com/gin-gonic/gin"
)

func CodisAdd(c *gin.Context) {
	code := hsc.SUCCESS
	var codisinfo CodisInfo
	var result int
	var ok bool
	err := c.BindJSON(&codisinfo)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Codis add error: ", err)
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(codisinfo)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		result, ok = mysql.DB.AddCodis(codisinfo.Curl, codisinfo.Cname)
		if !ok {
			code = hsc.ERROR
		}
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
			code = hsc.ERROR_NO_CONNEC
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

func CodisOpNode(c *gin.Context) {
	code := hsc.SUCCESS
	var codisnode model.CodisNode
	var topom codisapi.Topom
	var ok bool
	var result interface{}
	var clusterauth string
	err := c.BindJSON(&codisnode)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Codis op node error: ", err)
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(codisnode)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		topom, ok = codisapi.CodisTopom(codisnode.Curl, codisnode.ClusterName)
		if !ok {
			result = "Codis Get topom stats fails."
		}
		for _, v := range topom.Stats.Slots {
			if v.Action.State == "pending" || v.Action.State == "migrating" {
				result = "Codis Group Slot is mving!"
			}
		}
		clusterauth = tools.NewXAuth(codisnode.ClusterName)
		if clusterauth == "" {
			result = "Codis Get ID fail."
		}
		if result == nil {
			if codisnode.OpType == "dilatation" {
				result = opredis.Cdilatationn(codisnode, clusterauth, topom)
			} else if codisnode.OpType == "shrinkage" {
				result = opredis.Cshrinkage(codisnode, clusterauth, topom)
			} else {
				result = "Codis op type fails " + codisnode.OpType
			}
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
