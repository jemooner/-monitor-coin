# monitor-coin

监听全网新币发行


## 环境搭建相关
0. 设置临时环境变量
```
   set GO111MODULE=on
   set GOPROXY=https://goproxy.io
```
   
1. `go mod init github.com/jemooner/monitor-coin`
2. `go mod tidy`
3. `go mod vendor`

## 关于control.sh
使用`control.sh`管理
- 编译 `./control build`
- 服务启动 `./control start`
- 服务重启 `./control restart`
- 服务停止 `./control stop`

## 服务进程管理
使用 `supervisor` 管理。

### 配置部署环境
- 创建配置文件  
  `/etc/supervisord.d/monitor_coin.ini`，内容如下：

```
;monitor_coin.ini

[program:monitor_coin] ;程序名
directory = /opt/www/monitor_coin.deploy/bin/ ; 程序的启动目录
; 启动命令，可以看出与手动在命令行启动的命令是一样的
command = nohup /opt/www/monitor_coin.deploy/bin/monitor_coin -config=/opt/www/monitor_coin.deploy/config/ -env=prod &  
autostart = true     ; 在 supervisord 启动的时候也自动启动
startsecs = 10        ; 启动 10 秒后没有异常退出，就当作已经正常启动了
autorestart = true   ; 程序异常退出后自动重启
startretries = 10     ; 启动失败自动重试次数，默认是 3
user = root          ; 用哪个用户启动
redirect_stderr = true  ; 把 stderr 重定向到 stdout，默认 false
stdout_logfile_maxbytes = 100MB  ; stdout 日志文件大小，默认 50MB
stdout_logfile_backups = 10     ; stdout 日志文件备份数
; stdout 日志文件，需要注意当指定目录不存在时无法正常启动，所以需要手动创建目录（supervisord 会自动创建日志文件）
stdout_logfile = /opt/logs/monitor_coin/supervisor.log
```

## 定时任务

1、币安定时任务
```
curl -X POST 'http://127.0.0.1:9081/api/monitorBinanceListing' -d '{"action":"init"}'
* * * * * curl -X POST 'http://127.0.0.1:9081/api/monitorBinanceListing' -d '{"action":"monitor"}'

```

2、mexc定时任务
```
curl -X POST 'http://127.0.0.1:9081/api/monitorMexcListing' -d '{"action":"init"}'
* * * * * curl -X POST 'http://127.0.0.1:9081/api/monitorMexcListing' -d '{"action":"monitor"}'

```

3、bitget定时任务
```
curl -X POST 'http://127.0.0.1:9081/api/monitorBitgetListing' -d '{"action":"init"}'
* * * * * curl -X POST 'http://127.0.0.1:9081/api/monitorBitgetListing' -d '{"action":"monitor"}'

```

4、kucoin定时任务
```
curl -X POST 'http://127.0.0.1:9081/api/monitorKucoinListing' -d '{"action":"init"}'
* * * * * curl -X POST 'http://127.0.0.1:9081/api/monitorKucoinListing' -d '{"action":"monitor"}'

```

4、gate.io定时任务
```
curl -X POST 'http://127.0.0.1:9081/api/monitorGateioListing' -d '{"action":"init"}'
* * * * * curl -X POST 'http://127.0.0.1:9081/api/monitorGateioListing' -d '{"action":"monitor"}'

```
