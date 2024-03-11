package kucoinclient

import (
	"context"
	"encoding/json"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"net/http"
)

type GetAllCoinsResp struct {
	Code string        `json:"code"`
	Data []*CoinItemVo `json:"data"`
}

type CoinItemVo struct {
	BaseCurrency  string `json:"baseCurrency"`
	Market        string `json:"market"`
	EnableTrading bool   `json:"enableTrading"`
}

func GetAllCoins(ctx context.Context) (map[string]*CoinItemVo, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||kucoinclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("kucoinclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r GetAllCoinsResp
	err = json.Unmarshal(resp, &r)
	if err != nil || len(r.Data) == 0 {
		return nil, fmt.Errorf("kucoinclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	var ok bool
	coins := make(map[string]*CoinItemVo)
	for _, i := range r.Data {
		if _, ok = coins[i.BaseCurrency]; !ok {
			coins[i.BaseCurrency] = i
		}
	}
	return coins, nil
}
