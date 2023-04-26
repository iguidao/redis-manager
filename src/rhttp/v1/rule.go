package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/casbin"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/util"
)

func AllRule(c *gin.Context) {
	code := hsc.SUCCESS
	var result []interface{}
	casbinlist := casbin.RuleGet()
	if len(casbinlist) != 0 {
		for _, v := range casbinlist {
			rinfo := make(map[string]string)
			rinfo["identity"] = v.V0
			rinfo["path"] = v.V1
			rinfo["method"] = v.V2
			rinfo["note"] = util.ReturnDefaultModel(v.V1, model.DefaultPath)
			result = append(result, rinfo)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
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
	ppath := c.Query("path")
	pmethod := c.Query("method")
	pidentity := c.Query("identity")
	if ppath == "" || pmethod == "" || pidentity == "" {
		code = hsc.INVALID_PARAMS
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal("Path:" + ppath + " method:" + pmethod + " identity:" + pidentity)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		if !casbin.RuleDel(pidentity, ppath, pmethod) {
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
