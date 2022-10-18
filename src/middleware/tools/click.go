package tools

import (
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/opredis"
)

func BigKeyClick(clusterName, groupName, keyname string) (string, int) {
	// keyname := "Click-Bigkey-" + clusterName + "-" + groupName
	num, ok := opredis.IncrStringKey(keyname)
	if !ok {
		return "大key分析计数出错了执行出现了问题，请找sre服务台！！！", 0
	}
	if num == 1 {
		if !opredis.ExpireKey(keyname, cfg.Get_Info_Int("locktime")) {
			return "大key分析计数时间出错了执行出现了问题，请找sre服务台！！！", 0
		}
	}
	if num == 5 {
		return "成功激发隐藏功能，开始针对集群：" + clusterName + " 的组：" + groupName + " 执行大key分析操作！！！请等待，如若超过10min未查出结果，请找SRE服务台..", 1
	}
	if num > 5 {
		return "嗯？已经在执行后台大key分析了，还点？别着急，请等待10min！，如若超过10min未查出结果，请找SRE服务台..", 2
	}
	return "据听说1分钟点击5次查询，可以生成最新的大key分析数据", 3
}
