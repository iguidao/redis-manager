package model

var (
	DefaultUser     = make(map[string]string)
	USERTYPEADMIN   = "admin"                                                     //管理员身份
	USERTYPEVISITOR = "visitor"                                                   //访客身份
	USERTYPEMEMBER  = "member"                                                    //会员身份
	UserDefault     = [...]string{USERTYPEADMIN, USERTYPEVISITOR, USERTYPEMEMBER} //列出员工身份
)

func init() {
	DefaultUser[USERTYPEADMIN] = "管理员"
	DefaultUser[USERTYPEVISITOR] = "访客"
	DefaultUser[USERTYPEMEMBER] = "会员"
}
