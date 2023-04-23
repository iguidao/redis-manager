package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/cluster"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/opredis"
)

func ClusterList(c *gin.Context) {
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}
func NodeList(c *gin.Context) {
	code := hsc.SUCCESS
	cluster := c.Query("cluster")
	result := mysql.DB.GetClusterNode(cluster)
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func ClusterAdd(c *gin.Context) {
	var clusterinfo AddCluster
	result := make(map[string]interface{})
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&clusterinfo)
	code := hsc.ERROR
	if err != nil {
		logger.Error("Cluster add error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(clusterinfo)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		id, ok := mysql.DB.AddCluster(clusterinfo.Name, clusterinfo.Nodes, clusterinfo.Password)
		if ok || id != 0 {
			address := strings.Split(clusterinfo.Nodes, ",")
			for _, add := range address {
				if opredis.ConnectRedis(add, "") {
					nodeinfo := opredis.GetCluster()
					for _, v := range nodeinfo {
						if len(v) != 0 {
							cluster.WriteCluster(id, v)
						}
					}
				}
			}
		} else {
			logger.Error("添加集群到 cluster info 失败")
			code = hsc.ERROR_WRITE_MYSQL
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
