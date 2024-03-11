package commonlib

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

func Ping(addrs []string) ([]string, error) {
	if len(addrs) == 0 {
		return addrs, fmt.Errorf("httputil.Ping fail:addrs=null")
	}
	var reAddrIPv4Port = regexp.MustCompile(`^\d+\.\d+\.\d+\.\d+(:\d+)$`)

	for k, addr := range addrs {
		if reAddrIPv4Port.MatchString(addr) {
			addr = "http://" + addr
		}

		u, err := url.Parse(addr)

		if err != nil {
			return nil, fmt.Errorf(`invalid addr "%v" with error: %v`, addr, err)
		}

		if u.Scheme != "http" && u.Scheme != "https" {
			return nil, fmt.Errorf(`invalid addr "%v" with unsupported scheme "%v"`, addr, u.Scheme)
		}
		addrs[k] = addr
	}

	return addrs, nil
}

// 定义全局复用
// 避免每个client定义一个导致 socket: too many open files
// 参考https://www.jianshu.com/p/50ed36e98459
var defaultTransport *http.Transport

func initDefaultTransport() *http.Transport {
	defaultTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		DialTLSContext:         nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     false,
		MaxIdleConns:           500,
		MaxIdleConnsPerHost:    100,
		MaxConnsPerHost:        100,
		IdleConnTimeout:        time.Duration(60) * time.Second,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  1 * time.Second,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		GetProxyConnectHeader:  nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      true,
	}
	return defaultTransport
}

func GetDefaultHttpClient(timeout int) *http.Client {
	if defaultTransport == nil {
		defaultTransport = initDefaultTransport()
	}
	return &http.Client{
		Transport:     defaultTransport,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Duration(timeout) * time.Second, // 整体超时时间
	}
}
