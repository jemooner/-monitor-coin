package gateioclient

import (
	"context"
	"encoding/json"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"net/http"
)

type CoinItemVo struct {
	Currency string `json:"currency"`
}

func GetAllCoins(ctx context.Context) ([]*CoinItemVo, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||gateioclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("gateioclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r []*CoinItemVo
	err = json.Unmarshal(resp, &r)
	if err != nil || len(r) == 0 {
		return nil, fmt.Errorf("gateioclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	return r, nil
}
