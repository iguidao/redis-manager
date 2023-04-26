package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/casbin"
	"github.com/iguidao/redis-manager/src/middleware/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var token string
		var username string
		var usertype string
		var userid int
		Result := make(map[string]interface{})
		code = hsc.SUCCESS
		urlpath := c.Request.URL.Path
		method := c.Request.Method
		auth := c.Request.Header.Get("Authorization")
		if strings.Contains(auth, "Bearer ") {
			token = strings.Split(auth, "Bearer ")[1]
		}
		if token == "" {
			Result["result"] = "没有Authorization"
			code = hsc.NO_LOGIN
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				Result["result"] = "Token鉴权失败"
				code = hsc.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				Result["result"] = "Token已超时"
				code = hsc.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			} else {
				result := casbin.RuleCheck(claims.UserType, urlpath, method)
				if !result {
					code = hsc.WARN_NOT_PROMISE_RULE
					Result["result"] = "权限不够呀，找管理员开下权限！"
				} else {
					username = claims.UserName
					usertype = claims.UserType
					userid = claims.UserId
				}
			}
		}
		if code == hsc.WARN_NOT_PROMISE_RULE {
			c.JSON(http.StatusOK, gin.H{
				"errorCode": code,
				"msg":       hsc.GetMsg(code),
				"data":      Result,
			})
			c.Abort()
			return
		}
		if code != hsc.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"errorCode": code,
				"msg":       hsc.GetMsg(code),
				"data":      Result,
			})
			c.Abort()
			return
		}
		c.Set("UserName", username)
		c.Set("UserType", usertype)
		c.Set("UserId", userid)
		c.Next()
	}
}
