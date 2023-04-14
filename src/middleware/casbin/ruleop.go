package casbin

import (
	"log"
)

func RuleCheck(identity, path, method string) bool {
	// 简单认证
	res, err := Enforcer.Enforce(identity, path, method)
	if err != nil {
		return false
	}
	return res
}

func RuleAdd(identity, path, method string) bool {
	res, err := Enforcer.AddPolicy(identity, path, method)
	if err != nil {
		log.Println("create new article error ", err)
		return false
	}
	return res
}

func RuleDel(identity, path, method string) bool {
	res, err := Enforcer.RemovePolicy(identity, path, method)
	if err != nil {
		return false
	}
	return res
}

func RuleGet(page, size int) (casbinrule []CasbinRule) {
	DB.Offset(page).Limit(size).Find(&casbinrule)
	return
}
