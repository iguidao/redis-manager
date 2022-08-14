package opredis

import (
	"sort"
	"sync"
)

var (
	stringkeymap map[string]int64
	listkeymap   map[string]int64
	hashkeymap   map[string]int64
	setkeymap    map[string]int64
	zsetkeymap   map[string]int64
	wg           sync.WaitGroup
)

func BigKey() map[string]interface{} {
	stringkeymap = make(map[string]int64)
	listkeymap = make(map[string]int64)
	hashkeymap = make(map[string]int64)
	setkeymap = make(map[string]int64)
	zsetkeymap = make(map[string]int64)
	resultmap := make(map[string]interface{})
	// val, _ := BgsaveKey()
	val, num, scanok := GetScanKey(0, 1000)
	if !scanok {
		return nil
	}
	wg.Add(1)
	go Countkey(val)
	for {
		val, num, scanok = GetScanKey(num, 1000)
		if !scanok {
			break
		}
		if num != 0 {
			wg.Add(1)
			go Countkey(val)
		} else {
			break
		}
	}
	wg.Wait()
	resultmap["string-Top10"] = Sortkey(stringkeymap)
	resultmap["list-Top10"] = Sortkey(listkeymap)
	resultmap["hash-Top10"] = Sortkey(hashkeymap)
	resultmap["set-Top10"] = Sortkey(setkeymap)
	resultmap["zset-Top10"] = Sortkey(zsetkeymap)
	return resultmap
}

type keyperoson struct {
	Name string
	Age  int64
}

func Sortkey(keymap map[string]int64) map[string]int64 {
	var lstPerson []keyperoson
	resultkeymap := make(map[string]int64)
	for k, v := range keymap {
		lstPerson = append(lstPerson, keyperoson{k, v})
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

func Countkey(keylist []string) {
	cstringkeymap := make(map[string]int64)
	clistkeymap := make(map[string]int64)
	chashkeymap := make(map[string]int64)
	csetkeymap := make(map[string]int64)
	czsetkeymap := make(map[string]int64)
	for _, keyname := range keylist {
		// log.Println(keyname)
		keytype, ok := TypeKey(keyname)
		if !ok {
			continue
		}
		switch keytype {
		case "string":
			val, stringok := SizeStringKey(keyname)
			if stringok {
				cstringkeymap[keyname] = val
			}
		case "list":
			val, listok := SizeListKey(keyname)
			if listok {
				clistkeymap[keyname] = val
			}
		case "hash":
			val, hashok := SizeHashKey(keyname)
			if hashok {
				chashkeymap[keyname] = val
			}
		case "set":
			val, setok := SizeSetKey(keyname)
			if setok {
				csetkeymap[keyname] = val
			}
		case "zset":
			val, zsetok := SizeZsetKey(keyname)
			if zsetok {
				czsetkeymap[keyname] = val
			}
		case "none":
			continue
		default:
			continue
		}
	}
	stringkeymap = AppendMap(stringkeymap, cstringkeymap)
	listkeymap = AppendMap(listkeymap, clistkeymap)
	hashkeymap = AppendMap(hashkeymap, chashkeymap)
	setkeymap = AppendMap(setkeymap, csetkeymap)
	zsetkeymap = AppendMap(zsetkeymap, czsetkeymap)
	defer wg.Done()
}

func AppendMap(result, val map[string]int64) map[string]int64 {
	for i, v := range val {
		if i != "" || v != 0 {
			result[i] = v
		}
	}
	return result
}
