package mysql

import (
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQL refrence a mysql db
type MySQL struct {
	*gorm.DB
}

// DB as the mysql client
var DB MySQL

// Connect create db connection
func Connect(dsn string) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DSN:                       dsn,
		DefaultStringSize:         512,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{})
	if err != nil {
		logger.Error("Mysql Cannot open mysql database: ", err.Error())
		panic(err)
	}
	DB = MySQL{db}

}

// Migrate the db schema
func Migrate() {
	logger.Info("Mysql start check data table  exists...")
	if !DB.Migrator().HasTable(&UserInfo{}) {
		logger.Info("Mysql start create data table user migrate data schemas...")
		DB.AutoMigrate(&UserInfo{})
		logger.Info("Mysql Add User to  account:iguidao , password:123456")
		DB.CreatUser("iguidao", "iguidao@iguidao.com", "tXfP0JhWJgtaNQc/DcHF78yeI73RRR+35uFNDx4cIVA=")
	}
	if !DB.Migrator().HasTable(&UserGroup{}) {
		logger.Info("Mysql start create data table user_group migrate data schemas...")
		DB.AutoMigrate(&UserGroup{})
	}
	if !DB.Migrator().HasTable(&GroupContain{}) {
		logger.Info("Mysql start create data table group_contain migrate data schemas...")
		DB.AutoMigrate(&GroupContain{})
	}
	if !DB.Migrator().HasTable(&CloudInfo{}) {
		logger.Info("Mysql start create data table cloud_info migrate data schemas...")
		DB.AutoMigrate(&CloudInfo{})
	}
	if !DB.Migrator().HasTable(&ClusterInfo{}) {
		logger.Info("Mysql start create data table cluster_info migrate data schemas...")
		DB.AutoMigrate(&ClusterInfo{})
	}
	if !DB.Migrator().HasTable(&ClusterNode{}) {
		logger.Info("Mysql start create data table redis_node migrate data schemas...")
		DB.AutoMigrate(&ClusterNode{})
	}
	if !DB.Migrator().HasTable(&OpHistory{}) {
		logger.Info("Mysql start create data table ophistory migrate data schemas...")
		DB.AutoMigrate(&OpHistory{})
	}
	if !DB.Migrator().HasTable(&CodisInfo{}) {
		logger.Info("Mysql start create data table CodisInfo migrate data schemas...")
		DB.AutoMigrate(&CodisInfo{})
	}
	if !DB.Migrator().HasTable(&Rconfig{}) {
		logger.Info("Mysql start create data table Rconfig migrate data schemas...")
		DB.AutoMigrate(&Rconfig{})
	}
	logger.Info("Mysql auto check data table done.")
}
