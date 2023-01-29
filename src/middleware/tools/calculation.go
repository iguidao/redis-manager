package tools

import "github.com/iguidao/redis-manager/src/middleware/codisapi"

func CalculationGroup(slist []int, topomstats codisapi.TopomStats) int {
	var arrint []int
	for _, v := range topomstats.Group.Models {
		if !CheckInListInt(v.Id, slist) {
			arrint = append(arrint, v.Id)
		}

	}
	return CalculationArrMax(arrint)
}
func CalculationProxy(slist []int, topomstats codisapi.TopomStats) int {
	var arrint []int
	for _, v := range topomstats.Proxy.Models {
		if !CheckInListInt(v.Id, slist) {
			arrint = append(arrint, v.Id)
		}
	}
	return CalculationArrMax(arrint)
}
func CalculationArrMax(arrint []int) (max int) {
	max = arrint[0]
	for _, v := range arrint {
		if v > max {
			max = v
		}
	}
	return
}

func CheckInListInt(val int, slist []int) bool {
	for _, v := range slist {
		if v == val {
			return true
		}
	}
	return false
}

func CheckInListString(val string, slist []string) bool {
	for _, v := range slist {
		if v == val {
			return true
		}
	}
	return false
}

func DeleteListString(val string, slist []string) []string {
	var resultlist []string
	for _, v := range slist {
		if v != val {
			resultlist = append(resultlist, v)
		}
	}
	return resultlist
}
