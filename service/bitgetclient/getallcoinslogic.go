package bitgetclient

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
	CoinName string `json:"coinName"`
}

func GetAllCoins(ctx context.Context) (map[string]*CoinItemVo, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||bitgetclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("bitgetclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r GetAllCoinsResp
	err = json.Unmarshal([]byte(resp), &r)
	if err != nil || len(r.Data) == 0 {
		return nil, fmt.Errorf("bitgetclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	var ok bool
	coins := make(map[string]*CoinItemVo)
	for _, i := range r.Data {
		if _, ok = coins[i.CoinName]; !ok {
			coins[i.CoinName] = i
		}
	}
	return coins, nil
}
