package opredis

import (
	"context"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func HotKey(serverip, pw string) map[string]int {
	keydic := make(map[string]int)

	var monitor string
	var knowtime int64
	ch := make(chan string)
	timeout, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go func() {
		monitor, knowtime = TelnetCommond(serverip, "monitor", pw)
		ch <- "done"
	}()

	select {
	case res := <-ch:
		logger.Info("Telnet success: ", res)
	case <-timeout.Done():
		logger.Error("Telnet timout: ", timeout.Err())
	}

	// monitor, knowtime := TelnetCommond(serverip, "monitor")
	// log.Println("monitor, knowtime", monitor, knowtime)
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
			} else {
				keydic[vstring] = 1
			}
		}
	}
	return keydic
}
