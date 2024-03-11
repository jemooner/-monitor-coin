package bitfinexclient

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

func GetAllCoins(ctx context.Context) ([]string, error) {
	trc := commonlib.GetTrace(ctx)

	endpoint := getEndPoint("GetAllCoins")
	resp, err := doHttpRequest(ctx, http.MethodGet, endpoint, nil)
	dlog.Infof("%v||bitfinexclient->GetAllCoins done,err=%v", trc, err)
	if err != nil || len(resp) == 0 {
		return nil, fmt.Errorf("bitfinexclient->GetAllCoins fail, err=%v or resp=nil", err)
	}

	var r [][]string
	err = json.Unmarshal(resp, &r)
	if err != nil || len(r) == 0 {
		return nil, fmt.Errorf("bitfinexclient->GetAllCoins Unmarshal fail or no data, err=%v", err)
	}

	return r[0], nil
}

/*
[
	["1INCH", "AAA", "AAVE", "AAVEF0", "ADA", "ADAF0", "AIX", "ALBT", "ALG", "ALGF0", "AMP", "AMPF0", "ANC", "ANT", "APE", "APEF0", "APENFT", "APT", "APTF0", "ATLAS", "ATO",
"ATOF0", "AVAX", "AVAXC", "AVAXF0", "AXS", "AXSF0", "B2M", "BAL", "BAND", "BAT", "BBB", "BCH", "BCHABC", "BCHN", "BEST", "BFT", "BFX", "BG1", "BG2", "BMI", "BMN", "BNT", "BOBA",
"BOO", "BOSON", "BSV", "BTC", "BTCDOMF0", "BTCF0", "BTG", "BTSE", "BTT", "CAD", "CCD", "CEL", "CHEX", "CHF", "CHSB", "CHZ", "CLO", "CNH", "CNHT", "COMP", "COMPF0", "CONV", "CRV",
"CRVF0", "DAI", "DGB", "DOGE", "DOGEF0", "DORA", "DOT", "DOTF0", "DRK", "DSH", "DUSK", "DVF", "EDO", "EGLD", "EGLDF0", "ENJ", "EOS", "EOSF0", "ETC", "ETCF0", "ETH", "ETH2P", "ETH2R",
"ETH2X", "ETHF0", "ETHS", "ETHW", "ETP", "EUR", "EURF0", "EUS", "EUT", "EUTF0", "EXRD", "FBT", "FCL", "FET", "FIL", "FILF0", "FLR", "FORTH", "FTM", "FTMF0", "FUN", "GALA", "GALAF0",
"GBP", "GBPF0", "GNO", "GNT", "GRT", "GTX", "GXT", "HEC", "HIX", "HKD", "HMT", "HMTPLG", "HTX", "ICE", "ICP", "ICPF0", "IDX", "IOT", "IOTF0", "IQX", "JASMY", "JASMYF0", "JPY", "JPYF0",
"JST", "KAI", "KAN", "KNC", "KNCF0", "KSM", "LBT", "LEO", "LES", "LET", "LINK", "LINKF0", "LNX", "LRC", "LTC", "LTCF0", "LUNA", "LUNA2", "LUNAF0", "LUXO", "LYM", "MATIC", "MATICF0",
"MATICM", "MIM", "MIR", "MKR", "MKRF0", "MLN", "MNA", "MOB", "MXNT", "NEAR", "NEARF0", "NEO", "NEOF0", "NEOGAS", "NEXO", "OCEAN", "OGN", "OMG", "OMGF0", "OMN", "ONE", "OXY", "PAS",
"PAX", "PLANETS", "PLU", "PNG", "PNK", "POLC", "POLIS", "QRDO", "QTF", "QTM", "RBT", "REEF", "REP", "REQ", "RLY", "ROSE", "RRT", "SAND", "SANDF0", "SENATE", "SGB", "SGD", "SHFT",
"SHFTM", "SHIB", "SHIBF0", "SIDUS", "SMR", "SNT", "SNX", "SOL", "SOLF0", "SPELL", "SRM", "STG", "STGF0", "STJ", "SUKU", "SUN", "SUSHI", "SUSHIF0", "SWEAT", "SXX", "TERRAUST",
"TESTBTC", "TESTBTCF0", "TESTUSD", "TESTUSDT", "TESTUSDTF0", "THB", "THETA", "TLOS", "TRADE", "TREEB", "TRX", "TRXF0", "TRY", "TSD", "TWD", "UDC", "UNI", "UNIF0", "UOS", "USD",
"UST", "USTF0", "UTK", "VEE", "VELO", "VET", "VRA", "VSY", "WAVES", "WAVESF0", "WAX", "WBT", "WILD", "WNCG", "WOO", "XAGF0", "XAUT", "XAUTF0", "XCAD", "XCN", "XDC", "XLM", "XLMF0",
"XMR", "XMRF0", "XRA", "XRD", "XRP", "XRPF0", "XTZ", "XTZF0", "XVG", "YFI", "ZCN", "ZEC", "ZECF0", "ZIL", "ZMT", "ZRX"]
]

*/
