package v1

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/casbin"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/util"
)

func AllRule(c *gin.Context) {
	data := make(map[string]interface{})
	code := hsc.SUCCESS
	pages, size := util.GetPage(c)
	casbinlist := casbin.RuleGet(pages, size)
	data["result"] = casbinlist
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      data,
	})
}
func GetRuleCfg(c *gin.Context) {
	code := hsc.SUCCESS
	result := make(map[string]interface{})
	var cfglist []map[string]string
	for k, v := range model.DefaultPath {
		cfg := make(map[string]string)
		cfg["label"] = v
		cfg["value"] = k
		cfglist = append(cfglist, cfg)
	}
	var methodlist []map[string]string
	for k, v := range model.DefaultMethod {
		method := make(map[string]string)
		method["label"] = v
		method["value"] = k
		methodlist = append(methodlist, method)
	}
	result["url"] = cfglist
	result["method"] = methodlist
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func AddRule(c *gin.Context) {
	data := make(map[string]interface{})
	code := hsc.SUCCESS
	var Policy CasbinPolicyJson
	err := c.BindJSON(&Policy)
	if err != nil {
		log.Println(err)
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(Policy)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		if !casbin.RuleAdd(Policy.Identity, Policy.Path, Policy.Method) {
			code = hsc.ERROR
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      data,
	})
}

func DelRule(c *gin.Context) {
	data := make(map[string]interface{})
	code := hsc.SUCCESS
	var Policy CasbinPolicyJson
	err := c.BindJSON(&Policy)
	if err != nil {
		log.Println(err)
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(Policy)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		if !casbin.RuleDel(Policy.Identity, Policy.Path, Policy.Method) {
			code = hsc.ERROR
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      data,
	})
}

func RootCheck(c *gin.Context) {
	// var result auth.Result
	userinfo, ok := c.Get("UserName")
	log.Println(userinfo.(string))
	if !ok {
		data := make(map[string]interface{})
		code := hsc.NOT_PROMISE
		c.JSON(http.StatusOK, gin.H{
			"errorCode": code,
			"msg":       hsc.GetMsg(code),
			"data":      data,
		})
		return
	}

	// 请求用户id
	userid := userinfo.(string)
	// 请求的path
	path := c.Request.URL.Path
	// 请求的方法
	method := c.Request.Method
	if !casbin.RuleCheck(userid, path, method) {
		data := make(map[string]interface{})
		code := hsc.ERROR
		c.JSON(http.StatusOK, gin.H{
			"errorCode": code,
			"msg":       hsc.GetMsg(code),
			"data":      data,
		})
		return
	}
	c.Next()
}