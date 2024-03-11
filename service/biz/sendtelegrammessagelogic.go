package biz

import (
	"context"
	"encoding/json"

	"monitor-coin/commonlib"
	"monitor-coin/service/telegramclient"
)

type SendTeleGramMsgReq struct {
	Message string `json:"message"`
}

func SendTeleGramMessage(ctx context.Context, req *SendTeleGramMsgReq) commonlib.ErrCode {
	telegramclient.SendMsg(ctx, req.Message, "", "test")
	return commonlib.Success
}

func BindSendTeleGramMsgReq(p string) *SendTeleGramMsgReq {
	req := new(SendTeleGramMsgReq)
	if len(p) > 0 {
		_ = json.Unmarshal([]byte(p), req)
	}
	return req
}
