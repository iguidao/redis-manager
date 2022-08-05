package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func QueryKey(c *gin.Context) {
	var cliquery CliQuery
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&cliquery)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	log.Println(cliquery)
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}

func BigKey(c *gin.Context) {
	var clusterinfo AddCluster
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&clusterinfo)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	log.Println(clusterinfo)
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}

func HotKey(c *gin.Context) {
	var clusterinfo AddCluster
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&clusterinfo)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	log.Println(clusterinfo)
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}

func AllKey(c *gin.Context) {
	var clusterinfo AddCluster
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&clusterinfo)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	log.Println(clusterinfo)
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}

func SlowKey(c *gin.Context) {
	var clusterinfo AddCluster
	// staff_id, err := strconv.Atoi(UserId)
	err := c.BindJSON(&clusterinfo)
	if err != nil {
		logger.Error("Cluster add error: ", err)
	}
	log.Println(clusterinfo)
	code := hsc.SUCCESS
	result := mysql.DB.GetAllCluster()
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  hsc.GetMsg(code),
		"data": result,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}
