package main

import (
	"flag"

	"github.com/iguidao/redis-web-manager/server/src/cfg"
	"github.com/iguidao/redis-web-manager/server/src/http"
	"github.com/iguidao/redis-web-manager/server/src/middleware/mysql"
)

var migrate = flag.Bool("m", false, "migrate the database schemas.")

func init() {
	if err := cfg.Init(""); err != nil {
		panic(err)
	}
	mysql.Connect(cfg.Get_Info("MYSQL"))
}

func main() {
	flag.Parse()
	if *migrate {
		mysql.Migrate()
		return
	}
	listen := cfg.Get_Local("addr")
	if listen == "" {
		listen = ":8000"
	}
	http.NewServer().Run(listen)
}
