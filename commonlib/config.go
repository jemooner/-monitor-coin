package commonlib

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

var serviceConf *ServiceConfig
var once sync.Once

type ServiceConfig struct {
	Server   ServerConf     `mapstructure:"server"`
	Logger   LogConf        `mapstructure:"log"`
	Mysql    MysqlConf      `mapstructure:"mysql"`
	Redis    RedisConf      `mapstructure:"redis"`
	Binance  HttpClientConf `mapstructure:"binance"`
	Mexc     HttpClientConf `mapstructure:"mexc"`
	Bitget   HttpClientConf `mapstructure:"bitget"`
	Kucoin   HttpClientConf `mapstructure:"kucoin"`
	Gateio   HttpClientConf `mapstructure:"gateio"`
	Coinbase HttpClientConf `mapstructure:"coinbase"`
	Bitfinex HttpClientConf `mapstructure:"bitfinex"`
	Bitstamp HttpClientConf `mapstructure:"bitstamp"`
	TeleGram HttpClientConf `mapstructure:"teleGram"`
}

type ServerConf struct {
	Port        string `mapstructure:"port"`
	MaxProcess  int    `mapstructure:"max_process"`
	Version     string `mapstructure:"version"`
	ServiceName string `mapstructure:"name"`
	DataPath    string `mapstructure:"data_path"`
}

type LogConf struct {
	LogPath   string `mapstructure:"log_path"`
	KeepHours uint   `mapstructure:"keep_hours"`
}

type MysqlConf struct {
	Dsn             string `mapstructure:"dsn"`
	Retry           int    `mapstructure:"retry"`
	PoolMaxIdleConn int    `mapstructure:"pool_max_idle_conn"`
	PoolMaxOpenConn int    `mapstructure:"pool_max_open_conn"`
}

type RedisConf struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
}

type HttpClientConf struct {
	Host       []string `mapstructure:"host"`
	TimeoutSec int      `mapstructure:"timeout_sec"`
	SignKey    string   `mapstructure:"sign_key"`
}

func launchConfig(path string) {
	// 读取配置文件内容
	configPath := "./config"
	if len(path) > 0 {
		configPath = path
	}
	if FlagVar.ConfigPath != "" {
		configPath = FlagVar.ConfigPath
	}
	fmt.Printf("load config path=%v,name=%s\n", configPath, FlagVar.Env)

	// 解析配置
	vp := viper.New()
	vp.AddConfigPath(configPath)
	vp.SetConfigName(FlagVar.Env)
	vp.SetConfigType("toml")
	if err := vp.ReadInConfig(); err != nil {
		fmt.Printf("fail to read config,err=%+v,config=%+v\n", err, serviceConf)
		os.Exit(1)
	}

	var cf ServiceConfig
	err := vp.Unmarshal(&cf)
	if err != nil {
		fmt.Printf("fail to Unmarshal config,err=%+v\n", err)
		os.Exit(1)
	}
	serviceConf = &cf
}

func LaunchConfig() *ServiceConfig {
	if serviceConf == nil {
		once.Do(func() {
			launchConfig(FlagVar.ConfigPath)
		})
	}
	if serviceConf == nil {
		fmt.Printf("LaunchConfig fail. conf=nil")
		os.Exit(1)
	}
	return serviceConf
}
