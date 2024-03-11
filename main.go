package main

import (
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/handler"
	"monitor-coin/service/Bitstampclient"
	"monitor-coin/service/binanceclient"
	"monitor-coin/service/bitfinexclient"
	"monitor-coin/service/bitgetclient"
	"monitor-coin/service/coinbaseclient"
	"monitor-coin/service/gateioclient"
	"monitor-coin/service/kucoinclient"
	"monitor-coin/service/mexcclient"
	"monitor-coin/service/telegramclient"
	"net/http"
	"os"
)

func main() {
	// 加载配置
	commonlib.InitEnvVar()
	conf := commonlib.LaunchConfig()

	// 初始化日志组件
	commonlib.InitLogger(conf.Logger)
	defer dlog.Close()

	commonlib.InitMysql(&conf.Mysql)
	defer commonlib.ReleaseMysql()

	binanceclient.InitBinanceClient(&conf.Binance)
	mexcclient.InitMexcClient(&conf.Mexc)
	bitgetclient.InitBitgetClient(&conf.Bitget)
	kucoinclient.InitKucoinClient(&conf.Kucoin)
	gateioclient.InitGateioClient(&conf.Gateio)
	coinbaseclient.InitCoinbaseClient(&conf.Coinbase)
	bitfinexclient.InitBitfinexClient(&conf.Coinbase)
	Bitstampclient.InitBitstampClient(&conf.Coinbase)
	telegramclient.InitTelegramClient(&conf.TeleGram)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	//http.HandleFunc("/news/index.html", handlers.NewsIndexHandler)
	//http.HandleFunc("/news/list.html", handlers.NewsListHandler)
	//http.HandleFunc("/news/detail.html", handlers.NewsDetailHandler)

	http.HandleFunc("/version", commonlib.Wrapper(handler.VersionHandler))

	http.HandleFunc("/api/monitorBinanceListing", commonlib.Wrapper(handler.MonitorBinanceListingHandler))
	http.HandleFunc("/api/monitorMexcListing", commonlib.Wrapper(handler.MonitorMexcListingHandler))
	http.HandleFunc("/api/monitorBitgetListing", commonlib.Wrapper(handler.MonitorBitgetListingHandler))
	http.HandleFunc("/api/monitorKucoinListing", commonlib.Wrapper(handler.MonitorKucoinListingHandler))
	http.HandleFunc("/api/monitorGateioListing", commonlib.Wrapper(handler.MonitorGateioListingHandler))
	http.HandleFunc("/api/monitorCoinbaseListing", commonlib.Wrapper(handler.MonitorCoinbaseListingHandler))
	http.HandleFunc("/api/monitorBitfinexListing", commonlib.Wrapper(handler.MonitorBitfinexListingHandler))
	http.HandleFunc("/api/monitorBitstampListing", commonlib.Wrapper(handler.MonitorBitstampListingHandler))

	http.HandleFunc("/api/sendTeleGramMessage", commonlib.Wrapper(handler.SendTeleGramMessageHandler))

	dlog.Infof("service %s started", conf.Server.ServiceName)

	err := http.ListenAndServe(conf.Server.Port, nil) // 设置监听的端口
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
