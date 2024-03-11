package commonlib

import (
	"fmt"
	"testing"
)

func Test_LaunchConfig1(t *testing.T) {
	// 读取配置文件内容
	FlagVar.Env = "local"
	path := "www/monitor-coin/config/"
	launchConfig(path)
	fmt.Printf("%+v", serviceConf)
}
