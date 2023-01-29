package opredis

import (
	"log"
	"time"

	"github.com/iguidao/redis-manager/src/middleware/codisapi"
	"github.com/iguidao/redis-manager/src/middleware/tools"
	hstruct "github.com/iguidao/redis-manager/src/rhttp/v1"
)

func Cdilatationn(codisnode hstruct.CodisNode, auth string, topom codisapi.Topom) {
	uok, _ := UpClusterHost(codisnode, topom, auth)
	if !uok {
		panic("codis add host fails")
	}
	time.Sleep(time.Duration(5) * time.Second)
	if !CodisRebalanceAll(codisnode.Curl, codisnode.ClusterName, auth) {
		panic("codis Reabalance fails")
	}
}

// codis集群添加某个机器
func UpClusterHost(codisnode hstruct.CodisNode, topom codisapi.Topom, auth string) (bool, []string) {
	log.Println("将要上限的proxy节点：", codisnode.Proxy.List, " 将要上限的group节点：", codisnode.Group)
	var uplist []string
	var grouplist []int
	for _, v := range codisnode.Proxy.List {
		log.Println("开始上线proxy节点：", v, "ip: ", v)
		if codisapi.CodisProxyUp(codisnode.Curl, codisnode.ClusterName, auth, v, codisnode.Proxy.Port) {
			uplist = append(uplist, v)
		}
	}
	groupid := GetNextGroupId(topom) + 1
	log.Println("max group id: ", groupid)
	for i := 0; i < len(codisnode.Group)/2; i++ {
		codisapi.CodisAddGroup(codisnode.Curl, codisnode.ClusterName, auth, groupid)
		grouplist = append(grouplist, groupid)
		groupid++
	}
	log.Println("group list:", grouplist)
	// for i := 0; i < model.Gn; i++ {

	for _, gid := range grouplist {
		var count int = 1
		for _, v := range servernew {
			log.Println("开始上线group节点：", v, "ip: ", v)
			if codisapi.CodisGroupUp(gid, curl, cn, auth, v, sport) {
				if codisapi.CodisServerSync(curl, cn, auth, v, sport) {
					uplist = append(uplist, v)
				}
			}
			if count == 2 {
				break
			} else {
				servernew = tools.DeleteListString(v, servernew)
				count++
			}
			time.Sleep(time.Duration(1) * time.Second)
		}
	}
	time.Sleep(time.Duration(5) * time.Second)
	codisapi.CodisSync(curl, cn, auth)
	return true, uplist
}

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
				log.Println("检查失败，已经超过10min了，还在Rebalance")
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

func CheckSlotPending(curl, cn string) bool {
	topomstats, ok := codisapi.CodisInfo(curl, cn)
	if !ok {
		log.Println("获取codis的slot迁移状态失败!")
		return false
	}
	for _, v := range topomstats.Slots {
		if v.Action.State == "pending" || v.Action.State == "migrating" {
			return false
		}
	}
	return true
}
