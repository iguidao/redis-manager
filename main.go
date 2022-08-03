package main

import (
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/logger"
	"github.com/iguidao/redis-manager/src/middleware/mysql"
	"github.com/iguidao/redis-manager/src/rhttp"
)

func init() {
	if err := cfg.Init(""); err != nil {
		panic(err)
	}
	logger.SetupLogger()
	mysql.Connect(cfg.Get_Info("MYSQL"))
	mysql.Migrate()

}

func main() {
	listen := cfg.Get_Local("addr")
	if listen == "" {
		listen = ":8000"
	}
	rhttp.NewServer().Run(listen)
}
