package commonlib

import (
	"fmt"
	"os"

	"monitor-coin/commonlib/dlog"
)

func InitLogger(dlogConf LogConf) {
	dlogCfg := dlog.LogConfig{}
	dlogCfg.Type = "std"
	if FlagVar.Env != EnvLocal {
		dlogCfg.Type = "file"
		dlogCfg.FilePath = dlogConf.LogPath
		dlogCfg.RotateByHour = true
		dlogCfg.KeepHours = dlogConf.KeepHours
		dlogCfg.Level = "INFO"
	}

	err := dlog.Init(dlogCfg)
	if err != nil {
		fmt.Printf("init logger fail: %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("init logger done", dlogConf)
	dlog.Infof("init logger done,conf=%+v", dlogConf)
}
