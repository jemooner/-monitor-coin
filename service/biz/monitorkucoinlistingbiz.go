package biz

import (
	"context"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/dao"
	"monitor-coin/service/kucoinclient"
	"monitor-coin/service/telegramclient"
	"strings"
)

func MonitorKucoinListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}
	allCoins, err := kucoinclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorKucoinListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	mkt := commonlib.MarketTagKucoin
	var coins []*dao.MbCoinEntity
	where := map[string]interface{}{}
	for _, c := range allCoins {
		where["coin_name"] = c.BaseCurrency
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		if len(coins) != 0 {
			if !strings.Contains(coins[0].MarketTag, mkt) {
				updateCoinMarketTag(ctx, mkt, coins[0])
				go telegramclient.SendMsg(ctx, mkt, c.BaseCurrency, req.Action)
			}
			continue // 说明不是新币
		}

		// 说明是新币
		coin := buildMbCoinEntity(c.BaseCurrency, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorKucoinListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}

		go telegramclient.SendMsg(ctx, mkt, c.BaseCurrency, req.Action)
	}

	return commonlib.Success
}
