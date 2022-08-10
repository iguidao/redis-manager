package opredis

import (
	"sort"

	"github.com/iguidao/redis-manager/src/cfg"
)

func AllKey() []string {
	var keylist []string
	val, num, slowlogok := GetAllKey(0)
	if !slowlogok {
		return nil
	}
	keylist = append(keylist, val...)
	var fornum = 1
	for {
		val, num, slowlogok = GetAllKey(num)
		if !slowlogok {
			sort.Strings(keylist)
			return keylist
		}
		fornum++
		keylist = append(keylist, val...)
		if num == 0 || fornum >= cfg.Get_Info_Int("allkeyfornum") {
			break
		}
	}
	sort.Strings(keylist)
	return keylist
}
