package biz

import (
	"context"
	"monitor-coin/service/Bitstampclient"
	"strings"

	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/dao"
	"monitor-coin/service/telegramclient"
)

func MonitorBitstampListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}
	allCoins, err := Bitstampclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorBitstampListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	mkt := commonlib.MarketTagBitstamp
	var coins []*dao.MbCoinEntity
	//将字符转为小写
	var s = strings.ToLower("Enabled")
	where := map[string]interface{}{}
	for _, c := range allCoins {
		if len(c.Name) > 0 {
			//字符串分割，并获取下标0的字符
			ss := strings.Split(c.Name, "/")
			c.Name = ss[0]
		}
		//判断是否为法币
		if commonlib.CheckFiatCurrency(c.Name) {
			continue
		}
		//转态不为enabled，则不可交易
		if strings.ToLower(c.Trading) != s {
			dlog.Infof("%v||MonitorBitstampListing coin=%+v Unable to trade status=%v", trc, c.Name, c.Trading)
			continue
		}
		where["coin_name"] = c.Name
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		if len(coins) != 0 {
			//tag中没有该交易所，则添加
			if !strings.Contains(coins[0].MarketTag, mkt) {
				updateCoinMarketTag(ctx, mkt, coins[0])
				go telegramclient.SendMsg(ctx, mkt, c.Name, req.Action)
			}
			continue // 说明不是新币
		}

		// 说明是新币
		coin := buildMbCoinEntity(c.Name, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorBitstampListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}

		//发送消息TeleGram群组
		go telegramclient.SendMsg(ctx, mkt, c.Name, req.Action)
	}

	return commonlib.Success
}
