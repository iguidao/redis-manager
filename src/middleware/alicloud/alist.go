package alicloud

import (
	r_kvstore20150101 "github.com/alibabacloud-go/r-kvstore-20150101/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func AliListRegion() (string, bool) {
	describeRegionsRequest := &r_kvstore20150101.DescribeRegionsRequest{}
	runtime := &util.RuntimeOptions{}
	response, err := AliRedisApi.DescribeRegionsWithOptions(describeRegionsRequest, runtime)
	if err != nil {
		logger.Error("Get ali region error: ", err)
		return "", false
	}
	return response.Body.GoString(), true
}

func AliListRedis(region string) (string, bool) {
	describeInstancesRequest := &r_kvstore20150101.DescribeInstancesRequest{
		RegionId: tea.String(region),
	}
	runtime := &util.RuntimeOptions{}
	response, err := AliRedisApi.DescribeInstancesWithOptions(describeInstancesRequest, runtime)
	if err != nil {
		logger.Error("Get ali region error: ", err)
		return "", false
	}
	return response.Body.GoString(), true
}
