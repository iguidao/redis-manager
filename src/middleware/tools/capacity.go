package tools

import (
	"fmt"
	"strconv"

	"github.com/iguidao/redis-manager/src/middleware/codisapi"
)

func CapacityGroup(gn int, topom codisapi.Topom) bool {
	//计算内存大小
	var maxmemory, usememory, onemax, oneuse, maxnum, usenum int
	for _, v := range topom.Stats.Group.Stats {
		max, err := strconv.Atoi(v.Stats.Maxmemory)
		if err == nil {
			maxmemory = maxmemory + max
			maxnum = maxnum + 1
		}
		use, err := strconv.Atoi(v.Stats.Used_memory)
		if err == nil {
			usememory = usememory + use
			usenum = usenum + 1
		}
	}
	onemax = maxmemory / maxnum
	oneuse = usememory / usenum
	fenzi := ((oneuse*gn)/(len(topom.Stats.Group.Models)-gn) + oneuse)
	fenmu := onemax
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(fenzi)/float64(fenmu)), 64)
	if val < 0.8 {
		return true
	}
	return false

}

func CapacityProxy(pn int, topom codisapi.Topom) bool {
	// 计算proxy的qps
	var maxqps int
	for _, v := range topom.Stats.Proxy.Stats {
		maxqps = maxqps + v.Stats.Ops.Qps
		// log.Println(i, v.Stats.Ops.Qps)
	}
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(maxqps)/float64(len(topom.Stats.Proxy.Models)-pn)), 64)
	if val < 40000 {
		return true
	}
	return false
}
