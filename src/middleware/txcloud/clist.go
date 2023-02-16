package txcloud

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	tredis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
)

func TxListRedis() (string, bool) {
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tredis.NewDescribeInstancesRequest()
	// 返回的resp是一个DescribeInstancesResponse的实例，与请求对象对应
	response, err := TAPI.DescribeInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		logger.Error("An API error has returned: ", err)
		return "", false
	}
	if err != nil {
		logger.Error("List Tx Cloud Redis Error: ", err)
		return "", false
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), true
}
