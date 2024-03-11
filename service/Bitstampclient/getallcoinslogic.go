package Bitstampclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
)

type GetAllCoinsResp struct {
	Name    string `json:"name"`
	Trading string `json:"trading"`
}

func GetAllCoins(ctx context.Context) ([]*GetAllCoinsResp, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||Bitstampclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("Bitstampclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r []*GetAllCoinsResp
	err = json.Unmarshal(resp, &r)
	if err != nil || len(r) == 0 {
		return nil, fmt.Errorf("Bitstampclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	return r, nil
}

/*
[{
	"name": "BTC/USD",
	"url_symbol": "btcusd",
	"base_decimals": 8,
	"counter_decimals": 0,
	"instant_order_counter_decimals": 2,
	"minimum_order": "10 USD",
	"trading": "Enabled",
	"instant_and_market_orders": "Enabled",
	"description": "Bitcoin / U.S. dollar"
}, {
	"name": "DOGE/EUR",
	"url_symbol": "dogeeur",
	"base_decimals": 2,
	"counter_decimals": 5,
	"instant_order_counter_decimals": 5,
	"minimum_order": "10.00000 EUR",
	"trading": "Enabled",
	"instant_and_market_orders": "Disabled",
	"description": "Dogecoin / Euro"
}]
*/
