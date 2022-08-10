package opredis

func QueryKey(keyname string) QueryResult {
	keytype, ok := TypeKey(keyname)
	var result QueryResult
	if ok {
		result.Value = Query_value(keytype, keyname)
		ttl, tok := TtlKey(keyname)
		if tok {
			result.Ttl = ttl
		}
		result.Type = keytype
	}
	return result
}

func Query_value(keytype, keyname string) interface{} {
	switch keytype {
	case "string":
		val, stringok := GetStringKey(keyname)
		if stringok {
			return val
		}
	case "list":
		val, listok := GetListKey(keyname)
		if listok {
			return val
		}
	case "hash":
		val, hashok := GetHashKey(keyname)
		if hashok {
			return val
		}
	case "set":
		val, setok := GetSetKey(keyname)
		if setok {
			return val
		}
	case "zset":
		val, zsetok := GetZsetKey(keyname)
		if zsetok {
			return val
		}
	case "none":
		return "Not Found Key"
	default:
		return "This type " + keytype + " key, query is not supported"
	}
	return "Get Key Fail"
}
