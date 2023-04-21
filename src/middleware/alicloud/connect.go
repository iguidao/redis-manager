package alicloud

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

var AliRedisApi *r_kvstore20150101.Client

func CreateClient(accessKeyId *string, accessKeySecret *string) (_result *r_kvstore20150101.Client, _err error) {
	config := &openapi.Config{
		// 必填，您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 必填，您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}

	// 访问的域名
	config.Endpoint = tea.String(mysql.DB.GetOneCfgValue(model.ALIAPIURL))
	_result = &r_kvstore20150101.Client{}
	_result, _err = r_kvstore20150101.NewClient(config)
	return _result, _err
}

func AliRedisContent() bool {
	var err error
	// 工程代码泄露可能会导致AccessKey泄露，并威胁账号下所有资源的安全性。以下代码示例仅供参考，建议使用更安全的 STS 方式，更多鉴权访问方式请参见：https://help.aliyun.com/document_detail/378661.html
	AliRedisApi, err = CreateClient(tea.String(mysql.DB.GetOneCfgValue(model.ALIACCESSKEYID)), tea.String(mysql.DB.GetOneCfgValue(model.ALIALIACCESSKEYSECRET)))
	if err != nil {
		logger.Error("conenct ali cloud redis error: ", err)
		return false
	}
	return true
}
