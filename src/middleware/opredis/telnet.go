package opredis

import (
	"time"

	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/reiver/go-telnet"
)

func ReaderTelnet(conn *telnet.Conn) (out string, knowtime int64) {
	var result string
	var buffer [100]byte
	recvData := buffer[:]
	knowtime = time.Now().Unix()

	for {
		comtime := time.Now().Unix()
		_, err := conn.Read(recvData)
		if nil != err {
			logger.Error("ReaderTelnet error: ", err)
		}
		result = result + string(recvData)
		if comtime-knowtime > 1 {
			break
		}
	}

	return result, knowtime
}
func SenderTelnet(conn *telnet.Conn, command string) {
	var crlfBuffer [2]byte = [2]byte{'\r', '\n'}
	crlf := crlfBuffer[:]
	conn.Write([]byte(command))
	conn.Write(crlf)
}

func TelnetCommond(ip, command string) (string, int64) {
	conn, err := telnet.DialTo(ip)
	if nil != err {
		logger.Error("TelnetCommond error: ", err)
	}
	defer conn.Close()
	SenderTelnet(conn, command)
	monitor, knowtime := ReaderTelnet(conn)
	return monitor, knowtime
}
