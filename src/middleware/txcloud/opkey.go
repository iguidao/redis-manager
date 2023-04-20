package txcloud

import (
	"fmt"
	"time"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	dbbrain "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/dbbrain/v20210527"
	tredis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
)

func TxHostKey(instanceid string) (string, bool) {

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tredis.NewDescribeInstanceMonitorHotKeyRequest()

	request.InstanceId = common.StringPtr(instanceid)
	request.SpanType = common.Int64Ptr(1)

	// 返回的resp是一个DescribeInstanceMonitorHotKeyResponse的实例，与请求对象对应
	response, err := TxRedisApi.DescribeInstanceMonitorHotKey(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", false
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), true
}

func TxProxySlowKey(instanceid, starttime, endtime string) (string, bool) {
	// todaynow := time.Now().Format("2006") + "-" + time.Now().Format("01") + "-" + time.Now().Format("02")
	request := tredis.NewDescribeProxySlowLogRequest()

	request.InstanceId = common.StringPtr(instanceid)
	request.BeginTime = common.StringPtr(starttime)
	request.EndTime = common.StringPtr(endtime)

	// 返回的resp是一个DescribeProxySlowLogResponse的实例，与请求对象对应
	response, err := TxRedisApi.DescribeProxySlowLog(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", false
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), true
}
func TxRedisSlowKey(instanceid, starttime, endtime string) (string, bool) {

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := tredis.NewDescribeSlowLogRequest()

	request.InstanceId = common.StringPtr(instanceid)
	request.BeginTime = common.StringPtr(starttime)
	request.EndTime = common.StringPtr(endtime)

	// 返回的resp是一个DescribeSlowLogResponse的实例，与请求对象对应
	response, err := TxRedisApi.DescribeSlowLog(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", false
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	return response.ToJsonString(), true
}

func TxBigKey(instanceid string) (string, bool) {
	todaynow := time.Now().Format("2006") + "-" + time.Now().Format("01") + "-" + time.Now().Format("02")

	request := dbbrain.NewDescribeRedisTopBigKeysRequest()

	request.InstanceId = common.StringPtr(instanceid)
	request.Date = common.StringPtr(todaynow)
	request.Product = common.StringPtr("redis")

	// 返回的resp是一个DescribeRedisTopBigKeysResponse的实例，与请求对象对应
	response, err := TxDbrainApi.DescribeRedisTopBigKeys(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", false
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包	return response.ToJsonString(), true
	return response.ToJsonString(), true
}
