package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 首页
func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"head": "index_head",
	})
}
