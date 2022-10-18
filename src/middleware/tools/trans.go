package tools

import (
	"encoding/json"

	"github.com/iguidao/redis-manager/src/middleware/logger"
)

func JsonToMap(jsonstr string) map[string]interface{} {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonstr), &m)
	if err != nil {
		logger.Error("json to map fail: ", jsonstr, " error: ", err)
		return nil
	}
	return m
}
