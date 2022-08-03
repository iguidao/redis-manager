package cfg

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Get_Local(get_local string) string {
	switch get_local {
	case "addr":
		local_addr := viper.GetString("local.addr")
		return local_addr
	case "logapipath":
		local_logapipath := viper.GetString("local.logapipath")
		return local_logapipath
	case "logapppath":
		local_logapppath := viper.GetString("local.logapppath")
		return local_logapppath
	default:
		return "noconfig"
	}
}

func Get_Info(get_type string) string {
	switch get_type {
	case "MYSQL":
		mysql_name := viper.GetString("mysql.name")
		mysql_addr := viper.GetString("mysql.addr")
		mysql_username := viper.GetString("mysql.username")
		mysql_password := viper.GetString("mysql.password")
		mysql_url := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", mysql_username, mysql_password, mysql_addr, mysql_name)
		return mysql_url
	case "REDIS":
		redis_addr := viper.GetString("redis.addr")
		redis_port := viper.GetString("redis.port")
		redis_url := fmt.Sprintf("%s:%s", redis_addr, redis_port)
		return redis_url
	default:
		return "noconfig"
	}
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}
	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("yaml")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}
