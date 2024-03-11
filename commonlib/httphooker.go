package commonlib

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"monitor-coin/commonlib/dlog"
)

/*
*
http logging
*/
type httpHandler func(ctx context.Context, params string) []byte

func Wrapper(handle httpHandler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		st, tid := time.Now(), GetTraceId(r)
		ctx := SetTraceId(r.Context(), tid)
		ctx = SetSpanId(ctx, GetSpanId())
		ctx = SetMethod(ctx, r.Method)
		trc := GetTrace(ctx)

		errRb := []byte(fmt.Sprintf(`{"code":1, "msg":"system upgrade", "data":{}, "traceid":"%v"}`, tid))
		defer func(r []byte) {
			if rcv := recover(); rcv != nil {
				_, err := w.Write(r)
				dlog.Errorf("%v||httphooker.wrapper panic, recover=%+v,err=%v", trc, TracePanic(rcv), err)
			}
		}(errRb)

		s, err := ioutil.ReadAll(r.Body)
		str := string(s)

		dlog.Infof(HttpInLog(ctx, r, "", str))
		rb := handle(ctx, string(s))

		_, err = w.Write(rb)
		if err != nil {
			dlog.Infof(`%v||httphooker.wrapper->w.Write fail,err=%+v`, trc, err)
		}

		dlog.Infof(HttpOutLog(ctx, r, "", rb, st))
	}
}

type HttpResp struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceId interface{} `json:"traceid"`
}

func FormatResp(ctx context.Context, ec ErrCode, data interface{}) []byte {
	msg := Success.Err
	if ec.Code != Success.Code {
		msg = ErrUnknown.Err
		if len(ec.Err) != 0 {
			msg = ec.Err
		} else if data != nil {
			msg = fmt.Sprint(data)
		}
		data = nil
	}

	resp := HttpResp{
		Code:    ec.Code,
		Msg:     msg,
		Data:    data,
		TraceId: ctx.Value(trcid),
	}

	rb, err := json.Marshal(resp)
	if err != nil {
		rb = []byte(fmt.Sprintf(`{"code":1, "msg":"system upgrade", "data":{}, "traceid":"%v"}`, ctx.Value(trcid)))
	}
	return rb
}
