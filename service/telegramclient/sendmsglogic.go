package telegramclient

import (
	"context"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"net/http"
)

func SendMsg(ctx context.Context, mkt string, coin string, action string) {
	if action == "init" {
		return
	}
	// 消息体
	trc := commonlib.GetTrace(ctx)
	msg := fmt.Sprintf("%s 交易所上线新币种 %s", mkt, coin)
	endpoint := getEndPoint("SendMessage")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil, "-852864287,", msg)
	dlog.Infof("%v||SendMsg->SendMessage done,resp=%s,err=%v", trc, resp, err)
}
