package model

var (
	DefaultPath   = make(map[string]string)
	PATHCFG       = "/redis-manager/cfg/v1"
	PATHBOARD     = "/redis-manager/board/v1"
	PATHHISTORY   = "/redis-manager/ophistory/v1"
	PATHCODIS     = "/redis-manager/codis/v1"
	PATHCLOUD     = "/redis-manager/cloud/v1"
	PATHCLUSTER   = "/redis-manager/cluster/v1"
	PATHCLI       = "/redis-manager/cli/v1"
	PATHUSER      = "/redis-manager/user/v1"
	PATHRULE      = "/redis-manager/rule/v1"
	PATHAUTH      = "/redis-manager/auth/v1"
	DefaultMethod = make(map[string]string)
	METHODGET     = "GET"
	METHODPOST    = "POST"
	METHODDELETE  = "DELETE"
	METHODPUT     = "PUT"
)

func init() {
	DefaultPath[PATHCFG+"/*"] = "系统配置页面权限"
	DefaultPath[PATHBOARD+"/*"] = "概览页面权限"
	DefaultPath[PATHHISTORY+"/*"] = "历史记录页面权限"
	DefaultPath[PATHCODIS+"/*"] = "Redis集群/Codis页面权限"
	DefaultPath[PATHCLOUD+"/*"] = "Redis集群/腾讯和阿里Redis页面权限"
	DefaultPath[PATHCLUSTER+"/*"] = "Redis集群/Cluster页面权限"
	DefaultPath[PATHCLI+"/*"] = "数据查询页面权限"
	DefaultPath[PATHUSER+"/*"] = "用户管理/用户列表页面权限"
	DefaultPath[PATHRULE+"/*"] = "用户管理/权限管理页面权限"
	DefaultPath[PATHAUTH+"/*"] = "用户修改密码权限"
	DefaultMethod[METHODGET] = "读权限"
	DefaultMethod[METHODPOST] = "写权限"
	DefaultMethod[METHODDELETE] = "删除权限"

}
