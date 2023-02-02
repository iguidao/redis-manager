package opredis

import (
	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
)

// 缩容操作
func Cshrinkage(codisnode model.CodisNode, auth string, topom codisapi.Topom) map[string]interface{} {
	result := make(map[string]interface{})
	ok, downlist := DownClusterHost(codisnode, auth, topom.Stats)
	if !ok {
		logger.Error("codis down host fails")
	}
	result["status"] = ok
	result["downlist"] = downlist
	return result
}
