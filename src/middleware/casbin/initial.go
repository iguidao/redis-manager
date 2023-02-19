package casbin

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Enforcer *casbin.Enforcer
var err error

// DB as the mysql client
var DB *gorm.DB

func Connect(dsn string) {
	// Increase the column size to 512.

	// db, _ := gorm.Open(...)
	DB, err := gorm.Open(mysql.New(mysql.Config{
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
	// Initialize a Gorm adapter and use it in a Casbin enforcer:
	// The adapter will use an existing gorm.DB instnace.
	// a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(DB, &CasbinRule{})
	// a := gormadapter.NewAdapterByDB(db)
	m, err := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act || r.sub == "iguidao"
	`)
	if err != nil {
		logger.Error("error: model: ", err)
	}

	Enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		logger.Error("error: enforcer: ", err)
	}
	// 开启权限认证日志
	Enforcer.EnableLog(true)
	// Load the policy from DB.
	err = Enforcer.LoadPolicy()
	if err != nil {
		logger.Error("loadPolicy error")
		panic(err)
	}

}
