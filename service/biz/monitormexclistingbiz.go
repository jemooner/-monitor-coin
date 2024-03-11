package biz

import (
	"context"
	"strings"

	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/dao"
	"monitor-coin/service/mexcclient"
	"monitor-coin/service/telegramclient"
)

func MonitorMexcListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}

	allCoins, err := mexcclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorMexcListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	l := 0
	price := new(mexcclient.PriceTickerVo)
	mkt := commonlib.MarketTagMexc
	var coins []*dao.MbCoinEntity
	var s = strings.ToLower("ENABLED")
	where := map[string]interface{}{}
	for _, c := range allCoins {
		if strings.ToLower(c.Status) != s {
			dlog.Infof("%v||MonitorMexcListing ignore coin=%+v", trc, c)
			continue
		}

		where["coin_name"] = c.BaseAsset
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		l = len(coins)
		if l != 0 && strings.Contains(coins[0].MarketTag, mkt) {
			continue // 说明已经上市
		}

		price, err = mexcclient.Get24hrPriceTicker(ctx, c.BaseAsset+c.QuoteAsset)
		dlog.Infof("%v||MonitorMexcListing Get24hrPriceTicker price=%+v,err=%+v", trc, price, c)
		if err != nil || price == nil || price.OpenTime == 0 {
			continue // 说明可能是未正式上市
		}

		if l != 0 && !strings.Contains(coins[0].MarketTag, mkt) {
			updateCoinMarketTag(ctx, mkt, coins[0]) // 说明别的交易所已上线 但mexc刚上线
			go telegramclient.SendMsg(ctx, "抹茶[MEXC]", c.BaseAsset, req.Action)
		}

		// 说明是新币且已上市
		coin := buildMbCoinEntity(c.BaseAsset, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorMexcListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}

		//发送消息TeleGram群组
		go telegramclient.SendMsg(ctx, "抹茶[MEXC]", c.BaseAsset, req.Action)
	}

	return commonlib.Success
}

func updateCoinMarketTag(ctx context.Context, mt string, coin *dao.MbCoinEntity) {
	trc := commonlib.GetTrace(ctx)
	if strings.Contains(coin.MarketTag, mt) {
		dlog.Infof("%v||updateCoinMarketTag existed coin=%+v", trc, coin)
		return
	}
	update := map[string]interface{}{}
	update["market_tag"] = coin.MarketTag + "," + mt
	row, err := dao.UpdateMbCoin(ctx, update, map[string]interface{}{"id": coin.CoinId})
	dlog.Infof("%v||MonitorMexcListing UpdateMbCoin row=%d,err=%v", trc, row, err)
}
