package handler

import (
	"context"
	"monitor-coin/commonlib"
	"net/url"
)

/*
*
探活
*/
func PingHandler(ctx context.Context, params url.Values) []byte {
	return commonlib.FormatResp(ctx, commonlib.Success, "pong")
}

/*
*
查看服务版本
*/
func VersionHandler(ctx context.Context, params string) []byte {
	return commonlib.FormatResp(ctx, commonlib.Success, commonlib.LaunchConfig().Server.Version)
}

func NotFoundHandler(ctx context.Context, params string) []byte {
	return commonlib.FormatResp(ctx, commonlib.ErrApiNotFound, nil)
}
