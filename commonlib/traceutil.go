package commonlib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

const (
	trcid  = "traceid"
	spid   = "spanid"
	method = "method"
	code   = "code"
)

func GetTraceId(r *http.Request) string {
	if r != nil {
		traceId := r.Header.Get("X-Trace-ID")
		if len(traceId) != 0 {
			return traceId
		}
	}

	now := time.Now()
	return now.Format("060102150405X") + fmt.Sprint((now.UnixNano()/1000)%10000000)
}

func GenTraceId(ctx context.Context, r *http.Request) string {
	if c := ctx.Value(trcid); c != nil {
		return fmt.Sprintf("%v", c)
	}

	if r != nil {
		traceId := r.Header.Get("X-Trace-ID")
		if len(traceId) != 0 {
			return traceId
		}
	}

	now := time.Now()
	return now.Format("060102150405X") + fmt.Sprint((now.UnixNano()/1000)%10000000)
}

func GetSpanId() string {
	return fmt.Sprint((time.Now().UnixNano() / 1000) % 10000000)
}

func GetErrCode(ctx context.Context) string {
	if c := ctx.Value(code); c != nil {
		return fmt.Sprintf("%v", c)
	}
	return "1"
}

func GetTrace(ctx context.Context) string {
	return fmt.Sprintf("traceid=%v||spanid=%v", ctx.Value(trcid), ctx.Value(spid))
}

func TracePanic(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)

	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	return buf.String()
}

func SetCtx(ctx context.Context, k interface{}, v interface{}) context.Context {
	return context.WithValue(ctx, k, v)
}

func SetTraceId(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, trcid, v)
}

func SetSpanId(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, spid, v)
}

func SetMethod(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, method, v)
}

func SetLang(ctx context.Context, v interface{}) context.Context {
	return SetCtx(ctx, method, v)
}

func HttpInLog(ctx context.Context, r *http.Request, svc string, p interface{}) string {
	f := `%v||service=%s||tag=_http_request_in||ip=%v||uri=%v||method=%v||%+v`
	return fmt.Sprintf(f, GetTrace(ctx), svc, GetClientIp(r), r.URL.Path, r.Method, p)
}

func HttpOutLog(ctx context.Context, r *http.Request, svc string, resp []byte, t time.Time) string {
	c := fmt.Sprint(ErrUnknown.Code)
	var m map[string]interface{}
	if len(resp) > 0 {
		json.Unmarshal(resp, &m)
	}
	if v, ok := m[code]; ok {
		c = fmt.Sprint(v)
	}

	l := len(resp)
	if l > 1000 {
		l = 1000
	}
	f := `%v||service=%s||tag=_http_request_out||code=%v||uri=%v||latency=%dms||%v`
	return fmt.Sprintf(f, GetTrace(ctx), svc, c, r.URL.Path, time.Since(t).Milliseconds(), string(resp[0:l]))
}

func CallInLog(ctx context.Context, callee string, m string, ep string, p interface{}) string {
	f := `%v||tag=_call_in_%s||%s endpoint=%s||%+v`
	return fmt.Sprintf(f, GetTrace(ctx), callee, m, ep, p)
}

func CallOutLog(ctx context.Context, callee string, resp []byte, err error, t time.Time) string {
	l := len(resp)
	if l > 500 {
		l = 500
	}
	f := `%v||tag=_call_out_%s||latency=%dms||resp=%s||err=%v`
	return fmt.Sprintf(f, GetTrace(ctx), callee, time.Since(t).Milliseconds(), resp[0:l], err)
}

/**
获取上游ip
*/
func GetClientIp(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	} else {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		if ip == `::1` {
			return "127.0.0.1"
		}
		return ip
	}
}
