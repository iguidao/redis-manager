package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/iguidao/redis-manager/src/hsc"
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
				username = claims.UserName
				usertype = claims.UserType
				userid = claims.UserId
			}
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
