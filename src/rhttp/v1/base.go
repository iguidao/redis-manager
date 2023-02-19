package v1

import (
	//"log"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/iguidao/redis-manager/src/hsc"
	"github.com/iguidao/redis-manager/src/middleware/logger"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	code := hsc.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      true,
	})
	// c.JSON(http.StatusOK, gin.H{"ok": true})
}

func HandleNotFound(c *gin.Context) {
	code := hsc.NOT_FOUND
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      false,
	})
}

func MethodFails(c *gin.Context) {
	code := hsc.Method_FAILS
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      false,
	})
}

func RouterNotFound(c *gin.Context) {
	accept := c.Request.Header.Get("Accept")
	flag := strings.Contains(accept, "text/html")
	if flag {
		content, err := ioutil.ReadFile("./website/index.html")
		if (err) != nil {
			c.Writer.WriteHeader(404)
			c.Writer.WriteString("Not Found")
			return
		}
		c.Writer.WriteHeader(200)
		c.Writer.Header().Add("Accept", "text/html")
		c.Writer.Write((content))
		c.Writer.Flush()
	}
}

func HttpTemplate(c *gin.Context) {
	data := make(map[string]interface{})
	code := hsc.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"errorCode": code,
		"msg":       hsc.GetMsg(code),
		"data":      data,
	})
}

func Cookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")
	if err != nil {
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}
	c.JSON(http.StatusOK, gin.H{"cookie": cookie})
}
func AuthCheck(c *gin.Context) {
	// var result auth.Result
	userinfo, ok := c.Get("UserInfo")
	if ok {
		logger.Info("===========")
		logger.Info(userinfo.(string))
		logger.Info("===========")
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
