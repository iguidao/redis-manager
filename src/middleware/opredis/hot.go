package opredis

import (
	"regexp"
	"strconv"
	"strings"
)

// type kv struct {
// 	Key   string
// 	Value int
// }

func HotKey(serverip string) map[string]int {
	keydic := make(map[string]int)
	monitor, knowtime := TelnetCommond(serverip, "monitor")
	re := regexp.MustCompile("(?m)[\r\n]+^.*\"PING\"|\"INFO\".*$")
	monitor = re.ReplaceAllString(monitor, "")
	str := "(?m)[\r\n]+^.*" + strconv.FormatInt(knowtime+1, 10) + ".*$"
	re = regexp.MustCompile(str)
	monitora := re.FindAllString(monitor, -1)
	for _, v := range monitora {
		vlist := strings.Split(v, " ")
		if len(vlist) > 4 {
			vstring := strings.Replace(vlist[4], " ", "", -1)
			vstring = strings.Replace(vstring, "\n", "", -1)
			vstring = strings.Replace(vstring, "\r", "", -1)
			vstring = strings.Replace(vstring, "\"", "", -1)
			vstring = strings.Replace(vstring, "\\", "", -1)
			if _, ok := keydic[vstring]; ok {
				keydic[vstring] = keydic[vstring] + 1
			}
			keydic[vstring] = 1
		}
	}
	// var ss []kv
	// for k, v := range keydic {
	// 	ss = append(ss, kv{k, v})
	// }
	// sort.Slice(ss, func(i, j int) bool {
	// 	return ss[i].Value > ss[j].Value
	// })
	// for _, v := range ss {
	// 	fmt.Printf("%s, %d\n", v.Key, v.Value)
	// }

	// for i, v := range keydic {
	// 	log.Println("key:", i, "value:", v)
	// }
	// log.Println("wanle", knowtime+1)
	return keydic
}
