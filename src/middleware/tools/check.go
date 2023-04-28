package tools

import (
	"net"
	"time"
)

func CheckStringInArray(target string, str_array []string) bool {
	for _, element := range str_array {
		if target == element {
			return true
		}
	}
	return false
}

func IsIP(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	} else {
		return true
	}
}

func CheckIpPort(addr string, timeout int) bool {
	conn, err := net.DialTimeout("tcp", addr, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}
