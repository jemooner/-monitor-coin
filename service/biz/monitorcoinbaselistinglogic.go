package biz

import (
	"context"
	"strings"

	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"monitor-coin/service/coinbaseclient"
	"monitor-coin/service/dao"
	"monitor-coin/service/telegramclient"
)

func MonitorCoinbaseListing(ctx context.Context, req *MonitorNewListingReq) commonlib.ErrCode {
	trc := commonlib.GetTrace(ctx)
	if req == nil || len(req.Action) == 0 {
		return commonlib.ErrParam
	}
	allCoins, err := coinbaseclient.GetAllCoins(ctx)
	if err != nil || allCoins == nil || len(allCoins) == 0 {
		dlog.Errorf("%v||MonitorCoinbaseListing GetAllCoins no data or err=%v", trc, err)
		return commonlib.Success
	}

	mkt := commonlib.MarketTagCoinbase
	var coins []*dao.MbCoinEntity
	where := map[string]interface{}{}
	for _, c := range allCoins {
		//判断是否为法币
		if commonlib.CheckFiatCurrency(c.Id) {
			continue
		}
		//转态不为online，则不可交易
		if c.Status != "online" {
			dlog.Infof("%v||MonitorCoinbaseListing coin=%+v Unable to trade status=%v", trc, c.Id, c.Status)
			continue
		}
		where["coin_name"] = c.Id
		coins, err = dao.QueryMbCoinList(ctx, where)
		if err != nil {
			continue
		}
		if len(coins) != 0 {
			//tag中没有该交易所，则添加
			if !strings.Contains(coins[0].MarketTag, mkt) {
				updateCoinMarketTag(ctx, mkt, coins[0])
				go telegramclient.SendMsg(ctx, mkt, c.Id, req.Action)
			}
			continue // 说明不是新币
		}

		// 说明是新币
		coin := buildMbCoinEntity(c.Id, mkt)
		id, err := dao.InsertMbCoin(ctx, []*dao.MbCoinEntity{coin})
		dlog.Infof("%v||MonitorCoinbaseListing InsertMbCoin coin=%+v,id=%v,err=%v", trc, coin, id, err)
		if err != nil {
			continue
		}

		//发送消息TeleGram群组
		go telegramclient.SendMsg(ctx, mkt, c.Id, req.Action)
	}

	return commonlib.Success
}
