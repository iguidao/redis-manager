package opredis

import (
	"log"
	"strings"
	"time"

	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/model"
	"github.com/iguidao/redis-manager/src/middleware/tools"
)

// codis shrinkage ====================
func DownClusterHost(codisnode model.CodisChangeNode, auth string, topomstats codisapi.TopomStats) (bool, []string) {
	var proxylist []int
	var grouplist []int
	var downlist []string
	delproxynum := codisnode.DelProxy
	delgroupnum := codisnode.DelGroup
	for i := 0; i < delgroupnum; i++ {
		grouplist = append(grouplist, tools.CalculationGroup(grouplist, topomstats))
	}
	for i := 0; i < delproxynum; i++ {
		proxylist = append(proxylist, tools.CalculationProxy(proxylist, topomstats))
	}
	logger.Info("Codis Opnode shrinkage: 将要下线的proxy节点：", proxylist, " 将要下线的group节点：", grouplist)

	for _, v := range proxylist {
		for _, proxy := range topomstats.Proxy.Models {
			if proxy.Id == v {
				logger.Info("Codis Opnode shrinkage: 开始下掉proxy节点：", proxy.ProxyAddr)
				if codisapi.CodisProxyDown(codisnode.Curl, codisnode.ClusterName, auth, proxy.Token) {
					downlist = append(downlist, proxy.ProxyAddr)
				}
			}
		}
	}
	groupamp := make(map[int]int)
	for _, v := range topomstats.Group.Models {
		groupamp[v.Id] = 0
	}
	for _, v := range grouplist {
		delete(groupamp, v)
	}

	for _, v := range grouplist {
		grouphost := make(map[string]string)
		var checktime, slotnumber, slotmv, surplus int
		for _, slot := range topomstats.Slots {
			if slot.GroupId == v {
				slotnumber++
			}
		}
		surplus = slotnumber
		slotavg := slotnumber / (len(topomstats.Group.Models) - len(grouplist))
		for i := range groupamp {
			if slotnumber == surplus {
				groupamp[i] = slotavg
				surplus = surplus - slotavg
				slotmv = slotmv + slotavg
			} else if slotnumber-slotavg > surplus-slotavg {
				groupamp[i] = slotavg
				// slotlist = append(slotlist, slotavg)
				surplus = surplus - slotavg
				slotmv = slotmv + slotavg
			}
		}
		if slotnumber-slotmv > 0 {
			var fornum int
			for key, value := range groupamp {
				groupamp[key] = value + 1
				fornum++
				if fornum > slotnumber-slotmv-1 {
					break
				}
			}
		}

		for newid, mvslot := range groupamp {
			logger.Info("开始迁移分片：", mvslot, "到group节点:", newid)
			if codisapi.CodisSlotMv(codisnode.Curl, codisnode.ClusterName, auth, v, newid, mvslot) {
				for {
					if checktime > 600 {
						logger.Error("Codis Opnode shrinkage: 检查失败，已经超过40min了，还在迁移分片")
						return false, downlist
					}
					if CheckSlotPending(codisnode.Curl, codisnode.ClusterName) {
						break
					}
					checktime++
					time.Sleep(time.Duration(4) * time.Second)
				}
			}
		}

		for _, group := range topomstats.Group.Models {
			if group.Id == v && CheckSlotGroup(v, codisnode.Curl, codisnode.ClusterName) {
				for _, host := range group.Servers {
					grouphost[host.Server] = "salve"
				}
				for _, mastername := range topomstats.Sentinels.Masters {
					for name := range grouphost {
						if mastername == name {
							grouphost[name] = "master"
						}
					}
				}
				checktime = 0
				for {
					if len(grouphost) == 0 {
						break
					}
					if checktime > 120 {
						logger.Error("Codis Opnode shrinkage: 下节点group失败，已经超过10min了，还没下完")
						return false, downlist
					}

					for host, rule := range grouphost {
						if len(grouphost) == 1 {
							logger.Info("Codis Opnode shrinkage: 开始下掉group ", group.Id, " 节点的机器：", host)
							if codisapi.CodisGroupDown(v, codisnode.Curl, codisnode.ClusterName, auth, host) {
								delete(grouphost, host)
								downlist = append(downlist, host)
							}
						}
						if len(grouphost) > 1 && rule != "master" {
							logger.Info("Codis Opnode shrinkage: 开始下掉group ", group.Id, " 节点的机器：", host)
							if codisapi.CodisGroupDown(v, codisnode.Curl, codisnode.ClusterName, auth, host) {
								delete(grouphost, host)
								downlist = append(downlist, host)
							}
						}
					}
					checktime++
					time.Sleep(time.Duration(5) * time.Second)
				}
				if CheckGroup(v, codisnode.Curl, codisnode.ClusterName) {
					codisapi.CodisRmGroup(codisnode.Curl, codisnode.ClusterName, auth, v)
					codisapi.CodisSync(codisnode.Curl, codisnode.ClusterName, auth)
				}
			}
		}
	}
	return true, downlist
}

