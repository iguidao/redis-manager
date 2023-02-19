package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetPage get page parameters
func GetPage(c *gin.Context) (int, int) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		panic(err)
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil {
		panic(err)
	}

	return page, size

}
