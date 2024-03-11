package biz

import (
	"context"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/binanceclient"
	"monitor-coin/service/dao"
	"monitor-coin/service/telegramclient"
	"strings"
	"time"
)

func MonitorBinanceListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}
	allCoins, err := binanceclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorBinanceListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	mkt := commonlib.MarketTagBinance
	var coins []*dao.MbCoinEntity
	var s = strings.ToLower("TRADING")
	where := map[string]interface{}{}
	for _, c := range allCoins {
		if strings.ToLower(c.Status) != s {
			dlog.Infof("%v||MonitorBinanceListing ignore coin=%+v", trc, c)
			continue
		}
		where["coin_name"] = c.BaseAsset
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		if len(coins) != 0 {
			if !strings.Contains(coins[0].MarketTag, mkt) {
				updateCoinMarketTag(ctx, mkt, coins[0])
				go telegramclient.SendMsg(ctx, "币安[binance]", c.BaseAsset, req.Action)
			}
			continue // 说明不是新币
		}

		// 说明是新币
		coin := buildMbCoinEntity(c.BaseAsset, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorBinanceListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}
		//发送消息TeleGram群组
		go telegramclient.SendMsg(ctx, "币安[binance]", c.BaseAsset, req.Action)
	}

	return commonlib.Success
}

func buildMbCoinEntity(currency string, market string) *dao.MbCoinEntity {
	en := new(dao.MbCoinEntity)
	en.CoinName = currency
	en.CoinPrice = ""
	en.MarketTag = market
	en.ExtInfo = ""
	en.Remark = ""
	en.DelFlag = 0
	en.ListTime = time.Now().Format("2006-01-02 15:04:05")
	en.CreateTime = en.ListTime
	en.UpdateTime = en.ListTime
	return en
}
