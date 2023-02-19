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
	// casbinrule := &CasbinRule{
	// 	Ptype: "p",
	// 	V0:    identity,
	// 	V1:    path,
	// 	V2:    method,
	// }
	// result := DB.Create(&casbinrule)
	// if result.Error != nil {
	// 	log.Println("create new article error ", false)
	// 	return false
	// }
	// return true
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
	// log.Println("asd")
	// list := Enforcer.GetPolicy()
	// log.Println("list: ", list)
	// for _, vlist := range list {
	// 	log.Println("list range:", vlist)
	// 	for _, v := range vlist {
	// 		log.Println("value: , ", v)
	// 	}
	// }
}