//codis dilation ======================

// codis集群添加某个机器
func UpClusterHost(codisnode model.CodisChangeNode, topom codisapi.Topom, auth string) (bool, []string) {
	logger.Info("Codis Opnode dilatation: 将要上限的proxy节点：", codisnode.AddProxy, " 将要上限的group节点：", codisnode.AddServer)
	var uplist []string
	var grouplist []int
	serverlist := strings.Split(codisnode.AddServer, ",")
	proxylist := strings.Split(codisnode.AddProxy, ",")
	// serverlist := codisnode.AddServer
	// proxylist := codisnode.AddProxy
	for _, v := range proxylist {
		logger.Info("Codis Opnode dilatation: 开始上线proxy节点：", v, "ip: ", v)
		if codisapi.CodisProxyUp(codisnode.Curl, codisnode.ClusterName, auth, v) {
			uplist = append(uplist, v)
		}
	}
	groupid := GetNextGroupId(topom) + 1
	logger.Info("Codis Opnode dilatation: 最大 group id: ", groupid)
	for i := 0; i < len(serverlist)/2; i++ {
		codisapi.CodisAddGroup(codisnode.Curl, codisnode.ClusterName, auth, groupid)
		grouplist = append(grouplist, groupid)
		groupid++
	}
	logger.Info("Codis Opnode dilatation: group list:", grouplist)
	// for i := 0; i < model.Gn; i++ {
	for _, gid := range grouplist {
		var count int = 1
		for _, ip := range serverlist {
			serverlist = tools.DeleteListString(ip, serverlist)
			log.Println("Codis Opnode dilatation: 开始上线group节点ip:", ip)
			if codisapi.CodisGroupUp(gid, codisnode.Curl, codisnode.ClusterName, auth, ip) {
				if codisapi.CodisServerSync(codisnode.Curl, codisnode.ClusterName, auth, ip) {
					uplist = append(uplist, ip)
				}
			}
			if count == 2 {
				break
			} else {
				count++
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
	// for _, v := range codisnode.AddServer {
	// 	logger.Info("Codis Opnode dilatation: 开始上线group节点：", v)
	// 	for _, gid := range grouplist {
	// 		grouplist = tools.DeleteListint(gid, grouplist)
	// 		for _, ip := range v.List {
	// 			if codisapi.CodisGroupUp(gid, codisnode.Curl, codisnode.ClusterName, auth, ip, v.Port) {
	// 				if codisapi.CodisServerSync(codisnode.Curl, codisnode.ClusterName, auth, ip, v.Port) {
	// 					uplist = append(uplist, ip)
	// 				}
	// 			}
	// 		}
	// 		time.Sleep(time.Duration(1) * time.Second)
	// 		break
	// 	}

	// }

	time.Sleep(time.Duration(5) * time.Second)
	codisapi.CodisSync(codisnode.Curl, codisnode.ClusterName, auth)
	return true, uplist
}

// other

// get codis group id
func GetNextGroupId(topom codisapi.Topom) int {
	var list []int
	for _, v := range topom.Stats.Group.Models {
		list = append(list, v.Id)
	}
	max := tools.CalculationArrMax(list)
	return max
}

// Codis cluster rebalance
func CodisRebalanceAll(curl, cn, auth string) bool {
	var checktime int
	if codisapi.CodisRebalance(curl, cn, auth) {
		for {
			if checktime > 120 {
				logger.Info("检查失败，已经超过10min了，还在Rebalance")
				return false
			}
			if CheckSlotPending(curl, cn) {
				break
			}
			checktime++
			time.Sleep(time.Duration(5) * time.Second)
		}
		return true
	}
	return false
}

// codis pending check
func CheckSlotPending(curl, cn string) bool {
	topomstats, ok := codisapi.CodisInfo(curl, cn)
	if !ok {
		logger.Info("获取codis的slot迁移状态失败!")
		return false
	}
	for _, v := range topomstats.Slots {
		if v.Action.State == "pending" || v.Action.State == "migrating" {
			return false
		}
	}
	return true
}

// 检查group的slot
func CheckSlotGroup(groupid int, Curl, ClusterName string) bool {
	topomstats, ok := codisapi.CodisInfo(Curl, ClusterName)
	if !ok {
		logger.Info("获取codis的slot状态失败!")
		return false
	}
	for _, v := range topomstats.Slots {
		if v.GroupId == groupid {
			return false
		}
	}
	return true
}

// 检查group状态
func CheckGroup(groupid int, Curl, ClusterName string) bool {
	topomstats, ok := codisapi.CodisInfo(Curl, ClusterName)
	if !ok {
		logger.Info("获取codis的slot状态失败!")
		return false
	}
	for _, v := range topomstats.Group.Models {
		if v.Id == groupid && len(v.Servers) == 0 {
			return true
		}
	}
	return false
}
