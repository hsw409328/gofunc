package go_http

import (
	"fmt"
	"testing"
)

func TestHttpPost(t *testing.T) {
	a, err := HttpPost(&RequestOptions{
		UrlStr: "https://www.baidu.com",
		Ajax:   true,
	}, nil)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(a.GetBodyString())
}

func TestHttpGet(t *testing.T) {
	a, err := HttpGet(&RequestOptions{
		UrlStr: "http://www.baidu.com",
		Ajax:   true,
	}, nil)
	if err!=nil{
		t.Error(err)
	}
	fmt.Println(a.GetBodyString())
}
