package go_http

import (
	"context"
	"crypto/tls"
	"golang.org/x/net/publicsuffix"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type RequestOptions struct {
	Ajax               bool
	Cookies            string
	RequestTimeout     time.Duration
	InsecureSkipVerify bool
	DialTimeout        time.Duration
	Headers            map[string]string
	Host               string
	UserAgent          string
	ContentType        string
	ReqData            string
	UrlStr             string
	Method             string
}

func setHeader(req *http.Request, r *RequestOptions) {
	for k, v := range r.Headers {
		req.Header.Set(k, v)
	}
	if r.Ajax {
		req.Header.Set("X-Requested-With", "XMLHttpRequest")
	}
	if r.Host != "" {
		req.Host = r.Host
	}
	if r.ContentType != "" {
		req.Header.Set("Content-Type", r.ContentType)
	}
	req.Header.Set("UserAgent", "GoHttp/1.1")
	if r.UserAgent != "" {
		req.Header.Set("Content-Type", r.ContentType)
	}
	return
}

func createTls(r *RequestOptions) *http.Transport {
	if r.DialTimeout == 0 {
		r.DialTimeout = 5 * time.Second
	}
	tr := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       5 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: r.InsecureSkipVerify},
		DisableKeepAlives:     true,
	}
	tr.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		conn, e = net.DialTimeout(network, addr, r.DialTimeout)
		if e != nil {
			return nil, e
		}
		return
	}
	return tr
}

func client(r *RequestOptions, cJar http.CookieJar) (c *http.Client) {
	if r.RequestTimeout == 0 {
		r.RequestTimeout = 3 * time.Second
	}
	if cJar == nil {
		cJar, _ = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	}
	c = &http.Client{
		Jar:       cJar,
		Transport: createTls(r),
		Timeout:   r.RequestTimeout,
	}
	return
}

func newRequest(r *RequestOptions, cJar http.CookieJar) (*Response, error) {
	req, err := http.NewRequest(r.Method, r.UrlStr, strings.NewReader(r.ReqData))
	if err != nil {
		return nil, err
	}
	setHeader(req, r)
	c := client(r, cJar)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{Resp: resp}, nil
}
