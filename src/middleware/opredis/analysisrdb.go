package opredis

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"sort"
	"time"

	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/tommy351/rdb-go"
)

func Analysis(filename, serverip string) map[string]interface{} {
	checksize := cfg.Get_Info_Int("checksize")
	stringkeymap := make(map[string]int)
	listkeymap := make(map[string]int)
	hashkeymap := make(map[string]int)
	setkeymap := make(map[string]int)
	zsetkeymap := make(map[string]int)
	resultmap := make(map[string]interface{})
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	parser := rdb.NewParser(file)

	for {
		data, err := parser.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		switch data := data.(type) {
		case *rdb.StringData:
			stringkeysize := len(data.Value)
			if stringkeysize > checksize {
				stringkeymap[data.Key] = stringkeysize
			}
		case *rdb.ListData:
			listkeysize := 0
			for _, value := range data.Value {
				listkeysize = listkeysize + len(value)
			}
			if listkeysize > checksize {
				listkeymap[data.Key] = listkeysize
			}
		case *rdb.HashData:
			hashkeysize := 0
			for _, value := range data.Value {
				hashkeysize = hashkeysize + len(value)
			}
			if hashkeysize > checksize {
				hashkeymap[data.Key] = hashkeysize
			}
		case *rdb.SetData:
			setkeysize := 0
			for _, value := range data.Value {
				setkeysize = setkeysize + len(value)
			}
			if setkeysize > checksize {
				setkeymap[data.Key] = setkeysize
			}

		case *rdb.SortedSetData:
			sortsetkeysize := 0
			for _, value := range data.Value {
				sortsetkeysize = sortsetkeysize + len(value.Value)
			}
			if sortsetkeysize > checksize {
				zsetkeymap[data.Key] = sortsetkeysize
			}
		}
	}

	resultmap["String-Big-Key-Top10"] = SortTopkey(stringkeymap)
	resultmap["List-Big-Key-Top10"] = SortTopkey(listkeymap)
	resultmap["hash-Big-Key-Top10"] = SortTopkey(hashkeymap)
	resultmap["set-Big-Key-Top10"] = SortTopkey(setkeymap)
	resultmap["check-time"] = time.Now().Format("2006-01-02 15:04:05")
	jsonBody, _ := json.Marshal(resultmap)
	_, ok := SetStringKey(serverip, string(jsonBody))
	if ok {
		logger.Info("bigkey ", serverip, "设置成功 ", string(jsonBody))
	}
	return resultmap
}

type keyintperoson struct {
	Name string
	Age  int
}

func SortTopkey(keymap map[string]int) map[string]int {
	var lstPerson []keyintperoson
	resultkeymap := make(map[string]int)
	for k, v := range keymap {
		lstPerson = append(lstPerson, keyintperoson{k, v})
	}

	sort.Slice(lstPerson, func(i, j int) bool {
		return lstPerson[i].Age > lstPerson[j].Age // 降序
	})
	for i, v := range lstPerson {
		if i == 10 {
			break
		}
		resultkeymap[v.Name] = v.Age
	}
	return resultkeymap

}
