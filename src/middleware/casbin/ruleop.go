package casbin

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
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
		logger.Error("create new rule error ", err)
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

func RuleGet() (casbinrule []CasbinRule) {
	mysql.DB.DB.Find(&casbinrule)
	// mysql.DB.DB.Offset(page).Limit(size).Find(&casbinrule)
	return
}
