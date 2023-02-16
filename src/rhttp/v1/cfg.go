package v1

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func CfgListDefault(c *gin.Context) {
	code := hsc.SUCCESS
	var result []string
	custom := mysql.DB.GetOneCfg(model.CC)
	for _, v := range model.CfgDefault {
		result = append(result, v)
	}
	if custom.Value != "" {
		list := strings.Split(custom.Value, ",")
		result = append(result, list...)
	}

	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func CfgAddDefault(c *gin.Context) {
	code := hsc.SUCCESS
	var cfg ConfigInfo
	err := c.BindJSON(&cfg)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Config add Default error: ", err)
	} else {
		if mysql.DB.ExistCfg(model.CC) {
			custom := mysql.DB.GetOneCfg(model.CC)
			value := custom.Value + "," + cfg.Value
			if !mysql.DB.UpdateCfg(model.CC, value) {
				log.Println("update error")
				code = hsc.ERROR
			}
		} else {
			_, ok := mysql.DB.AddCfg(model.CC, cfg.Value, model.CN)
			if !ok {
				log.Println("add error")
				code = hsc.ERROR
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      hsc.GetMsg(code),
	})
}

func CfgList(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	cfglist := mysql.DB.GetAllCfg()
	result["lists"] = cfglist
	result["total"] = len(cfglist)
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func CfgAdd(c *gin.Context) {
	code := hsc.SUCCESS
	var cfg ConfigInfo
	var result int
	var ok bool
	err := c.BindJSON(&cfg)
	if err != nil {
		code = hsc.INVALID_PARAMS
		logger.Error("Config add error: ", err)
	} else {
		result, ok = mysql.DB.AddCfg(cfg.Name, cfg.Value, cfg.Note)
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
