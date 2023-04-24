package v1

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/cluster"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
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
	var result []*model.ClusterNodeTables
	clusterid := c.Query("cluster_id")
	nodes := mysql.DB.GetClusterNode(clusterid)
	for _, v := range nodes {
		if v.Flags == "master" {
			result = append(result, cluster.ClusterConvergeTree(v))
		}
	}
	for _, v := range nodes {
		if v.Flags == "slave" {
			result = cluster.ClusterUpdateTree(result, v)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func MasterList(c *gin.Context) {
	code := hsc.SUCCESS
	clusterid := c.Query("cluster_id")
	nodes := mysql.DB.GetClusterNodeMaster(clusterid)
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      nodes,
	})
}
func ClusterAdd(c *gin.Context) {
	var clusterinfo AddCluster
	result := make(map[string]interface{})
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&clusterinfo)
	code := hsc.ERROR
	if err != nil || clusterinfo.Name == "" || clusterinfo.Nodes == "" {
		logger.Error("Cluster add error: ", err)
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(clusterinfo)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))

		address := strings.Split(clusterinfo.Nodes, ",")
		connectok := false
		if opredis.ConnectRedisCluster(address, clusterinfo.Password) {
			connectok = true
		}
		if connectok {
			id, ok := mysql.DB.AddCluster(clusterinfo.Name, clusterinfo.Nodes, clusterinfo.Password)
			if ok || id != 0 {
				nodeinfo := opredis.CGetClusterNode()
				for _, v := range nodeinfo {
					if len(v) != 0 {
						cluster.WriteCluster(id, v)
					}
				}
				code = hsc.SUCCESS
			} else {
				logger.Error("添加集群到 cluster info 失败")
				code = hsc.ERROR_WRITE_MYSQL
			}
		} else {
			logger.Error("链接目标redis异常")
			code = hsc.ERROR_NO_CONNEC
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
