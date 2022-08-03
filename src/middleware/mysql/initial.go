package mysql

import (
	"log"

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
		log.Println("Cannot open mysql database: ", err.Error())
		panic(err)
	}
	DB = MySQL{db}
}

// Migrate the db schema
func Migrate() {
	log.Println("start to auto migrate data schemas...")
	DB.Debug().AutoMigrate(&User{})
	DB.Debug().AutoMigrate(&ClusterInfo{})
	DB.Debug().AutoMigrate(&RedisNodes{})
	DB.Debug().AutoMigrate(&UserGroup{})
	DB.Debug().AutoMigrate(&OpHistory{})
	log.Println("auto migrate data schemas done.")
}
