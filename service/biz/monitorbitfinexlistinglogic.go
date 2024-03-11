package biz

import (
	"context"
	"strings"

	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/bitfinexclient"
	"monitor-coin/service/dao"
	"monitor-coin/service/telegramclient"
)

func MonitorBitfinexListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}
	allCoins, err := bitfinexclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorBitfinexListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	mkt := commonlib.MarketTagBitfinex
	var coins []*dao.MbCoinEntity
	where := map[string]interface{}{}
	for _, c := range allCoins {
		//判断是否为法币
		if commonlib.CheckFiatCurrency(c) {
			continue
		}

		where["coin_name"] = c
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		if len(coins) != 0 {
			//tag中没有该交易所，则添加
			if !strings.Contains(coins[0].MarketTag, mkt) {
				updateCoinMarketTag(ctx, mkt, coins[0])
				go telegramclient.SendMsg(ctx, mkt, c, req.Action)
			}
			continue // 说明不是新币
		}

		// 说明是新币
		coin := buildMbCoinEntity(c, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorBitfinexListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}

		//发送消息TeleGram群组
		go telegramclient.SendMsg(ctx, mkt, c, req.Action)
	}
	return commonlib.Success
}
