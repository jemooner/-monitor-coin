package handler

import (
	"context"
	"monitor-coin/commonlib"
	"monitor-coin/service/biz"
)

/*
*
监听binance新币
*/
func MonitorBinanceListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorBinanceListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}

/*
*
监听mexc新币
*/
func MonitorMexcListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)

	ec := biz.MonitorMexcListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}

/*
*
监听bitget新币
*/
func MonitorBitgetListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorBitgetListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)

}

/*
*
监听kucoin新币
*/
func MonitorKucoinListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorKucoinListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}

/*
*
监听Gateio新币
*/
func MonitorGateioListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorGateioListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}

/*
*
监听Coinbase新币

curl -X POST 'http://127.0.0.1:9081/api/monitorCoinbaseListing' -d '{"action":"monitor"}'
*/
func MonitorCoinbaseListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorCoinbaseListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}

/*
*
监听Bitfinex新币

curl -X POST 'http://127.0.0.1:9081/api/monitorBitfinexListing' -d '{"action":"monitor"}'
*/
func MonitorBitfinexListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorBitfinexListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}

/*
*
监听Bitstamp新币

curl -X POST 'http://127.0.0.1:9081/api/monitorBitstampListing' -d '{"action":"monitor"}'
*/
func MonitorBitstampListingHandler(ctx context.Context, params string) []byte {
	req := biz.BindMonitorNewListingReq(params)
	ec := biz.MonitorBitstampListing(ctx, req)
	return commonlib.FormatResp(ctx, ec, nil)
}
