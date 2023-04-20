package txcloud

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	dbbrain "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain/v20210527"
	tredis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
)

// DB as the mysql client
var TxRedisApi *tredis.Client

func TxRedisContent(region string) bool {
	var err error
	credential := common.NewCredential(
		mysql.DB.GetOneCfgValue(model.TXSECRETID),
		mysql.DB.GetOneCfgValue(model.TXSECRETKEY),
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = mysql.DB.GetOneCfgValue(model.TXAPIURL)
	// 实例化要请求产品的client对象,clientProfile是可选的
	TxRedisApi, err = tredis.NewClient(credential, region, cpf)
	if err != nil {
		logger.Error("conenct tx cloud redis error: ", err)
		return false
	}
	return true
}

// CVM
var TxCvmApi *cvm.Client

func TxCvmContent() bool {
	var err error
	credential := common.NewCredential(
		mysql.DB.GetOneCfgValue(model.TXSECRETID),
		mysql.DB.GetOneCfgValue(model.TXSECRETKEY),
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	TxCvmApi, err = cvm.NewClient(credential, "", cpf)
	if err != nil {
		logger.Error("conenct tx cloud region error: ", err)
		return false
	}
	return true
}

// Dbrain
var TxDbrainApi *dbbrain.Client

func TxDbrainContent(region string) bool {
	var err error
	credential := common.NewCredential(
		mysql.DB.GetOneCfgValue(model.TXSECRETID),
		mysql.DB.GetOneCfgValue(model.TXSECRETKEY),
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "dbbrain.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	TxDbrainApi, err = dbbrain.NewClient(credential, region, cpf)
	if err != nil {
		logger.Error("conenct tx cloud region error: ", err)
		return false
	}
	return true
}
