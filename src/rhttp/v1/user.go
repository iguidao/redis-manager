package v1

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/middleware/useride"
	"github.com/iguidao/redis-manager/src/middleware/util"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	code := hsc.SUCCESS
	result := mysql.DB.GetAllUser()
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      result,
	})
}
func ListUserType(c *gin.Context) {
	code := hsc.SUCCESS
	var typelist []map[string]string
	for k, v := range model.DefaultUser {
		t := make(map[string]string)
		t["label"] = v
		t["value"] = k
		typelist = append(typelist, t)
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      typelist,
	})
}
func AddUser(c *gin.Context) {
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
		code = hsc.WARN_USER_MAIL_EXIST
		Result["result"] = "邮箱已经注册"
	} else if mysql.DB.FindUser(rduser.UserName) {
		code = hsc.WARN_USER_NAME_EXIST
		Result["result"] = "用户名已经注册"
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(rduser)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		scrypt_password := useride.Get_scrypt(rduser.Password)
		result := mysql.DB.CreatUser(rduser.UserName, rduser.Mail, scrypt_password)
		if !result {
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

func DelUser(c *gin.Context) {
	Result := make(map[string]interface{})
	var code int
	userid := c.Query("userid")
	id, err := strconv.Atoi(userid)
	if userid == "" || err != nil {
		code = hsc.INVALID_PARAMS
		Result["result"] = "参数错误"
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(userid)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		result := mysql.DB.DelUser(id)
		if !result {
			code = hsc.ERROR
			Result["result"] = "删除用户失败"
		} else {
			code = hsc.SUCCESS
			Result["result"] = "删除用户成功"
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      Result,
	})
}

func ChangUserPassword(c *gin.Context) {
	Result := make(map[string]interface{})
	var code int
	var upasssword UserPassword
	err := c.BindJSON(&upasssword)

	if err != nil || upasssword.Old == "" || upasssword.New == "" {
		code = hsc.INVALID_PARAMS
		Result["result"] = "参数错误"
	} else {
		username, _ := c.Get("UserName")
		userid, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(upasssword)
		method := c.Request.Method
		go mysql.DB.AddHistory(userid.(int), method+":"+urlinfo.Path, string(jsonBody))

		if useride.Gd_login(username.(string), upasssword.Old) {
			scrypt_password := useride.Get_scrypt(upasssword.New)
			result := mysql.DB.UpdateUserPassword(username.(string), scrypt_password)
			if !result {
				code = hsc.ERROR
				Result["result"] = "更改密码失败"
			} else {
				code = hsc.SUCCESS
				Result["result"] = "更改密码成功"
			}
		} else {
			code = hsc.WARN_USER_PASSWORD_CHECK
			Result["result"] = "旧的账号密码不对"
		}

	}
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      Result,
	})
}
func ChangUserType(c *gin.Context) {
	Result := make(map[string]interface{})
	var code int
	var rduser UserInfo
	err := c.BindJSON(&rduser)
	if err != nil || rduser.UserName == "" || rduser.UserType == "" {
		code = hsc.INVALID_PARAMS
		Result["result"] = "参数错误"
	} else {
		username, _ := c.Get("UserId")
		urlinfo := c.Request.URL
		jsonBody, _ := json.Marshal(rduser)
		method := c.Request.Method
		go mysql.DB.AddHistory(username.(int), method+":"+urlinfo.Path, string(jsonBody))
		result := false
		if mysql.DB.ExistUserName(rduser.UserName) {
			result = mysql.DB.UpdateUserType(rduser.UserName, rduser.UserType)
		}
		if !result {
			code = hsc.ERROR
			Result["result"] = "变更用户失败"
		} else {
			code = hsc.SUCCESS
			Result["result"] = "变更用户成功"
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
	if err != nil || rduser.UserName == "" || rduser.Password == "" {
		logger.Error("login error:", err)
		code = hsc.INVALID_PARAMS
		Result["result"] = "参数错误"
	} else if useride.Gd_login(rduser.UserName, rduser.Password) {
		token, usertype, err := util.GenerateToken(rduser.UserName, rduser.Password)
		if err != nil {
			code = hsc.ERROR_AUTH_TOKEN
			Result["result"] = "获取Token失败"
		} else {
			Result["token"] = "Bearer " + token
			Result["result"] = "登录成功"
			Result["username"] = rduser.UserName
			Result["usertype"] = usertype
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
