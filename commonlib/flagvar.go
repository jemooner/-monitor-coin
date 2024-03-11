package commonlib

import (
	"flag"
	"fmt"
)

type FlagVarTy struct {
	Env        string
	ConfigPath string
}

var FlagVar FlagVarTy

const (
	EnvProd  = "prod"
	EnvDev   = "dev"
	EnvLocal = "local"
	EnvTest  = "test"
)

func InitEnvVar() {
	flag.StringVar(&FlagVar.Env, "env", "prod", "env local/dev/prod")
	flag.StringVar(&FlagVar.ConfigPath, "config", "", "server config.")

	flag.Parse()
	fmt.Printf("%+v\n", FlagVar)
}

func InitLocalEnvVar() {
	FlagVar.Env = "local"
	FlagVar.ConfigPath = "www/wallet-proxy/config/"
}
