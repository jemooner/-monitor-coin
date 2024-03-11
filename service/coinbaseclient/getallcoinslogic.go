package coinbaseclient

import (
	"context"
	"encoding/json"
	"fmt"
	"monitor-coin/commonlib"
	"monitor-coin/commonlib/dlog"
	"net/http"
)

type GetAllCoinsResp struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

func GetAllCoins(ctx context.Context) ([]*GetAllCoinsResp, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||coinbaseclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("coinbaseclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r []*GetAllCoinsResp
	err = json.Unmarshal(resp, &r)
	if err != nil || len(r) == 0 {
		return nil, fmt.Errorf("coinbaseclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	return r, nil
}

/*
[{
	"id": "EUR",
	"name": "Euro",
	"min_size": "0.01",
	"status": "online",
	"message": "",
	"max_precision": "0.01",
	"convertible_to": [],
	"details": {
		"type": "fiat",
		"symbol": "€",
		"network_confirmations": null,
		"sort_order": 2,
		"crypto_address_link": null,
		"crypto_transaction_link": null,
		"push_payment_methods": ["sepa_bank_account"],
		"group_types": ["fiat", "eur"],
		"display_name": null,
		"processing_time_seconds": null,
		"min_withdrawal_amount": null,
		"max_withdrawal_amount": null
	},
	"default_network": "",
	"supported_networks": []
}, {
	"id": "00",
	"name": "00 Token",
	"min_size": "0.00000001",
	"status": "online",
	"message": "",
	"max_precision": "0.00000001",
	"convertible_to": [],
	"details": {
		"type": "crypto",
		"symbol": null,
		"network_confirmations": 14,
		"sort_order": 0,
		"crypto_address_link": "https://etherscan.io/token/0x881ba05de1e78f549cc63a8f6cabb1d4ad32250d?a={{address}}",
		"crypto_transaction_link": "https://etherscan.io/tx/0x{{txId}}",
		"push_payment_methods": [],
		"group_types": [],
		"display_name": null,
		"processing_time_seconds": null,
		"min_withdrawal_amount": 1e-8,
		"max_withdrawal_amount": 410000
	},
	"default_network": "ethereum",
	"supported_networks": [{
		"id": "ethereum",
		"name": "Ethereum",
		"status": "online",
		"contract_address": "0x881ba05de1e78f549cc63a8f6cabb1d4ad32250d",
		"crypto_address_link": "https://etherscan.io/token/0x881ba05de1e78f549cc63a8f6cabb1d4ad32250d?a={{address}}",
		"crypto_transaction_link": "https://etherscan.io/tx/0x{{txId}}",
		"min_withdrawal_amount": 1e-8,
		"max_withdrawal_amount": 410000,
		"network_confirmations": 14,
		"processing_time_seconds": null
	}]
}]
*/
