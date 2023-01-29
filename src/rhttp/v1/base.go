package v1

import (
	//"log"

	"io/ioutil"
	"net/http"
	"strings"

	"github.com/iguidao/redis-manager/src/hsc"

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
