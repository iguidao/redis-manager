package opredis

func DeleteKey(keyname string) string {
	val, stringok := DelKey(keyname)
	if stringok {
		if val == 1 {
			return "删除成功"
		}
		if val == -1 {
			return "删除失败"
		}
		if val == 0 {
			return "没有这个key"
		}
	}
	return "删除失败"
}

func CDeleteKey(keyname string) string {
	val, stringok := CDelKey(keyname)
	if stringok {
		if val == 1 {
			return "删除成功"
		}
		if val == -1 {
			return "删除失败"
		}
		if val == 0 {
			return "没有这个key"
		}
	}
	return "删除失败"
}
