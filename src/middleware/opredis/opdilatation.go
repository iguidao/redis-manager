package opredis

import (
	"time"

	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
)

// Codis扩容操作
func Cdilatation(codisnode model.CodisChangeNode, auth string, topom codisapi.Topom) map[string]interface{} {
	result := make(map[string]interface{})
	ok, uplist := UpClusterHost(codisnode, topom, auth)

	if !ok {
		logger.Error("Codis Opnode dilatation: codis add host fails")
	}
	time.Sleep(time.Duration(5) * time.Second)
	if !CodisRebalanceAll(codisnode.Curl, codisnode.ClusterName, auth) {
		logger.Error("Codis Opnode dilatation: codis Reabalance fails")
	}
	result["status"] = ok
	result["downlist"] = uplist
	return result
}
