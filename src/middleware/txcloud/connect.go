package txcloud

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tredis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
)

// DB as the mysql client
var TAPI *tredis.Client

func TxContent() bool {
	var err error
	credential := common.NewCredential(
		mysql.DB.GetOneCfgValue(model.TXSECRETID),
		mysql.DB.GetOneCfgValue(model.TXSECRETKEY),
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = mysql.DB.GetOneCfgValue(model.TXAPIURL)
	// 实例化要请求产品的client对象,clientProfile是可选的
	TAPI, err = tredis.NewClient(credential, mysql.DB.GetOneCfgValue(model.TXREGION), cpf)
	if err != nil {
		logger.Error("conenct tx cloud error: ", err)
		return false
	}
	return true
}
