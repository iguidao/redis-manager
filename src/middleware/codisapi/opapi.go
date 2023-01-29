package codisapi

import (
	"encoding/json"
	"log"
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

// 获取集群信息
func CodisTopom(codisurl, cn string) (Topom, bool) {
	var topom Topom
	// log.Println(clusterAuth)
	url := codisurl + "/topom"
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.GetDefault(url, geturi, nil)
	if !httpCode {
		log.Println("获取codis集群信息失败：", httpCode, httpResult)
		return topom, false
	}
	json.Unmarshal([]byte(httpResult), &topom)
	return topom, true
}

// 添加proxy
func CodisProxyUp(curl, cn, clusterAuth, proxyip, port string) bool {
	// log.Println(clusterAuth)
	url := curl + "/api/topom/proxy/create/" + clusterAuth + "/" + proxyip + ":" + port
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		log.Println("添加proxy节点", proxyip, "失败：", httpCode, httpResult)
		return false
	}
	log.Println("添加proxy节点", proxyip, "成功", httpCode, httpResult)
	return true
}

// 添加group组
func CodisAddGroup(curl, cn, clusterAuth string, groupid int) bool {
	groupname := strconv.Itoa(groupid)
	url := curl + "/api/topom/group/create/" + clusterAuth + "/" + groupname
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		log.Println("添加group组", groupid, "失败：", httpCode, httpResult)
		return false
	}
	log.Println("添加group组", groupid, "成功", httpCode, httpResult)
	return true
}

// 添加group中的节点
func CodisGroupUp(groupid int, curl, cn, clusterAuth, hostname, port string) bool {
	groupname := strconv.Itoa(groupid)
	url := curl + "/api/topom/group/add/" + clusterAuth + "/" + groupname + "/" + hostname + ":" + port
	log.Println(url)
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		log.Println("添加group节点", groupid, "失败：", httpCode, httpResult)
		return false
	}
	log.Println("添加group节点", groupid, "成功", httpCode, httpResult)
	return true
}

// 执行 server sync
func CodisServerSync(curl, cn, clusterAuth, hostname, port string) bool {
	url := curl + "/api/topom/group/action/create/" + clusterAuth + "/" + hostname + ":" + port
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		log.Println("Codis机器"+hostname+"sync失败：", httpCode, httpResult)
		return false
	}
	log.Println("Codis机器"+hostname+"sync成功：", httpCode, httpResult)
	return true
}

// 执行group sync
func CodisSync(curl, cn, clusterAuth string) bool {
	url := curl + "/api/topom/sentinels/resync-all/" + clusterAuth
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		log.Println("Codis集群sync失败：", httpCode, httpResult)
		return false
	}
	log.Println("Codis集群sync成功：", httpCode, httpResult)
	return true
}

// 执行rebalance
func CodisRebalance(curl, cn, clusterAuth string) bool {
	url := curl + "/api/topom/slots/rebalance/" + clusterAuth + "/" + "1"
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		log.Println("执行集群Rebalance失败: ", httpCode, httpResult)
		return false
	}
	log.Println("执行集群Rebalance任务开始")
	return true
}

// 获取状态
func CodisInfo(curl, cn string) (TopomStats, bool) {
	var topomstats TopomStats
	// log.Println(clusterAuth)
	url := curl + "/topom/stats"
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.GetDefault(url, geturi, nil)
	if !httpCode {
		log.Println("获取codis集群信息失败：", httpCode, httpResult)
		return topomstats, false
	}
	json.Unmarshal([]byte(httpResult), &topomstats)
	return topomstats, true
}
