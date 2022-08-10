package codisapi

import (
	"encoding/json"
	"strconv"

	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/httpapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func GetList() []string {
	var listresult []string
	codisurl := cfg.Get_Local("codisurl") + "/list"
	UriData := map[string]string{}
	httpCode, httpResult := httpapi.GetDefault(codisurl, UriData, nil)
	if httpCode {
		json.Unmarshal([]byte(httpResult), &listresult)
	}
	return listresult
}

func GetProxy(ClusterName string) []string {
	var resultforward ResultForward
	var result []string
	codisurl := cfg.Get_Local("codisurl") + "/topom"
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

func GetGroup(ClusterName string) []string {
	var resultforward ResultForward
	var result []string
	codisurl := cfg.Get_Local("codisurl") + "/topom"
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
	for i := range resultforward.Stats.Sentinels.Masters.(map[string]interface{}) {
		// groupname := "Group" + i + "-" + v.(string)
		result = append(result, i)
	}
	return result
}

func GetMaster(ClusterName string, id string) string {
	var resultforward ResultForward
	var result string
	codisurl := cfg.Get_Local("codisurl") + "/topom"
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
	for i, v := range resultforward.Stats.Sentinels.Masters.(map[string]interface{}) {
		if i == id {
			result = v.(string)
		}
	}
	return result
}

func GetSlave(ClusterName string, id string) string {
	var resultforward ResultForward
	var result string
	codisurl := cfg.Get_Local("codisurl") + "/topom"
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
				if server.Action.State == "synced" {
					result = server.Server
				}
			}
		}
	}
	return result
}
