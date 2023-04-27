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

// 获取集群信息
func CodisTopom(codisurl, cn string) (Topom, bool) {
	var topom Topom
	url := codisurl + "/topom"
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.GetDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("获取codis集群信息失败：", httpCode, httpResult)
		return topom, false
	}
	json.Unmarshal([]byte(httpResult), &topom)
	return topom, true
}

// 添加proxy
func CodisProxyUp(curl, cn, clusterAuth, proxyip string) bool {
	url := curl + "/api/topom/proxy/create/" + clusterAuth + "/" + proxyip
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("添加proxy节点", proxyip, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("添加proxy节点", proxyip, "成功", httpCode, httpResult)
	return true
}

// 下掉proxy
func CodisProxyDown(curl, cn, clusterAuth, proxyid string) bool {
	url := curl + "/api/topom/proxy/remove/" + clusterAuth + "/" + proxyid + "/0"
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("删除proxy节点", proxyid, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("删除proxy节点", proxyid, "成功", httpCode, httpResult)
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
		logger.Error("添加group组", groupid, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("添加group组", groupid, "成功", httpCode, httpResult)
	return true
}

// 添加group中的节点
func CodisGroupUp(groupid int, curl, cn, clusterAuth, ip string) bool {
	groupname := strconv.Itoa(groupid)
	url := curl + "/api/topom/group/add/" + clusterAuth + "/" + groupname + "/" + ip
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("添加group节点", groupid, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("添加group节点", groupid, "成功", httpCode, httpResult)
	return true
}

// 删除group中的节点
func CodisGroupDown(groupid int, curl, cn, clusterAuth, hostname string) bool {
	groupname := strconv.Itoa(groupid)
	url := curl + "/api/topom/group/del/" + clusterAuth + "/" + groupname + "/" + hostname
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("删除group节点", groupid, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("删除group节点", groupid, "成功", httpCode, httpResult)
	return true
}

// 删除group组
func CodisRmGroup(curl, cn, clusterAuth string, groupid int) bool {
	groupname := strconv.Itoa(groupid)
	url := curl + "/api/topom/group/remove/" + clusterAuth + "/" + groupname
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("删除group组", groupid, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("删除group组", groupid, "成功", httpCode, httpResult)
	return true
}

// 执行 server sync
func CodisServerSync(curl, cn, clusterAuth, ip string) bool {
	url := curl + "/api/topom/group/action/create/" + clusterAuth + "/" + ip
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("Codis机器"+ip+"sync失败：", httpCode, httpResult)
		return false
	}
	logger.Info("Codis机器"+ip+"sync成功：", httpCode, httpResult)
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
		logger.Error("Codis集群sync失败：", httpCode, httpResult)
		return false
	}
	logger.Info("Codis集群sync成功：", httpCode, httpResult)
	return true
}

// 迁移slot
func CodisSlotMv(curl, cn, clusterAuth string, oldid, newid, slotnumber int) bool {
	oldgroup := strconv.Itoa(oldid)
	newgroup := strconv.Itoa(newid)
	slotstring := strconv.Itoa(slotnumber)
	url := curl + "/api/topom/slots/action/create-some/" + clusterAuth + "/" + oldgroup + "/" + newgroup + "/" + slotstring
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.PutDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("迁移slot：", slotnumber, "从", oldid, "到", newid, "失败：", httpCode, httpResult)
		return false
	}
	logger.Info("迁移slot：", slotnumber, "从", oldid, "到", newid, "任务开始")
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
		logger.Error("执行集群Rebalance失败: ", httpCode, httpResult)
		return false
	}
	logger.Info("执行集群Rebalance任务开始")
	return true
}

// 获取状态
func CodisInfo(curl, cn string) (TopomStats, bool) {
	var topomstats TopomStats
	url := curl + "/topom/stats"
	geturi := map[string]string{
		"forward": cn,
	}
	httpCode, httpResult := httpapi.GetDefault(url, geturi, nil)
	if !httpCode {
		logger.Error("获取codis集群信息失败：", httpCode, httpResult)
		return topomstats, false
	}
	json.Unmarshal([]byte(httpResult), &topomstats)
	return topomstats, true
}
