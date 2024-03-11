package handler

import (
	"context"
	"monitor-coin/commonlib"
	"monitor-coin/service/biz"
)

func SendTeleGramMessageHandler(ctx context.Context, params string) []byte {
	req := biz.BindSendTeleGramMsgReq(params)
	ec := biz.SendTeleGramMessage(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}
