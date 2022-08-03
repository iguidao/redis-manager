package mysql

import (
	"time"

	"gorm.io/gorm"
)

// cluster
type Base struct {
	ID        uint      `gorm:"primary_key"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	//创建索引`sql:"index"`
}

type ClusterInfo struct {
	Base
	ClusterName    string `gorm:"not null;index"`
	ClusterMode    string // 集群(Cluster)；单点(Single)；哨兵(Sentinel)
	ClusterVersion string
	NodesAll       int
	NodesMaster    int
	NodesSlave     int
	RedisPassword  string
	Environment    string // 主机 Machine；容器 Container
	From           string //导入Import；平台创建Self
}

// redis cluster node
type RedisNode struct {
	Base
	CluserId   uint   `gorm:"not null;index"`
	NodeId     string `gorm:"type:varchar(50)"`
	MasterId   string `gorm:"type:varchar(50)"`
	Identity   string `gorm:"type:varchar(50)"`
	Ip         string `gorm:"type:varchar(50)"`
	Port       int
	SlotNumber int
	Name       string `gorm:"not null;index"`
}

//user
type User struct {
	UserName string `gorm:"not null;index"`
	Password string `gorm:"not null"`
	Email    string `gorm:"type:varchar(255)"`
	Mobile   string `gorm:"type:varchar(255)"`
	GroupId  uint   `gorm:"not null;index"`
	UserType string `gorm:"not null;index"` //admin 管理员；visitor 访客；staff 员工
}

type UserGroup struct {
	Name string `gorm:"not null;type:varchar(255);index"`
	Info string `gorm:"type:varchar(255)"`
}

//history
type OpHistory struct {
	Base
	UserId   uint
	OpInfo   string `gorm:"type:varchar(50)"` // 操作动作
	OpParams string `gorm:"type:text"`        //操作参数属组或者对象
}
