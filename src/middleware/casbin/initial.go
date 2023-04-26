package casbin

import (
	"github.com/casbin/casbin/v2"
	cmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	rmysql "github.com/iguidao/redis-manager/src/middleware/mysql"
)

var Enforcer *casbin.Enforcer

func Connect() {
	var err error
	a, _ := gormadapter.NewAdapterByDBWithCustomTable(rmysql.DB.DB, &CasbinRule{})
	// a := gormadapter.NewAdapterByDB(db)
	m, err := cmodel.NewModelFromString(`
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act) || r.sub == "admin"
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

// // DB as the mysql client
// var DB *gorm.DB

// // Connect create db connection

// func Connect(dsn string) {
// 	// Increase the column size to 512.

// 	// db, _ := gorm.Open(...)
// 	DB, err := gorm.Open(mysql.New(mysql.Config{
// 		// DSN: "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
// 		DSN:                       dsn,
// 		DefaultStringSize:         512,   // string 类型字段的默认长度
// 		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
// 		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
// 		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
// 		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
// 	}), &gorm.Config{})
// 	if err != nil {
// 		logger.Error("Cannot open mysql database: ", err.Error())
// 		panic(err)
// 	}
// 	// Initialize a Gorm adapter and use it in a Casbin enforcer:
// 	// The adapter will use an existing gorm.DB instnace.
// 	// a, _ := gormadapter.NewAdapterByDBWithCustomTable(db, &CasbinRule{})
// 	a, _ := gormadapter.NewAdapterByDBWithCustomTable(DB, &CasbinRule{})
// 	// a := gormadapter.NewAdapterByDB(db)
// 	m, err := model.NewModelFromString(`
// 	[request_definition]
// 	r = sub, obj, act

// 	[policy_definition]
// 	p = sub, obj, act

// 	[policy_effect]
// 	e = some(where (p.eft == allow))

// 	[matchers]
// 	m = r.sub == p.sub && r.obj == p.obj && r.act == p.act || r.sub == "iguidao"
// 	`)
// 	if err != nil {
// 		logger.Error("error: model: ", err)
// 	}

// 	Enforcer, err = casbin.NewEnforcer(m, a)
// 	if err != nil {
// 		logger.Error("error: enforcer: ", err)
// 	}
// 	// 开启权限认证日志
// 	Enforcer.EnableLog(true)
// 	// Load the policy from DB.
// 	err = Enforcer.LoadPolicy()
// 	if err != nil {
// 		logger.Error("loadPolicy error")
// 		panic(err)
// 	}

// }
