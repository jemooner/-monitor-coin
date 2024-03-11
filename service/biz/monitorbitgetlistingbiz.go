package biz

import (
	"context"
	"encoding/json"
	"strings"

	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/bitgetclient"
	"monitor-coin/service/dao"
	"monitor-coin/service/telegramclient"
)

type MonitorNewListingReq struct {
	Action string `json:"action"` // init、monitor
}

func BindMonitorNewListingReq(p string) *MonitorNewListingReq {
	req := new(MonitorNewListingReq)
	if len(p) > 0 {
		_ = json.Unmarshal([]byte(p), req)
	}
	return req
}

func MonitorBitgetListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}
	allCoins, err := bitgetclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorBitgetListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	mkt := commonlib.MarketTagBitget
	var coins []*dao.MbCoinEntity
	where := map[string]interface{}{}
	for _, c := range allCoins {
		where["coin_name"] = c.CoinName
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		if len(coins) != 0 {
			if !strings.Contains(coins[0].MarketTag, mkt) {
				updateCoinMarketTag(ctx, mkt, coins[0])
				go telegramclient.SendMsg(ctx, mkt, c.CoinName, req.Action)
			}
			continue // 说明不是新币
		}

		// 说明是新币
		coin := buildMbCoinEntity(c.CoinName, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorBitgetListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}

		//发送消息TeleGram群组
		go telegramclient.SendMsg(ctx, mkt, c.CoinName, req.Action)
	}

	return commonlib.Success
}
