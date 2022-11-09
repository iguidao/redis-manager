package codisapi

import (
	"encoding/json"
	"strconv"

	"github.com/iguidao/redis-manager/src/middleware/httpapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func GeClusterList(curl string) []string {
	var listresult []string
	codisurl := curl + "/list"
	UriData := map[string]string{}
	httpCode, httpResult := httpapi.GetDefault(codisurl, UriData, nil)
	if httpCode {
		json.Unmarshal([]byte(httpResult), &listresult)
	}
	return listresult
}

func GetProxy(curl, ClusterName string) []string {
	var resultforward ResultForward
	var result []string
	codisurl := curl + "/topom"
	UriData := map[string]string{
		"forward": ClusterName,
	}
	httpCode, httpResult := httpapi.GetDefault(codisurl, UriData, nil)
	if httpCode {

		json.Unmarshal([]byte(httpResult), &resultforward)
	} else {
		logger.Error("codis connect fail: ", httpResult)
	}
	for _, v := range resultforward.Stats.Proxy.Models {
		result = append(result, v.ProxyAddr)
	}
	return result
}

func GetGroup(curl, ClusterName string) []string {
	var resultforward ResultForward
	var result []string
	codisurl := curl + "/topom"
	UriData := map[string]string{
		"forward": ClusterName,
	}
	httpCode, httpResult := httpapi.GetDefault(codisurl, UriData, nil)
	if httpCode {
		json.Unmarshal([]byte(httpResult), &resultforward)
	} else {
		logger.Error("codis connect fail: ", httpResult)
		return result
	}
	if resultforward.Stats.Sentinels.Masters == nil {
		return result
	}
	for i := range resultforward.Stats.Sentinels.Masters.(map[string]interface{}) {
		// groupname := "Group" + i + "-" + v.(string)
		result = append(result, i)
	}
	return result
}

func GetMaster(curl, ClusterName string, id string) string {
	var resultforward ResultForward
	var result string
	codisurl := curl + "/topom"
	UriData := map[string]string{
		"forward": ClusterName,
	}
	httpCode, httpResult := httpapi.GetDefault(codisurl, UriData, nil)
	if httpCode {
		json.Unmarshal([]byte(httpResult), &resultforward)
	} else {
		logger.Error("codis connect fail: ", httpResult)
		return result
	}
	if resultforward.Stats.Sentinels.Masters == nil {
		return result
	}
	for i, v := range resultforward.Stats.Sentinels.Masters.(map[string]interface{}) {
		if i == id {
			result = v.(string)
		}
	}
	return result
}

func GetSlave(curl, ClusterName string, id string) string {
	var resultforward ResultForward
	var result string
	var redis_list []string
	codisurl := curl + "/topom"
	UriData := map[string]string{
		"forward": ClusterName,
	}
	httpCode, httpResult := httpapi.GetDefault(codisurl, UriData, nil)
	if httpCode {
		json.Unmarshal([]byte(httpResult), &resultforward)
	} else {
		logger.Error("codis connect fail: ", httpResult)
		return result
	}
	for _, v := range resultforward.Stats.Group.Models {
		if strconv.Itoa(v.Id) == id {
			for _, server := range v.Servers {
				redis_list = append(redis_list, server.Server)
			}
		}
	}
	if resultforward.Stats.Sentinels.Masters == nil {
		return result
	}
	for i, v := range resultforward.Stats.Sentinels.Masters.(map[string]interface{}) {
		if i == id {
			result = v.(string)
		}
	}
	for _, v := range redis_list {
		if v != result {
			result = v
		}
	}
	return result
}
