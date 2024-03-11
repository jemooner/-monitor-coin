package commonlib

type ErrCode struct {
	Code int
	Err  string
}

var (
	Success     = ErrCode{Code: 0, Err: "success"}                // 成功
	ErrUnknown  = ErrCode{Code: 1, Err: "unknown error"}          // 未知异常
	ErrFallback = ErrCode{Code: 2, Err: "service is unavailable"} // 降级

	// 参数类异常
	ErrApiNotFound = ErrCode{Code: 10000, Err: "api not found"}
	ErrParam       = ErrCode{Code: 10001, Err: "invalid param"}

	// 基础依赖类异常
	ErrQueryDB  = ErrCode{Code: 11001, Err: "query db fail"}
	ErrInsertDB = ErrCode{Code: 11002, Err: "insert db fail"}
	ErrUpdateDB = ErrCode{Code: 11003, Err: "update db fail"}

	// 业务逻辑
	ErrInvalidBalance     = ErrCode{Code: 12005, Err: "insufficient balance"}
	ErrNoAccount          = ErrCode{Code: 12006, Err: "account not exist"}
	ErrFreezeBalance      = ErrCode{Code: 12007, Err: "freeze balance fail"}
	ErrReceiptOrder       = ErrCode{Code: 12008, Err: "receipt order fail"}
	ErrCancelOrder        = ErrCode{Code: 12009, Err: "cancel order fail"}
	ErrConfirmOrder       = ErrCode{Code: 12010, Err: "confirm fail"}
	ErrTransfer           = ErrCode{Code: 12011, Err: "transfer fail"}
	ErrCollect            = ErrCode{Code: 12012, Err: "collect fail"}
	ErrInvalidStatus      = ErrCode{Code: 12013, Err: "invalid status"}
	ErrPay                = ErrCode{Code: 12014, Err: "pay fail"}
	ErrReward             = ErrCode{Code: 12015, Err: "reward fail"}
	ErrInvalidRule        = ErrCode{Code: 12016, Err: "invalid rules"}
	ErrHasProcessingEvent = ErrCode{Code: 12017, Err: "there is a processing event"}
	ErrEnoughGas          = ErrCode{Code: 12018, Err: "no need to transfer gas"}
	ErrPledge             = ErrCode{Code: 12019, Err: "operate pledge fail"}
	ErrRefund             = ErrCode{Code: 12020, Err: "refund fail"}
	ErrInvalidPayChannel  = ErrCode{Code: 12021, Err: "invalid payChannel"}

	// 数据类异常
	ErrNoData = ErrCode{Code: 13000, Err: "no data"}
)
