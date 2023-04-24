package mysql

import (
	"time"

	"gorm.io/gorm"
)

// base
type Base struct {
	ID        int       `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//创建索引`sql:"index"`
}

// 用户
type UserInfo struct {
	Base
	UserName string `gorm:"not null;index;unique"`
	Password string `gorm:"not null"`
	Email    string `gorm:"type:varchar(255)"`
	UserType string `gorm:"not null;index;type:varchar(50)"` //admin 管理员；visitor 访客；staff 员工
	Enable   bool   // 0是封禁False，1是可登录True
}

// 用户组
type UserGroup struct {
	Base
	GroupName     string `gorm:"not null;index;unique"`
	GroupDescribe string `gorm:"type:varchar(255)"`
	GroupType     string `gorm:"not null;index;type:varchar(50)"` //admin 管理组；visitor 访客组；staff 员工组
}

// 组与用户的关系
type GroupContain struct {
	Base
	GroupId int `gorm:"not null;index"`
	UserId  int `gorm:"not null;index"`
}

// 集群信息
type ClusterInfo struct {
	Base
	Name     string `gorm:"not null;index;unique"`
	Nodes    string `gorm:"type:varchar(255)"` //ip:port,ip:port
	Password string `gorm:"type:varchar(255)"`
}

// node信息
type ClusterNode struct {
	Base
	CluserId   int    `gorm:"not null;index"`          //集群ID
	NodeId     string `gorm:"type:varchar(50);unique"` //node的ID
	Ip         string `gorm:"type:varchar(50)"`        //node的IP
	Port       string `gorm:"type:varchar(25)"`        //node的端口
	Flags      string `gorm:"type:varchar(50)"`        //node的身份
	MasterId   string `gorm:"type:varchar(50)"`        //如果是从的话master的ID
	LinkState  string `gorm:"type:varchar(50)"`        //链接状态
	RunStatus  bool   `gorm:"type:varchar(25)"`        //运行状态
	SlotRange  string `gorm:"type:varchar(50)"`        //slot区间
	SlotNumber int    `gorm:"type:varchar(25)"`        //solt个数
}

// cloud redis信息
type CloudInfo struct {
	Base
	Cloud            string `gorm:"type:varchar(10)"`      //云厂商
	InstanceId       string `gorm:"not null;index;unique"` //实例ID
	InstanceName     string `gorm:"type:varchar(100)"`     //实例名称
	PrivateIp        string `gorm:"type:varchar(20)"`      //内网IP
	Port             int    `gorm:"type:varchar(10)"`      //端口
	Region           string `gorm:"type:varchar(20)"`      //region
	Createtime       string `gorm:"type:varchar(20)"`      //创建时间
	Size             int    `gorm:"type:varchar(10)"`      //实例大小
	InstanceStatus   string `gorm:"type:varchar(10)"`      //实例状态
	RedisShardSize   int    `gorm:"type:varchar(10)"`      //分片大小
	RedisShardNum    int    `gorm:"type:varchar(10)"`      //分练数量
	RedisReplicasNum int    `gorm:"type:varchar(10)"`      //副本个数
	NoAuth           bool   `gorm:"type:varchar(10)"`      //是否需要密码
	PublicIp         string `gorm:"type:varchar(20)"`      //外网IP
	Password         string `gorm:"type:varchar(50)"`      //密码
}

// codis信息
type CodisInfo struct {
	Base
	Curl  string `gorm:"not null;index;unique"`
	Cname string `gorm:"type:varchar(50)"`
}

// config信息
type Rconfig struct {
	Base
	Name  string `gorm:"type:varchar(255)"`
	Key   string `gorm:"not null:index:primary_key;unique"`
	Value string `gorm:"type:varchar(255)"`
}

// 操作历史
type OpHistory struct {
	Base
	UserId   int    `gorm:"not null;index"`
	OpInfo   string `gorm:"type:varchar(100)"` // 操作动作
	OpParams string `gorm:"type:text"`         //操作参数属组或者对象
}

type Tabler interface {
	TableName() string
}

func (UserInfo) TableName() string {
	return "user_info"
}

func (UserGroup) TableName() string {
	return "user_group"
}

func (GroupContain) TableName() string {
	return "group_contain"
}

func (CloudInfo) TableName() string {
	return "cloud_info"
}

func (ClusterInfo) TableName() string {
	return "cluster_info"
}
func (CodisInfo) TableName() string {
	return "codis_info"
}

func (ClusterNode) TableName() string {
	return "cluster_node"
}

func (OpHistory) TableName() string {
	return "op_history"
}

func (Rconfig) TableName() string {
	return "rconfig"
}
