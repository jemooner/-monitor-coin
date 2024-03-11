package mexcclient

import (
	"context"
	"encoding/json"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"net/http"
)

type PriceTickerVo struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
}

func Get24hrPriceTicker(ctx context.Context, symbol string) (*PriceTickerVo, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("Get24hrPriceTicker")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, map[string]interface{}{"symbol": symbol})
	dlog.Infof("%v||mexcclient->Get24hrPriceTicker done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("mexcclient->Get24hrPriceTicker fail, err=%v or resp=nil", err)
	}

	var r PriceTickerVo
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return nil, fmt.Errorf("mexcclient->Get24hrPriceTicker Unmarshal fail or no data, err=%v", err)
	}

	return &r, nil
}
