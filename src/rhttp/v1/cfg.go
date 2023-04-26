package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func CfgListDefault(c *gin.Context) {
	code := hsc.SUCCESS
	listcfg := model.DefaultName
	var result []map[string]string
	for k, v := range listcfg {
		cfg := make(map[string]string)
		cfg["label"] = v
		cfg["value"] = k
		result = append(result, cfg)
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
				logger.Error("update cfg error")
				code = hsc.ERROR
			}
		} else {
			username, _ := c.Get("UserId")
			urlinfo := c.Request.URL
			jsonBody, _ := json.Marshal(cfg)
			method := c.Request.Method
			go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
			_, ok := mysql.DB.AddCfg(model.CC, cfg.Value, model.CN)
			if !ok {
				logger.Error("add cfg error")
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

func CfgUpdate(c *gin.Context) {
	code := hsc.SUCCESS
	var cfg ConfigInfo
	var result int
	var ok bool
	err := c.BindJSON(&cfg)
	if err != nil || cfg.Key == "" || cfg.Value == "" {
		code = hsc.INVALID_PARAMS
		logger.Error("Config add error: ", err)
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(cfg)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		name := model.DefaultName[cfg.Key]
		if mysql.DB.ExistCfg(cfg.Key) {
			if !mysql.DB.UpdateCfg(cfg.Key, cfg.Value) {
				code = hsc.ERROR
			}
		} else {
			result, ok = mysql.DB.AddCfg(name, cfg.Key, cfg.Value)
			if !ok {
				code = hsc.ERROR
			}
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}

func CfgDelete(c *gin.Context) {
	code := hsc.SUCCESS
	key := c.Query("key")
	result := false
	if key != "" {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(key)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		if !mysql.DB.DelCfg(key) {
			result = false
			code = hsc.SERVER_ERROR
		}
	} else {
		result = false
		code = hsc.INVALID_PARAMS
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
