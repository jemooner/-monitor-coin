package binanceclient

import (
	"context"
	"encoding/json"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"net/http"
)

type GetAllCoinsResp struct {
	Timezone   string        `json:"timezone"`
	ServerTime int64         `json:"serverTime"`
	Symbols    []*CoinItemVo `json:"symbols"`
}

type CoinItemVo struct {
	Status    string `json:"status"`
	BaseAsset string `json:"baseAsset"`
}

func GetAllCoins(ctx context.Context) (map[string]*CoinItemVo, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||binanceclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("binanceclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r GetAllCoinsResp
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil || len(r.Symbols) == 0 {
		return nil, fmt.Errorf("binanceclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	var ok bool
	coins := make(map[string]*CoinItemVo)
	for _, i := range r.Symbols {
		if _, ok = coins[i.BaseAsset]; !ok {
			coins[i.BaseAsset] = i
		}
	}
	return coins, nil
}
