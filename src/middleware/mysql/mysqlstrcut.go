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

//用户
type UserInfo struct {
	Base
	UserName string `gorm:"not null;index"`
	Password string `gorm:"not null"`
	Email    string `gorm:"type:varchar(255)"`
	Mobile   string `gorm:"type:varchar(20)"`
	UserType string `gorm:"not null;index;type:varchar(50)"` //admin 管理员；visitor 访客；staff 员工
}

//用户组
type UserGroup struct {
	Base
	GroupName     string `gorm:"not null;index"`
	GroupDescribe string `gorm:"type:varchar(255)"`
	GroupType     string `gorm:"not null;index;type:varchar(50)"` //admin 管理组；visitor 访客组；staff 员工组
}

//组与用户的关系
type GroupContain struct {
	Base
	GroupId int `gorm:"not null;index"`
	UserId  int `gorm:"not null;index"`
}

//集群信息
type ClusterInfo struct {
	Base
	GroupId              int    `gorm:"not null;index"`
	UserId               int    `gorm:"not null;index"`
	ClusterName          string `gorm:"not null;index"`
	RedisNodes           string `gorm:"type:varchar(255)"`
	ClusterNotes         string `gorm:"type:varchar(255)"`
	ClusterMode          string `gorm:"type:varchar(25)"` // 集群(Cluster)；单点(Single)；哨兵(Sentinel)
	ClusterOs            string `gorm:"type:varchar(255)"`
	ClusterVersion       string `gorm:"type:varchar(25)"`
	Initialized          bool
	Clusterstate         string `gorm:"type:varchar(50)"`
	ClusterSlotsAssigned int
	ClusterSlotsOk       int
	ClusterNodes         int
	RedisPassword        string `gorm:"type:varchar(255)"`
	Environment          string `gorm:"type:varchar(50)"` // 主机 Machine；容器 Container
	From                 string `gorm:"type:varchar(50)"` //导入Import；平台创建Self
}

// node信息
type RedisNode struct {
	Base
	CluserId   int    `gorm:"not null;index"`
	NodeId     string `gorm:"type:varchar(50);index""`
	MasterId   string `gorm:"type:varchar(50)"`
	Host       string `gorm:"type:varchar(50);index""`
	Port       int
	NodeRole   string `gorm:"type:varchar(50)"`
	Flags      string `gorm:"type:varchar(50)"`
	LinkState  string `gorm:"type:varchar(50)"`
	Identity   string `gorm:"type:varchar(50)"`
	InCluster  bool
	RunStatus  bool
	SlotRange  string `gorm:"type:varchar(50)"`
	SlotNumber int
}

//操作历史
type OpHistory struct {
	Base
	GroupId  int    `gorm:"not null;index"`
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

func (ClusterInfo) TableName() string {
	return "cluster_info"
}

func (RedisNode) TableName() string {
	return "redis_node"
}

func (OpHistory) TableName() string {
	return "op_history"
}
