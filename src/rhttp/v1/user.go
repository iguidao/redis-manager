package v1

import (
	"net/http"
	"strings"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/useride"
	"github.com/iguidao/redis-manager/src/middleware/util"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	Result := make(map[string]interface{})
	var code int
	var rduser UserInfo
	err := c.BindJSON(&rduser)
	if err != nil {
		code = hsc.INVALID_PARAMS
		Result["result"] = "参数错误"
	} else if !util.VerifyEmailFormat(rduser.Mail) {
		code = hsc.INVALID_PARAMS
		Result["result"] = "邮箱写错了"
	} else if mysql.DB.FindEmail(rduser.Mail) {
		code = hsc.INVALID_PARAMS
		Result["result"] = "邮箱已经注册"
	} else if mysql.DB.FindUser(rduser.UserName) {
		code = hsc.INVALID_PARAMS
		Result["result"] = "用户名已经注册"
	} else {
		scrypt_password := useride.Get_scrypt(rduser.Password)
		result := mysql.DB.CreatUser(rduser.UserName, rduser.Mail, rduser.UserType, scrypt_password)
		if result == false {
			code = hsc.ERROR
			Result["result"] = "创建用户失败"
		} else {
			code = hsc.SUCCESS
			Result["result"] = "创建用户成功"
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      Result,
	})
}

func Login(c *gin.Context) {
	Result := make(map[string]interface{})
	var code int
	var rduser UserInfo
	err := c.BindJSON(&rduser)
	if err != nil {
		logger.Error("login error:", err)
		code = hsc.INVALID_PARAMS
		Result["result"] = "参数错误"
	} else if useride.Gd_login(rduser.UserName, rduser.Password) {
		token, err := util.GenerateToken(rduser.UserName, rduser.Password)
		if err != nil {
			code = hsc.ERROR_AUTH_TOKEN
			Result["result"] = "获取Token失败"
		} else {
			Result["token"] = "Bearer " + token
			Result["result"] = "登录成功"
			Result["username"] = rduser.UserName
			code = hsc.SUCCESS
			// go useride.CacheUserinfo(token, Phonenum)
		}
	} else {
		Result["result"] = "认证失败"
		code = hsc.ERROR_AUTH
	}

	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      Result,
	})
}

func Refresh(c *gin.Context) {
	Result := make(map[string]interface{})
	var code int
	username, _ := c.Get("UserName")
	auth := c.Request.Header.Get("Authorization")
	authtoken := strings.Split(auth, "Bearer ")[1]
	token, err := util.RefreshToken(authtoken)
	if err != nil {
		logger.Error("refresh fails:", err)
		code = hsc.ERROR_AUTH_TOKEN
		Result["result"] = "获取Token失败"
	} else {
		Result["token"] = "Bearer " + token
		Result["result"] = "刷新token成功"
		Result["username"] = username.(string)
		code = hsc.SUCCESS
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      Result,
	})
}
