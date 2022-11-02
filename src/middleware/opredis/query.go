package opredis

import "strings"

func QueryKey(keyname string) QueryResult {
	keytype, ok := TypeKey(keyname)
	var result QueryResult
	if ok {
		result.Value, result.Len = Query_value(keytype, keyname)
		ttl, tok := TtlKey(keyname)
		if tok {
			result.Ttl = ttl
		}
		result.Type = keytype
	}
	return result
}

func Query_value(keytype, keyname string) (interface{}, int) {
	switch keytype {
	case "string":
		val, stringok := GetStringKey(keyname)
		if stringok {
			return val, strings.Count(val, "")
		}
	case "list":
		val, listok := GetListKey(keyname)
		if listok {
			return val, len(val)
		}
	case "hash":
		val, hashok := GetHashKey(keyname)
		if hashok {
			return val, len(val)
		}
	case "set":
		val, setok := GetSetKey(keyname)
		if setok {
			return val, len(val)
		}
	case "zset":
		val, zsetok := GetZsetKey(keyname)
		if zsetok {
			return val, len(val)
		}
	case "none":
		return "Not Found Key", 0
	default:
		return "This type " + keytype + " key, query is not supported", 0
	}
	return "Get Key Fail", 0
}
