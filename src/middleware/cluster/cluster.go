package cluster

import (
	"strconv"
	"strings"

	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
)

func WriteCluster(clusterid int, v string) {
	var flags string
	var masterid string
	var linkstate string
	var slotrange string
	var slotnumber int
	nodelist := strings.Split(v, " ")
	nodeid := nodelist[0]
	nodeaddress := strings.Split(nodelist[1], "@")
	ip := strings.Split(nodeaddress[0], ":")[0]
	pord := strings.Split(nodeaddress[0], ":")[1]
	flags_info := nodelist[2]
	flags_list := strings.Split(flags_info, ",")
	if len(flags_list) == 1 {
		if flags_list[0] == "master" {
			flags = "master"
			linkstate = nodelist[7]
			slotrange = nodelist[8]
			slotnumber = getslotnumber(slotrange)
		} else {
			flags = "slave"
			masterid = nodelist[3]
		}
	} else {
		flags = flags_list[1]
		if flags == "master" {
			linkstate = nodelist[7]
			slotrange = nodelist[8]
			slotnumber = getslotnumber(slotrange)
		}
	}
	id, ok := mysql.DB.AddClusterNode(nodeid, ip, pord, flags, masterid, linkstate, slotrange, clusterid, slotnumber)
	if ok {
		logger.Info("write ", ip, " redis to mysql ok: ", id, "nodeid: ", nodeid)
	} else {
		logger.Error("write ", ip, " redis to mysql false: ", id, "nodeid: ", nodeid)
	}
}

func getslotnumber(slotrange string) int {
	var err error
	var start, end, slot int
	num := strings.Split(slotrange, ",")
	start, err = strconv.Atoi(num[0])
	if err != nil {
		return slot
	}
	end, err = strconv.Atoi(num[1])
	if err != nil {
		return slot
	}
	slot = end - start
	return slot
}
