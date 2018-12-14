package go_http

import (
	"net/http"
)

func HttpPost(r *RequestOptions, c http.CookieJar) (*Response, error) {
	r.Method = "POST"
	return newRequest(r, c)
}

func HttpGet(r *RequestOptions, c http.CookieJar) (*Response, error) {
	r.Method = "GET"
	return newRequest(r, c)
}
